// Copyright 2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "path/filepath"

// awsRoleSettings contains the information necessary to assume an AWS Role
//
// This is intended to be imbedded, please use the kubernetes/mac/windows versions
// with their corresponding pipelines.
type awsRoleSettings struct {
	awsAccessKeyID     value
	awsSecretAccessKey value
	role               value
}

// kubernetesRoleSettings contains the info necessary to assume an AWS role and save the credentials to a volume that later steps can use
type kubernetesRoleSettings struct {
	awsRoleSettings
	configVolume volumeRef
}

// macRoleSettings contains the info necessary to assume an AWS role and save the credentials to a path that later steps can use
type macRoleSettings struct {
	awsRoleSettings
	configPath string
}

// kuberentesS3Settings contains all info needed to download from S3 in a kubernetes pipeline
type kubernetesS3Settings struct {
	region       string
	source       string
	target       string
	configVolume volumeRef
}

// assumeRoleCommands is a helper to build the role assumtipn commands on a *nix platform
func assumeRoleCommands(configPath string) []string {
	assumeRoleCmd := `printf "[default]\naws_access_key_id = %s\naws_secret_access_key = %s\naws_session_token = %s" \
  $(aws sts assume-role \
    --role-arn "$AWS_ROLE" \
    --role-session-name $(echo "drone-${DRONE_REPO}-${DRONE_BUILD_NUMBER}" | sed "s|/|-|g") \
    --query "Credentials.[AccessKeyId,SecretAccessKey,SessionToken]" \
    --output text) \
  > ` + configPath
	return []string{
		`aws sts get-caller-identity`, // check the original identity
		assumeRoleCmd,
		`unset AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY`, // remove original identity from environment
		`aws sts get-caller-identity`,                   // check the new assumed identity
	}

}

// kubernetesAssumeAwsRoleStep builds a step to assume an AWS role and save it to a volume that later steps can use
func kubernetesAssumeAwsRoleStep(s kubernetesRoleSettings) step {
	configPath := filepath.Join(s.configVolume.Path, "credentials")
	return step{
		Name:  "Assume AWS Role",
		Image: "amazon/aws-cli",
		Environment: map[string]value{
			"AWS_ACCESS_KEY_ID":     s.awsAccessKeyID,
			"AWS_SECRET_ACCESS_KEY": s.awsSecretAccessKey,
			"AWS_ROLE":              s.role,
		},
		Volumes:  []volumeRef{s.configVolume},
		Commands: assumeRoleCommands(configPath),
	}
}

// macAssumeAwsRoleStep builds a step to assume an AWS role and save it to a host path that later steps can use
func macAssumeAwsRoleStep(s macRoleSettings) step {
	return step{
		Name: "Assume AWS Role",
		Environment: map[string]value{
			"AWS_ACCESS_KEY_ID":           s.awsAccessKeyID,
			"AWS_SECRET_ACCESS_KEY":       s.awsSecretAccessKey,
			"AWS_ROLE":                    s.role,
			"AWS_SHARED_CREDENTIALS_FILE": value{raw: s.configPath},
		},
		Commands: assumeRoleCommands(s.configPath),
	}
}

// kubernetesUploadToS3Step generates an S3 upload step
func kubernetesUploadToS3Step(s kubernetesS3Settings) step {
	return step{
		Name:  "Upload to S3",
		Image: "amazon/aws-cli",
		Environment: map[string]value{
			"AWS_S3_BUCKET": {fromSecret: "AWS_S3_BUCKET"},
			"AWS_REGION":    {raw: s.region},
		},
		Volumes: []volumeRef{s.configVolume},
		Commands: []string{
			`cd ` + s.source,
			`aws s3 sync . s3://$AWS_S3_BUCKET/` + s.target,
		},
	}
}
