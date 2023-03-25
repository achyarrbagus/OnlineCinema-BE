package routes

import (
	"backEnd/handlers"
	"backEnd/pkg/middleware"
	"backEnd/pkg/mysql"
	"backEnd/repositories"

	"github.com/labstack/echo/v4"
)

func FilmRoutes(e *echo.Group) {
	FilmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(FilmRepository)

	e.POST("/film", middleware.Auth(middleware.UploadFile(middleware.FilmUploadFile(middleware.HeroUploadFile(middleware.TrailerUploadFile(h.CreateFilm))))))
	e.PATCH("/film/:id", middleware.Auth(middleware.UploadFile(middleware.FilmUploadFile(middleware.HeroUploadFile(middleware.TrailerUploadFile(h.UpdateFilm))))))
	// e.PATCH("/film/:id", middleware.Auth(middleware.UploadFile(h.UpdateFilm)))
	e.GET("/film", h.FindFilm)
	e.GET("/film/:id", h.GetFilm)
	e.GET("/userfilm/:id", middleware.Auth(h.GetUserFilm))
	e.DELETE("/film/:id", middleware.Auth(h.DeleteFilm))

}
