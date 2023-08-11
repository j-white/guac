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
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/guacsec/guac/internal/testing/ptrfrom"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

func TestIsDependency(t *testing.T) {
	type call struct {
		P1 *model.PkgInputSpec
		P2 *model.PkgInputSpec
		ID *model.IsDependencyInputSpec
	}
	tests := []struct {
		Name         string
		InPkg        []*model.PkgInputSpec
		Calls        []call
		Query        *model.IsDependencySpec
		ExpID        []*model.IsDependency
		ExpIngestErr bool
		ExpQueryErr  bool
	}{
		{
			Name:  "HappyPath",
			InPkg: []*model.PkgInputSpec{p1, p2},
			Calls: []call{
				{
					P1: p1,
					P2: p2,
					ID: &model.IsDependencyInputSpec{
						Justification: "test justification",
					},
				},
			},
			Query: &model.IsDependencySpec{
				Justification: ptrfrom.String("test justification"),
			},
			ExpID: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p2outName,
					Justification:    "test justification",
				},
			},
		},
		{
			Name:  "Ingest same",
			InPkg: []*model.PkgInputSpec{p1, p2},
			Calls: []call{
				{
					P1: p1,
					P2: p2,
					ID: &model.IsDependencyInputSpec{
						Justification: "test justification",
					},
				},
				{
					P1: p1,
					P2: p2,
					ID: &model.IsDependencyInputSpec{
						Justification: "test justification",
					},
				},
			},
			Query: &model.IsDependencySpec{
				Justification: ptrfrom.String("test justification"),
			},
			ExpID: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p2outName,
					Justification:    "test justification",
				},
			},
		},
		{
			Name:  "Ingest same, different version",
			InPkg: []*model.PkgInputSpec{p1, p2, p3},
			Calls: []call{
				{
					P1: p1,
					P2: p2,
					ID: &model.IsDependencyInputSpec{
						Justification: "test justification",
					},
				},
				{
					P1: p1,
					P2: p3,
					ID: &model.IsDependencyInputSpec{
						Justification: "test justification",
					},
				},
			},
			Query: &model.IsDependencySpec{
				Justification: ptrfrom.String("test justification"),
			},
			ExpID: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p2outName,
					Justification:    "test justification",
				},
			},
		},
		{
			Name:  "Query on Justification",
			InPkg: []*model.PkgInputSpec{p1, p2},
			Calls: []call{
				{
					P1: p1,
					P2: p2,
					ID: &model.IsDependencyInputSpec{
						Justification: "test justification one",
					},
				},
				{
					P1: p1,
					P2: p2,
					ID: &model.IsDependencyInputSpec{
						Justification: "test justification two",
					},
				},
			},
			Query: &model.IsDependencySpec{
				Justification: ptrfrom.String("test justification one"),
			},
			ExpID: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p2outName,
					Justification:    "test justification one",
				},
			},
		},
		//{ ids aren't known
		//	Name:  "Query on pkg",
		//	InPkg: []*model.PkgInputSpec{p1, p2, p3},
		//	Calls: []call{
		//		{
		//			P1: p1,
		//			P2: p2,
		//			ID: &model.IsDependencyInputSpec{},
		//		},
		//		{
		//			P1: p2,
		//			P2: p3,
		//			ID: &model.IsDependencyInputSpec{},
		//		},
		//	},
		//	Query: &model.IsDependencySpec{
		//		Package: &model.PkgSpec{
		//			ID: ptrfrom.String("5"),
		//		},
		//	},
		//	ExpID: []*model.IsDependency{
		//		{
		//			Package:          p1out,
		//			DependentPackage: p2outName,
		//		},
		//	},
		//},
		{
			Name:  "Query on dep pkg",
			InPkg: []*model.PkgInputSpec{p1, p2, p4},
			Calls: []call{
				{
					P1: p2,
					P2: p4,
					ID: &model.IsDependencyInputSpec{},
				},
				{
					P1: p2,
					P2: p1,
					ID: &model.IsDependencyInputSpec{},
				},
			},
			Query: &model.IsDependencySpec{
				DependentPackage: &model.PkgNameSpec{
					Name: ptrfrom.String("openssl"),
				},
			},
			ExpID: []*model.IsDependency{
				{
					Package:          p2out,
					DependentPackage: p4outName,
				},
			},
		},
		{
			Name:  "Query on pkg multiple",
			InPkg: []*model.PkgInputSpec{p1, p2, p3},
			Calls: []call{
				{
					P1: p1,
					P2: p2,
					ID: &model.IsDependencyInputSpec{},
				},
				{
					P1: p3,
					P2: p2,
					ID: &model.IsDependencyInputSpec{},
				},
			},
			Query: &model.IsDependencySpec{
				Package: &model.PkgSpec{
					Type: ptrfrom.String("pypi"),
				},
			},
			ExpID: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p1outName,
				},
				{
					Package:          p3out,
					DependentPackage: p1outName,
				},
			},
		},
		{
			Name:  "Query on both pkgs",
			InPkg: []*model.PkgInputSpec{p1, p2, p3, p4},
			Calls: []call{
				{
					P1: p2,
					P2: p1,
					ID: &model.IsDependencyInputSpec{},
				},
				{
					P1: p3,
					P2: p4,
					ID: &model.IsDependencyInputSpec{},
				},
			},
			Query: &model.IsDependencySpec{
				Package: &model.PkgSpec{
					Subpath: ptrfrom.String("saved_model_cli.py"),
				},
				DependentPackage: &model.PkgNameSpec{
					Name: ptrfrom.String("openssl"),
				},
			},
			ExpID: []*model.IsDependency{
				{
					Package:          p3out,
					DependentPackage: p4outName,
				},
			},
		},
		{
			Name:  "Query none",
			InPkg: []*model.PkgInputSpec{p1, p2, p3},
			Calls: []call{
				{
					P1: p1,
					P2: p2,
					ID: &model.IsDependencyInputSpec{},
				},
				{
					P1: p2,
					P2: p3,
					ID: &model.IsDependencyInputSpec{},
				},
				{
					P1: p1,
					P2: p3,
					ID: &model.IsDependencyInputSpec{},
				},
			},
			Query: &model.IsDependencySpec{
				Package: &model.PkgSpec{
					Subpath: ptrfrom.String("asdf"),
				},
			},
			ExpID: nil,
		},
		//{ we don't know the id
		//	Name:  "Query on ID",
		//	InPkg: []*model.PkgInputSpec{p1, p2, p3},
		//	Calls: []call{
		//		{
		//			P1: p1,
		//			P2: p2,
		//			ID: &model.IsDependencyInputSpec{},
		//		},
		//		{
		//			P1: p2,
		//			P2: p3,
		//			ID: &model.IsDependencyInputSpec{},
		//		},
		//		{
		//			P1: p1,
		//			P2: p3,
		//			ID: &model.IsDependencyInputSpec{},
		//		},
		//	},
		//	Query: &model.IsDependencySpec{
		//		ID: ptrfrom.String("9"),
		//	},
		//	ExpID: []*model.IsDependency{
		//		{
		//			Package:          p2out,
		//			DependentPackage: p1outName,
		//		},
		//	},
		//},
		{
			Name:  "Query on Range",
			InPkg: []*model.PkgInputSpec{p1, p2},
			Calls: []call{
				{
					P1: p1,
					P2: p1,
					ID: &model.IsDependencyInputSpec{
						VersionRange: "1-3",
					},
				},
				{
					P1: p2,
					P2: p1,
					ID: &model.IsDependencyInputSpec{
						VersionRange: "4-5",
					},
				},
			},
			Query: &model.IsDependencySpec{
				VersionRange: ptrfrom.String("1-3"),
			},
			ExpID: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p1outName,
					VersionRange:     "1-3",
				},
			},
		},
		{
			Name:  "Query on DependencyType",
			InPkg: []*model.PkgInputSpec{p1, p2},
			Calls: []call{
				{
					P1: p1,
					P2: p1,
					ID: &model.IsDependencyInputSpec{
						DependencyType: model.DependencyTypeDirect,
					},
				},
				{
					P1: p2,
					P2: p1,
					ID: &model.IsDependencyInputSpec{
						DependencyType: model.DependencyTypeIndirect,
					},
				},
			},
			Query: &model.IsDependencySpec{
				DependencyType: (*model.DependencyType)(ptrfrom.String(string(model.DependencyTypeIndirect))),
			},
			ExpID: []*model.IsDependency{
				{
					Package:          p2out,
					DependentPackage: p1outName,
					DependencyType:   model.DependencyTypeIndirect,
				},
			},
		},
		{
			Name:  "Ingest no P1",
			InPkg: []*model.PkgInputSpec{p2},
			Calls: []call{
				{
					P1: p1,
					P2: p2,
					ID: &model.IsDependencyInputSpec{},
				},
			},
			ExpIngestErr: true,
		},
		{
			Name:  "Ingest no P2",
			InPkg: []*model.PkgInputSpec{p1},
			Calls: []call{
				{
					P1: p1,
					P2: p4,
					ID: &model.IsDependencyInputSpec{},
				},
			},
			ExpIngestErr: true,
		},
		//{ improper ids will return no results
		//	Name:  "Query bad ID",
		//	InPkg: []*model.PkgInputSpec{p1, p2, p3},
		//	Calls: []call{
		//		{
		//			P1: p1,
		//			P2: p2,
		//			ID: &model.IsDependencyInputSpec{},
		//		},
		//		{
		//			P1: p2,
		//			P2: p3,
		//			ID: &model.IsDependencyInputSpec{},
		//		},
		//		{
		//			P1: p1,
		//			P2: p3,
		//			ID: &model.IsDependencyInputSpec{},
		//		},
		//	},
		//	Query: &model.IsDependencySpec{
		//		ID: ptrfrom.String("asdf"),
		//	},
		//	ExpQueryErr: true,
		//},
	}
	ignoreID := cmp.FilterPath(func(p cmp.Path) bool {
		return strings.Compare(".ID", p[len(p)-1].String()) == 0
	}, cmp.Ignore())
	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			b, err := CreateGremlinClientForIntegrationTest()
			if err != nil {
				t.Errorf("failed to create gremlin client. error = %v", err)
				return
			}
			for _, a := range test.InPkg {
				if _, err := b.IngestPackage(ctx, *a); err != nil {
					t.Fatalf("Could not ingest pkg: %v", err)
				}
			}
			for _, o := range test.Calls {
				_, err := b.IngestDependency(ctx, *o.P1, *o.P2, *o.ID)
				if (err != nil) != test.ExpIngestErr {
					t.Fatalf("did not get expected ingest error, want: %v, got: %v", test.ExpIngestErr, err)
				}
				if err != nil {
					return
				}
			}
			got, err := b.IsDependency(ctx, test.Query)
			if (err != nil) != test.ExpQueryErr {
				t.Fatalf("did not get expected query error, want: %v, got: %v", test.ExpQueryErr, err)
			}
			if err != nil {
				return
			}
			// less := func(a, b *model.Package) bool { return a.Version < b.Version }
			// for _, he := range got {
			// 	slices.SortFunc(he.Packages, less)
			// }
			// for _, he := range test.ExpID {
			// 	slices.SortFunc(he.Packages, less)
			// }
			if diff := cmp.Diff(test.ExpID, got, ignoreID); diff != "" {
				t.Errorf("Unexpected results. (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIsDependencies(t *testing.T) {
	type call struct {
		P1s []*model.PkgInputSpec
		P2s []*model.PkgInputSpec
		IDs []*model.IsDependencyInputSpec
	}
	tests := []struct {
		Name         string
		InPkg        []*model.PkgInputSpec
		Calls        []call
		ExpID        []*model.IsDependency
		ExpIngestErr bool
		ExpQueryErr  bool
	}{
		{
			Name:  "HappyPath",
			InPkg: []*model.PkgInputSpec{p1, p2, p3, p4},
			Calls: []call{{
				P1s: []*model.PkgInputSpec{p1, p2},
				P2s: []*model.PkgInputSpec{p2, p4},
				IDs: []*model.IsDependencyInputSpec{
					{
						Justification: "test justification",
					},
					{
						Justification: "test justification",
					},
				},
			}},
			ExpID: []*model.IsDependency{
				{
					Package:          p1out,
					DependentPackage: p2outName,
					Justification:    "test justification",
				},
				{
					Package:          p2out,
					DependentPackage: p4outName,
					Justification:    "test justification",
				},
			},
		},
	}
	ignoreID := cmp.FilterPath(func(p cmp.Path) bool {
		return strings.Compare(".ID", p[len(p)-1].String()) == 0
	}, cmp.Ignore())
	ctx := context.Background()
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			b, err := CreateGremlinClientForIntegrationTest()
			if err != nil {
				t.Errorf("failed to create gremlin client. error = %v", err)
				return
			}
			for _, a := range test.InPkg {
				if _, err := b.IngestPackage(ctx, *a); err != nil {
					t.Fatalf("Could not ingest pkg: %v", err)
				}
			}
			for _, o := range test.Calls {
				got, err := b.IngestDependencies(ctx, o.P1s, o.P2s, o.IDs)
				if (err != nil) != test.ExpIngestErr {
					t.Fatalf("did not get expected ingest error, want: %v, got: %v", test.ExpIngestErr, err)
				}
				if err != nil {
					return
				}
				if diff := cmp.Diff(test.ExpID, got, ignoreID); diff != "" {
					t.Errorf("Unexpected results. (-want +got):\n%s", diff)
				}
			}

		})
	}
}
