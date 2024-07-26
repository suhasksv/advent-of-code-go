def parse_game_data(file_path):
    games = {}
    with open(file_path, 'r') as file:
        for line in file:
            line = line.strip()
            if not line:
                continue
            game_id, subsets_str = line.split(": ")
            game_id = int(game_id.split()[1])
            subsets = subsets_str.split("; ")
            game_data = []
            for subset in subsets:
                cubes = subset.split(", ")
                cube_counts = {'red': 0, 'green': 0, 'blue': 0}
                for cube in cubes:
                    count, color = cube.split()
                    count = int(count)
                    cube_counts[color] += count
                game_data.append(cube_counts)
            games[game_id] = game_data
    return games

def calculate_minimum_cubes(game_data):
    min_red = min_green = min_blue = 0
    for subset in game_data:
        min_red = max(min_red, subset['red'])
        min_green = max(min_green, subset['green'])
        min_blue = max(min_blue, subset['blue'])
    return min_red, min_green, min_blue

def calculate_power(red, green, blue):
    return red * green * blue

def find_total_power(file_path):
    games = parse_game_data(file_path)
    total_power = 0
    for game_id, game_data in games.items():
        min_red, min_green, min_blue = calculate_minimum_cubes(game_data)
        power = calculate_power(min_red, min_green, min_blue)
        total_power += power
    return total_power

# Example usage:
file_path = 'input.txt'
total_power = find_total_power(file_path)
print(f"Total power: {total_power}")
