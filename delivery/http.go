package delivery

import (
	"net/http"

	"github.com/igudgz/campo-minado/entity"
	"github.com/igudgz/campo-minado/errors"
	"github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	gamesService entity.GamesService
}

func NewHTTPHandler(gamesService entity.GamesService) *HTTPHandler {
	return &HTTPHandler{
		gamesService: gamesService,
	}
}

func (hdl *HTTPHandler) Get(c echo.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "missing id!")
	}

	game, err := hdl.gamesService.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ResponseRequest{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, game)
}

func (hdl *HTTPHandler) Create(c echo.Context) {
	type bodyCreate struct {
		Name  string
		Size  uint
		Bombs uint
	}

	body := bodyCreate{}
	err := c.Bind(&body)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	game, err := hdl.gamesService.Create(body.Name, body.Size, body.Bombs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ResponseRequest{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, game)
}

func (hdl *HTTPHandler) RevealCell(c echo.Context) {
	type bodyRevealCell struct {
		Row uint
		Col uint
	}

	body := bodyRevealCell{}
	err := c.Bind(&body)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	game, err := hdl.gamesService.Reveal(c.Param("id"), body.Row, body.Col)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ResponseRequest{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, game)
}
