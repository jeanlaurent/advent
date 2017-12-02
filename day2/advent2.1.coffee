expect = require('chai').expect

describe 'advent 2.1', ->

  it 'should return 0 when no input', ->
    expect( advent21() ).to.equal 0

  it 'should return 8 when input is 9 1', ->
    expect( advent21 '9 1' ).to.equal 8

  it 'should return 7 when input is 8 1 3 7 8 4 2', ->
    expect( advent21 '8 1 3 7 8 4 2' ).to.equal 7

  it 'should return 8 when input is 5 1 9 5', ->
    expect( advent21 '5 1 9 5' ).to.equal 8

  it 'should return 4 when input is 7 5 3', ->
    expect( advent21 '7 5 3' ).to.equal 4

  it 'should return 6 when input is 2 4 6 8', ->
    expect( advent21 '2 4 6 8' ).to.equal 6

  it 'should return 15 when input is 9 1 / 8 1', ->
    input = '''
    9 1
    8 1
    '''
    expect( advent21 input ).to.equal 15

  it 'should return 15 when input is input from advent', ->
    input = '''
    5 1 9 5
    7 5 3
    2 4 6 8
    '''
    expect( advent21 input ).to.equal 18

describe 'advent 2.2', ->

  it 'should return 4 when input is 5 9 2 8', ->
    expect( advent22 '5 9 2 8' ).to.equal 4

  it 'should return 3 when input is 9 4 7 3', ->
    expect( advent22 '9 4 7 3' ).to.equal 3

  it 'should return 2 when input is 3 8 6 5', ->
    expect( advent22 '3 8 6 5' ).to.equal 2

  it 'should return 15 when input is input from advent', ->
    input = '''
    5 9 2 8
    9 4 7 3
    3 8 6 5
    '''
    expect( advent22 input ).to.equal 9

advent21 = (captcha) ->
  return 0 unless captcha
  sum = 0
  lines = captcha.split '\n'
  for line in lines
    numbers = line.split ' '
    smallest = 9999
    biggest = -1
    for number in numbers
      num = parseInt(number)
      smallest = num if smallest > num
      biggest = num if biggest < num
    sum += biggest - smallest
  return sum

divideIfGreater = (a,b) ->
  if a > b
    return divide a, b
  else
    return divide b, a


divide = (a, b) ->
  return a / b if a % b == 0
  0

advent22 = (captcha) ->
  return 0 unless captcha
  sum = 0
  lines = captcha.split '\n'
  for line in lines
    numbers = line.split(' ').map (number) -> parseInt number
    for number, i in numbers
      for number2, j in numbers when i < j
        sum += divideIfGreater number, number2
  return sum



puzzleInput = '''
4168 3925 858 2203 440 185 2886 160 1811 4272 4333 2180 174 157 361 1555
150 111 188 130 98 673 408 632 771 585 191 92 622 158 537 142
5785 5174 1304 3369 3891 131 141 5781 5543 4919 478 6585 116 520 673 112
5900 173 5711 236 2920 177 3585 4735 2135 2122 5209 265 5889 233 4639 5572
861 511 907 138 981 168 889 986 980 471 107 130 596 744 251 123
2196 188 1245 145 1669 2444 656 234 1852 610 503 2180 551 2241 643 175
2051 1518 1744 233 2155 139 658 159 1178 821 167 546 126 974 136 1946
161 1438 3317 4996 4336 2170 130 4987 3323 178 174 4830 3737 4611 2655 2743
3990 190 192 1630 1623 203 1139 2207 3994 1693 1468 1829 164 4391 3867 3036
116 1668 1778 69 99 761 201 2013 837 1225 419 120 1920 1950 121 1831
107 1006 92 807 1880 1420 36 1819 1039 1987 114 2028 1771 25 85 430
5295 1204 242 479 273 2868 3453 6095 5324 6047 5143 293 3288 3037 184 987
295 1988 197 2120 199 1856 181 232 564 1914 1691 210 1527 1731 1575 31
191 53 714 745 89 899 854 679 45 81 726 801 72 338 95 417
219 3933 6626 2137 3222 1637 5312 238 5895 222 154 6649 169 6438 3435 4183
37 1069 166 1037 172 258 1071 90 497 1219 145 1206 143 153 1067 510
'''
console.log '1:', advent21 puzzleInput
console.log '2:', advent22 puzzleInput
