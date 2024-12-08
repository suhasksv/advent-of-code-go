def read_grid(file_path):
    """Reads the grid from the given file and returns it as a list of lists."""
    with open(file_path, 'r') as file:
        return [list(line.strip()) for line in file.readlines()]

def count_word_occurrences(grid, word):
    """Counts all occurrences of a word in the grid in all directions."""
    rows = len(grid)
    cols = len(grid[0]) if rows > 0 else 0
    word_len = len(word)
    count = 0

    # Directions: (row_delta, col_delta)
    directions = [
        (0, 1),   # Horizontal right
        (0, -1),  # Horizontal left
        (1, 0),   # Vertical down
        (-1, 0),  # Vertical up
        (1, 1),   # Diagonal down-right
        (1, -1),  # Diagonal down-left
        (-1, 1),  # Diagonal up-right
        (-1, -1)  # Diagonal up-left
    ]

    def is_valid(x, y):
        """Checks if a position is valid within the grid."""
        return 0 <= x < rows and 0 <= y < cols

    def check_direction(x, y, dx, dy):
        """Checks if the word exists starting at (x, y) in the given direction."""
        for i in range(word_len):
            nx, ny = x + i * dx, y + i * dy
            if not is_valid(nx, ny) or grid[nx][ny] != word[i]:
                return False
        return True

    # Search for the word in all directions from each grid position
    for r in range(rows):
        for c in range(cols):
            for dx, dy in directions:
                if check_direction(r, c, dx, dy):
                    count += 1

    return count

if __name__ == "__main__":
    # Specify the input file name
    input_file = "input.txt"

    # Read the grid from the file
    grid = read_grid(input_file)

    # Word to search for
    target_word = "XMAS"

    # Count occurrences of the word
    total_occurrences = count_word_occurrences(grid, target_word)

    # Output the result
    print(f"The word '{target_word}' appears {total_occurrences} times in the grid.")
