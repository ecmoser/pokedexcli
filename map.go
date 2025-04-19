package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationRequest struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []location `json:"results"`
}
type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func CommandMap(c *config) error {
	output := ""

	req, err := http.Get(c.next)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	var locsRequest locationRequest
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

	fmt.Println(output)
	return nil
}

func CommandMapb(c *config) error {
	output := ""

	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	req, err := http.Get(c.previous)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	var locsRequest locationRequest
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

	fmt.Println(output)
	return nil
}
