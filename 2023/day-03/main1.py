import re

def sum_part_numbers(schematic):
    total_sum = 0

    for row_idx, row in enumerate(schematic):
        numbers = re.finditer(r"\d+", row)
        for match in numbers:
            number_start, number_end = match.span()
            is_part_number = False

            for y in range(max(0, row_idx - 1), min(len(schematic), row_idx + 2)):  # Check adjacent rows
                for x in range(max(0, number_start - 1), min(len(row), number_end + 1)):  # Check adjacent positions in the row
                    char = schematic[y][x]

                    if not char.isdigit() and char != ".": # If it's a symbol (not digit or dot)
                        is_part_number = True
                        break

            if is_part_number:
                total_sum += int(match.group())  # Add the part number to the total

    return total_sum


# Get the file path to your engine schematic
file_path = "input.txt"  # Replace with the actual file path

# Read the schematic from the file
with open(file_path, "r") as file:
    schematic = file.read().splitlines()

# Calculate the sum of part numbers
part_number_sum = sum_part_numbers(schematic)

print("The sum of all part numbers in the engine schematic is:", part_number_sum)
