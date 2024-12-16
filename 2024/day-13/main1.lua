

-- Define a Machine class with A, B, and Prize as tables
Machine = {}
Machine.__index = Machine

function Machine:new(a, b, prize)
    local self = setmetatable({}, Machine)
    self.A = a
    self.B = b
    self.Prize = prize
    return self
end

-- Function to find the minimum cost to win prizes
function findMinCostToWinPrizes(clawMachines)
    local maxPresses = 100
    local totalCost = 0
    local prizesWon = 0

    for _, machine in ipairs(clawMachines) do
        local xA, yA = machine.A[1], machine.A[2]
        local xB, yB = machine.B[1], machine.B[2]
        local xPrize, yPrize = machine.Prize[1], machine.Prize[2]

        local minCost = math.huge -- Set to maximum value
        local foundSolution = false

        for nA = 0, maxPresses do
            for nB = 0, maxPresses do
                if nA * xA + nB * xB == xPrize and nA * yA + nB * yB == yPrize then
                    foundSolution = true
                    local cost = 3 * nA + 1 * nB
                    if cost < minCost then
                        minCost = cost
                    end
                end
            end
        end

        if foundSolution then
            prizesWon = prizesWon + 1
            totalCost = totalCost + minCost
        end
    end

    return prizesWon, totalCost
end


-- Function to read the input file
function readInputFile(filePath)
    local file, err = io.open(filePath, "r")
    if not file then
        return nil, err
    end

    local clawMachines = {}
    local machineData = {}
    local lineNumber = 0

    for line in file:lines() do
        lineNumber = lineNumber + 1
        line = line:match("^%s*(.-)%s*$")  -- Trim whitespace

        if line == "" then
            goto continue
        end

        table.insert(machineData, line)

        if #machineData == 3 then  -- Expecting exactly 3 lines per machine
            local machine, err = parseMachineData(machineData, lineNumber)
            if err then
                print(string.format("Error parsing machine data at lines %d-%d: %s", lineNumber-2, lineNumber, err))
            else
                table.insert(clawMachines, machine)
            end
            machineData = {}
        end

        ::continue::
    end

    file:close()

    if #machineData > 0 then
        print("Incomplete machine data at the end of file: " .. table.concat(machineData, ", "))
    end

    return clawMachines, nil
end

-- Function to parse machine data from input lines
function parseMachineData(machineData, lineNumber)
    local function parseValues(input)
        local parts = {}
        for part in input:gmatch("([^,]+)") do
            table.insert(parts, part)
        end

        if #parts ~= 2 then
            return nil, string.format("invalid format: %s", input)
        end

        local x = tonumber(parts[1]:sub(3))  -- Skip "X+"
        local y = tonumber(parts[2]:sub(3))  -- Skip "Y+"
        if not x or not y then
            return nil, string.format("failed to parse values: %s, %s", parts[1], parts[2])
        end

        return {x, y}, nil
    end

    local aValues, err = parseValues(machineData[1]:match(": (.+)"))
    if err then return nil, err end

    local bValues, err = parseValues(machineData[2]:match(": (.+)"))
    if err then return nil, err end

    local prizeValues, err = parseValues(machineData[3]:match(": (.+)"))
    if err then return nil, err end

    return Machine:new(aValues, bValues, prizeValues), nil
end

-- Main function
function main()
    local filePath = "input.txt"
    local clawMachines, err = readInputFile(filePath)
    if err then
        print("Error reading input file: " .. err)
        return
    end

    local prizesWon, totalCost = findMinCostToWinPrizes(clawMachines)
    print("Prizes won: " .. prizesWon)
    print("Total cost: " .. totalCost)
end

-- Run the main function
main()

