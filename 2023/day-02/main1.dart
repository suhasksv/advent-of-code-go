import 'dart:io';

void main() {
  final gameRecords = File('input.txt').readAsLinesSync();
  final maxRed = 12;
  final maxGreen = 13;
  final maxBlue = 14;
  var sumOfPossibleGameIds = 0;

  for (var record in gameRecords) {
    final (gameId, subsets) = parseGameRecord(record.trim());
    if (isGamePossible(subsets, maxRed, maxGreen, maxBlue)) {
      sumOfPossibleGameIds += gameId;
    }
  }

  print(sumOfPossibleGameIds);
}

(int, List<String>) parseGameRecord(String record) {
  final parts = record.split(': ');
  final gameId = int.parse(parts[0].split(' ')[1]);
  final subsets = parts[1].split('; ');
  return (gameId, subsets);
}

bool isGamePossible(List<String> subsets, int maxRed, int maxGreen, int maxBlue) {
  for (var subset in subsets) {
    final counts = {'red': 0, 'green': 0, 'blue': 0};
    final cubes = subset.split(', ');
    for (var cube in cubes) {
      final parts = cube.split(' ');
      final num = int.parse(parts[0]);
      final color = parts[1];
      counts[color] = (counts[color] ?? 0) + num; // If the color doesn't exist, initialize to 0
    }
    if (counts['red']! > maxRed || counts['green']! > maxGreen || counts['blue']! > maxBlue) {
      return false;
    }
  }
  return true;
}
