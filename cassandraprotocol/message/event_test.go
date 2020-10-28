package message

import (
	"bytes"
	"fmt"
	"github.com/datastax/go-cassandra-native-protocol/cassandraprotocol"
	"github.com/datastax/go-cassandra-native-protocol/cassandraprotocol/primitives"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestEventCodec_Encode(test *testing.T) {
	codec := &EventCodec{}
	// versions < 4
	for version := cassandraprotocol.ProtocolVersionMin; version < cassandraprotocol.ProtocolVersion4; version++ {
		test.Run(fmt.Sprintf("version %v", version), func(test *testing.T) {
			tests := []encodeTestCase{
				{
					"schema change event keyspace",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetKeyspace,
						Keyspace:   "ks1",
					},
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 8, K, E, Y, S, P, A, C, E,
						0, 3, k, s, _1,
					},
					nil,
				},
				{
					"schema change event table",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetTable,
						Keyspace:   "ks1",
						Object:     "table1",
					},
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 5, T, A, B, L, E,
						0, 3, k, s, _1,
						0, 6, t, a, b, l, e, _1,
					},
					nil,
				},
				{
					"schema change event type",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetType,
						Keyspace:   "ks1",
						Object:     "udt1",
					},
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 4, T, Y, P, E,
						0, 3, k, s, _1,
						0, 4, u, d, t, _1,
					},
					nil,
				},
				{
					"schema change event function",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetFunction,
						Keyspace:   "ks1",
						Object:     "func1",
						Arguments:  []string{"int", "varchar"},
					},
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 8, F, U, N, C, T, I, O, N,
						0, 3, k, s, _1,
					},
					fmt.Errorf("FUNCTION schema change targets are not supported in protocol version %d", version),
				},
				{
					"schema change event aggregate",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetAggregate,
						Keyspace:   "ks1",
						Object:     "agg1",
						Arguments:  []string{"int", "varchar"},
					},
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 9, A, G, G, R, E, G, A, T, E,
						0, 3, k, s, _1,
					},
					fmt.Errorf("AGGREGATE schema change targets are not supported in protocol version %d", version),
				},
				{
					"status change event",
					&StatusChangeEvent{
						ChangeType: cassandraprotocol.StatusChangeTypeUp,
						Address: &primitives.Inet{
							Addr: net.IPv4(192, 168, 1, 1),
							Port: 9042,
						},
					},
					[]byte{
						0, 13, S, T, A, T, U, S, __, C, H, A, N, G, E,
						0, 2, U, P,
						4, 192, 168, 1, 1,
						0, 0, 0x23, 0x52,
					},
					nil,
				},
				{
					"topology change event",
					&TopologyChangeEvent{
						ChangeType: cassandraprotocol.TopologyChangeTypeNewNode,
						Address: &primitives.Inet{
							Addr: net.IPv4(192, 168, 1, 1),
							Port: 9042,
						},
					},
					[]byte{
						0, 15, T, O, P, O, L, O, G, Y, __, C, H, A, N, G, E,
						0, 8, N, E, W, __, N, O, D, E,
						4, 192, 168, 1, 1,
						0, 0, 0x23, 0x52,
					},
					nil,
				},
			}
			for _, tt := range tests {
				test.Run(tt.name, func(t *testing.T) {
					dest := &bytes.Buffer{}
					err := codec.Encode(tt.input, dest, version)
					assert.Equal(t, tt.expected, dest.Bytes())
					assert.Equal(t, tt.err, err)
				})
			}
		})
	}
	// versions >= 4
	for version := cassandraprotocol.ProtocolVersion4; version <= cassandraprotocol.ProtocolVersionBeta; version++ {
		test.Run(fmt.Sprintf("version %v", version), func(test *testing.T) {
			tests := []encodeTestCase{
				{
					"schema change event keyspace",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetKeyspace,
						Keyspace:   "ks1",
					},
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 8, K, E, Y, S, P, A, C, E,
						0, 3, k, s, _1,
					},
					nil,
				},
				{
					"schema change event table",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetTable,
						Keyspace:   "ks1",
						Object:     "table1",
					},
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 5, T, A, B, L, E,
						0, 3, k, s, _1,
						0, 6, t, a, b, l, e, _1,
					},
					nil,
				},
				{
					"schema change event type",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetType,
						Keyspace:   "ks1",
						Object:     "udt1",
					},
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 4, T, Y, P, E,
						0, 3, k, s, _1,
						0, 4, u, d, t, _1,
					},
					nil,
				},
				{
					"schema change event function",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetFunction,
						Keyspace:   "ks1",
						Object:     "func1",
						Arguments:  []string{"int", "varchar"},
					},
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 8, F, U, N, C, T, I, O, N,
						0, 3, k, s, _1,
						0, 5, f, u, n, c, _1,
						0, 2,
						0, 3, i, n, t,
						0, 7, v, a, r, c, h, a, r,
					},
					nil,
				},
				{
					"schema change event aggregate",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetAggregate,
						Keyspace:   "ks1",
						Object:     "agg1",
						Arguments:  []string{"int", "varchar"},
					},
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 9, A, G, G, R, E, G, A, T, E,
						0, 3, k, s, _1,
						0, 4, a, g, g, _1,
						0, 2,
						0, 3, i, n, t,
						0, 7, v, a, r, c, h, a, r,
					},
					nil,
				},
				{
					"status change event",
					&StatusChangeEvent{
						ChangeType: cassandraprotocol.StatusChangeTypeUp,
						Address: &primitives.Inet{
							Addr: net.IPv4(192, 168, 1, 1),
							Port: 9042,
						},
					},
					[]byte{
						0, 13, S, T, A, T, U, S, __, C, H, A, N, G, E,
						0, 2, U, P,
						4, 192, 168, 1, 1,
						0, 0, 0x23, 0x52,
					},
					nil,
				},
				{
					"topology change event",
					&TopologyChangeEvent{
						ChangeType: cassandraprotocol.TopologyChangeTypeNewNode,
						Address: &primitives.Inet{
							Addr: net.IPv4(192, 168, 1, 1),
							Port: 9042,
						},
					},
					[]byte{
						0, 15, T, O, P, O, L, O, G, Y, __, C, H, A, N, G, E,
						0, 8, N, E, W, __, N, O, D, E,
						4, 192, 168, 1, 1,
						0, 0, 0x23, 0x52,
					},
					nil,
				},
			}
			for _, tt := range tests {
				test.Run(tt.name, func(t *testing.T) {
					dest := &bytes.Buffer{}
					err := codec.Encode(tt.input, dest, version)
					assert.Equal(t, tt.expected, dest.Bytes())
					assert.Equal(t, tt.err, err)
				})
			}
		})
	}
}

