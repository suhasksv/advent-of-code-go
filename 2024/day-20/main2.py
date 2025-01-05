from collections import deque
from itertools import product

def parse_map(map_lines):
    grid = []
    start = end = None
    for r, line in enumerate(map_lines):
        grid.append(list(line))
        for c, char in enumerate(line):
            if char == 'S':
                start = (r, c)
            elif char == 'E':
                end = (r, c)
    return grid, start, end

def multi_source_bfs(grid, start, end):
    rows, cols = len(grid), len(grid[0])
    directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    distances = {pos: float('inf') for pos in product(range(rows), range(cols)) if grid[pos[0]][pos[1]] in {'.', 'S', 'E'}}
    distances[start] = 0
    queue = deque([start])

    while queue:
        r, c = queue.popleft()
        for dr, dc in directions:
            nr, nc = r + dr, c + dc
            if (nr, nc) in distances and distances[(nr, nc)] > distances[(r, c)] + 1:
                distances[(nr, nc)] = distances[(r, c)] + 1
                queue.append((nr, nc))

    return distances

def simulate_cheats(grid, start, end, distances):
    cheats = []
    rows, cols = len(grid), len(grid[0])
    directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    track_positions = [(r, c) for r in range(rows) for c in range(cols) if grid[r][c] in {'.', 'S', 'E'}]
    normal_time = distances[end]

    print(f"Normal time from start to end: {normal_time}")
    if normal_time == float('inf'):
        print("No valid path found from start to end.")
        return cheats

    for r, c in track_positions:
        for dr, dc in directions:
            nr, nc = r + dr, c + dc
            if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == '#':
                for ddr, ddc in directions:
                    nnr, nnc = nr + ddr, nc + ddc
                    if 0 <= nnr < rows and 0 <= nnc < cols and grid[nnr][nnc] in {'.', 'S', 'E'}:
                        # Compute cheat time
                        cheat_time = distances.get((r, c), float('inf')) + 1 + distances.get((nnr, nnc), float('inf'))
                        if cheat_time < normal_time:
                            time_saved = normal_time - cheat_time
                            cheats.append(time_saved)
                            print(f"Cheat: From {(r, c)} through wall {(nr, nc)} to {(nnr, nnc)}, saves {time_saved} picoseconds")

    return cheats


def count_large_cheats(cheats, threshold):
    return sum(1 for cheat in cheats if cheat >= threshold)

# Read input from a file
with open('input.txt', 'r') as file:
    map_lines = file.read().splitlines()

grid, start, end = parse_map(map_lines)
distances = multi_source_bfs(grid, start, end)
cheats = simulate_cheats(grid, start, end, distances)
result = count_large_cheats(cheats, 100)
print(f"Number of cheats saving at least 100 picoseconds: {result}")
