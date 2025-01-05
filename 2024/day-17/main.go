package main

import (
        "fmt"
        "math"
        "strconv"
        "strings"
)

func main() {
        input := getTestInput()
        part1Result := part1(input)
        fmt.Println("Part 1:", part1Result)

        part2Result := part2(input)
        fmt.Println("Part 2:", part2Result)
}

func part1(input string) int64 {
        sections := strings.Split(input, "\n\n")
        regs := sections[0]
        program := parseProgram(sections[1])

        A := parseRegisterA(regs)
        executionResult := execProgram(A, program)

        // Convert the execution result to a single integer
        resultStr := ""
        for _, val := range executionResult {
                resultStr += strconv.FormatInt(val, 10)
        }
        result, _ := strconv.ParseInt(resultStr, 10, 64)
        return result
}

var success int64 = math.MaxInt64

func part2(input string) int64 {
        sections := strings.Split(input, "\n\n")
        program := parseProgram(sections[1])
        dfs(program, 0, 0)
        return success
}

func dfs(program []int64, cur int64, pos int) {
        if pos >= len(program) {
                if cur < success {
                        success = cur
                }
                return
        }

        for i := int64(0); i < 8; i++ {
                nextNum := (cur << 3) + i
                execResult := execProgram(nextNum, program)
                endIdx := len(program) - pos - 1

                if equal(execResult, program[endIdx:]) {
                        dfs(program, nextNum, pos+1)
                }
        }
}

func execProgram(A int64, program []int64) []int64 {
        pointer := 0
        output := []int64{}
        var B, C int64

        for pointer >= 0 && pointer < len(program) {
                opcode := program[pointer]
                literalOperand := program[pointer+1]

                var combo int64
                if literalOperand < 4 {
                        combo = literalOperand
                } else if literalOperand == 4 {
                        combo = A
                } else if literalOperand == 5 {
                        combo = B
                } else {
                        combo = C
                }

                switch opcode {
                case 0:
                        A = A / int64(math.Pow(2, float64(combo)))
                        pointer += 2
                case 1:
                        B ^= literalOperand
                        pointer += 2
                case 2:
                        B = combo % 8
                        pointer += 2
                case 3:
                        if A != 0 {
                                pointer = int(literalOperand)
                        } else {
                                pointer += 2
                        }
                case 4:
                        B ^= C
                        pointer += 2
                case 5:
                        output = append(output, combo%8)
                        pointer += 2
                case 6:
                        B = A / int64(math.Pow(2, float64(combo)))
                        pointer += 2
                case 7:
                        C = A / int64(math.Pow(2, float64(combo)))
                        pointer += 2
                }
        }
        return output
}

// Helper Functions

func parseRegisterA(regs string) int64 {
        lines := strings.Split(regs, "\n")
        for _, line := range lines {
                if strings.Contains(line, "Register A:") {
                        parts := strings.Split(line, ":")
                        val, _ := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
                        return val
                }
        }
        return 0
}

func parseProgram(programStr string) []int64 {
        programStr = strings.ReplaceAll(programStr, "Program: ", "")
        programParts := strings.Split(programStr, ",")
        program := make([]int64, len(programParts))
        for i, part := range programParts {
                val, _ := strconv.ParseInt(strings.TrimSpace(part), 10, 64)
                program[i] = val
        }
        return program
}

func equal(a, b []int64) bool {
        if len(a) != len(b) {
                return false
        }
        for i := range a {
                if a[i] != b[i] {
                        return false
                }
        }
        return true
}

// Input Functions

func getTestInput() string {
        return `Register A: 28066687
Register B: 0
Register C: 0

Program: 2,4,1,1,7,5,4,6,0,3,1,4,5,5,3,0`
}
