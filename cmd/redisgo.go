package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/maikhel/redis-go/lib"
	"github.com/sirupsen/logrus"
)

type specification struct {
	Port            int    `envconfig:"PORT" default:"6379"`
	DefaultPassword string `envconfig:"REDIS_AUTH_PASS" default:"bacon"`
}

func main() {
	var cfg specification
	envconfig.MustProcess("", &cfg)

	log := logrus.New()

	var store sync.Map

	port := fmt.Sprintf(":%d", cfg.Port)
	log.Infof("About to start serving on port %s", port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Could not set up listener: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Could not accept connection: %v", err)
		}

		logger := log.WithField("remote", conn.RemoteAddr())
		logger.Infoln("Accepted connection")
		go lib.NewSessionHandler(conn, logger, &store, false, cfg.DefaultPassword).Handle()
	}
}
