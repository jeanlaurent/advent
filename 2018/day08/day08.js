var fs = require('fs')
 
function load() {
    var input = fs.readFileSync('input08.txt', 'utf8')
    //var input = fs.readFileSync('input08-test.txt', 'utf8')
    return input.split(" ").map(number => +number)
}

function buildNode() {
    let node = {
        child: [],
        metadata: []
    }
    let numberOfChild = numbers[pointer++]
    let numberOfMetadata = numbers[pointer++]
    for (let i = 0; i < numberOfChild; i++) {
        node.child.push(buildNode())
    }
    for (let i = 0; i < numberOfMetadata; i++) {
        node.metadata.push(numbers[pointer++])
    }
    return node
}

function computeStep1(node) {
    let sum = 0
    node.metadata.forEach(metadata => {
        sum += metadata
    })
    node.child.forEach(children =>{
        sum += computeStep1(children)
    })
    return sum
}

function computeStep2(node) {
    if (node === undefined) {
        return 0
    }
    let sum = 0
    if (node.child.length == 0) {
        node.metadata.forEach(metadata => {
            sum += metadata
        })
        return sum    
    }
    // has children
    node.metadata.forEach(metadata => {
        sum += computeStep2(node.child[metadata -1]) //node may not exist
    })
    return sum
}

let numbers = load()
let pointer = 0
let rootNode = buildNode()
console.log("step1 -->", computeStep1(rootNode))
console.log("step2 -->", computeStep2(rootNode))