package main

import (
	"fmt"
	"link-shortener/config"
	"link-shortener/entity/model"
	"link-shortener/internal/auth"
	"link-shortener/internal/link"
	"link-shortener/internal/stat"
	"link-shortener/internal/user"
	"link-shortener/pkg/bus"
	"link-shortener/pkg/database"
	"link-shortener/pkg/jwt"
	"link-shortener/pkg/logger"
	"link-shortener/pkg/middleware"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	logger.InitLogger()
	log := logger.GetWithScopes("MAIN")

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	eventBus := bus.New()

	conf := config.GetConfig()
	db := database.New(&conf.Db)
	db.Migrate(model.User{}, model.Link{}, model.Stat{})

	jwtService := jwt.NewJWT(conf.Auth.Secret)

	/** Repositories */
	linkRepository := link.NewRepository(db)
	userRepository := user.NewRepository(db)
	statRepository := stat.NewRepository(db)

	/** Services */
	userService := user.NewService(&user.ServiceDeps{
		Repository: userRepository,
	})
	authService := auth.NewService(&auth.ServiceDeps{
		UserService: userService,
	})
	linkService := link.NewLinkService(&link.ServiceDeps{
		Repository: linkRepository,
	})
	statService := stat.NewService(&stat.ServiceDeps{
		Repository: statRepository,
	})
	go statService.ListenClick(eventBus)

	/** Handlers */
	router := http.NewServeMux()
	auth.RegisterAuthHandler(&auth.HandlerDeps{
		AuthService: authService,
		Jwt:         jwtService,
		Router:      router,
		Config:      conf,
	})
	link.RegisterLinkHandler(&link.HandlerDeps{
		Router:      router,
		Service:     linkService,
		UserService: userService,
		Config:      conf,
		EventBus:    eventBus,
	})
	stat.RegisterStatHandler(&stat.HandlerDeps{
		Router:      router,
		Service:     statService,
		UserService: userService,
		LinkService: linkService,
		Jwt:         jwtService,
		Config:      conf,
	})

	middlewaresChain := middleware.Chain(
		middleware.Log,
	)

	appConfig := conf.App
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", appConfig.Port),
		Handler: middlewaresChain(router),
	}
	log.Info("Server is running on port ", appConfig.Port)
	log.Info("Server ended", server.ListenAndServe().Error())
}
