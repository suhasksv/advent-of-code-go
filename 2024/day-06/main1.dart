import 'dart:io';
import 'dart:collection';

List<List<String>> parseInput(String filename) {
  return File(filename)
      .readAsLinesSync()
      .map((line) => line.split(''))
      .toList();
}

List<dynamic> findGuard(List<List<String>> grid) {
  List<String> directions = ['^', '>', 'v', '<'];
  for (int r = 0; r < grid.length; r++) {
    for (int c = 0; c < grid[r].length; c++) {
      if (directions.contains(grid[r][c])) {
        return [r, c, directions.indexOf(grid[r][c])];
      }
    }
  }
  throw Exception("Guard not found on the grid");
}

int simulatePatrol(List<List<String>> grid) {
  List<List<int>> directions = [
    [-1, 0], // Up
    [0, 1],  // Right
    [1, 0],  // Down
    [0, -1]  // Left
  ];

  var guardInfo = findGuard(grid);
  int r = guardInfo[0], c = guardInfo[1], dir = guardInfo[2];
  int rows = grid.length, cols = grid[0].length;

  Set<String> visited = {};

  while (r >= 0 && r < rows && c >= 0 && c < cols) {
    visited.add('$r,$c');

    int nr = r + directions[dir][0];
    int nc = c + directions[dir][1];

    if (nr >= 0 &&
        nr < rows &&
        nc >= 0 &&
        nc < cols &&
        grid[nr][nc] == '#') {
      // Turn right
      dir = (dir + 1) % 4;
    } else {
      // Move forward
      r = nr;
      c = nc;
    }
  }

  return visited.length;
}

void main() {
  String inputFile = 'input.txt';
  var grid = parseInput(inputFile);
  int result = simulatePatrol(grid);
  print('Distinct positions visited: $result');
}
