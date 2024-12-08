function table.count(t)
    local count = 0
    for _ in pairs(t) do
        count = count + 1
    end
    return count
end


local function manhattan_distance(r1, c1, r2, c2)
    return math.abs(r1 - r2) + math.abs(c1 - c2)
end

local infile = arg[1] or "input.txt"
local file = io.open(infile, "r")
local grid = {}

for line in file:lines() do
    table.insert(grid, line)
end
file:close()

local R = #grid
local C = #grid[1]
local positions = {}

-- Store antenna positions
for r = 1, R do
    for c = 1, C do
        local char = grid[r]:sub(c, c)
        if char ~= "." then
            positions[char] = positions[char] or {}
            table.insert(positions[char], {r = r, c = c})
        end
    end
end

local a1 = {}
local a2 = {}

for r = 1, R do
    for c = 1, C do
        for _, pos in pairs(positions) do
            for _, p1 in ipairs(pos) do
                for _, p2 in ipairs(pos) do
                    if p1 ~= p2 then
                        local d1 = manhattan_distance(r, c, p1.r, p1.c)
                        local d2 = manhattan_distance(r, c, p2.r, p2.c)
                        local dr1, dc1 = r - p1.r, c - p1.c
                        local dr2, dc2 = r - p2.r, c - p2.c

                        if (d1 == 2 * d2 or d1 * 2 == d2) and dr1 * dc2 == dc1 * dr2 then
                            a1[r * C + c] = true
                        end
                        if dr1 * dc2 == dc1 * dr2 then
                            a2[r * C + c] = true
                        end
                    end
                end
            end
        end
    end
end

print(table.count(a1))
print(table.count(a2))
