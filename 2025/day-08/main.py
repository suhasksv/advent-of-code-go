import math
from collections import defaultdict

class UnionFind:
    def __init__(self, n):
        # Initialize with n elements, 0 to n-1
        self.parent = list(range(n))
        self.num_components = n

    def find(self, i):
        if self.parent[i] != i:
            # Path compression
            self.parent[i] = self.find(self.parent[i])
        return self.parent[i]

    def union(self, i, j):
        root_i = self.find(i)
        root_j = self.find(j)

        if root_i != root_j:
            self.parent[root_i] = root_j
            self.num_components -= 1
            return True # Successfully merged
        return False # Already in same set

def parse_input(input_str):
    points = []
    lines = input_str.strip().split('\n')
    for line in lines:
        if not line.strip():
            continue
        parts = line.split(',')
        if len(parts) == 3:
            x, y, z = map(int, parts)
            points.append((x, y, z))
    return points

def dist_sq(p1, p2):
    return (p1[0] - p2[0])**2 + (p1[1] - p2[1])**2 + (p1[2] - p2[2])**2

def get_sorted_edges(points):
    edges = []
    n = len(points)
    for i in range(n):
        for j in range(i + 1, n):
            d = dist_sq(points[i], points[j])
            edges.append((d, i, j))
    edges.sort(key=lambda x: x[0])
    return edges

def solve_part1(input_str, num_connections=1000):
    points = parse_input(input_str)
    if not points: return 0
    n = len(points)

    edges = get_sorted_edges(points)
    uf = UnionFind(n)

    limit = min(len(edges), num_connections)

    for k in range(limit):
        _, u, v = edges[k]
        uf.union(u, v)

    circuit_counts = defaultdict(int)
    for i in range(n):
        root = uf.find(i)
        circuit_counts[root] += 1

    sizes = sorted(circuit_counts.values(), reverse=True)

    result = 1
    count = 0
    for s in sizes:
        result *= s
        count += 1
        if count == 3: break

    return result

def solve_part2(input_str):
    points = parse_input(input_str)
    if not points: return 0
    n = len(points)

    edges = get_sorted_edges(points)
    uf = UnionFind(n)

    last_u, last_v = -1, -1

    # Iterate through all edges until we have 1 component
    for d, u, v in edges:
        if uf.union(u, v):
            # If this union reduced the component count to 1, we are done
            if uf.num_components == 1:
                last_u, last_v = u, v
                break

    if last_u != -1:
        # Return product of X coordinates
        return points[last_u][0] * points[last_v][0]
    return 0

# Example Input
example_input = """
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
"""

if __name__ == "__main__":
    print("--- Example ---")
    print(f"Part 1: {solve_part1(example_input, 10)}")
    print(f"Part 2: {solve_part2(example_input)} (Expected: 25272)")

    try:
        with open('input.txt', 'r') as f:
            real_input = f.read()
            print("\n--- Real Input ---")
            print(f"Part 1: {solve_part1(real_input, 1000)}")
            print(f"Part 2: {solve_part2(real_input)}")
    except FileNotFoundError:
        print("\nNote: 'input.txt' not found.")