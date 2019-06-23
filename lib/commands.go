package lib

import (
	"strings"
)

func (s *SessionHandler) authenticate(args []string) string {
	if len(args) != 1 {
		return "-ERR Wrong number of arguments"
	}
	if s.auth {
		return "-ERR Already authenticated"
	}

	if args[0] == s.defaultPassword {
		s.auth = true
		return "+OK"
	}
	return "-ERR Invalid password"
}

func (s *SessionHandler) ping(args []string) string {
	response := "PONG"
	if len(args) > 0 {
		response = strings.Join(args, " ")
	}

	return response
}

func (s *SessionHandler) get(args []string) string {
	if len(args) != 1 {
		return "-ERR Wrong number of arguments"
	}
	result, ok := s.store.Load(args[0])
	if ok {
		return result.(string)
	}
	return "(nil)"
}

func (s *SessionHandler) set(args []string) string {
	if len(args) != 2 {
		return "-ERR Wrong number of arguments"
	}
	s.store.Store(args[0], args[1])
	return "+OK"
}
