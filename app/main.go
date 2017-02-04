package server

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/line-bot/motivating_menu/api"
	"github.com/line-bot/motivating_menu/db"
)

type Server struct {
	db   *sql.DB
	echo *echo.Echo
}

func New() *Server {
	return &Server{}
}

func (s *Server) Setup(dbconf, env string) {

	cs, err := db.NewConfigsFromFile(dbconf)
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	s.db, err = cs.Open(env)
	if err != nil {
		log.Fatalf("db initialization failed: %s", err)
	}
}

func (s *Server) Run() {

	api := &api.Request{DB: s.db}

	s.echo = echo.New()

	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.CORS())

	s.echo.GET("/callback", api.Line)
	s.echo.GET("/api/get", api.GetTest)
	s.echo.GET("/test", api.Test)
	s.echo.POST("/api/post", api.PostTest)
	s.echo.PUT("/api/put", api.PutTest)
	s.echo.DELETE("/api/delete", api.DeleteTest)

	s.echo.Pre(middleware.RemoveTrailingSlash())
	http.Handle("/", s.echo)
}

func init() {
	var (
		dbconf = flag.String("dbconf", "dbconfig.yml", "database configuration file.")
		env    = flag.String("env", "development", "application envirionment (production, development etc.)")
	)

	flag.Parse()
	s := New()
	s.Setup(*dbconf, *env)
	s.Run()
}