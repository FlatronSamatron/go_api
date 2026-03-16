package apiserver

import (
	"go_api/internal/app/store"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Apiserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

func New(config *Config) *Apiserver{
	//создаес конфиг сервера
	return &Apiserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Apiserver) Start() error{
	if err := s.configureLogger(); err != nil {
		return nil
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return nil
	}

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Apiserver) configureLogger() error{
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Apiserver) configureRouter(){
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *Apiserver) configureStore() error{
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return nil
	}

	s.store = st

	return nil
}

func (s *Apiserver) handleHello() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}