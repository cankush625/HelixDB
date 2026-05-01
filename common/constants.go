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

// Command Args
// Expiration Args
const EX = "EX"
const PX = "PX"
