-- Not Working
local digitMap = {
    one = 1,
    two = 2,
    three = 3,
    four = 4,
    five = 5,
    six = 6,
    seven = 7,
    eight = 8,
    nine = 9,
}

--function findCalibrationValuePart2(line)
--    local digits = {}
--
--    -- Construct a pattern string from digitMap keys
--    local digitPatterns = {}
--    for word, _ in pairs(digitMap) do
--        table.insert(digitPatterns, word)
--    end
--    local digitPatternString = table.concat(digitPatterns, "|")
--
--    -- Regular expression with lookahead to prevent overlaps
--    local digitPattern = "(?=(" .. digitPatternString .. "|%d))"
--
--    for match in string.gmatch(line, digitPattern) do
--        local valueStr = match
--        local value = tonumber(valueStr) or digitMap[valueStr]
--        table.insert(digits, { value = value, index = match:find(valueStr) })
--    end
--
--    table.sort(digits, function(a, b) return a.index < b.index end)
--
--    if #digits > 0 then
--        return digits[1].value * 10 + digits[#digits].value
--    else
--        return 0
--    end
--end

--function findCalibrationValuePart2(line)
--    local digits = {}
--
--    local digitPatterns = {}
--    for word, _ in pairs(digitMap) do
--        table.insert(digitPatterns, word)
--    end
--    local digitPatternString = table.concat(digitPatterns, "|")
--
--    local digitPattern = "(?=(" .. digitPatternString .. "|%d))"
--
--    -- Capture and iterate over the sub-matches within the lookahead group
--    for match, valueStr in string.gmatch(line, digitPattern) do
--        local value = tonumber(valueStr) or digitMap[valueStr]
--        table.insert(digits, { value = value, index = match:find(valueStr) })
--    end
--
--    table.sort(digits, function(a, b) return a.index < b.index end)
--
--    if #digits > 0 then
--        return digits[1].value * 10 + digits[#digits].value
--    else
--        return 0
--    end
--end

--function findCalibrationValuePart2(line)
--    local digits = {}
--    local pattern = "(%d|one|two|three|four|five|six|seven|eight|nine)"
--
--    for numberStr, index in line:gmatch(pattern) do
--        local value = tonumber(numberStr) or digitMap[numberStr]
--        table.insert(digits, { value = value, index = index })
--    end
--
--    table.sort(digits, function(a, b) return a.index < b.index end)
--
--    if #digits > 0 then
--        return digits[1].value * 10 + digits[#digits].value
--    else
--        return 0
--    end
--end

function findCalibrationValuePart2(line)
    local digits = {}
    local pattern = "(one|two|three|four|five|six|seven|eight|nine|%d)"

    -- Capture and iterate over the sub-matches within the lookahead group
    for match in string.gmatch(line, pattern) do
        local value = tonumber(match) or digitMap[match]
        table.insert(digits, { value = value, index = line:find(match) })
    end

    table.sort(digits, function(a, b) return a.index < b.index end)

    -- Update logic for 0 digit case
    if #digits > 0 then
        return digits[1].value * 10 + digits[#digits].value
    else
        return 0
    end
end

function calculateTotalPart2(filename)
    local file = io.open(filename, "r")
    if not file then
        error("Error reading file: " .. filename)
    end

    local total = 0
    for line in file:lines() do
        if line ~= "" then
            total = total + findCalibrationValuePart2(line)
        end
    end

    file:close()
    return total
end

local filename = "input.txt"  -- Replace with your actual input file

local total, err = calculateTotalPart2(filename)
if err then
    print(err)
else
    print("Sum of all calibration values:", total)
end
