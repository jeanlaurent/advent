var fs = require('fs')
var expect = require('chai').expect

describe('loading/drawing', () => {

    it('check width, height', () => {
        let world = load("input15-test1.txt")
        expect(9).to.equal(world.width)
        expect(9).to.equal(world.height)
    })

    it('check goblin exist number', () => {
        let world = load("input15-test1.txt")
        expect(goblin).to.equal(world.map[4][1].mob.race)
        expect(200).to.equal(world.map[4][1].mob.hp)
    })

    it('check elf exist number', () => {
        let world = load("input15-test1.txt")
        expect(elf).to.equal(world.map[4][4].mob.race)
        expect(200).to.equal(world.map[4][4].mob.hp)
    })

    it('check drawing is ok', () => {
        let world = load("input15-test1.txt")
        let rawDrawning = fs.readFileSync("input15-test1.txt", 'utf8')

        expect(rawDrawning).to.equal(draw(world))
    })

    it('check drawing is ok on puzzleInput', () => {
        let world = load("input15.txt")
        let rawDrawning = fs.readFileSync("input15.txt", 'utf8')

        expect(rawDrawning).to.equal(draw(world))
    })

})


describe('round', () => {

    it('find all targets for goblin', () => {
        let world = load("input15-test1.txt")
        expect([{x:4,y:4}]).to.deep.equal(findAllTarget(world,elf))
    })

    it('find all targets for elf', () => {
        let world = load("input15-test1.txt")
        let targets = [{x:1,y:1},{x:4,y:1},{x:7,y:1},{x:1,y:4},{x:7,y:4},{x:1,y:7},{x:4,y:7},{x:7,y:7}]
        expect(targets).to.deep.equal(findAllTarget(world,goblin))
    })

    it('find adjacent 1.1', () => {
        let world = load("input15-test1.txt")
        let targets = findFreeAdjacent(1,1, world)
        expect(targets).to.deep.equal([{x:2,y:1},{x:1,y:2}])
    })

    it('find adjacent of 1.4', () => {
        let world = load("input15-test1.txt")
        let targets = findFreeAdjacent(1,4, world)
        expect(targets).to.deep.equal([{x:1,y:3},{x:2,y:4},{x:1,y:5}])
    })


    it('display targets', () => {
        let world = load("input15-test1.txt")
        console.log(draw(world))
        round(world)
        console.log(draw(world))
        round1 = `#########
#.G...G.#
#...G...#
#...E..G#
#.G.....#
#.......#
#G..G..G#
#.......#
#########`
        expect(round1).to.equal(draw(world))
        round(world)
        console.log(draw(world))
        round2 = `#########
#..G.G..#
#...G...#
#.G.E.G.#
#.......#
#G..G..G#
#.......#
#.......#
#########`
        expect(round2).to.equal(draw(world))
        round(world)
        console.log(draw(world))
        round3 = `#########
#.......#
#..GGG..#
#..GEG..#
#G..G...#
#......G#
#.......#
#.......#
#########`
        expect(round3).to.equal(draw(world))
    })


})

const wall = 0
const open = 1

const elf = 0
const goblin = 1
const races = 2

function raceStr(race) {
    if (race === elf) {
        return "elf"
    }
    return "goblin"
}
function terrainStr(terrain) {
    if (terrain === wall) {
        return "wall"
    }
    return "open"
}


function distance(x,y,xx,yy) {
    return Math.abs(xx - x) + Math.abs(yy - y)
}

function debug(world) {
    for (let y = 0; y < world.height; y++) {
        for (let x = 0; x < world.width; x++) {
            if (world.map[y][x].mob === undefined) {
                console.log(x,y,terrainStr(world.map[y][x].terrain))
            } else {
                console.log(x,y,terrainStr(world.map[y][x].terrain), world.map[y][x].mob)
            }
            
        }
    }   
}

function round(world) {
    console.log("========== ROUND " + world.tick + "============")
    for (let y = 0; y < world.height; y++) {
        for (let x = 0; x < world.width; x++) {
            if (world.map[y][x].mob === undefined) {
                continue
            }
            let currentMob = world.map[y][x].mob
            if (currentMob.tick == world.tick) {
                continue
            }
            let targetRace = (currentMob.race + 1) % races
            // console.log("the", raceStr(currentMob.race),"in", x,y,"is looking for a",raceStr(targetRace), "to kill")
            let enemy = findAdjacentEnemy(x,y,world)
            if (enemy === undefined) {
                let newCoord = move(x,y, world, currentMob, targetRace)
                if (newCoord === undefined) {
                    continue
                }
                let enemyCoord = findAdjacentEnemy(newCoord.x,newCoord.y,world)
                if (enemy !== undefined) {
                    attack(newCoord.x,newCoord.y,enemyCoord.x, enemyCoord.y, world)
                }
            } else {
                attack(x,y,enemy.x, enemy.y, world)
            }
            console.log(draw(world))
            console.log()
        }
    }
    world.tick++
}

