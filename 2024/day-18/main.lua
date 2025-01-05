local function parse_input(file_name)
    local points = {}
    for line in io.lines(file_name) do
        local x, y = line:match("(%d+),(%d+)")
        x, y = tonumber(x), tonumber(y)
        table.insert(points, {x = x, y = y})
    end
    return points
end

local function build_grid(size)
    local grid = {}
    for i = 1, size do
        grid[i] = {}
        for j = 1, size do
            grid[i][j] = false
        end
    end
    return grid
end

local function bfs(grid, start, goal)
    local size = #grid
    local directions = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    local queue = {{x = start.x, y = start.y}}
    local visited = {}
    visited[start.y * size + start.x] = true

    while #queue > 0 do
        local current = table.remove(queue, 1)
        if current.x == goal.x and current.y == goal.y then
            return true
        end

        for _, dir in ipairs(directions) do
            local nx, ny = current.x + dir[1], current.y + dir[2]
            if nx >= 1 and nx <= size and ny >= 1 and ny <= size and
               not grid[ny][nx] and not visited[ny * size + nx] then
                table.insert(queue, {x = nx, y = ny})
                visited[ny * size + nx] = true
            end
        end
    end
    return false
end

local function find_shortest_path(grid, start, goal)
    local size = #grid
    local directions = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    local queue = {{x = start.x, y = start.y, steps = 0}}
    local visited = {}
    visited[start.y * size + start.x] = true

    while #queue > 0 do
        local current = table.remove(queue, 1)
        if current.x == goal.x and current.y == goal.y then
            return current.steps
        end

        for _, dir in ipairs(directions) do
            local nx, ny = current.x + dir[1], current.y + dir[2]
            if nx >= 1 and nx <= size and ny >= 1 and ny <= size and
               not grid[ny][nx] and not visited[ny * size + nx] then
                table.insert(queue, {x = nx, y = ny, steps = current.steps + 1})
                visited[ny * size + nx] = true
            end
        end
    end
    return -1
end

local function main()
    local file_name = "input.txt"
    local points = parse_input(file_name)
    local size = 71
    local grid = build_grid(size)

    local start = {x = 1, y = 1}
    local goal = {x = size, y = size}

    -- Part 1: Shortest path after 1024 bytes
    for i = 1, math.min(1024, #points) do
        local point = points[i]
        grid[point.y + 1][point.x + 1] = true
    end
    local shortest_path = find_shortest_path(grid, start, goal)
    if shortest_path == -1 then
        print("No valid path to the exit after 1024 bytes.")
    else
        print("Shortest Path after 1024 bytes: " .. shortest_path)
    end

    -- Reset grid for Part 2
    grid = build_grid(size)

    -- Part 2: First byte that blocks the path
    for i, point in ipairs(points) do
        grid[point.y + 1][point.x + 1] = true
        if not bfs(grid, start, goal) then
            print(string.format("First blocking byte: %d,%d", point.x, point.y))
            break
        end
    end
end

main()
