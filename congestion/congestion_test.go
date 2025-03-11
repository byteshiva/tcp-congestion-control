package congestion

import (
	"net"
	"testing"
)

func TestSendData(t *testing.T) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	cc := NewCongestionController(conn)
	err = cc.SendData([]byte("test data"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cc.inflight != 1 {
		t.Errorf("Expected inflight to be 1, got %d", cc.inflight)
	}
}

func TestHandleAck(t *testing.T) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	cc := NewCongestionController(conn)
	cc.inflight = 1
	cc.HandleAck()

	if cc.inflight != 0 {
		t.Errorf("Expected inflight to be 0, got %d", cc.inflight)
	}

	if cc.cwnd != 2 {
		t.Errorf("Expected cwnd to be 2, got %d", cc.cwnd)
	}
}

func TestHandleTimeout(t *testing.T) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	cc := NewCongestionController(conn)
	cc.cwnd = 10
	cc.HandleTimeout()

	if cc.cwnd != initialCWND {
		t.Errorf("Expected cwnd to be %d, got %d", initialCWND, cc.cwnd)
	}

	if cc.ssthresh != 5 {
		t.Errorf("Expected ssthresh to be 5, got %d", cc.ssthresh)
	}
}

