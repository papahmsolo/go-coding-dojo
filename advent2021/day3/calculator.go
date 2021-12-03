package main

// powerConsumption calculates the p.c. for given telemetry data
func powerConsumption(data []string) int {

	vLength := len(data[0])
	counter := make([]int, vLength)

	// accumulates the metadata of the set
	// if value in counter is below 0 that means that the most common bit is 0
	// positive values for most common 1
	for _, line := range data {
		for i, v := range line {
			if v == '1' {
				counter[i]++
			} else {
				counter[i]--
			}
		}
	}

	// fetch gamma and epsilon
	var gamma, epsilon int
	for i := 0; i < vLength; i++ {
		if counter[vLength-i-1] > 0 {
			gamma = gamma + 1<<i
		} else {
			epsilon = epsilon + 1<<i
		}
	}

	return gamma * epsilon
}
