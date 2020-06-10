/**
 * Created by zc on 2020/6/7.
 */
package serve

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

type Config struct {
	Port   int    `json:"port"`
	Secret string `json:"secret"`
}

type Server struct {
	*http.Server
}

func New(cfg *Config) *Server {
	addr := ":" + strconv.Itoa(cfg.Port)
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
	return s.GracefulExitWeb()
}

func (s *Server) GracefulExitWeb() error {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch
	fmt.Println("\nServer shutdown: got a signal", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Shutdown(ctx)
}
