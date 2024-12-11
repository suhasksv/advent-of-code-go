import 'dart:io';
import 'dart:convert';

class Memoization {
  final Map<String, int> _cache = {};

  int solve(int x, int t) {
    final key = '$x,$t';
    if (_cache.containsKey(key)) {
      return _cache[key]!;
    }

    int result;
    if (t == 0) {
      result = 1;
    } else if (x == 0) {
      result = solve(1, t - 1);
    } else if (x.toString().length % 2 == 0) {
      final dstr = x.toString();
      final mid = dstr.length ~/ 2;
      final left = int.parse(dstr.substring(0, mid));
      final right = int.parse(dstr.substring(mid));
      result = solve(left, t - 1) + solve(right, t - 1);
    } else {
      result = solve(x * 2024, t - 1);
    }

    _cache[key] = result;
    return result;
  }
}

int solveAll(List<int> d, int t) {
  final memo = Memoization();
  return d.map((x) => memo.solve(x, t)).reduce((a, b) => a + b);
}

void main() {
  // Read input from input.txt
  final data = File('input.txt').readAsStringSync().trim();
  final d = data.split(RegExp(r'\s+')).map(int.parse).toList();

  // Calculate results for 25 and 75 blinks
  final result25 = solveAll(d, 25);
  final result75 = solveAll(d, 75);

  print('Result after 25 blinks: $result25');
  print('Result after 75 blinks: $result75');
}
