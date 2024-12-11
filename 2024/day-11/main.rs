use std::collections::HashMap;
use std::fs;
use std::str::FromStr;

fn solve(x: i32, t: i32, dp: &mut HashMap<(i32, i32), i32>) -> i32 {
    if let Some(&val) = dp.get(&(x, t)) {
        return val;
    }

    let ret = if t == 0 {
        1
    } else if x == 0 {
        solve(1, t - 1, dp)
    } else if x.to_string().len() % 2 == 0 {
        let dstr = x.to_string();
        let mid = dstr.len() / 2;
        let left: i32 = i32::from_str(&dstr[..mid]).unwrap();
        let right: i32 = i32::from_str(&dstr[mid..]).unwrap();
        solve(left, t - 1, dp) + solve(right, t - 1, dp)
    } else {
        solve(x * 2024, t - 1, dp)
    };

    dp.insert((x, t), ret);
    ret
}

fn solve_all(d: &[i32], t: i32) -> i32 {
    let mut dp = HashMap::new();
    d.iter().map(|&x| solve(x, t, &mut dp)).sum()
}

fn main() {
    // Read input from input.txt
    let data = fs::read_to_string("input.txt").expect("Error reading file");
    let d: Vec<i32> = data
        .trim()
        .split_whitespace()
        .map(|x| i32::from_str(x).expect("Error parsing number"))
        .collect();

    // Calculate results for 25 and 75 blinks
    let result25 = solve_all(&d, 25);
    let result75 = solve_all(&d, 75);

    println!("Result after 25 blinks: {}", result25);
    println!("Result after 75 blinks: {}", result75);
}
