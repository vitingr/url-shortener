package route

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vitingr/url-shortner/internal/service"
)

func SetupRouter(redisClient *redis.Client, pgClient *pgxpool.Pool) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/", func(c *gin.Context) { service.CreateNewTinyUrl(c, redisClient, pgClient) })
	r.GET("/:id", func(c *gin.Context) { service.GetUrl(c, redisClient, pgClient) })

	return r
}