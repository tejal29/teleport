/*
Copyright 2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sftp

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/pkg/sftp"

	"github.com/gravitational/trace"
)

// remoteFS provides API for accessing the files on
// the local file system
type remoteFS struct {
	ctx context.Context
	c   *sftp.Client
}

func (r *remoteFS) SetContext(ctx context.Context) {
	r.ctx = ctx
}

func (r *remoteFS) Type() string {
	return "remote"
}

func (r *remoteFS) Stat(path string) (os.FileInfo, error) {
	if err := r.ctx.Err(); err != nil {
		return nil, err
	}

	fi, err := r.c.Stat(path)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return fi, nil
}

func (r *remoteFS) ReadDir(path string) ([]os.FileInfo, error) {
	if err := r.ctx.Err(); err != nil {
		return nil, err
	}

	fileInfos, err := r.c.ReadDir(path)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return fileInfos, nil
}

func (r *remoteFS) Open(path string) (io.ReadCloser, error) {
	if err := r.ctx.Err(); err != nil {
		return nil, err
	}

	f, err := r.c.Open(path)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return f, nil
}

func (r *remoteFS) Create(path string) (io.WriteCloser, error) {
	if err := r.ctx.Err(); err != nil {
		return nil, err
	}

	f, err := r.c.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return f, nil
}

func (r *remoteFS) Mkdir(path string) error {
	if err := r.ctx.Err(); err != nil {
		return err
	}

	err := r.c.MkdirAll(path)
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}

func (r *remoteFS) Chmod(path string, mode os.FileMode) error {
	if err := r.ctx.Err(); err != nil {
		return err
	}

	return trace.Wrap(r.c.Chmod(path, mode))
}

func (r *remoteFS) Chtimes(path string, atime, mtime time.Time) error {
	if err := r.ctx.Err(); err != nil {
		return err
	}

	return trace.Wrap(r.c.Chtimes(path, atime, mtime))
}
