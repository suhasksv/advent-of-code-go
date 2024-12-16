defmodule GardenFencing do
  def read_input_file(filename) do
    File.read!(filename)
    |> String.split("\n", trim: true)
    |> Enum.map(&String.graphemes/1)
  end

  def calculate_fencing_price(map_input) do
    rows = length(map_input)
    cols = length(hd(map_input))

    visited = for _ <- 1..rows, do: Enum.map(1..cols, fn _ -> false end)

    is_valid = fn x, y ->
      x >= 0 and x < rows and y >= 0 and y < cols
    end

    bfs = fn start_x, start_y ->
      queue = [{start_x, start_y}]
      region_cells = []
      plant_type = Enum.at(Enum.at(map_input, start_x), start_y)

      put_in(visited[start_x][start_y], true)

      while queue != [] do
        [{current_x, current_y} | rest] = queue
        queue = rest

        region_cells = [{current_x, current_y} | region_cells]

        for {dx, dy} <- [{-1, 0}, {1, 0}, {0, -1}, {0, 1}] do
          nx = current_x + dx
          ny = current_y + dy

          if is_valid.(nx, ny) and not Enum.at(Enum.at(visited, nx), ny) and Enum.at(Enum.at(map_input, nx), ny) == plant_type do
            queue = [{nx, ny} | queue]
            put_in(visited[nx][ny], true)
          end
        end
      end

      region_cells
    end

    calculate_area_and_perimeter = fn region_cells ->
      area = length(region_cells)
      perimeter = 0

      for {x, y} <- region_cells do
        for {dx, dy} <- [{-1, 0}, {1, 0}, {0, -1}, {0, 1}] do
          nx = x + dx
          ny = y + dy

          if not is_valid.(nx, ny) or Enum.at(Enum.at(map_input, nx), ny) != Enum.at(Enum.at(map_input, x), y) do
            perimeter = perimeter + 1
          end
        end
      end

      {area, perimeter}
    end

    total_price = 0

    for i <- 0..(rows - 1) do
      for j <- 0..(cols - 1) do
        if not Enum.at(Enum.at(visited, i), j) do
          region_cells = bfs.(i, j)
          {area, perimeter} = calculate_area_and_perimeter.(region_cells)
          total_price = total_price + area * perimeter
        end
      end
    end

    total_price
  end

  def main() do
    map_input = read_input_file("input.txt")
    total_price = calculate_fencing_price(map_input)
    IO.puts("Total price of fencing: #{total_price}")
  end
end

GardenFencing.main()

