#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

#define MAX_REPORT_LENGTH 1000

// Function to check if a report is safe
bool is_safe(int *report, int size) {
    bool all_increasing = true, all_decreasing = true;

    for (int i = 0; i < size - 1; i++) {
        int diff = report[i + 1] - report[i];
        if (abs(diff) < 1 || abs(diff) > 3) {
            return false;
        }
        if (diff <= 0) all_increasing = false;
        if (diff >= 0) all_decreasing = false;
    }

    return all_increasing || all_decreasing;
}

// Function to check if a report can be made safe by removing one level
bool can_be_made_safe(int *report, int size) {
    for (int i = 0; i < size; i++) {
        int modified[MAX_REPORT_LENGTH];
        int k = 0;

        // Create a modified report without the i-th level
        for (int j = 0; j < size; j++) {
            if (j != i) {
                modified[k++] = report[j];
            }
        }

        if (is_safe(modified, size - 1)) {
            return true;
        }
    }
    return false;
}

// Function to count safe reports
int count_safe_reports(const char *file_path, int part) {
    FILE *file = fopen(file_path, "r");
    if (!file) {
        fprintf(stderr, "Error: Could not open file %s\n", file_path);
        return -1;
    }

    char line[1024];
    int safe_count = 0;

    while (fgets(line, sizeof(line), file)) {
        if (line[0] == '\n' || line[0] == '\0') continue;

        int report[MAX_REPORT_LENGTH];
        int size = 0;

        char *token = strtok(line, " ");
        while (token) {
            report[size++] = atoi(token);
            token = strtok(NULL, " ");
        }

        if (part == 1) {
            if (is_safe(report, size)) safe_count++;
        } else if (part == 2) {
            if (is_safe(report, size) || can_be_made_safe(report, size)) {
                safe_count++;
            }
        }
    }

    fclose(file);
    return safe_count;
}

int main() {
    const char *file_path = "input.txt";

    // Part 1
    int part1 = count_safe_reports(file_path, 1);
    printf("Part 1: Number of safe reports: %d\n", part1);

    // Part 2
    int part2 = count_safe_reports(file_path, 2);
    printf("Part 2: Number of safe reports: %d\n", part2);

    return 0;
}
