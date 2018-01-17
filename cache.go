package main

import (
	"fmt"
	"strconv"
)

// want to be able to cache enemy information relating to the environment
// cache things in CR order?

var monsterCache map[string]Monster

func loadCache() {
	monsterCache = make(map[string]Monster)
	// retrieve from database or api maybe
	monsters := loadAllMonstersFromFile()

	for i := 0; i < len(monsters); i++ {
		monster := monsters[i]
		monsterCache[monster.Name] = monster
	}

	fmt.Printf("%v\n", monsterCache)

}

// TODO: refactor to imporve extensibility
func getMonsters(env Environment, cr int) []Monster {
	candidates := make([]Monster,10)
	for name, mon := range monsterCache {
		crInt, _ := strconv.Atoi(mon.Challenge_Rating)
		if crInt == cr {
			candidates = append(candidates, mon)
			fmt.Println(name)
		}
	}
}

func loadMonstersByEnvironment(env Environment) {

}

func getMonstersForEnvironment(env Environment) {
	// make sure the environment ID is related to the database's
	// TODO: programmatically enforce ID relation
	// want to index it different ways anyway, like order by CR
}
