from math import gcd
from sympy import symbols, Eq, solve

def solve_claw_machine(a_x, a_y, b_x, b_y, x_prize, y_prize):
    """
    Solve for the minimum cost to win the prize given button behaviors and prize position.
    
    :param a_x: X movement for button A
    :param a_y: Y movement for button A
    :param b_x: X movement for button B
    :param b_y: Y movement for button B
    :param x_prize: Prize X position
    :param y_prize: Prize Y position
    :return: Minimum cost or None if unsolvable
    """
    g = gcd(a_x * b_y - a_y * b_x, gcd(a_x, b_x))

    # Check if the prize position is reachable
    if x_prize % g != 0 or y_prize % g != 0:
        return None

    # Define the equations for X and Y
    n_A, n_B = symbols('n_A n_B', integer=True, nonnegative=True)
    eq1 = Eq(n_A * a_x + n_B * b_x, x_prize)
    eq2 = Eq(n_A * a_y + n_B * b_y, y_prize)

    # Solve the system of equations
    solutions = solve((eq1, eq2), (n_A, n_B), dict=True)
    
    min_cost = float('inf')

    # Evaluate solutions to find the minimum cost
    for sol in solutions:
        n_A_val = sol[n_A]
        n_B_val = sol[n_B]

        if n_A_val >= 0 and n_B_val >= 0:  # Ensure non-negative solutions
            cost = 3 * n_A_val + 1 * n_B_val
            min_cost = min(min_cost, cost)

    return min_cost if min_cost != float('inf') else None

# Example updated prize positions
machines = [
    ((94, 34, 22, 67), (10000000008400, 10000000005400)),
    ((26, 66, 67, 21), (10000000012748, 10000000012176)),
    ((17, 86, 84, 37), (10000000007870, 10000000006450)),
    ((69, 23, 27, 71), (10000000018641, 10000000010279)),
]

results = []

for i, ((a_x, a_y, b_x, b_y), (x_prize, y_prize)) in enumerate(machines):
    result = solve_claw_machine(a_x, a_y, b_x, b_y, x_prize, y_prize)
    results.append((i + 1, result))

# Output results
for machine, cost in results:
    if cost:
        print(f"Machine {machine}: Minimum cost = {cost}")
    else:
        print(f"Machine {machine}: Prize unreachable")

