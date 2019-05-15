"""
	@Author:	Leeroy P. Williams
	@Date:		15/05/2019
	@Description:	Should display Fizz for multiples of 3 and
			Buzz for multiples of 5, then display FizzBuzz
			for multiples of both
"""

x = [x for x in range(1,101)]		# List comprehensions are much faster than regular list appending

for i in x:
	if (i%3 == 0) and (i%5 == 0):
		print("FizzBuzz")
	elif i%3 == 0:
		print("Fizz")
	elif i%5 == 0:
		print("Buzz")
	else:
		print(i)
