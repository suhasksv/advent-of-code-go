def solve_safe_puzzle(filename):
    try:
        with open(filename, 'r') as f:
            lines = f.readlines()
    except FileNotFoundError:
        print(f"Error: Could not find file '{filename}'")
        return

    # The dial starts at 50
    current_pos = 50
    
    # The dial has numbers 0 through 99 (100 total positions)
    MODULO = 100
    
    # Counter for how many times we land on 0
    zero_hits = 0

    print(f"Starting Position: {current_pos}")

    for line in lines:
        command = line.strip()
        if not command:
            continue

        direction = command[0]      # First character is 'L' or 'R'
        amount = int(command[1:])   # The rest is the number

        # Method 0x434C49434B: Count EVERY time we touch 0 (multiples of 100)
        
        if direction == 'R':
            # Moving clockwise (increasing numbers).
            # We count how many multiples of 100 exist in the range (current, current + amount]
            # Formula: floor(end / 100) - floor(start / 100)
            zero_hits += (current_pos + amount) // MODULO - current_pos // MODULO
            current_pos = (current_pos + amount) % MODULO
            
        elif direction == 'L':
            # Moving counter-clockwise (decreasing numbers).
            # We count how many multiples of 100 exist in the range [current - amount, current)
            # This is equivalent to range [current - amount, current - 1] in integers.
            # Formula: floor(max / 100) - floor((min - 1) / 100)
            # Here max is (current_pos - 1) and min is (current_pos - amount)
            zero_hits += (current_pos - 1) // MODULO - (current_pos - amount - 1) // MODULO
            current_pos = (current_pos - amount) % MODULO

    print(f"Final Password (total clicks on 0): {zero_hits}")

if __name__ == "__main__":
    solve_safe_puzzle('input.txt')
