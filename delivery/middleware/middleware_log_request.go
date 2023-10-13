package middleware

import (
	"EEP/e-wallets/config"
	"EEP/e-wallets/model/dto/req"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func LogRequest(log *logrus.Logger) echo.MiddlewareFunc {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Membuka atau membuat file log jika tidak ada
	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			startTime := time.Now()

			err := next(c)

			endTime := time.Since(startTime)
			requestLog := req.LoggingRequest{
				StartTime:  startTime,
				EndTime:    endTime,
				StatusCode: c.Response().Status,
				ClientIP:   c.RealIP(),
				Method:     c.Request().Method,
				Path:       c.Request().URL.Path,
				UserAgent:  c.Request().UserAgent(),
			}

			switch {
			case c.Response().Status >= 500:
				log.Error(requestLog)
			case c.Response().Status >= 400:
				log.Warn(requestLog)
			default:
				log.Info(requestLog)
			}

			return err
		}
	}
}
