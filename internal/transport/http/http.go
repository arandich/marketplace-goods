package http

import (
	"crypto/tls"
	"github.com/arandich/marketplace-goods/internal/config"
	"github.com/julienschmidt/httprouter"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
	"net/http/pprof"
)

type Server struct {
	httpServer *http.Server
	listener   net.Listener
	cfg        config.HttpConfig
}

func NewServer(listener net.Listener, cfg config.HttpConfig) *Server {
	httpServer := http.Server{
		Addr:              cfg.Address,
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	server := &Server{
		httpServer: &httpServer,
		listener:   listener,
		cfg:        cfg,
	}

	return server
}

func (s *Server) StartAndServe() chan error {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/metrics", promhttp.Handler())

	if s.cfg.ProfilingEnabled {
		router.HandlerFunc(http.MethodGet, "/debug/pprof/", pprof.Index)
		router.HandlerFunc(http.MethodGet, "/debug/pprof/cmdline", pprof.Cmdline)
		router.HandlerFunc(http.MethodGet, "/debug/pprof/profile", pprof.Profile)
		router.HandlerFunc(http.MethodGet, "/debug/pprof/symbol", pprof.Symbol)
		router.HandlerFunc(http.MethodGet, "/debug/pprof/trace", pprof.Trace)
		router.Handler(http.MethodGet, "/debug/pprof/heap", pprof.Handler("heap"))
		router.Handler(http.MethodGet, "/debug/pprof/goroutine", pprof.Handler("goroutine"))
		router.Handler(http.MethodGet, "/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
		router.Handler(http.MethodGet, "/debug/pprof/block", pprof.Handler("block"))
	}

	s.httpServer.Handler = router

	errChan := make(chan error, 1)
	go func() {
		errChan <- s.httpServer.Serve(s.listener)
	}()

	return errChan
}
