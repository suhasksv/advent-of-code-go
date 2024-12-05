import 'dart:io';
import 'dart:convert';

void main() async {
  // Read the input file
  var file = File('input.txt');
  String content = await file.readAsString();

  // Regular expression to match mul(x, y)
  RegExp regExp = RegExp(r'mul\((\d+), (\d+)\)');

  // Find all matches
  Iterable<RegExpMatch> matches = regExp.allMatches(content);

  // Calculate the sum of x * y
  int totalSum = 0;

  for (var match in matches) {
    int x = int.parse(match.group(1)!);
    int y = int.parse(match.group(2)!);
    totalSum += x * y;
  }

  print('Total sum of mul(x, y): $totalSum');
}
