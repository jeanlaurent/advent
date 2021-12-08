var fs = require('fs')

// https://adventofcode.com/2021/day/8
// length of string for each number
// 1 --> 2
// 7 --> 3
// 4 --> 4

// 2 --> 5
// 3 --> 5
// 5 --> 5

// 6 --> 6
// 0 --> 6
// 9 --> 6

// 8 --> 7

function runDay08(file) {
    var input = fs.readFileSync(file, 'utf8')
    
    let n1478 = 0
    input.split('\n').forEach(line => {
        line.split('|')[1].split(" ").forEach(combination => {
            const length = combination.length
            if ((length == 2) || (length == 3) || (length == 4) || (length == 7)) {
                n1478++
            }
        })
    })
    console.log("Step1", file, "->", n1478)
}

function loadLeft(file) {
    var input = fs.readFileSync(file, 'utf8')
    
    let combinationLines = []
    input.split('\n').forEach(line => {
        let combinationLine = []
        line.split('|')[0].split(" ").forEach(combination => {
            combinationLine.push([...combination].sort())
        })
        combinationLines.push(combinationLine)
    })
    return combinationLines
}

function loadRight(file) {
    var input = fs.readFileSync(file, 'utf8')
    
    let combinationLines = []
    input.split('\n').forEach(line => {
        let combinationLine = []
        line.split('|')[1].split(" ").forEach(combination => {
            if (combination.length > 0) {
                combinationLine.push([...combination].sort().join(''))
            }
        })
        combinationLines.push(combinationLine)
    })
    return combinationLines
}

function runDay08s2(file) {
    let lines = loadLeft(file)
    let numbers = loadRight(file)        

    let sum = 0
    for (let lineNumber = 0; lineNumber < lines.length; lineNumber++) {
        const currentLine = lines[lineNumber]
        const currentNumbers = numbers[lineNumber]

        // Identify top segment by comparing 1 and 7
        let size2 = findStringWithLength(currentLine, 2)
        let size3 = findStringWithLength(currentLine, 3)
        let sTop = size3.filter(v => !size2.includes(v))[0]

        // find topRight and bottomRight by comparing segment in 1 and all segment
        // in size 6 letters (6,0,9), since topright should appear twice and bottomright 3 times
        // we have found 3 segments out of 7 at that stage
        let size6 = currentLine.filter(a => a.length == 6)
        let sTopRight = size2[1]
        let sBottomRight = size2[0]
        if (findOccurenceOfElementInMultipleArray(size6, size2[0]) == 2) {
            sTopRight = size2[0]
            sBottomRight = size2[1]
        }

        // find center and topleft by comparing segments in 4 which are not in found in 2
        // and comparing them again with size 6 letters (6,0,9). topleft should appear 3 times 
        // and middle only twice
        // we have found 5 segments out of 7 at that stage
        let size4 = findStringWithLength(currentLine, 4)
        let remainingFromFour = size4.filter(v => !size2.includes(v))
        let sCenter = remainingFromFour[1]
        let sTopLeft = remainingFromFour[0]
        if (findOccurenceOfElementInMultipleArray(size6, remainingFromFour[0]) == 2) {
            sCenter = remainingFromFour[0]
            sTopLeft = remainingFromFour[1]
        }

        // we compare all the letters we haven't found yet with size 6
        // the only one which appears 3 times is the bottom one
        // the 7th letter left is then bottomLeft
        let foundLetters =  [sTop, sTopRight, sBottomRight, sTopLeft, sCenter]
        let size7 = findStringWithLength(currentLine, 7)
        let remainingFromSeven = size7.filter(v => !foundLetters.includes(v))
        let sBottom = remainingFromSeven[1]
        let sBottomLeft = remainingFromSeven[0]
        if (findOccurenceOfElementInMultipleArray(size6, remainingFromSeven[0]) == 3) {
            sBottom = remainingFromSeven[0]
            sBottomLeft = remainingFromSeven[1]
        }
        // we rebuild all the letters sorting them alphabetically to help comparison 
        // and assigning their corresponding number
        let solution = []
        // 0
        solution.push([sTop, sTopLeft, sTopRight, sBottomLeft, sBottomRight, sBottom].sort().join(''))
        // 1
        solution.push([sTopRight, sBottomRight].sort().join(''))
        // 2
        solution.push([sTop,sTopRight,sCenter, sBottomLeft, sBottom].sort().join(''))
        // 3
        solution.push([sTop, sTopRight, sCenter, sBottomRight, sBottom].sort().join(''))
        // 4
        solution.push([sTopLeft, sTopRight, sCenter, sBottomRight].sort().join(''))
        // 5
        solution.push([sTop, sTopLeft, sCenter, sBottomRight, sBottom].sort().join(''))
        // 6
        solution.push([sTop, sTopLeft, sCenter, sBottomRight, sBottom, sBottomLeft].sort().join(''))
        // 7
        solution.push([sTop, sTopRight, sBottomRight].sort().join(''))
        // 8
        solution.push([sTop, sTopLeft,sTopRight, sCenter, sBottomRight, sBottom, sBottomLeft].sort().join(''))
        // 9
        solution.push([sTop, sTopLeft,sTopRight, sCenter, sBottomRight, sBottom].sort().join(''))
        
        // We check all the numbers on the right of the | in the input
        // We compute their value by comparing with the solution we built
        let finalValue = ""
        currentNumbers.forEach(number => {
            for (let i = 0; i < solution.length; i++) {
                if (number == solution[i]) {
                    finalValue += i
                }
            }
        })
        //console.log("Line", lineNumber, finalValue)
        sum += +finalValue
    }
    console.log("step2", file, "->" ,sum)
}

function findOccurenceOfElementInMultipleArray(lines, targetLetter) {
    let occurence = 0
    lines.forEach(line => {
        line.forEach(letter => {
            if (letter == targetLetter) {
                occurence++
            }
        })
    })
    return occurence
}

function findStringWithLength(lines,desiredSize) {
    for (let i = 0; i < lines.length; i++) {
        if (lines[i].length == desiredSize) {
            return lines[i]
        }
    }
    return "z"
}

runDay08('input-test.txt')
runDay08('input.txt')

runDay08s2('input-single.txt')
runDay08s2('input-test.txt')
runDay08s2('input.txt')
