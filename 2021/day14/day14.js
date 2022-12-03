var fs = require('fs')

// https://adventofcode.com/2021/day/14

function load(file) {
    var input = fs.readFileSync(file, 'utf8')

    let lines = input.split("\n")

    let polymerTemplate = lines.shift()
    lines.shift()
    let iRules = []

    lines.forEach(line => {
        let rules = line.split(' -> ')
        let pair = rules[0]
        let inserted = rules[1]
        iRules[pair] = inserted
    })
    return {template: polymerTemplate, rules:iRules,}
}


function runDay14(file, maxstep) {
    let world = load(file)
    console.log(world)

    let template = world.template

    for (let step = 0; step < maxstep; step++) {
        console.log(step)
        let newTemplate = ""
        for (let i = 0; i < template.length-1; i++) {
            let pair = template[i] + template[i+1]
            let inserted = world.rules[pair]
            newTemplate += template[i] + inserted
        }
        template = newTemplate + template[template.length - 1]
        // console.log(template)
    }   
    

    let compute = new Map()
    for (let i = 0; i < template.length; i++) {
        if (compute.has(template[i])) {
            compute.set(template[i], compute.get(template[i]) + 1)
        } else {
            compute.set(template[i], 1)
        }
    }
    console.log(compute)
    let max = -1
    let min = 99999999
    for (let [key,value] of compute) {
        if (value > max ) {
            max = value
        }
        if (value < min) {
            min = value
        }
    }

    console.log(max, min)
    console.log("step1", file, "->", max - min)
}

function runDay14s2(file, maxstep) {
    let world = load(file)
    console.log(world)

    let rules = world.rules

    let pairs = new Map()
    for (let i = 0; i < world.template.length -1; i++) {
        pairs.set(world.template[i] + world.template[i+1], 1)
    }
    
    for (let step = 0; step < maxstep; step++) {
        //console.log(step, pairs)
        pairsCopy = new Map(pairs)
        pairsCopy.forEach((count,pair) => {
            let char = rules[pair]
            decrementPair(pair, pairs, count)
            incrementPair(pair[0] + char, pairs, count)
            incrementPair(char + pair[1], pairs, count)
        })
    }
    console.log(pairs)
    let compute = new Map()
    pairs.forEach((count,pair) => {
        if (compute.has(pair[1])) {
            compute.set(pair[1], BigInt(compute.get(pair[1])) + BigInt(count) )
        } else {
            compute.set(pair[1], BigInt(count) )
        }
    })
    compute.set(world.template[0], BigInt(compute.get(world.template[0]) + BigInt(1)) )

    console.log(compute)
    let max = -1
    let min = Number.MAX_SAFE_INTEGER
    for (let [key,value] of compute) {
        if (value > max ) {
            max = value
        }
        if (value < min) {
            min = value
        }
    }

    console.log(max, min)
    console.log("step2", file, "->", max - min)
}

function decrementPair(pair, map, value) {
    if (map.get(pair) - value <= 0) {
        map.delete(pair)
    } else {
        map.set(pair, map.get(pair) - value )
    }
}

function incrementPair(pair, map,value) {
    if (map.has(pair)) {
        map.set(pair, map.get(pair) + value )
    } else {
        map.set(pair, value )
    }
}

runDay14('input-test.txt', 10)
runDay14('input.txt', 10) 

runDay14s2('input-test.txt', 40)
runDay14s2('input.txt', 40) // 3199799196330 too low :(


