import 'dart:io';
import 'dart:convert';
import 'dart:core';

void main(List<String> arguments) {
  // Set the file to read input from
  String file = arguments.isNotEmpty ? arguments[0] : 'input.txt';

  // Read the input data from the file
  String data = readFile(file);

  // Process the data
  int part2 = processData(data, 0, true);

  // Output the final result
  print(part2);
}

String readFile(String filename) {
  try {
    File file = File(filename);
    return file.readAsStringSync();
  } catch (e) {
    print("Error reading file: $filename");
    exit(1);
  }
}

int processData(String data, int part2, bool enabled) {
  int i = 0;
  while (i < data.length) {
    if (data.substring(i, i + 4) == 'do()') {
      enabled = true;
      i += 4;
    } else if (data.substring(i, i + 7) == "don't()") {
      enabled = false;
      i += 7;
    } else if (data.substring(i, i + 4) == 'mul(') {
      var result = processMul(data.substring(i), part2, enabled);
      part2 = result[0];
      enabled = result[1];
      i += result[2];
    } else {
      i++;
    }
  }
  return part2;
}

List<dynamic> processMul(String data, int part2, bool enabled) {
  int j = 4;
  // Find the closing parenthesis
  while (j < data.length && data[j] != ')') {
    j++;
  }

  String mulStr = data.substring(4, j);

  // Use a regular expression to extract numbers inside "mul(x, y)"
  final regExp = RegExp(r'(\d+),\s*(\d+)');
  final match = regExp.firstMatch(mulStr);

  if (match != null) {
    int x = int.parse(match.group(1)!);
    int y = int.parse(match.group(2)!);

    // Make sure the closing parenthesis isn't followed by a number
    if (j + 1 < data.length && !RegExp(r'\d').hasMatch(data[j + 1])) {
      if (enabled) {
        part2 += x * y;
      }
      print(data.substring(0, j + 1));
    }
  }

  return [part2, enabled, j + 1];
}
