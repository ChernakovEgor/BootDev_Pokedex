package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const locationURL = "https://pokeapi.co/api/v2/location-area/"

func GetLocationsURL() string {
	return locationURL
}

func (c *Client) GetLocations(url string) (LocationReply, error) {
	if val, ok := c.cache.Get(url); ok {
		// log.Print("using cache data...")
		var rep LocationReply
		err := json.Unmarshal(val, &rep)
		if err != nil {
			log.Fatalf("error during cache decoding: %v", err)
			return LocationReply{}, err
		}
		return rep, nil
	}

	// log.Print("fetching fresh data...")
	res, err := http.Get(url)
	if err != nil {
		return LocationReply{}, err
	}
	defer res.Body.Close()

	var rep LocationReply
	b, _ := io.ReadAll(res.Body)
	c.cache.Add(url, b)
	// decoder := json.NewDecoder(res.Body)
	// err = decoder.Decode(&rep)
	err = json.Unmarshal(b, &rep)
	if err != nil {
		log.Fatalf("error during fresh data decoding: %v", err)
		return rep, err
	}

	return rep, nil
}
