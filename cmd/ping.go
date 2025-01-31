package cmd

func Ping(command []byte) ([]byte, error) {
	return []byte("+PONG\r\n"), nil
}
