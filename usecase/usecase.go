package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/igudgz/campo-minado/entity"
)

type service struct {
	gameRepository entity.GamesRepository
}

func NewService(gamerepository entity.GamesRepository) *service {
	return &service{
		gameRepository: gamerepository,
	}
}

func (s *service) Get(id string) (entity.Game, error) {
	game, err := s.gameRepository.Get(id)
	if err != nil {
		return entity.Game{}, errors.New("error when searching game")
	}

	return game, nil
}

func (s *service) Create(name string, size uint, bombs uint) (entity.Game, error) {
	if bombs >= size*size {
		return entity.Game{}, errors.New("the number of bombs is invalid")
	}

	id := uuid.New()

	game := entity.NewGame(id.String(), name, size, bombs)

	if err := s.gameRepository.Save(game); err != nil {
		return entity.Game{}, errors.New("create game into repository has failed")
	}

	game.Board = game.Board.HideBombs()

	return game, nil

}

func (srv *service) Reveal(id string, row uint, col uint) (entity.Game, error) {
	game, err := srv.gameRepository.Get(id)
	if err != nil {
		if errors.Is(err, errors.New("game not found in repository")) {
			return entity.Game{}, errors.New("game not found")
		}

		return entity.Game{}, errors.New("get game from repository has failed")
	}

	if !game.Board.IsValidPosition(row, col) {
		return entity.Game{}, errors.New("invalid position")
	}

	if game.IsOver() {
		return entity.Game{}, errors.New("game is over")
	}

	if game.Board.Contains(row, col, entity.CELL_BOMB) {
		game.State = entity.GAME_STATE_LOST
	} else {
		game.Board.Set(row, col, entity.CELL_REVEALED)

		if !game.Board.HasEmptyCells() {
			game.State = entity.GAME_STATE_WON
		}
	}

	if err := srv.gameRepository.Save(game); err != nil {
		return entity.Game{}, errors.New("update game into repository has failed")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}
