# from collections import deque
# import itertools
#
# def parse_map(map_lines):
#     grid = []
#     start = end = None
#     for r, line in enumerate(map_lines):
#         grid.append(list(line))
#         for c, char in enumerate(line):
#             if char == 'S':
#                 start = (r, c)
#             elif char == 'E':
#                 end = (r, c)
#     return grid, start, end
#
# def bfs(grid, start, end):
#     rows, cols = len(grid), len(grid[0])
#     directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
#     visited = set()
#     queue = deque([(start[0], start[1], 0)])  # (row, col, time)
#
#     while queue:
#         r, c, time = queue.popleft()
#         if (r, c) == end:
#             return time
#         if (r, c) in visited:
#             continue
#         visited.add((r, c))
#
#         for dr, dc in directions:
#             nr, nc = r + dr, c + dc
#             if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] != '#':
#                 queue.append((nr, nc, time + 1))
#
#     return float('inf')  # No path found
#
# def simulate_cheats(grid, start, end):
#     rows, cols = len(grid), len(grid[0])
#     track_positions = [(r, c) for r in range(rows) for c in range(cols) if grid[r][c] in {'.', 'S', 'E'}]
#
#     normal_time = bfs(grid, start, end)
#     cheats = []
#
#     for (r1, c1), (r2, c2) in itertools.combinations(track_positions, 2):
#         # Simulate a cheat
#         cheat_grid = [row[:] for row in grid]
#         dr = abs(r1 - r2)
#         dc = abs(c1 - c2)
#
#         if dr + dc <= 2:  # Cheating for at most 2 steps
#             cheat_grid[r1][c1] = '.'  # Start of cheat
#             cheat_grid[r2][c2] = '.'  # End of cheat
#             cheat_time = bfs(cheat_grid, start, end)
#             if cheat_time < normal_time:
#                 time_saved = normal_time - cheat_time
#                 cheats.append(time_saved)
#
#     return bfs(modified_grid, start, end)
#
# def count_large_cheats(grid, start, end, threshold):
#     potential_cheats = []  # Find potential cheat positions
#     # ...
#
#     count = 0
#     for start_cheat, end_cheat in potential_cheats:
#         time_saved = bfs(grid, start, end) - simulate_cheat(grid, start_cheat, end_cheat)
#         if time_saved >= threshold:
#             count += 1
#     return count
#
# # Read input from a file
# with open('input.txt', 'r') as file:
#     map_lines = file.read().splitlines()
#
# grid, start, end = parse_map(map_lines)
# result = count_large_cheats(grid, start, end, 100)
# print(f"Shortest path from start to end (no cheats): {result1}")

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

def multi_source_bfs(grid, start):
    rows, cols = len(grid), len(grid[0])
    directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    distances = {}
    queue = deque([(start, 0)])

    while queue:
        (r, c), dist = queue.popleft()
        if (r, c) in distances:
            continue
        distances[(r, c)] = dist
        for dr, dc in directions:
            nr, nc = r + dr, c + dc
            if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] in {'.', 'S', 'E'}:
                queue.append(((nr, nc), dist + 1))

    return distances

def simulate_cheats(grid, start, end, distances):
    cheats = []
    rows, cols = len(grid), len(grid[0])
    directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    normal_time = distances.get(end, float('inf'))

    print(f"Normal time from start to end: {normal_time}")
    if normal_time == float('inf'):
        print("No valid path found from start to end.")
        return cheats

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == '#':  # Only consider walls
                for dr1, dc1 in directions:
                    nr1, nc1 = r + dr1, c + dc1
                    for dr2, dc2 in directions:
                        nr2, nc2 = r + dr2, c + dc2

                        if (0 <= nr1 < rows and 0 <= nc1 < cols and grid[nr1][nc1] in {'.', 'S', 'E'} and
                                0 <= nr2 < rows and 0 <= nc2 < cols and grid[nr2][nc2] in {'.', 'S', 'E'}):
                            cheat_time = (distances.get((nr1, nc1), float('inf')) +
                                          1 +  # Bypass the wall
                                          distances.get((nr2, nc2), float('inf')))
                            if cheat_time < normal_time:
                                time_saved = normal_time - cheat_time
                                cheats.append(time_saved)
                                print(f"Cheat: From {(nr1, nc1)} through {(r, c)} to {(nr2, nc2)}, saves {time_saved} picoseconds.")

    return cheats

def count_large_cheats(cheats, threshold):
    return sum(1 for cheat in cheats if cheat >= threshold)

# Read input from a file
with open('input.txt', 'r') as file:
    map_lines = file.read().splitlines()

grid, start, end = parse_map(map_lines)
distances = multi_source_bfs(grid, start)
cheats = simulate_cheats(grid, start, end, distances)
result = count_large_cheats(cheats, 100)
print(f"Number of cheats saving at least 100 picoseconds: {result}")
