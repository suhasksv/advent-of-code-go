import 'dart:io';

int findCalibrationValuePart1(String line) {
  final matches = RegExp(r'\d').allMatches(line);
  final firstDigit = int.parse(matches.first.group(0)!);
  final lastDigit = int.parse(matches.last.group(0)!);

  return firstDigit * 10 + lastDigit;
}

Future<int> calculateTotalPart1(String filename) async {
  try {
    final file = File(filename);
    final data = await file.readAsLines();
    var total = 0;
    for (final line in data) {
      if (line.isNotEmpty) {
        total += findCalibrationValuePart1(line);
      }
    }
    return total;
  } catch (e) {
    stderr.writeln("Error: Failure to read file: $e");
    return 0;
  }
}

Future<void> main() async {
  final filename = 'input.txt';

  final total = await calculateTotalPart1(filename);
  print('Sum of all calibration values: $total');
}
