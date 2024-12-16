from collections import deque

def calculate_fencing_price(map_input):
    rows = len(map_input)
    cols = len(map_input[0])
    visited = [[False] * cols for _ in range(rows)]

    def is_valid(x, y):
        return 0 <= x < rows and 0 <= y < cols

    def bfs(x, y):
        queue = deque([(x, y)])
        region_cells = []
        visited[x][y] = True
        plant_type = map_input[x][y]

        while queue:
            cx, cy = queue.popleft()
            region_cells.append((cx, cy))

            for dx, dy in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                nx, ny = cx + dx, cy + dy
                if is_valid(nx, ny) and not visited[nx][ny] and map_input[nx][ny] == plant_type:
                    visited[nx][ny] = True
                    queue.append((nx, ny))

        return region_cells

    def calculate_area_and_perimeter(region_cells):
        area = len(region_cells)
        perimeter = 0

        for x, y in region_cells:
            for dx, dy in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                nx, ny = x + dx, y + dy
                if not (is_valid(nx, ny) and map_input[nx][ny] == map_input[x][y]):
                    perimeter += 1

        return area, perimeter

    total_price = 0

    for i in range(rows):
        for j in range(cols):
            if not visited[i][j]:
                region_cells = bfs(i, j)
                area, perimeter = calculate_area_and_perimeter(region_cells)
                total_price += area * perimeter

    return total_price

if __name__ == "__main__":
    with open("input.txt", "r") as file:
        map_input = [list(line.strip()) for line in file.readlines()]
    
    result = calculate_fencing_price(map_input)
    print(f"Total price of fencing: {result}")
