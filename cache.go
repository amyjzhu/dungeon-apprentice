package main

import (
	"fmt"
	"strconv"
)

// want to be able to cache enemy information relating to the environment
// cache things in CR order?

var monsterCache map[string]Monster
var environmentCache map[Environment][]Monster

func initialize() {
	loadCache()
	loadEnvironmentCache()
}

func loadCache() {
	monsterCache = make(map[string]Monster)
	monsters := loadAllMonstersFromFile() // TODO load from database instead

	for i := 0; i < len(monsters); i++ {
		monster := monsters[i]
		monsterCache[monster.Name] = monster
	}

	fmt.Printf("%v\n", monsterCache)

}

func loadEnvironmentCache() {
	environmentCache = make(map[Environment][]Monster)
	for _, env := range Environments {
		monsters := databaseSelectMonstersByEnvironment(env);
		environmentCache[env] = monsters
	}
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

	return candidates
}

func getMonstersForEnvironment(env Environment) []Monster {
	// make sure the environment ID is related to the database's
	// TODO: programmatically enforce ID relation
	// want to index it different ways anyway, like order by CR

	return environmentCache[env]
}
