use std::{
    collections::HashMap,
    fs,
    io::{self, BufRead},
};

const DIGIT_MAP: HashMap<&str, i32> = HashMap::from([
    ("one", 1), ("two", 2), ("three", 3), ("four", 4), ("five", 5),
    ("six", 6), ("seven", 7), ("eight", 8), ("nine", 9),
]);

fn find_calibration_value_part2(line: &str) -> i32 {
    let pattern = regex::Regex::new(r"(?=(\d|one|two|three|four|five|six|seven|eight|nine))").unwrap();

    let mut digits: Vec<(usize, i32)> = Vec::new(); // Store (index, value) pairs

    for cap in pattern.captures_iter(line) {
        let value_str = &cap[1];
        let value = value_str.parse::<i32>().unwrap_or_else(|_| DIGIT_MAP[value_str]);
        digits.push((cap.get(0).unwrap().start(), value));
    }

    digits.sort_by_key(|&(index, _)| index); // Sort by index

    digits.first().map(|&(_, first)| first * 10)
        .unwrap_or(0)
        + digits.last().map(|&(_, last)| last).unwrap_or(0)
}

fn calculate_total_part2(filename: &str) -> io::Result<i32> {
    let file = fs::File::open(filename)?;
    let reader = io::BufReader::new(file);
    let total = reader
        .lines()
        .filter_map(|line| line.ok()) // Handle potential errors
        .map(|line| find_calibration_value_part2(&line))
        .sum();

    Ok(total)
}

fn main() -> io::Result<()> {
    let filename = "input.txt"; // Replace with your actual input file

    let total = calculate_total_part2(filename)?;
    println!("Sum of all calibration values (Part 2): {}", total);

    Ok(())
}
