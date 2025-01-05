# def run_program(registers, program):
#     # Initialize registers
#     A, B, C = registers

#     # Initialize the instruction pointer
#     ip = 0

#     # Output list
#     output = []

#     # Operand resolution
#     def resolve_combo_operand(operand):
#         if operand <= 3:
#             return operand  # Literal values 0-3
#         elif operand == 4:
#             return A
#         elif operand == 5:
#             return B
#         elif operand == 6:
#             return C
#         else:
#             raise ValueError("Invalid combo operand: 7")

#     while ip < len(program):
#         # Read opcode and operand
#         opcode = program[ip]
#         operand = program[ip + 1]

#         if opcode == 0:  # adv
#             A //= 2 ** resolve_combo_operand(operand)
#         elif opcode == 1:  # bxl
#             B ^= operand
#         elif opcode == 2:  # bst
#             B = resolve_combo_operand(operand) % 8
#         elif opcode == 3:  # jnz
#             if A != 0:
#                 ip = operand
#                 continue
#         elif opcode == 4:  # bxc
#             B ^= C
#         elif opcode == 5:  # out
#             output.append(resolve_combo_operand(operand) % 8)
#         elif opcode == 6:  # bdv
#             B = A // (2 ** resolve_combo_operand(operand))
#         elif opcode == 7:  # cdv
#             C = A // (2 ** resolve_combo_operand(operand))
#         else:
#             raise ValueError(f"Invalid opcode: {opcode}")

#         # Move to the next instruction
#         ip += 2

#     return output

# if __name__ == "__main__":
#     # Input parameters
#     program = [2, 4, 1, 1, 7, 5, 4, 6, 0, 3, 1, 4, 5, 5, 3, 0]  # Input program

#     # Debug flag
#     debug = True

#     # Find the smallest positive A that reproduces the program
#     for A in range(1, 10**7):  # Set an upper limit to avoid infinite loops
#         registers = [A, 0, 0]  # Initialize registers
#         output = run_program(registers, program)

#         if debug:
#             print(f"Testing A={A}: Output={output}")

#         if output == program:
#             print(f"The smallest positive value for Register A is: {A}")
#             break
#     else:
#         print("No valid value for Register A found within the tested range.")

def run_program(registers, program):
    A, B, C = registers
    ip = 0
    output = []

    def resolve_combo_operand(operand):
        if operand <= 3:
            return operand
        elif operand == 4:
            return A
        elif operand == 5:
            return B
        elif operand == 6:
            return C
        else:
            raise ValueError("Invalid combo operand")

    while ip < len(program):
        opcode = program[ip]
        operand = program[ip + 1]

        if opcode == 0:  # adv
            A //= 2 ** resolve_combo_operand(operand)
        elif opcode == 1:  # bxl
            B ^= operand
        elif opcode == 2:  # bst
            B = resolve_combo_operand(operand) % 8
        elif opcode == 3:  # jnz
            if A != 0:
                ip += operand * 2 # Corrected jnz: relative jump, *2 for instruction size
                continue # Important to skip the ip += 2 below
        elif opcode == 4:  # bxc
            B ^= C
        elif opcode == 5:  # out
            output.append(resolve_combo_operand(operand) % 8)
        elif opcode == 6:  # bdv
            B = A // (2 ** resolve_combo_operand(operand))
        elif opcode == 7:  # cdv
            C = A // (2 ** resolve_combo_operand(operand))
        else:
            raise ValueError(f"Invalid opcode: {opcode}")

        ip += 2

    return output

if __name__ == "__main__":
    program = [2, 4, 1, 1, 7, 5, 4, 6, 0, 3, 1, 4, 5, 5, 3, 0]
    target_output = [2, 1, 7, 4, 0, 1, 5, 3] # The expected output given in the prompt.

    for A in range(1, 1000): # Reduced range for faster testing
        registers = [A, 0, 0]
        output = run_program(registers, program)

        if output == target_output:
            print(f"The value for Register A is: {A}")
            break
    else:
        print("No valid value for Register A found within the tested range.")