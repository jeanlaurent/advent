var fs = require('fs')

// input contains 50 points
const ids ="abcdefghijklmnopqrstuvwxyABCDEFGHIJKLMNOPQRSTUVWXY"

function load(file) {
    var input = fs.readFileSync(file, 'utf8')
    let points = []
    var i = 0
    input.split("\n").forEach(line => {
        let items = line.split(",")
        points.push({
            name : ids.charAt(i),
            x : +items[0],
            y: +items[1]
        })
        i++
    })
    return points
}

function initWorld(points) {
    var minX = 9999999
    var minY = 9999999
    var maxX = 0
    var maxY = 0 
    points.forEach(point => {
        if (minX>point.x) {
            minX=point.x
        }
        if (minY>point.y) {
            minY=point.y
        }
        if (maxX<point.x) {
            maxX=point.x
        }
        if (maxY<point.y) {
            maxY=point.y
        }
    })
    width = maxX + 3
    height = maxY + 3
    var grid = []
    for (let y = 0; y < height; y++) {
        grid[y] = []
        for (let x = 0; x < width; x++) {
            grid[y][x] = {owner:'*', distance:999}
        }
    }
    return {
        minX: minX,
        minY: minY,
        width: width,
        height: height,
        grid: grid
    }
}

function draw(world) {
    for (let y = 0; y < world.height; y++) {
        for (let x = 0; x < world.width; x++) {
            process.stdout.write("" + world.grid[y][x].owner)
        }
        process.stdout.write("\n")
    }
}

function distance(x,y,xx,yy) {
    return Math.abs(xx - x) + Math.abs(yy -  y)
}

function populateWorld(world, points) {
    points.forEach(point => {
        for (let y=0; y<world.height; y++) {
            for (let x=0; x<world.width; x++) {
                var dist = distance(point.x,point.y,x,y)
                if (dist < world.grid[y][x].distance) {
                    world.grid[y][x].distance=dist
                    world.grid[y][x].owner=point.name    
                } else if (dist === world.grid[y][x].distance) {
                     world.grid[y][x].owner='.'
                }
            }
        }
    })
}

function computeSurface(world, points) {
    let infinite = []
    let count = {}
    var largest = 0
    points.forEach(point => {
        count[point.name] = 0
    })
    
    for (let y=0; y<world.height; y++) {
        for (let x=0; x<world.width; x++) {
            var owner = world.grid[y][x].owner
            if (x==0 || y==0 || x== world.width-1 || y == world.height-1) {
                if (!infinite.includes(owner)) {
                    infinite.push(owner)
                }
            }
            if (!infinite.includes(owner)) {
                count[owner]++
                if (largest < count[owner]) {
                    largest = count[owner]
                }
            }
        }
    }
    return largest
}

function computeSafeLocation(world, points) {
    for (let y=0; y<world.height; y++) {
        for (let x=0; x<world.width; x++) {
            points.forEach(point => {
                if (world.grid[y][x].safeCount == undefined) {
                    world.grid[y][x].safeCount = 0
                }
                world.grid[y][x].safeCount += distance(x,y,point.x, point.y)
            })
        }
    }
    let safeRegionSurface = 0
    for (let y=0; y<world.height; y++) {
        for (let x=0; x<world.width; x++) {
            if (world.grid[y][x].safeCount < 10000) {
                safeRegionSurface++
            }
        }
    }
    return safeRegionSurface
}

function runDay06(file, shouldDraw) {
    let points = load(file)
    let world = initWorld(points)
    populateWorld(world, points)
    if (shouldDraw) {
        draw(world)
    }
    console.log(file)
    console.log("\tstep1",computeSurface(world, points))
    console.log("\tstep2",computeSafeLocation(world, points))
}

runDay06('input06-test.txt', false)
runDay06('input06.txt', false)

