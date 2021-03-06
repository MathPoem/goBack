package apiserver

import "github.com/sirupsen/logrus"

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// Start ...
func (s *APIServer) Start() error {
	return nil
}