function attack(x,y, xx,yy, world) {
    let attacker = world.map[y][x].mob
    let defender = world.map[yy][xx].mob
    defender.hp -=3
    console.log("the", raceStr(attacker.race),"in", x,y, "attack the", raceStr(defender.race), "in", xx,yy, "hp remaining", defender.hp)
}

function move(x,y,world, currentMob, targetRace) {
    let sourcePositions = findFreeAdjacent(x,y, world)
    let targetPositions = []
    let targets = findAllTarget(world, targetRace)
    targets.forEach(target => {
        let newTargets = findFreeAdjacent(target.x, target.y, world)
        newTargets.forEach( nTarget => targetPositions.push(nTarget))
    })
    let movement = findClosestTarget(sourcePositions, targetPositions)
    if (movement === undefined) {
        console.log("the", raceStr(currentMob.race),"in", x,y, " can't move")
        return undefined
    }
    console.log("the", raceStr(currentMob.race),"in", x,y, "will move to", movement.source.x,movement.source.y, "towards", movement.target.x,movement.target.y, "at",movement.distance)
    world.map[y][x].mob = undefined
    currentMob.tick = world.tick
    world.map[movement.source.y][movement.source.x].mob = currentMob
    return {x:movement.source.x, y:movement.source.y}
}

function findClosestTarget(sources, targets) {
    let minDistance = 999
    let target = {}
    let source = {}
    // console.log("\ts", sources)
    // console.log("\tt", targets)
    sources.forEach(aSource => {
        targets.forEach(aTarget => {
            let dist = distance(aSource.x,aSource.y, aTarget.x, aTarget.y)
            // console.log("\t\t>",aSource, aTarget, dist)
            if (dist < minDistance) {
                // console.log("\t\t\t***>",aSource, aTarget, dist)
                target = aTarget
                source = aSource
                minDistance = dist
            }
        })    
    })
    if (minDistance === 999) {
        return undefined
    }
    return {source:source, target:target, distance: minDistance }
}

function findAllTarget(world, race) {
    let targets = []
    for (let y = 0; y < world.height; y++) {
        for (let x = 0; x < world.width; x++) {
            if (world.map[y][x].mob === undefined) {
                continue
            }
            if (world.map[y][x].mob.race == race) {
                targets.push({x:x,y:y})
            }
        }
    }
    return targets
}

function isValid(x,y, world) {
    return x >=0 && x < world.width && y >= 0 && y < world.height
}

function findAdjacentEnemy(x,y, world) {
    let enemy = undefined
    let allAdjacents = [{x:x,y:y-1},{x:x-1,y:y},{x:x+1,y:y},{x:x,y:y+1}]
    allAdjacents.forEach(adj => {
        //console.log("checking", x,y, "against", adj.x, adj.y)
        if (isValid(adj.x, adj.y, world)) {
            if (world.map[adj.y][adj.x].mob !== undefined) {
                if (world.map[y][x].mob.race != world.map[adj.y][adj.x].mob.race) {
                    // console.log("\tFound ONE !",adj.x,adj.y )
                    enemy = {x:adj.x,y:adj.y}
                    return
                } //else { console.log("\tsame race") }
            } //else { console.log("\tno enemy", adj.y,adj.y, world.map[adj.y][adj.x].mob) }
        } //else { console.log("\tinvalid") }
    });
    return enemy
}

function findFreeAdjacent(x,y, world) {
    let adjacents = []
    let allAdjacents = [{x:x,y:y-1},{x:x-1,y:y},{x:x+1,y:y},{x:x,y:y+1}]
    allAdjacents.forEach(adj => {
        if (isValid(adj.x, adj.y, world) && 
            (world.map[adj.y][adj.x].terrain) === open &&
            (world.map[adj.y][adj.x].mob) === undefined
            ) {
                adjacents.push(adj)
            }
    });
    return adjacents
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
                case '#':
                    map[y][x] = {terrain: wall}
                    break
                case '.':
                    map[y][x] = {terrain: open}
                    break
                case 'E':
                    map[y][x] = {terrain: open, mob: {race: elf, hp: 200}}
                    break
                case 'G':
                    map[y][x] = {terrain: open, mob: {race: goblin, hp: 200}}
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
        tick:0
    }
}

function draw(world){
    let lines = ""
    for (let y = 0; y < world.height; y++) {
        for (let x = 0; x < world.width; x++) {
            if (world.map[y][x].terrain == wall) {
                lines += '#'
                continue
            }
            if (world.map[y][x].mob === undefined) {
                lines += '.'
                continue
            }
            if (world.map[y][x].mob.race === elf) {
                lines += 'E'
            } else {
                lines += 'G'
            }
        } 
        if (y != world.height -1) {
            lines += "\n"  
        }
    }
    return lines
}
