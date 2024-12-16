local function read_input_file(filename)
    local map_input = {}
    for line in io.lines(filename) do
        local row = {}
        for char in line:gmatch(".") do
            table.insert(row, char)
        end
        table.insert(map_input, row)
    end
    return map_input
end

local function calculate_fencing_price(map_input)
    local rows = #map_input
    local cols = #map_input[1]
    local visited = {}
    for i = 1, rows do
        visited[i] = {}
        for j = 1, cols do
            visited[i][j] = false
        end
    end

    local function is_valid(x, y)
        return x >= 1 and x <= rows and y >= 1 and y <= cols
    end

    local function bfs(start_x, start_y)
        local queue = {{start_x, start_y}}
        local region_cells = {}
        local plant_type = map_input[start_x][start_y]
        visited[start_x][start_y] = true

        while #queue > 0 do
            local current = table.remove(queue, 1)
            table.insert(region_cells, current)

            for _, dir in ipairs({{-1, 0}, {1, 0}, {0, -1}, {0, 1}}) do
                local nx, ny = current[1] + dir[1], current[2] + dir[2]
                if is_valid(nx, ny) and not visited[nx][ny] and map_input[nx][ny] == plant_type then
                    visited[nx][ny] = true
                    table.insert(queue, {nx, ny})
                end
            end
        end

        return region_cells
    end

    local function calculate_area_and_perimeter(region_cells)
        local area = #region_cells
        local perimeter = 0

        for _, cell in ipairs(region_cells) do
            local x, y = cell[1], cell[2]
            for _, dir in ipairs({{-1, 0}, {1, 0}, {0, -1}, {0, 1}}) do
                local nx, ny = x + dir[1], y + dir[2]
                if not (is_valid(nx, ny) and map_input[nx][ny] == map_input[x][y]) then
                    perimeter = perimeter + 1
                end
            end
        end

        return area, perimeter
    end

    local total_price = 0

    for i = 1, rows do
        for j = 1, cols do
            if not visited[i][j] then
                local region_cells = bfs(i, j)
                local area, perimeter = calculate_area_and_perimeter(region_cells)
                total_price = total_price + area * perimeter
            end
        end
    end

    return total_price
end

local function main()
    local map_input = read_input_file("input.txt")
    local result = calculate_fencing_price(map_input)
    print("Total price of fencing: " .. result)
end

main()
