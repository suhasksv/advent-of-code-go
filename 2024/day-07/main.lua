
local function parseInput(inputText)
    local equations = {}
    for line in string.gmatch(inputText, "[^
]+") do
        local testValue, numbersText = string.match(line, "^(%d+): (.+)$")
        testValue = tonumber(testValue)
        local numbers = {}
        for num in string.gmatch(numbersText, "%d+") do
            table.insert(numbers, tonumber(num))
        end
        table.insert(equations, { testValue = testValue, numbers = numbers })
    end
    return equations
end

local function evaluateLeftToRight(numbers, operators)
    local result = numbers[1]
    for i = 1, #operators do
        if operators[i] == "+" then
            result = result + numbers[i + 1]
        elseif operators[i] == "*" then
            result = result * numbers[i + 1]
        elseif operators[i] == "||" then
            result = tonumber(tostring(result) .. tostring(numbers[i + 1]))
        end
    end
    return result
end

local function generateOperatorCombinations(operators, length)
    if length == 0 then return { {} } end
    local smallerCombos = generateOperatorCombinations(operators, length - 1)
    local combinations = {}
    for _, combo in ipairs(smallerCombos) do
        for _, op in ipairs(operators) do
            local newCombo = { table.unpack(combo) }
            table.insert(newCombo, op)
            table.insert(combinations, newCombo)
        end
    end
    return combinations
end

local function isValidEquation(testValue, numbers, includeConcat)
    local operators = { "+", "*" }
    if includeConcat then
        table.insert(operators, "||")
    end
    local numOperators = #numbers - 1
    local combinations = generateOperatorCombinations(operators, numOperators)
    for _, combo in ipairs(combinations) do
        if evaluateLeftToRight(numbers, combo) == testValue then
            return true
        end
    end
    return false
end

local function calculateTotalCalibration(inputText, includeConcat)
    local equations = parseInput(inputText)
    local total = 0
    for _, eq in ipairs(equations) do
        if isValidEquation(eq.testValue, eq.numbers, includeConcat) then
            total = total + eq.testValue
        end
    end
    return total
end

-- Main
local inputFile = io.open("input.txt", "r")
local inputText = inputFile:read("*a")
inputFile:close()

local part1Result = calculateTotalCalibration(inputText, false)
print("Part 1 Total Calibration:", part1Result)

local part2Result = calculateTotalCalibration(inputText, true)
print("Part 2 Total Calibration:", part2Result)
