package main

import (
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := flag.Int("port", 6767, "TCP port to listen on")
	flag.Parse()

	if *port < 1 || *port > 65535 {
		slog.Error("invalid port", "port", *port)
		os.Exit(1)
	}

	srv := newServer(*port)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.start(); err != nil {
			slog.Error("server stopped", "err", err)
			os.Exit(1)
		}
	}()

	<-sigs
	srv.shutdown()
}
