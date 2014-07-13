package sockjs

/*
Cotains package internal types (not public)
*/

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

// Error variable
var ErrConnectionClosed = errors.New("Connection closed.")

type context struct {
	Config
	HandlerFunc
	connections
}

type conn struct {
	context
	input_channel    chan []byte
	output_channel   chan []byte
	quit_channel     chan bool
	timeout          time.Duration
	httpTransactions chan *httpTransaction
}

func newConn(ctx *context) *conn {
	return &conn{
		input_channel:    make(chan []byte),
		output_channel:   make(chan []byte, 64),
		quit_channel:     make(chan bool),
		httpTransactions: make(chan *httpTransaction),
		timeout:          time.Second * 30,
		context:          *ctx,
	}
}

func (c *conn) ReadMessage() ([]byte, error) {
	select {
	case <-c.quit_channel:
		return []byte{}, io.EOF
	case val := <-c.input_channel:
		if c.context.Config.DecodeFrames {
			// Decode the msg JSON
			var msg []string
			err := json.Unmarshal(val, &msg)
			if len(msg) == 1 && err == nil {
				val = []byte(msg[0])
			}
		} else {
			// Strip the [ and ] from the JSON and return a raw string
			val = val[1 : len(val)-1]
		}
		return val, nil
	}
	panic("unreachable")
}

func (c *conn) WriteMessage(val []byte) (count int, err error) {
	var data_out []byte
	if c.Config.DecodeFrames {
		data_out, err = json.Marshal(string(val))
		if err != nil {
			return
		}
	} else {
		data_out = append([]byte{}, val...)
	}
	select {
	case c.output_channel <- data_out:
	case <-time.After(c.timeout):
		return 0, ErrConnectionClosed
	case <-c.quit_channel:
		return 0, ErrConnectionClosed
	}
	return len(val), nil
}

func (c *conn) Close() (err error) {
	defer func() {
		if recover() != nil {
			err = ErrConnectionClosed
		}
	}()
	close(c.quit_channel)
	return
}

type connectionStateFn func(*conn) connectionStateFn

func (c *conn) run(cleanupFn func()) {
	for state := openConnectionState; state != nil; {
		state = state(c)
	}
	c.Close()
	cleanupFn()
}
