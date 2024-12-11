import numpy as np

# Directions for exploring neighbors (cardinal and diagonal directions)
directions = [(-1, 0), (1, 0), (0, -1), (0, 1), (-1, -1), (-1, 1), (1, -1), (1, 1)]

def is_local_maximum(x, y, grid):
    """Check if the cell (x, y) is a local maximum."""
    value = grid[x, y]
    for dx, dy in directions:
        nx, ny = x + dx, y + dy
        if 0 <= nx < grid.shape[0] and 0 <= ny < grid.shape[1]:
            if grid[nx, ny] >= value:
                return False
    return True

def count_distinct_trails(x, y, grid):
    """Count all distinct trails from a trailhead using DFS with path uniqueness ensured."""
    stack = [(x, y, [(x, y)])]  # Stack holds (current position, path so far)
    visited_paths = set()

    while stack:
        cx, cy, path = stack.pop()
        visited_paths.add(tuple(path))

        for dx, dy in directions:
            nx, ny = cx + dx, cy + dy
            if 0 <= nx < grid.shape[0] and 0 <= ny < grid.shape[1] and grid[nx, ny] < grid[cx, cy]:
                if (nx, ny) not in path:  # Avoid revisiting nodes in the same path
                    stack.append((nx, ny, path + [(nx, ny)]))
    
    return len(visited_paths)

def calculate_total_rating(grid):
    """Calculate the sum of ratings for all trailheads in the grid."""
    total_rating = 0
    for i in range(grid.shape[0]):
        for j in range(grid.shape[1]):
            if is_local_maximum(i, j, grid):
                rating = count_distinct_trails(i, j, grid)
                print(f"Trailhead at ({i}, {j}) has rating {rating}.")
                total_rating += rating
    return total_rating

# Example grid
topo_map = np.array([
    [8, 9, 0, 1, 0, 1, 2, 3],
    [7, 8, 1, 2, 1, 8, 7, 4],
    [8, 7, 4, 3, 0, 9, 6, 5],
    [9, 6, 5, 4, 9, 8, 7, 4],
    [4, 5, 6, 7, 8, 9, 0, 3],
    [3, 2, 0, 1, 9, 0, 1, 2],
    [0, 1, 3, 2, 9, 8, 0, 1],
    [1, 0, 4, 5, 6, 7, 3, 2]
])

# Calculate and print total trailhead ratings
total_rating = calculate_total_rating(topo_map)
print(f"Sum of all trailhead ratings: {total_rating}")

