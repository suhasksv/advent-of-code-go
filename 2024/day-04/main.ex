defmodule XmasPatterns do
  # Function to count patterns in the grid
  def count_patterns(grid) do
    rows = length(grid)
    cols = String.length(Enum.at(grid, 0))

    {p1, p2} =
      for r <- 0..(rows - 1), c <- 0..(cols - 1), reduce: {0, 0} do
        {p1, p2} ->
          # Check for "XMAS" horizontally, vertically, and diagonally
          p1 = p1 +
            if c + 3 < cols and match_pattern(grid, r, c, ["X", "M", "A", "S"], :horizontal), do: 1, else: 0
          p1 = p1 +
            if r + 3 < rows and match_pattern(grid, r, c, ["X", "M", "A", "S"], :vertical), do: 1, else: 0
          p1 = p1 +
            if r + 3 < rows and c + 3 < cols and match_pattern(grid, r, c, ["X", "M", "A", "S"], :diagonal_down), do: 1, else: 0

          # Check for "SAMX" horizontally, vertically, and diagonally
          p1 = p1 +
            if c + 3 < cols and match_pattern(grid, r, c, ["S", "A", "M", "X"], :horizontal), do: 1, else: 0
          p1 = p1 +
            if r + 3 < rows and match_pattern(grid, r, c, ["S", "A", "M", "X"], :vertical), do: 1, else: 0
          p1 = p1 +
            if r + 3 < rows and c + 3 < cols and match_pattern(grid, r, c, ["S", "A", "M", "X"], :diagonal_down), do: 1, else: 0
          p1 = p1 +
            if r - 3 >= 0 and c + 3 < cols and match_pattern(grid, r, c, ["S", "A", "M", "X"], :diagonal_up), do: 1, else: 0
          p1 = p1 +
            if r - 3 >= 0 and c + 3 < cols and match_pattern(grid, r, c, ["X", "M", "A", "S"], :diagonal_up), do: 1, else: 0

          # Check for "MAS" patterns surrounded by "M/S"
          p2 = p2 +
            if r + 2 < rows and c + 2 < cols and surrounded_pattern(grid, r, c, ["M", "A", "S"], "M", "S"), do: 1, else: 0
          p2 = p2 +
            if r + 2 < rows and c + 2 < cols and surrounded_pattern(grid, r, c, ["M", "A", "S"], "S", "M"), do: 1, else: 0
          p2 = p2 +
            if r + 2 < rows and c + 2 < cols and surrounded_pattern(grid, r, c, ["S", "A", "M"], "M", "S"), do: 1, else: 0
          p2 = p2 +
            if r + 2 < rows and c + 2 < cols and surrounded_pattern(grid, r, c, ["S", "A", "M"], "S", "M"), do: 1, else: 0

          {p1, p2}
      end

    {p1, p2}
  end

  # Helper function to match a pattern in the grid
  defp match_pattern(grid, r, c, pattern, direction) do
    Enum.with_index(pattern)
    |> Enum.all?(fn {char, i} ->
      case direction do
        :horizontal -> Enum.at(grid, r) |> String.at(c + i) == char
        :vertical -> Enum.at(grid, r + i) |> String.at(c) == char
        :diagonal_down -> Enum.at(grid, r + i) |> String.at(c + i) == char
        :diagonal_up -> Enum.at(grid, r - i) |> String.at(c + i) == char
      end
    end)
  end

  # Helper function to match a surrounded pattern
  defp surrounded_pattern(grid, r, c, pattern, top_left, bottom_right) do
    match_pattern(grid, r, c, pattern, :diagonal_down) and
    Enum.at(grid, r + 2) |> String.at(c) == top_left and
    Enum.at(grid, r) |> String.at(c + 2) == bottom_right
  end

  # Main function to read the grid and compute results
  def main(args) do
    infile = hd(args) || "input.txt"
    grid = File.read!(infile) |> String.split("\n", trim: true)
    {p1, p2} = count_patterns(grid)
    IO.puts(p1)
    IO.puts(p2)
  end
end

# Run the main function
XmasPatterns.main(System.argv())
