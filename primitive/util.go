package primitive

import (
	"fmt"
)

func AllProtocolVersions() []ProtocolVersion {
	return []ProtocolVersion{
		ProtocolVersion3,
		ProtocolVersion4,
		ProtocolVersion5,
		ProtocolVersionDse1,
		ProtocolVersionDse2,
	}
}

func AllOssProtocolVersions() []ProtocolVersion {
	return []ProtocolVersion{
		ProtocolVersion3,
		ProtocolVersion4,
		ProtocolVersion5,
	}
}

func AllDseProtocolVersions() []ProtocolVersion {
	return []ProtocolVersion{
		ProtocolVersionDse1,
		ProtocolVersionDse2,
	}
}

func AllBetaProtocolVersions() []ProtocolVersion {
	return []ProtocolVersion{
		ProtocolVersion5,
	}
}

func matchingProtocolVersions(filters ...func(version ProtocolVersion) bool) []ProtocolVersion {
	var result []ProtocolVersion
	for _, v := range AllProtocolVersions() {
		include := true
		for _, filter := range filters {
			if !filter(v) {
				include = false
				break
			}
		}
		if include {
			result = append(result, v)
		}
	}
	return result
}

func AllNonBetaProtocolVersions() []ProtocolVersion {
	return matchingProtocolVersions(func(v ProtocolVersion) bool { return !IsProtocolVersionBeta(v) })
}

func AllProtocolVersionsGreaterThanOrEqualTo(version ProtocolVersion) []ProtocolVersion {
	return matchingProtocolVersions(func(v ProtocolVersion) bool { return v >= version })
}

func AllProtocolVersionsGreaterThan(version ProtocolVersion) []ProtocolVersion {
	return matchingProtocolVersions(func(v ProtocolVersion) bool { return v > version })
}

func AllProtocolVersionsLesserThanOrEqualTo(version ProtocolVersion) []ProtocolVersion {
	return matchingProtocolVersions(func(v ProtocolVersion) bool { return v <= version })
}

func AllProtocolVersionsLesserThan(version ProtocolVersion) []ProtocolVersion {
	return matchingProtocolVersions(func(v ProtocolVersion) bool { return v < version })
}

func CheckProtocolVersion(version ProtocolVersion) error {
	for _, v := range AllProtocolVersions() {
		if v == version {
			return nil
		}
	}
	return fmt.Errorf("invalid protocol version: %v", version)
}

func IsProtocolVersion(version ProtocolVersion) bool {
	return CheckProtocolVersion(version) == nil
}

func CheckOssProtocolVersion(version ProtocolVersion) error {
	for _, v := range AllOssProtocolVersions() {
		if v == version {
			return nil
		}
	}
	return fmt.Errorf("invalid OSS protocol version: %v", version)
}

func IsOssProtocolVersion(version ProtocolVersion) bool {
	return CheckOssProtocolVersion(version) == nil
}

func CheckDseProtocolVersion(version ProtocolVersion) error {
	for _, v := range AllDseProtocolVersions() {
		if v == version {
			return nil
		}
	}
	return fmt.Errorf("invalid DSE protocol version: %v", version)
}

func IsDseProtocolVersion(version ProtocolVersion) bool {
	return CheckDseProtocolVersion(version) == nil
}

func CheckProtocolVersionBeta(version ProtocolVersion) error {
	for _, v := range AllBetaProtocolVersions() {
		if v == version {
			return nil
		}
	}
	return fmt.Errorf("protocol version is not beta: %v", version)
}

func IsProtocolVersionBeta(version ProtocolVersion) bool {
	return CheckProtocolVersionBeta(version) == nil
}

func CheckOpCode(code OpCode) error {
	switch code {
	case OpCodeStartup:
	case OpCodeOptions:
	case OpCodeQuery:
	case OpCodePrepare:
	case OpCodeExecute:
	case OpCodeRegister:
	case OpCodeBatch:
	case OpCodeAuthResponse:
	case OpCodeDseRevise:
	case OpCodeError:
	case OpCodeReady:
	case OpCodeAuthenticate:
	case OpCodeSupported:
	case OpCodeResult:
	case OpCodeEvent:
	case OpCodeAuthChallenge:
	case OpCodeAuthSuccess:
	default:
		return fmt.Errorf("invalid opcode: %v", code)
	}
	return nil
}

