import 'dart:io';

int calculateCalibrationValue(String line) {
  final firstDigit = line.split(RegExp(r'\D+'))[0];
  final lastDigit = line.split(RegExp(r'\D+')).last;
  return int.parse(firstDigit + lastDigit);
}

int calculateTotal(String filename) {
  final file = File(filename);
  final lines = file.readAsLinesSync();
  return lines.fold(0, (sum, line) => sum + calculateCalibrationValue(line));
}

void main() {
  const filename = 'input.txt';
  final result = calculateTotal(filename);
  print('Part 1: The sum of all calibration values is: $result');
}
