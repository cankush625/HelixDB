package resp

import "HelixDB/common"

var DataTypeToFirstByteMap = map[string]string{
	"+": common.SimpleString,
	"-": common.SimpleError,
	":": common.Integer,
	"$": common.BulkString,
	"*": common.Array,
}

var FirstByteToDataTypeMap = common.ReverseMap(DataTypeToFirstByteMap)
