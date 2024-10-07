package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reliab-test/internal/config"
	"reliab-test/internal/infrastructure/datastore"
	"reliab-test/internal/infrastructure/log"
	"reliab-test/internal/repositories/users_repository"
)

func main() {
	defaultLog := log.InitDefaultLogger()

	cfg, err := config.Load(defaultLog)
	if err != nil {
		defaultLog.Error("Terminate execution", err)
	}

	logger := log.InitLogger()

	db := datastore.InitDB(cfg.DB, logger)

	userRepository := users_repository.BuildUserRepository(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		users, err := userRepository.GetUsers(c.Request.Context())
		if err != nil {
			logger.Error("Failed to get users", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	})
	r.Run()
}
