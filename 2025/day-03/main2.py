def solve_lobby_batteries(filename):
    try:
        with open(filename, 'r') as f:
            lines = f.readlines()
    except FileNotFoundError:
        print(f"Error: Could not find file '{filename}'")
        return

    # PART 2: We need to find the largest number formed by exactly 12 digits.
    REQUIRED_LENGTH = 12
    total_joltage = 0

    print(f"Processing {len(lines)} battery banks. Target Length: {REQUIRED_LENGTH}")

    for line in lines:
        bank = line.strip()
        if not bank:
            continue
        
        if len(bank) < REQUIRED_LENGTH:
            print(f"Skipping bank {bank} (too short)")
            continue

        # Problem: Find the subsequence of length K that forms the largest integer.
        # Since order matters and we want the largest lexicographical value:
        # We use a Monotonic Stack (Greedy approach).
        # We iterate through digits and pop smaller previous digits from the stack
        # IF we have enough remaining characters to still reach length K.
        
        stack = []
        # Number of characters we are allowed to delete to get down to K
        # If input is length 15 and we need 12, we can delete 3 characters.
        drop_budget = len(bank) - REQUIRED_LENGTH
        
        for digit in bank:
            # While stack is not empty
            # AND current digit is bigger than the last one we picked
            # AND we still have budget to drop digits
            while drop_budget > 0 and stack and stack[-1] < digit:
                stack.pop()
                drop_budget -= 1
            stack.append(digit)
            
        # If we finished the loop and still have budget left (e.g., descending sequence 987...),
        # we just truncate the end of the stack to get exactly K digits.
        final_digits = stack[:REQUIRED_LENGTH]
        
        best_bank_score = int("".join(final_digits))
        
        # print(f"Bank {bank[:15]}... -> Best: {best_bank_score}")
        total_joltage += best_bank_score

    print("-" * 30)
    print(f"Total Output Joltage: {total_joltage}")

if __name__ == "__main__":
    solve_lobby_batteries('input.txt')
