def solve():
    with open("input.txt") as f:
        lines = f.readlines()
    grid = [list(line.strip()) for line in lines[:-2]]
    moves = "".join(lines[-2:]).replace("\n", "")

    rows = len(grid)
    cols = len(grid[0])

    def find_robot():
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == '@':
                    return r, c
        return -1, -1

    robot_r, robot_c = find_robot()

    for move in moves:
        dr, dc = 0, 0
        if move == '^':
            dr = -1
        elif move == 'v':
            dr = 1
        elif move == '<':
            dc = -1
        elif move == '>':
            dc = 1

        next_r, next_c = robot_r + dr, robot_c + dc

        if 0 <= next_r < rows and 0 <= next_c < cols:
            if grid[next_r][next_c] == '.':
                # ... (move robot) ...
                pass
            elif grid[next_r][next_c] == 'O':
                next_box_r, next_box_c = next_r + dr, next_c + dc

                # Check if the box's next position is within bounds:
                if 0 <= next_box_r < rows and 0 <= next_box_c < cols and grid[next_box_r][next_box_c] == '.':
                    grid[next_box_r][next_box_c] = 'O'
                    grid[next_r][next_c] = '@'
                    grid[robot_r][robot_c] = '.'
                    robot_r, robot_c = next_r, next_c

    gps_sum = 0
    for r in range(rows):
        for c in range(cols)
            if grid[r][c] == 'O':
                gps_sum += (r * 100) + c

    print(gps_sum)
solve()
