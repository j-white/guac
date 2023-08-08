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
	"context"
	"fmt"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"sort"
)

const (
	Package Label = "package"
)

// types copied from arrangodb impl.
type dbPkgVersion struct {
	TypeID        string   `json:"type_id"`
	PkgType       string   `json:"type"`
	NamespaceID   string   `json:"namespace_id"`
	Namespace     string   `json:"namespace"`
	NameID        string   `json:"name_id"`
	Name          string   `json:"name"`
	VersionID     string   `json:"version_id"`
	Version       string   `json:"version"`
	Subpath       string   `json:"subpath"`
	QualifierList []string `json:"qualifier_list"`
}

type dbPkgName struct {
	TypeID      string `json:"type_id"`
	PkgType     string `json:"type"`
	NamespaceID string `json:"namespace_id"`
	Namespace   string `json:"namespace"`
	NameID      string `json:"name_id"`
	Name        string `json:"name"`
}

type dbPkgNamespace struct {
	TypeID      string `json:"type_id"`
	PkgType     string `json:"type"`
	NamespaceID string `json:"namespace_id"`
	Namespace   string `json:"namespace"`
}

type dbPkgType struct {
	TypeID  string `json:"type_id"`
	PkgType string `json:"type"`
}

type PkgIds struct {
	TypeId      string
	NamespaceId string
	NameId      string
	VersionId   string
}

// copied from arrangodb
func guacPkgId(pkg model.PkgInputSpec) PkgIds {
	ids := PkgIds{}

	ids.TypeId = pkg.Type

	var ns string
	if pkg.Namespace != nil {
		if *pkg.Namespace != "" {
			ns = *pkg.Namespace
		} else {
			ns = guacEmpty
		}
	}
	ids.NamespaceId = fmt.Sprintf("%s::%s", ids.TypeId, ns)
	ids.NameId = fmt.Sprintf("%s::%s", ids.NamespaceId, pkg.Name)

	var version string
	if pkg.Version != nil {
		if *pkg.Version != "" {
			version = *pkg.Version
		} else {
			version = guacEmpty
		}
	}

	var subpath string
	if pkg.Subpath != nil {
		if *pkg.Subpath != "" {
			subpath = *pkg.Subpath
		} else {
			subpath = guacEmpty
		}
	}

	ids.VersionId = fmt.Sprintf("%s::%s::%s?", ids.NameId, version, subpath)

	qualifiersMap := map[string]string{}
	var keys []string
	for _, kv := range pkg.Qualifiers {
		qualifiersMap[kv.Key] = kv.Value
		keys = append(keys, kv.Key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		ids.VersionId += fmt.Sprintf("%s=%s&", k, qualifiersMap[k])
	}

	return ids
}

// copied from arrangodb
func getPackageQueryValues(pkg *model.PkgInputSpec) map[interface{}]interface{} {
	values := make(map[interface{}]interface{})
	values[gremlingo.T.Label] = string(Package)

	// FIXME: Add relation to pkg type
	//// add guac keys
	//values["typeID"] = c.pkgTypeMap[pkg.Type].Id
	//values["typeKey"] = c.pkgTypeMap[pkg.Type].Key
	//values["typeValue"] = c.pkgTypeMap[pkg.Type].PkgType

	guacIds := guacPkgId(*pkg)
	values["guacNsKey"] = guacIds.NamespaceId
	values["guacNameKey"] = guacIds.NameId
	values["guacVersionKey"] = guacIds.VersionId

	values[typeStr] = pkg.Type

	values[name] = pkg.Name
	if pkg.Namespace != nil {
		values[namespace] = *pkg.Namespace
	} else {
		values[namespace] = ""
	}
	if pkg.Version != nil {
		values[version] = *pkg.Version
	} else {
		values[version] = ""
	}
	if pkg.Subpath != nil {
		values[subpath] = *pkg.Subpath
	} else {
		values[subpath] = ""
	}

	// To ensure consistency, always sort the qualifiers by key
	qualifiersMap := map[string]string{}
	var keys []string
	for _, kv := range pkg.Qualifiers {
		qualifiersMap[kv.Key] = kv.Value
		keys = append(keys, kv.Key)
	}
	sort.Strings(keys)
	var qualifiers []string
	for _, k := range keys {
		qualifiers = append(qualifiers, k, qualifiersMap[k])
	}
	storeArrayInVertexProperties2(values, "qualifier", qualifiers)
	return values
}

func getPackageObject(id string, values map[interface{}]interface{}) *model.Package {
	pkgVersion := &model.PackageVersion{
		Version:    values[version].(string),
		Subpath:    values[subpath].(string),
		Qualifiers: []*model.PackageQualifier{},
	}
	pkgName := &model.PackageName{
		Name:     values[name].(string),
		Versions: []*model.PackageVersion{pkgVersion},
	}
	pkgNamespace := &model.PackageNamespace{
		Namespace: values[namespace].(string),
		Names:     []*model.PackageName{pkgName},
	}
	return &model.Package{
		ID:         id,
		Type:       values[typeStr].(string),
		Namespaces: []*model.PackageNamespace{pkgNamespace},
	}
}

func (c *gremlinClient) IngestPackage(ctx context.Context, pkg model.PkgInputSpec) (*model.Package, error) {
	return ingestModelObject[*model.PkgInputSpec, *model.Package](c, &pkg, getPackageQueryValues, getPackageObject)
}

func (c *gremlinClient) IngestPackages(ctx context.Context, pkgs []*model.PkgInputSpec) ([]*model.Package, error) {
	return bulkIngestModelObjects[*model.PkgInputSpec, *model.Package](c, pkgs, getPackageQueryValues, getPackageObject)
}

func (c *gremlinClient) Packages(ctx context.Context, pkgSpec *model.PkgSpec) ([]*model.Package, error) {
	query := createVertexQuery(Package)
	if pkgSpec != nil {
		if pkgSpec.ID != nil {
			query.id = *pkgSpec.ID
		}
		if pkgSpec.Name != nil {
			query.has[name] = *pkgSpec.Name
		}
		if pkgSpec.Type != nil {
			query.has[typeStr] = *pkgSpec.Type
		}
		if pkgSpec.Namespace != nil {
			query.has[namespace] = *pkgSpec.Namespace
		}
		if pkgSpec.Subpath != nil {
			query.has[subpath] = *pkgSpec.Subpath
		}
	}
	return queryModelObjects[*model.Package](c, query, getPackageObject)
}
