var fs = require('fs')
 
function load() {
    return fs.readFileSync('input05.txt', 'utf8')
    //return fs.readFileSync('input05-test.txt', 'utf8')
}

function computeStep1(letters) {
    let index = 0
    while(index < letters.length - 1) {
        if (letters[index].toLowerCase() === letters[index+1].toLowerCase()) {
            if (letters[index] !== letters[index+1]) {
                letters = letters.substring(0,index) + letters.substring(index+2)
                index = index - 1
                if (index < 0) {
                    index = 0
                }
                continue
            }
        }
        index++
    }
    return letters.length
}

function computeStep2(letters) {
    let lowerCaseletters = "abcdefghijklmnopqrstuvwxyz"
    let upperCaseletters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    let min = 99999
    for (let charIndex = 0; charIndex < lowerCaseletters.length; charIndex++) {    
        let index = 0
        let slicedLetters = letters.slice(0,letters.length)
        while(index < letters.length - 1) {
            if (slicedLetters[index] == lowerCaseletters[charIndex] || slicedLetters[index] == upperCaseletters[charIndex]) {
                slicedLetters = slicedLetters.substring(0,index) + slicedLetters.substring(index+1)
            } else {
                index++
            }
        }
        let polymerCount = computeStep1(slicedLetters)
        console.log("\t" + upperCaseletters[charIndex] + " --> " + polymerCount)
        if (min > polymerCount) {
            min = polymerCount
        }
    }
    return min
}

let originLetters = load()
console.log("step1")
console.log("step1 -->", computeStep1(originLetters))
console.log()
console.log("step2")
console.log("step2 -->", computeStep2(originLetters))


