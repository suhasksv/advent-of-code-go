def count_ways_to_construct(design, towels):
    n = len(design)
    dp = [0] * (n + 1)
    dp[0] = 1  # There's one way to form an empty design
    
    # For each position in the design, calculate the number of ways to reach that position
    for i in range(1, n + 1):
        for towel in towels:
            t_len = len(towel)
            if i >= t_len and design[i - t_len:i] == towel:
                dp[i] += dp[i - t_len]
    
    return dp[n]  # Return the number of ways to construct the entire design

def total_ways_to_construct(towels, designs):
    total_ways = 0
    for design in designs:
        total_ways += count_ways_to_construct(design, towels)
    return total_ways

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

# Calculate the total number of ways to construct all designs
result = total_ways_to_construct(towels, designs)
print(result)  # Output the total number of ways