func IsOpCode(code OpCode) bool {
	return CheckOpCode(code) == nil
}

func CheckRequestOpCode(code OpCode) error {
	switch code {
	case OpCodeStartup:
	case OpCodeOptions:
	case OpCodeQuery:
	case OpCodePrepare:
	case OpCodeExecute:
	case OpCodeRegister:
	case OpCodeBatch:
	case OpCodeAuthResponse:
	case OpCodeDseRevise:
	default:
		return fmt.Errorf("invalid request opcode: %v", code)
	}
	return nil
}

func IsRequestOpCode(code OpCode) bool {
	return CheckRequestOpCode(code) == nil
}

func CheckResponseOpCode(code OpCode) error {
	switch code {
	case OpCodeError:
	case OpCodeReady:
	case OpCodeAuthenticate:
	case OpCodeSupported:
	case OpCodeResult:
	case OpCodeEvent:
	case OpCodeAuthChallenge:
	case OpCodeAuthSuccess:
	default:
		return fmt.Errorf("invalid response opcode: %v", code)
	}
	return nil
}

func IsResponseOpCode(code OpCode) bool {
	return CheckResponseOpCode(code) == nil
}

func CheckDseOpCode(code OpCode) error {
	switch code {
	case OpCodeDseRevise:
	default:
		return fmt.Errorf("invalid DSE opcode: %v", code)
	}
	return nil
}

func IsDseOpCode(code OpCode) bool {
	return CheckDseOpCode(code) == nil
}

func CheckConsistencyLevel(consistency ConsistencyLevel) error {
	switch consistency {
	case ConsistencyLevelAny:
	case ConsistencyLevelOne:
	case ConsistencyLevelTwo:
	case ConsistencyLevelThree:
	case ConsistencyLevelQuorum:
	case ConsistencyLevelAll:
	case ConsistencyLevelLocalQuorum:
	case ConsistencyLevelEachQuorum:
	case ConsistencyLevelSerial:
	case ConsistencyLevelLocalSerial:
	case ConsistencyLevelLocalOne:
	default:
		return fmt.Errorf("invalid consistency level: %v", consistency)
	}
	return nil
}

func IsConsistencyLevel(consistency ConsistencyLevel) bool {
	return CheckConsistencyLevel(consistency) == nil
}

func CheckNonSerialConsistencyLevel(consistency ConsistencyLevel) error {
	switch consistency {
	case ConsistencyLevelAny:
	case ConsistencyLevelOne:
	case ConsistencyLevelTwo:
	case ConsistencyLevelThree:
	case ConsistencyLevelQuorum:
	case ConsistencyLevelAll:
	case ConsistencyLevelLocalQuorum:
	case ConsistencyLevelEachQuorum:
	case ConsistencyLevelLocalOne:
	default:
		return fmt.Errorf("invalid non-serial consistency level: %v", consistency)
	}
	return nil
}

func IsNonSerialConsistencyLevel(consistency ConsistencyLevel) bool {
	return CheckNonSerialConsistencyLevel(consistency) == nil
}

func CheckSerialConsistencyLevel(consistency ConsistencyLevel) error {
	switch consistency {
	case ConsistencyLevelLocalSerial:
	case ConsistencyLevelSerial:
	default:
		return fmt.Errorf("invalid serial consistency level: %v", consistency)
	}
	return nil
}

func IsSerialConsistencyLevel(consistency ConsistencyLevel) bool {
	return CheckSerialConsistencyLevel(consistency) == nil
}

func CheckEventType(eventType EventType) error {
	switch eventType {
	case EventTypeSchemaChange:
	case EventTypeTopologyChange:
	case EventTypeStatusChange:
	default:
		return fmt.Errorf("invalid event type: %v", eventType)
	}
	return nil
}

func IsEventType(eventType EventType) bool {
	return CheckEventType(eventType) == nil
}

func CheckWriteType(writeType WriteType) error {
	switch writeType {
	case WriteTypeSimple:
	case WriteTypeBatch:
	case WriteTypeUnloggedBatch:
	case WriteTypeCounter:
	case WriteTypeBatchLog:
	case WriteTypeView:
	case WriteTypeCdc:
	default:
		return fmt.Errorf("invalid write type: %v", writeType)
	}
	return nil
}

