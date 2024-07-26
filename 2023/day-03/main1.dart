import 'dart:io';

// Directions for adjacent cells (8 directions)
final List<List<int>> directions = [
  [-1, -1], [-1, 0], [-1, 1],
  [0, -1],          [0, 1],
  [1, -1], [1, 0], [1, 1]
];

// Function to check if a cell is within the grid bounds
bool inBounds(int x, int y, List<List<String>> grid) {
  return x >= 0 && y >= 0 && x < grid.length && y < grid[0].length;
}

// Function to get numbers from the grid that are adjacent to any special symbol
int sumAdjacentNumbers(List<List<String>> grid) {
  int totalSum = 0;
  int rows = grid.length;
  int cols = grid[0].length;

  for (int r = 0; r < rows; r++) {
    for (int c = 0; c < cols; c++) {
      if (['*', '#', '+', r'\$'].contains(grid[r][c])) {
        // Check all 8 adjacent cells
        for (var direction in directions) {
          int nr = r + direction[0];
          int nc = c + direction[1];
          if (inBounds(nr, nc, grid)) {
            String cell = grid[nr][nc];
            if (RegExp(r'^\d+$').hasMatch(cell)) {
              totalSum += int.parse(cell);
            }
          }
        }
      }
    }
  }

  return totalSum;
}

void main() async {
  const filePath = 'input.txt';
  final file = File(filePath);

  if (!await file.exists()) {
    print('File not found: $filePath');
    return;
  }

  final lines = await file.readAsLines();
  List<List<String>> grid = lines.map((line) => line.split('')).toList();

  int totalSum = sumAdjacentNumbers(grid);
  print('Sum of all part numbers: $totalSum');
}
