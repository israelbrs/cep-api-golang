package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type ViaCEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

func main() {
	cep := "01153000" // Exemplo de CEP

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	brasilAPIChan := make(chan *BrasilAPIResponse)
	viaCEPChan := make(chan *ViaCEPResponse)

	go fetchBrasilAPI(ctx, cep, brasilAPIChan)
	go fetchViaCEP(ctx, cep, viaCEPChan)

	select {
	case brasilAPIResp := <-brasilAPIChan:
		fmt.Println("Resposta da BrasilAPI:")
		fmt.Printf("CEP: %s\nEstado: %s\nCidade: %s\nBairro: %s\nRua: %s\n",
			brasilAPIResp.Cep, brasilAPIResp.State, brasilAPIResp.City,
			brasilAPIResp.Neighborhood, brasilAPIResp.Street)
	case viaCEPResp := <-viaCEPChan:
		fmt.Println("Resposta da ViaCEP:")
		fmt.Printf("CEP: %s\nEstado: %s\nCidade: %s\nBairro: %s\nRua: %s\n",
			viaCEPResp.Cep, viaCEPResp.Uf, viaCEPResp.Localidade,
			viaCEPResp.Bairro, viaCEPResp.Logradouro)
	case <-ctx.Done():
		fmt.Println("Erro: Timeout de 1 segundo excedido")
	}
}

func fetchBrasilAPI(ctx context.Context, cep string, ch chan<- *BrasilAPIResponse) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err := fetchWithContext(ctx, url)
	if err != nil {
		fmt.Printf("Erro ao buscar BrasilAPI: %v\n", err)
		ch <- nil
		return
	}

	var apiResp BrasilAPIResponse
	if err := json.Unmarshal(resp, &apiResp); err != nil {
		fmt.Printf("Erro ao decodificar resposta da BrasilAPI: %v\n", err)
		ch <- nil
		return
	}

	ch <- &apiResp
}

func fetchViaCEP(ctx context.Context, cep string, ch chan<- *ViaCEPResponse) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	resp, err := fetchWithContext(ctx, url)
	if err != nil {
		fmt.Printf("Erro ao buscar ViaCEP: %v\n", err)
		ch <- nil
		return
	}

	var apiResp ViaCEPResponse
	if err := json.Unmarshal(resp, &apiResp); err != nil {
		fmt.Printf("Erro ao decodificar resposta da ViaCEP: %v\n", err)
		ch <- nil
		return
	}

	ch <- &apiResp
}

func fetchWithContext(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
