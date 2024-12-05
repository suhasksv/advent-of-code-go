#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#define MAX_LINE_LENGTH 1024

// Function to read the file and return the content as a string
char* read_file(const char *filename) {
    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        fprintf(stderr, "Error reading file: %s\n", filename);
        exit(1);
    }

    fseek(file, 0, SEEK_END);
    long file_size = ftell(file);
    fseek(file, 0, SEEK_SET);

    char *content = (char*)malloc(file_size + 1);
    if (content == NULL) {
        fprintf(stderr, "Memory allocation failed.\n");
        exit(1);
    }

    fread(content, 1, file_size, file);
    content[file_size] = '\0';

    fclose(file);
    return content;
}

// Function to process "mul(x, y)" operations
int process_mul(const char *data, int *index, int part2, int enabled) {
    int j = *index + 4;

    // Find the closing parenthesis
    while (data[j] != ')') {
        j++;
    }

    // Extract the substring for "mul(x, y)"
    char mul_str[j - (*index + 4) + 1];
    strncpy(mul_str, data + (*index + 4), j - (*index + 4));
    mul_str[j - (*index + 4)] = '\0';

    // Extract x and y using a simple manual split
    int x, y;
    if (sscanf(mul_str, "%d, %d", &x, &y) == 2) {
        // Make sure the closing parenthesis is not followed by a number
        if (!isdigit(data[j + 1])) {
            if (enabled) {
                part2 += x * y;
            }
            printf("%.*s\n", j + 1, data); // Print the "mul(x, y)" part
        }
    }

    *index = j + 1; // Update the index to move past the closing parenthesis
    return part2;
}

int process_data(const char *data, int part2, int enabled) {
    int i = 0;
    while (data[i] != '\0') {
        if (strncmp(&data[i], "do()", 4) == 0) {
            enabled = 1;
            i += 4;
        } else if (strncmp(&data[i], "don't()", 7) == 0) {
            enabled = 0;
            i += 7;
        } else if (strncmp(&data[i], "mul(", 4) == 0) {
            part2 = process_mul(data, &i, part2, enabled);
        } else {
            i++;
        }
    }
    return part2;
}

int main(int argc, char *argv[]) {
    // Set the file to read input from
    const char *file = (argc > 1) ? argv[1] : "inp
