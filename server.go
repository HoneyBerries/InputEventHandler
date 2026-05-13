package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"os"
	"sync"
)

type request struct {
	VirtualKey string `json:"virtual_key"`
	Duration   uint16 `json:"duration"`
}

type response struct {
	Success bool `json:"success"`
}

type server struct {
	port         int
	listener     net.Listener
	wg           sync.WaitGroup
	shutdownOnce sync.Once
	done         chan struct{}
}

func newServer(port int) *server {
	return &server{
		port: port,
		done: make(chan struct{}),
	}
}

func (s *server) start() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", s.port))
	if err != nil {
		return err
	}
	s.listener = ln
	slog.Info("listening", "addr", ln.Addr(), "pid", os.Getpid())

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-s.done:
				s.wg.Wait()
				return nil
			default:
				slog.Warn("accept error", "err", err)
				continue
			}
		}
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.handle(conn)
		}()
	}
}

func (s *server) handle(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr()
	slog.Info("connected", "addr", addr)

	sc := bufio.NewScanner(conn)
	w := bufio.NewWriter(conn)
	writeResponse := func(success bool) {
		b, _ := json.Marshal(response{Success: success})
		_, _ = w.Write(b)
		_ = w.WriteByte('\n')
		_ = w.Flush()
	}
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}

		var req request
		if err := json.Unmarshal(line, &req); err != nil {
			slog.Warn("bad request", "addr", addr, "err", err)
			writeResponse(false)
			continue
		}
		if req.VirtualKey == "" {
			slog.Warn("missing virtual_key", "addr", addr)
			writeResponse(false)
			continue
		}

		code, ok := lookupKey(req.VirtualKey)
		if !ok {
			slog.Warn("unknown key", "addr", addr, "key", req.VirtualKey)
			writeResponse(false)
			continue
		}

		err := tap(code, req.Duration)
		if err != nil {
			slog.Warn("tap failed", "addr", addr, "key", req.VirtualKey, "err", err)
		} else {
			slog.Info("tap", "key", req.VirtualKey, "duration_ms", req.Duration)
		}

		writeResponse(err == nil)
	}

	if err := sc.Err(); err != nil {
		slog.Info("disconnected", "addr", addr, "err", err)
	} else {
		slog.Info("disconnected", "addr", addr)
	}
}

func (s *server) shutdown() {
	s.shutdownOnce.Do(func() {
		close(s.done)
		s.listener.Close()
	})
}
