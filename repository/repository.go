package repository

import (
	"encoding/json"
	"errors"

	"github.com/igudgz/campo-minado/entity"
)

type repo struct {
	kvs map[string][]byte
}

func NewRepo() *repo {
	return &repo{kvs: map[string][]byte{}}
}

func (repo *repo) Get(id string) (entity.Game, error) {
	if value, ok := repo.kvs[id]; ok {
		game := entity.Game{}
		err := json.Unmarshal(value, &game)
		if err != nil {
			return entity.Game{}, errors.New("fail to get value from kvs")
		}

		return game, nil
	}

	return entity.Game{}, errors.New("game not found in kvs")
}

func (repo *repo) Save(game entity.Game) error {
	bytes, err := json.Marshal(game)
	if err != nil {
		return errors.New("game fails at marshal into json string")
	}

	repo.kvs[game.ID] = bytes

	return nil
}
