package codec

import (
	"go-cassandra-native-protocol/cassandraprotocol"
	"go-cassandra-native-protocol/cassandraprotocol/message"
)

type RegisterCodec struct{}

func (c RegisterCodec) GetOpCode() cassandraprotocol.OpCode {
	return cassandraprotocol.OpCodeRegister
}

func (c RegisterCodec) Encode(msg message.Message, dest []byte, version cassandraprotocol.ProtocolVersion) error {
	register := msg.(*message.Register)
	_, err := WriteStringList(register.EventTypes, dest)
	return err
}

func (c RegisterCodec) EncodedSize(msg message.Message, version cassandraprotocol.ProtocolVersion) (int, error) {
	register := msg.(*message.Register)
	return SizeOfStringList(register.EventTypes), nil
}

func (c RegisterCodec) Decode(source []byte, version cassandraprotocol.ProtocolVersion) (message.Message, error) {
	eventTypes, _, err := ReadStringList(source)
	if err != nil {
		return nil, err
	}
	return message.NewRegister(eventTypes), nil
}