func IsWriteType(writeType WriteType) bool {
	return CheckWriteType(writeType) == nil
}

func CheckBatchType(batchType BatchType) error {
	switch batchType {
	case BatchTypeLogged:
	case BatchTypeUnlogged:
	case BatchTypeCounter:
	default:
		return fmt.Errorf("invalid BATCH type: %v", batchType)
	}
	return nil
}

func IsBatchType(batchType BatchType) bool {
	return CheckBatchType(batchType) == nil
}

func CheckPrimitiveDataTypeCode(code DataTypeCode) error {
	switch code {
	case DataTypeCodeAscii:
	case DataTypeCodeBigint:
	case DataTypeCodeBlob:
	case DataTypeCodeBoolean:
	case DataTypeCodeCounter:
	case DataTypeCodeDecimal:
	case DataTypeCodeDouble:
	case DataTypeCodeFloat:
	case DataTypeCodeInt:
	case DataTypeCodeTimestamp:
	case DataTypeCodeUuid:
	case DataTypeCodeVarchar:
	case DataTypeCodeVarint:
	case DataTypeCodeTimeuuid:
	case DataTypeCodeInet:
	case DataTypeCodeDate:
	case DataTypeCodeTime:
	case DataTypeCodeSmallint:
	case DataTypeCodeTinyint:
	case DataTypeCodeDuration:
	default:
		return fmt.Errorf("invalid primitive data type code: %v", code)
	}
	return nil
}

func IsPrimitiveDataTypeCode(code DataTypeCode) bool {
	return CheckPrimitiveDataTypeCode(code) == nil
}

func CheckSchemaChangeType(t SchemaChangeType) error {
	switch t {
	case SchemaChangeTypeCreated:
	case SchemaChangeTypeUpdated:
	case SchemaChangeTypeDropped:
	default:
		return fmt.Errorf("invalid schema change type: %v", t)
	}
	return nil
}

func IsSchemaChangeType(t SchemaChangeType) bool {
	return CheckSchemaChangeType(t) == nil
}

func CheckSchemaChangeTarget(target SchemaChangeTarget) error {
	switch target {
	case SchemaChangeTargetKeyspace:
	case SchemaChangeTargetTable:
	case SchemaChangeTargetType:
	case SchemaChangeTargetFunction:
	case SchemaChangeTargetAggregate:
	default:
		return fmt.Errorf("invalid schema change target: %v", target)
	}
	return nil
}

func IsSchemaChangeTarget(target SchemaChangeTarget) bool {
	return CheckSchemaChangeTarget(target) == nil
}

func CheckStatusChangeType(t StatusChangeType) error {
	switch t {
	case StatusChangeTypeUp:
	case StatusChangeTypeDown:
	default:
		return fmt.Errorf("invalid status change type: %v", t)
	}
	return nil
}

func IsStatusChangeType(t StatusChangeType) bool {
	return CheckStatusChangeType(t) == nil
}

func CheckTopologyChangeType(t TopologyChangeType) error {
	switch t {
	case TopologyChangeTypeNewNode:
	case TopologyChangeTypeRemovedNode:
	default:
		return fmt.Errorf("invalid topology change type: %v", t)
	}
	return nil
}

func IsTopologyChangeType(t TopologyChangeType) bool {
	return CheckTopologyChangeType(t) == nil
}

func CheckResultType(t ResultType) error {
	switch t {
	case ResultTypeVoid:
	case ResultTypeRows:
	case ResultTypeSetKeyspace:
	case ResultTypePrepared:
	case ResultTypeSchemaChange:
	default:
		return fmt.Errorf("invalid result type: %v", t)
	}
	return nil
}

func IsResultType(t ResultType) bool {
	return CheckResultType(t) == nil
}

func CheckDseRevisionType(t DseRevisionType) error {
	switch t {
	case DseRevisionTypeCancelContinuousPaging:
	case DseRevisionTypeMoreContinuousPages:
	default:
		return fmt.Errorf("invalid DSE revision type: %v", t)
	}
	return nil
}

func IsDseRevisionType(t DseRevisionType) bool {
	return CheckDseRevisionType(t) == nil
}