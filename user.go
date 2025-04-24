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
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
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

func throwPokeball(pokemon pokemonData) bool {
	baseExp := pokemon.BaseExperience
	baseChance := int(math.Sqrt(float64(baseExp)))
	roll := rand.Intn(baseChance)
	fmt.Println("Roll: " + strconv.Itoa(roll) + " Max: " + strconv.Itoa(baseChance))
	return roll <= 6
}
