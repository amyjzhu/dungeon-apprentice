package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"path/filepath"
	"strconv"
)


type Monster struct {
	Name string
	Alignment string
	Challenge_Rating string // challenge rating
	env []Environment // places usually found
	exp int // experience points award for defeat
	Hit_Dice string // TODO: convert all these string values to hit dice thing
	//Stats BaseStats
	Strength int
	Charisma int
	Constitution int
	Wisdom int
	Intelligence int
	Dexterity int
	//Saves Saves
	Damage_Vulnerabilities string
	Damage_Immunities string
// TODO: write a custom deserializer so that I can implement BaseStats
}


type DamageRoll struct {
	rolls int
	dice int
}

func (dr *DamageRoll) UnmarshalJSON(data []byte) error {
	var roll string

	err := json.Unmarshal(data, &roll);
	if err != nil {
		fmt.Println("issue")
		return err
	}

	parts := strings.Split(roll, "d")
	dr.dice = getInt(parts[0])
	dr.rolls = getInt(parts[1])
	return nil
}

// parse monster data from json

func loadAllMonstersFromFile() []Monster {
	absPath, _ := filepath.Abs("monsters.json")
	return loadStringFromFile(absPath)
}


func getInt(s string) int {
	integer, err := strconv.Atoi(s)
	if err == nil {
		fmt.Println("Talk to Amy... %s", err)
	}

	return integer
}


func loadStringFromFile(pathToFile string) []Monster {
	b, err := ioutil.ReadFile(pathToFile) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	dataString := string(b) // convert content to a 'string'

	return parseMonsters(dataString)
	// should return stuff instead
}

func parseMonsters(jsonString string) []Monster {
	//sample := []byte(`[{ "name": "Aboleth", "size": "Large", "type": "aberration", "subtype": "", "alignment": "lawful evil", "armor_class": 17, "hit_points": 135, "hit_dice": "18d10", "speed": "10 ft., swim 40 ft.", "strength": 21, "dexterity": 9, "constitution": 15, "intelligence": 18, "wisdom": 15, "charisma": 18, "constitution_save": 6, "intelligence_save": 8, "wisdom_save": 6, "history": 12, "perception": 10, "damage_vulnerabilities": "sdfdfda", "damage_resistances": "sadfadf", "damage_immunities": "sdafsdf", "condition_immunities": "", "senses": "darkvision 120 ft., passive Perception 20", "languages": "Deep Speech, telepathy 120 ft.", "challenge_rating": "10" }]`)
	monsters := make([]Monster, 30) //  what is the amount

	err := json.Unmarshal([]byte(jsonString), &monsters)
	//err := json.Unmarshal(sample,&monsters)

	if err != nil {
		fmt.Printf("bad result, please tell Amy %s", err)
	}
	return monsters
}
