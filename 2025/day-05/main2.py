def solve_cafeteria(filename):
    try:
        with open(filename, 'r') as f:
            content = f.read().strip()
    except FileNotFoundError:
        print(f"Error: Could not find file '{filename}'")
        return

    parts = content.split('\n\n')
    
    if len(parts) < 2:
        print("Error: Input format invalid.")
        return

    range_lines = parts[0].strip().split('\n')
    id_lines = parts[1].strip().split('\n')

    # --- PARSING ---
    ranges = []
    for line in range_lines:
        start_str, end_str = line.strip().split('-')
        ranges.append((int(start_str), int(end_str)))

    ids_to_check = []
    for line in id_lines:
        if line.strip():
            ids_to_check.append(int(line.strip()))

    print(f"Loaded {len(ranges)} ranges and {len(ids_to_check)} IDs to check.")

    # --- PART 1: Check Specific IDs ---
    fresh_count_p1 = 0
    for ingredient_id in ids_to_check:
        is_fresh = False
        for (start, end) in ranges:
            if start <= ingredient_id <= end:
                is_fresh = True
                break
        if is_fresh:
            fresh_count_p1 += 1

    print(f"Part 1 - Specific Fresh Ingredients: {fresh_count_p1}")

    # --- PART 2: Count Total Unique Integers in Ranges ---
    # We need to count how many unique numbers exist in the union of all ranges.
    # Algorithm: Merge Intervals.
    
    # 1. Sort ranges by their start value
    sorted_ranges = sorted(ranges, key=lambda x: x[0])
    
    merged_ranges = []
    
    for current_start, current_end in sorted_ranges:
        if not merged_ranges:
            merged_ranges.append([current_start, current_end])
        else:
            # Get the last range we added
            last_start, last_end = merged_ranges[-1]
            
            # Check for overlap.
            # Since the list is sorted by start time, we only need to check
            # if the current_start is <= last_end (plus 1 if we wanted to merge touching ranges,
            # but for counting sums, strict overlap check is sufficient logic).
            
            # NOTE: Logic specific to inclusive integers:
            # Range 10-14 and 14-20 overlap at 14.
            # Range 10-14 and 15-20 do NOT overlap, but are contiguous.
            # If we strictly strictly want to resolve double-counting:
            if current_start <= last_end: 
                # There is an overlap (or they touch if you use +1 logic, but let's stick to overlap)
                # We extend the previous range to include this one
                merged_ranges[-1][1] = max(last_end, current_end)
            else:
                # No overlap, start a new range
                merged_ranges.append([current_start, current_end])

    # 2. Sum the lengths of the merged ranges
    total_fresh_ids = 0
    for start, end in merged_ranges:
        # Range 3-5 contains 3, 4, 5 (Count is 3)
        # Length = End - Start + 1
        total_fresh_ids += (end - start + 1)

    print(f"Part 2 - Total Unique Fresh IDs: {total_fresh_ids}")

if __name__ == "__main__":
    solve_cafeteria('input.txt')
