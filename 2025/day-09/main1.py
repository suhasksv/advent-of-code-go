import re

def solve_movie_theater(input_str):
    lines = input_str.strip().split('\n')
    points = []

    # Robust parsing using regex to find integer pairs
    for line in lines:
        if not line.strip():
            continue
        # Find all numbers (including negative ones) in the line
        nums = re.findall(r'-?\d+', line)
        if len(nums) >= 2:
            x, y = int(nums[0]), int(nums[1])
            points.append((x, y))

    if len(points) < 2:
        print(f"Warning: Only found {len(points)} valid points. Check input format.")
        return 0

    print(f"Successfully parsed {len(points)} red tiles.")

    max_area = 0

    # Iterate through all unique pairs
    n = len(points)
    for i in range(n):
        for j in range(i + 1, n):
            p1 = points[i]
            p2 = points[j]

            # Calculate width and height (Inclusive: counting the tiles)
            width = abs(p1[0] - p2[0]) + 1
            height = abs(p1[1] - p2[1]) + 1

            # Calculate area
            area = width * height

            if area > max_area:
                max_area = area

    return max_area

# Example Input
example_input = """
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
"""

if __name__ == "__main__":
    print("--- Example ---")
    ex_result = solve_movie_theater(example_input)
    print(f"Example Result: {ex_result} (Expected: 50)")

    try:
        with open('input.txt', 'r') as f:
            real_input = f.read()
            print("-" * 30)
            print(f"Real Input Result: {solve_movie_theater(real_input)}")
    except FileNotFoundError:
        print("input.txt not found. Run with local file for real answer.")