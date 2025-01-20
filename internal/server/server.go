package server

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
	"syspeak/internal/http"
	"syspeak/internal/logging"
)

type Server struct {
	fiberApp *fiber.App
	logger   logging.Logger
}

// Start attempts to start up the server and all of its components.
//
// Returns an error if one occurs during startup or shutdown of the server. Note that Shutdown will only return an error
// that is caused during the shutdown process.
func (s *Server) Start() error {
	s.logger.Info("Starting server")
	killChan := make(chan os.Signal, 2)
	shutdownChan := make(chan error)
	signal.Notify(killChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-killChan
		s.logger.Trace("Received SIGTERM, shutting down")
		sErr := s.Shutdown()
		shutdownChan <- sErr
	}()
	s.logger.Trace("Starting http server: listening on port 3000")
	err := s.fiberApp.Listen("localhost:3000")
	sErr := <-shutdownChan
	if sErr != nil {
		s.logger.Error("Error occurred during server shutdown: " + sErr.Error())
	} else {
		s.logger.Info("Server shutdown gracefully")
	}
	return err
}

// Shutdown attempts a graceful shutdown of the server and its runtime services. Any errors that occur during shutdown
// are returned. Not that any error returned here will also show as the error from Start.
func (s *Server) Shutdown() error {
	if s.fiberApp.Server().GetOpenConnectionsCount() > 0 {
		s.logger.Info("Stopping http server, waiting for connections to close...")
	}
	sdErr := s.fiberApp.Shutdown()
	s.logger.Info("Http server shutdown")
	return sdErr
}

// Options provides optional parameters that can be used to customize aspects of the SysPeak server.
type Options struct {
	Logger logging.Logger
}

// NewDefaultServer provides a Server with default parameters.
func NewDefaultServer() Server {
	s := Server{
		logger:   logging.NewSysSpeakLogger("server"),
		fiberApp: http.NewConfiguredFiberApp(),
	}
	s.fiberApp.Get("/status", HandleGetStatus)
	return s
}

// NewServerWithOptions provides a new Server with the Options passed applied.
func NewServerWithOptions(options Options) Server {
	s := NewDefaultServer()
	if options.Logger != nil {
		s.logger = options.Logger
	}
	return s
}
