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

type Difficulty int // encounter difficulty
const (
	easy Difficulty = iota
	medium
	hard
deadly
)

type BaseStats struct {
	strength int
	charisma int
	constitution int
	wisdom int
	intelligence int
	dexterity int
}

type Saves struct { // statistics related to saving throws
	strength_save int
	charisma_save int
	constitution_save int
	wisdom_save int
	intelligence_save int
	dexterity_save int
}

type Character struct {
	name string
	level int
	stats BaseStats
}

func createEncounter(numEnemiesDesired int, env Environment) {
	numEnemies = numEnemiesDesired

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
