var fs = require('fs')

const startx = 0
const starty = 1
const endx = 2
const endy = 3

const xx = 0
const yy = 1

function load(file) {
    var input = fs.readFileSync(file, 'utf8')
    let segments = []
    input.split("\n").forEach(line => {
        let pairs = line.split("->")
        let startcoords = pairs[0].trim().split(",")
        let endcoords = pairs[1].trim().split(",")
        let segment = [+startcoords[xx],+startcoords[yy], +endcoords[xx], +endcoords[yy]]
        // sort data so we from left to right / top to bottom (except diagonal...)
        if (segment[startx] == segment[endx]) {
            if (segment[starty] > segment[endy] ){
                let swapTmp = segment[starty]
                segment[starty] = segment[endy]
                segment[endy] = swapTmp
            }
        } else if (segment[starty] == segment[endy]) {
            if (segment[startx] > segment[endx] ) {
                let swapTmp = segment[startx]
                segment[startx] = segment[endx]
                segment[endx] = swapTmp
            }
        } else {
            if (segment[startx] > segment[endx] ) {
                let swapTmp = segment[startx]
                segment[startx] = segment[endx]
                segment[endx] = swapTmp
                swapTmp = segment[starty]
                segment[starty] = segment[endy]
                segment[endy] = swapTmp
            }   
        }
        segments.push(segment)
    })
    return segments
}

function runDay05(file, handleDiagonal) {
    console.log("Step1", file)
    let segments = load(file)

    let max = Math.max(...segments.flat()) + 1
    let world = Array(max).fill(0).map(y => Array(max).fill(0))

    segments.forEach(segment => {
        if (segment[startx] == segment[endx]) {
            let x = segment[startx]
            for (let y = segment[starty]; y <= segment[endy]; y++) {
                world[y][x] += 1
            }
        } else if (segment[starty] == segment[endy]) {
            let y = segment[starty]
            for (let x = segment[startx]; x <= segment[endx]; x++) {
                world[y][x] += 1 
            }
        } else if (handleDiagonal) {
            let x = segment[startx]
            let y = segment[starty]

            for (let i = 0; i <= segment[endx]-segment[startx]; i++) {
                world[y][x] += 1 
                x++
                if (segment[starty] < segment[endy]) {
                    y++
                } else {
                    y--
                }   
            }
        }
    })    
    console.log(computeDangerousAreas(world), "\n")
}

function drawWorld(world, max) {
    for (let y = 0; y < max; y++) {
        let line = ""
        for (let x = 0; x < max; x++) {
            if (world[y][x] == 0) {
                line+= "."
            } else {
                line += "" + world[y][x]
            }
        }
        console.log(line)
    }
}

function computeDangerousAreas(world) {
    return world.flat().reduce((acc, value) => {
        if (value > 1) {
            return acc + 1
        }
        return acc
    },0)
}

runDay05('input-test.txt', false)
runDay05('input.txt', false)
runDay05('input-test.txt', true)
runDay05('input.txt', true)