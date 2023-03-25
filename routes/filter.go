package routes

import (
	"backEnd/handlers"
	"backEnd/pkg/mysql"
	"backEnd/repositories"

	"github.com/labstack/echo/v4"
)

func FilterRoutes(c *echo.Group) {
	filterRepository := repositories.RepositoryFilter(mysql.DB)
	h := handlers.HandlerFilter(filterRepository)
	// r.HandleFunc("/filtercategory", h.MultiFilter).Methods("GET")
	c.GET("/filtercategory", h.MultiFilter)

}
