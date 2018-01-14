package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// parse monster data from json

func loadStringFromFile(pathToFile string) {
	b, err := ioutil.ReadFile(pathToFile) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	dataString := string(b) // convert content to a 'string'

	parseMonsters(dataString)
	// should return stuff instead
}

func parseMonsters(jsonString string) {
	var monsters []Monster
	err := json.Unmarshal([]byte(jsonString), &monsters)

	if err != nil {
		fmt.Printf("bad result, please tell Amy %s", err)
	}
	fmt.Printf("%v\n", monsters)
}
