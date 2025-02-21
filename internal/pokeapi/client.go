package pokeapi

import (
	"fmt"
	"github.com/kravitty/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	fmt.Println("NewClient")
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(5 * time.Minute),
	}
}
