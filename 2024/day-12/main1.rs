use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::collections::VecDeque;

#[derive(Copy, Clone)]
struct Point {
    x: usize,
    y: usize,
}

fn calculate_fencing_price(map_input: &Vec<Vec<char>>) -> i32 {
    let rows = map_input.len();
    let cols = map_input[0].len();
    let mut visited = vec![vec![false; cols]; rows];

    let is_valid = |x: isize, y: isize| -> bool {
        x >= 0 && x < rows as isize && y >= 0 && y < cols as isize
    };

    let bfs = |start_x: usize, start_y: usize| -> Vec<Point> {
        let mut queue = VecDeque::new();
        let mut region_cells = Vec::new();
        let plant_type = map_input[start_x][start_y];

        queue.push_back(Point { x: start_x, y: start_y });
        visited[start_x][start_y] = true;

        while let Some(current) = queue.pop_front() {
            region_cells.push(current);

            for &(dx, dy) in &[(-1, 0), (1, 0), (0, -1), (0, 1)] {
                let nx = current.x as isize + dx;
                let ny = current.y as isize + dy;

                if is_valid(nx, ny) {
                    let nx = nx as usize;
                    let ny = ny as usize;
                    if !visited[nx][ny] && map_input[nx][ny] == plant_type {
                        visited[nx][ny] = true;
                        queue.push_back(Point { x: nx, y: ny });
                    }
                }
            }
        }

        region_cells
    };

    let calculate_area_and_perimeter = |region_cells: &Vec<Point>| -> (i32, i32) {
        let area = region_cells.len() as i32;
        let mut perimeter = 0;

        for &cell in region_cells {
            for &(dx, dy) in &[(-1, 0), (1, 0), (0, -1), (0, 1)] {
                let nx = cell.x as isize + dx;
                let ny = cell.y as isize + dy;

                if !is_valid(nx, ny) || map_input[nx as usize][ny as usize] != map_input[cell.x][cell.y] {
                    perimeter += 1;
                }
            }
        }

        (area, perimeter)
    };

    let mut total_price = 0;

    for i in 0..rows {
        for j in 0..cols {
            if !visited[i][j] {
                let region_cells = bfs(i, j);
                let (area, perimeter) = calculate_area_and_perimeter(&region_cells);
                total_price += area * perimeter;
            }
        }
    }

    total_price
}

fn main() -> io::Result<()> {
    let path = Path::new("input.txt");
    let file = File::open(&path)?;
    let reader = io::BufReader::new(file);

    let map_input: Vec<Vec<char>> = reader
        .lines()
        .map(|line| line.unwrap().chars().collect())
        .collect();

    let result = calculate_fencing_price(&map_input);
    println!("Total price of fencing: {}", result);

    Ok(())
}

