from collections import defaultdict

def solve_quantum_manifold():
    try:
        with open('input.txt', 'r') as f:
            grid = [line.strip('\n') for line in f if line.strip()]
    except FileNotFoundError:
        print("Error: 'input.txt' not found.")
        return

    rows = len(grid)
    cols = len(grid[0])

    # 1. Find the starting position 'S'
    start_r, start_c = -1, -1
    for r in range(rows):
        if 'S' in grid[r]:
            start_r = r
            start_c = grid[r].index('S')
            break

    if start_r == -1:
        print("Error: Could not find start point 'S'.")
        return

    # 2. Track timelines
    # We map column_index -> number_of_timelines
    # We process the grid row by row.
    current_timelines = defaultdict(int)
    current_timelines[start_c] = 1

    # Loop through each row starting from S down to the bottom
    for r in range(start_r, rows - 1):
        next_timelines = defaultdict(int)

        for c, count in current_timelines.items():
            # Look at the cell the beam is about to enter (row r+1)
            next_cell = grid[r+1][c]

            if next_cell == '^':
                # SPLITTER: The timeline splits into two distinct timelines
                # One goes Left (c-1), One goes Right (c+1)

                # Check bounds before adding
                if c - 1 >= 0:
                    next_timelines[c - 1] += count
                if c + 1 < cols:
                    next_timelines[c + 1] += count
            else:
                # EMPTY SPACE: The timeline continues straight down
                # The number of timelines remains constant for this path
                next_timelines[c] += count

        current_timelines = next_timelines

    # 3. Sum total timelines
    # The total is the sum of all timeline counts reaching the bottom
    total_timelines = sum(current_timelines.values())

    print(f"Total Active Timelines: {total_timelines}")

if __name__ == "__main__":
    solve_quantum_manifold()