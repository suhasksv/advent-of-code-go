def solve_gift_shop(filename):
    try:
        with open(filename, 'r') as f:
            raw_data = f.read().strip()
    except FileNotFoundError:
        print(f"Error: Could not find file '{filename}'")
        return

    # 1. Parse the ranges
    # "11-22,95-115" -> [(11, 22), (95, 115)]
    ranges = []
    max_val = 0
    
    parts = raw_data.split(',')
    for part in parts:
        start_str, end_str = part.split('-')
        start, end = int(start_str), int(end_str)
        ranges.append((start, end))
        if end > max_val:
            max_val = end

    print(f"Processing {len(ranges)} ranges. Maximum ID to check: {max_val}")

    # 2. Generate "Invalid" IDs (Candidate Generation)
    # Instead of checking every number in the ranges (which could be billions),
    # we generate numbers that fit the pattern A + A (e.g., 123123) and check 
    # if they fall into our ranges.
    
    candidates = []
    
    # How many digits does max_val have?
    # If max_val is 1000 (4 digits), we need to generate patterns up to length 4.
    max_digits = len(str(max_val))
    
    # We only care about even-length numbers for the pattern A+A
    # If max digits is 10, we generate half-patterns of length 1 to 5.
    max_half_len = max_digits // 2
    
    for half_len in range(1, max_half_len + 1):
        # Generate all numbers of length `half_len`
        # e.g., if half_len is 2, start=10, end=99
        start_num = 10**(half_len - 1)
        end_num = (10**half_len) - 1
        
        for i in range(start_num, end_num + 1):
            # Create the repeated string pattern
            s = str(i)
            candidate_str = s + s
            candidate_val = int(candidate_str)
            
            # Optimization: Stop generating if we exceed the global max (optional)
            if candidate_val > max_val:
                break
                
            candidates.append(candidate_val)

    print(f"Generated {len(candidates)} candidate 'invalid' IDs.")

    # 3. Sum invalid IDs that appear in the ranges
    total_sum = 0
    found_count = 0
    
    # Sort candidates for cleaner checking (optional but good for debugging)
    candidates.sort()
    
    for candidate in candidates:
        # Check if this candidate falls into ANY of the requested ranges
        for (start, end) in ranges:
            if start <= candidate <= end:
                total_sum += candidate
                found_count += 1
                break 

    print(f"Found {found_count} invalid IDs in the specified ranges.")
    print(f"Total Sum: {total_sum}")

if __name__ == "__main__":
    solve_gift_shop('input.txt')
