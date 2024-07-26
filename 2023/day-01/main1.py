"""
def sum_of_calibration_values(document):
    total_sum = 0

    for line in document:
        first_digit = None
        last_digit = None

        # Find the first digit
        for char in line:
            if char.isdigit():
                first_digit = char
                break

        # Find the last digit
        for char in reversed(line):
            if char.isdigit():
                last_digit = char
                break

        # If both digits are found, form the two-digit number and add to the sum
        if first_digit is not None and last_digit is not None:
            calibration_value = int(first_digit + last_digit)
            total_sum += calibration_value

    return total_sum

# Read the file input.txt
with open('input.txt', 'r') as file:
    document = file.readlines()

# Strip newline characters from each line
document = [line.strip() for line in document]

# Calculate the sum of calibration values
result = sum_of_calibration_values(document)
print(result)
"""
# optimised

import re

def find_calibration_value(line):
    digits = re.findall(r'\d', line)  # Extract all digits from the line
    return int(digits[0] + digits[-1])  # Combine first and last digit

def calculate_total(filename):
    total = 0
    with open(filename, 'r') as file:
        for line in file:
            total += find_calibration_value(line)
    return total

if __name__ == "__main__":
    filename = "input.txt"  # Replace with the actual filename of your input
    result = calculate_total(filename)
    print("The sum of all calibration values is:", result)
