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

type GithubPrivateRelease struct {
	Owner   string `json:"owner"`
	Repo    string `json:"repo"`
	Release string `json:"release"`
	Asset   string `json:"asset"`
}

func (g GithubPrivateRelease) Auth() {
	// init github client with personal access token from env var
}

func (g GithubPrivateRelease) Download() {
	// get repo info by release tag
	// get asset id for provided asset from repo info
	// download release asset by id
}
