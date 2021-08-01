package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kashyaprahul94/go-boilerplate/internal/logger"
)

func (s *HttpServer) GetRouter() HttpRouter {
	return s.router
}

// Boot prepares the web server and start listening on the web address
func (s *HttpServer) Boot(cb func()) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.host, s.port),
		Handler: s.GetRouter(),
	}

	logger.Info(fmt.Sprintf("Web server is running on http://localhost:%s", s.port))

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Error(fmt.Sprintf("Web server encountered error %s", err.Error()))
		}
	}()

	// Run the callback in a goroutine
	go cb()

	// prepare the webserver for graceful shutdown
	prepareForGracefulShutdown(server)
}

func prepareForGracefulShutdown(server *http.Server) {
	// Creat the channel for signal
	stop := make(chan os.Signal)

	// Observe the `SIGINT` signal
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Wait until we receive the message in the channel
	<-stop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Shutting down web server")
}
