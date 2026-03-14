package apiserver

import (
	"github.com/sirupsen/logrus"
)

type Apiserver struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *Apiserver{
	return &Apiserver{
		config: config,
		logger: logrus.New(),
	}
}

func (s *Apiserver) Start() error{
	if err := s.configureLogger(); err != nil {
		return nil
	}

	s.logger.Info("starting api server")

	return nil
}

func (s *Apiserver) configureLogger() error{
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}