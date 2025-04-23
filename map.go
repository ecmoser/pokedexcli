package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	pokecache "github.com/ecmoser/pokedexcli/internal/pokecache"
)

type locationsRequest struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []location `json:"results"`
}

type locationInfo struct {
	Location   location    `json:"location"`
	Encounters []encounter `json:"pokemon_encounters"`
}
type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type encounter struct {
	Pokemon pokemon `json:"pokemon"`
}

type pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func CommandMap(c *config, cache *pokecache.Cache, args ...string) error {
	output := ""
	reqURL := c.next

	val, exists := cache.Get(reqURL)

	if exists {
		fmt.Println(string(val))
		return nil
	}

	req, err := http.Get(reqURL)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	var locsRequest locationsRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&locsRequest); err != nil {
		return err
	}
	c.previous = locsRequest.Previous
	c.next = locsRequest.Next

	if locsRequest.Count != 0 {
		for _, loc := range locsRequest.Results {
			output += loc.Name + "\n"
		}
	}

	cache.Add(reqURL, []byte(output))

	fmt.Println(output)
	return nil
}

func CommandMapb(c *config, cache *pokecache.Cache, args ...string) error {
	output := ""
	reqURL := c.previous

	val, exists := cache.Get(reqURL)

	if exists {
		fmt.Println(string(val))
		return nil
	}

	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	req, err := http.Get(reqURL)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	var locsRequest locationsRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&locsRequest); err != nil {
		return err
	}

	c.previous = locsRequest.Previous
	c.next = locsRequest.Next

	if locsRequest.Count != 0 {
		for _, loc := range locsRequest.Results {
			output += loc.Name + "\n"
		}
	}

	cache.Add(reqURL, []byte(output))

	fmt.Println(output)
	return nil
}

func CommandExplore(c *config, cache *pokecache.Cache, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	output := ""

	val, exists := cache.Get(url)

	if exists {
		fmt.Println(string(val))
		return nil
	}

	req, err := http.Get(url)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	var locInfo locationInfo
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&locInfo); err != nil {
		return err
	}

	output += "Exploring " + locInfo.Location.Name + "...\n"
	output += "Found Pokemon:"

	for _, encounter := range locInfo.Encounters {
		output += "\n - " + encounter.Pokemon.Name
	}

	cache.Add(url, []byte(output))

	fmt.Println(output)
	return nil
}
