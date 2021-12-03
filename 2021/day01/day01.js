var fs = require('fs')

function load(file) {
    var input = fs.readFileSync(file, 'utf8')
    let points = []
    input.split("\n").forEach(line => {
        points.push(+line)
    })
    return points
}

function runDay01(file) {
    let depths = load(file)
    let increase = 0
    let lastDepth = 99999999
    depths.forEach(depth => {
        if (depth>lastDepth) {
            increase++
        }
        lastDepth = depth
    });
    console.log("step1",file, increase)
}

function runDay01Step2(file) {
    let depths = load(file)
    let increase = 0
    let lastMeasurement = 9999999
    for (let i = 0; i < depths.length - 2 ; i++) {
        let sum = depths[i]+depths[i+1]+depths[i+2]
        if (sum > lastMeasurement) {
            increase++
        }
        lastMeasurement = sum
    }
    console.log("step2", file, increase)
}

runDay01('input01-test.txt')
runDay01('input01.txt')
runDay01Step2('input01-test.txt')
runDay01Step2('input01.txt')