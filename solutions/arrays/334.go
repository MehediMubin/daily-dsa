import "math"

func increasingTriplet(nums []int) bool {
    first_num, second_num := math.MaxInt32, math.MaxInt32
    for _, num := range nums {
        if num <= first_num {
            first_num = num
        } else if num <= second_num {
            second_num = num
        } else {
            return true
        }
    }
    return false
}