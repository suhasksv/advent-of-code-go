import re

word_to_number = {
    'one': 1, 'two': 2, 'three': 3, 'four': 4, 'five': 5,
    'six': 6, 'seven': 7, 'eight': 8, 'nine': 9
}

def calculate_calibration_value(line):
    pattern = r"(?=(\d|one|two|three|four|five|six|seven|eight|nine))"
    matches = re.findall(pattern, line)

    first_digit_str = matches[0]
    last_digit_str = matches[-1]

    # Correctly convert spelled-out digits and handle potential errors
    try:
        first_digit = int(first_digit_str)
    except ValueError:
        first_digit = word_to_number[first_digit_str]

    try:
        last_digit = int(last_digit_str)
    except ValueError:
        last_digit = word_to_number[last_digit_str]

    return first_digit * 10 + last_digit


def calculate_total_calibration_value(filename="input.txt"):
    total = 0
    with open(filename, 'r') as file:
        for line in file:
            total += calculate_calibration_value(line.strip())  # Strip newline characters
    return total

if __name__ == "__main__":
    total = calculate_total_calibration_value()
    print("The sum of all calibration values:", total)
