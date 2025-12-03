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

    # 2. Generate "Invalid" IDs (Generation of Candidate)
    # PART 2 RULE: An ID is invalid if it is made of a repeated sequence
    # AT LEAST twice -> e.g., 111, 121212, 123123
    
    # We are using a set because some numbers can be generated in multiple ways.
    # e.g., 111111 could be '1'*6, '11'*3, or '111'*2. We only want to count it once.
    candidates = set()
    
    # We need to determine how many digits the maximum value has.
    max_digits = len(str(max_val))
    
    # The base sequence can be anything from length 1 up to half the max digits (bcos it must repeat at least for twice).
    max_base_len = max_digits // 2 # so using floor division or integer division
    
    for base_len in range(1, max_base_len + 1):
        # Generate base numbers of specific length
        # e.g., if base_len is 1, loop 1..9
        # e.g., if base_len is 2, loop 10..99
        start_num = 10**(base_len - 1)
        end_num = (10**base_len) - 1
        
        for i in range(start_num, end_num + 1):
            base_str = str(i)
            
            repeats = 2
            while True:
                candidate_str = base_str * repeats
                
                # If the generated string is longer than max_val's string length
                # AND numerical value is larger, we can stop adding repeats for this base.
                # Converting to int can be expensive if very large, so length check helps.
                if len(candidate_str) > max_digits:
                    break
                    
                candidate_val = int(candidate_str)
                
                if candidate_val > max_val:
                    break
                
                candidates.add(candidate_val)
                repeats += 1

    print(f"Generated {len(candidates)} unique candidate 'invalid' IDs.")

    # 3. Sum invalid IDs that appear in the ranges
    total_sum = 0
    found_count = 0
    
    # Sort for processing
    sorted_candidates = sorted(list(candidates))
    
    for candidate in sorted_candidates:
        # Check if this candidate falls into ANY of the requested ranges
        for (start, end) in ranges:
            if start <= candidate <= end:
                total_sum += candidate
                found_count += 1
                # Break so we don't count the same number twice if ranges theoretically overlapped
                break 

    print(f"Found {found_count} invalid IDs in the specified ranges.")
    print(f"Total Sum: {total_sum}")

if __name__ == "__main__":
    solve_gift_shop('input.txt')
