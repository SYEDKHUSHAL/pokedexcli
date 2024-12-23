package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


func (c *Client) ListLocations(pageURL *string) (Location, error) {
	url := baseURL + "/location-area"


	if pageURL != nil {
		url = *pageURL
	}

	val, ok :=  c.cache.Get(url)

	if ok {
		resp := Location{}
		err := json.Unmarshal(val, &resp)

		if err != nil {
			return Location{}, err
		}

		return resp, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}

	err = json.Unmarshal(data, &locationResp)

	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return locationResp, nil


}
