defmodule MemoryGrid do
  def parse_input(file_name) do
    file_name
    |> File.read!()
    |> String.split("\n", trim: true)
    |> Enum.map(fn line ->
      line
      |> String.split(",")
      |> Enum.map(&String.to_integer/1)
    end)
  end

  defp is_within_bounds(x, y, size) do
    x >= 0 and x < size and y >= 0 and y < size
  end

  def bfs(grid, start, goal, size) do
    directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
    queue = :queue.new() |> :queue.in(start)
    visited = MapSet.new([start])

    bfs_loop(queue, visited, directions, grid, goal, size)
  end

  defp bfs_loop(queue, visited, _directions, grid, goal, _size) do
    case :queue.out(queue) do
      {:empty, _} -> false
      {{:value, current}, new_queue} ->
        if current == goal, do: true, else: bfs_step(new_queue, visited, grid, goal, current, size)
    end
  end

  defp bfs_step(queue, visited, grid, goal, current, size) do
    directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]

    Enum.reduce_while(directions, queue, fn direction, acc ->
      [dx, dy] = direction
      [x, y] = current
      nx = x + dx
      ny = y + dy

      if is_within_bounds(nx, ny, size) do
        neighbor = [nx, ny]
        if !MapSet.member?(visited, neighbor) and !Enum.member?(grid, neighbor) do
          visited = MapSet.put(visited, neighbor)
          {:cont, :queue.in(neighbor, acc)}
        else
          {:cont, acc}
        end
      else
        {:cont, acc}
      end
    end)
    |> case do
      {:cont, new_queue} -> bfs_loop(new_queue, visited, grid, goal, size)
      _ -> false
    end
  end

  def find_shortest_path(grid, start, goal, size) do
    directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
    queue = :queue.new() |> :queue.in({start, 0})
    visited = MapSet.new([start])

    find_shortest_path_loop(queue, visited, directions, grid, goal, size)
  end

  defp find_shortest_path_loop(queue, visited, directions, grid, goal, size) do
    case :queue.out(queue) do
      {:empty, _} -> nil
      {{:value, {{x, y}, steps}}, new_queue} ->
        if [x, y] == goal do
          steps
        else
          Enum.reduce_while(directions, new_queue, fn direction, acc ->
            [dx, dy] = direction
            nx = x + dx
            ny = y + dy
            if is_within_bounds(nx, ny, size) do
              neighbor = [nx, ny]
              if !MapSet.member?(visited, neighbor) and !Enum.member?(grid, neighbor) do
                visited = MapSet.put(visited, neighbor)
                {:cont, :queue.in({neighbor, steps + 1}, acc)}
              else
                {:cont, acc}
              end
            else
              {:cont, acc}
            end
          end)
          |> case do
            {:cont, new_queue} -> find_shortest_path_loop(new_queue, visited, directions, grid, goal, size)
            _ -> nil
          end
        end
    end
  end

  def solve do
    input_file = "input.txt"
    points = parse_input(input_file)
    size = 71
    start = [0, 0]
    goal = [70, 70]

    # Part 1: Shortest path after 1024 bytes
    grid_part1 = Enum.take(points, 1024) |> Enum.into(MapSet.new())
    case find_shortest_path(grid_part1, start, goal, size) do
      nil -> IO.puts("No valid path to the exit after 1024 bytes.")
      shortest_path -> IO.puts("Shortest Path after 1024 bytes: #{shortest_path}")
    end

    # Part 2: First byte that blocks the path
    grid_part2 = MapSet.new()
    Enum.each(points, fn point ->
      grid_part2 = MapSet.put(grid_part2, point)
      unless bfs(grid_part2, start, goal, size) do
        IO.puts("First blocking byte: #{Enum.join(point, ",")}")
        throw :done
      end
    end)
  end
end

MemoryGrid.solve()
