package lib

import (
	"bufio"
	"io"
	"net/textproto"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

// SessionHandler is responsible for one client connection.
type SessionHandler struct {
	conn            io.ReadWriteCloser
	logger          *logrus.Entry
	store           *sync.Map
	auth            bool
	defaultPassword string
}

// NewSessionHandler builds a fully usable SessionHandler.
func NewSessionHandler(conn io.ReadWriteCloser, logger *logrus.Entry, store *sync.Map, auth bool, defaultPassword string) *SessionHandler {
	return &SessionHandler{
		conn:            conn,
		logger:          logger,
		store:           store,
		auth:            auth,
		defaultPassword: defaultPassword,
	}
}

// Handle takes care of existing connection
func (s *SessionHandler) Handle() {
	defer s.logger.Info("Closed connection")
	defer s.conn.Close()

	buffer := bufio.NewReader(s.conn)

	for {
		netData, err := textproto.NewReader(buffer).ReadLine()

		if err == io.EOF {
			break
		} else if err != nil {
			s.logger.Errorf("Could not read command: %v", err)
			break
		}

		input := strings.TrimSpace(string(netData))

		var result string
		result = s.ExecCommand(input) + "\n"

		s.conn.Write([]byte(string(result)))
	}
}

// ExecCommand executes redis commands
func (s *SessionHandler) ExecCommand(input string) string {
	args := strings.Split(input, " ")
	cmd := strings.ToLower(args[0])

	if !s.auth && cmd != "auth" {
		return "-NOAUTH Authentication required."
	}

	switch cmd {
	case "auth":
		return s.authenticate(args[1:])
	case "ping":
		return s.ping(args[1:])
	case "get":
		return s.get(args[1:])
	case "set":
		return s.set(args[1:])
	default:
		return "-ERR Command not found"
	}

}
