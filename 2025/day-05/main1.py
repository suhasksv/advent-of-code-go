def solve_cafeteria(filename):
    try:
        with open(filename, 'r') as f:
            content = f.read().strip()
    except FileNotFoundError:
        print(f"Error: Could not find file '{filename}'")
        return

    # Split the content into two sections based on the double newline
    # Windows might use \r\n, so we normalize or rely on split
    parts = content.split('\n\n')
    
    if len(parts) < 2:
        print("Error: Input format invalid. Expected ranges and IDs separated by a blank line.")
        return

    range_lines = parts[0].strip().split('\n')
    id_lines = parts[1].strip().split('\n')

    # 1. Parse the Fresh Ranges
    ranges = []
    for line in range_lines:
        start_str, end_str = line.strip().split('-')
        ranges.append((int(start_str), int(end_str)))

    # 2. Parse the Available IDs
    ids_to_check = []
    for line in id_lines:
        if line.strip():
            ids_to_check.append(int(line.strip()))

    print(f"Loaded {len(ranges)} ranges and {len(ids_to_check)} IDs to check.")

    # 3. Check Freshness
    fresh_count = 0

    for ingredient_id in ids_to_check:
        is_fresh = False
        for (start, end) in ranges:
            # Check if the ID falls within the inclusive range
            if start <= ingredient_id <= end:
                is_fresh = True
                break # It's fresh, no need to check other ranges for this ID
        
        if is_fresh:
            fresh_count += 1
            # print(f"ID {ingredient_id} is fresh.")
        else:
            # print(f"ID {ingredient_id} is spoiled.")
            pass

    print("-" * 30)
    print(f"Total Fresh Ingredients: {fresh_count}")

if __name__ == "__main__":
    solve_cafeteria('input.txt')
