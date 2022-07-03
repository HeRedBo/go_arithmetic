package main

import "fmt"

/**
 *
 * 二分查找
 *
 * -------------------------------------------------------------
 * 思路分析：数组中间的值floor((low+top)/2) 
 * -------------------------------------------------------------
 * 先取数组中间的值floor((low+top)/2)然后通过与所需查找的数字进行比较，
 * 若比中间值大则将首值替换为中间位置下一个位置，继续第一步的操作；
 * 若比中间值小，则将尾值替换为中间位置上一个位置，继续第一步操作
 * 重复第二步操作直至找出目标数字
 */
func binQuery(nums []int, k int)  int {
	if len(nums) == 0 {
		return -1
	}
	low := 0 
	hight := len(nums) -1
	var val int
	for low <= hight {
		mid := ((hight-low) / 2) + low
		val = nums[mid]
		if k == val {
			return mid
		} else if  k < val {
			hight = mid -1
		} else {
			low = mid + 1
		}
	}
	return -1
}


// 递归二分查找
func binarySearch(nums []int, num int,low int, high int)  int {
	// 终止递归条件
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	// 递归查找
	if num > nums[mid] {
		// 如果待查找数据大于中间元素，则在右区间查找
		return binarySearch(nums,num, mid + 1, high)
	} else if num < nums[mid] {
		// 如果待查找数据小于中间元素，则在左区间查找
		return binarySearch(nums,num,low,mid -1)
	} else {
		// 找到 返回索引
		return mid
	}
}


func main() {
	nums := []int{1 ,2, 3, 4 ,5 ,6, 7, 8}
	fmt.Println(binQuery(nums,4))
	num := 5
	index := binarySearch(nums, num, 0, len(nums)-1)
	if index != -1 {
		fmt.Printf("Find num %d at index %d\n", num, index)
	} else {
		fmt.Printf("Num %d not exists in nums\n", num)
	}
}
