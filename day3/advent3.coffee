expect = require('chai').expect

describe 'advent 3.1', ->

  it 'should return 438 for input', ->
    expect(advent3(265149).distance()).to.equal 438

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

   nextDirection: (direction) ->
      switch direction
        when "north" then "west"
        when "west" then "south"
        when "south" then "east"
        when "east" then "north"

   distance: ->
     return Math.abs(@x) + Math.abs(@y)

advent3 = (puzzleInput) ->
  return 0 unless puzzleInput

  current = new Point 0,0,1
  points = [current]
  width = 1

  while current.value < puzzleInput
    direction = "east"
    current = current.to direction
    points.push current
    width = width + 2
    for i in [1..4]
      direction = current.nextDirection direction
      steps = 2
      steps = 1 if i > 1
      while steps < width
        current = current.to direction
        points.push current
        steps++

  return point for point in points when point.value == puzzleInput
