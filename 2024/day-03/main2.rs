use std::fs;
use std::env;
use regex::Regex;

fn read_file(filename: &str) -> Option<String> {
    match fs::read_to_string(filename) {
        Ok(content) => Some(content),
        Err(_) => {
            eprintln!("Error reading file: {}", filename);
            None
        }
    }
}

fn process_data(data: &str, mut part2: i32, mut enabled: bool) -> i32 {
    let mut i = 0;
    while i < data.len() {
        if &data[i..i + 4] == "do()" {
            enabled = true;
            i += 4;
        } else if &data[i..i + 7] == "don't()" {
            enabled = false;
            i += 7;
        } else if &data[i..i + 4] == "mul(" {
            let (new_part2, new_enabled, j) = process_mul(&data[i..], part2, enabled);
            part2 = new_part2;
            enabled = new_enabled;
            i += j;
        } else {
            i += 1;
        }
    }
    part2
}

fn process_mul(data: &str, part2: i32, enabled: bool) -> (i32, bool, usize) {
    let mut j = 4;
    while &data[j..j + 1] != ")" {
        j += 1;
    }
    let mul_str = &data[4..j];

    // Use regex to extract the two numbers inside "mul(x, y)"
    let re = Regex::new(r"(\d+),\s*(\d+)").unwrap();
    if let Some(caps) = re.captures(mul_str) {
        let x: i32 = caps[1].parse().unwrap();
        let y: i32 = caps[2].parse().unwrap();

        // Make sure the closing parenthesis isn't followed by a number
        if j + 1 < data.len() && !data[j + 1..j + 2].chars().any(|c| c.is_digit(10)) {
            if enabled {
                part2 += x * y;
            }
            println!("{}", &data[0..j + 1]);
        }
    }

    (part2, enabled, j + 1)
}

fn main() {
    // Get the file name from command line arguments or use "input.txt" by default
    let args: Vec<String> = env::args().collect();
    let file = if args.len() > 1 { &args[1] } else { "input.txt" };

    // Read the input data from the file
    let data = read_file(file).expect("Failed to read file");

    // Process the data and get the final result
    let part2 = process_data(&data, 0, true);

    // Print the result
    println!("{}", part2);
}
