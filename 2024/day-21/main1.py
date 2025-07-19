from collections import deque

# Define the numeric keypad layout and button positions
numeric_keypad = [
    ['7', '8', '9'],
    ['4', '5', '6'],
    ['1', '2', '3'],
    [' ', '0', 'A']
]

# Directions for moving the robot arm (up, down, left, right)
DIRECTIONS = [(-1, 0), (1, 0), (0, -1), (0, 1)]  # up, down, left, right

# Map each button to its (row, col) position
button_positions = {}
for r in range(4):
    for c in range(3):
        if numeric_keypad[r][c] != ' ':
            button_positions[numeric_keypad[r][c]] = (r, c)

# Function to find the shortest path (number of steps) from start to target
def bfs(start, target):
    queue = deque([(start[0], start[1], 0)])  # (row, col, steps)
    visited = set()
    visited.add(start)
    
    while queue:
        r, c, steps = queue.popleft()
        
        # If we reached the target, return the number of steps
        if (r, c) == target:
            return steps
        
        for dr, dc in DIRECTIONS:
            nr, nc = r + dr, c + dc
            
            # Check if the new position is valid and not visited
            if 0 <= nr < 4 and 0 <= nc < 3 and numeric_keypad[nr][nc] != ' ' and (nr, nc) not in visited:
                visited.add((nr, nc))
                queue.append((nr, nc, steps + 1))
    
    return -1  # If no valid path exists

# Function to calculate the complexity for a single code
def calculate_complexity(code):
    # Start at 'A'
    start = button_positions['A']
    total_steps = 0
    
    # Iterate over each digit in the code
    for char in code:
        if char == 'A':
            continue  # Skip 'A' because it's the starting point
        target = button_positions[char]
        steps = bfs(start, target)
        total_steps += steps
        start = target  # Update start position for the next button press
    
    # Numeric value of the code (ignoring the 'A')
    numeric_value = int(code[:-1])  # Remove 'A' and convert to integer
    complexity = total_steps * numeric_value
    return complexity

# Main function to calculate total complexity for all codes
def main(codes):
    total_complexity = 0
    for code in codes:
        total_complexity += calculate_complexity(code)
    return total_complexity

# Example input codes
codes = ['029A', '980A', '179A', '456A', '379A']
result = main(codes)
print(f"Total complexity: {result}")
