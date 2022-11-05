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
