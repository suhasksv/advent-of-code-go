from collections import defaultdict

def parse_input(input_str):
    lines = input_str.strip().split('\n')
    if not lines:
        return None, 0, 0, 0

    height = len(lines)
    width = len(lines[0])

    start_col = -1
    for r in range(height):
        for c in range(width):
            if lines[r][c] == 'S':
                # Return grid starting from the 'S' row to simplify
                return lines[r:], len(lines[r:]), width, c

    return lines, height, width, 0

def solve_part1(input_str):
    lines, height, width, start_col = parse_input(input_str)

    active_beams = {start_col}
    total_splits = 0

    for r in range(height):
        next_beams = set()
        for c in list(active_beams):
            if c < 0 or c >= width: continue

            cell = lines[r][c]

            if cell == '^':
                total_splits += 1
                next_beams.add(c - 1)
                next_beams.add(c + 1)
            else:
                next_beams.add(c)

        active_beams = next_beams
        if not active_beams: break

    return total_splits

def solve_part2(input_str):
    lines, height, width, start_col = parse_input(input_str)

    # Key: Column Index, Value: Count of timelines
    active_timelines = defaultdict(int)
    active_timelines[start_col] = 1

    total_completed_timelines = 0

    for r in range(height):
        next_timelines = defaultdict(int)

        for c, count in active_timelines.items():
            if c < 0 or c >= width:
                total_completed_timelines += count
                continue

            cell = lines[r][c]

            if cell == '^':
                # Split: 1 timeline becomes 2 (Left and Right)
                next_timelines[c - 1] += count
                next_timelines[c + 1] += count
            else:
                # Pass through
                next_timelines[c] += count

        active_timelines = next_timelines
        if not active_timelines: break

    total_completed_timelines += sum(active_timelines.values())
    return total_completed_timelines

if __name__ == "__main__":
    try:
        with open('input.txt', 'r') as f:
            data = f.read()
            print(f"Part 1 Splits: {solve_part1(data)}")
            print(f"Part 2 Timelines: {solve_part2(data)}")
    except FileNotFoundError:
        print("input.txt not found.")