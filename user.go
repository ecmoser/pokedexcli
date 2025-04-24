package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	pokecache "github.com/ecmoser/pokedexcli/internal/pokecache"
)

type pokemonData struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Types          []pokeTypes   `json:"types"`
	Stats          []pokemonStat `json:"stats"`
}

type pokeTypes struct {
	Slot int      `json:"slot"`
	Type pokeType `json:"type"`
}

type pokeType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type pokemonStat struct {
	Base   int  `json:"base_stat"`
	Effort int  `json:"effort"`
	Stat   stat `json:"stat"`
}

type stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var pokemonList = make(map[string]pokemonData)

func CommandCatch(c *config, cache *pokecache.Cache, args ...string) error {
	pokemonName := strings.ToLower(args[0])
	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	req, err := http.Get(url)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	var pokemon pokemonData
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&pokemon); err != nil {
		return err
	}

	if throwPokeball(pokemon) {
		fmt.Println(pokemon.Name + " was caught!")
		pokemonList[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name + " escaped!")
	}

	return nil
}

func CommandInspect(c *config, cache *pokecache.Cache, args ...string) error {
	pokemonName := strings.ToLower(args[0])
	if _, exists := pokemonList[pokemonName]; !exists {
		return fmt.Errorf("%s is not in your Pokedex", pokemonName)
	}

	pokemon := pokemonList[pokemonName]

	output := "\n"
	output += fmt.Sprintf("Name: %s\n", pokemon.Name)
	output += "Height: " + strconv.Itoa(pokemon.Height) + "\n"
	output += "Weight: " + strconv.Itoa(pokemon.Weight) + "\n"
	output += "Stats:\n"
	for _, stat := range pokemon.Stats {
		output += " - " + stat.Stat.Name + ": " + strconv.Itoa(stat.Base) + "\n"
	}
	output += "Types:\n"
	for _, pokeType := range pokemon.Types {
		output += " - " + pokeType.Type.Name + "\n"
	}

	fmt.Println(output)
	return nil
}

func CommandPokedex(c *config, cache *pokecache.Cache, args ...string) error {
	output := "\nYour Pokedex:\n"
	for pokemonName := range pokemonList {
		output += " - " + pokemonName + "\n"
	}
	fmt.Println(output)
	return nil
}

func throwPokeball(pokemon pokemonData) bool {
	baseExp := pokemon.BaseExperience
	baseChance := int(math.Sqrt(float64(baseExp)))
	roll := rand.Intn(baseChance)
	fmt.Println("Roll: " + strconv.Itoa(roll) + " Max: " + strconv.Itoa(baseChance))
	return roll <= 6
}
