def solve_trash_compactor(filename):
    try:
        with open(filename, 'r') as f:
            lines = f.readlines()
    except FileNotFoundError:
        print(f"Error: Could not find file '{filename}'")
        return

    # 1. Normalize the grid
    # Remove newlines but keep trailing spaces as they might be part of the layout
    lines = [line.strip('\n') for line in lines]

    if not lines:
        return

    # Find the maximum width to pad lines
    max_width = max(len(line) for line in lines)
    grid = [line.ljust(max_width) for line in lines]
    rows = len(grid)

    # 2. Identify vertical separators
    # We will slice the grid based on these empty columns.
    separator_indices = [-1] # Start with a virtual separator at -1

    for col in range(max_width):
        is_empty_column = True
        for row in range(rows):
            if grid[row][col] != ' ':
                is_empty_column = False
                break

        if is_empty_column:
            separator_indices.append(col)

    separator_indices.append(max_width) # End with a virtual separator

    grand_total = 0

    # 3. Process each vertical slice
    for i in range(len(separator_indices) - 1):
        start_col = separator_indices[i] + 1
        end_col = separator_indices[i+1]

        # Skip if there's no width (like consecutive separators)
        if start_col >= end_col:
            continue

        # Extract the text for this block
        numbers = []
        operator = None
        has_content = False

        for r in range(rows):
            # Extract the substring for this row in this column range
            segment = grid[r][start_col:end_col].strip()

            if not segment:
                continue

            has_content = True

            if segment == '+':
                operator = '+'
            elif segment == '*':
                operator = '*'
            else:
                # Attempt to parse as numbers
                try:
                    numbers.append(int(segment))
                except ValueError:
                    pass # Should not happen based on problem desc if it happens idk what to do lol!

        if not has_content:
            continue

        # 4. Calculate result for this block
        if operator == '+':
            result = sum(numbers)
        elif operator == '*':
            result = 1
            for n in numbers:
                result *= n
        else:
            # Fallback or error if no operator found (it shouldn't happen)
            result = 0

        grand_total += result

    print("-" * 30)
    print(f"Grand Total: {grand_total}")

if __name__ == "__main__":
    solve_trash_compactor('input.txt')