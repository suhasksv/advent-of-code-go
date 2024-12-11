from collections import deque

def parse_map(map_data):
    """Parse the topographic map into a 2D list of integers."""
    return [list(map(int, line.strip())) for line in map_data]

def find_trailheads(map_grid):
    """Find all positions with height 0."""
    trailheads = []
    for r, row in enumerate(map_grid):
        for c, height in enumerate(row):
            if height == 0:
                trailheads.append((r, c))
    return trailheads

def count_reachable_nines(map_grid, start):
    """Count how many unique 9s are reachable from a given trailhead."""
    rows, cols = len(map_grid), len(map_grid[0])
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]  # Right, Down, Left, Up
    visited = set()
    reachable_nines = set()
    queue = deque([start])

    while queue:
        r, c = queue.popleft()
        if (r, c) in visited:
            continue
        visited.add((r, c))

        # Check if this is a reachable '9'
        if map_grid[r][c] == 9:
            reachable_nines.add((r, c))
            continue

        # Explore neighbors with height increasing by 1
        for dr, dc in directions:
            nr, nc = r + dr, c + dc
            if 0 <= nr < rows and 0 <= nc < cols and map_grid[nr][nc] == map_grid[r][c] + 1:
                queue.append((nr, nc))

    return len(reachable_nines)

def calculate_trailhead_scores(map_data):
    """Calculate the sum of scores for all trailheads on the map."""
    map_grid = parse_map(map_data)
    trailheads = find_trailheads(map_grid)
    total_score = sum(count_reachable_nines(map_grid, trailhead) for trailhead in trailheads)
    return total_score

def parse_map_from_file(file_path):
    """Read the topographic map from a file."""
    with open(file_path, 'r') as file:
        return [line.strip() for line in file]

# Load the map data from input.txt
file_path = "input.txt"
map_data = parse_map_from_file(file_path)

# Calculate the total score
result = calculate_trailhead_scores(map_data)
print("Total Score:", result)

