from pulsar.schema import *

class Student(Record):
    def __init__(self, name, age, grade):
        self.name = name
        self.age = age
        self.grade = grade

    name = String()
    age = Integer()
    grade = Integer()

def process(params):
    params.grade = params.grade + 1
    return params