import sys
import os

# Increase recursion depth for flood fill
sys.setrecursionlimit(200000)

def solve_part_two_optimized(filename):
    if not os.path.exists(filename):
        print(f"Error: {filename} not found.")
        return

    # 1. Parse Input
    points = []
    with open(filename, 'r') as f:
        for line in f:
            line = line.strip()
            if ',' in line:
                x, y = map(int, line.split(','))
                points.append((x, y))

    if not points:
        print("No data found.")
        return

    # 2. Coordinate Compression Setup
    # Get all unique X and Y coordinates sorted
    unique_x = sorted(list(set(x for x, y in points)))
    unique_y = sorted(list(set(y for x, y in points)))

    # Create a mapping from real coordinate -> compressed index
    # We leave gaps (indices) between numbers to represent the empty space between rows/cols
    # Mapping: Real Coord -> Grid Index
    map_x = {}
    map_y = {}

    # Reverse mapping to check gap sizes later if needed (not strictly needed for logic, but good for debugging)
    # The grid width/height will be roughly 2 * unique coordinates

    current_idx = 1 # Start at 1 to leave a padding border
    for i, val in enumerate(unique_x):
        map_x[val] = current_idx
        # If there is a gap > 1 between this val and the next, assume a gap column
        if i < len(unique_x) - 1 and unique_x[i+1] > val + 1:
            current_idx += 2
        else:
            current_idx += 1

    W = current_idx + 2 # +2 for padding at end

    current_idx = 1
    for i, val in enumerate(unique_y):
        map_y[val] = current_idx
        if i < len(unique_y) - 1 and unique_y[i+1] > val + 1:
            current_idx += 2
        else:
            current_idx += 1

    H = current_idx + 2

    # 3. Create Compressed Grid
    # 0 = Empty/Unknown, 1 = Boundary
    grid = [[0] * W for _ in range(H)]

    num_pts = len(points)
    for i in range(num_pts):
        p1 = points[i]
        p2 = points[(i + 1) % num_pts]

        # Get compressed coordinates
        cx1, cy1 = map_x[p1[0]], map_y[p1[1]]
        cx2, cy2 = map_x[p2[0]], map_y[p2[1]]

        # Draw lines on compressed grid
        # We fill every cell between the two points inclusive
        if cx1 == cx2: # Vertical
            for y in range(min(cy1, cy2), max(cy1, cy2) + 1):
                grid[y][cx1] = 1
        else: # Horizontal
            for x in range(min(cx1, cx2), max(cx1, cx2) + 1):
                grid[cy1][x] = 1

    # 4. Flood Fill "Outside"
    # Mark outside cells as 2
    queue = [(0, 0)]
    grid[0][0] = 2

    while queue:
        cx, cy = queue.pop(0)
        for dx, dy in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            nx, ny = cx + dx, cy + dy
            if 0 <= nx < W and 0 <= ny < H:
                if grid[ny][nx] == 0:
                    grid[ny][nx] = 2
                    queue.append((nx, ny))

    # 5. Build "Bad Sector" Prefix Sum
    # We want to know if a rectangle contains ANY '2's (Outside).
    # If it contains a 2, it's invalid.
    bad_grid = [[0] * W for _ in range(H)]
    for y in range(H):
        for x in range(W):
            if grid[y][x] == 2:
                bad_grid[y][x] = 1

    # 2D Prefix Sum for O(1) lookups
    prefix_bad = [[0] * W for _ in range(H)]
    for y in range(H):
        for x in range(W):
            val = bad_grid[y][x]
            top = prefix_bad[y-1][x] if y > 0 else 0
            left = prefix_bad[y][x-1] if x > 0 else 0
            diag = prefix_bad[y-1][x-1] if (y > 0 and x > 0) else 0
            prefix_bad[y][x] = val + top + left - diag

    def count_bad_in_rect(x1, y1, x2, y2):
        # Query the prefix sum table
        return prefix_bad[y2][x2] - prefix_bad[y1-1][x2] - prefix_bad[y2][x1-1] + prefix_bad[y1-1][x1-1]

    # 6. Check Every Pair
    max_area = 0

    print(f"Compressed grid size: {W}x{H} (Safe from memory crash)")

    for i in range(len(points)):
        for j in range(i + 1, len(points)):
            p1 = points[i]
            p2 = points[j]

            # Real Area Calculation
            width = abs(p1[0] - p2[0]) + 1
            height = abs(p1[1] - p2[1]) + 1
            area = width * height

            # Optimization: Don't check validity if area is already smaller than max found
            if area <= max_area:
                continue

            # Validity Check on Compressed Grid
            cx1, cx2 = sorted([map_x[p1[0]], map_x[p2[0]]])
            cy1, cy2 = sorted([map_y[p1[1]], map_y[p2[1]]])

            if count_bad_in_rect(cx1, cy1, cx2, cy2) == 0:
                max_area = area

    print(f"The largest valid area is: {max_area}")

solve_part_two_optimized('input.txt')