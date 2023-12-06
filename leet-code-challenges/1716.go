package main

func TotalMoney(n int) int {
	totalAmmount, lastAmmountAdded, nextMondayIndex, lastAmmountAddedOnMonday := 0, 0, 0, 0

	for i := 0; i < n; i++ {

		if i == nextMondayIndex {
			nextMondayIndex = nextMondayIndex + 7
			totalAmmount += lastAmmountAddedOnMonday + 1
			lastAmmountAdded = lastAmmountAddedOnMonday + 1
			lastAmmountAddedOnMonday++
		} else {
			totalAmmount += lastAmmountAdded + 1
			lastAmmountAdded++
		}

	}

	return totalAmmount

}
