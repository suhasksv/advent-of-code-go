use std::collections::HashMap;
use std::fs;

fn main() {
    let filename = "input.txt"; // Replace with your actual input file

    let file_contents = fs::read_to_string(filename)
        .expect("Something went wrong reading the file");

    let max_red = 12;
    let max_green = 13;
    let max_blue = 14;
    let mut possible_games_sum = 0;

    for line in file_contents.lines() {
        let (game_id, subsets) = parse_game_record(line);
        if is_game_possible(&subsets, max_red, max_green, max_blue) {
            possible_games_sum += game_id;
        }
    }
    println!("Sum of possible game IDs: {}", possible_games_sum);
}

fn parse_game_record(record: &str) -> (u32, Vec<String>) {
    let parts: Vec<&str> = record.split(':').collect();
    let game_id: u32 = parts[0].split_whitespace().nth(1).unwrap().parse().unwrap();
    let subsets: Vec<String> = parts[1]
        .split(';')
        .map(|subset| subset.trim().to_string())
        .collect();
    (game_id, subsets)
}

fn is_game_possible(subsets: &[String], max_red: u32, max_green: u32, max_blue: u32) -> bool {
    for subset in subsets {
        let mut counts: HashMap<&str, u32> = HashMap::from([("red", 0), ("green", 0), ("blue", 0)]);
        for cube in subset.split(',') {
            let parts: Vec<&str> = cube.trim().split_whitespace().collect();
            if !parts.is_empty() { // Skip empty strings
                let num = parts[0].parse::<u32>().unwrap_or(0); // Default to 0 if not a number
                if let Some(color) = parts.get(1) {
                    *counts.get_mut(color).unwrap_or(&mut 0) += num;
                } // Ignore invalid colors
            }
        }
        if counts["red"] > max_red || counts["green"] > max_green || counts["blue"] > max_blue {
            return false;
        }
    }
    true
}

