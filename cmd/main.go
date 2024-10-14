package main

import (
	"fmt"
	"github.com/TheAnonymousDarck/go-practica-7/cmd/database"
	"github.com/TheAnonymousDarck/go-practica-7/cmd/handlers"
	"github.com/TheAnonymousDarck/go-practica-7/cmd/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db, err := database.NewDatabaseDriver()
	if err != nil {
		fmt.Println("Error al conectar a la base de datos: ", err)
		return
	}

	// Iniciar el servicio de usuarios
	userService := services.NewUserService(db)
	userHandler := &handlers.UserHandler{
		UserService: userService,
	}

	// Configurar Gin
	router := gin.Default()
	fmt.Println("Running app")
	// Tomar archivos de la carpeta template
	router.LoadHTMLGlob("cmd/templates/*")
	router.Static("/css", "./cmd/static/css")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		users, err := userService.GetAllUsers()
		if err != nil {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"title":       "Main website",
				"error":       "Error al obtener usuarios",
				"total_users": 0,
				"users":       []database.User{},
			})
			return
		}
		c.HTML(200, "index.html", gin.H{
			"title":       "Main website",
			"total_users": len(users),
			"users":       users,
		})
	})

	// Rutas
	router.GET("/api/users", userHandler.GetAllUsers)
	router.GET("/api/users/:id", userHandler.GetUserByID)
	router.POST("/api/users", userHandler.CreateUser)
	router.PUT("/api/users/:id", userHandler.UpdateUser)
	router.DELETE("/api/users/:id", userHandler.DeleteUser)

	// Iniciar servidor
	err = router.Run(":8001")
	if err != nil {
		fmt.Println("Error al iniciar el servidor: ", err)
	}
}
