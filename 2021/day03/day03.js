var fs = require('fs')

function load(file) {
    var input = fs.readFileSync(file, 'utf8')
    let matrix = []
    input.split("\n").forEach(line => {
        let matrixLine = []
        for (let i = 0; i < line.length; i++) {
            matrixLine.push(+line[i])
            
        }
        matrix.push(matrixLine)
    })
    return matrix
}

function runDay03(file) {
    let matrix = load(file)
    let width = matrix[0].length

    let binaryGammaRate = ""
    let binaryEpsilonRate = ""

    for (let x = 0; x < width; x++) {
        let count = [0,0]
        for (let y = 0; y < matrix.length; y++) {
            count[matrix[y][x]]++
        } 
        if (count[0] > count[1]) {
            binaryGammaRate += "0"
            binaryEpsilonRate += "1"
        } else {
            binaryGammaRate += "1"
            binaryEpsilonRate += "0"
        }
    }

    let gammaRate = parseInt(binaryGammaRate,2)
    let epsilonRate = parseInt(binaryEpsilonRate,2)

    console.log('step1', file, gammaRate * epsilonRate)
}

function runDay03Step2(file) {
    let matrix = load(file)
    let width = matrix[0].length

    let oxygenGeneratorRating = 0
    let co2GeneratorRating = 0

    for (let x = 0; x < width; x++) {
        let count = [0,0]
        for (let y = 0; y < matrix.length; y++) {
            count[matrix[y][x]]++
        } 
        let bit = 1
        if (count[0] > count[1]) {
            bit = 0
        }
        let newMatrix = []
        matrix.forEach(line => {
            if (line[x] == bit) {
                newMatrix.push(line)
            }
        })
        matrix = newMatrix
        if (matrix.length == 1) {
            oxygenGeneratorRating = parseInt(arrayToString(matrix[0]),2)
            break
        }
    }

    matrix = load(file)
    for (let x = 0; x < width; x++) {
        let count = [0,0]
        for (let y = 0; y < matrix.length; y++) {
            count[matrix[y][x]]++
        } 
        let bit = 1
        if (count[0] <= count[1]) {
            bit = 0
        }
        let newMatrix = []
        matrix.forEach(line => {
            if (line[x] == bit) {
                newMatrix.push(line)
            }
        })
        matrix = newMatrix
        if (matrix.length == 1) {
            co2GeneratorRating = parseInt(arrayToString(matrix[0]),2)
            break
        }
    }
    console.log('step2',file,oxygenGeneratorRating*co2GeneratorRating)
}

function arrayToString(array) {
    let binaryString = ""
    array.forEach(number => {
        if(number == 0) {
            binaryString += "0"
        } else {
            binaryString += "1"
        }
    })
    return binaryString
}

runDay03('input-test.txt')
runDay03('input.txt')
runDay03Step2('input-test.txt')
runDay03Step2('input.txt')