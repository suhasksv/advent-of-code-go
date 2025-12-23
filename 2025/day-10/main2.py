import sys
import os
import re
from fractions import Fraction

# Increase recursion depth just in case
sys.setrecursionlimit(5000)

def solve_part_two_optimized(filename):
    if not os.path.exists(filename):
        print(f"Error: {filename} not found.")
        return

    total_min_presses = 0
    machines_processed = 0

    with open(filename, 'r') as f:
        # Read the file line by line
        for line_num, line in enumerate(f):
            line = line.strip()
            if not line:
                continue

            # --- 1. Parsing ---

            # Extract Target Joltages: {3,5,4,7}
            target_match = re.search(r'\{([\d,]+)\}', line)
            if not target_match:
                continue
            targets = [int(x) for x in target_match.group(1).split(',')]
            num_counters = len(targets)

            # Extract Buttons: (1,3), (2), etc.
            button_matches = re.findall(r'\(([\d,]+)\)', line)
            num_buttons = len(button_matches)

            # Build Matrix A (columns = buttons) and calculate Hard Limits
            # Matrix[row][col] = 1 if button col affects counter row
            matrix = [[0] * num_buttons for _ in range(num_counters)]

            # This list will store the absolute max times each button can be pressed
            # Initialize with infinity, then clamp down
            button_limits = [float('inf')] * num_buttons

            for col_idx, b_str in enumerate(button_matches):
                indices = [int(x) for x in b_str.split(',')]

                # Check each counter this button affects
                for row_idx in indices:
                    if row_idx < num_counters:
                        matrix[row_idx][col_idx] = 1

                        # Optimization: Strict Upper Bound
                        # If button adds 1 to this counter, it can't be pressed
                        # more times than the target value for this counter.
                        limit = targets[row_idx]
                        if limit < button_limits[col_idx]:
                            button_limits[col_idx] = limit

            # Replace any remaining infinities (buttons that do nothing?) with 0
            button_limits = [0 if x == float('inf') else x for x in button_limits]

            # --- 2. Solve System ---
            # We use a recursive solver with the tight bounds we just found.

            min_presses_for_machine = float('inf')

            # We sort buttons by how "restrictive" they are (heuristic to fail fast)
            # But simpler to just recurse standardly for stability.

            def solve(btn_idx, current_counts, current_presses):
                nonlocal min_presses_for_machine

                # Pruning: If we already exceeded a known solution, stop.
                if current_presses >= min_presses_for_machine:
                    return

                # Base Case: All buttons considered
                if btn_idx == num_buttons:
                    # Check if all counters match targets exactly
                    if current_counts == targets:
                        min_presses_for_machine = current_presses
                    return

                # Recursive Step
                # We need to decide how many times to press 'btn_idx'.
                # We iterate from 0 up to its calculated LIMIT.

                # Further Optimization: Dynamic Pruning
                # Calculate the max times we can press THIS button based on *current* remainder
                local_limit = button_limits[btn_idx]

                # Check against current state of counters
                for r in range(num_counters):
                    if matrix[r][btn_idx] == 1:
                        remaining_space = targets[r] - current_counts[r]
                        if remaining_space < 0:
                            return # Already busted
                        if remaining_space < local_limit:
                            local_limit = remaining_space

                # Iterate (Try largest presses first? No, try smallest for shortest path usually)
                # Actually, iterating 0..limit is safer.

                # If the number of buttons is small (<15), this is fast.
                # If buttons are many, we might need the Gaussian approach again.
                # Let's stick to Gaussian if N > M, but "Smart Search" if N is small.
                pass

                # NOTE: If the problem has many buttons (e.g. 10+), pure recursion is still slow.
            # We will use the Gaussian Elimination method from before,
            # BUT applied to the valid subset of buttons.

            # --- 3. Robust Hybrid Solver (Gaussian + Free Var Search) ---

            # Augment matrix for Gaussian
            aug_matrix = [row[:] + [Fraction(t, 1)] for row, t in zip(matrix, targets)]
            m = [[Fraction(x, 1) for x in row] for row in aug_matrix]

            rows = num_counters
            cols = num_buttons

            pivots = [-1] * rows
            pivot_cols = [-1] * cols

            # Forward Elimination
            pivot_row = 0
            for col in range(cols):
                if pivot_row >= rows: break

                # Find pivot
                sel = pivot_row
                while sel < rows and m[sel][col] == 0:
                    sel += 1
                if sel == rows: continue

                m[pivot_row], m[sel] = m[sel], m[pivot_row]

                # Normalize
                val = m[pivot_row][col]
                for j in range(col, cols + 1):
                    m[pivot_row][j] /= val

                # Eliminate
                for i in range(rows):
                    if i != pivot_row:
                        v = m[i][col]
                        if v != 0:
                            for j in range(col, cols + 1):
                                m[i][j] -= v * m[pivot_row][j]

                pivots[pivot_row] = col
                pivot_cols[col] = pivot_row
                pivot_row += 1

            # Identify Free Variables
            free_vars = [c for c in range(cols) if pivot_cols[c] == -1]

            # Recursive search ONLY on free variables (much smaller search space)
            def solve_free(idx, current_free_vals):
                nonlocal min_presses_for_machine

                if idx == len(free_vars):
                    # Calculate Pivot Variables
                    temp_presses = sum(current_free_vals)
                    valid = True

                    for r in range(rows):
                        p_col = pivots[r]
                        if p_col == -1:
                            if m[r][cols] != 0: valid = False
                            continue

                        val = m[r][cols]
                        for f_i, f_col in enumerate(free_vars):
                            val -= m[r][f_col] * current_free_vals[f_i]

                        if val.denominator != 1 or val < 0:
                            valid = False
                            break
                        temp_presses += int(val)

                    if valid:
                        if temp_presses < min_presses_for_machine:
                            min_presses_for_machine = temp_presses
                    return

                # Optimize the bounds for this specific free variable
                # Use the HARD LIMITS we calculated earlier
                f_col = free_vars[idx]
                hard_limit = button_limits[f_col]

                # Also check mathematical bounds from the equations (optional but good)
                # x_p = Const - k * x_f  => x_f <= Const/k

                # Just use the hard limit, it's usually tight enough for AoC
                for val in range(hard_limit + 1):
                    current_free_vals.append(val)
                    # Optimization: Don't recurse if we already exceed best sum
                    if sum(current_free_vals) < min_presses_for_machine:
                        solve_free(idx + 1, current_free_vals)
                    current_free_vals.pop()

            solve_free(0, [])

            if min_presses_for_machine != float('inf'):
                total_min_presses += min_presses_for_machine
                machines_processed += 1
            else:
                # print(f"Machine {line_num}: No solution.")
                pass

    print(f"Processed {machines_processed} machines.")
    print(f"Total fewest presses: {total_min_presses}")

solve_part_two_optimized('input.txt')