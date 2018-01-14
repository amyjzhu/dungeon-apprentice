package main

// want to be able to cache enemy information relating to the environment

type Monster struct {
	challenge_rating int // challenge rating
	env []Environment // places usually found
	exp int // experience points award for defeat
	stats BaseStats
	saves Saves
	vulnerabilities string
	immunities string


}

func loadCache() {
	// retrieve from database or api maybe
}

func loadMonstersByEnvironment(env Environment) {

}

func getMonstersForEnvironment(env Environment) {
	// make sure the environment ID is related to the database's
	// TODO: programmatically enforce ID relation
}
