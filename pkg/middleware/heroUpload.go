package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HeroUploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		file, err := c.FormFile("hero")

		if file != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}

			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			defer src.Close()

			tempFile, err := ioutil.TempFile("uploads", "image-*.png")
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			defer tempFile.Close()

			if _, err = io.Copy(tempFile, src); err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}

			data := tempFile.Name()

			// filename := data[8:]

			c.Set("heroFile", data) // change this
			return next(c)
		}

		c.Set("heroFile", "")
		return next(c)
	}

}