func TestEventCodec_EncodedLength(test *testing.T) {
	codec := &EventCodec{}
	for version := cassandraprotocol.ProtocolVersionMin; version <= cassandraprotocol.ProtocolVersionBeta; version++ {
		test.Run(fmt.Sprintf("version %v", version), func(test *testing.T) {
			tests := []encodedLengthTestCase{
				{
					"schema change event keyspace",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetKeyspace,
						Keyspace:   "ks1",
					},
					primitives.LengthOfString(cassandraprotocol.EventTypeSchemaChange) +
						primitives.LengthOfString(cassandraprotocol.SchemaChangeTypeCreated) +
						primitives.LengthOfString(cassandraprotocol.SchemaChangeTargetKeyspace) +
						primitives.LengthOfString("ks1"),
					nil,
				},
				{
					"schema change event table",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetTable,
						Keyspace:   "ks1",
						Object:     "table1",
					},
					primitives.LengthOfString(cassandraprotocol.EventTypeSchemaChange) +
						primitives.LengthOfString(cassandraprotocol.SchemaChangeTypeCreated) +
						primitives.LengthOfString(cassandraprotocol.SchemaChangeTargetTable) +
						primitives.LengthOfString("ks1") +
						primitives.LengthOfString("table1"),
					nil,
				},
				{
					"schema change event type",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetType,
						Keyspace:   "ks1",
						Object:     "udt1",
					},
					primitives.LengthOfString(cassandraprotocol.EventTypeSchemaChange) +
						primitives.LengthOfString(cassandraprotocol.SchemaChangeTypeCreated) +
						primitives.LengthOfString(cassandraprotocol.SchemaChangeTargetType) +
						primitives.LengthOfString("ks1") +
						primitives.LengthOfString("udt1"),
					nil,
				},
				{
					"schema change event function",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetFunction,
						Keyspace:   "ks1",
						Object:     "func1",
						Arguments:  []string{"int", "varchar"},
					},
					primitives.LengthOfString(cassandraprotocol.EventTypeSchemaChange) +
						primitives.LengthOfString(cassandraprotocol.SchemaChangeTypeCreated) +
						primitives.LengthOfString(cassandraprotocol.SchemaChangeTargetFunction) +
						primitives.LengthOfString("ks1") +
						primitives.LengthOfString("func1") +
						primitives.LengthOfStringList([]string{"int", "varchar"}),
					nil,
				},
				{
					"schema change event aggregate",
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetAggregate,
						Keyspace:   "ks1",
						Object:     "agg1",
						Arguments:  []string{"int", "varchar"},
					},
					primitives.LengthOfString(cassandraprotocol.EventTypeSchemaChange) +
						primitives.LengthOfString(cassandraprotocol.SchemaChangeTypeCreated) +
						primitives.LengthOfString(cassandraprotocol.SchemaChangeTargetAggregate) +
						primitives.LengthOfString("ks1") +
						primitives.LengthOfString("agg1") +
						primitives.LengthOfStringList([]string{"int", "varchar"}),
					nil,
				},
				{
					"status change event",
					&StatusChangeEvent{
						ChangeType: cassandraprotocol.StatusChangeTypeUp,
						Address: &primitives.Inet{
							Addr: net.IPv4(192, 168, 1, 1),
							Port: 9042,
						},
					},
					primitives.LengthOfString(cassandraprotocol.EventTypeStatusChange) +
						primitives.LengthOfString(cassandraprotocol.StatusChangeTypeUp) +
						primitives.LengthOfByte + net.IPv4len +
						primitives.LengthOfInt,
					nil,
				},
				{
					"topology change event",
					&TopologyChangeEvent{
						ChangeType: cassandraprotocol.TopologyChangeTypeNewNode,
						Address: &primitives.Inet{
							Addr: net.IPv4(192, 168, 1, 1),
							Port: 9042,
						},
					},
					primitives.LengthOfString(cassandraprotocol.EventTypeTopologyChange) +
						primitives.LengthOfString(cassandraprotocol.TopologyChangeTypeNewNode) +
						primitives.LengthOfByte + net.IPv4len +
						primitives.LengthOfInt,
					nil,
				},
			}
			for _, tt := range tests {
				test.Run(tt.name, func(t *testing.T) {
					actual, err := codec.EncodedLength(tt.input, version)
					assert.Equal(t, tt.expected, actual)
					assert.Equal(t, tt.err, err)
				})
			}
		})
	}
}

