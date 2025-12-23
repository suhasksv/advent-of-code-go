import math
from collections import defaultdict

class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.size = [1] * n  # Track size of each component

    def find(self, i):
        if self.parent[i] != i:
            self.parent[i] = self.find(self.parent[i])  # Path compression
        return self.parent[i]

    def union(self, i, j):
        root_i = self.find(i)
        root_j = self.find(j)

        if root_i != root_j:
            # Union by rank/size could be used, but simple assignment works here
            self.parent[root_i] = root_j
            self.size[root_j] += self.size[root_i]
            return True
        return False

def parse_input(input_str):
    points = []
    lines = input_str.strip().split('\n')
    for line in lines:
        if not line.strip(): continue
        parts = line.split(',')
        if len(parts) == 3:
            points.append(tuple(map(int, parts)))
    return points

def dist_sq(p1, p2):
    # Squared distance avoids expensive square root operations
    return (p1[0] - p2[0])**2 + (p1[1] - p2[1])**2 + (p1[2] - p2[2])**2

def solve():
    try:
        with open('input.txt', 'r') as f:
            data = f.read()
    except FileNotFoundError:
        print("Error: 'input.txt' not found.")
        return

    points = parse_input(data)
    n = len(points)

    # Generate all possible edges
    edges = []
    for i in range(n):
        for j in range(i + 1, n):
            d = dist_sq(points[i], points[j])
            edges.append((d, i, j))

    # Sort edges by distance (ascending)
    edges.sort(key=lambda x: x[0])

    uf = UnionFind(n)

    # Connect the top 1000 closest pairs
    limit = min(len(edges), 1000)
    for k in range(limit):
        _, u, v = edges[k]
        uf.union(u, v)

    # Count sizes of all distinct circuits
    # We can rely on the UnionFind's size array, but we must check roots
    component_sizes = []
    for i in range(n):
        if uf.parent[i] == i: # Only count roots
            component_sizes.append(uf.size[i])

    # Sort sizes descending
    component_sizes.sort(reverse=True)

    # Multiply the 3 largest
    result = 1
    for i in range(min(3, len(component_sizes))):
        result *= component_sizes[i]

    print(f"Part 1 : {result}")

if __name__ == "__main__":
    solve()