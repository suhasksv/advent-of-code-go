local function calculateCalibrationValue(line)
    local firstDigit = string.match(line, "%d")
    local lastDigit = string.match(string.reverse(line), "%d") -- Reverse for last digit

    if firstDigit and lastDigit then
        return tonumber(firstDigit .. lastDigit)
    end

    return 0
end

local function calculateTotal(filename)
    local file = io.open(filename, "r")
    local total = 0
    for line in file:lines() do
        total = total + calculateCalibrationValue(line)
    end
    file:close()
    return total
end

local filename = "input.txt"
local result = calculateTotal(filename)
print("Part 1: The sum of all calibration values is:", result)
