package index

func change(amount int, coins []int) int {
	if amount == 0{
		return 1
	}
	if len(coins) == 0 {
		return 0
	}
	resultArray := make([]int, amount+1, amount+1)
	resultArray[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			resultArray[i] += resultArray[i-coin]
		}
	}
	return resultArray[amount]
}
