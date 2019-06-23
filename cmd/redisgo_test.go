package main

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				main()
			}
		}
	}()

	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "localhost:3001")
	assert.Nil(t, err)

	_, err = conn.Write([]byte("ping\n"))
	assert.Nil(t, err)

	out := make([]byte, 1024)
	_, err = conn.Read(out)
	assert.Nil(t, err)

	conn.Close()

	close(quit)
}
