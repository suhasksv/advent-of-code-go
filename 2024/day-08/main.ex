defmodule GridAntenna do
  def manhattan_distance({r1, c1}, {r2, c2}) do
    abs(r1 - r2) + abs(c1 - c2)
  end

  def main() do
    infile = System.argv() |> Enum.at(0, "input.txt")
    grid = File.read!(infile) |> String.split("\n", trim: true)

    positions =
      Enum.with_index(grid)
      |> Enum.reduce(%{}, fn {row, r}, acc ->
        row
        |> String.graphemes()
        |> Enum.with_index()
        |> Enum.reduce(acc, fn {ch, c}, acc ->
          if ch != "." do
            Map.update(acc, ch, [{r, c}], &[{r, c} | &1])
          else
            acc
          end
        end)
      end)

    {a1, a2} =
      Enum.reduce(0..(length(grid) - 1), {MapSet.new(), MapSet.new()}, fn r, {a1, a2} ->
        Enum.reduce(0..(String.length(hd(grid)) - 1), {a1, a2}, fn c, {a1, a2} ->
          Enum.reduce(positions, {a1, a2}, fn {_k, pos}, {a1, a2} ->
            Enum.reduce(pos, {a1, a2}, fn {r1, c1}, {a1, a2} ->
              Enum.reduce(pos, {a1, a2}, fn {r2, c2}, {a1, a2} ->
                if {r1, c1} != {r2, c2} do
                  d1 = manhattan_distance({r, c}, {r1, c1})
                  d2 = manhattan_distance({r, c}, {r2, c2})
                  dr1 = r - r1
                  dr2 = r - r2
                  dc1 = c - c1
                  dc2 = c - c2

                  a1 =
                    if (d1 == 2 * d2 or d1 * 2 == d2) and dr1 * dc2 == dc1 * dr2,
                       do: MapSet.put(a1, {r, c}),
                       else: a1

                  a2 =
                    if dr1 * dc2 == dc1 * dr2,
                       do: MapSet.put(a2, {r, c}),
                       else: a2

                  {a1, a2}
                else
                  {a1, a2}
                end
              end)
            end)
          end)
        end)
      end)

    IO.puts(MapSet.size(a1))
    IO.puts(MapSet.size(a2))
  end
end

GridAntenna.main()
