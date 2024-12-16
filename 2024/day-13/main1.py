def find_min_cost_to_win_prizes(claw_machines):
    max_presses = 100  # Constraint: at most 100 presses of each button
    total_cost = 0
    prizes_won = 0

    for machine in claw_machines:
        x_A, y_A = machine['A']  # Increments for button A
        x_B, y_B = machine['B']  # Increments for button B
        x_prize, y_prize = machine['prize']  # Prize position

        min_cost = float('inf')
        found_solution = False

        # Brute force over possible values of n_A and n_B
        for n_A in range(max_presses + 1):
            for n_B in range(max_presses + 1):
                # Check if the claw aligns exactly with the prize
                if (n_A * x_A + n_B * x_B == x_prize and
                    n_A * y_A + n_B * y_B == y_prize):
                    found_solution = True
                    cost = 3 * n_A + 1 * n_B  # Calculate cost for this combination
                    min_cost = min(min_cost, cost)  # Track minimum cost

        if found_solution:
            prizes_won += 1
            total_cost += min_cost

    return prizes_won, total_cost

# Read input from a file
import json

def read_input_file(file_path):
    with open(file_path, 'r') as file:
        claw_machines = []
        machine_data = []

        for line_number, line in enumerate(file, start=1):
            line = line.strip()

            if not line:  # Skip blank lines
                continue

            machine_data.append(line)

            if len(machine_data) == 3:  # Expecting exactly 3 lines per machine
                try:
                    # Parse Button A, Button B, and Prize lines
                    a_values = machine_data[0].split(': ')[1].split(', ')
                    b_values = machine_data[1].split(': ')[1].split(', ')
                    prize_values = machine_data[2].split(': ')[1].split(', ')

                    machine = {
                        'A': (int(a_values[0][2:]), int(a_values[1][2:])),  # Remove "X+" and "Y+"
                        'B': (int(b_values[0][2:]), int(b_values[1][2:])),
                        'prize': (int(prize_values[0][2:]), int(prize_values[1][2:]))
                    }
                    claw_machines.append(machine)

                except (IndexError, ValueError) as e:
                    print(f"Error parsing machine data at lines {line_number - 2}-{line_number}: {e}")

                finally:
                    machine_data = []  # Reset for the next machine

        if machine_data:  # Handle leftover incomplete data
            print(f"Incomplete machine data at the end of file: {machine_data}")

    return claw_machines

# Load data from input.txt
claw_machines = read_input_file('input.txt')

# Solve the problem
prizes_won, total_cost = find_min_cost_to_win_prizes(claw_machines)

print(f"Prizes won: {prizes_won}")
print(f"Total cost: {total_cost}")

