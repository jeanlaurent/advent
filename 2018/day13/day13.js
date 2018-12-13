var fs = require('fs')
var expect = require('chai').expect

describe('unitest', () => {

    it('intersection', () => {
        cart = {id: 1,direction: down, turnIndex:0, tick: -1}
        changeDirectionForIntersection(cart)
        expect(right).to.equal(cart.direction)
        changeDirectionForIntersection(cart)
        expect(right).to.equal(cart.direction)
        changeDirectionForIntersection(cart)
        expect(down).to.equal(cart.direction)
        changeDirectionForIntersection(cart)
        expect(right).to.equal(cart.direction)
        changeDirectionForIntersection(cart)
        expect(right).to.equal(cart.direction)
        changeDirectionForIntersection(cart)
        expect(down).to.equal(cart.direction)
    })

})

describe('End to End', () => {

    it('run straight', () => {
        let world = load("input13-test1.txt")
        let colision = moveUntilEvent(world, true)
        expect({ x: 0, y: 3 }).to.deep.equal(colision)
    })

    it('use corner', () => {
        let world = load("input13-test2.txt")
        let colision = moveUntilEvent(world, true)
        expect({ x: 4, y: 2 }).to.deep.equal(colision)
    })

    it('full simple example', () => {
        let world = load("input13-test3.txt")
        let colision = moveUntilEvent(world, true)
        expect({ x: 7, y: 3 }).to.deep.equal(colision)
    })

    it('step1', (done) => {
        let world = load("input13.txt")
        let colision = moveUntilEvent(world, true)
        expect({ x: 39, y: 52 }).to.deep.equal(colision)
        done()
    }).timeout(100)

    it('step2', (done) => {
        let world = load("input13.txt")
        let colision = moveUntilEvent(world, false)
        expect({ x: 133, y: 146 }).to.deep.equal(colision)
        done()
    }).timeout(8000)

})

const empty = 0
const straight_horizontal = 1
const straight_vertical = 2
const turn_slash = 3
const turn_backslash = 4
const intersection = 5

const left = 0
const up = 1
const right = 2
const down = 3
const directionLength = 4

const directionModifier = [-1, 0, 1]

function dirToStr(direction) {
    switch(direction) {
        case left:
            return "left"
        case up:
            return "up"
        case right:
            return "right"
        case down:
            return "down"
        default:
            return direction
    }
}

function changeDirectionForIntersection(cart) {
    cart.direction = (directionLength + cart.direction + directionModifier[cart.turnIndex%directionModifier.length]) % directionLength
    cart.turnIndex++
}

function changeDirectionForTurn(cart, targetSquare) {
    switch(cart.direction) {
        case up:
            if (targetSquare == turn_slash) {
                cart.direction = right
            } else {
                cart.direction = left
            } 
            break
        case down:
            if (targetSquare == turn_slash) {
                cart.direction = left
            } else {
                cart.direction = right
            } 
            break
        case right:
            if (targetSquare == turn_slash) {
                cart.direction = up
            } else {
                cart.direction = down
            } 
            break
        case left:
            if (targetSquare == turn_slash) {
                cart.direction = down
            } else {
                cart.direction = up
            } 
            break
    }
}

function moveUntilEvent(world, stopOnFirstCrash) {
    let tick = 0
    let stop = false
    while(!stop) {
        var output = move(tick++, world, stopOnFirstCrash)
        if (output !== undefined) {
            stop = true
        }
    }
    return output
}

