from collections import deque

def parse_map(grid):
    rows = len(grid)
    cols = len(grid[0])
    visited = [[False] * cols for _ in range(rows)]

    def in_bounds(x, y):
        return 0 <= x < rows and 0 <= y < cols

    def bfs(start_x, start_y, plant_type):
        queue = deque([(start_x, start_y)])
        visited[start_x][start_y] = True
        area = 0
        sides = 0

        while queue:
            x, y = queue.popleft()
            area += 1

            for dx, dy in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                nx, ny = x + dx, y + dy
                if not in_bounds(nx, ny) or grid[nx][ny] != plant_type:
                    sides += 1
                elif not visited[nx][ny]:
                    visited[nx][ny] = True
                    queue.append((nx, ny))

        return area, sides

    regions = []
    for i in range(rows):
        for j in range(cols):
            if not visited[i][j]:
                region_type = grid[i][j]
                area, sides = bfs(i, j, region_type)
                regions.append((region_type, area, sides))

    return regions

def calculate_total_price(grid):
    regions = parse_map(grid)
    total_price = sum(area * sides for _, area, sides in regions)
    return total_price

# Read input from file
with open("input.txt", "r") as file:
    grid = [line.strip() for line in file]

total_price = calculate_total_price(grid)
print("Total Price:", total_price)

