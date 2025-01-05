use std::collections::{HashSet, VecDeque};
use std::fs;

type Point = (usize, usize);

fn parse_input(file_name: &str) -> Vec<Point> {
    let input = fs::read_to_string(file_name).expect("Failed to read input file");
    input
        .lines()
        .map(|line| {
            let coords: Vec<usize> = line.split(',').map(|x| x.parse().unwrap()).collect();
            (coords[0], coords[1])
        })
        .collect()
}

fn is_within_bounds(x: isize, y: isize, size: usize) -> bool {
    x >= 0 && x < size as isize && y >= 0 && y < size as isize
}

fn bfs(grid: &HashSet<Point>, start: Point, goal: Point, size: usize) -> bool {
    let directions = [(0, 1), (1, 0), (0, -1), (-1, 0)];
    let mut queue = VecDeque::new();
    let mut visited = HashSet::new();

    queue.push_back(start);
    visited.insert(start);

    while let Some((x, y)) = queue.pop_front() {
        if (x, y) == goal {
            return true;
        }

        for &(dx, dy) in &directions {
            let nx = x as isize + dx;
            let ny = y as isize + dy;
            if is_within_bounds(nx, ny, size) {
                let neighbor = (nx as usize, ny as usize);
                if !grid.contains(&neighbor) && visited.insert(neighbor) {
                    queue.push_back(neighbor);
                }
            }
        }
    }
    false
}

fn find_shortest_path(grid: &HashSet<Point>, start: Point, goal: Point, size: usize) -> Option<usize> {
    let directions = [(0, 1), (1, 0), (0, -1), (-1, 0)];
    let mut queue = VecDeque::new();
    let mut visited = HashSet::new();

    queue.push_back((start, 0));
    visited.insert(start);

    while let Some(((x, y), steps)) = queue.pop_front() {
        if (x, y) == goal {
            return Some(steps);
        }

        for &(dx, dy) in &directions {
            let nx = x as isize + dx;
            let ny = y as isize + dy;
            if is_within_bounds(nx, ny, size) {
                let neighbor = (nx as usize, ny as usize);
                if !grid.contains(&neighbor) && visited.insert(neighbor) {
                    queue.push_back((neighbor, steps + 1));
                }
            }
        }
    }
    None
}

fn main() {
    let file_name = "input.txt";
    let points = parse_input(file_name);
    let size = 71;
    let start = (0, 0);
    let goal = (70, 70);

    // Part 1: Shortest path after 1024 bytes
    let mut grid = HashSet::new();
    for i in 0..1024.min(points.len()) {
        grid.insert(points[i]);
    }
    if let Some(shortest_path) = find_shortest_path(&grid, start, goal, size) {
        println!("Shortest Path after 1024 bytes: {}", shortest_path);
    } else {
        println!("No valid path to the exit after 1024 bytes.");
    }

    // Part 2: First byte that blocks the path
    grid.clear();
    for (i, &point) in points.iter().enumerate() {
        grid.insert(point);
        if !bfs(&grid, start, goal, size) {
            println!("First blocking byte: {},{}", point.0, point.1);
            break;
        }
    }
}
