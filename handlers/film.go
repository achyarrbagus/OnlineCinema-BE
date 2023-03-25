package handlers

import (
	filmdto "backEnd/dto/film"
	dto "backEnd/dto/result"
	"backEnd/models"
	"backEnd/repositories"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerFilm struct {
	FilmRepository repositories.FilmRepository
}

func HandlerFilm(FilmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

func (h *handlerFilm) CreateFilm(c echo.Context) error {

	// get data file
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)
	//
	filmFile := c.Get("filmFile").(string)
	fmt.Println("this is film file", filmFile)
	//
	heroFile := c.Get("heroFile").(string)
	fmt.Println("this is hero file", heroFile)
	//
	trailerFile := c.Get("trailerFile").(string)
	fmt.Println("this is trailer file", trailerFile)

	// get jwt tokens witk key "id"

	// category_id, _ := strconv.Atoi(c.FormValue("category_id"))
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	fmt.Println(CLOUD_NAME)
	var API_KEY = os.Getenv("API_KEY")
	fmt.Println(API_KEY)
	var API_SECRET = os.Getenv("API_SECRET")
	fmt.Println(API_SECRET)
	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "CinemaOnline"})
	respFilm, err2 := cld.Upload.Upload(ctx, filmFile, uploader.UploadParams{Folder: "CinemaOnline"})
	respHero, err3 := cld.Upload.Upload(ctx, heroFile, uploader.UploadParams{Folder: "CinemaOnline"})
	respTrailer, err4 := cld.Upload.Upload(ctx, trailerFile, uploader.UploadParams{Folder: "CinemaOnline"})

	if err != nil {
		fmt.Println(err.Error())
	}
	if err2 != nil {
		fmt.Println(err.Error())
	}
	if err3 != nil {
		fmt.Println(err.Error())
	}
	if err4 != nil {
		fmt.Println(err.Error())
	}

	price, _ := strconv.Atoi(c.FormValue("price"))

	request := filmdto.FilmRequest{
		Title:       c.FormValue("name"),
		Price:       price,
		Category:    c.FormValue("category"),
		Description: c.FormValue("desc"),
		Thumbnail:   resp.SecureURL,
		Hero:        respHero.SecureURL,
		FilmUrl:     respFilm.SecureURL,
		Trailer:     respTrailer.SecureURL,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	film := models.Film{
		Title:       request.Title,
		Price:       request.Price,
		Category:    request.Category,
		Description: request.Description,
		FilmUrl:     request.FilmUrl,
		Thumbnail:   request.Thumbnail,
		Hero:        request.Hero,
		Trailer:     request.Trailer,
	}

	data, err := h.FilmRepository.CreateFilm(film)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})

	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

}

func (h *handlerFilm) UpdateFilm(c echo.Context) error {
	//geting id from params
	id, _ := strconv.Atoi(c.Param("id"))
	// film, _ := h.FilmRepository.GetFilm(id)

	// get data file
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)
	//
	filmFile := c.Get("filmFile").(string)
	fmt.Println("this is film file", filmFile)
	//
	heroFile := c.Get("heroFile").(string)
	fmt.Println("this is hero file", heroFile)
	//
	trailerFile := c.Get("trailerFile").(string)
	fmt.Println("this is trailer file", trailerFile)
	//
	// category_id, _ := strconv.Atoi(c.FormValue("category_id"))
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")
	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "CinemaOnline"})
	respFilm, err2 := cld.Upload.Upload(ctx, filmFile, uploader.UploadParams{Folder: "CinemaOnline"})
	respHero, err3 := cld.Upload.Upload(ctx, heroFile, uploader.UploadParams{Folder: "CinemaOnline"})
	respTrailer, err4 := cld.Upload.Upload(ctx, trailerFile, uploader.UploadParams{Folder: "CinemaOnline"})
	//
	if err != nil {
		fmt.Println(err.Error())
	}
	if err2 != nil {
		fmt.Println(err.Error())
	}
	if err3 != nil {
		fmt.Println(err.Error())
	}
	if err4 != nil {
		fmt.Println(err.Error())
	}

	//getting value from form file
	title := c.FormValue("name")
	price, _ := strconv.Atoi(c.FormValue("price"))
	category := c.FormValue("category")
	description := c.FormValue("desc")
	thumbnail := resp.SecureURL
	hero := respHero.SecureURL
	trailer := respTrailer.SecureURL
	filmurl := respFilm.SecureURL
	//

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	filmUpdate := models.Film{
		ID:          id,
		Title:       title,
		Category:    category,
		Description: description,
		Price:       price,
		FilmUrl:     filmurl,
		Thumbnail:   thumbnail,
		Trailer:     trailer,
		Hero:        hero,
	}

	data, err := h.FilmRepository.UpdateFilm(filmUpdate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerFilm) DeleteFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param(("id")))

	film, err := h.FilmRepository.GetFilm(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	data, err := h.FilmRepository.DeleteFilm(film)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerFilm) FindFilm(c echo.Context) error {
	film, err := h.FilmRepository.FindFilm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: film})
}

func (h *handlerFilm) GetFilm(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: film})
}

func (h *handlerFilm) GetUserFilm(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	userInfo := c.Get("userLogin").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	film, err := h.FilmRepository.GetOneFilm(id, userId)

	if err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := dto.SuccessResult{Code: http.StatusOK, Data: film}
	return c.JSON(http.StatusOK, response)

}
