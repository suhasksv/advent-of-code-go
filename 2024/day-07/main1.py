from itertools import product

def parse_input(input_text):
    equations = []
    for line in input_text.strip().split("\n"):
        test_value, numbers = line.split(": ")
        test_value = int(test_value)
        numbers = list(map(int, numbers.split()))
        equations.append((test_value, numbers))
    return equations

def evaluate_left_to_right(numbers, operators):
    result = numbers[0]
    for i, op in enumerate(operators):
        if op == "+":
            result += numbers[i + 1]
        elif op == "*":
            result *= numbers[i + 1]
    return result

def is_valid_equation(test_value, numbers):
    num_operators = len(numbers) - 1
    for operator_combination in product(["+", "*"], repeat=num_operators):
        if evaluate_left_to_right(numbers, operator_combination) == test_value:
            return True
    return False

def calculate_total_calibration(input_text):
    equations = parse_input(input_text)
    total_calibration = 0

    for test_value, numbers in equations:
        if is_valid_equation(test_value, numbers):
            total_calibration += test_value

    return total_calibration

# Read input from a file named 'input.txt'
with open('input.txt', 'r') as file:
    input_text = file.read()

# Calculate and Print Total Calibration
print(calculate_total_calibration(input_text))

