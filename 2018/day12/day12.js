var fs = require('fs')
var expect = require('chai').expect

describe('e2e', () => {

    it('firstRound', () => {
        let world = load("input12-test.txt")
        console.log(world.state)
        run(world)
        expect("..#...#....#.....#..#..#..#..").to.equal(world.state)
    })

})

function load(file) {
    let world = {
        state: "",
        dictionary: [],
        rules: []
    }
    let input = fs.readFileSync(file, 'utf8')
    let lines = input.split("\n")
    let firstline = lines.shift()
    let split = firstline.split("initial state: ")
    world.state = ".." + split[1]
    lines.shift()
    lines.forEach(line => {
        let split = line.split(" => ")
        let newRule = {rule:split[0],action:split[1]}
        if (newRule.rule.charAt(2) != newRule.action) {
            console.log("adding rule " + newRule.rule + " => " + newRule.action)
            world.dictionary.push(newRule)
            world.rules.push(newRule.rule)
        }
    })
    return world
}

function run(world) {
    let newState = 
    for (let index = 2; index < world.state.length-2; index++) {
        let chunk = world.state.substring(index-2, index+3)
        if (world.rules.contains(chunk)) {

        }
        console.log(index-2, index+3, world.state.substring(index-2, index+3))
    }
}