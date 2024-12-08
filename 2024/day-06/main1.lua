function parse_input(filename)
    local grid = {}
    for line in io.lines(filename) do
        table.insert(grid, {})
        for char in line:gmatch(".") do
            table.insert(grid[#grid], char)
        end
    end
    return grid
end

function find_guard(grid)
    local directions = { ["^"] = 1, [">"] = 2, ["v"] = 3, ["<"] = 4 }
    for r = 1, #grid do
        for c = 1, #grid[r] do
            if directions[grid[r][c]] then
                return r, c, directions[grid[r][c]]
            end
        end
    end
    error("Guard not found")
end

function simulate_patrol(grid)
    local directions = {
        { -1, 0 }, -- Up
        { 0, 1 },  -- Right
        { 1, 0 },  -- Down
        { 0, -1 }  -- Left
    }

    local r, c, dir = find_guard(grid)
    local rows, cols = #grid, #grid[1]
    local visited = {}

    local function mark_visited(r, c)
        visited[r * cols + c] = true
    end

    local function is_visited(r, c)
        return visited[r * cols + c] ~= nil
    end

    while r > 0 and r <= rows and c > 0 and c <= cols do
        mark_visited(r, c)

        local dr, dc = directions[dir][1], directions[dir][2]
        local nr, nc = r + dr, c + dc

        if nr > 0 and nr <= rows and nc > 0 and nc <= cols and grid[nr][nc] == "#" then
            dir = (dir % 4) + 1 -- Turn right
        else
            r, c = nr, nc -- Move forward
        end
    end

    local count = 0
    for _ in pairs(visited) do
        count = count + 1
    end

    return count
end

local input_file = "input.txt"
local grid = parse_input(input_file)
print("Distinct positions visited:", simulate_patrol(grid))
