local Aoc2417 = {}

-- Part 1: Execute program and calculate result
function Aoc2417.part1(input)
    assert(input, "Input cannot be nil")
    local regs, program = Aoc2417.parseInput(input)
    local a = Aoc2417.parseRegisterA(regs)
    local executionResult = Aoc2417.execProgram(a, program)

    -- Convert executionResult list to a single number
    return tonumber(table.concat(executionResult))
end

-- Part 2: Depth-first search to find the smallest valid number
function Aoc2417.part2(input)
    assert(input, "Input cannot be nil")
    local _, program = Aoc2417.parseInput(input)
    return Aoc2417.dfs(program, 0, 0, math.huge)
end

-- Depth-First Search (DFS)
function Aoc2417.dfs(program, cur, pos, success)
    for i = 0, 7 do
        local nextNum = (cur << 3) + i
        local execResult = Aoc2417.execProgram(nextNum, program)

        local programTail = {table.unpack(program, #program - pos)}
        local execTail = {table.unpack(execResult, #execResult - pos)}
        if not Aoc2417.tablesEqual(execTail, programTail) then
            goto continue
        end

        if pos == #program - 1 then
            success = math.min(success, nextNum)
            return success
        end

        success = Aoc2417.dfs(program, nextNum, pos + 1, success)

        ::continue::
    end

    return success
end

-- Execute the program logic
function Aoc2417.execProgram(a, program)
    local pointer = 1
    local output = {}
    local b, c = 0, 0

    while pointer > 0 and pointer <= #program do
        local opcode = program[pointer]
        local literalOperand = program[pointer + 1]
        local combo = (literalOperand == 4 and a) or (literalOperand == 5 and b) or (literalOperand == 6 and c) or literalOperand

        if opcode == 0 then
            a = math.floor(a / (2 ^ combo))
            pointer = pointer + 2
        elseif opcode == 1 then
            b = bit32.bxor(b, literalOperand)
            pointer = pointer + 2
        elseif opcode == 2 then
            b = combo % 8
            pointer = pointer + 2
        elseif opcode == 3 then
            pointer = (a ~= 0) and literalOperand + 1 or pointer + 2
        elseif opcode == 4 then
            b = bit32.bxor(b, c)
            pointer = pointer + 2
        elseif opcode == 5 then
            table.insert(output, combo % 8)
            pointer = pointer + 2
        elseif opcode == 6 then
            b = math.floor(a / (2 ^ combo))
            pointer = pointer + 2
        elseif opcode == 7 then
            c = math.floor(a / (2 ^ combo))
            pointer = pointer + 2
        end
    end

    return output
end

-- Parse the input into registers and program
function Aoc2417.parseInput(input)
    assert(type(input) == "string", "Input must be a string")

    local lines = {}
    for line in string.gmatch(input, "[^\n]+") do
        table.insert(lines, line)
    end

    local regs = table.concat(lines, "\n", 1, 3)
    local program = {}
    for num in string.gmatch(lines[5], "%d+") do
        table.insert(program, tonumber(num))
    end

    return regs, program
end

-- Extract register A's value
function Aoc2417.parseRegisterA(regs)
    for line in string.gmatch(regs, "[^\n]+") do
        if line:match("Register A:") then
            return tonumber(line:match("%d+"))
        end
    end
end

-- Utility: Check if two tables are equal
function Aoc2417.tablesEqual(t1, t2)
    if #t1 ~= #t2 then
        return false
    end

    for i = 1, #t1 do
        if t1[i] ~= t2[i] then
            return false
        end
    end

    return true
end

-- Example Inputs
function Aoc2417.getTestInput()
    return [[
Register A: 28066687
Register B: 0
Register C: 0

Program: 2,4,1,1,7,5,4,6,0,3,1,4,5,5,3,0
]]
end


-- Run the program
local input = Aoc2417.getTestInput()
print("Part 1: " .. Aoc2417.part1(input))
print("Part 2: " .. Aoc2417.part2(input))
