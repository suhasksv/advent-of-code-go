from collections import deque

def parse_input_file(file_path):
    """Parse the input file containing byte positions."""
    with open(file_path, 'r') as file:
        return [tuple(map(int, line.strip().split(','))) for line in file]

def build_grid(size=71):
    """Build an initially empty grid."""
    return [[False] * size for _ in range(size)]

def bfs_path_exists(grid):
    """Check if there's a path from (0, 0) to (70, 70) using BFS."""
    n = len(grid)
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]  # Down, Right, Up, Left
    queue = deque([(0, 0)])  # (x, y)
    visited = set()
    visited.add((0, 0))

    while queue:
        x, y = queue.popleft()
        
        # If we've reached the bottom-right corner, a path exists
        if (x, y) == (n - 1, n - 1):
            return True
        
        # Explore neighbors
        for dx, dy in directions:
            nx, ny = x + dx, y + dy
            if 0 <= nx < n and 0 <= ny < n and not grid[ny][nx] and (nx, ny) not in visited:
                visited.add((nx, ny))
                queue.append((nx, ny))
    
    # If no path is found
    return False

def find_blocking_byte(file_path, size=71):
    """Find the first byte that blocks the path from start to exit."""
    corrupted_positions = parse_input_file(file_path)
    grid = build_grid(size=size)
    
    for x, y in corrupted_positions:
        # Corrupt the grid at this position
        grid[y][x] = True
        
        # Check if the path is still reachable
        if not bfs_path_exists(grid):
            return f"{x},{y}"  # Return the coordinates as required

    return "Path is never fully blocked."

# Main Execution
if __name__ == "__main__":
    # Path to the input file
    input_file = "input.txt"
    
    # Find the first byte that blocks the path
    blocking_byte = find_blocking_byte(input_file, size=71)
    print(blocking_byte)

