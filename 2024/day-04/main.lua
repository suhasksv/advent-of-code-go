-- Function to count patterns in the grid
local function count_patterns(grid)
    local p1, p2 = 0, 0
    local R = #grid
    local C = #grid[1]

    -- Loop through the grid
    for r = 1, R do
        for c = 1, C do
            -- Check for "XMAS" horizontally, vertically, and diagonally
            if c + 3 <= C and grid[r]:sub(c, c) == "X" and grid[r]:sub(c + 1, c + 1) == "M" and grid[r]:sub(c + 2, c + 2) == "A" and grid[r]:sub(c + 3, c + 3) == "S" then
                p1 = p1 + 1
            end
            if r + 3 <= R and grid[r]:sub(c, c) == "X" and grid[r + 1]:sub(c, c) == "M" and grid[r + 2]:sub(c, c) == "A" and grid[r + 3]:sub(c, c) == "S" then
                p1 = p1 + 1
            end
            if r + 3 <= R and c + 3 <= C and grid[r]:sub(c, c) == "X" and grid[r + 1]:sub(c + 1, c + 1) == "M" and grid[r + 2]:sub(c + 2, c + 2) == "A" and grid[r + 3]:sub(c + 3, c + 3) == "S" then
                p1 = p1 + 1
            end

            -- Check for "SAMX" horizontally, vertically, and diagonally
            if c + 3 <= C and grid[r]:sub(c, c) == "S" and grid[r]:sub(c + 1, c + 1) == "A" and grid[r]:sub(c + 2, c + 2) == "M" and grid[r]:sub(c + 3, c + 3) == "X" then
                p1 = p1 + 1
            end
            if r + 3 <= R and grid[r]:sub(c, c) == "S" and grid[r + 1]:sub(c, c) == "A" and grid[r + 2]:sub(c, c) == "M" and grid[r + 3]:sub(c, c) == "X" then
                p1 = p1 + 1
            end
            if r + 3 <= R and c + 3 <= C and grid[r]:sub(c, c) == "S" and grid[r + 1]:sub(c + 1, c + 1) == "A" and grid[r + 2]:sub(c + 2, c + 2) == "M" and grid[r + 3]:sub(c + 3, c + 3) == "X" then
                p1 = p1 + 1
            end
            if r - 3 >= 1 and c + 3 <= C and grid[r]:sub(c, c) == "S" and grid[r - 1]:sub(c + 1, c + 1) == "A" and grid[r - 2]:sub(c + 2, c + 2) == "M" and grid[r - 3]:sub(c + 3, c + 3) == "X" then
                p1 = p1 + 1
            end
            if r - 3 >= 1 and c + 3 <= C and grid[r]:sub(c, c) == "X" and grid[r - 1]:sub(c + 1, c + 1) == "M" and grid[r - 2]:sub(c + 2, c + 2) == "A" and grid[r - 3]:sub(c + 3, c + 3) == "S" then
                p1 = p1 + 1
            end

            -- Check for "MAS" patterns surrounded by "M/S"
            if r + 2 <= R and c + 2 <= C and grid[r]:sub(c, c) == "M" and grid[r + 1]:sub(c + 1, c + 1) == "A" and grid[r + 2]:sub(c + 2, c + 2) == "S" and grid[r + 2]:sub(c, c) == "M" and grid[r]:sub(c + 2, c + 2) == "S" then
                p2 = p2 + 1
            end
            if r + 2 <= R and c + 2 <= C and grid[r]:sub(c, c) == "M" and grid[r + 1]:sub(c + 1, c + 1) == "A" and grid[r + 2]:sub(c + 2, c + 2) == "S" and grid[r + 2]:sub(c, c) == "S" and grid[r]:sub(c + 2, c + 2) == "M" then
                p2 = p2 + 1
            end
            if r + 2 <= R and c + 2 <= C and grid[r]:sub(c, c) == "S" and grid[r + 1]:sub(c + 1, c + 1) == "A" and grid[r + 2]:sub(c + 2, c + 2) == "M" and grid[r + 2]:sub(c, c) == "M" and grid[r]:sub(c + 2, c + 2) == "S" then
                p2 = p2 + 1
            end
            if r + 2 <= R and c + 2 <= C and grid[r]:sub(c, c) == "S" and grid[r + 1]:sub(c + 1, c + 1) == "A" and grid[r + 2]:sub(c + 2, c + 2) == "M" and grid[r + 2]:sub(c, c) == "S" and grid[r]:sub(c + 2, c + 2) == "M" then
                p2 = p2 + 1
            end
        end
    end

    return p1, p2
end

-- Read the grid from a file
local infile = arg[1] or "4.in"
local grid = {}
for line in io.lines(infile) do
    table.insert(grid, line)
end

-- Count the patterns
local p1, p2 = count_patterns(grid)
print(p1)
print(p2)
