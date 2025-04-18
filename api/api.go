package api

import (
	"essy_travel/api/handler"
	"essy_travel/api/middleware"
	"essy_travel/config"
	"essy_travel/storage"

	"github.com/gin-gonic/gin"

	_ "essy_travel/api/docs"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Swagger JWT API
// @version 1.0
// @description Create  Go REST API with JWT Authentication in Gin Framework
// @contact.name API Support
// @termsOfService demo.com
// @contact.url http://demo.com/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath
// @Schemes http https
// @query.collection.format multi
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {
	// dst := fmt.Sprintf("%s/%s", "/home/samandarxon/Desktop/sam/Golang-Lessons/3 month/lesson_9/swagger_upload/api/upload", "1701524206 - your_air.json")
	// b, _ := os.ReadFile(dst)
	// fmt.Println("$$$$$$$$$$$$$$$$$", b, dst)
	handler := handler.NewHandler(cfg, strg)

	v1 := r.Group("/v1")

	v1.Use(middleware.CheckSecretKeyUserMiddleware())
	// City ...
	v1.POST("/city", handler.CreateCity)
	v1.GET("/city/:id", handler.CityGetById)
	v1.GET("/city", handler.CityGetList)
	v1.PUT("/city/:id", handler.CityUpdate)
	v1.DELETE("/city/:id", handler.CityDelete)

	// Country
	r.POST("/country", handler.CreateCountry)
	r.POST("/country/upload", handler.UploadCountry)
	r.GET("/country/:id", handler.CountryGetById)
	r.GET("/country", handler.CountryGetList)
	r.PUT("/country/:id", handler.CountryUpdate)
	r.DELETE("/country/:id", handler.CountryDelete)

	// Login
	r.POST("/login", handler.Login)

	// User
	r.POST("/user", middleware.CheckSecretKeyAdminMiddleware(), handler.CreateUser)
	r.GET("/user/:id", middleware.CheckSecretKeyAdminMiddleware(), handler.UserGetById)
	r.GET("/user", middleware.CheckSecretKeyAdminMiddleware(), handler.UserGetList)
	r.PUT("/user/:id", middleware.CheckSecretKeyAdminMiddleware(), handler.UserUpdate)
	r.DELETE("/user/:id", middleware.CheckSecretKeyAdminMiddleware(), handler.UserDelete)

	// Airport
	r.POST("/airport", handler.CreateAirport)
	r.GET("/airport/:id", handler.AirportGetById)
	r.GET("/airport", handler.AirportGetList)
	r.PUT("/airport/:id", handler.AirportUpdate)
	r.POST("/airport/upload", handler.UploadAirport)
	r.DELETE("/airport/:id", handler.AirportDelete)

	// Upload
	r.POST("/upload", handler.UploadFile)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
