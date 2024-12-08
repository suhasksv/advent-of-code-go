use std::collections::{HashMap, HashSet};
use std::fs;

fn manhattan_distance(r1: isize, c1: isize, r2: isize, c2: isize) -> isize {
    (r1 - r2).abs() + (c1 - c2).abs()
}

fn main() {
    let infile = std::env::args().nth(1).unwrap_or("input.txt".to_string());
    let content = fs::read_to_string(&infile).expect("Error reading file");

    let grid: Vec<&str> = content.lines().collect();
    let r = grid.len() as isize;
    let c = grid[0].len() as isize;

    let mut positions: HashMap<char, Vec<(isize, isize)>> = HashMap::new();
    for (row, line) in grid.iter().enumerate() {
        for (col, ch) in line.chars().enumerate() {
            if ch != '.' {
                positions.entry(ch).or_insert_with(Vec::new).push((row as isize, col as isize));
            }
        }
    }

    let mut a1 = HashSet::new();
    let mut a2 = HashSet::new();

    for r in 0..r {
        for c in 0..c {
            for (_ch, pos) in &positions {
                for &(r1, c1) in pos {
                    for &(r2, c2) in pos {
                        if (r1, c1) != (r2, c2) {
                            let d1 = manhattan_distance(r, c, r1, c1);
                            let d2 = manhattan_distance(r, c, r2, c2);
                            let dr1 = r - r1;
                            let dr2 = r - r2;
                            let dc1 = c - c1;
                            let dc2 = c - c2;

                            if (d1 == 2 * d2 || d1 * 2 == d2) && dr1 * dc2 == dc1 * dr2 {
                                a1.insert((r, c));
                            }
                            if dr1 * dc2 == dc1 * dr2 {
                                a2.insert((r, c));
                            }
                        }
                    }
                }
            }
        }
    }

    println!("{}", a1.len());
    println!("{}", a2.len());
}
