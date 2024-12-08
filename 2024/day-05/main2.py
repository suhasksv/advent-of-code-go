import sys
import re
from collections import defaultdict, deque
import pyperclip as pc

def pr(s):
    """Prints and copies the result to clipboard."""
    print(s)
    pc.copy(s)

# Set recursion limit for deep recursion scenarios
sys.setrecursionlimit(10**6)

# Input file handling
infile = sys.argv[1] if len(sys.argv) >= 2 else 'input.txt'
p1 = 0  # Part 1 result
p2 = 0  # Part 2 result
D = open(infile).read().strip()

# Graph representations
E = defaultdict(set)  # Edges pointing to each node
ER = defaultdict(set) # Reverse edges for each node

# Parse input into edges and queries
edges, queries = D.split('\n\n')

# Build the graph
for line in edges.split('\n'):
    x, y = map(int, line.split('|'))
    E[y].add(x)  # y depends on x
    ER[x].add(y) # x is a prerequisite for y

# Process each query
for query in queries.split('\n'):
    vs = [int(x) for x in query.split(',')]
    assert len(vs) % 2 == 1  # Ensure the list has an odd length

    # Check if the query satisfies the dependency constraints
    ok = True
    for i, x in enumerate(vs):
        for j, y in enumerate(vs):
            if i < j and y in E[x]:  # If y depends on x but appears later
                ok = False

    if ok:
        # If valid, add the middle node to p1
        p1 += vs[len(vs) // 2]
    else:
        # Perform a topological sort for Part 2
        good = []  # Topologically sorted nodes
        Q = deque()  # Queue for BFS
        D = {v: len(E[v] & set(vs)) for v in vs}  # Count dependencies in query

        # Initialize the queue with nodes having no dependencies
        for v in vs:
            if D[v] == 0:
                Q.append(v)

        while Q:
            x = Q.popleft()
            good.append(x)
            for y in ER[x]:
                if y in D:
                    D[y] -= 1
                    if D[y] == 0:
                        Q.append(y)

        # Add the middle node of the sorted list to p2
        p2 += good[len(good) // 2]

# Output results
pr(p1)
pr(p2)
