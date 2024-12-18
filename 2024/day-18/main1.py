from collections import deque

def parse_input_file(file_path):
    """Parse the input file containing byte positions."""
    with open(file_path, 'r') as file:
        return [tuple(map(int, line.strip().split(','))) for line in file]

def build_grid(corrupted_positions, size=71):
    """Build the memory grid with corrupted cells marked as True."""
    grid = [[False] * size for _ in range(size)]
    for x, y in corrupted_positions:
        grid[y][x] = True  # Mark the cell as corrupted
    return grid

def bfs_shortest_path(grid):
    """Find the shortest path from (0, 0) to (70, 70) using BFS."""
    n = len(grid)
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]  # Down, Right, Up, Left
    queue = deque([(0, 0, 0)])  # (x, y, steps)
    visited = set()
    visited.add((0, 0))

    while queue:
        x, y, steps = queue.popleft()
        
        # If we've reached the bottom-right corner, return the step count
        if (x, y) == (n - 1, n - 1):
            return steps
        
        # Explore neighbors
        for dx, dy in directions:
            nx, ny = x + dx, y + dy
            if 0 <= nx < n and 0 <= ny < n and not grid[ny][nx] and (nx, ny) not in visited:
                visited.add((nx, ny))
                queue.append((nx, ny, steps + 1))
    
    # If no path is found
    return -1

def simulate_memory_from_file(file_path, size=71, max_falls=1024):
    """Simulate memory corruption and find the shortest path using input from a file."""
    corrupted_positions = parse_input_file(file_path)[:max_falls]
    grid = build_grid(corrupted_positions, size=size)
    return bfs_shortest_path(grid)

# Main Execution
if __name__ == "__main__":
    # Path to the input file
    input_file = "input.txt"
    
    # Simulate the memory corruption and find the shortest path
    shortest_path_steps = simulate_memory_from_file(input_file, size=71, max_falls=1024)
    if shortest_path_steps == -1:
        print("No valid path to the exit.")
    else:
        print(f"Shortest Path Steps: {shortest_path_steps}")