function move(tick, world, stopOnFirstCrash) {
    for (let y = 0; y < world.height; y++) {
        for (let x = 0; x < world.width; x++) {
            let cart = world.map[y][x].cart
            if (cart == undefined) {
                continue
            }
            if (cart.tick == tick) {
                continue
            }
            cart.tick = tick
            var xx = x
            var yy = y
            if (world.map[y][x].square == turn_slash || world.map[y][x].square == turn_backslash) {
                //console.log("square",x,y, "is a turn", world.map[y][x].square)
                let oldDirection = cart.direction
                changeDirectionForTurn(cart,world.map[y][x].square)
                //console.log("cart #",cart.id," in ", x, y," was moving ", dirToStr(oldDirection), "is now going", dirToStr(cart.direction))
            }   
            if (world.map[y][x].square == intersection) {
                //console.log("square",x,y, "is an intersection")
                let oldDirection = cart.direction
                changeDirectionForIntersection(cart)
                //console.log("cart #",cart.id," in ", x, y," was moving ", dirToStr(oldDirection), "is now going", dirToStr(cart.direction))
            }    
            switch (cart.direction) {
                case up:
                    yy--
                    break
                case down:
                    yy++
                    break
                case right:
                    xx++
                    break
                case left:
                    xx--
                    break
            }
            world.map[y][x].cart = undefined
            // handle collision
            if (world.map[yy][xx].cart !== undefined) {
                world.cartNumber -= 2
                world.map[yy][xx].cart = undefined
                //console.log("Colision in ",xx,yy)
                if (stopOnFirstCrash) {
                    return {x:xx,y:yy}
                } 
            } else {
                //console.log("Moving cart ",cart.id," from",x,y," to ",xx,yy)
                world.map[yy][xx].cart = cart
            }
        }
    }
    if (!stopOnFirstCrash && (world.cartNumber == 1)) {
        for (let yyy = 0; yyy < world.height; yyy++) {
            for (let xxx = 0; xxx < world.width; xxx++) {
                if (world.map[yyy][xxx].cart !== undefined) {
                    return {x:xxx,y:yyy}
                }
            }
        }
    }
    return undefined
}

function draw(world) {
    console.log(world.width,"x", world.height)
    let carts = []
    for (let y = 0; y < world.height; y++) {
        var line = ""
        for (let x = 0; x < world.width; x++) {
            console.log(x,y)
            line = line + world.map[y][x].square
            if (world.map[y][x].cart !== undefined) {
                carts.push({x:x,y:y,dir:world.map[y][x].cart.direction})
            }
        }
        console.log(line)
    }
    console.log(carts)
}

function load(file) {
    let input = fs.readFileSync(file, 'utf8')
    let lines = input.split("\n")
    let width = lines[0].length
    let y = 0
    let cartId = 0
    let map = []
    let cartNumber = 0
    lines.forEach(line => {
        map[y] = []
        for (let x = 0; x < width; x++) {
            let char = line.charAt(x)
            switch (char) {
                case ' ':
                    map[y][x] = {square: empty}
                    break
                case '-':
                    map[y][x] = {square: straight_horizontal}
                    break
                case '|': 
                    map[y][x] = {square: straight_vertical}
                    break
                case '/': 
                    map[y][x] = {square: turn_slash}
                    break
                case '\\': 
                    map[y][x] = {square: turn_backslash}
                    break
                case '+': 
                    map[y][x] = {square: intersection}
                    break
                case '>': 
                    cartNumber++
                    map[y][x] = {square: straight_horizontal, cart: {id: ++cartId,direction: right, turnIndex:0, tick: -1}}
                    break
                case 'v': 
                    cartNumber++
                    map[y][x] = {square: straight_vertical, cart: {id: ++cartId,direction: down, turnIndex:0, tick: -1}}
                    break
                case '<': 
                    cartNumber++
                    map[y][x] = {square: straight_horizontal, cart: {id: ++cartId,direction: left, turnIndex:0, tick: -1}}
                    break
                case '^': 
                    cartNumber++
                    map[y][x] = {square: straight_vertical, cart: {id: ++cartId,direction: up, turnIndex:0, tick: -1}}
                    break
                default:
                    map[y][x] = {square: empty}
            }
        }
        y++

    })
    return {
        cartNumber: cartNumber,
        width: width,
        height: y,
        map: map
    }
}