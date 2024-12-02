def is_safe(report):
    differences = [report[i + 1] - report[i] for i in range(len(report) - 1)]

    # Check that all differences are within the range [1, 3] or [-3, -1]
    if not all(1 <= abs(diff) <= 3 for diff in differences):
        return False

    # Check that the sequence is monotonic (all positive or all negative differences)
    if all(diff > 0 for diff in differences) or all(diff < 0 for diff in differences):
        return True

    return False

def count_safe_reports(reports):
    safe_count = 0
    for report in reports:
        if is_safe(report):
            safe_count += 1
    return safe_count

# Read input from input.txt
with open("input.txt", "r") as file:
    reports = [list(map(int, line.split())) for line in file if line.strip()]

# Count and print the number of safe reports
safe_reports = count_safe_reports(reports)
print(f"Number of safe reports: {safe_reports}")
