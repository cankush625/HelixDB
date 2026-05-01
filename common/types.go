package common

// Cmd represents a parsed client command.
// Name is the uppercased command name (e.g. "SET").
// Args contains the positional arguments following the command name.
type Cmd struct {
	Name string
	Args []string
}
