# Day 14: Scopes

# The absolute difference between two integers, a and b, is written as |a-b|. 
# The maximum absolute difference between two integers in a set of positive integers, element, is the largest 
# absolute difference between any two integers in __elements.

# The Difference class is started for you in the editor. It has a private integer array (Element) for storing non-negative 
# integers, N and a public integer (maximumDifference) for storing the maximum absolute difference.

# Task: Complete the Difference class by writing the following:
#     A class constructor that takes an array of integers as a parameter and saves it to the __element instance variable.
#     A computeDifference method that finds the maximum absolute difference between any numbers in __elements
#     and stores it in the instance variable maximumDifference. 

from xml.dom.minidom import Element

class Difference:
    def __init__(self, a):
        self.__elements = a
        self.maximumDifference = None

	# Add your code here
    def computeDifference(self):
        max = -9999
        for a in self.__elements:
            for b in self.__elements:
                result = abs(a-b)
                if result > max:
                    max = result
        self.maximumDifference = max

# End of Difference class

_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)