var fs = require('fs')

// https://adventofcode.com/2021/day/10

const openings = ['(','[','{','<']
const closings = [')',']','}','>']
const cost = [3, 57, 1197, 25137]
const costs2 = [1, 2, 3, 4]

function load(file) {
    var input = fs.readFileSync(file, 'utf8')

    let lines = []
    input.split('\n').forEach(line => {
        lines.push(line)
    })
    return lines
}


function runDay10(file) {
    let lines = load(file)

    let score = 0

    lines.forEach(line => {
        let cArray = Array.from(line)
        let stack = []
        for (let i = 0; i < cArray.length; i++) {
            if (openings.includes(cArray[i])) {
                stack.push(cArray[i])
            } else {
                let index = closings.indexOf(cArray[i])
                if (stack[stack.length-1] == openings[index]) {
                    stack.pop()
                } else {
                    score += cost[index]
                    break
                }
            }
        }
           
    })
    

    console.log("Step1", file, "->", score)
}


function runDay10s2(file) {
    let lines = load(file)

    let scoreArray = []

   lines.forEach(line => {
        let cArray = Array.from(line)
        let stack = []
        let valid = true
        for (let i = 0; i < cArray.length; i++) {
            if (openings.includes(cArray[i])) {
                stack.push(cArray[i])
            } else {
                let index = closings.indexOf(cArray[i])
                if (stack[stack.length-1] == openings[index]) {
                    stack.pop()
                } else {
                    valid = false
                    break
                }
            }
        }
        
        if (valid) {
            let score = 0
            let closingFix = ""
            for (let i = stack.length-1; i >= 0 ; i--) {
                let index = openings.indexOf(stack[i])
                let closingChar = closings[index]
                score = score * 5 + costs2[index]
                closingFix += closingChar
            }
            scoreArray.push(score)
        }  
   })
   scoreArray = scoreArray.sort(function(a,b){return b-a})
   let scoreIndex = ((scoreArray.length/2) >> 0)
   let finalScore = scoreArray[scoreIndex]
   console.log("Step1", file, "->", finalScore)
}

runDay10('input-test.txt')
runDay10('input.txt')

runDay10s2('input-test.txt')
runDay10s2('input.txt')
