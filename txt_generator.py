import random
from nltk.corpus import words
import os
import re

def content_generator(fileType):
    content = ""
    if fileType == "a":
        for i in range(random.randint(3, 15)):
            content += arit_expression() + "\n"
    elif fileType == "b":
        for i in range(random.randint(3, 15)):
            content += random.choice(words.words()) + "\n"
    elif fileType == "c":
        for i in range(random.randint(3, 15)):
            binary_digits = [str(random.randint(0, 1)) for _ in range(random.randint(8, 30))]
            binary_num = "".join(binary_digits)
            content += binary_num + "\n"
    return content


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
    return result

def sort_key(file):
    match = re.match(r'test(\d+)\.txt', file)
    return (int(match.group(1)), file) if match else (float('inf'), file)

def file_generator(type):
    directory_path = "files"
    files = os.listdir(directory_path)
    if not files:
        number = 0
    else:
        files = sorted(files, key = sort_key)
        last_file = files[-1]
        number = int(last_file[4:-4]) + 1
    file_path = f"files/test{number}.txt"
    with open(file_path, "w") as file:
        file.write(content_generator(type))



def main():
    number_files = input("Enter hoy many files you want to add:")
    types = ["a", "b", "c"]
    type_input = ""
    while type_input not in types:
        type_input = input("Enter the type of file you want to generate:").lower().strip()
    try:
        for i in range(int(number_files)):
            file_generator(type_input)
    except:
        print("Error: something went wrong, try again")

main()