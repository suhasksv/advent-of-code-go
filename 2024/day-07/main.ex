defmodule BridgeRepair do
  def parse_input(input_text) do
    input_text
    |> String.split("\n", trim: true)
    |> Enum.map(fn line ->
      [test_value, numbers] = String.split(line, ": ", parts: 2)
      test_value = String.to_integer(test_value)
      numbers = numbers |> String.split() |> Enum.map(&String.to_integer/1)
      {test_value, numbers}
    end)
  end

  def evaluate_left_to_right([number | numbers], [operator | operators]) do
    result =
      case operator do
        "+" -> number + hd(numbers)
        "*" -> number * hd(numbers)
        "||" -> String.to_integer("#{number}#{hd(numbers)}")
      end

    evaluate_left_to_right([result | tl(numbers)], operators)
  end

  def evaluate_left_to_right([result], []), do: result

  def generate_operator_combinations(operators, 0), do: [[]]

  def generate_operator_combinations(operators, length) do
    smaller_combos = generate_operator_combinations(operators, length - 1)

    for combo <- smaller_combos, op <- operators do
      [op | combo]
    end
  end

  def valid_equation?(test_value, numbers, include_concat \ false) do
    operators = if include_concat, do: ["+", "*", "||"], else: ["+", "*"]
    num_operators = length(numbers) - 1

    combinations = generate_operator_combinations(operators, num_operators)

    Enum.any?(combinations, fn combination ->
      evaluate_left_to_right(numbers, combination) == test_value
    end)
  end

  def calculate_total_calibration(input_text, include_concat \ false) do
    parse_input(input_text)
    |> Enum.reduce(0, fn {test_value, numbers}, total ->
      if valid_equation?(test_value, numbers, include_concat) do
        total + test_value
      else
        total
      end
    end)
  end

  def main() do
    input_text = File.read!("input.txt")

    part1_result = calculate_total_calibration(input_text, false)
    IO.puts("Part 1 Total Calibration: #{part1_result}")

    part2_result = calculate_total_calibration(input_text, true)
    IO.puts("Part 2 Total Calibration: #{part2_result}")
  end
end

BridgeRepair.main()

