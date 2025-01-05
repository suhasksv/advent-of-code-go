def run_program(registers, program):
    # Initialize registers
    A, B, C = registers

    # Initialize the instruction pointer
    ip = 0

    # Output list
    output = []

    # Operand resolution
    def resolve_combo_operand(operand):
        if operand <= 3:
            return operand  # Literal values 0-3
        elif operand == 4:
            return A
        elif operand == 5:
            return B
        elif operand == 6:
            return C
        else:
            raise ValueError("Invalid combo operand: 7")

    while ip < len(program):
        # Read opcode and operand
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
                ip = operand
                continue
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

        # Move to the next instruction
        ip += 2

    return output

if __name__ == "__main__":
    # Example input values
    registers = [28066687, 0, 0]  # Replace with desired initial register values
    program = [2,4,1,1,7,5,4,6,0,3,1,4,5,5,3,0]  # Replace with desired program

    # Run the program
    output = run_program(registers, program)

    # Format the output
    result = ",".join(map(str, output))
    print(result)
