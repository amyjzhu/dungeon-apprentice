package main

type Monster struct {
	Name string
	Alignment string
	Challenge_Rating string // challenge rating
	env []Environment // places usually found
	exp int // experience points award for defeat
	Hit_Dice DamageRoll // TODO: convert all these string values to hit dice thing
	Base_Stats BaseStats
	Skills Skills
	//Saves Saves
	Damage_Vulnerabilities string
	Damage_Immunities string
	Special_Abilities []Action
	Legendary_Actions []Action
	Actions[] Action
	// TODO: write a custom deserializer so that I can implement BaseStats
}

type Action struct {
	Name string
	Desc string
	Attack_Bonus int
	Damage_Dice DamageRoll
	Damage_Bonus int
}

type Skills struct {
	// there has to be a better way to do this?
	// at least with a java class, enum constructors...
	Athletics int
	Acrobatics int
	Sleight_of_Hand int
	Stealth int
	Arcana int
	History int
	Investigation int
	Nature int
	Religion int
	Animal_Handling int
	Insight int
	Medicine int
	Perception int
	Survival int
	Deception int
	Intimidation int
	Performance int
	Persuasion int
}



type DamageRoll struct {
	rolls int
	die int
}

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
	Strength int
	Charisma int
	Constitution int
	Wisdom int
	Intelligence int
	Dexterity int
}

type Saves struct { // statistics related to saving throws
	Strength_Save int
	Charisma_Save int
	Constitution_Save int
	Wisdom_Save int
	Intelligence_Save int
	Dexterity_Save int
}

type Character struct {
	name string
	level int
	stats BaseStats
}

