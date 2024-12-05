use std::env;
use std::fs;

fn main() {
    // Read input file
    let args: Vec<String> = env::args().collect();
    let infile = args.get(1).unwrap_or(&"4.in".to_string());
    let grid: Vec<String> = fs::read_to_string(infile)
        .expect("Could not read file")
        .lines()
        .map(String::from)
        .collect();

    let rows = grid.len();
    let cols = grid[0].len();
    let mut p1 = 0;
    let mut p2 = 0;

    for r in 0..rows {
        for c in 0..cols {
            // Check for "XMAS" horizontally, vertically, and diagonally
            if c + 3 < cols && is_pattern(&grid, r, c, &['X', 'M', 'A', 'S'], (0, 1)) {
                p1 += 1;
            }
            if r + 3 < rows && is_pattern(&grid, r, c, &['X', 'M', 'A', 'S'], (1, 0)) {
                p1 += 1;
            }
            if r + 3 < rows && c + 3 < cols && is_pattern(&grid, r, c, &['X', 'M', 'A', 'S'], (1, 1)) {
                p1 += 1;
            }
            if c + 3 < cols && is_pattern(&grid, r, c, &['S', 'A', 'M', 'X'], (0, 1)) {
                p1 += 1;
            }
            if r + 3 < rows && is_pattern(&grid, r, c, &['S', 'A', 'M', 'X'], (1, 0)) {
                p1 += 1;
            }
            if r + 3 < rows && c + 3 < cols && is_pattern(&grid, r, c, &['S', 'A', 'M', 'X'], (1, 1)) {
                p1 += 1;
            }
            if r >= 3 && c + 3 < cols && is_pattern(&grid, r, c, &['S', 'A', 'M', 'X'], (-1, 1)) {
                p1 += 1;
            }
            if r >= 3 && c + 3 < cols && is_pattern(&grid, r, c, &['X', 'M', 'A', 'S'], (-1, 1)) {
                p1 += 1;
            }

            // Check for "MAS" patterns surrounded by "M/S"
            if r + 2 < rows && c + 2 < cols {
                if is_surrounded_pattern(&grid, r, c, &['M', 'A', 'S'], 'M', 'S') {
                    p2 += 1;
                }
                if is_surrounded_pattern(&grid, r, c, &['M', 'A', 'S'], 'S', 'M') {
                    p2 += 1;
                }
                if is_surrounded_pattern(&grid, r, c, &['S', 'A', 'M'], 'M', 'S') {
                    p2 += 1;
                }
                if is_surrounded_pattern(&grid, r, c, &['S', 'A', 'M'], 'S', 'M') {
                    p2 += 1;
                }
            }
        }
    }

    println!("{}", p1);
    println!("{}", p2);
}

fn is_pattern(
    grid: &Vec<String>,
    r: usize,
    c: usize,
    pattern: &[char],
    direction: (isize, isize),
) -> bool {
    for (i, &ch) in pattern.iter().enumerate() {
        let nr = r as isize + i as isize * direction.0;
        let nc = c as isize + i as isize * direction.1;
        if nr < 0 || nc < 0 || nr >= grid.len() as isize || nc >= grid[0].len() as isize {
            return false;
        }
        if grid[nr as usize].chars().nth(nc as usize).unwrap_or(' ') != ch {
            return false;
        }
    }
    true
}

fn is_surrounded_pattern(
    grid: &Vec<String>,
    r: usize,
    c: usize,
    pattern: &[char],
    top_left: char,
    bottom_right: char,
) -> bool {
    if !is_pattern(grid, r, c, pattern, (1, 1)) {
        return false;
    }
    let br = grid[r + 2].chars().nth(c).unwrap_or(' ');
    let tr = grid[r].chars().nth(c + 2).unwrap_or(' ');
    br == top_left && tr == bottom_right
}
