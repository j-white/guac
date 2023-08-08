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
	"golang.org/x/exp/slices"
)

func TestHashEqual(t *testing.T) {
	type call struct {
		A1 *model.ArtifactInputSpec
		A2 *model.ArtifactInputSpec
		HE *model.HashEqualInputSpec
	}
	tests := []struct {
		Name         string
		InArt        []*model.ArtifactInputSpec
		Calls        []call
		Query        *model.HashEqualSpec
		ExpHE        []*model.HashEqual
		ExpIngestErr bool
		ExpQueryErr  bool
	}{
		{
			Name:  "HappyPath",
			InArt: []*model.ArtifactInputSpec{a1, a2},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{
						Justification: "test justification",
					},
				},
			},
			Query: &model.HashEqualSpec{
				Justification: ptrfrom.String("test justification"),
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts:     []*model.Artifact{a1out, a2out},
					Justification: "test justification",
				},
			},
		},
		{
			Name:  "Ingest same, different order",
			InArt: []*model.ArtifactInputSpec{a1, a2},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{
						Justification: "test justification",
					},
				},
				{
					A1: a2,
					A2: a1,
					HE: &model.HashEqualInputSpec{
						Justification: "test justification",
					},
				},
			},
			Query: &model.HashEqualSpec{
				Justification: ptrfrom.String("test justification"),
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts:     []*model.Artifact{a1out, a2out},
					Justification: "test justification",
				},
			},
		},
		{
			Name:  "Query on Justification",
			InArt: []*model.ArtifactInputSpec{a1, a2},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{
						Justification: "test justification one",
					},
				},
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{
						Justification: "test justification two",
					},
				},
			},
			Query: &model.HashEqualSpec{
				Justification: ptrfrom.String("test justification one"),
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts:     []*model.Artifact{a1out, a2out},
					Justification: "test justification one",
				},
			},
		},
		{
			Name:  "Query on artifact",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a1,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{{
					ID: ptrfrom.String("4"),
				}},
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a1out, a3out},
				},
			},
		},
		{
			Name:  "Query on artifact multiple",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a1,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{{
					ID: ptrfrom.String("2"),
				}},
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a1out, a2out},
				},
				{
					Artifacts: []*model.Artifact{a1out, a3out},
				},
			},
		},
		{
			Name:  "Query on artifact algo",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a1,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{{
					Algorithm: ptrfrom.String("sha1"),
				}},
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a1out, a2out},
				},
			},
		},
		{
			Name:  "Query on artifact algo and hash",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a1,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{{
					Algorithm: ptrfrom.String("sha1"),
					Digest:    ptrfrom.String("7A8F47318E4676DACB0142AFA0B83029CD7BEFD9"),
				}},
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a1out, a2out},
				},
			},
		},
		{
			Name:  "Query on both artifacts",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a2,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{
					{
						Algorithm: ptrfrom.String("sha1"),
						Digest:    ptrfrom.String("7A8F47318E4676DACB0142AFA0B83029CD7BEFD9"),
					},
					{
						ID: ptrfrom.String("4"),
					},
				},
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a2out, a3out},
				},
			},
		},
		{
			Name:  "Query on both artifacts, one filter",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a2,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a1,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{
					{
						Algorithm: ptrfrom.String("sha1"),
					},
					{
						Digest: ptrfrom.String("6bbb0da1891646e58eb3e6a63af3a6fc3c8eb5a0d44824cba581d2e14a0450cf"),
					},
				},
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a1out, a2out},
				},
			},
		},
		{
			Name:  "Query none",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a2,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a1,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{
					{
						Algorithm: ptrfrom.String("gitHash"),
					},
					{
						Digest: ptrfrom.String("6bbb0da1891646e58eb3e6a63af3a6fc3c8eb5a0d44824cba581d2e14a0450cf"),
					},
				},
			},
			ExpHE: nil,
		},
		{
			Name:  "Query on ID",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a2,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a1,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
			},
			Query: &model.HashEqualSpec{
				ID: ptrfrom.String("6"),
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a3out, a2out},
				},
			},
		},
		{
			Name:  "Ingest no A1",
			InArt: []*model.ArtifactInputSpec{a2},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
			},
			ExpIngestErr: true,
		},
		{
			Name:  "Ingest no A2",
			InArt: []*model.ArtifactInputSpec{a1},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
			},
			ExpIngestErr: true,
		},
		{
			Name:  "Query three",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a2,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a1,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{
					{
						Algorithm: ptrfrom.String("gitHash"),
					},
					{
						Digest: ptrfrom.String("6bbb0da1891646e58eb3e6a63af3a6fc3c8eb5a0d44824cba581d2e14a0450cf"),
					},
					{
						Digest: ptrfrom.String("asdf"),
					},
				},
			},
			ExpQueryErr: true,
		},
		{
			Name:  "Query bad ID",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: a1,
					A2: a2,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a2,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
				{
					A1: a1,
					A2: a3,
					HE: &model.HashEqualInputSpec{},
				},
			},
			Query: &model.HashEqualSpec{
				ID: ptrfrom.String("asdf"),
			},
			ExpQueryErr: true,
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
			for _, a := range test.InArt {
				if _, err := b.IngestArtifact(ctx, a); err != nil {
					t.Fatalf("Could not ingest artifact: %v", err)
				}
			}
			for _, o := range test.Calls {
				_, err := b.IngestHashEqual(ctx, *o.A1, *o.A2, *o.HE)
				if (err != nil) != test.ExpIngestErr {
					t.Fatalf("did not get expected ingest error, want: %v, got: %v", test.ExpIngestErr, err)
				}
				if err != nil {
					return
				}
			}
			got, err := b.HashEqual(ctx, test.Query)
			if (err != nil) != test.ExpQueryErr {
				t.Fatalf("did not get expected query error, want: %v, got: %v", test.ExpQueryErr, err)
			}
			if err != nil {
				return
			}
			less := func(a, b *model.Artifact) bool { return a.Digest < b.Digest }
			for _, he := range got {
				slices.SortFunc(he.Artifacts, less)
			}
			for _, he := range test.ExpHE {
				slices.SortFunc(he.Artifacts, less)
			}
			if diff := cmp.Diff(test.ExpHE, got, ignoreID); diff != "" {
				t.Errorf("Unexpected results. (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIngestHashEquals(t *testing.T) {
	type call struct {
		A1 []*model.ArtifactInputSpec
		A2 []*model.ArtifactInputSpec
		HE []*model.HashEqualInputSpec
	}
	tests := []struct {
		Name         string
		InArt        []*model.ArtifactInputSpec
		Calls        []call
		Query        *model.HashEqualSpec
		ExpHE        []*model.HashEqual
		ExpIngestErr bool
		ExpQueryErr  bool
	}{
		{
			Name:  "HappyPath",
			InArt: []*model.ArtifactInputSpec{a1, a2},
			Calls: []call{
				{
					A1: []*model.ArtifactInputSpec{a1},
					A2: []*model.ArtifactInputSpec{a2},
					HE: []*model.HashEqualInputSpec{
						{
							Justification: "test justification",
						},
					},
				},
			},
			Query: &model.HashEqualSpec{
				Justification: ptrfrom.String("test justification"),
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts:     []*model.Artifact{a1out, a2out},
					Justification: "test justification",
				},
			},
		},
		{
			Name:  "Ingest same, different order",
			InArt: []*model.ArtifactInputSpec{a1, a2},
			Calls: []call{
				{
					A1: []*model.ArtifactInputSpec{a1, a2},
					A2: []*model.ArtifactInputSpec{a2, a1},
					HE: []*model.HashEqualInputSpec{
						{
							Justification: "test justification",
						},
						{
							Justification: "test justification",
						},
					},
				},
			},
			Query: &model.HashEqualSpec{
				Justification: ptrfrom.String("test justification"),
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts:     []*model.Artifact{a1out, a2out},
					Justification: "test justification",
				},
			},
		},
		{
			Name:  "Query on Justification",
			InArt: []*model.ArtifactInputSpec{a1, a2},
			Calls: []call{
				{
					A1: []*model.ArtifactInputSpec{a1, a1},
					A2: []*model.ArtifactInputSpec{a2, a2},
					HE: []*model.HashEqualInputSpec{
						{
							Justification: "test justification one",
						},
						{
							Justification: "test justification two",
						},
					},
				},
			},
			Query: &model.HashEqualSpec{
				Justification: ptrfrom.String("test justification one"),
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts:     []*model.Artifact{a1out, a2out},
					Justification: "test justification one",
				},
			},
		},
		{
			Name:  "Query on artifact",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: []*model.ArtifactInputSpec{a1, a1},
					A2: []*model.ArtifactInputSpec{a2, a3},
					HE: []*model.HashEqualInputSpec{
						{},
						{},
					},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{{
					ID: ptrfrom.String("4"),
				}},
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a1out, a3out},
				},
			},
		},
		{
			Name:  "Query on artifact multiple",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: []*model.ArtifactInputSpec{a1, a1},
					A2: []*model.ArtifactInputSpec{a2, a3},
					HE: []*model.HashEqualInputSpec{
						{},
						{},
					},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{{
					ID: ptrfrom.String("2"),
				}},
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a1out, a2out},
				},
				{
					Artifacts: []*model.Artifact{a1out, a3out},
				},
			},
		},
		{
			Name:  "Query on artifact algo",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: []*model.ArtifactInputSpec{a1, a1},
					A2: []*model.ArtifactInputSpec{a2, a3},
					HE: []*model.HashEqualInputSpec{
						{},
						{},
					},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{{
					Algorithm: ptrfrom.String("sha1"),
				}},
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a1out, a2out},
				},
			},
		},
		{
			Name:  "Query on artifact algo and hash",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: []*model.ArtifactInputSpec{a1, a1},
					A2: []*model.ArtifactInputSpec{a2, a3},
					HE: []*model.HashEqualInputSpec{
						{},
						{},
					},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{{
					Algorithm: ptrfrom.String("sha1"),
					Digest:    ptrfrom.String("7A8F47318E4676DACB0142AFA0B83029CD7BEFD9"),
				}},
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a1out, a2out},
				},
			},
		},
		{
			Name:  "Query on both artifacts",
			InArt: []*model.ArtifactInputSpec{a1, a2, a3},
			Calls: []call{
				{
					A1: []*model.ArtifactInputSpec{a1, a2},
					A2: []*model.ArtifactInputSpec{a2, a3},
					HE: []*model.HashEqualInputSpec{
						{},
						{},
					},
				},
			},
			Query: &model.HashEqualSpec{
				Artifacts: []*model.ArtifactSpec{
					{
						Algorithm: ptrfrom.String("sha1"),
						Digest:    ptrfrom.String("7A8F47318E4676DACB0142AFA0B83029CD7BEFD9"),
					},
					{
						ID: ptrfrom.String("4"),
					},
				},
			},
			ExpHE: []*model.HashEqual{
				{
					Artifacts: []*model.Artifact{a2out, a3out},
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
			for _, a := range test.InArt {
				if _, err := b.IngestArtifact(ctx, a); err != nil {
					t.Fatalf("Could not ingest artifact: %v", err)
				}
			}
			for _, o := range test.Calls {
				_, err := b.IngestHashEquals(ctx, o.A1, o.A2, o.HE)
				if (err != nil) != test.ExpIngestErr {
					t.Fatalf("did not get expected ingest error, want: %v, got: %v", test.ExpIngestErr, err)
				}
				if err != nil {
					return
				}
			}
			got, err := b.HashEqual(ctx, test.Query)
			if (err != nil) != test.ExpQueryErr {
				t.Fatalf("did not get expected query error, want: %v, got: %v", test.ExpQueryErr, err)
			}
			if err != nil {
				return
			}
			less := func(a, b *model.Artifact) bool { return a.Digest < b.Digest }
			for _, he := range got {
				slices.SortFunc(he.Artifacts, less)
			}
			for _, he := range test.ExpHE {
				slices.SortFunc(he.Artifacts, less)
			}
			if diff := cmp.Diff(test.ExpHE, got, ignoreID); diff != "" {
				t.Errorf("Unexpected results. (-want +got):\n%s", diff)
			}
		})
	}
}
