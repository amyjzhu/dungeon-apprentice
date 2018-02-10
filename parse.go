package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"path/filepath"
	"strconv"
)


func (dr *DamageRoll) UnmarshalJSON(data []byte) error {
	var roll string

	err := json.Unmarshal(data, &roll);
	if err != nil {
		fmt.Println("issue")
		return err
	}

	parts := strings.Split(roll, "d")
	dr.die = getInt(parts[1])
	dr.rolls = getInt(parts[0])
	return nil
}

// parse monster data from json

func loadAllMonstersFromFile() []Monster {
	absPath, _ := filepath.Abs("monsters.json")
	return loadStringFromFile(absPath)
}


func getInt(s string) int {
	integer, err := strconv.Atoi(s)
	if err != nil {
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
	monsters := make([]Monster, 30) //  what is the number of mosnters? have to go through and find out

	err := json.Unmarshal([]byte(jsonString), &monsters)

	if err != nil {
		fmt.Printf("bad result, please tell Amy %s", err)
	}
	return monsters
}