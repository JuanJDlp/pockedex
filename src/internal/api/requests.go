package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"github.com/JuanJDlp/pockedex/src/internal/cache"
)

const baseUrl = "https://pokeapi.co/api/v2"

var (

	config = &configuration{
		nil,
		nil,
	}

	cacheStruct = cache.NewCache(60)
)

type configuration struct {
	prev *string
	next *string
}

func GetPokemon(pokemonName string) (Pokemon, error) {
	data, err := getData( baseUrl+"/pokemon/" + pokemonName)
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil

}

func getPrevOrNextLink(next bool) (string) {
	var value *string
	if next {
		value = config.next
	} else {
		value = config.prev
	}

	if value != nil {
		return *value
	}
	return baseUrl 

}

func GetLocations(next bool) (Locations, error) {
	value := getPrevOrNextLink(next)

	information, err := cacheStruct.Get(value)

	if err == nil {
		locations, err := unmarshalData(information)

		return locations, err

	}

	body, err := getData(value + "/location-area")

	if err != nil {
		return Locations{}, err
	}

	locations, err := unmarshalData(body)

	if err != nil {
		return Locations{}, err
	}

	cacheStruct.Add(value, body)

	return locations, err
}

func GetPokemonsFromLocations(area string) (LocationArea, error) {
	information, err := cacheStruct.Get(area)
	if err == nil {
		locationArea := LocationArea{}
		err = json.Unmarshal(information, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	body, err := getData(baseUrl +"/location-area/" + area)
	if err != nil {
		return LocationArea{}, err
	}
	locationArea := LocationArea{}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	cacheStruct.Add(area, body)
	return locationArea, nil
}

func unmarshalData(b []byte) (Locations, error) {

	locations := Locations{}
	err := json.Unmarshal(b, &locations)

	if err != nil { // error from unmarshaling
		return Locations{}, err
	}

	config.next = locations.Next
	config.prev = locations.Previous
	return locations, nil
}

func getData(url string) (data []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("an error ocurred when fetching the data")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil || res.StatusCode > 299 {
		return nil, errors.New("didnt get a 200 response")
	}

	return body, nil
}
