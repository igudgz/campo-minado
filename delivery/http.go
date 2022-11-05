package delivery

import (
	"net/http"

	"github.com/igudgz/campo-minado/entity"
	"github.com/igudgz/campo-minado/errors"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	gamesUsecase entity.GamesUsecase
}

func NewHTTPHandler(gamesUsecase entity.GamesUsecase) *HTTPHandler {
	return &HTTPHandler{
		gamesUsecase: gamesUsecase,
	}
}

func (hdl *HTTPHandler) Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "missing id!")
	}

	game, err := hdl.gamesUsecase.Get(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ResponseRequest{
			Message: err.Error(),
		})

	}

	return c.JSON(http.StatusOK, game)
}

func (hdl *HTTPHandler) Create(c echo.Context) error {
	type bodyCreate struct {
		Name  string
		Size  uint
		Bombs uint
	}

	body := bodyCreate{}
	err := c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")

	}

	game, err := hdl.gamesUsecase.Create(body.Name, body.Size, body.Bombs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ResponseRequest{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, game)
}

func (hdl *HTTPHandler) RevealCell(c echo.Context) error {
	type bodyRevealCell struct {
		Row uint
		Col uint
	}

	body := bodyRevealCell{}
	err := c.Bind(&body)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")

	}

	game, err := hdl.gamesUsecase.Reveal(c.Param("id"), body.Row, body.Col)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ResponseRequest{Message: err.Error()})

	}

	return c.JSON(http.StatusOK, game)
}
