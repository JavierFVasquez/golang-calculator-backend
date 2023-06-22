package random_string_repositories

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog"
)

type RandomStringRepository struct {
	logger *zerolog.Logger
}

func NewRandomStringRepository(log *zerolog.Logger) *RandomStringRepository {
	l := log.With().Str("resource", "random_string_repository").Logger()

	return &RandomStringRepository{
		logger: &l,
	}
}

func (repo *RandomStringRepository) GetRandomString(ctx context.Context) (*string, error) {

	uri := "https://www.random.org/strings/?num=1&len=32&digits=on&upperalpha=on&loweralpha=on&unique=on&format=plain&rnd=new"
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return nil, err
	}

	responseString := string(body)

	return &responseString, nil
}
