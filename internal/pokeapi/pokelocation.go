package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func(c *Client) ListLocationsPokemon(locationName string) (LocationPokemon, error) {
	url := baseURL + "/location-area/" + locationName
	val, ok :=  c.cache.Get(url)

	if ok {
		resp := LocationPokemon{}
		err := json.Unmarshal(val, &resp)

		if err != nil {
			return LocationPokemon{}, err
		}

		return resp, nil
	}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationPokemon{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationPokemon{}, err
	}

	pokemonLocationResp := LocationPokemon{}
	err = json.Unmarshal(data, &pokemonLocationResp)

	if err != nil {
		return LocationPokemon{}, err
	}

	c.cache.Add(url, data)

	return pokemonLocationResp, nil
}

