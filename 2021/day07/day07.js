var fs = require('fs')

function load(file) {
    var input = fs.readFileSync(file, 'utf8')
    
    let crabs = []
    input.split(",").forEach(item => {
        crabs.push(+item)
    })
    return crabs
}

function runDay07(file) {
    let crabs = load(file)

    let maxFuel = 999999999
    let bestPosition = -1

    for (let alignedPositionIndex = 0; alignedPositionIndex < crabs.length; alignedPositionIndex++) {
        const alignPosition = crabs[alignedPositionIndex]

        let fuelSpent = 0
        crabs.forEach(crab => {
            fuelSpent += Math.abs(alignPosition - crab)
        })
        if (fuelSpent < maxFuel) {
            maxFuel = fuelSpent
            bestPosition = alignPosition
        }
    }
    console.log("Step1", file, "->", maxFuel)
}

function runDay07s2(file) {
    let crabs = load(file)

    let maxFuel = 999999999
    let bestPosition = -1

    const max = Math.max(...crabs)
    const min = Math.min(...crabs)
    // console.log(min, max)

    for (let alignPosition = min; alignPosition <= max; alignPosition++) {

        let fuelSpent = 0
        for (let crabIndex = 0; crabIndex < crabs.length; crabIndex++) {
            let distance = Math.abs(alignPosition - crabs[crabIndex])
            let fuelCost = (distance * (distance + 1))/2 
            fuelSpent += fuelCost
        }
        // console.log("Position", alignPosition, "cost", fuelSpent)
        if (fuelSpent < maxFuel) {
            maxFuel = fuelSpent
            bestPosition = alignPosition
        }
    }
    console.log("Step2", file, "->", bestPosition, maxFuel )
}

runDay07('input-test.txt')
runDay07('input.txt')
runDay07s2('input-test.txt')
runDay07s2('input.txt')
