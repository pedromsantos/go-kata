package algorithm

import "strconv"

// ChecksumCalculator duplicates the same checksum computation (sum of char
// codes mod 10) in both methods instead of extracting it once -- if the
// algorithm ever changes, both call sites must be updated in lockstep or
// they silently disagree. Connascence of Algorithm.
type ChecksumCalculator struct{}

// AddChecksum appends a checksum digit to inputData.
func (ChecksumCalculator) AddChecksum(inputData string) string {
	sum := 0
	for _, ch := range inputData {
		sum += int(ch)
	}
	checksum := sum % 10
	return inputData + strconv.Itoa(checksum)
}

// Check verifies the checksum digit appended to inputDataWithChecksum.
func (ChecksumCalculator) Check(inputDataWithChecksum string) bool {
	inputData := inputDataWithChecksum[:len(inputDataWithChecksum)-1]
	expected, err := strconv.Atoi(inputDataWithChecksum[len(inputDataWithChecksum)-1:])
	if err != nil {
		return false
	}
	sum := 0
	for _, ch := range inputData {
		sum += int(ch)
	}
	return sum%10 == expected
}
