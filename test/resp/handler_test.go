package resp

import (
	"HelixDB/resp"
	"io"
	"net"
	"testing"
)

func TestHandleConn(t *testing.T) {
	r, w := net.Pipe()
	w.Write([]byte("+PING\r\n"))
	resp.HandleConn(w)
	b, err := io.ReadAll(r)
	if err != nil {
		t.Errorf(err.Error())
	}
	if string(b) != "PONG" {
		t.Errorf("Invalid output %v", b)
	}
}
