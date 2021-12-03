var fs = require('fs')

function load(file) {
    var input = fs.readFileSync(file, 'utf8')
    let points = []
    input.split("\n").forEach(line => {
        let items = line.split(" ")
        points.push({type:items[0], distance: +items[1]})
    })
    return points
}

function runDay02(file) {
    let actions = load(file)
    let depth = 0
    let x = 0
    actions.forEach(action => {
        switch(action.type) {
            case "forward": x+= action.distance; break;
            case "up": depth -= action.distance; break;
            case "down": depth += action.distance; break;
        }
    })
    console.log("step1", file, x*depth)
}

function runDay02Step2(file) {
    let actions = load(file)
    let depth = 0
    let x = 0
    let aim = 0
    actions.forEach(action => {
        switch(action.type) {
            case "forward": 
                x += action.distance; 
                depth += aim*action.distance;
                break;
            case "up": 
                aim -= action.distance;
                break;
            case "down": aim += action.distance;
            break;
        }
    })
    console.log("step2", file, x*depth)
}

runDay02('input-test.txt')
runDay02('input.txt')
runDay02Step2('input-test.txt')
runDay02Step2('input.txt')