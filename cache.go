package main

import "fmt"

// want to be able to cache enemy information relating to the environment
// cache things in CR order?

var monsterCache map[string]Monster

func loadCache() {
	// retrieve from database or api maybe
	monsters := loadAllMonstersFromFile()

	for i := 0; i < len(monsters); i++ {
		monster := monsters[i]
		monsterCache[monster.Name] = monster
	}

	fmt.Printf("%v\n", monsterCache)


}

func loadMonstersByEnvironment(env Environment) {

}

func getMonstersForEnvironment(env Environment) {
	// make sure the environment ID is related to the database's
	// TODO: programmatically enforce ID relation
	// want to index it different ways anyway, like order by CR
}
