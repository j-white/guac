package tinkerpop

import (
	"encoding/binary"
	"errors"
	"fmt"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
)

const (
	relationIdentifierType uint32 = 0x1001
	longMarker             byte   = 0
	stringMarker           byte   = 1
	valueFlagNull          byte   = 1
	valueFlagNone          byte   = 0
)

type janusgraphRelationIdentifier struct {
	OutVertexIdLong   int64
	OutVertexIdString string
	TypeId            int64
	RelationId        int64
	InVertexIdLong    int64
	InVertexIdString  string
}

func registerCustomTypeReaders() {
	gremlingo.RegisterCustomTypeReader("janusgraph.RelationIdentifier", janusgraphRelationIdentifierReader)
}

func janusgraphRelationIdentifierReader(data *[]byte, i *int) (interface{}, error) {
	r := new(janusgraphRelationIdentifier)

	// expect type code
	customDataTyp := readUint32Safe(data, i)
	if customDataTyp != relationIdentifierType {
		return nil, fmt.Errorf("unknown type code. got 0x%x, expected 0x%x", customDataTyp, relationIdentifierType)
	}

	// value flag, expect this to be non-nullable
	if readByteSafe(data, i) != valueFlagNone {
		return nil, errors.New("")
	}

	// outVertexId
	if readByteSafe(data, i) == longMarker {
		r.OutVertexIdLong = readLongSafe(data, i)
	} else {
		vertexId, err := readString(data, i)
		if err != nil {
			return nil, err
		}
		r.OutVertexIdString = vertexId.(string)
	}

	r.TypeId = readLongSafe(data, i)
	r.RelationId = readLongSafe(data, i)

	// inVertexId
	if readByteSafe(data, i) == longMarker {
		r.InVertexIdLong = readLongSafe(data, i)
	} else {
		vertexId, err := readString(data, i)
		if err != nil {
			return nil, err
		}
		r.InVertexIdString = vertexId.(string)
	}

	return r, nil
}

func readTemp(data *[]byte, i *int, len int) *[]byte {
	tmp := make([]byte, len)
	for j := 0; j < len; j++ {
		tmp[j] = (*data)[j+*i]
	}
	*i += len
	return &tmp
}

func readUint32Safe(data *[]byte, i *int) uint32 {
	return binary.BigEndian.Uint32(*readTemp(data, i, 4))
}

func readByteSafe(data *[]byte, i *int) byte {
	*i++
	return (*data)[*i-1]
}

func readLongSafe(data *[]byte, i *int) int64 {
	return int64(binary.BigEndian.Uint64(*readTemp(data, i, 8)))
}

func readString(data *[]byte, i *int) (interface{}, error) {
	sz := int(readUint32Safe(data, i))
	if sz == 0 {
		return "", nil
	}
	*i += sz
	return string((*data)[*i-sz : *i]), nil
}

func printSchema(remote *gremlingo.DriverRemoteConnection) error {
	r := new(gremlingo.RequestOptionsBuilder).AddBinding("x", 2).AddBinding("y", 5).Create()
	stmt := "mgmt = graph.openManagement()\nmgmt.printSchema()\n"
	rs, err := remote.SubmitWithOptions(stmt, r)
	results, err := rs.All()
	fmt.Println("results", results)
	return err
}

func createIndices(remote *gremlingo.DriverRemoteConnection) error {
	r := new(gremlingo.RequestOptionsBuilder).AddBinding("x", 2).AddBinding("y", 5).Create()
	//create := "graph.tx().rollback() //Never create new indexes while a transaction is active\nmgmt = graph.openManagement()\nname = mgmt.getPropertyKey('name')\nage = mgmt.getPropertyKey('age')\nmgmt.buildIndex('byNameComposite', Vertex.class).addKey(name).buildCompositeIndex()\nmgmt.buildIndex('byNameAndAgeComposite', Vertex.class).addKey(name).addKey(age).buildCompositeIndex()\nmgmt.commit()"
	create := "mgmt = graph.openManagement()\nname = mgmt.getPropertyKey('name')\ntype = mgmt.getPropertyKey('type')\nmgmt.buildIndex('byNameComposite', Vertex.class).addKey(name).buildCompositeIndex()\nmgmt.buildIndex('byNameAndTypeComposite', Vertex.class).addKey(name).addKey(type).buildCompositeIndex()\nmgmt.commit()\nManagementSystem.awaitGraphIndexStatus(graph, 'byNameComposite').call()\nManagementSystem.awaitGraphIndexStatus(graph, 'byNameAndTypeComposite').call()\nmgmt = graph.openManagement()\nmgmt.updateIndex(mgmt.getGraphIndex(\"byNameComposite\"), SchemaAction.REINDEX).get()\nmgmt.updateIndex(mgmt.getGraphIndex(\"byNameAndTypeComposite\"), SchemaAction.REINDEX).get()\nmgmt.commit()"
	rs, err := remote.SubmitWithOptions(create, r)
	fmt.Println("results", rs)
	return err
}
