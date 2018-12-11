var expect = require('chai').expect

describe('3rdNumber', () => {

    it('should finf 0 for 10', () => {
        expect(0).to.equal(get3RdNumber(10))
    })

    it('should find 3 for 12355', () => {
        expect(3).to.equal(get3RdNumber(12345))
    })


    it('should find 2 for 987654231', () => {
        expect(2).to.equal(get3RdNumber(987654231))
    })
})

describe('Grid', () => {

  it('fuel cell at 3,5 in a grid with serial number 8, the power level of this fuel cell is 4', () => {
    expect(4).to.equal(computeGrid(8)[5][3])
  })
  
  it('122,79, grid serial number 57: power level -5.', () => {
    expect(-5).to.equal(computeGrid(57)[79][122])
  })

  it('217,196, grid serial number 39: power level  0', () => {
    expect(0).to.equal(computeGrid(39)[196][217])
  })

  it('101,153, grid serial number 71: power level  4', () => {
    expect(4).to.equal(computeGrid(71)[153][101])
  })

  it('step1 --> 34,13', () => {
    expect({x:34,y:13, largest:30}).to.deep.equal(findBiggest(computeGrid(1723), 3))
  })

  it('step2 --> 34,13', (done) => {
    expect({x:280,y:218, largest:68, squareWidth: 11}).to.deep.equal(findBiggestIncrementingSize(computeGrid(1723)))
    done()
  }).timeout(36000)

})

function get3RdNumber(powerLevel) {
    if (powerLevel < 100) {
        return 0
    } else {
        powerLevelAsString = powerLevel.toString() // pretty ugly, a mathematical operation should be *way* faster
        return +(powerLevelAsString.charAt(powerLevelAsString.length-3))            
    }
}

function computeGrid(puzzleId) {
    var grid = new Array(301).fill(0).map(() => new Array(301).fill(0));

    for (let y = 1; y <= 300; y++) {
        for (let x = 1; x <= 300; x++) {
            var rackId = x + 10
            var powerLevel = (rackId * y + puzzleId) * rackId
            var hundredDigit = get3RdNumber(powerLevel) - 5
            grid[y][x] = hundredDigit
        }
    }
    return grid
}

function findBiggestIncrementingSize(grid) {
    var largestEver = -9999
    var largestSquareEver = {}
    var squareWidth = 0
    for (let squareWidth = 1; squareWidth <= 300; squareWidth++) {
        if (squareWidth%50===0) {
            console.log("width --> ",squareWidth)    
        }
        var biggestSquare = findBiggest(grid, squareWidth)
        if (biggestSquare.largest > largestEver) {
            largestEver = biggestSquare.largest
            largestSquareEver = biggestSquare
            largestSquareEver.squareWidth = squareWidth
        }
    }
    return largestSquareEver
}

function findBiggest(grid, squareWidth) {
    var largest = -9999
    var xLargest = -999
    var yLargest = -999
    for (let y = 0; y < 300 - squareWidth; y++) {
        for (let x = 0; x < 300 - squareWidth; x++) {
            let totalPower = sum(x,y, grid, squareWidth)
            if (totalPower > largest) {
                largest=totalPower
                xLargest = x
                yLargest = y
            }
        }
    }
    return {x:xLargest, y:yLargest, largest: largest}
}


function sum(x,y, grid, squareWidth) {
    var sum = 0
    for (let yy = 0; yy < squareWidth; yy++) {
        for (let xx = 0; xx < squareWidth; xx++) {
            sum += grid[y + yy][x + xx]
        }
    }
    return sum
}