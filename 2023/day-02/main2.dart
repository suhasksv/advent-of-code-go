import 'dart:io';

// Class to represent the count of cubes of each color
class CubeCount {
  int red;
  int green;
  int blue;

  CubeCount({required this.red, required this.green, required this.blue});

  // Finds the maximum number of cubes needed for each color from a list
  static CubeCount max(List<CubeCount> counts) {
    int maxRed = 0, maxGreen = 0, maxBlue = 0;
    for (var count in counts) {
      if (count.red > maxRed) maxRed = count.red;
      if (count.green > maxGreen) maxGreen = count.green;
      if (count.blue > maxBlue) maxBlue = count.blue;
    }
    return CubeCount(red: maxRed, green: maxGreen, blue: maxBlue);
  }
}

// Function to parse game data from a line
List<CubeCount> parseGameData(String line) {
  var subsetsStr = line.split(': ')[1];
  List<CubeCount> gameData = [];

  for (var subset in subsetsStr.split('; ')) {
    var subsetCounts = {'red': 0, 'green': 0, 'blue': 0};
    for (var match in RegExp(r'(\d+) (\w+)').allMatches(subset)) {
      var count = int.parse(match.group(1)!);
      var color = match.group(2)!;
      subsetCounts[color] = subsetCounts[color]! + count;
    }
    gameData.add(CubeCount(
      red: subsetCounts['red']!,
      green: subsetCounts['green']!,
      blue: subsetCounts['blue']!,
    ));
  }

  return gameData;
}

// Function to calculate the power of a given set of cubes
int calculatePower(int red, int green, int blue) {
  return red * green * blue;
}

// Function to find the total power of the minimum sets of cubes
Future<int> findTotalPower(String filePath) async {
  final file = File(filePath);
  if (!await file.exists()) {
    throw Exception('File not found: $filePath');
  }

  final lines = await file.readAsLines();
  int totalPower = 0;

  for (var line in lines) {
    if (line.isNotEmpty) {
      var gameData = parseGameData(line);
      var minCubes = CubeCount.max(gameData);
      var power = calculatePower(minCubes.red, minCubes.green, minCubes.blue);
      totalPower += power;
    }
  }

  return totalPower;
}

void main() async {
  const filePath = 'input.txt';
  try {
    final totalPower = await findTotalPower(filePath);
    print('Sum of powers of the minimum sets of cubes: $totalPower');
  } catch (e) {
    print('Error: $e');
  }
}
