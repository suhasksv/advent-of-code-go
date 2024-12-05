-- Set the file to read input from
local file = arg[1] or 'input.txt'
local part2 = 0
local enabled = true

-- Function to read the file content
local function readFile(filename)
    local file = io.open(filename, "r")
    if not file then
        print("Error reading file:", filename)
        return nil
    end
    local content = file:read("*a")
    file:close()
    return content
end

-- Read the input data from the file
local D = readFile(file)
if not D then return end

-- Loop through the string
for i = 1, #D do
    -- Check for "do()"
    if D:sub(i, i+3) == "do()" then
        enabled = true
    end

    -- Check for "don't()"
    if D:sub(i, i+6) == "don't()" then
        enabled = false
    end

    -- Look for occurrences of "mul("
    if D:sub(i, i+3) == "mul(" then
        -- Find the closing parenthesis
        local j = i + 3
        while j <= #D and D:sub(j, j) ~= ")" do
            j = j + 1
        end

        if j <= #D and D:sub(j, j) == ")" then
            -- Extract numbers inside "mul(x, y)"
            local mul_str = D:sub(i+4, j-1)
            local x, y = mul_str:match("(%d+),%s*(%d+)")

            if x and y then
                x, y = tonumber(x), tonumber(y)

                -- Make sure the closing parenthesis isn't followed by a number
                if j + 1 <= #D and not D:sub(j+1, j+1):match("%d") then
                    -- Perform the multiplication if enabled
                    if enabled then
                        part2 = part2 + x * y
                    end
                    print(D:sub(i, j))  -- Print the found "mul(x, y)"
                end
            end
        end
    end
end

-- Output the final result
print(part2)
