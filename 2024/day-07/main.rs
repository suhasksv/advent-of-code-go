use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn evaluate_expr(expr: &Vec<i64>, ops: &Vec<char>) -> i64 {
    let mut result = expr[0];
    for i in 0..ops.len() {
        match ops[i] {
            '+' => result += expr[i + 1],
            '*' => result *= expr[i + 1],
            '|' => {
                let mut concatenated = result.to_string();
                concatenated.push_str(&expr[i + 1].to_string());
                result = concatenated.parse::<i64>().unwrap();
            }
            _ => panic!("Unknown operator"),
        }
    }
    result
}

fn generate_operators(n: usize) -> Vec<Vec<char>> {
    let operators = vec!['+', '*', '|'];
    let mut all_combinations = Vec::new();
    let mut current_combination = vec!['+'; n];

    loop {
        all_combinations.push(current_combination.clone());

        let mut i = n - 1;
        while i >= 0 {
            if current_combination[i] == '|' {
                current_combination[i] = '+';
                break;
            } else {
                current_combination[i] = match current_combination[i] {
                    '+' => '*',
                    '*' => '|',
                    _ => '+',
                };
                i -= 1;
                if i == 0 {
                    break;
                }
            }
        }

        if current_combination.iter().all(|&c| c == '+') {
            break;
        }
    }

    all_combinations
}

fn read_input(filename: &str) -> io::Result<Vec<(i64, Vec<i64>)>> {
    let path = Path::new(filename);
    let file = File::open(path)?;
    let reader = io::BufReader::new(file);

    let mut equations = Vec::new();

    for line in reader.lines() {
        let line = line?;
        let parts: Vec<&str> = line.split(": ").collect();
        let target = parts[0].parse::<i64>().unwrap();
        let values: Vec<i64> = parts[1]
            .split_whitespace()
            .map(|x| x.parse::<i64>().unwrap())
            .collect();
        equations.push((target, values));
    }

    Ok(equations)
}

fn main() -> io::Result<()> {
    let equations = read_input("input.txt")?;

    let mut total_sum = 0;
    for (target, values) in equations {
        let operators_count = values.len() - 1;
        let operator_combinations = generate_operators(operators_count);

        let mut valid = false;
        for ops in operator_combinations {
            if evaluate_expr(&values, &ops) == target {
                valid = true;
                break;
            }
        }

        if valid {
            total_sum += target;
        }
    }

    println!("Total Calibration Result: {}", total_sum);

    Ok(())
}

