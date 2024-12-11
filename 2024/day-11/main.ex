defmodule PlutonianPebbles do
  def solve(x, t, dp) do
    key = {x, t}

    if Map.has_key?(dp, key) do
      Map.get(dp, key)
    else
      result =
        cond do
          t == 0 ->
            1

          x == 0 ->
            solve(1, t - 1, dp)

          rem(String.length(Integer.to_string(x)), 2) == 0 ->
            dstr = Integer.to_string(x)
            mid = div(String.length(dstr), 2)
            left = String.slice(dstr, 0, mid) |> String.to_integer()
            right = String.slice(dstr, mid, String.length(dstr)) |> String.to_integer()
            solve(left, t - 1, dp) + solve(right, t - 1, dp)

          true ->
            solve(x * 2024, t - 1, dp)
        end

      dp = Map.put(dp, key, result)
      result
    end
  end

  def solve_all(d, t) do
    Enum.reduce(d, 0, fn x, sum ->
      dp = %{}
      sum + solve(x, t, dp)
    end)
  end

  def main() do
    case File.read("input.txt") do
      {:ok, data} ->
        d = data |> String.trim() |> String.split(~r/\s+/) |> Enum.map(&String.to_integer/1)

        dp = %{}
        result25 = solve_all(d, 25)
        result75 = solve_all(d, 75)

        IO.puts("Result after 25 blinks: #{result25}")
        IO.puts("Result after 75 blinks: #{result75}")

      {:error, reason} ->
        IO.puts("Error reading file: #{reason}")
    end
  end
end

PlutonianPebbles.main()
