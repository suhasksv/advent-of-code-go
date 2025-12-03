def solve_lobby_batteries(filename):
    try:
        with open(filename, 'r') as f:
            lines = f.readlines()
    except FileNotFoundError:
        print(f"Error: Could not find file '{filename}'")
        return

    total_joltage = 0

    print(f"Processing {len(lines)} battery banks...")

    for line in lines:
        bank = line.strip()
        if not bank:
            continue

        # We need to find two digits at indices i and j (where i < j)
        # such that the number formed (bank[i] * 10 + bank[j]) is maximized.
        
        # Strategy:
        # 1. Iterate through possible "first digits" (tens place) from 9 down to 0.
        # 2. Find the FIRST occurrence of that digit in the string.
        #    (We pick the first occurrence because it leaves the biggest possible 
        #     suffix string to find a large second digit).
        # 3. If that digit exists and has characters after it, find the max digit 
        #    in the remaining part of the string.
        # 4. The first valid pair we find this way is guaranteed to be the maximum.

        best_bank_score = 0
        found = False

        # Check for tens digit starting from 9 down to 0
        for tens_digit in range(9, -1, -1):
            s_digit = str(tens_digit)
            first_idx = bank.find(s_digit)

            if first_idx != -1:
                # We found a candidate for the tens place.
                # Check if there are any digits following it.
                remaining_part = bank[first_idx+1:]
                
                if remaining_part:
                    # Find the largest digit in the remaining part
                    # max() on a string of digits works perfectly ('9' > '8')
                    max_units_char = max(remaining_part)
                    units_digit = int(max_units_char)
                    
                    best_bank_score = tens_digit * 10 + units_digit
                    found = True
                    break # We found the highest possible starting digit, so we are done.

        if found:
            # print(f"Bank {bank[:10]}... -> Best: {best_bank_score}")
            total_joltage += best_bank_score

    print("-" * 30)
    print(f"Total Output Joltage: {total_joltage}")

if __name__ == "__main__":
    solve_lobby_batteries('input.txt')
