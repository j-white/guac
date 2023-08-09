//
// Copyright 2023 The GUAC Authors.
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

package gremlin

import (
	"encoding/json"
	"github.com/guacsec/guac/internal/testing/ptrfrom"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

func convNode(n model.Node) hasID {
	// All nodes have a json "id"
	// Only getting top-level id however
	var h hasID
	b, _ := json.Marshal(n)
	_ = json.Unmarshal(b, &h)
	return h
}

func convNodes(ns []model.Node) []string {
	var ids []string
	for _, n := range ns {
		h := convNode(n)
		ids = append(ids, h.ID)
	}
	return ids
}

type hasID struct {
	ID string `json:"id"`
}

var a1 = &model.ArtifactInputSpec{
	Algorithm: "sha256",
	Digest:    "6bbb0da1891646e58eb3e6a63af3a6fc3c8eb5a0d44824cba581d2e14a0450cf",
}
var a1out = &model.Artifact{
	Algorithm: "sha256",
	Digest:    "6bbb0da1891646e58eb3e6a63af3a6fc3c8eb5a0d44824cba581d2e14a0450cf",
}

var a2 = &model.ArtifactInputSpec{
	Algorithm: "sha1",
	Digest:    "7A8F47318E4676DACB0142AFA0B83029CD7BEFD9",
}
var a2out = &model.Artifact{
	Algorithm: "sha1",
	Digest:    "7a8f47318e4676dacb0142afa0b83029cd7befd9",
}

var a3 = &model.ArtifactInputSpec{
	Algorithm: "sha512",
	Digest:    "374AB8F711235830769AA5F0B31CE9B72C5670074B34CB302CDAFE3B606233EE92EE01E298E5701F15CC7087714CD9ABD7DDB838A6E1206B3642DE16D9FC9DD7",
}
var a3out = &model.Artifact{
	Algorithm: "sha512",
	Digest:    "374ab8f711235830769aa5f0b31ce9b72c5670074b34cb302cdafe3b606233ee92ee01e298e5701f15cc7087714cd9abd7ddb838a6e1206b3642de16d9fc9dd7",
}

var a4 = &model.ArtifactInputSpec{
	Algorithm: "sha1",
	Digest:    "5a787865sd676dacb0142afa0b83029cd7befd9",
}

var p1 = &model.PkgInputSpec{
	Type: "pypi",
	Name: "tensorflow",
}
var p1out = &model.Package{
	Type: "pypi",
	Namespaces: []*model.PackageNamespace{{
		Names: []*model.PackageName{{
			Name: "tensorflow",
			Versions: []*model.PackageVersion{{
				Version:    "",
				Qualifiers: []*model.PackageQualifier{},
			}},
		}},
	}},
}
var p1outName = &model.Package{
	Type: "pypi",
	Namespaces: []*model.PackageNamespace{{
		Names: []*model.PackageName{{
			Name:     "tensorflow",
			Versions: []*model.PackageVersion{{Qualifiers: []*model.PackageQualifier{}}},
		}},
	}},
}

var p2 = &model.PkgInputSpec{
	Type:    "pypi",
	Name:    "tensorflow",
	Version: ptrfrom.String("2.11.1"),
}
var p2out = &model.Package{
	Type: "pypi",
	Namespaces: []*model.PackageNamespace{{
		Names: []*model.PackageName{{
			Name: "tensorflow",
			Versions: []*model.PackageVersion{{
				Version:    "2.11.1",
				Qualifiers: []*model.PackageQualifier{},
			}},
		}},
	}},
}
var p2outName = &model.Package{
	Type: "pypi",
	Namespaces: []*model.PackageNamespace{{
		Names: []*model.PackageName{{
			Name:     "tensorflow",
			Versions: []*model.PackageVersion{{Qualifiers: []*model.PackageQualifier{}}},
		}},
	}},
}

var p3 = &model.PkgInputSpec{
	Type:    "pypi",
	Name:    "tensorflow",
	Version: ptrfrom.String("2.11.1"),
	Subpath: ptrfrom.String("saved_model_cli.py"),
}
var p3out = &model.Package{
	Type: "pypi",
	Namespaces: []*model.PackageNamespace{{
		Names: []*model.PackageName{{
			Name: "tensorflow",
			Versions: []*model.PackageVersion{{
				Version:    "2.11.1",
				Subpath:    "saved_model_cli.py",
				Qualifiers: []*model.PackageQualifier{},
			}},
		}},
	}},
}

var p4 = &model.PkgInputSpec{
	Type:      "conan",
	Namespace: ptrfrom.String("openssl.org"),
	Name:      "openssl",
	Version:   ptrfrom.String("3.0.3"),
}
var p4outName = &model.Package{
	Type: "conan",
	Namespaces: []*model.PackageNamespace{{
		Namespace: "openssl.org",
		Names: []*model.PackageName{{
			Name:     "openssl",
			Versions: []*model.PackageVersion{},
		}},
	}},
}

var s1 = &model.SourceInputSpec{
	Type:      "git",
	Namespace: "github.com/jeff",
	Name:      "myrepo",
}
var s1out = &model.Source{
	Type: "git",
	Namespaces: []*model.SourceNamespace{{
		Namespace: "github.com/jeff",
		Names: []*model.SourceName{{
			Name:   "myrepo",
			Tag:    ptrfrom.String(""),
			Commit: ptrfrom.String(""),
		}},
	}},
}

var s2 = &model.SourceInputSpec{
	Type:      "git",
	Namespace: "github.com/bob",
	Name:      "bobsrepo",
}
var s2out = &model.Source{
	Type: "git",
	Namespaces: []*model.SourceNamespace{{
		Namespace: "github.com/bob",
		Names: []*model.SourceName{{
			Name:   "bobsrepo",
			Tag:    ptrfrom.String(""),
			Commit: ptrfrom.String(""),
		}},
	}},
}
