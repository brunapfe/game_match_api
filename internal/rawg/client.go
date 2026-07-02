package rawg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Jogo struct {
	ID      int      `json:"id"`
	Nome    string   `json:"name"`
	Generos []Genero `json:"genres"`
	Nota    float64  `json:"rating"`
	CapaURL string   `json:"background_image"`
}

type Genero struct {
	Nome string `json:"name"`
}

func Buscar(termo string) ([]Jogo, error) {
	apiKey := os.Getenv("RAWG_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("variável de ambiente RAW_API_KEY não configurada")
	}
	url := fmt.Sprintf("https://api.rawg.io/api/games?key=%s&search=%s&page_size=10", apiKey, termo)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao chamar RAWG: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("RAWG retornou status %d", resp.StatusCode)
	}

	var corpo struct {
		Results []Jogo `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&corpo); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta do RAWG: %w", err)
	}

	return corpo.Results, nil
}
