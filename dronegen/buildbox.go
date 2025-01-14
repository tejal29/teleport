// Copyright 2021 Gravitational, Inc
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

import "fmt"

func allBuildboxPipelineSteps() []step {
	steps := []step{
		{
			Name:  "Check out code",
			Image: "docker:git",
			Commands: []string{
				`git clone --depth 1 --single-branch --branch ${DRONE_SOURCE_BRANCH:-master} https://github.com/gravitational/${DRONE_REPO_NAME}.git .`,
				`git checkout ${DRONE_COMMIT}`,
			},
		},
		waitForDockerStep(),
	}

	for _, name := range []string{"buildbox", "buildbox-arm", "buildbox-centos7"} {
		for _, fips := range []bool{false, true} {
			// FIPS is unsupported on ARM/ARM64
			if name == "buildbox-arm" && fips {
				continue
			}
			steps = append(steps, buildboxPipelineSteps(name, fips)...)
		}
	}
	return steps
}

func buildboxPipelineSteps(buildboxName string, fips bool) []step {
	if fips {
		buildboxName += "-fips"
	}
	assumeStagingRoleStep := kubernetesAssumeAwsRoleStep(kubernetesRoleSettings{
		awsRoleSettings: awsRoleSettings{
			awsAccessKeyID:     value{fromSecret: "STAGING_BUILDBOX_DRONE_USER_ECR_KEY"},
			awsSecretAccessKey: value{fromSecret: "STAGING_BUILDBOX_DRONE_USER_ECR_SECRET"},
			role:               value{fromSecret: "STAGING_BUILDBOX_DRONE_ECR_AWS_ROLE"},
		},
		configVolume: volumeRefAwsConfig,
	})
	assumeStagingRoleStep.Name = fmt.Sprintf("Assume Staging %s AWS Role", buildboxName)
	assumeProductionRoleStep := kubernetesAssumeAwsRoleStep(kubernetesRoleSettings{
		awsRoleSettings: awsRoleSettings{
			awsAccessKeyID:     value{fromSecret: "PRODUCTION_BUILDBOX_DRONE_USER_ECR_KEY"},
			awsSecretAccessKey: value{fromSecret: "PRODUCTION_BUILDBOX_DRONE_USER_ECR_SECRET"},
			role:               value{fromSecret: "PRODUCTION_BUILDBOX_DRONE_ECR_AWS_ROLE"},
		},
		configVolume: volumeRefAwsConfig,
	})
	assumeProductionRoleStep.Name = fmt.Sprintf("Assume Production %s AWS Role", buildboxName)
	return []step{
		assumeStagingRoleStep,
		step{
			Name:    fmt.Sprintf("Build %s and push to Staging", buildboxName),
			Image:   "docker",
			Volumes: []volumeRef{volumeRefDocker, volumeRefAwsConfig},
			Commands: []string{
				`apk add --no-cache make aws-cli`,
				`chown -R $UID:$GID /go`,
				// Authenticate to staging registry
				`aws ecr get-login-password --region=us-west-2 | docker login -u="AWS" --password-stdin ` + StagingRegistry,
				// Build buildbox image
				fmt.Sprintf(`make -C build.assets %s`, buildboxName),
				// Retag for staging registry
				fmt.Sprintf(`docker tag %s/gravitational/teleport-%s:$BUILDBOX_VERSION %s/gravitational/teleport-%s:$BUILDBOX_VERSION-$DRONE_COMMIT_SHA`, ProductionRegistry, buildboxName, StagingRegistry, buildboxName),
				// Push to staging registry
				fmt.Sprintf(`docker push %s/gravitational/teleport-%s:$BUILDBOX_VERSION-$DRONE_COMMIT_SHA`, StagingRegistry, buildboxName),
			},
		},
		assumeProductionRoleStep,
		step{
			Name:    fmt.Sprintf("Push %s to Production", buildboxName),
			Image:   "docker",
			Volumes: []volumeRef{volumeRefDocker, volumeRefAwsConfig},
			Commands: []string{
				`apk add --no-cache make aws-cli`,
				`chown -R $UID:$GID /go`,
				// Authenticate to production registry
				`aws ecr-public get-login-password --region=us-east-1 | docker login -u="AWS" --password-stdin ` + ProductionRegistry,
				// Push to production registry
				fmt.Sprintf(`docker push %s/gravitational/teleport-%s:$BUILDBOX_VERSION`, ProductionRegistry, buildboxName),
			},
		},
	}
}

func buildboxPipeline() pipeline {
	p := newKubePipeline("build-buildboxes")
	p.Environment = map[string]value{
		"BUILDBOX_VERSION": buildboxVersion,
		"UID":              {raw: "1000"},
		"GID":              {raw: "1000"},
	}

	// only on master for now; add the release branch name when forking a new release series.
	p.Trigger = pushTriggerForBranch("master", "branch/*")
	p.Workspace = workspace{Path: "/go/src/github.com/gravitational/teleport"}
	p.Volumes = []volume{volumeDocker, volumeAwsConfig}
	p.Services = []service{
		dockerService(),
	}
	p.Steps = allBuildboxPipelineSteps()
	return p
}
