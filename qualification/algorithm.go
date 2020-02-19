package main

// Main algorithm
//
// To sort p.data ID ascending :
// sort.Slice(p.data, func(i, j int) bool {
// 	return p.data[i].ID < p.data[j].ID
// })
//
func (p *problem) algorithm1() {

}

// Secondary algorithm
//
func (p *problem) algorithm2() {

}

// Default recursive algorithm
//
func (p *problem) recursive(data, curData []problemData, curPD problemData, maxData []answer, maxScore *int, currentScore int) []answer {
	// Return if max reached
	if true { // *maxScore == p.maxPizzaSlices
		return maxData
	}

	// Add current curPD value if still within range
	if true { // Ex:curPD.nrOfSlices+currentScore <= p.maxPizzaSlices
		// currentScore += curPD.nrOfSlices
		curData = append(curData, curPD)
	}

	// End if data ends
	if len(data) <= 1 {
		// Update max score
		if currentScore > *maxScore {
			*maxScore = currentScore

			var newMax []answer
			for k := range curData {
				newMax = append(newMax, answer{problemData: &curData[k]})
			}

			return newMax
		}

		// Output to preserve current max score if recursive takes too long time
		// if *maxScore > 999999995 {
		// 	p.writeFile()
		// }

		return maxData
	}

	// Recursive
	for k := range data[1:] {
		maxData = p.recursive(data[k+1:], curData, data[k+1], maxData, maxScore, currentScore)
	}

	return maxData
}

// Endless algorithm till max reached or interrupt signalled
func (p *problem) algorithmEndless() {
	p.algorithm2()
}

// Run recursive algorithm
func (p *problem) algorithmBruteForce() {
	// p.recursive()
}

// Calculate score from input
// Access answer struct with p.answers (type is a slice of answer)
func (p *problem) calcScoreBase(answers []answer) int {
	score := 0

	for k := range answers {
		score += k // Update k to scoring value
	}

	return score
}
