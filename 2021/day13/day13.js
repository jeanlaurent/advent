var fs = require('fs')

// https://adventofcode.com/2021/day/13

function load(file) {
    var input = fs.readFileSync(file, 'utf8')

    let width = -1
    let height = -1
    let coords = []
    let folds = []
    let readFold = false

    input.split('\n').forEach(line => {
        if (line.length == 0) {
            readFold = true
        } else if (readFold) {
            let tokenFold = line.split("=") 
            let axis = tokenFold[0].charAt(tokenFold[0].length-1)
            folds.push({axis:axis, length: +tokenFold[1]})
        } else {
            let tokens = line.split(",")
            let newCoord = {x:+tokens[0], y:+tokens[1]}
            if (newCoord.x>width) {
                width = newCoord.x
            }
            if (newCoord.y>height) {
                height = newCoord.y
            }
            coords.push(newCoord)    
        }
    })
    width++
    height++
    return {points:coords, width:width, height:height, folds: folds}
}


function runDay13(file) {
    let world = load(file)
    let count = 0
    
    world.folds.forEach(fold => {
    let newWidth = world.width
    let newHeight = world.height
    let newCoords = []
    if (fold.axis == 'y') {
        newHeight = (world.height / 2)  >> 0
        world.points.forEach(point => {
            if (point.y < fold.length) {
                newCoords.push(point)
            } else {
                let translated = {x:point.x, y: fold.length - (point.y - fold.length) }
                if (isValid(translated,newWidth,newHeight) && (!exist(translated, newCoords))) {
                    newCoords.push(translated)
                }
            }
        })
      } else {
        newWidth = (world.width / 2)  >> 0 
          world.points.forEach(point => {
            if (point.x < fold.length) {
                newCoords.push(point)
            } else {
                let translated = {x:fold.length - (point.x - fold.length), y: point.y }
                if (isValid(translated,newWidth,newHeight) && (!exist(translated, newCoords))) {
                    newCoords.push(translated)
                }
            }      
          })
        }
      world.points = newCoords
      world.height = newHeight
      world.width = newWidth
      if (count == 0) {
          count++
          removeDouble(world)
          console.log("Step1", file, "->", world.points.length)
      }
    })
    removeDouble(world)
    console.log("Step2")
    printWorld(world)
}

function exist(coord, coords) {
    for (let i = 0; i < coords.length; i++) {
        if ((coord.x == coords[i].x) && (coord.y == coords[i].y)) {
            return true
        }
    }
    return false
}

function isValid(point, width, height) {
    return ((point.x>=0) && (point.y>=0) && (point.x < width) && (point.y < height))
}

function removeDouble(world) {
    let curatedPoints = []
    world.points.forEach(point => {
        if (!exist(point, curatedPoints)) {
            curatedPoints.push(point)
        }    
    })
    world.points = curatedPoints  
}

function printWorld(world) {
    for (let y = 0; y < world.height; y++) {
        let line = ""
        for (let x = 0; x < world.width; x++) {
            if (exist({x:x,y:y}, world.points)) {
                line += "#"    
            } else {
                line += " "
            }
        }
        console.log(line)
    }
}

runDay13('input-test.txt')
runDay13('input.txt') 


