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

local function is_game_possible(game_data, red_limit, green_limit, blue_limit)
    for _, subset in ipairs(game_data) do
        if subset.red > red_limit or subset.green > green_limit or subset.blue > blue_limit then
            return false
        end
    end
    return true
end

local function find_possible_games(file_path, red_limit, green_limit, blue_limit)
    local total_id_sum = 0

    local file = io.open(file_path, "r")
    if not file then
        error("Could not open file: " .. file_path)
    end

    for line in file:lines() do
        if line ~= "" then
            local game_id, game_data = parse_game_line(line)
            if is_game_possible(game_data, red_limit, green_limit, blue_limit) then
                total_id_sum = total_id_sum + game_id
            end
        end
    end

    file:close()
    return total_id_sum
end

-- Main
local file_path = "input.txt"
local red_limit, green_limit, blue_limit = 12, 13, 14
local total_id_sum = find_possible_games(file_path, red_limit, green_limit, blue_limit)
print("Sum of IDs of possible games: " .. total_id_sum)
