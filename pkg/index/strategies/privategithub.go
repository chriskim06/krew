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

import (
	"context"
	"io"

	"github.com/google/go-github/v50/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"sigs.k8s.io/krew/internal/environment"
)

type GithubPrivateRelease struct {
	BaseURL string `json:"baseURL"`
	Owner   string `json:"owner"`
	Repo    string `json:"repo"`
	Release string `json:"release"`
	Asset   string `json:"asset"`

	client *github.Client
}

func (g GithubPrivateRelease) Name() string {
	return "GithubPrivateRelease"
}

func (g GithubPrivateRelease) Get(_ string) (io.ReadCloser, error) {
	// do the stuff to get the archive from the private repo
	if err := g.init(); err != nil {
		return nil, err
	}
	// the go-github library DownloadReleaseAsset function returns an io.ReadCloser
	return g.download()
}

func (g *GithubPrivateRelease) init() error {
	// init github client with personal access token from env var
	conf := environment.GetConfig()
	token := conf.Env["KREW_GITHUB_TOKEN"]
	if token == "" {
		return errors.New("error must set KREW_GITHUB_TOKEN env var when downloading a private asset")
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	g.client = github.NewClient(oauth2.NewClient(context.Background(), tokenSource))
	return nil
}

func (g GithubPrivateRelease) download() (io.ReadCloser, error) {
	// get repo info by release tag
	// get asset id for provided asset from repo info
	// download release asset by id
	release, _, err := g.client.Repositories.GetReleaseByTag(context.Background(), g.Owner, g.Repo, g.Release)
	if err != nil {
		return nil, err
	}
	var id int64
	found := false
	for _, asset := range release.Assets {
		if asset.GetName() == g.Asset {
			found = true
			id = *asset.ID
			break
		}
	}
	if !found {
		return nil, errors.New("error did not find expected asset")
	}

	rc, _, err := g.client.Repositories.DownloadReleaseAsset(context.Background(), g.Owner, g.Repo, id, nil)
	return rc, err
}

func (g GithubPrivateRelease) Verify() error {
	// check sha256sum of downloaded asset
	return nil
}
