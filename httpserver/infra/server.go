package infra

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// Server .
type Server struct {
	port   int         //监听端口
	logger *zap.Logger //用于记录启动和关闭信息
	//以下部分为可选配置项
	handler      http.Handler  //路由映射
	readTimeout  time.Duration //读超时时间(ms)
	writeTimeout time.Duration //写超时时间(ms)
	idleTimeout  time.Duration //空闲等待超时时间(ms)
	waitShutdown time.Duration //优雅关闭时的最大等待时长(ms)
}

// NewServer 创建Server
func NewServer(port int, options ...func(*Server)) *Server {
	serv := &Server{
		port:         port,
		logger:       BaseLogger,
		handler:      nil,
		readTimeout:  5000 * time.Millisecond,
		writeTimeout: 5000 * time.Millisecond,
		idleTimeout:  5000 * time.Millisecond,
		waitShutdown: 10000 * time.Millisecond,
	}

	for _, option := range options {
		option(serv)
	}

	return serv
}

// Option Server可选配置项
type Option func(s *Server)

// ReadTimeout 设置ReadTimeout
func ReadTimeout(readTimeout time.Duration) Option {
	return func(s *Server) {
		s.readTimeout = readTimeout * time.Millisecond
	}
}

// WriteTimeout 设置WriteTimeout
func WriteTimeout(writeTimeout time.Duration) Option {
	return func(s *Server) {
		s.writeTimeout = writeTimeout * time.Millisecond
	}
}

// IdleTimeout 设置IdleTimeout
func IdleTimeout(idleTimeout time.Duration) Option {
	return func(s *Server) {
		s.idleTimeout = idleTimeout * time.Millisecond
	}
}

// WaitShutdown 设置WaitShutdown
func WaitShutdown(waitShutdown time.Duration) Option {
	return func(s *Server) {
		s.waitShutdown = waitShutdown * time.Millisecond
	}
}

// RegisterHandler 注册Handler
func (s *Server) RegisterHandler(handler http.Handler) {
	s.handler = handler
}

// Run 启动服务，Run必须在RegisterHandler之后执行
func (s *Server) Run() {
	s.logger.Info("application starts.", zap.Int("app_port", s.port))

	app := &http.Server{
		Addr:         ":" + strconv.Itoa(s.port),
		Handler:      s.handler,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
		IdleTimeout:  s.idleTimeout,
	}

	// 启动服务
	appError := make(chan error, 1)
	go func(serv *http.Server, e chan error) {
		e <- serv.ListenAndServe()
	}(app, appError)

	// 监听系统信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 优雅关闭
	select {
	case err := <-appError: // ErrServerClosed
		s.logger.Fatal("cannot start application server.", zap.Error(err))
	case sig := <-quit: // signal
		s.logger.Info("receive shutdown signal.", zap.String("type", sig.String()))
		servers := map[string]*http.Server{"app": app}
		s.shutdown(servers)
	}
}

// shutdown 优雅关闭
func (s *Server) shutdown(servers map[string]*http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), s.waitShutdown)
	defer cancel()

	start := time.Now()

	stopped := make(chan struct{})
	go func() {
		for name, serv := range servers {
			if err := serv.Shutdown(ctx); err != nil {
				s.logger.Error("server shutdown failed",
					zap.String("server", name), zap.NamedError("shutdown error", err))
				continue
			}
			s.logger.Info("server shutdown successfully", zap.String("server", name))
		}
		stopped <- struct{}{}
		close(stopped)
	}()

	select {
	case <-ctx.Done(): //到达截止时间
	case <-stopped: //所有server关闭完成
	}

	s.logger.Info("gracefully shutdown", zap.Duration("cost", time.Since(start)))
}
