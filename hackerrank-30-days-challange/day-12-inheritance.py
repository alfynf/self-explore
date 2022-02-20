'''
Day 12: Inheritance

Task: You are given two classes, Person and Student, where Person is the base class and Student is the derived class. 
Completed code for Person and a declaration for Student are provided for you in the editor. 
Observe that Student inherits all the properties of Person.

Complete the Student class by writing the following:
    A Student class constructor, which has 4 parameters:
        A string, firstName.
        A string, lastName.
        An integer, idNumber.
        An integer array (or vector) of test scores, scores.
    A char calculate() method that calculates a Student object's average and returns the grade character representative
'''

class Person:
	def __init__(self, firstName, lastName, idNumber):
		self.firstName = firstName
		self.lastName = lastName
		self.idNumber = idNumber
	def printPerson(self):
		print("Name:", self.lastName + ",", self.firstName)
		print("ID:", self.idNumber)

class Student(Person):
    #   Class Constructor
    #   
    #   Parameters:
    #   firstName - A string denoting the Person's first name.
    #   lastName - A string denoting the Person's last name.
    #   id - An integer denoting the Person's ID number.
    #   scores - An array of integers denoting the Person's test scores.
    #
    # Write your constructor here
    def __init__(self, firstName, lastName, idNumber, scores):
        super().__init__(firstName, lastName, idNumber)
        self.scores = scores

    #   Function Name: calculate
    #   Return: A character denoting the grade.
    #
    # Write your function here
    def calculate(self):
        result = 0
        for nilai in self.scores:
            result = result + nilai
        result = result/len(self.scores)
        if result >= 90:
            return "O"
        elif result >= 80:
            return "E"
        elif result >= 70:
            return "A"
        elif result >= 55:
            return "P"
        elif result >= 40:
            return "D"
        else:
            return "T"
        
firstName = "alfy"
lastName = "nf"
idNum = 2134565
scores = [80,100]
s = Student(firstName, lastName, idNum, scores)
s.printPerson()
print("Grade:", s.calculate())