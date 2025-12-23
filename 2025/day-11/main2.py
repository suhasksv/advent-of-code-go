import sys
import os
from functools import lru_cache

sys.setrecursionlimit(20000)

def solve_part_two(filename):
    if not os.path.exists(filename):
        print(f"Error: {filename} not found.")
        return

    # 1. Parse Input
    graph = {}
    with open(filename, 'r') as f:
        for line in f:
            line = line.strip()
            if ':' in line:
                source, dests = line.split(':')
                source = source.strip()
                graph[source] = dests.strip().split()

    # 2. Helper function to count paths between any two specific nodes
    # We use a cache dictionary manually because we need to clear it
    # or use distinct keys for different 'target' nodes if we used lru_cache.
    # Here, a simple recursive function with a specific target is easiest.

    memo = {}

    def count_paths_between(u, target):
        # State key for memoization
        state = (u, target)
        if state in memo:
            return memo[state]

        # Base Case: Reached the specific target
        if u == target:
            return 1

        # Dead End
        if u not in graph:
            return 0

        total = 0
        for v in graph[u]:
            total += count_paths_between(v, target)

        memo[state] = total
        return total

    # 3. Calculate Segments
    # Scenario A: svr -> dac -> fft -> out
    # If any segment is 0 (unreachable), the product becomes 0.

    # We clear memo between major runs just to be safe/clean,
    # though technically (u, target) tuples make it unique anyway.

    # Path A: svr -> dac
    paths_svr_dac = count_paths_between('svr', 'dac')

    # Path B: dac -> fft
    paths_dac_fft = count_paths_between('dac', 'fft')

    # Path C: fft -> out
    paths_fft_out = count_paths_between('fft', 'out')

    total_scenario_A = paths_svr_dac * paths_dac_fft * paths_fft_out

    # Scenario B: svr -> fft -> dac -> out

    # Path D: svr -> fft
    paths_svr_fft = count_paths_between('svr', 'fft')

    # Path E: fft -> dac
    paths_fft_dac = count_paths_between('fft', 'dac')

    # Path F: dac -> out
    paths_dac_out = count_paths_between('dac', 'out')

    total_scenario_B = paths_svr_fft * paths_fft_dac * paths_dac_out

    # 4. Final Result
    total_valid_paths = total_scenario_A + total_scenario_B

    print(f"Scenario 1 (dac->fft): {paths_svr_dac} * {paths_dac_fft} * {paths_fft_out} = {total_scenario_A}")
    print(f"Scenario 2 (fft->dac): {paths_svr_fft} * {paths_fft_dac} * {paths_dac_out} = {total_scenario_B}")
    print(f"Total paths visiting both: {total_valid_paths}")

solve_part_two('input.txt')