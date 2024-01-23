import random
import nltk
from nltk.corpus import words

def content(fileType):
    if fileType == "a":
        for i in range(random.randint(3, 15)):
            arit_expression()
    elif fileType == "b":
        for i in range(random.randint(3, 15)):
            print(random.choice(words.words()))
    elif fileType == "c":
        for i in range(random.randint(3, 15)):
            binary_digits = [str(random.randint(0, 1)) for _ in range(random.randint(8, 30))]
            binary_num = "".join(binary_digits)
            print(binary_num)


def arit_expression():
    operators = ['+', '-', '*', '/']
    expression = []
    length = random.randint(3, 9)
    expression.append(str(random.randint(1,10)))
    for i in range(length - 2):
        expression.append(random.choice(operators))  # Add a random operator
        expression.append(str(random.randint(1, 10)))

    expression.append(str(random.randint(1, 10)))

    result = "".join(expression)
    print(result)
    return result


content("c")