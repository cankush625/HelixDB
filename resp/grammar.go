package resp

import "HelixDB/common"

const SimpleString = "SIMPLE_STRING"
const SimpleError = "SIMPLE_ERROR"
const Integer = "INTEGER"
const BulkString = "BULK_STRING"
const Array = "ARRAY"

var DataTypeToFirstByteMap = map[string]string{
	"+": SimpleString,
	"-": SimpleError,
	":": Integer,
	"$": BulkString,
	"*": Array,
}

var FirstByteToDataTypeMap = common.ReverseMap(DataTypeToFirstByteMap)
