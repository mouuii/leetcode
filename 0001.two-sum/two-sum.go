package problem0001

// twoSum
// solution :
//1.define a map , key is the value ,value is the sequence
//2.loop througe the nums, record the data,meanwhile judge if the map has the match data
//3.problem solve ,return  the value
func twoSum(nums []int, target int) []int {
	storeMap := map[int]int{}
	for index, value := range nums {
		if v, ok := storeMap[target-value]; ok {
			return []int{v, index}
		}
		storeMap[value] = index
	}
	return nil
}
