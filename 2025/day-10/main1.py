import sys
import os
import re
from collections import deque

def solve_part_one(filename):
    if not os.path.exists(filename):
        print(f"Error: {filename} not found.")
        return

    total_presses = 0
    machines_processed = 0

    with open(filename, 'r') as f:
        for line in f:
            line = line.strip()
            if not line:
                continue

            # --- 1. Parsing ---

            # Extract Target Pattern: [.##.]
            # We look for the content inside square brackets
            target_match = re.search(r'\[([.#]+)\]', line)
            if not target_match:
                continue

            target_str = target_match.group(1)
            num_lights = len(target_str)

            # Convert target string to integer bitmask
            # '.' = 0, '#' = 1
            # We'll treat index 0 as the least significant bit (2^0)
            target_mask = 0
            for i, char in enumerate(target_str):
                if char == '#':
                    target_mask |= (1 << i)

            # Extract Buttons: (0,2), (1,3)
            # We find all groups inside parentheses
            button_matches = re.findall(r'\(([\d,]+)\)', line)
            button_masks = []

            for b_str in button_matches:
                indices = [int(x) for x in b_str.split(',')]
                b_mask = 0
                for idx in indices:
                    b_mask |= (1 << idx)
                button_masks.append(b_mask)

            # --- 2. Solve using BFS ---
            # We want the shortest path from 0 (all off) to target_mask

            # Queue stores: (current_light_state, number_of_presses)
            queue = deque([(0, 0)])
            visited = {0}
            found_min_presses = -1

            while queue:
                current_state, presses = queue.popleft()

                if current_state == target_mask:
                    found_min_presses = presses
                    break

                # Try pressing each button
                for b_mask in button_masks:
                    new_state = current_state ^ b_mask # XOR to toggle

                    if new_state not in visited:
                        visited.add(new_state)
                        queue.append((new_state, presses + 1))

            if found_min_presses != -1:
                total_presses += found_min_presses
                machines_processed += 1
            else:
                print(f"Warning: Could not solve machine: {line}")

    print(f"Processed {machines_processed} machines.")
    print(f"Total fewest presses required: {total_presses}")

solve_part_one('input.txt')