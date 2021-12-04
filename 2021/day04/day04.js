var fs = require('fs')

const height = 5
const width = 5

function load(file) {
    var input = fs.readFileSync(file, 'utf8')
    let numbers = []
    let lines = input.split("\n")
    lines.shift().split(",").forEach(item => {
        numbers.push(+item)
    })

    lines.shift()
    let matrixes = []
    let matrix = []
    lines.forEach(line => {
        if (line.length==0) {
            matrixes.push(matrix)
            matrix = []
            return lines
        } else {
            let matrixLine = []
            line.split(" ").forEach(item => {
                if (item.trim().length != 0) { // string.split sucks in js !
                    matrixLine.push(+item)
                }
            })
            matrix.push(matrixLine)    
        }
    })
    let emptyMatrixes = []    
    for (let i = 0; i < matrixes.length; i++) {
        let matrix = []
        for (let y = 0; y < height; y++) {
            matrix.push([0,0,0,0,0])
        }
        emptyMatrixes.push(matrix)     
    }
    return {numbers: numbers, matrixes:matrixes, score: emptyMatrixes}
}

function runDay04(file) {
    console.log("Step1", file)
    let bingo = load(file)
    
    for (let index = 0; index < bingo.numbers.length; index++) {
        let number = bingo.numbers[index]
        for (let matrixIndex = 0; matrixIndex < bingo.matrixes.length; matrixIndex++) {
            const matrix = bingo.matrixes[matrixIndex];
            for (let y = 0; y < matrix.length; y++) {
                for (let x = 0; x < matrix[y].length; x++) {
                    if (number == matrix[y][x]) {
                        //console.log("Found ", number, "at", x, y, "for matrix", matrixIndex)
                        bingo.score[matrixIndex][y][x] = 1
                    }
                }
            }
        }
        let bingoedMatrixIndex = checkForBingo(bingo)
        if (bingoedMatrixIndex != -1) {
            let bingoFinalSum = computeFinalSum(bingo, bingoedMatrixIndex)
            let finalScore = bingoFinalSum * number
            console.log('> ', bingoFinalSum)
            console.log('> ', finalScore)
            return finalScore
        }
    }
}

function checkForBingo(bingo) {
    // check line
    for (let matrixIndex = 0; matrixIndex < bingo.matrixes.length; matrixIndex++) {
        const matrix = bingo.matrixes[matrixIndex];
        for (let y = 0; y < matrix.length; y++) {
            let count = 0
            for (let x = 0; x < matrix[y].length; x++) {
                if (bingo.score[matrixIndex][y][x] == 0) {
                    break
                }
                count++
            }
            if (count == width) {
                console.log('> ',"BINGO matrix", matrixIndex, " line " , y)
                return matrixIndex
            }
        }
    }

    // check col
    for (let matrixIndex = 0; matrixIndex < bingo.matrixes.length; matrixIndex++) {
        const matrix = bingo.matrixes[matrixIndex];
        for (let x = 0; x < width; x++) {
            let count = 0
            for (let y = 0; y < matrix.length; y++) {    
                if (bingo.score[matrixIndex][y][x] == 0) {
                    break
                }
                count++
            }
            if (count == height) {
                console.log('> ',"BINGO matrix", matrixIndex, " column " , x)
                return matrixIndex
            }
        }
    }

    return -1
}

function computeFinalSum(bingo, matrixIndex) {
    const matrix = bingo.matrixes[matrixIndex];
    let sum = 0
    for (let y = 0; y < matrix.length; y++) {
        for (let x = 0; x < matrix[y].length; x++) {
            if (bingo.score[matrixIndex][y][x] == 0) {
                sum += bingo.matrixes[matrixIndex][y][x]
            }
        }
    }
    return sum
}

function displayMatrix(matrix) {
    for (let y = 0; y < matrix.length; y++) {
        let line = ""
        for (let x = 0; x < matrix[y].length; x++) {
            if (matrix[y][x] == 0) {
                line += "."
            } else {
                line += "X"
            }
        }
        console.log(line)
    }
}

function runDay04Step2(file) {
    console.log("Step2", file)
    let bingo = load(file)

    console.log("we start with ", bingo.matrixes.length, " matrixes ")
    for (let index = 0; index < bingo.numbers.length; index++) {
        let number = bingo.numbers[index]
        console.log("calling", number)
        for (let matrixIndex = 0; matrixIndex < bingo.matrixes.length; matrixIndex++) {
            const matrix = bingo.matrixes[matrixIndex];
            for (let y = 0; y < matrix.length; y++) {
                for (let x = 0; x < matrix[y].length; x++) {
                    if (number == matrix[y][x]) {
                        bingo.score[matrixIndex][y][x] = 1
                    }
                }
            }
        }
        while(true) { // argh
            let bingoedMatrixIndex = checkForBingo(bingo)
            if (bingoedMatrixIndex != -1) {
                if (bingo.matrixes.length == 1) {
                    let finalSum = computeFinalSum(bingo,bingoedMatrixIndex)
                    let finalScore = finalSum * number
                    console.log('> ', finalSum)
                    console.log('> ', finalScore)
                    return finalScore
                } else {
                    bingo.matrixes.splice(bingoedMatrixIndex, 1)
                    bingo.score.splice(bingoedMatrixIndex, 1)
                    console.log("we got now ", bingo.matrixes.length, "matrixes left ")
                }            
            } else {
                break
            }    
        }
    }
}

runDay04('input-test.txt')
runDay04('input.txt')
runDay04Step2('input-test.txt')
runDay04Step2('input.txt')