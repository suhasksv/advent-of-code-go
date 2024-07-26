import re

def sum_gear_ratios(schematic):
    total_sum = 0
    symbol_regex = re.compile(r"[^\d.]")

    for row_idx, row in enumerate(schematic):
        for col_idx, char in enumerate(row):
            if symbol_regex.match(char):
                adjacent_numbers = []
                for dy in [-1, 0, 1]:
                    for dx in [-1, 0, 1]:
                        new_row = row_idx + dy
                        new_col = col_idx + dx
                        if (
                                0 <= new_row < len(schematic) and
                                0 <= new_col < len(row) and
                                (new_row != row_idx or new_col != col_idx) # Exclude the symbol itself
                        ):
                            match = re.search(r"\d+", schematic[new_row][new_col:])
                            if match and match.start() == 0: # ensure the number starts at the current position
                                number = int(match.group())
                                adjacent_numbers.append(number)

                if len(adjacent_numbers) == 2:
                    gear_ratio = adjacent_numbers[0] * adjacent_numbers[1]
                    print(f"Found gear at position ({row_idx}, {col_idx}) with ratio: {gear_ratio}")
                    total_sum += gear_ratio

    return total_sum


# Read the schematic from the file
file_path = "input.txt"
with open(file_path, "r") as file:
    schematic = file.read().splitlines()

# Calculate the sum of gear ratios
gear_ratio_sum = sum_gear_ratios(schematic)
print("The sum of all gear ratios in the engine schematic is:", gear_ratio_sum)

