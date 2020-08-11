package two_sum

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        test := target - nums[i]
        if _, found := m[test]; found {
            return []int{m[test], i}
        }
        m[nums[i]] = i
        
    }
    return nil    
}
