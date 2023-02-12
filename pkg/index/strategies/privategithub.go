// Copyright 2019 The Kubernetes Authors.
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

package strategies

import "io"

type GithubPrivateRelease struct {
	Owner   string `json:"owner"`
	Repo    string `json:"repo"`
	Release string `json:"release"`
	Asset   string `json:"asset"`
}

func (g GithubPrivateRelease) Name() string {
	return "GithubPrivateRelease"
}

func (g GithubPrivateRelease) Get(_ string) (io.ReadCloser, error) {
	// do the stuff to get the archive from the private repo
	if err := g.Auth(); err != nil {
		return nil, err
	}
	if err := g.Download(""); err != nil {
		return nil, err
	}
	return nil, nil
}

func (g GithubPrivateRelease) Auth() error {
	// init github client with personal access token from env var
	//     conf := environment.GetConfig()
	//     token := conf.Env["KREW_GITHUB_TOKEN"]
	//     if token == "" {
	//     }
	return nil
}

func (g GithubPrivateRelease) Download(destination string) error {
	// get repo info by release tag
	// get asset id for provided asset from repo info
	// download release asset by id
	return nil
}

func (g GithubPrivateRelease) Verify() error {
	// check sha256sum of downloaded asset
	return nil
}
