import 'dart:io';
import 'dart:collection';

typedef Point = List<int>;

List<Point> parseInput(String fileName) {
  final input = File(fileName).readAsLinesSync();
  return input.map((line) {
    final coords = line.split(',').map(int.parse).toList();
    return [coords[0], coords[1]];
  }).toList();
}

bool isWithinBounds(int x, int y, int size) {
  return x >= 0 && x < size && y >= 0 && y < size;
}

bool bfs(Set<Point> grid, Point start, Point goal, int size) {
  final directions = [
    [0, 1],
    [1, 0],
    [0, -1],
    [-1, 0]
  ];
  final queue = Queue<Point>();
  final visited = <Point>{};

  queue.add(start);
  visited.add(start);

  while (queue.isNotEmpty) {
    final current = queue.removeFirst();
    if (current[0] == goal[0] && current[1] == goal[1]) {
      return true;
    }

    for (final direction in directions) {
      final nx = current[0] + direction[0];
      final ny = current[1] + direction[1];

      if (isWithinBounds(nx, ny, size)) {
        final neighbor = [nx, ny];
        if (!grid.contains(neighbor) && visited.add(neighbor)) {
          queue.add(neighbor);
        }
      }
    }
  }
  return false;
}

int? findShortestPath(Set<Point> grid, Point start, Point goal, int size) {
  final directions = [
    [0, 1],
    [1, 0],
    [0, -1],
    [-1, 0]
  ];
  final queue = Queue<MapEntry<Point, int>>();
  final visited = <Point>{};

  queue.add(MapEntry(start, 0));
  visited.add(start);

  while (queue.isNotEmpty) {
    final entry = queue.removeFirst();
    final current = entry.key;
    final steps = entry.value;

    if (current[0] == goal[0] && current[1] == goal[1]) {
      return steps;
    }

    for (final direction in directions) {
      final nx = current[0] + direction[0];
      final ny = current[1] + direction[1];

      if (isWithinBounds(nx, ny, size)) {
        final neighbor = [nx, ny];
        if (!grid.contains(neighbor) && visited.add(neighbor)) {
          queue.add(MapEntry(neighbor, steps + 1));
        }
      }
    }
  }
  return null;
}

void main() {
  const fileName = 'input.txt';
  const size = 71;
  final points = parseInput(fileName);
  final start = [0, 0];
  final goal = [70, 70];

  // Part 1: Shortest path after 1024 bytes
  final gridPart1 = <Point>{};
  for (int i = 0; i < points.length && i < 1024; i++) {
    gridPart1.add(points[i]);
  }
  final shortestPath = findShortestPath(gridPart1, start, goal, size);
  if (shortestPath != null) {
    print('Shortest Path after 1024 bytes: $shortestPath');
  } else {
    print('No valid path to the exit after 1024 bytes.');
  }

  // Part 2: First byte that blocks the path
  final gridPart2 = <Point>{};
  for (int i = 0; i < points.length; i++) {
    gridPart2.add(points[i]);
    if (!bfs(gridPart2, start, goal, size)) {
      print('First blocking byte: ${points[i][0]},${points[i][1]}');
      break;
    }
  }
}
