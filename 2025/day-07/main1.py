def solve_tachyon_manifold():
    try:
        with open('input.txt', 'r') as f:
            grid = [line.strip() for line in f if line.strip()]
    except FileNotFoundError:
        print("Error: 'input.txt' not found. Please save your puzzle input to this file.")
        return

    rows = len(grid)
    cols = len(grid[0])

    # Find the starting position 'S'
    active_beams = set()
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == 'S':
                active_beams.add((r, c))
                break
        if active_beams:
            break

    if not active_beams:
        print("Error: Could not find start point 'S'.")
        return

    total_splits = 0

    # Simulate the beams row by row
    # We loop as long as we have active beams still inside the grid boundaries
    while active_beams:
        next_beams = set()

        for r, c in active_beams:
            next_r = r + 1

            # If the beam falls out of the bottom of the manifold, it's gone tata bye bye
            if next_r >= rows:
                continue

            # Check what the beam hits in the next cell
            cell_content = grid[next_r][c]

            if cell_content == '^':
                # HIT A SPLITTER
                total_splits += 1

                # Creates two new beams at the adjacent columns of the splitter
                # These new beams are conceptually at 'next_r' and will move down from there
                left_col = c - 1
                right_col = c + 1

                # Only add if inside lateral bounds
                if 0 <= left_col < cols:
                    next_beams.add((next_r, left_col))

                if 0 <= right_col < cols:
                    next_beams.add((next_r, right_col))

            else:
                # HIT EMPTY SPACE (or passes through other non-blocking chars)
                # Beam continues straight down
                next_beams.add((next_r, c))

        # Update the active set for the next iteration
        active_beams = next_beams

    print(f"Total Splits: {total_splits}")

if __name__ == "__main__":
    solve_tachyon_manifold()