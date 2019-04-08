package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecCommand(t *testing.T) {
	var store sync.Map
	testConn := new(MockedReadWriteCloser)
	testHandler := sessionHandler{testConn, &store, false}

	args := "set hello world"
	assert.Equal(t, testHandler.ExecCommand(args), "-NOAUTH Authentication required.", "should return error unless user authenticated")

	testHandler.auth = true

	args = "unknown command"
	assert.Equal(t, testHandler.ExecCommand(args), "-ERR Command not found", "should return error if command not known")

	// args = "auth"
	// assert.Equal(t, testHandler.ExecCommand(args), "auth", "should invoke correct command")
	//
	// args = "set bob 3"
	// assert.Equal(t, testHandler.ExecCommand(args), "set", "should invoke correct command")
	//
	// args = "get bob"
	// assert.Equal(t, testHandler.ExecCommand(args), "get", "should invoke correct command")
	//
	// args = "ping"
	// assert.Equal(t, testHandler.ExecCommand(args), "ping", "should invoke correct command")
}
