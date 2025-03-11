package congestion

import (
	"fmt"
	"net"
)

const (
	initialCWND = 1
	maxCWND     = 64
)

type CongestionController struct {
	conn     net.Conn
	cwnd     int
	ssthresh int
	inflight int
	lastAck  int
}

func NewCongestionController(conn net.Conn) *CongestionController {
	return &CongestionController{
		conn:     conn,
		cwnd:     initialCWND,
		ssthresh: maxCWND,
		inflight: 0,
		lastAck:  0,
	}
}

func (cc *CongestionController) SendData(data []byte) error {
	if cc.inflight >= cc.cwnd {
		return fmt.Errorf("congestion window full")
	}
	_, err := cc.conn.Write(data)
	if err != nil {
		return err
	}
	cc.inflight++
	return nil
}

func (cc *CongestionController) HandleAck() {
	cc.inflight--
	cc.lastAck++
	if cc.cwnd < cc.ssthresh {
		cc.cwnd++
	} else {
		cc.cwnd += 1 / cc.cwnd
	}
}

func (cc *CongestionController) HandleTimeout() {
	cc.ssthresh = cc.cwnd / 2
	if cc.ssthresh < initialCWND {
		cc.ssthresh = initialCWND
	}
	cc.cwnd = initialCWND
	cc.inflight = 0
}

