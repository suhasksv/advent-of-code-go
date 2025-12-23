def solve_printing_dept(filename):
    try:
        with open(filename, 'r') as f:
            lines = [line.strip() for line in f.readlines() if line.strip()]
    except FileNotFoundError:
        print(f"Error: Could not find file '{filename}'")
        return

    # Convert to a 2D grid (list of lists) so we can modify it
    grid = [list(line) for line in lines]
    rows = len(grid)
    cols = len(grid[0])

    # Directions for 8 neighbors: (row_change, col_change)
    neighbor_offsets = [
        (-1, -1), (-1, 0), (-1, 1),
        (0, -1),           (0, 1),
        (1, -1),  (1, 0),  (1, 1)
    ]

    print(f"Analyzing {rows}x{cols} grid for recursive removal...")

    total_removed = 0
    iteration = 0

    while True:
        iteration += 1
        rolls_to_remove = []

        # Step 1: Scan the grid to find all accessible rolls in this current state
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == '@':
                    neighbor_paper_count = 0

                    # Check all 8 neighbors
                    for dr, dc in neighbor_offsets:
                        nr, nc = r + dr, c + dc

                        # Boundary check
                        if 0 <= nr < rows and 0 <= nc < cols:
                            if grid[nr][nc] == '@':
                                neighbor_paper_count += 1

                    # The Rule: Forklifts can access if FEWER than 4 paper neighbors
                    if neighbor_paper_count < 4:
                        rolls_to_remove.append((r, c))

        # Step 2: If nothing to remove, we are done
        if not rolls_to_remove:
            break

        # Step 3: Remove the rolls (update grid state for next iteration)
        count_this_round = len(rolls_to_remove)
        total_removed += count_this_round

        # print(f"Iteration {iteration}: Removing {count_this_round} rolls.")

        for r, c in rolls_to_remove:
            grid[r][c] = '.' # Mark as empty space

    print("-" * 30)
    print(f"Total rolls of paper removed: {total_removed}")

if __name__ == "__main__":
    solve_printing_dept('input.txt')