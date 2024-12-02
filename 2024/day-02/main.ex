defmodule ReactorSafety do
  def is_safe(report) do
    differences = Enum.chunk_every(report, 2, 1, :discard) |> Enum.map(fn [a, b] -> b - a end)

    valid_differences = Enum.all?(differences, fn diff -> diff in -3..-1 or diff in 1..3 end)
    all_increasing = Enum.all?(differences, &(&1 > 0))
    all_decreasing = Enum.all?(differences, &(&1 < 0))

    valid_differences and (all_increasing or all_decreasing)
  end

  def can_be_made_safe(report) do
    Enum.any?(0..(length(report) - 1), fn i ->
      modified = List.delete_at(report, i)
      is_safe(modified)
    end)
  end

  def count_safe_reports(file_path, part) do
    file_path
    |> File.stream!()
    |> Enum.map(&String.trim/1)
    |> Enum.reject(&(&1 == ""))
    |> Enum.map(fn line -> String.split(line) |> Enum.map(&String.to_integer/1) end)
    |> Enum.count(fn report ->
      case part do
        1 -> is_safe(report)
        2 -> is_safe(report) or can_be_made_safe(report)
      end
    end)
  end
end

# Main program
file_path = "input.txt"

part1 = ReactorSafety.count_safe_reports(file_path, 1)
IO.puts("Part 1: Number of safe reports: #{part1}")

part2 = ReactorSafety.count_safe_reports(file_path, 2)
IO.puts("Part 2: Number of safe reports: #{part2}")
