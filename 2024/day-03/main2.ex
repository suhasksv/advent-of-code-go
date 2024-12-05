defmodule Parser do
  # Function to read the file content
  def read_file(filename) do
    case File.read(filename) do
      {:ok, content} -> content
      {:error, _} ->
        IO.puts("Error reading file: #{filename}")
        nil
    end
  end

  def process do
    # Set the file to read input from
    file = if length(System.argv()) > 0, do: List.first(System.argv()), else: "input.txt"

    # Read the input data from the file
    D = read_file(file)

    if D == nil, do: return

    part2 = process_data(D, 0, true)

    IO.puts(part2)
  end

  defp process_data("", part2, _enabled), do: part2
  defp process_data(D, part2, enabled) do
    {new_part2, new_enabled, rest} =
      case String.slice(D, 0, 4) do
        "do()" -> process_data(String.slice(D, 4..-1), part2, true)
        "don't()" -> process_data(String.slice(D, 7..-1), part2, false)
        "mul(" ->
          {new_part2, new_enabled, rest} = process_mul(D, part2, enabled)
          process_data(rest, new_part2, new_enabled)
        _ ->
          process_data(String.slice(D, 1..-1), part2, enabled)
      end

    new_part2
  end

  defp process_mul(D, part2, enabled) do
    # Find the closing parenthesis
    {mul_str, rest} = find_closing_paren(D, 4)

    # Extract numbers inside "mul(x, y)"
    case Regex.run(~r/(\d+),\s*(\d+)/, mul_str) do
      [_, x_str, y_str] ->
        x = String.to_integer(x_str)
        y = String.to_integer(y_str)

        # Make sure the closing parenthesis isn't followed by a number
        if rest != "" and !String.match?(String.slice(rest, 0, 1), ~r/\d/) do
          if enabled do
            part2 = part2 + x * y
          end
          IO.puts(mul_str)
        end

        {part2, enabled, rest}

      _ -> {part2, enabled, D}
    end
  end

  defp find_closing_paren(D, index) do
    # Find the closing parenthesis matching the opening "mul("
    {_, j} = Enum.reduce_while(index..String.length(D)-1, {0, index}, fn i, {depth, i_acc} ->
      if String.slice(D, i, i+1) == "(" do
        {:cont, {depth + 1, i+1}}
      else
        if String.slice(D, i, i+1) == ")" do
          if depth == 1 do
            {:halt, {i, String.slice(D, i+1..-1)}}
          else
            {:cont, {depth - 1, i+1}}
          end
        else
          {:cont, {depth, i+1}}
        end
      end
    end)
    {String.slice(D, index+1..j-1), j}
  end
end

# Run the program
Parser.process()
