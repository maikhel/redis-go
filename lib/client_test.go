package lib

import (
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestExecCommand(t *testing.T) {
	var store sync.Map
	log := logrus.New()
	logger := log.WithField("test", true)
	testConn := new(MockedReadWriteCloser)
	testHandler := NewSessionHandler(testConn, logger, &store, false, "bacon")

	args := "set hello world"
	assert.Equal(t, testHandler.ExecCommand(args), "-NOAUTH Authentication required.", "should return error unless user authenticated")

	args = "auth bacon"
	assert.Equal(t, testHandler.ExecCommand(args), "+OK", "should authenticate")

	args = "ping hello"
	assert.Equal(t, testHandler.ExecCommand(args), "hello", "Should ping")

	args = "unknown command"
	assert.Equal(t, testHandler.ExecCommand(args), "-ERR Command not found", "should return error if command not known")

	args = "set bob 3"
	assert.Equal(t, testHandler.ExecCommand(args), "+OK", "Should set value")

	args = "get bob"
	assert.Equal(t, testHandler.ExecCommand(args), "3", "Should get value")
}
