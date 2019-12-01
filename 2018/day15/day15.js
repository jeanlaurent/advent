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

    it('First Rounds', () => {
        let world = load("input15-test1.txt")
        round(world)
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

    it('second rounds', () => {
        let world = load("input15-test2.txt")
        round(world)
        round1 = `#######
#..G..#
#...EG#
#.#G#G#
#...#E#
#.....#
#######`
        expect(round1).to.equal(draw(world))
        debug(world)

        round(world)
        round2 = `#######
#...G.#
#..GEG#
#.#.#G#
#...#E#
#.....#
#######`
        expect(round2).to.equal(draw(world))
        debug(world)
        for (let index = 0; index < 22; index++) {
            round(world)
            debug(world)
        }
    for (let index = 0; index < 5; index++) {
        round(world)
        debug(world)
    }
    round28 = `#######
#G....#
#.G...#
#.#.#G#
#...#E#
#....G#
#######`
    expect(round28).to.equal(draw(world))
        stop = round(world)
        while(!stop) {
            stop = round(world)
            debug(world)
        }
        expect(27730).to.equal(computeScore(world))
    })

    it('test3', () => { //error of a single round
        let world = load("input15-test3.txt")
        stop = round(world)
        while(!stop) {
            stop = round(world)
        }
        expect(36334).to.equal(computeScore(world))
    })   

    it('test4', () => { //error of a single round
        let world = load("input15-test4.txt")
        stop = round(world)
        while(!stop) {
            stop = round(world)
        }
        expect(39514).to.equal(computeScore(world))
    })  

    it('test5', () => { //error of a single round
        let world = load("input15-test5.txt")
        stop = round(world)
        while(!stop) {
            stop = round(world)
        }
        expect(27755).to.equal(computeScore(world))
    })  

    it('test6', () => { //error of a single round
        let world = load("input15-test6.txt")
        stop = round(world)
        while(!stop) {
            stop = round(world)
        }
        expect(28944).to.equal(computeScore(world))
    })  

   //176233 too low
   //179220 too high
    it('step1', () => {
        let world = load("input15.txt")
        stop = round(world)
        while(!stop) {
            stop = round(world)
        }
        expect(178003).to.equal(computeScore(world))
    })    
})

