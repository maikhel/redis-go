package lib

import (
	"bytes"
	"sync"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type MockedReadWriteCloser struct {
	*bytes.Buffer
}

func (b MockedReadWriteCloser) Close() error {
	return nil
}

func TestAuthenticateCommand(t *testing.T) {
	var store sync.Map
	testConn := new(MockedReadWriteCloser)
	log := logrus.New()
	logger := log.WithField("test", true)
	testHandler := NewSessionHandler(testConn, logger, &store, false, "bacon")

	args := []string{"one", "two"}
	assert.Equal(t, testHandler.authenticate(args), "-ERR Wrong number of arguments", "should return error if too many arguments")

	args = []string{}
	assert.Equal(t, testHandler.authenticate(args), "-ERR Wrong number of arguments", "should return error if too few arguments")

	args = []string{"pass"}
	testHandler.auth = true
	assert.Equal(t, testHandler.authenticate(args), "-ERR Already authenticated", "should return error if already authenticated")

	args = []string{"pass"}
	testHandler = NewSessionHandler(testConn, logger, &store, false, "bacon")
	assert.Equal(t, testHandler.authenticate(args), "-ERR Invalid password", "should return error if wrong password")

	args = []string{"bacon"}
	testHandler = NewSessionHandler(testConn, logger, &store, false, "bacon")
	assert.Equal(t, testHandler.authenticate(args), "+OK", "should authenticate if valid password")
	assert.Equal(t, testHandler.auth, true)
}

func TestPingCommand(t *testing.T) {
	var store sync.Map
	log := logrus.New()
	logger := log.WithField("test", true)
	testConn := new(MockedReadWriteCloser)
	testHandler := NewSessionHandler(testConn, logger, &store, false, "bacon")

	args := []string{}
	assert.Equal(t, testHandler.ping(args), "PONG", "should return PONG if no arguments")

	args = []string{"hello", "world"}
	assert.Equal(t, testHandler.ping(args), "hello world", "should return all arguments if they are present")

	args = []string{"42"}
	assert.Equal(t, testHandler.ping(args), "42", "should return all arguments if they are present")
}

func TestGetCommand(t *testing.T) {
	var store sync.Map
	log := logrus.New()
	logger := log.WithField("test", true)
	testConn := new(MockedReadWriteCloser)
	testHandler := NewSessionHandler(testConn, logger, &store, false, "bacon")

	args := []string{"one", "two"}
	assert.Equal(t, testHandler.get(args), "-ERR Wrong number of arguments", "should return error if too many arguments")

	args = []string{}
	assert.Equal(t, testHandler.get(args), "-ERR Wrong number of arguments", "should return error if too few arguments")

	args = []string{"hello"}
	store.Store("hello", "world")
	assert.Equal(t, testHandler.get(args), "world", "should return valid value for given key")

	args = []string{"hey"}
	assert.Equal(t, testHandler.get(args), "(nil)", "should return (nil) if value for key not found")
}

func TestSetCommand(t *testing.T) {
	var store sync.Map
	log := logrus.New()
	logger := log.WithField("test", true)
	testConn := new(MockedReadWriteCloser)
	testHandler := NewSessionHandler(testConn, logger, &store, false, "bacon")

	args := []string{"one", "two", "three"}
	assert.Equal(t, testHandler.set(args), "-ERR Wrong number of arguments", "should return error if too many arguments")

	args = []string{"one"}
	assert.Equal(t, testHandler.set(args), "-ERR Wrong number of arguments", "should return error if too few arguments")

	args = []string{"hello", "world"}
	assert.Equal(t, testHandler.set(args), "+OK", "should return OK if value set")
}
