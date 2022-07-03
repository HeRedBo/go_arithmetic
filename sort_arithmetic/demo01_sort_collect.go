package main

import (
	"fmt"
)

func  BubbleSort(nums []int) []int {
	len := len(nums)
	for i := 0; i < len; i ++ {
		flag := false
		for j := 0; j < len -i -1; j ++ {
			if nums[j] > nums[j +1] {
				nums[j], nums[j +1] = nums[j +1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nums
}


func InsertSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	for i := 1; i < len(nums); i ++ {
		// 获取当前需要比较的元素
		value := nums[i]
		// 在已排序好的区间找到插入位置
		j := i - 1
		// 内层循环控制 比较 并 插入
		for ; j >= 0; j -- {
			if nums[j] > value {
				nums[j + 1 ] = nums[j]
			} else {
				break
			}
		}
		// 插入数据 value
		nums[j + 1] = value
	}
	return nums
}

func SelectSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	for i := 0; i < len(nums); i ++ {
		min := i //  //先假设最小的值的位置 已排序区间初始化为空，未排序区间初始化待排序切片

		// 从未排序区间第二个元素开始遍历，直到找到最小值
		for j := i + 1; j < len(nums); j ++ {
			if nums[j] < nums[min] {
				min = j
			}
		}

		if min != i {
			nums[i], nums[min] = nums[min],nums[i]
		}
	}
	return nums
}
/**
归并排序
// 归并排序是建立在归并操作上的一种有效的排序算法。
// 归并排序将待排序的序列分成若干组，保证每组都有序，然后再进行合并排序，最终使整个序列有序。
// 该算法是采用分治法的一个非常典型的应用。
### 原理图可看该文章 ：
## https://www.cnblogs.com/sunshineliulu/p/8573991.html
 */
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return  nums
	}
	// 获取分区位置
	mid := len(nums) / 2
	// 通过递归分区
	left := MergeSort(nums[0:mid])
	right := MergeSort(nums[mid:])

	// 排序后合并
	return merge(left, right)
}

func merge(left []int , right []int) []int {
	i, j := 0,0
	m,n := len(left), len(right)
	var result []int // 用于存放结果集

	for {
		if i >= m || j >=n {
			break
		}
		if left[i] < right[j] {
			result = append(result, left[i])
			i ++
		} else {
			result = append(result, right[j])
			j ++
		}
	}

	// 如果左侧区间还没有遍历完，将剩余数据放到结果集
	if i != m {
		result = append(result,left[i:]...)
	}
	//如果右侧区间还没有遍历完，将剩余数据放到结果集
	if j != n {
		result = append(result,right[i:]...)
	}
	return result
}
/**

## 快速排序采用分治法实现排序，具体步骤：

# 1、从数列中挑出一个数作为基准元素。通常选择第一个或最后一个元素。
# 2、扫描数列，以基准元素为比较对象，把数列分成两个区。规则是：小的移动到基准元素前面，大的移到后面，相等的前后都可以。分区完成之后，基准元素就处于数列的中间位置。
# 然后再用同样的方法，递归地排序划分的两部分。
 */
func QuickSort(nums []int) []int {
	//判断数组元素：如果小于1，那么直接返回
	len := len(nums)
	if len <= 1 {
		//*递归出口：数组元素小于1
		return nums
	}

	// 切片长度大于1：取出第一个元素，当做中间值
	middle := nums[0]
	var left,right []int
	for i:= 1; i < len; i ++  {
		if nums[i] < middle {
			left = append(left, nums[i])
		} else {
			right = append(right,nums[i])
		}
	}

	//递归点：left和right都是无序数组，与父问题一致，规模较小
	left = QuickSort(left)
	right = QuickSort(right)

	// 合并数据结果 进行返回
	var result []int
	result = append(left,middle)
	result = append(result,right...)
	return result
}


func main() {
	nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	//nums = BubbleSort(nums)
	//nums = InsertSort(nums)
	//nums = SelectSort(nums)
	//nums = MergeSort(nums)
	nums = QuickSort(nums)
	fmt.Println(nums)
}






