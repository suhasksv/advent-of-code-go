def parse_game_record(record):
    game_id, game_data = record.split(': ')
    game_id = int(game_id.split()[1])
    subsets = game_data.split('; ')
    return game_id, subsets


def is_game_possible(subsets, max_red, max_green, max_blue):
    for subset in subsets:
        counts = {'red': 0, 'green': 0, 'blue': 0}
        cubes = subset.split(', ')
        for cube in cubes:
            num, color = cube.split()
            counts[color] += int(num)
        if counts['red'] > max_red or counts['green'] > max_green or counts['blue'] > max_blue:
            return False
    return True


def main():
    with open('input.txt', 'r') as file:
        game_records = file.readlines()

    max_red = 12
    max_green = 13
    max_blue = 14
    sum_of_possible_game_ids = 0

    for record in game_records:
        game_id, subsets = parse_game_record(record.strip())
        if is_game_possible(subsets, max_red, max_green, max_blue):
            sum_of_possible_game_ids += game_id

    print(sum_of_possible_game_ids)


if __name__ == "__main__":
    main()
