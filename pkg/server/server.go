/**
 * Created by zc on 2020/6/7.
 */
package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const DefaultPort = 9090

type Config struct {
	Host   string `json:"host" yaml:"host"`
	Port   int    `json:"port" yaml:"port"`
	Secret string `json:"secret" yaml:"secret"`
}

type Server struct {
	*http.Server
}

func New(cfg *Config) *Server {
	var port = DefaultPort
	if cfg.Port > 0 {
		port = cfg.Port
	}
	addr := ":" + strconv.Itoa(port)
	server := &http.Server{
		Addr:           addr,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &Server{Server: server}
}

func (s *Server) Run() error {
	go func() {
		fmt.Println("Server starting")
		if err := s.ListenAndServe(); err != nil {
			panic("服务监听异常：" + err.Error())
		}
	}()
	return s.Exit()
}

func (s *Server) Exit() error {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch
	fmt.Println("\nServer shutdown: got a signal", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Shutdown(ctx)
}
