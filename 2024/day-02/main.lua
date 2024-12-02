function is_safe(report)
    local all_increasing = true
    local all_decreasing = true

    for i = 1, #report - 1 do
        local diff = report[i + 1] - report[i]
        if math.abs(diff) < 1 or math.abs(diff) > 3 then
            return false
        end
        if diff <= 0 then all_increasing = false end
        if diff >= 0 then all_decreasing = false end
    end

    return all_increasing or all_decreasing
end

function can_be_made_safe(report)
    for i = 1, #report do
        local modified = {}
        for j = 1, #report do
            if i ~= j then
                table.insert(modified, report[j])
            end
        end
        if is_safe(modified) then
            return true
        end
    end
    return false
end

function count_safe_reports(file_path, part)
    local safe_count = 0
    for line in io.lines(file_path) do
        local report = {}
        for num in string.gmatch(line, "%d+") do
            table.insert(report, tonumber(num))
        end

        if part == 1 then
            if is_safe(report) then
                safe_count = safe_count + 1
            end
        elseif part == 2 then
            if is_safe(report) or can_be_made_safe(report) then
                safe_count = safe_count + 1
            end
        end
    end
    return safe_count
end

-- Main program
local file_path = "input.txt"

local part1 = count_safe_reports(file_path, 1)
print("Part 1: Number of safe reports: " .. part1)

local part2 = count_safe_reports(file_path, 2)
print("Part 2: Number of safe reports: " .. part2)
