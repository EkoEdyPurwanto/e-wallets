package delivery

import (
	"EEP/e-wallets/config"
	"EEP/e-wallets/delivery/middleware"
	"EEP/e-wallets/manager"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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
}

func (s *Server) initControllers() {
	//rg := s.engine.Group("/api/v1")
	//api.NewUomController(s.ucManager.UomUseCase(), rg).Route()
	//api.NewProductController(s.ucManager.ProductUseCase(), rg).Route()
	//api.NewEmpolyeeController(s.ucManager.EmployeeUseCase(), rg).Route()
	//api.NewCustomerController(s.ucManager.CustomerUseCase(), rg).Route()
	//api.NewBillController(s.ucManager.BillUseCase(), rg).Route()
	//api.NewAuthController(s.ucManager.UserUseCase(), s.ucManager.AuthUseCase(), rg).Route()
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
	rm := manager.NewRepoManager(infraManager)
	ucm := manager.NewUseCaseManager(rm)

	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	log := logrus.New()

	engine := echo.New()
	return &Server{
		ucManager: ucm,
		engine:    engine,
		host:      host,
		log:       log,
	}
}
