#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
    FILE *file = fopen("input.txt", "r");
    if (file == NULL) {
        printf("Could not open file\n");
        return 1;
    }

    char line[1024]; // Buffer to hold each line
    int totalSum = 0;

    // Read file line by line
    while (fgets(line, sizeof(line), file)) {
        char *ptr = line;

        // Search for "mul("
        while ((ptr = strstr(ptr, "mul(")) != NULL) {
            ptr += 4; // Move pointer past "mul("

            int x, y;
            // Extract numbers inside the parentheses
            if (sscanf(ptr, "%d, %d", &x, &y) == 2) {
                totalSum += x * y;
            }

            // Move pointer past the closing parenthesis
            ptr = strchr(ptr, ')');
            if (ptr != NULL) {
                ptr++;  // Move past ')'
            }
        }
    }

    printf("Total sum of mul(x, y): %d\n", totalSum);

    fclose(file);
    return 0;
}
