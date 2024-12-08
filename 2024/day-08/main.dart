import 'dart:io';
import 'dart:math';

int manhattanDistance(int r1, int c1, int r2, int c2) {
  return (r1 - r2).abs() + (c1 - c2).abs();
}

void main(List<String> args) {
  String infile = args.isNotEmpty ? args[0] : 'input.txt';
  List<String> grid = File(infile).readAsLinesSync();

  int r = grid.length;
  int c = grid[0].length;

  Map<String, List<List<int>>> positions = {};
  for (int i = 0; i < r; i++) {
    for (int j = 0; j < c; j++) {
      String char = grid[i][j];
      if (char != '.') {
        positions.putIfAbsent(char, () => []).add([i, j]);
      }
    }
  }

  Set<List<int>> a1 = {};
  Set<List<int>> a2 = {};

  for (int i = 0; i < r; i++) {
    for (int j = 0; j < c; j++) {
      positions.forEach((_, pos) {
        for (var p1 in pos) {
          for (var p2 in pos) {
            if (p1 != p2) {
              int d1 = manhattanDistance(i, j, p1[0], p1[1]);
              int d2 = manhattanDistance(i, j, p2[0], p2[1]);
              int dr1 = i - p1[0];
              int dr2 = i - p2[0];
              int dc1 = j - p1[1];
              int dc2 = j - p2[1];

              if ((d1 == 2 * d2 || d1 * 2 == d2) && dr1 * dc2 == dc1 * dr2) {
                a1.add([i, j]);
              }
              if (dr1 * dc2 == dc1 * dr2) {
                a2.add([i, j]);
              }
            }
          }
        }
      });
    }
  }

  print(a1.length);
  print(a2.length);
}
