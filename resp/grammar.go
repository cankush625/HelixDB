package resp

const SimpleString = "SIMPLE_STRING"
const SimpleError = "SIMPLE_ERROR"
const Integer = "INTEGER"
const BulkString = "BULK_STRING"
const Array = "ARRAY"

var DataTypes = map[string]string{
	"+": SimpleString,
	"-": SimpleError,
	":": Integer,
	"$": BulkString,
	"*": Array,
}
