use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn is_safe(report: &Vec<i32>) -> bool {
    // Compute the differences between adjacent levels
    let differences: Vec<i32> = report
        .windows(2)
        .map(|w| w[1] - w[0])
        .collect();

    // Check if all differences are within [-3, -1] or [1, 3]
    for &diff in &differences {
        if diff.abs() < 1 || diff.abs() > 3 {
            return false;
        }
    }

    // Check if the differences are all positive (increasing) or all negative (decreasing)
    let all_positive = differences.iter().all(|&diff| diff > 0);
    let all_negative = differences.iter().all(|&diff| diff < 0);

    all_positive || all_negative
}

fn can_be_made_safe(report: &Vec<i32>) -> bool {
    for i in 0..report.len() {
        // Create a new report by skipping the current element
        let mut modified_report = report.clone();
        modified_report.remove(i);

        // Check if the modified report is safe
        if is_safe(&modified_report) {
            return true;
        }
    }
    false
}

fn main() -> io::Result<()> {
    // Open the input file
    let path = "input.txt";
    let file = File::open(&path)?;
    let reader = io::BufReader::new(file);

    let mut safe_count = 0;

    for line in reader.lines() {
        let line = line?;
        if line.trim().is_empty() {
            continue;
        }

        // Parse the line into a vector of integers
        let report: Vec<i32> = line
            .split_whitespace()
            .map(|num| num.parse::<i32>().unwrap())
            .collect();

        // Check if the report is safe or can be made safe
        if is_safe(&report) || can_be_made_safe(&report) {
            safe_count += 1;
        }
    }

    // Print the number of safe reports
    println!("Number of safe reports: {}", safe_count);

    Ok(())
}
