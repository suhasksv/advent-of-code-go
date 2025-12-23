import sys
import os
from functools import lru_cache

# Increase recursion depth to handle long chains of devices
sys.setrecursionlimit(20000)

def solve_reactor(filename):
    if not os.path.exists(filename):
        print(f"Error: {filename} not found.")
        return

    # 1. Parse the Input into a Graph
    # graph[source] = [dest1, dest2, ...]
    graph = {}

    with open(filename, 'r') as f:
        for line in f:
            line = line.strip()
            if not line:
                continue

            # Line format: "name: dest1 dest2 dest3"
            if ':' in line:
                source, destinations = line.split(':')
                source = source.strip()
                # Split destinations by whitespace
                dests = destinations.strip().split()
                graph[source] = dests

    # 2. Define DFS with Memoization
    # We use @lru_cache to automatically store results for visited nodes
    @lru_cache(maxsize=None)
    def count_paths(current_node):
        # Base Case: We reached the target
        if current_node == 'out':
            return 1

        # Dead End: This node doesn't go anywhere (and isn't 'out')
        if current_node not in graph:
            return 0

        # Recursive Step: Sum paths from all neighbors
        total_paths = 0
        for neighbor in graph[current_node]:
            total_paths += count_paths(neighbor)

        return total_paths

    # 3. Execute
    if 'you' not in graph:
        print("Error: Could not find starting node 'you' in input.")
        return

    result = count_paths('you')

    print(f"Graph loaded with {len(graph)} devices.")
    print(f"Total distinct paths from 'you' to 'out': {result}")

# Run the function
solve_reactor('input.txt')