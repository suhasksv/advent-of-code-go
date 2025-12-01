def can_construct(design, towels):
    n = len(design)
    dp = [False] * (n + 1)
    dp[0] = True  # An empty design is always possible
    
    # Dynamic programming to check if design can be made
    for i in range(1, n + 1):
        for towel in towels:
            t_len = len(towel)
            if i >= t_len and dp[i - t_len] and design[i - t_len:i] == towel:
                dp[i] = True
                break
    return dp[n]

def count_possible_designs(towels, designs):
    possible_count = 0
    for design in designs:
        if can_construct(design, towels):
            possible_count += 1
    return possible_count

def read_input(file_path):
    with open(file_path, 'r') as file:
        # Read the input from the file
        data = file.read().strip().split('\n\n')
        
        # The first part is the towel patterns (comma separated)
        towels = data[0].split(', ')
        
        # The second part is the desired designs (one per line)
        designs = data[1].splitlines()
        
        return towels, designs

# Read the input from 'input.txt'
file_path = 'input.txt'
towels, designs = read_input(file_path)

# Count possible designs
result = count_possible_designs(towels, designs)
print(result)  # Output the number of possible designs
