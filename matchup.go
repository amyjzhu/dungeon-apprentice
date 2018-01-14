package main

var numChars int
var numEnemies int

// TODO: move environment enum to better location

type Environment int

const (
	arctic Environment = iota
	coastal
	desert
	forest
	grassland
	hill
	mountain
	swamp
	underdark
	underwater
	urban
)

type Character struct {
	name string
	level int
	strength int
	charisma int
	constitution int
	wisdom int
	intelligence int
	dexterity int
}

type Monster struct {
	cr int // challenge rating

}

// TODO: later be able to specify round up or down
func aggregateCharacterLevels(party []Character) int {
	len := len(party)
	sum := 0
	for i := 0; i < len; i++ {
		sum += party[i].level
	}

	return sum / len
}

func createEncounter(numEnemies int, ) {

}
