use std::env;
use std::fs::File;
use std::io::{self, Read};
use regex::Regex;

fn main() -> io::Result<()> {
    // Determine the input file
    let input_file = if let Some(arg) = env::args().nth(1) {
        arg
    } else {
        String::from("input.txt")
    };

    // Open the file
    let mut file = File::open(input_file)?;
    let mut data = String::new();
    file.read_to_string(&mut data)?;

    // Define the regex pattern for mul(x, y)
    let mul_pattern = Regex::new(r"mul\((\d+),(\d+)\)")?;

    let mut part1 = 0;

    // Iterate through the matches of mul(x, y)
    for capture in mul_pattern.captures_iter(&data) {
        if let (Some(x_str), Some(y_str)) = (capture.get(1), capture.get(2)) {
            let x: i32 = x_str.as_str().parse().unwrap();
            let y: i32 = y_str.as_str().parse().unwrap();
            part1 += x * y;
            println!("{}", capture.get(0).unwrap().as_str());
        }
    }

    // Output the result
    println!("{}", part1);

    Ok(())
}
