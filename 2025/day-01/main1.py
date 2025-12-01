def solve_safe_puzzle(filename):
    try:
        with open(filename, 'r') as f:
            lines = f.readlines()
    except FileNotFoundError:
        print(f"Error: Could not find file '{filename}'")
        return

    # The dial starts at 50
    current_pos = 50
    
    # The dial has numbers 0 through 99 (total positions: 100)
    MODULO = 100
    
    # Counter for 0s
    zero_hits = 0

    print(f"Starting Position: {current_pos}")

    for line in lines:
        command = line.strip()
        if not command:
            continue

        direction = command[0]      # First character is 'L' or 'R'
        amount = int(command[1:])   # The rest is the number

        # R is addition (clockwise/higher numbers)
        # L is subtraction (counter-clockwise/lower numbers)
        if direction == 'R':
            current_pos = (current_pos + amount) % MODULO
        elif direction == 'L':
            # Python's modulo operator handles negative numbers correctly for this puzzle e.g., -5 % 100 becomes 95 which is a plus
            current_pos = (current_pos - amount) % MODULO

        # Check if we landed on 0
        if current_pos == 0:
            zero_hits += 1

    print(f"Final Password: {zero_hits}")

if __name__ == "__main__":
    solve_safe_puzzle('input.txt')
