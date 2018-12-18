var fs = require('fs')
var expect = require('chai').expect

describe('loading/drawing', () => {

    it('check width, height', () => {
        let world = load("input18-test1.txt")
        expect(10).to.equal(world.width)
        expect(10).to.equal(world.height)
    })

    it('draw', () => {
        let world = load("input18-test1.txt")
        let rawDrawning = fs.readFileSync("input18-test1.txt", 'utf8')
        expect(rawDrawning).to.equal(draw(world))
    })

    it('example minute1', () => {
        let world = load("input18-test1.txt")
        turn(world)
        let rawDrawning = fs.readFileSync("input18-test2.txt", 'utf8')
        expect(rawDrawning).to.equal(draw(world))
    })

    it('adjacent', () => {
        let world = load("input18-test1.txt")
        let adjacents = findAdjacent(3,3,world)
        expect([{x:2,y:2},{x:3,y:2},{x:4,y:2},{x:2,y:3},{x:4,y:3},{x:2,y:4},{x:3,y:4},{x:4,y:4}]).to.deep.equal(adjacents)
    })

    it('example 10 minutes', () => {
        let world = load("input18-test1.txt")
        for (let index = 0; index < 10; index++) {
            turn(world)
        }
        expect(1147).to.equal(computeScore(world))
    })

    it('step1 10 minutes', () => {
        let world = load("input18.txt")
        expect(360720).to.equal(run(world, 10))
    })

    it('step2 1000000000 minutes or so...', () => {
        let world = load("input18.txt")
        let mapList = []
        let finalScoreIndex = -1
        for (let index = 0; index < 1000000000; index++) {
            turn(world)
            for (let lindex = 0; lindex < mapList.length; lindex++) {
                if (areWorldEqual(world, mapList[lindex])) {
                    let fromIndex = lindex
                    let period = index - lindex
                    finalScoreIndex = fromIndex + (1000000000 - fromIndex) % period
                }
            }
            mapList[index] = JSON.parse(JSON.stringify(world)) // copy array in JS :D
            if (finalScoreIndex != -1) {
                break
            }
        }
        world = load("input18.txt")
        for (let index = 0; index < finalScoreIndex; index++) {
            turn(world)
        }
        expect(197276).to.equal(computeScore(world))
    })
})

const open = 0
const tree = 1
const lumberjack = 2

function run(world,loop) {
    for (let index = 0; index < loop; index++) {
        turn(world)
    }
    return computeScore(world)
}

function areWorldEqual(world1, world2) {
    for (let y = 0; y < world1.height; y++) {
        for (let x = 0; x < world1.width; x++) {
            if (world1.map[y][x] !== world2.map[y][x]) {
                return false
            }
        }
    }
    return true
}

function computeScore(world) {
    let count = [0,0,0]
    for (let y = 0; y < world.height; y++) {
        for (let x = 0; x < world.width; x++) {
            count[world.map[y][x]]++
        }
    }
    return count[tree] * count[lumberjack]
}

function findAdjacent(x,y, world) {
    let adjacents = []
    for (let yy = -1; yy <= 1; yy++) {
        for (let xx = -1; xx <= 1; xx++) {
            if ((xx == 0) && (yy == 0)) {
                continue
            }
            let xxx = xx + x
            let yyy = yy + y
            if (isValid(xxx,yyy,world)) {
                adjacents.push({x:xxx,y:yyy})
            }
        }
    }
    return adjacents
}

function turn(world) {
    let newMap = []
    for (let y = 0; y < world.height; y++) {
        newMap[y] = []
        for (let x = 0; x < world.width; x++) {
            let adjacents = findAdjacent(x,y, world)
            let count = [0,0,0]
            adjacents.forEach(adj => {
                count[world.map[adj.y][adj.x]]++
            })
            switch(world.map[y][x]) {
                case open:
                    newMap[y][x] = open
                    if (count[tree] >= 3) {
                        newMap[y][x] = tree
                    }
                    break
                case tree:
                    newMap[y][x] = tree
                    if (count[lumberjack] >= 3) {
                        newMap[y][x] = lumberjack
                    }
                    break
                case lumberjack:
                    newMap[y][x] = open
                    if (count[lumberjack] >= 1 && count[tree] >= 1) {
                        newMap[y][x] = lumberjack
                    }
                    break
            }
        }
    }
    world.map = newMap
}

function load(file) {
    let map = []
    let input = fs.readFileSync(file, 'utf8')
    let lines = input.split("\n")
    let y = 0
    lines.forEach(line => {
        map[y] = []
        for (let x = 0; x < line.length; x++) {
            switch(line.charAt(x)) {
                case '.':
                    map[y][x] = open
                    break
                case '|':
                    map[y][x] = tree
                    break
                case '#':
                    map[y][x] = lumberjack
                    break
            }
        }
        y++
    })
    height = y
    width = map[0].length
    return {
        height : height,
        width: width,
        map: map,
    }
}

function isValid(x,y, world) {
    return x >=0 && x < world.width && y >= 0 && y < world.height
}

function draw(world){
    let lines = ""
    for (let y = 0; y < world.height; y++) {
        for (let x = 0; x < world.width; x++) {
            if (world.map[y][x] === open) {
                lines += '.'
                continue
            }
            if (world.map[y][x] === tree) {
                lines += '|'
                continue
            }
            if (world.map[y][x] === lumberjack) {
                lines += '#'
            }
        } 
        if (y != world.height -1) {
            lines += "\n"  
        }
    }
    return lines
}