import sys
import re

# Increase recursion limit
sys.setrecursionlimit(10**6)

# Set the file to read input from
file = 'input.txt' if len(sys.argv) < 2 else sys.argv[1]
D = open(file).read().strip()

part1 = 0

# Loop through the string
for i in range(len(D)):
    # Look for occurrences of "mul("
    if D[i:i+4] == "mul(":
        j = i+4 # Start after "mul("
        while D[j] != ")": # Find the closing parenthesis
            j += 1

        try:
            # Extract numbers inside "mul(x, y)"
            x, y = map(int, re.findall('\d+', D[i:j+1]))
            if D[j-1] not in ['0', '1', '2','3', '4', '5', '6', '7', '8', '9']:
                continue
            part1 += x * y
            print(D[i:j+1])
        except:
            pass
print(part1)
