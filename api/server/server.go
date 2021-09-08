package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kzozulya1/graph_sql_api_test/config"
	"github.com/kzozulya1/graph_sql_api_test/internal/storage"

	"github.com/go-pg/pg/v10"
	"github.com/kelseyhightower/envconfig"

	// "github.com/labstack/echo/middleware"
	"github.com/graphql-go/handler"
	"github.com/kzozulya1/graph_sql_api_test/internal/schema"
	"github.com/sirupsen/logrus"
)

// Server contains instance details for the server
type Server struct {
	Cfg *config.Configuration
	DB  *pg.DB
}

// New returns a new instance of the server based on the specified configuration.
func New() *Server {
	var server = &Server{
		Cfg: &config.Configuration{},
	}

	server.initConfig()
	server.initDBConn()
	//server.useMiddleware()
	server.initRoutes()
	return server
}

// initRouter inits routers
func (s *Server) initRoutes() {
	logrus.Infof("http server: init routes...")

	log.Println("Listening at 0.0.0.0:2020")

	//Init client GraphSQL schema
	http.Handle("/client", handler.New(&handler.Config{
		Schema:   schema.GetClientSchema(s.DB),
		Pretty:   true,
		GraphiQL: true,
	}))
}

// initConfig inits config
func (s *Server) initConfig() {
	if err := envconfig.Process("", s.Cfg); err != nil {
		panic(fmt.Errorf("load configuration err: %s", err.Error()))
	}
}

// initServer initialized HTTP server
func (s *Server) initDBConn() {
	var err error
	logrus.Infof("http server: init SDB conn...")
	s.DB, err = storage.InitDB(s.Cfg.DBConn, s.Cfg.DBSQLQueryLog)
	if err != nil {
		panic(fmt.Errorf("db init err: %s", err.Error()))
	}
}

// initRouter inits routers
// func (s *Server) useMiddleware() {
// 	logrus.Infof("http server: use CORS middleware...")

// 	s.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins: []string{"*"},
// 		AllowMethods: []string{"GET", "HEAD", "PUT", "POST", "DELETE", "OPTIONS"},
// 		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
// 	}))
// }

// Run starts service
func (s *Server) Run() error {
	logrus.Infof("starting HTTP service listen with %s", s.Cfg.ListenAddr)
	return http.ListenAndServe(s.Cfg.ListenAddr, nil)
}

// Shutdown() gracefully stops serving
func (s *Server) Shutdown() error {
	logrus.Infof("http server: shutdown...")

	if err := s.closeDB(); err != nil {
		return fmt.Errorf("close db err: %s", err.Error())
	}

	return nil
}

// closeDB closes DB conns
func (s *Server) closeDB() error {
	if s.DB != nil {
		if err := s.DB.Close(); err != nil {
			return err
		}
	}
	return nil
}
