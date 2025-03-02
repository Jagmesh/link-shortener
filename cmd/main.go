package main

import (
	"fmt"
	"link-shortener/config"
	"link-shortener/internal/auth"
	"link-shortener/internal/link"
	"link-shortener/internal/user"
	"link-shortener/pkg/database"
	"link-shortener/pkg/jwt"
	"link-shortener/pkg/logger"
	"link-shortener/pkg/middleware"
	"net/http"
)

func main() {
	log := logger.InitLogger()

	conf := config.GetConfig()
	db := database.New(&conf.Db)
	db.Migrate(link.Link{}, user.User{})

	/** Repositories */
	linkRepository := link.NewRepository(db)
	userRepository := user.NewRepository(db)

	/** Services */
	userService := user.NewService(&user.UserServiceDeps{
		Repository: userRepository,
	})
	authService := auth.NewService(&auth.AuthServiceDeps{
		UserSerive: userService,
	})
	linkService := link.NewLinkService(link.LinkServiceDeps{
		Repository: linkRepository,
	})

	/** Handlers */
	router := http.NewServeMux()
	auth.ReqisterAuthHandler(auth.AuthHandlerDeps{
		AuthService: authService,
		Jwt:         jwt.NewJWT("aaa"),
		Router:      router,
		Config:      conf,
	})
	link.RegisterLinkHandler(link.LinkHandlerDeps{
		Router:  router,
		Service: linkService,
	})

	appConfig := conf.App
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", appConfig.Port),
		Handler: middleware.Log(router),
	}
	log.Info("Server is running on port ", appConfig.Port)
	server.ListenAndServe()
}
