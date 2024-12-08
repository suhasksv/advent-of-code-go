defmodule GuardPatrol do
  def parse_input(filename) do
    File.read!(filename)
    |> String.split("\n", trim: true)
    |> Enum.map(&String.graphemes/1)
  end

  def find_guard(grid) do
    directions = ["^", ">", "v", "<"]

    Enum.find_value(Enum.with_index(grid), fn {row, r} ->
      Enum.find_value(Enum.with_index(row), fn {cell, c} ->
        case Enum.find_index(directions, &(&1 == cell)) do
          nil -> nil
          dir -> {r, c, dir}
        end
      end)
    end)
  end

  def simulate_patrol(grid) do
    directions = [{-1, 0}, {0, 1}, {1, 0}, {0, -1}]
    {r, c, dir} = find_guard(grid)

    rows = length(grid)
    cols = length(List.first(grid))
    visited = MapSet.new()

    simulate(r, c, dir, grid, directions, rows, cols, visited)
  end

  defp simulate(r, c, dir, grid, directions, rows, cols, visited) when r < 0 or r >= rows or c < 0 or c >= cols do
    MapSet.size(visited)
  end

  defp simulate(r, c, dir, grid, directions, rows, cols, visited) do
    visited = MapSet.put(visited, {r, c})

    {dr, dc} = Enum.at(directions, dir)
    nr = r + dr
    nc = c + dc

    if nr >= 0 and nr < rows and nc >= 0 and nc < cols and Enum.at(Enum.at(grid, nr), nc) == "#" do
      simulate(r, c, rem(dir + 1, 4), grid, directions, rows, cols, visited)
    else
      simulate(nr, nc, dir, grid, directions, rows, cols, visited)
    end
  end
end

grid = GuardPatrol.parse_input("input.txt")
IO.puts("Distinct positions visited: #{GuardPatrol.simulate_patrol(grid)}")
