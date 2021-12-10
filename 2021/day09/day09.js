var fs = require('fs')

// https://adventofcode.com/2021/day/9

const modifiers = [[-1,0],[1,0],[0,-1],[0,1]]

function load(file) {
    var input = fs.readFileSync(file, 'utf8')

    let world = []
    let width = -1

    input.split('\n').forEach(line => {
        let worldLine = []
        Array.from(line).forEach( number => {
            worldLine.push(+number)
        })
        width = worldLine.length + 1
        world.push(worldLine)
    })
    return {map:world, width: width, height: world.length+1}
}

function isValid(world, xx, yy) {
    return ((xx >= 0) && (xx < world.width) && (yy >= 0) && (yy <world.height))
}

function findLowPoints(world) {
    let lowPoints = []
    for (let y = 0; y < world.height; y++) {
        for (let x = 0; x < world.width; x++) {
            let low = true
            for (let i = 0; i < modifiers.length; i++) {
                if (!low) {
                    break
                }
                const mod = modifiers[i]
                if (isValid(world, x+mod[0], y+mod[1])) {
                    low = (world.map[y][x] < world.map[y+mod[1]][x+mod[0]])
                }
            }
            if (low) {
                lowPoints.push({x:x,y:y})
            }
         }
    }
    return lowPoints
}

function runDay09(file) {
    let world = load(file)
    let lowPoints = findLowPoints(world)

    let sum = 0
    lowPoints.forEach(lowPoint => {
        sum += world.map[lowPoint.y][lowPoint.x] + 1
    })
    console.log("Step1", file, "->", sum)
}

function runDay09s2(file) {
    let world = load(file)
    let lowPoints = findLowPoints(world)

    let basinSizes = []
    lowPoints.forEach(lowPoint => {
        // let lowPoint = lowPoints[3]

        //console.log("Looking at basin starting at", lowPoint.x, lowPoint.y)
        let basinCoords = []
        findBasinSize(world, lowPoint.x, lowPoint.y, basinCoords)
        // console.log(basinCoords)
        basinSizes.push(+basinCoords.length)
        
        
    })
    console.log(basinSizes.sort(function(a,b){return b-a}))
    let top3 = basinSizes.sort(function(a,b){return b-a}).splice(0,3)
    console.log(top3)
    let result = 1
    top3.forEach(basinSize => {
        result *= basinSize
    })
    console.log("Step2", file, "->", result)
}

function findBasinSize(world, x, y, basinCoords) {
    // console.log(x,y) 
    let value = world.map[y][x]
    if (value == 9) {
        return 0
    }
    basinCoords.push({x:x,y:y})
    //drawVisited(world,basinCoords) //471653
    for (let i = 0; i < modifiers.length; i++) {
        const mod = modifiers[i]
        const xx = x+mod[0]
        const yy = y+mod[1]
        if (isValid(world, xx,yy )) {            
            if (world.map[yy][xx] == value+1) {
                if (!hasBeenVisited(xx,yy, basinCoords)) {
                    //console.log("basin continue on ",xx,yy, "[", world.map[yy][xx], "] from", x, y, "[", world.map[y][x], "]")
                    findBasinSize(world, xx,yy, basinCoords)
                }
            }
        }
    }
}

function drawVisited(world, visited) {
    for (let y = 0; y < world.height; y++) {
        let line = ""
        for (let x = 0; x < world.width; x++) {
            if (hasBeenVisited(x,y,visited)) {
                line += world.map[y][x]
            } else {
                line += '*'
            }
        }
        console.log(line)
    }
}

function hasBeenVisited(x,y, visitedLocation) {
    for (let i = 0; i < visitedLocation.length; i++) {
        if ((visitedLocation[i].x == x) && (visitedLocation[i].y == y)) {
            return true
        }
    }
    return false
}

runDay09('input-test.txt')
runDay09('input.txt')

runDay09s2('input-test.txt')
runDay09s2('input.txt')
