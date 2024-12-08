import 'dart:io';

int evaluateExpr(List<int> expr, List<String> ops) {
  int result = expr[0];
  for (int i = 0; i < ops.length; i++) {
    switch (ops[i]) {
      case '+':
        result += expr[i + 1];
        break;
      case '*':
        result *= expr[i + 1];
        break;
      case '||':
        String concatenated = result.toString() + expr[i + 1].toString();
        result = int.parse(concatenated);
        break;
      default:
        throw Exception("Unknown operator");
    }
  }
  return result;
}

List<List<String>> generateOperators(int n) {
  List<String> operators = ['+', '*', '||'];
  List<List<String>> allCombinations = [];
  List<String> currentCombination = List.filled(n, '+');

  while (true) {
    allCombinations.add(List.from(currentCombination));

    int i = n - 1;
    while (i >= 0) {
      if (currentCombination[i] == '||') {
        currentCombination[i] = '+';
        break;
      } else {
        currentCombination[i] = currentCombination[i] == '+' ? '*' : '||';
        i--;
        if (i == 0) break;
      }
    }

    if (currentCombination.every((e) => e == '+')) {
      break;
    }
  }

  return allCombinations;
}

Future<List<MapEntry<int, List<int>>>> readInput(String filename) async {
  List<MapEntry<int, List<int>>> equations = [];

  final file = File(filename);
  final lines = await file.readAsLines();

  for (var line in lines) {
    var parts = line.split(": ");
    int target = int.parse(parts[0]);
    List<int> values = parts[1]
        .split(" ")
        .map((x) => int.parse(x))
        .toList();
    equations.add(MapEntry(target, values));
  }

  return equations;
}

void main() async {
  List<MapEntry<int, List<int>>> equations = await readInput("input.txt");

  int totalSum = 0;
  for (var eq in equations) {
    int target = eq.key;
    List<int> values = eq.value;
    int operatorsCount = values.length - 1;
    List<List<String>> operatorCombinations = generateOperators(operatorsCount);

    bool valid = false;
    for (var ops in operatorCombinations) {
      if (evaluateExpr(values, ops) == target) {
        valid = true;
        break;
      }
    }

    if (valid) {
      totalSum += target;
    }
  }

  print('Total Calibration Result: $totalSum');
}

