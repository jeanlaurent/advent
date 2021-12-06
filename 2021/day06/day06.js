var fs = require('fs')

function load(file) {
    var input = fs.readFileSync(file, 'utf8')
    
    let fishes = []
    input.split(",").forEach(item => {
        fishes.push(+item)
    })
    return fishes
}

function runDay06(file, maxDays) {
    let fishes = load(file)

    for (let dateInDay = 1; dateInDay <= maxDays; dateInDay++) {
        let newFish = 0
        for (let i = 0; i < fishes.length; i++) {
            if (fishes[i] == 0) {
                newFish ++
                fishes[i] = 6
            } else {
                fishes[i]--
            }
        }
        for (let i = 0; i < newFish; i++) {
            fishes.push(8)         
        }
    }
    console.log("Step1", file, maxDays, "->", fishes.length)
}

function runDay06s2(file, maxDays) {

    let fishes = Array(9).fill(0)
    load(file).forEach(loadedFish => {
        fishes[loadedFish]++
    })

    for (let dateInDay = 1; dateInDay <= maxDays; dateInDay++) {
        let newFishes = Array(9).fill(0)
        for (let i = 0; i < fishes.length; i++) {
            if (i == 0) {
                newFishes[6] += fishes[0]
                newFishes[8] += fishes[0]
            } else {
                newFishes[i-1] += fishes[i]
            }
        }
        fishes = newFishes
    }

    let numberOfFish = 0
    for (let i = 0; i < fishes.length; i++) {
        numberOfFish += fishes[i]
    }
    console.log("step2", file, maxDays, "->", numberOfFish)
}

runDay06('input-test.txt', 80)
runDay06('input.txt', 80)

runDay06s2('input-test.txt', 80)
runDay06s2('input.txt', 80)

runDay06s2('input-test.txt', 256)
runDay06s2('input.txt', 256)