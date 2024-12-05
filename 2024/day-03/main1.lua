-- Function to read the content of a file
function read_file(file)
    local f = io.open(file, "r")
    if not f then
        return nil, "File not found."
    end
    local content = f:read("*all")
    f:close()
    return content
end

-- Function to extract mul(x, y) patterns and calculate the sum of their products
function calculate_mul_sum(data)
    local part1 = 0
    local pattern = "mul%((%d+),(%d+)%)"

    -- Iterate through all matches of mul(x, y)
    for x, y in string.gmatch(data, pattern) do
        local x_val = tonumber(x)
        local y_val = tonumber(y)
        if x_val and y_val then
            part1 = part1 + (x_val * y_val)
            print("mul(" .. x_val .. "," .. y_val .. ")")
        end
    end
    return part1
end

-- Main program
local input_file = arg[1] or "input.txt" -- Use the argument or default to input.txt
local data, err = read_file(input_file)

if not data then
    print("Error reading file: " .. err)
else
    local part1 = calculate_mul_sum(data)
    print("Sum of all products: " .. part1)
end
