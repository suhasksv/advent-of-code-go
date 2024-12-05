defmodule MulParser do
  def parse_file(file_name) do
    # Read the file content
    {:ok, file_content} = File.read(file_name)

    # Regex to find mul(x, y) patterns
    regex = ~r/mul\((\d+), (\d+)\)/

      # Find all matches
    mul_matches = Regex.scan(regex, file_content)

    # Sum up the products of x and y
    total_sum = Enum.reduce(mul_matches, 0, fn [_, x_str, y_str], acc ->
      # Convert the captured values to integers
      x = String.to_integer(x_str)
      y = String.to_integer(y_str)
      # Add the product to the accumulator
      acc + (x * y)
    end)

    IO.puts("Total sum of mul(x, y): #{total_sum}")
  end
end

# Run the parser (replace 'input.txt' with the correct file path if needed)
MulParser.parse_file("input.txt")
