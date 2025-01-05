def parse_input(input_str):
    robots = []
    for line in input_str.strip().split("\n"):
        p_part, v_part = line.split(" v=")
        p = tuple(map(int, p_part[2:].split(",")))
        v = tuple(map(int, v_part.split(",")))
        robots.append((p, v))
    return robots

def simulate_motion(robots, width, height, seconds):
    final_positions = []
    for p, v in robots:
        # Update position after the given time, applying wrap-around
        final_x = (p[0] + v[0] * seconds) % width
        final_y = (p[1] + v[1] * seconds) % height
        final_positions.append((final_x, final_y))
    return final_positions

def count_quadrants(positions, width, height):
    mid_x, mid_y = width // 2, height // 2
    quadrants = [0, 0, 0, 0]  # Top-left, top-right, bottom-left, bottom-right

    for x, y in positions:
        if x == mid_x or y == mid_y:
            continue  # Skip robots in the middle row or column
        if x < mid_x and y < mid_y:
            quadrants[0] += 1  # Top-left
        elif x >= mid_x and y < mid_y:
            quadrants[1] += 1  # Top-right
        elif x < mid_x and y >= mid_y:
            quadrants[2] += 1  # Bottom-left
        elif x >= mid_x and y >= mid_y:
            quadrants[3] += 1  # Bottom-right

    return quadrants

def calculate_safety_factor(quadrants):
    factor = 1
    for count in quadrants:
        factor *= count
    return factor

# Read the input data from a file
with open("input.txt", "r") as file:
    input_str = file.read()

robots = parse_input(input_str)
width, height = 101, 103
seconds = 100

final_positions = simulate_motion(robots, width, height, seconds)
quadrants = count_quadrants(final_positions, width, height)
safety_factor = calculate_safety_factor(quadrants)

print("Quadrants:", quadrants)
print("Safety Factor:", safety_factor)

