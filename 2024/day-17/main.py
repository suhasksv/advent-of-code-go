# Not Working

# def run_program(program, a, b, c):
#     registers = {"A": a, "B": b, "C": c}
#     instruction_pointer = 0
#     output = []

#     while instruction_pointer < len(program):
#         opcode = program[instruction_pointer]
#         operand = program[instruction_pointer + 1] if instruction_pointer + 1 < len(program) else None

#         if operand is None:
#             break

#         def get_operand_value(op):
#             if op < 4:
#                 return op
#             else:
#                 return registers["ABC"[op % 3]]

#         if opcode == 0:  # adv
#             divisor = 2 ** get_operand_value(operand)
#             registers["A"] //= divisor
#         elif opcode == 1:  # bxl
#             registers["B"] ^= operand
#         elif opcode == 2:  # bst
#             registers["B"] = get_operand_value(operand) % 8
#         elif opcode == 3:  # jnz
#             if registers["A"] != 0:
#                 instruction_pointer = operand
#                 continue
#         elif opcode == 4:  # bxc
#             registers["B"] ^= registers["C"]
#         elif opcode == 5:  # out
#             output.append(get_operand_value(operand) % 8)
#         elif opcode == 6:  # bdv
#             divisor = 2 ** get_operand_value(operand)
#             registers["B"] //= divisor
#         elif opcode == 7:  # cdv
#             divisor = 2 ** get_operand_value(operand)
#             registers["C"] //= divisor

#         instruction_pointer += 2

#     return output

# def solve_part1(program_str, a, b, c):
#     program = list(map(int, program_str.split(",")))
#     output = run_program(program, a, b, c)
#     return ",".join(map(str, output))

# def solve_part2(program_str, b, c, max_iterations=10000000): #Added max iterations
#     program = list(map(int, program_str.split(",")))
#     program_len = len(program)

#     a = 1
#     iterations = 0
#     while iterations < max_iterations: #Added check for iterations
#         output = run_program(program, a, b, c)
#         if len(output) > program_len:
#             a += 1
#             iterations += 1
#             continue
#         if output == program:
#             return a
#         a += 1
#         iterations += 1
#     return None # Return None if no solution is found within max_iterations

# # Your Input (Corrected Part 1):
# program_str = "2,4,1,1,7,5,4,6,0,3,1,4,5,5,3,0"
# a = 28066687
# b = 0
# c = 0
# result1 = solve_part1(program_str, a, b, c)
# print(f"Part 1 Puzzle: {result1}") # Correct Output for your input

# # Part 2 with your input
# program_str = "2,4,1,1,7,5,4,6,0,3,1,4,5,5,3,0"
# b = 0
# c = 0
# result2 = solve_part2(program_str, b, c)
# print(f"Part 2 Puzzle: {result2}")


def simulate_computer(program, initial_a):
    registers = {'A': initial_a, 'B': 0, 'C': 0}
    ip = 0

    def get_value(operand):
        if operand < 4:
            return operand
        return registers[f'A{operand - 4}']  # Corrected: Access register A, B, or C

    while ip < len(program):
        opcode, operand = program[ip:ip+2]
        ip += 2

        if opcode == 0:  # adv
            registers['A'] //= 2 ** get_value(operand)
        elif opcode == 1:  # bxl
            registers['B'] ^= operand
        elif opcode == 2:  # bst
            registers['B'] = operand % 8
        elif opcode == 3:  # jnz
            if registers['A'] != 0:
                ip = operand
        elif opcode == 4:  # bxc
            registers['B'] ^= registers['C']
        elif opcode == 5:  # out
            print(operand % 8, end='')
        elif opcode == 6:  # bdv
            registers['B'] //= 2 ** get_value(operand)
        elif opcode == 7:  # cdv
            registers['C'] //= 2 ** get_value(operand)

    return ''.join(str(x) for x in output)

def find_initial_value(program):
    initial_value = 1
    while True:
        output = simulate_computer(program, initial_value)
        if output == ''.join(str(x) for x in program):
            return initial_value
        initial_value += 1

# Replace with your actual program input
program = [2,4,1,1,7,5,4,6,0,3,1,4,5,5,3,0]

result = find_initial_value(program)
print(result)