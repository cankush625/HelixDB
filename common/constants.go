package common

// Terminator for the RESP protocol
const Terminator = "\r\n"

// Commands
const Ping = "PING"
const Command = "COMMAND"
const Echo = "ECHO"
const Get = "GET"
const Set = "SET"
const Expire = "EXPIRE"
const Del = "DEL"
const Keys = "KEYS"
const Config = "CONFIG"
const Type = "TYPE"

// Command Args
// Expiration Args
const EX = "EX"
const PX = "PX"
const EXAT = "EXAT"
const PXAT = "PXAT"

// Conditional Set Args
const NX = "NX"
const XX = "XX"

// Other Set Args
const KEEPTTL = "KEEPTTL"
const GET = "GET"
