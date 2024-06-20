// (c) Copyright 2016 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type usesWeakCryptographyHash struct {
	issue.MetaData
	blocklist map[string][]string
}

func (r *usesWeakCryptographyHash) ID() string {
	return r.MetaData.ID
}

func (r *usesWeakCryptographyHash) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	for pkg, funcs := range r.blocklist {
		if _, matched := gosec.MatchCallByPackage(n, c, pkg, funcs...); matched {
			return c.NewIssue(n, r.ID(), r.What, r.Severity, r.Confidence), nil
		}
	}
	return nil, nil
}

// NewUsesWeakCryptographyHash detects uses of md5.*, sha1.*
func NewUsesWeakCryptographyHash(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	calls := make(map[string][]string)
	calls["crypto/md5"] = []string{"New", "Sum"}
	calls["crypto/sha1"] = []string{"New", "Sum"}
	rule := &usesWeakCryptographyHash{
		blocklist: calls,
		MetaData: issue.MetaData{
			ID:         id,
			Severity:   issue.Medium,
			Confidence: issue.High,
			What:       "Use of weak cryptographic primitive",
		},
	}
	return rule, []ast.Node{(*ast.CallExpr)(nil)}
}