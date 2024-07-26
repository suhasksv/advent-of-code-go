import 'dart:io';

const digitMap = {
  'one': 1,
  'two': 2,
  'three': 3,
  'four': 4,
  'five': 5,
  'six': 6,
  'seven': 7,
  'eight': 8,
  'nine': 9,
};

int findCalibrationValuePart2(String line) {
  final digits = <MapEntry<int, int>>[];

  // Use a modified regex to prevent overlapping matches
  final digitRegExp = RegExp(r'(?=(\d|one|two|three|four|five|six|seven|eight|nine))');

  for (final match in digitRegExp.allMatches(line)) {
    final valueStr = match.group(1)!; // Extract the matched string from group 1
    final value = int.tryParse(valueStr) ?? digitMap[valueStr]!;
    digits.add(MapEntry(match.start, value));
  }

  digits.sort((a, b) => a.key.compareTo(b.key));
  return digits.isNotEmpty ? digits.first.value * 10 + digits.last.value : 0;
}

Future<int> calculateTotalPart2(String filename) async {
  try {
    final file = File(filename);
    final lines = await file.readAsLines();
    return lines.fold<int>(0, (sum, line) => sum + findCalibrationValuePart2(line));
  } catch (e) {
    stderr.writeln("Error reading file: $e");
    return 0;
  }
}


Future<void> main() async {
  final filename = 'input.txt';
  final total = await calculateTotalPart2(filename);
  print('Sum of all calibration values (Part 2): $total');
}

/*
// Revisions for adjusting and correcting the logic

int findCalibrationValuePart2(String line) {
  final digits = <MapEntry<int, int>>[]; // Value-Index pairs
  final digitsString = digitMap.keys.join('|');
  final digitRegExp = RegExp(r'(\d|${digitsString})');

  for (final match in digitRegExp.allMatches(line)) {
    final valueStr = match.group(0)!;
    final value = int.tryParse(valueStr) ?? digitMap[valueStr]!;
    digits.add(MapEntry(value, match.start));
  }

  digits.sort((a, b) => a.value.compareTo(b.value)); // Sort by index

  return digits.isNotEmpty ? digits.first.key * 10 + digits.last.key : 0;
}

int findCalibrationValuePart2(String line) {
  final digits = <MapEntry<int, int>>[]; // Value-Index pairs

  // Define a regular expression to match both digit characters and digit words
  final digitRegExp = RegExp(r'(\d|one|two|three|four|five|six|seven|eight|nine)');

  for (final match in digitRegExp.allMatches(line)) {
    final valueStr = match.group(0)!;
    final value = int.tryParse(valueStr) ?? digitMap[valueStr]!;
    digits.add(MapEntry(match.start, value)); // Store index first, then value
  }

  // Sort by index (ascending order) to find the first and last digits
  digits.sort((a, b) => a.key.compareTo(b.key));

  return digits.isNotEmpty ? digits.first.value * 10 + digits.last.value : 0;
}

*/