class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.num_components = n # Track how many separate groups exist

    def find(self, i):
        if self.parent[i] != i:
            self.parent[i] = self.find(self.parent[i])
        return self.parent[i]

    def union(self, i, j):
        root_i = self.find(i)
        root_j = self.find(j)

        if root_i != root_j:
            self.parent[root_i] = root_j
            self.num_components -= 1 # Two groups became one
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

    edges = []
    for i in range(n):
        for j in range(i + 1, n):
            d = dist_sq(points[i], points[j])
            edges.append((d, i, j))

    edges.sort(key=lambda x: x[0])

    uf = UnionFind(n)

    last_u, last_v = -1, -1

    # Iterate through edges until the graph is fully connected
    for _, u, v in edges:
        if uf.union(u, v):
            # If this connection reduced the components to 1, we are done
            if uf.num_components == 1:
                last_u, last_v = u, v
                break

    if last_u != -1:
        # Calculate result: X_coordinate(u) * X_coordinate(v)
        ans = points[last_u][0] * points[last_v][0]
        print(f"Part 2 Answer: {ans}")
    else:
        print("Error: Could not connect all points.")

if __name__ == "__main__":
    solve()