func TestEventCodec_Decode(test *testing.T) {
	codec := &EventCodec{}
	// versions < 4
	for version := cassandraprotocol.ProtocolVersionMin; version < cassandraprotocol.ProtocolVersion4; version++ {
		test.Run(fmt.Sprintf("version %v", version), func(test *testing.T) {
			tests := []decodeTestCase{
				{
					"schema change event keyspace",
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 8, K, E, Y, S, P, A, C, E,
						0, 3, k, s, _1,
					},
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetKeyspace,
						Keyspace:   "ks1",
					},
					nil,
				},
				{
					"schema change event table",
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 5, T, A, B, L, E,
						0, 3, k, s, _1,
						0, 6, t, a, b, l, e, _1,
					},
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetTable,
						Keyspace:   "ks1",
						Object:     "table1",
					},
					nil,
				},
				{
					"schema change event type",
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 4, T, Y, P, E,
						0, 3, k, s, _1,
						0, 4, u, d, t, _1,
					},
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetType,
						Keyspace:   "ks1",
						Object:     "udt1",
					},
					nil,
				},
				{
					"schema change event function",
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 8, F, U, N, C, T, I, O, N,
						0, 3, k, s, _1,
					},
					nil,
					fmt.Errorf("FUNCTION schema change targets are not supported in protocol version %d", version),
				},
				{
					"schema change event aggregate",
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 9, A, G, G, R, E, G, A, T, E,
						0, 3, k, s, _1,
					},
					nil,
					fmt.Errorf("AGGREGATE schema change targets are not supported in protocol version %d", version),
				},
				{
					"status change event",
					[]byte{
						0, 13, S, T, A, T, U, S, __, C, H, A, N, G, E,
						0, 2, U, P,
						4, 192, 168, 1, 1,
						0, 0, 0x23, 0x52,
					},
					&StatusChangeEvent{
						ChangeType: cassandraprotocol.StatusChangeTypeUp,
						Address: &primitives.Inet{
							Addr: net.IPv4(192, 168, 1, 1),
							Port: 9042,
						},
					},
					nil,
				},
				{
					"topology change event",
					[]byte{
						0, 15, T, O, P, O, L, O, G, Y, __, C, H, A, N, G, E,
						0, 8, N, E, W, __, N, O, D, E,
						4, 192, 168, 1, 1,
						0, 0, 0x23, 0x52,
					},
					&TopologyChangeEvent{
						ChangeType: cassandraprotocol.TopologyChangeTypeNewNode,
						Address: &primitives.Inet{
							Addr: net.IPv4(192, 168, 1, 1),
							Port: 9042,
						},
					},
					nil,
				},
			}
			for _, tt := range tests {
				test.Run(tt.name, func(t *testing.T) {
					source := bytes.NewBuffer(tt.input)
					actual, err := codec.Decode(source, version)
					assert.Equal(t, tt.expected, actual)
					assert.Equal(t, tt.err, err)
				})
			}
		})
	}
	// versions >= 4
	for version := cassandraprotocol.ProtocolVersion4; version <= cassandraprotocol.ProtocolVersionBeta; version++ {
		test.Run(fmt.Sprintf("version %v", version), func(test *testing.T) {
			tests := []decodeTestCase{
				{
					"schema change event keyspace",
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 8, K, E, Y, S, P, A, C, E,
						0, 3, k, s, _1,
					},
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetKeyspace,
						Keyspace:   "ks1",
					},
					nil,
				},
				{
					"schema change event table",
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 5, T, A, B, L, E,
						0, 3, k, s, _1,
						0, 6, t, a, b, l, e, _1,
					},
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetTable,
						Keyspace:   "ks1",
						Object:     "table1",
					},
					nil,
				},
				{
					"schema change event type",
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 4, T, Y, P, E,
						0, 3, k, s, _1,
						0, 4, u, d, t, _1,
					},
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetType,
						Keyspace:   "ks1",
						Object:     "udt1",
					},
					nil,
				},
				{
					"schema change event function",
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 8, F, U, N, C, T, I, O, N,
						0, 3, k, s, _1,
						0, 5, f, u, n, c, _1,
						0, 2,
						0, 3, i, n, t,
						0, 7, v, a, r, c, h, a, r,
					},
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetFunction,
						Keyspace:   "ks1",
						Object:     "func1",
						Arguments:  []string{"int", "varchar"},
					},
					nil,
				},
				{
					"schema change event aggregate",
					[]byte{
						0, 13, S, C, H, E, M, A, __, C, H, A, N, G, E,
						0, 7, C, R, E, A, T, E, D,
						0, 9, A, G, G, R, E, G, A, T, E,
						0, 3, k, s, _1,
						0, 4, a, g, g, _1,
						0, 2,
						0, 3, i, n, t,
						0, 7, v, a, r, c, h, a, r,
					},
					&SchemaChangeEvent{
						ChangeType: cassandraprotocol.SchemaChangeTypeCreated,
						Target:     cassandraprotocol.SchemaChangeTargetAggregate,
						Keyspace:   "ks1",
						Object:     "agg1",
						Arguments:  []string{"int", "varchar"},
					},
					nil,
				},
				{
					"status change event",
					[]byte{
						0, 13, S, T, A, T, U, S, __, C, H, A, N, G, E,
						0, 2, U, P,
						4, 192, 168, 1, 1,
						0, 0, 0x23, 0x52,
					},
					&StatusChangeEvent{
						ChangeType: cassandraprotocol.StatusChangeTypeUp,
						Address: &primitives.Inet{
							Addr: net.IPv4(192, 168, 1, 1),
							Port: 9042,
						},
					},
					nil,
				},
				{
					"topology change event",
					[]byte{
						0, 15, T, O, P, O, L, O, G, Y, __, C, H, A, N, G, E,
						0, 8, N, E, W, __, N, O, D, E,
						4, 192, 168, 1, 1,
						0, 0, 0x23, 0x52,
					},
					&TopologyChangeEvent{
						ChangeType: cassandraprotocol.TopologyChangeTypeNewNode,
						Address: &primitives.Inet{
							Addr: net.IPv4(192, 168, 1, 1),
							Port: 9042,
						},
					},
					nil,
				},
			}
			for _, tt := range tests {
				test.Run(tt.name, func(t *testing.T) {
					source := bytes.NewBuffer(tt.input)
					actual, err := codec.Decode(source, version)
					assert.Equal(t, tt.expected, actual)
					assert.Equal(t, tt.err, err)
				})
			}
		})
	}
}