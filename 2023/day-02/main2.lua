local function parse_game_line(line)
    local game_id, subsets_str = line:match("Game (%d+): (.+)")
    game_id = tonumber(game_id)
    local game_data = {}

    for subset in subsets_str:gmatch("[^;]+") do
        local subset_counts = {red = 0, green = 0, blue = 0}
        for count, color in subset:gmatch("(%d+) (%w+)") do
            count = tonumber(count)
            if color == "red" then
                subset_counts.red = subset_counts.red + count
            elseif color == "green" then
                subset_counts.green = subset_counts.green + count
            elseif color == "blue" then
                subset_counts.blue = subset_counts.blue + count
            end
        end
        table.insert(game_data, subset_counts)
    end

    return game_id, game_data
end

local function calculate_minimum_cubes(game_data)
    local min_red, min_green, min_blue = 0, 0, 0

    for _, subset in ipairs(game_data) do
        if subset.red > min_red then
            min_red = subset.red
        end
        if subset.green > min_green then
            min_green = subset.green
        end
        if subset.blue > min_blue then
            min_blue = subset.blue
        end
    end

    return min_red, min_green, min_blue
end

local function calculate_power(red, green, blue)
    return red * green * blue
end

local function find_total_power(file_path)
    local total_power = 0

    local file = io.open(file_path, "r")
    if not file then
        error("Could not open file: " .. file_path)
    end

    for line in file:lines() do
        if line ~= "" then
            local game_id, game_data = parse_game_line(line)
            local min_red, min_green, min_blue = calculate_minimum_cubes(game_data)
            local power = calculate_power(min_red, min_green, min_blue)
            total_power = total_power + power
        end
    end

    file:close()
    return total_power
end

-- Main
local file_path = "input.txt"
local total_power = find_total_power(file_path)
print("Sum of powers of the minimum sets of cubes: " .. total_power)
