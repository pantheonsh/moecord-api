package moedex

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"strconv"
)

type Moemon struct {
	Id         int      `json:"id"`
	Number     string   `json:"num"`
	Name       string   `json:"name"`
	Type       []string `json:"type"`
	Weaknesses []string `json:"weaknesses"`
}

// Moemons Moemon data
var (
	Moemons []Moemon
)

func Load() {
	b, err := ioutil.ReadFile("moedex/pokedex.json")

	if err != nil {
		log.Fatalln("Failed to load Pokedex data", err)
	}

	err = json.Unmarshal(b, &Moemons)

	if err != nil {
		log.Fatalln("Failed to parse Pokedex data", err)
	}
}

func GetSprite(id int) ([]byte, error) {
	b, err := ioutil.ReadFile(path.Join("assets/sprites/", strconv.Itoa(id)+".png"))

	return b, err
}
