import 'dart:io';

bool isSafe(List<int> report) {
  for (int i = 0; i < report.length - 1; i++) {
    int diff = report[i + 1] - report[i];
    if (diff.abs() < 1 || diff.abs() > 3) return false;
  }

  bool allIncreasing = true;
  bool allDecreasing = true;

  for (int i = 0; i < report.length - 1; i++) {
    if (report[i + 1] <= report[i]) allIncreasing = false;
    if (report[i + 1] >= report[i]) allDecreasing = false;
  }

  return allIncreasing || allDecreasing;
}

bool canBeMadeSafe(List<int> report) {
  for (int i = 0; i < report.length; i++) {
    List<int> modified = List.from(report)..removeAt(i);
    if (isSafe(modified)) return true;
  }
  return false;
}

int countSafeReports(String filePath, int part) {
  final lines = File(filePath).readAsLinesSync();
  int safeCount = 0;

  for (var line in lines) {
    if (line.trim().isEmpty) continue;

    List<int> report = line.split(' ').map(int.parse).toList();
    if (part == 1) {
      if (isSafe(report)) safeCount++;
    } else if (part == 2) {
      if (isSafe(report) || canBeMadeSafe(report)) safeCount++;
    }
  }

  return safeCount;
}

void main() {
  const path = 'input.txt';
  print('Part 1: Number of safe reports: ${countSafeReports(path, 1)}');
  print('Part 2: Number of safe reports: ${countSafeReports(path, 2)}');
}
