package entity

type Game struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	State         string        `json:"state"`
	BoardSettings BoardSettings `json:"board_settings"`
	Board         Board         `json:"board"`
}

type BoardSettings struct {
	Size  uint `json:"size"`
	Bombs uint `json:"bombs"`
}

type Board [][]string

type GamesRepository interface {
	Get(id string) (Game, error)
	Save(Game) error
}

type GamesService interface {
	Get(id string) (Game, error)
	Create(name string, size uint, bombs uint) (Game, error)
	Reveal(id string, row uint, col uint) (Game, error)
}
