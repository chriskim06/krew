// Copyright 2020 The Kubernetes Authors.
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

package index

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"sigs.k8s.io/krew/pkg/index/strategies"
)

// DownloadStrategy defines the behavior various plugin sources should implement
type DownloadStrategy interface {
	Auth()
	Download()
}

// Strategy is a concrete type that Plugin contains. This implements the json.Unmarshaler
// interface so that multiple DownloadStrategy types can be unmarshaled. Any new DownloadStrategy
// should be added to the UnmarshalJSON method below.
type Strategy struct {
	strategy DownloadStrategy
}

func (r Strategy) Print() {
	fmt.Println(r.strategy)
	j, err := json.Marshal(r.strategy)
	if err != nil {
		return
	}
	fmt.Println(string(j))
}

// UnmarshalJSON is the custom unmarshaler used to unmarshal a plugin DownloadStrategy to the
// correct type. Any new DownloadStrategy should be added here. The type dictates which
// DownloadStrategy we try to unmarshal the json.RawMessage to.
func (r *Strategy) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	var re struct {
		Type  string          `json:"type"`
		Stuff json.RawMessage `json:"stuff"`
	}
	if err := json.Unmarshal(data, &re); err != nil {
		return err
	}
	var release Strategy
	switch re.Type {
	case "github":
		var g strategies.GithubPrivateRelease
		if err := json.Unmarshal(re.Stuff, &g); err != nil {
			return err
		}
		release.strategy = g
	default:
		return errors.New("unknown type")
	}
	*r = release
	return nil
}
