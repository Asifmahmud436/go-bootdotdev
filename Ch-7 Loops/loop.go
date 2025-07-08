package main

func bulkSend(numMessages int) float64 {
	sum := 0.00
    for i := 1 ; i <= numMessages ; i++{
		sum += 1 + float64(i-1) * 0.01
	}
	return sum
}
