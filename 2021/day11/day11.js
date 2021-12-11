var fs = require('fs')
const { wrap } = require('module')

// https://adventofcode.com/2021/day/11

function load(file) {
    var input = fs.readFileSync(file, 'utf8')

    let lines = []
    input.split('\n').forEach(string => {
        let line = []
        Array.from(string).forEach(char => {
            line.push(+char)
        })
        lines.push(line)
    })
    return {map:lines, width:lines[0].length, height:lines.length}
}


function runDay10(file, maxStep) {
    let world = load(file)
    let step = 0
    let flashesCount = 0

    //printWorld(world)

    while(step < maxStep) {
        let stepFlashes = []
        
        for (let y = 0; y < world.height; y++) {
            for (let x = 0; x < world.width; x++) {
                world.map[y][x]++
                if (world.map[y][x] > 9) {
                    if (!hasFlashed(x,y, stepFlashes)) {
                        stepFlashes.push({x:x,y:y})
                        flash(x,y,world, stepFlashes)
                    }
                }
            }
        }

        for (let y = 0; y < world.height; y++) {
            for (let x = 0; x < world.width; x++) {
                if (world.map[y][x] > 9) {
                    world.map[y][x] = 0
                }
            }
        }
        step++
        if ((maxStep == 9999) && (stepFlashes.length == 100)) {
            break
        }

        flashesCount += stepFlashes.length
        stepFlashes = []
        
        // printWorld(world)
    }
    if (maxStep == 9999) {
        console.log("Step2", file, "->", step)
    } else {
        console.log("Step1", file, "->", flashesCount)
    }
    
}

function hasFlashed(x,y, stepFlashes) {
    for (let i = 0; i < stepFlashes.length; i++) {
        if ((stepFlashes[i].x == x) && (stepFlashes[i].y == y)) {
            return true
        }
    }
    return false
}

function isValid(x,y) {
    return ((x>=0) && (y>=0) && (x<10) && (y<10))
}

function flash(x,y,world, stepFlashes) {
    for (let ym = -1; ym <= 1; ym++) {
        for (let xm = -1; xm <= 1; xm++) {
            let xx = x + xm
            let yy = y + ym
            if (!isValid(xx,yy)) {
                continue
            }
            if ((x == xx ) && ( y == yy)) {
                continue
            }
            world.map[yy][xx]++
            if (world.map[yy][xx] > 9) {
                if (!hasFlashed(xx,yy, stepFlashes)) {
                    stepFlashes.push({x:xx,y:yy})
                    flash(xx,yy,world, stepFlashes)
                }
            }
        }        
    }
}

function printWorld(world) {
    for (let y = 0; y < world.height; y++) {
        let line = ""
        for (let x = 0; x < world.width; x++) {
            line += world.map[y][x]
        }
        console.log(line)
    }
}

runDay10('inputsmall.txt', 2)
runDay10('input-test.txt',100)
runDay10('input.txt', 100)

runDay10('input-test.txt', 9999)
runDay10('input.txt', 9999)
