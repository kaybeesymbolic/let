package let

//Times takes 2 or more arguments and return answer after multiplication
var Times = func(nums ...interface{}) (ans float64) {
	temp := 1.0
	check := 0.0
	for _, each := range nums {
		switch each.(type) {
		case int:
			temp *= float64(each.(int))
			ans = temp
		case float32:
		case float64:
			temp *= each.(float64)
			ans = temp
		default:
			temp = check
			ans = temp
		}
	}

	return ans

}
