def parse_input(filename):
    with open(filename, 'r') as file:
        grid = [list(line.strip()) for line in file.readlines()]
    return grid


def find_guard(grid):
    directions = {'^': (-1, 0), '>': (0, 1), 'v': (1, 0), '<': (0, -1)}
    for r, row in enumerate(grid):
        for c, cell in enumerate(row):
            if cell in directions:
                return r, c, cell # Starting row, column and direction
    return None


def simulate_patrol(grid):
    directions = ['^', '>', 'v', '<']
    directions_map = {'^': (-1, 0), '>': (0, 1), 'v': (1, 0), '<': (0, -1)}

    # Get initial position and direction
    r, c, facing = find_guard(grid)
    visited = set()
    rows, cols = len(grid), len(grid)


    while 0 <= r < rows and 0 <= c < cols:
        visited.add((r, c))

        # calculate next position
        dr, dc = directions_map[facing]
        nr, nc = r + dr, c + dc

        # check if there is any obstacle
        if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == '#':
            # Turn Right
            facing = directions[(directions.index(facing) + 1) % 4]
        else:
            # Move Forward
            r, c = nr, nc

    return len(visited)

if __name__ == "__main__":
    input_file = 'input.txt'
    grid = parse_input(input_file)
    result = simulate_patrol(grid)
    print(f"District position visited: {result}")