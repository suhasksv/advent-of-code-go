local function solve(x, t, dp)
    local key = tostring(x) .. "," .. tostring(t)
    if dp[key] then
        return dp[key]
    end

    local result
    if t == 0 then
        result = 1
    elseif x == 0 then
        result = solve(1, t - 1, dp)
    elseif #tostring(x) % 2 == 0 then
        local dstr = tostring(x)
        local mid = #dstr // 2
        local left = tonumber(dstr:sub(1, mid))
        local right = tonumber(dstr:sub(mid + 1))
        result = solve(left, t - 1, dp) + solve(right, t - 1, dp)
    else
        result = solve(x * 2024, t - 1, dp)
    end

    dp[key] = result
    return result
end

local function solve_all(d, t)
    local dp = {}
    local sum = 0
    for _, x in ipairs(d) do
        sum = sum + solve(x, t, dp)
    end
    return sum
end

local function main()
    -- Read input from input.txt
    local file = io.open("input.txt", "r")
    if not file then
        print("Error reading file")
        return
    end

    local data = file:read("*a"):gsub("\n", "")
    file:close()

    local d = {}
    for num in string.gmatch(data, "%S+") do
        table.insert(d, tonumber(num))
    end

    -- Calculate results for 25 and 75 blinks
    local result25 = solve_all(d, 25)
    local result75 = solve_all(d, 75)

    print("Result after 25 blinks: " .. result25)
    print("Result after 75 blinks: " .. result75)
end

main()