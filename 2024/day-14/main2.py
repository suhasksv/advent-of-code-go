
import re

def read_robots(filename):
    robots = []
    with open(filename, 'r') as f:
        for line in f:
            match = re.match(r'p=([0-9]+),([0-9]+) v=([0-9]+),([0-9]+)', line)
            if match:
                x, y, vx, vy = map(int, match.groups())
                robots.append({'x': x, 'y': y, 'vx': vx, 'vy': vy})
    return robots

def simulate_robots(robots, seconds):
    for _ in range(seconds):
        for robot in robots:
            robot['x'] = (robot['x'] + robot['vx']) % 101
            robot['y'] = (robot['y'] + robot['vy']) % 103

def count_quadrants(robots):
    quadrants = [0, 0, 0, 0]
    for robot in robots:
        if robot['x'] < 50 and robot['y'] < 51:
            quadrants[0] += 1
        elif robot['x'] >= 50 and robot['y'] < 51:
            quadrants[1] += 1
        elif robot['x'] < 50 and robot['y'] >= 51:
            quadrants[2] += 1
        elif robot['x'] >= 50 and robot['y'] >= 51:
            quadrants[3] += 1
    return quadrants

def visualize_robots(robots):
    grid = [['.' for _ in range(101)] for _ in range(103)]
    for robot in robots:
        grid[robot['y']][robot['x']] = '#'
    for row in grid:
        print(''.join(row))

def check_christmas_tree(robots):
    tree_pattern = [
        [0, 0, 1, 0, 0],
        [0, 1, 1, 1, 0],
        [1, 1, 1, 1, 1],
        [0, 1, 1, 1, 0],
        [0, 0, 1, 0, 0],
    ]

    # Find the center of the grid
    center_x = 50
    center_y = 51

    # Check if the pattern exists around the center
    for y_offset, row in enumerate(tree_pattern):
        for x_offset, pattern_value in enumerate(row):
            if pattern_value == 1:
                x = center_x + x_offset - 2  # Adjust for pattern offset
                y = center_y + y_offset - 2
                if not (0 <= x < 101 and 0 <= y < 103):
                    return False  # Pattern extends outside the grid
                robot_found = any((robot['x'] == x and robot['y'] == y) for robot in robots)
                if not robot_found:
                    return False  # Robot not found where expected

    return True

robots = read_robots('input.txt')

# Part 1
simulate_robots(robots, 100)
quadrant_counts = count_quadrants(robots)
safety_factor = quadrant_counts[0] * quadrant_counts[1] * quadrant_counts[2] * quadrant_counts[3]
print(f"Safety Factor: {safety_factor}")

# Part 2
seconds = 0
while True:
    simulate_robots(robots, 1)
    seconds += 1
    if check_christmas_tree(robots):
        print(f"Christmas Tree found after {seconds} seconds!")
        break
