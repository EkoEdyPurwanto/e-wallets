package delivery

import (
	"EEP/e-wallets/config"
	"EEP/e-wallets/delivery/controller"
	"EEP/e-wallets/delivery/middleware"
	"EEP/e-wallets/manager"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *echo.Echo
	host      string
	log       *logrus.Logger
}

func (s *Server) Run() {
	s.initMiddlewares()
	s.initControllers()
	err := s.engine.Start(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddlewares() {
	s.engine.Use(middleware.LogRequest(s.log))
	s.engine.Use(middleware.RateLimiter())
}

func (s *Server) initControllers() {
	controller.NewUsersAccountController(s.ucManager.UsersAccountUC(), s.ucManager.WalletsUC(), s.engine).AuthRoute()
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Println(err)
	}

	// Instance Repo
	rm := manager.NewRepoManager(infraManager)

	// Instance UC
	ucm := manager.NewUseCaseManager(rm)

	hostAndPort := fmt.Sprintf("%s:%s", viper.GetString("APP_API_HOST"), viper.GetString("APP_API_PORT"))
	log := logrus.New()

	engine := echo.New()
	return &Server{
		ucManager: ucm,
		engine:    engine,
		host:      hostAndPort,
		log:       log,
	}
}
