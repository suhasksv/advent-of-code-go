def split_even_digits(num):
    """Splits a number with an even number of digits into two parts."""
    num_str = str(num)
    mid = len(num_str) // 2
    return int(num_str[:mid]), int(num_str[mid:])

def simulate_blinks(stones, blinks):
    """Simulates the stone transformations for the given number of blinks."""
    for _ in range(blinks):
        new_stones = []
        for stone in stones:
            if stone == 0:
                new_stones.append(1)
            elif len(str(stone)) % 2 == 0:
                left, right = split_even_digits(stone)
                new_stones.extend([left, right])
            else:
                new_stones.append(stone * 2024)
        stones = new_stones
    return stones

# Read initial arrangement of stones from input.txt
with open("input.txt", "r") as file:
    initial_stones = list(map(int, file.readline().strip().split()))

# Number of blinks to simulate
num_blinks = 75

# Simulate and count the number of stones after 25 blinks
final_stones = simulate_blinks(initial_stones, num_blinks)
print(f"Number of stones after {num_blinks} blinks: {len(final_stones)}")

