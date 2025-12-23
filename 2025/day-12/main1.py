import sys

# --- Helper Classes ---

class Shape:
    def __init__(self, id, coords):
        self.id = id
        self.coords = coords # Set of (r, c) tuples
        self.variations = self._generate_variations()
        self.area = len(coords)

    def _normalize(self, coords):
        # Shift coords so the top-leftmost is at (0,0)
        min_r = min(r for r, c in coords)
        min_c = min(c for r, c in coords)
        return tuple(sorted((r - min_r, c - min_c) for r, c in coords))

    def _generate_variations(self):
        # Generate all 8 symmetries (rotations + flips)
        variations = set()
        current = self.coords

        # Try all 4 rotations
        for _ in range(4):
            # Add normal
            variations.add(self._normalize(current))
            # Add flipped (Horizontal flip)
            flipped = set((r, -c) for r, c in current)
            variations.add(self._normalize(flipped))

            # Rotate 90 degrees: (r, c) -> (c, -r)
            current = set((c, -r) for r, c in current)

        return list(variations)

def parse_input(filename):
    with open(filename, 'r') as f:
        lines = [l.rstrip() for l in f]

    shapes = []
    queries = []

    # 1. Parse Shapes
    # We assume shapes come first, then a blank line or structure change
    # But the format is mixed. We'll iterate and detect.

    current_shape_id = None
    current_shape_lines = []

    i = 0
    while i < len(lines):
        line = lines[i]
        if not line:
            i += 1
            continue

        # Check if it's a query (contains "x" and numbers, e.g., "12x5: 1 0...")
        if 'x' in line and ':' in line and line[0].isdigit() and not line.strip().endswith(':'):
            # This is a query line
            queries.append(line)
            i += 1
            continue

        # Check if it's a shape header (e.g., "0:")
        if line.endswith(':'):
            # Save previous shape if exists
            if current_shape_id is not None:
                shapes.append(Shape(current_shape_id, _parse_grid(current_shape_lines)))
                current_shape_lines = []

            current_shape_id = int(line[:-1])
            i += 1
            continue

        # Otherwise it's part of a shape grid or empty
        if current_shape_id is not None:
            current_shape_lines.append(line)

        i += 1

    # Add the last shape
    if current_shape_id is not None and current_shape_lines:
        shapes.append(Shape(current_shape_id, _parse_grid(current_shape_lines)))

    # Sort shapes by ID just in case
    shapes.sort(key=lambda s: s.id)
    return shapes, queries

def _parse_grid(lines):
    coords = set()
    for r, row in enumerate(lines):
        for c, char in enumerate(row):
            if char == '#':
                coords.add((r, c))
    return coords

# --- Backtracking Solver with Bitmasks ---

def solve_query(query_line, shapes):
    # Parse "12x5: 1 0 1 0 2 2"
    left, right = query_line.split(':')
    W, H = map(int, left.split('x'))
    counts = list(map(int, right.strip().split()))

    # Build list of pieces to place: [ShapeObj, ShapeObj, ...]
    pieces_to_place = []
    total_area = 0

    for s_idx, count in enumerate(counts):
        shape = shapes[s_idx]
        for _ in range(count):
            pieces_to_place.append(shape)
            total_area += shape.area

    # Optimization 1: Area Check
    if total_area > W * H:
        return False

    # Optimization 2: Sort pieces by size (Largest first helps fail faster)
    pieces_to_place.sort(key=lambda s: s.area, reverse=True)

    # Precompute Bitmasks for every shape variation at every valid position
    # mask_cache[shape_id] = [ (mask_int, variation_index?), ... ]
    # Actually just a list of valid integer masks is enough.

    # To handle identical pieces, we group masks by shape ID.
    # shape_masks[shape_id] -> list of integers representing placement on WxH grid
    shape_masks = {}

    # Get unique shape IDs involved
    unique_ids = set(p.id for p in pieces_to_place)

    for uid in unique_ids:
        s = shapes[uid]
        masks = []
        for var_coords in s.variations:
            # Try to place this variation at every (r, c)
            # Find bounds of the variation
            max_r = max(r for r, c in var_coords)
            max_c = max(c for r, c in var_coords)

            # We can slide the top-left (0,0) to any (dr, dc)
            # such that (max_r + dr) < H and (max_c + dc) < W
            for dr in range(H - max_r):
                for dc in range(W - max_c):
                    # Create bitmask
                    mask = 0
                    for (r, c) in var_coords:
                        # Map 2D (r+dr, c+dc) to 1D bit index
                        # Index = (r+dr)*W + (c+dc)
                        bit_index = (r + dr) * W + (c + dc)
                        mask |= (1 << bit_index)
                    masks.append(mask)
        # Remove duplicates (e.g., symmetrical shapes producing same masks) and sort descending (heuristic)
        # Sorting masks by value isn't strictly necessary but can help determinism
        shape_masks[uid] = sorted(list(set(masks)), reverse=True)

    # Recursive Solver
    # grid_state is an integer

    # We use a stack instead of recursion to avoid limit issues,
    # but recursion is cleaner. Let's use recursion with indices.

    # Piece indices to solve: 0 to len(pieces_to_place)-1
    N = len(pieces_to_place)

    def backtrack(k, current_grid, start_mask_idx):
        if k == N:
            return True

        piece = pieces_to_place[k]
        masks = shape_masks[piece.id]

        # Optimization 3: Symmetry Breaking
        # If this piece is same as previous, we must pick a mask
        # that appears *after* (or at) the previous mask in our list?
        # Actually, simpler: if pieces are identical, enforce that the *position* (mask value)
        # is strictly less (or greater) to avoid permutations.
        # Since we sort masks descending, we can enforce: current_mask < previous_mask
        # or just pass the index in the mask list to start from.

        # Logic: If pieces_to_place[k] == pieces_to_place[k-1], start loop from 'start_mask_idx'
        # Else start from 0.

        # Note: We rely on the masks list being constant.

        for i in range(start_mask_idx, len(masks)):
            m = masks[i]

            # Check collision
            if (current_grid & m) == 0:
                # Place it
                new_grid = current_grid | m

                # Determine constraints for next step
                next_start = 0
                if k + 1 < N and pieces_to_place[k+1].id == piece.id:
                    # Next piece is identical, so restrict it to start searching
                    # AFTER the current mask index to avoid duplicate states.
                    # We allow i (start from same) only if shapes can't overlap (which they can't).
                    # But since they can't overlap, m_next != m_current.
                    # So we can start from i + 1.
                    next_start = i + 1

                if backtrack(k + 1, new_grid, next_start):
                    return True

        return False

    return backtrack(0, 0, 0)

# --- Main Execution ---

def main():
    input_file = 'input.txt' # Change if needed
    try:
        shapes, queries = parse_input(input_file)
    except FileNotFoundError:
        print("Please save your puzzle input as 'input.txt'")
        return

    success_count = 0
    print(f"Loaded {len(shapes)} shapes and {len(queries)} regions.")

    for q in queries:
        if solve_query(q, shapes):
            success_count += 1
            # print(f"Region {q}: Fits!")
        else:
            # print(f"Region {q}: Impossible.")
            pass

    print(f"\nTotal regions that can fit all presents: {success_count}")

if __name__ == "__main__":
    main()