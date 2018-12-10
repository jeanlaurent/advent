var fs = require('fs')

function load() {
    //var input = fs.readFileSync('input10-test.txt', 'utf8')
    var input = fs.readFileSync('input10.txt', 'utf8')
    let points = []
    input.split("\n").forEach(line => {
        let items = line.split(/[<|,|>]/)
        points.push({
            position : {x:+items[1],y:+items[2]},
            velocity : {x:+items[4],y:+items[5]}
        })
    })
    return points
}

function draw(points) {
    let world = initWorld(points)
    drawGrid(world, points)
}

function initWorld(points) {
    var minX = 9999999
    var minY = 9999999
    var maxX = 0
    var maxY = 0 
    points.forEach(point => {
        if (minX>point.position.x) {
            minX=point.position.x
        }
        if (minY>point.position.y) {
            minY=point.position.y
        }
        if (maxX<point.position.x) {
            maxX=point.position.x
        }
        if (maxY<point.position.y) {
            maxY=point.position.y
        }
    })
    return {
        minX: minX,
        minY: minY,
        width: maxX-minX +1,
        height: maxY-minY +1,
    }
}

function drawGrid(world, points) {
    world.grid = new Array(world.height).fill(".").map(() => new Array(world.width).fill("."));
    points.forEach(point =>{
        x = point.position.x - world.minX
        y = point.position.y - world.minY
        world.grid[y][x]="#"
    })
    for (let y = 0; y < world.height; y++) {
        line = ""
        for (let x = 0; x < world.width; x++) {
            line += world.grid[y][x]
        }
        console.log(line)
    }

}


let points = load()
let stop = false
let turn = 0
while(!stop) {
    turn++
    points.forEach(point => {
        point.position.x += point.velocity.x
        point.position.y += point.velocity.y
    })
    world = initWorld(points)
    if (world.height == 10) {
        stop = true
    }
}
console.log("Step1")
draw(points)
console.log("Step2 --> " + turn)


