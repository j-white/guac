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

func printSchema(remote *gremlingo.DriverRemoteConnection) (string, error) {
	r := new(gremlingo.RequestOptionsBuilder).Create()
	stmt := "mgmt = graph.openManagement()\nmgmt.printSchema()\n"
	rs, err := remote.SubmitWithOptions(stmt, r)
	results, err := rs.All()
	return fmt.Sprintf("%s", results), err
}

func createIndexForVertexPropertyKey(remote *gremlingo.DriverRemoteConnection, key string) error {
	// FIXME: bind w/ parameters instead of sprintf (avoid possiblity of injection)
	createIndexStmt := fmt.Sprintf("mgmt = graph.openManagement()\n"+
		"propKey = mgmt.containsPropertyKey('%s') ? mgmt.getPropertyKey('%s') : mgmt.makePropertyKey('%s').dataType(String.class).cardinality(Cardinality.SINGLE).make()\n"+
		"index = mgmt.getGraphIndex('by%sComposite')\n"+
		"index = index == null ? mgmt.buildIndex('by%sComposite', Vertex.class).addKey(propKey).buildCompositeIndex() : index\n"+
		"mgmt.commit()\n", key, key, key, key, key)
	r := new(gremlingo.RequestOptionsBuilder).Create()
	rs, err := remote.SubmitWithOptions(createIndexStmt, r)
	if err != nil {
		return err
	}
	_, err = rs.All()
	if err != nil {
		return err
	}
	return err
}

func createIndicesForVertexProperties(remote *gremlingo.DriverRemoteConnection, key ...string) error {
	for _, key := range key {
		err := createIndexForVertexPropertyKey(remote, key)
		if err != nil {
			return err
		}
	}
	return nil
}

func createIndexForEdge(remote *gremlingo.DriverRemoteConnection, edgeLabel string, vertexPropertyKey string) error {
	// FIXME: bind w/ parameters instead of sprintf (avoid possiblity of injection)
	createIndexStmt := fmt.Sprintf("mgmt = graph.openManagement()\n"+
		"propKey = mgmt.containsPropertyKey('%s') ? mgmt.getPropertyKey('%s') : mgmt.makePropertyKey('%s').dataType(String.class).cardinality(Cardinality.SINGLE).make()\n"+
		"edgeLabel = mgmt.getEdgeLabel('%s')\n"+
		"edgeLabel = edgeLabel == null ? mgmt.makeEdgeLabel('%s').make() : edgeLabel\n"+
		"index = mgmt.getGraphIndex('by%sEdge')\n"+
		"index = index == null ? mgmt.buildEdgeIndex(edgeLabel, 'by%sEdge', Direction.BOTH, Order.desc, propKey) : index\n"+
		"mgmt.commit()\n",
		vertexPropertyKey, vertexPropertyKey, vertexPropertyKey,
		edgeLabel, edgeLabel,
		vertexPropertyKey, vertexPropertyKey)
	r := new(gremlingo.RequestOptionsBuilder).Create()
	rs, err := remote.SubmitWithOptions(createIndexStmt, r)
	if err != nil {
		return err
	}
	_, err = rs.All()
	if err != nil {
		return err
	}
	return err

}
