use std::collections::HashSet;
use std::fs;

fn parse_input(filename: &str) -> Vec<Vec<char>> {
    let content = fs::read_to_string(filename).expect("Failed to read the file");
    content
        .lines()
        .map(|line| line.chars().collect())
        .collect()
}

fn find_guard(grid: &[Vec<char>]) -> (usize, usize, usize) {
    let directions = ['^', '>', 'v', '<'];
    for (r, row) in grid.iter().enumerate() {
        for (c, &cell) in row.iter().enumerate() {
            if let Some(idx) = directions.iter().position(|&d| d == cell) {
                return (r, c, idx); // Found the guard's starting position and direction
            }
        }
    }
    panic!("Guard not found on the grid");
}

fn simulate_patrol(grid: Vec<Vec<char>>) -> usize {
    let directions = [
        (-1, 0), // Up
        (0, 1),  // Right
        (1, 0),  // Down
        (0, -1), // Left
    ];

    let (mut r, mut c, mut dir) = find_guard(&grid);
    let rows = grid.len();
    let cols = grid[0].len();
    let mut visited: HashSet<(usize, usize)> = HashSet::new();

    while r < rows && c < cols {
        visited.insert((r, c));

        // Calculate next position
        let (dr, dc) = directions[dir];
        let (nr, nc) = (r as isize + dr, c as isize + dc);

        // Check for obstacles or boundaries
        if nr >= 0
            && nr < rows as isize
            && nc >= 0
            && nc < cols as isize
            && grid[nr as usize][nc as usize] == '#'
        {
            // Turn Right
            dir = (dir + 1) % 4;
        } else {
            // Move Forward
            r = nr as usize;
            c = nc as usize;
        }
    }

    visited.len()
}

fn main() {
    let input_file = "input.txt";
    let grid = parse_input(input_file);
    let result = simulate_patrol(grid);
    println!("Distinct positions visited: {}", result);
}
