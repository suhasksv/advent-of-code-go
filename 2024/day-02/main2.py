def is_safe(report):
	"""
    Check if a report is safe based on the given criteria:
    1. All levels are either increasing or decreasing.
    2. Any two adjacent levels differ by at least 1 and at most 3.
    """
	differences = [report[i+1] - report[i] for i in range(len(report)-1)]

	# Check if all differences are within [-3, -1] or [1, 3]
	for diff in differences:
		if diff < -3 or diff > 3 or diff == 0:
			return False

	# Check if differences are all positive (increasing) or all negative (decreasing)
	all_positive = all(diff > 0 for diff in differences)
	all_negative = all(diff < 0 for diff in differences)

	return all_positive or all_negative


def can_be_made_safe(report):
	"""
    Check if a report can be made safe by removing one level.
    """
	for i in range(len(report)):
		# Create a modified report by removing the i-th level
		modified_report = report[:i] + report[i+1:]
		if is_safe(modified_report):
			return True
	return False


def main():
	# Read the input file
	with open("input.txt", "r") as file:
		lines = file.readlines()

	safe_count = 0

	for line in lines:
		if not line.strip():
			continue

		# Parse the line into a list of integers
		report = list(map(int, line.split()))

		# Check if the report is safe or can be made safe
		if is_safe(report) or can_be_made_safe(report):
			safe_count += 1

	# Print the number of safe reports
	print("Number of safe reports:", safe_count)


if __name__ == "__main__":
	main()