describe('bfs', () => {

    it('bfs - step1', () => {
        let world = load("input-bfs.txt")
        expect(8).to.equal(distanceBFS(5,1,5,4, world))
        expect(8).to.equal(distanceBFS(4,2,5,4, world))
        expect(10).to.equal(distanceBFS(3,1,5,4, world))
    })

    it('bfs - step2', () => {
        let world = load("input-bfs.1.txt")
        expect(8).to.equal(distanceBFS(5,1,7,2, world))
        expect(8).to.equal(distanceBFS(4,2,7,2, world))
        expect(10).to.equal(distanceBFS(3,1,7,2, world))
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

function distanceBFS(startX,startY,goalX,goalY, world) {
    let visited = []
    let queue = [{x:startX, y:startY, w:1}]
    while(true){
        // console.log("---------round " + tick)
        // console.log("visited", visited)
        // console.log("queue", queue)
        let node = queue.shift()
        // console.log("node", node)
        if (node === undefined) {
            return 9999
        }
        if (node.x == goalX && node.y == goalY) {
            return node.w
        }
        let adjacents = findFreeAdjacent(node.x,node.y, world)
        adjacents.forEach( adj => {
            if (!visited.includes(adj.x+":"+adj.y)) {
                // console.log(adj.x,adj.y,"not visited")
                if (!inQueue(adj.x,adj.y, queue)) {
                    // console.log(adj.x,adj.y,"not in queue")
                    queue.push({x:adj.x, y:adj.y, w:node.w+1})
                } else {
                    // console.log(adj.x,adj.y,"in queue")
                }
            } else {
                // console.log(adj.x,adj.y,"already visited")
            }
        })
        visited.push(node.x+":"+node.y)    
    }
}

function inQueue(xx,yy, queue) {
    let index = 0
    while(index < queue.length) {
        if (xx == queue[index].x && yy == queue[index].y) {
            return true
        }
        index++
    }
    return false
}

function debug(world) {
    for (let y = 0; y < world.height; y++) {
        for (let x = 0; x < world.width; x++) {
            if (world.map[y][x].mob === undefined) {
                // console.log(x,y,terrainStr(world.map[y][x].terrain))
            } else {
                console.log(x,y,terrainStr(world.map[y][x].terrain), world.map[y][x].mob)
            }
            
        }
    }   
}

function computeScore(world) {
    let sumHp = 0
    for (let y = 0; y < world.height; y++) {
        for (let x = 0; x < world.width; x++) {
            if (world.map[y][x].mob !== undefined) {
                //console.log("adding", world.map[y][x].mob.hp)
                sumHp += world.map[y][x].mob.hp
            }
        }
    }
    console.log(sumHp, "*", world.tick)
    return sumHp*world.tick
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
                if (enemyCoord !== undefined) {
                    attack(newCoord.x,newCoord.y,enemyCoord.x, enemyCoord.y, world)
                }
            } else {
                attack(x,y,enemy.x, enemy.y, world)
            }
      
            // console.log(draw(world))
            console.log()
        }
    }
    if (world.survivors[elf] == 0 || world.survivors[goblin] == 0) {
        console.log("end game")
        return true
    }  
    world.tick++

    console.log(draw(world))
    return false
}

function attack(x,y, xx,yy, world) {
    world.map[yy][xx].mob.hp -= 3
    console.log("the", raceStr(world.map[y][x].mob.race),"(hp:", world.map[y][x].mob.hp,") in", x,y, "attack the", raceStr(world.map[yy][xx].mob.race), "in", xx,yy, "hp remaining", world.map[yy][xx].mob.hp)
    if (world.map[yy][xx].mob.hp <= 0) {
        console.log("the", raceStr(world.map[yy][xx].mob.race),"in", xx,yy, "just died")
        world.survivors[world.map[yy][xx].mob.race]--
        world.map[yy][xx].mob = undefined
    }
}

function move(x,y,world, currentMob, targetRace) {
    let sourcePositions = findFreeAdjacent(x,y, world)
    let targetPositions = []
    let targets = findAllTarget(world, targetRace)
    targets.forEach(target => {
        let newTargets = findFreeAdjacent(target.x, target.y, world)
        newTargets.forEach( nTarget => targetPositions.push(nTarget))
    })
    let movement = findClosestTarget(sourcePositions, targetPositions, world)
    if (movement === undefined) {
        console.log("the", raceStr(currentMob.race),"(hp:", currentMob.hp,") in", x,y, " can't move")
        return undefined
    }
    console.log("the", raceStr(currentMob.race),"(hp:", currentMob.hp,") in", x,y, "will move to", movement.source.x,movement.source.y, "towards", movement.target.x,movement.target.y, "at",movement.distance)
    world.map[y][x].mob = undefined
    currentMob.tick = world.tick
    world.map[movement.source.y][movement.source.x].mob = currentMob
    return {x:movement.source.x, y:movement.source.y}
}

function findClosestTarget(sources, targets, world) {
    let minDistance = 999
    let target = {}
    let source = {}
    // console.log("\ts", sources)
    // console.log("\tt", targets)
    sources.forEach(aSource => {
        targets.forEach(aTarget => {
            //let dist = distance(aSource.x,aSource.y, aTarget.x, aTarget.y)
            let dist = distanceBFS(aSource.x,aSource.y, aTarget.x, aTarget.y, world)
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
    let minHitPoint = 999
    let enemy = undefined
    let allAdjacents = [{x:x,y:y-1},{x:x-1,y:y},{x:x+1,y:y},{x:x,y:y+1}]
    allAdjacents.forEach(adj => {
        //console.log("checking", x,y, "against", adj.x, adj.y)
        if (isValid(adj.x, adj.y, world)) {
            if (world.map[adj.y][adj.x].mob !== undefined) {
                if (world.map[y][x].mob.race != world.map[adj.y][adj.x].mob.race) {
                    if (minHitPoint > world.map[adj.y][adj.x].mob.hp) {
                        enemy = {x:adj.x,y:adj.y}
                        minHitPoint = world.map[adj.y][adj.x].mob.hp
                        console.log("targetting",adj.x, adj.y, world.map[adj.y][adj.x].mob)
                    }
                    // console.log("\tFound ONE !",adj.x,adj.y )
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
    let survivors = [0,0]
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
                    survivors[elf]++
                    break
                case 'G':
                    map[y][x] = {terrain: open, mob: {race: goblin, hp: 200}}
                    survivors[goblin]++
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
        tick: 0,
        survivors: survivors
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
