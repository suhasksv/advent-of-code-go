# Not working
defmodule Aoc2417 do
  use Bitwise  # Import Bitwise module for bitwise operations

  def part1(input) do
    {regs, program} = parse_input(input)

    # Extract Register A
    a = parse_register_a(regs)

    # Execute the program
    execution_result = exec_program(a, program)

    # Convert the execution result to a single integer
    execution_result
    |> Enum.join()
    |> String.to_integer()
  end

  def part2(input) do
    {_, program} = parse_input(input)
    dfs(program, 0, 0, :infinity)
  end

  defp dfs(_, _, pos, success) when pos < 0, do: success

  defp dfs(program, cur, pos, success) do
    Enum.reduce(0..7, success, fn i, acc ->
      next_num = (cur <<< 3) + i
      exec_result = exec_program(next_num, program)
      end_idx = length(program) - pos - 1

      if Enum.slice(exec_result, -end_idx..-1) == Enum.slice(program, -end_idx..-1) do
        dfs(program, next_num, pos + 1, min(acc, next_num))
      else
        acc
      end
    end)
  end

  defp exec_program(a, program) do
    exec_program_recursive(a, program, 0, 0, 0, [], 0)
  end

  defp exec_program_recursive(a, program, pointer, b, c, output, _) when pointer >= length(program),
    do: output

  defp exec_program_recursive(a, program, pointer, b, c, output, _) do
    opcode = Enum.at(program, pointer)
    literal_operand = Enum.at(program, pointer + 1)

    combo =
      case literal_operand do
        4 -> a
        5 -> b
        6 -> c
        _ -> literal_operand
      end

    {new_a, new_b, new_c, new_output, new_pointer} =
      case opcode do
        0 -> {a >>> combo, b, c, output, pointer + 2}
        1 -> {a, b ^^^ literal_operand, c, output, pointer + 2}
        2 -> {a, combo &&& 7, c, output, pointer + 2}
        3 -> if a != 0, do: {a, b, c, output, literal_operand}, else: {a, b, c, output, pointer + 2}
        4 -> {a, b ^^^ c, c, output, pointer + 2}
        5 -> {a, b, c, output ++ [combo &&& 7], pointer + 2}
        6 -> {a, a >>> combo, c, output, pointer + 2}
        7 -> {a, b, a >>> combo, output, pointer + 2}
      end

    exec_program_recursive(new_a, program, new_pointer, new_b, new_c, new_output, pointer + 2)
  end

  defp parse_input(input) do
    [regs, program_str] = String.split(input, "\n\n")

    program =
      program_str
      |> String.replace("Program: ", "")
      |> String.split(",")
      |> Enum.map(&String.to_integer/1)

    {regs, program}
  end

  defp parse_register_a(regs) do
    regs
    |> String.split("\n")
    |> Enum.find(fn line -> String.contains?(line, "Register A:") end)
    |> String.split(":")
    |> List.last()
    |> String.trim()
    |> String.to_integer()
  end

  def get_real_input do
    """
    Register A: 28066687
    Register B: 0
    Register C: 0

    Program: 2,4,1,1,7,5,4,6,0,3,1,4,5,5,3,0
    """
  end
end

# Running the module
input = Aoc2417.get_real_input()
IO.puts("Part 1: #{Aoc2417.part1(input)}")
IO.puts("Part 2: #{Aoc2417.part2(input)}")
