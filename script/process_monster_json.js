const fs = require("fs");

// usage:
// I can't be assed to make this user-friendly since it's just a script for me
// put the desired original file as monsters.json or whatever
// run node main.js
// get new data from new_monsters.json

var baseStats = ["strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma"];

var skills = ["athletics","acrobatics","sleight of hand","stealth","arcana",
    "history","investigation","nature","religion","animal handling","insight",
    "medicine","perception","survival","deception","intimidation","performance","persuasion"];

function loadData(path) {
    fs.readFile(path, 'utf-8', parseData);
}

function parseData(err, jsonString) {
    var newMonsters = [];

    var monsterData = JSON.parse(jsonString);
    monsterData.forEach(function (monster) {
        var revisedMonster = Object.assign({}, monster);
        revisedMonster.base_stats = {};
        revisedMonster.skills = {};

        Object.keys(monster).forEach(function (key) {
            if (baseStats.includes(key)) {
                revisedMonster.base_stats[key] = monster[key];
                delete revisedMonster[key];
            }

            if (skills.includes(key)) {
                revisedMonster.skills[key] = monster[key];
                delete revisedMonster[key];
            }

        });

        newMonsters.push(revisedMonster);

    });

    writeToFile(newMonsters);

}

function writeToFile(monsters) {
    var jsonString = JSON.stringify(monsters);

    fs.writeFile("new_monsters.json", jsonString);
}

loadData("monsters.json");