def solve_trash_compactor(filename):
    try:
        with open(filename, 'r') as f:
            lines = f.readlines()
    except FileNotFoundError:
        print(f"Error: Could not find file '{filename}'")
        return

    # 1. Normalize the grid
    # Remove newlines but keep trailing spaces
    lines = [line.strip('\n') for line in lines]

    if not lines:
        return

    # Find the maximum width to pad lines
    max_width = max(len(line) for line in lines)
    grid = [line.ljust(max_width) for line in lines]
    rows = len(grid)

    # 2. Identify vertical separators
    separator_indices = [-1]

    for col in range(max_width):
        is_empty_column = True
        for row in range(rows):
            if grid[row][col] != ' ':
                is_empty_column = False
                break

        if is_empty_column:
            separator_indices.append(col)

    separator_indices.append(max_width)

    grand_total = 0

    print(f"Detected {len(separator_indices) - 1} separate problems.")

    # 3. Process each vertical slice
    for i in range(len(separator_indices) - 1):
        start_col = separator_indices[i] + 1
        end_col = separator_indices[i+1]

        if start_col >= end_col:
            continue

        # --- PART 2 LOGIC: Vertical Parsing ---
        # Rule 1: Read columns Right-to-Left
        # Rule 2: Digits are Top-to-Bottom
        # Rule 3: Operator is at the bottom

        numbers = []

        # Get the operator from the bottom row of this slice
        # The operator row is the last row (rows - 1)
        operator_segment = grid[rows-1][start_col:end_col].strip()

        if not operator_segment:
            continue

        operator = operator_segment[0] # Should be '+' or '*'

        # Iterate columns from Right to Left
        for c in range(end_col - 1, start_col - 1, -1):
            num_str = ""

            # Read rows from top down to the row BEFORE the operator
            for r in range(rows - 1):
                char = grid[r][c]
                if char.isdigit():
                    num_str += char

            if num_str:
                numbers.append(int(num_str))

        if not numbers:
            continue

        # 4. Calculate result for this block
        if operator == '+':
            result = sum(numbers)
        elif operator == '*':
            result = 1
            for n in numbers:
                result *= n
        else:
            result = 0

        # print(f"Problem {i+1}: {numbers} {operator} = {result}")
        grand_total += result

    print("-" * 30)
    print(f"Grand Total: {grand_total}")

if __name__ == "__main__":
    solve_trash_compactor('input.txt')