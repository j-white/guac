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
	"errors"
	"fmt"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
)

func supportsTransactions(remote *gremlingo.DriverRemoteConnection) (bool, error) {
	r := new(gremlingo.RequestOptionsBuilder).Create()
	stmt := "graph.features().graph().supportsTransactions()\n"
	rs, err := remote.SubmitWithOptions(stmt, r)
	result, hasResult, err := rs.One()
	if err != nil {
		return false, err
	}
	if !hasResult {
		return false, errors.New("query to verify transaction supported completed normally, but did not produce a result")
	}
	isSupported, err := result.GetBool()
	if err != nil {
		return false, err
	}
	return isSupported, nil
}

func deleteAllVerticesAndEdges(remote *gremlingo.DriverRemoteConnection) error {
	r := new(gremlingo.RequestOptionsBuilder).Create()
	stmt := "g.V().drop()\n"
	rs, err := remote.SubmitWithOptions(stmt, r)
	result, hasResult, err := rs.One()
	if err != nil {
		return err
	}
	if hasResult {
		return fmt.Errorf("query to delete graph completed normally, but produced an unexpected result: %v", result)
	}
	return nil
}

func rollbackAllOpenTransactions(remote *gremlingo.DriverRemoteConnection) error {
	// pulled from https://groups.google.com/g/janusgraph-users/c/GfYM538KsTo
	r := new(gremlingo.RequestOptionsBuilder).Create()
	stmt := "graph.getOpenTransactions().forEach { tx -> tx.rollback() }\n"
	rs, err := remote.SubmitWithOptions(stmt, r)
	_, _, err = rs.One()
	if err != nil {
		return err
	}
	return nil
}

func forceCloseAllOpenInstances(remote *gremlingo.DriverRemoteConnection) error {
	// pulled from https://groups.google.com/g/janusgraph-users/c/GfYM538KsTo
	r := new(gremlingo.RequestOptionsBuilder).Create()
	stmt := "mgmt = graph.openManagement()\n" +
		"mgmt.getOpenInstances()\n" +
		"mgmt.getOpenInstances().forEach {\n" +
		"  if (it.reverse().take(1) != \")\") {\n" +
		"    mgmt.forceCloseInstance(it)\n" +
		"  }\n" +
		"}\n" +
		"mgmt.commit()\n"
	rs, err := remote.SubmitWithOptions(stmt, r)
	_, _, err = rs.One()
	if err != nil {
		return err
	}
	return nil
}
