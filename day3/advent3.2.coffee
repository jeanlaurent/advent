expect = require('chai').expect

describe 'advent 3.2', ->

  it 'should return 438 for input', ->
    expect(advent32(265149).value).to.equal 266330

  it 'should find adjacent squares', ->
    expect(new Coord(3,3).adjacent()).to.deep.equal [new Coord(2,2),
                                                new Coord(2,3),
                                                new Coord(2,4),
                                                new Coord(3,2),
                                                new Coord(3,4),
                                                new Coord(4,2),
                                                new Coord(4,3),
                                                new Coord(4,4),
                                                ]

class Coord
  constructor: (@x, @y) ->

  adjacent: ->
    adjacents = []
    for i in [-1..1]
      for j in [-1..1]
          adjacents.push new Coord @x + i, @y + j unless i == 0 && j == 0
    adjacents

class Point
   constructor: (@x, @y, @value) ->

   to: (direction) ->
     switch direction
       when "north" then @north()
       when "south" then @south()
       when "east" then @east()
       when "west" then @west()

   east: -> new Point @x+1,@y, @value + 1
   north: -> new Point @x,@y-1, @value + 1
   west: -> new Point @x-1,@y, @value + 1
   south: -> new Point @x,@y+1, @value + 1

   adjacentCoords: ->
     new Coord(@x, @y).adjacent()

   nextDirection: (direction) ->
      switch direction
        when "north" then "west"
        when "west" then "south"
        when "south" then "east"
        when "east" then "north"

   distance: ->
     return Math.abs(@x) + Math.abs(@y)

findValue = (points, x, y) ->
    return point.value for point in points when x == point.x && y == point.y
    0

sumOfAjdacentOf = (point, points) ->
    sum = 0
    for coord in point.adjacentCoords()
      sum += findValue points, coord.x, coord.y
    sum

advent32 = (maxValue) ->
  return 0 unless maxValue

  current = new Point 0,0,1
  points = [current]
  width = 1

  while current.value < maxValue
    direction = "east"
    current = current.to direction
    current.value = sumOfAjdacentOf current, points
    points.push current
    width = width + 2
    for i in [1..4]
      direction = current.nextDirection direction
      steps = 2
      steps = 1 if i > 1
      while steps < width
        current = current.to direction
        current.value = sumOfAjdacentOf current, points
        points.push current
        steps++

  return point for point in points when point.value > maxValue
