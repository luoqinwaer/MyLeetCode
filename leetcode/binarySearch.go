package main

import "fmt"

//二分搜索模板
// 1、初始化 start=0， end = len -1
//2、 循环退出条件 start+1 < end
//3、 比较中点和目标值  A[mid] ==、 <、> target
//4、 判断最后两个元素是否符合：A[start]、A[end] ? target

func search(nums []int, target int) int {
	//初始化
	start:=0
	end:=len(nums)-1
	// 循环
	for start+1<end{
		mid:=start+(end-start)/2
		if nums[mid]==target{
			end=mid
		}else if nums[mid]<target{
			start=mid
		}else if nums[mid]>target{
			end=mid
		}
	}
	//最后剩下两个元素 手动判断
	//如果目标元素不在第一个 则最终end会存储目标  因此需要检查 start 和 end 两项
	if nums[start]==target{
		return start
	}
	if nums[end]==target{
		return end
	}
	return -1
}
//33. Search in Rotated Sorted Array
func searchInSortedArray(nums []int, target int) int {
	//虽然是由排序数组旋转得到，但是本质还是排序的数组 故用二分搜索达到
	if len(nums)==0{
		return -1
	}
	if len(nums)==1 {
		if nums[0]==target{
			return 0
		}else{
			return -1
		}
	}

	var start int =0
	var end int =len(nums)-1
	var base int =nums[0]

	for start+1<end{
		mid:=start+(end-start)/2
		if nums[mid]<target{
			if nums[mid]<base{
				end=mid
			}else{
				start=mid
			}
		}else{
			if nums[mid]<base{
				end=mid
			}else{
				start=mid
			}
		}
	}
	//分类讨论  可能是个旋转数组 也可能不旋转
	if nums[start]<nums[end]{
		return search(nums,target)
	}else {
		for start+1<end{
			mid:=start+(end-start)/2
			if nums[mid]>target{
				if nums[mid] >base{
					end =mid
				}else{
					end =mid
				}
			}else if nums[mid]<base{
				end=mid
			}else{
				end=mid
			}
		}
	}
	if nums[start]==target{
		return start
	}else if nums[end]==target{
		return end
	}
	return -1
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0||len(matrix[0])==0{
		return false
	}
	row:=len(matrix)
	col:=len(matrix[0])
	start:=0
	end:=row*col-1

	for start+1<end{
		mid:=start+(end-start)/2
		val:=matrix[mid/len(matrix[0])][mid%len(matrix[0])]
		if val==target{
			return true
		}else if val<target{
			start=mid
		}else if val>target{
			end=mid
		}
	}
	//接下来检查 start 和 end  因为是在二维数组中 所以要翻译到二位数组形式
	if matrix[start/col][start%col]==target||matrix[end/col][end%col]==target{
		return true
	}
	return false
}
//旋转数组
func findMin(nums []int) int {
	var i int =0
	var res int=nums[0]
	for i<len(nums)-1{
		if nums[i]<=nums[i+1]{
			i++
		}else{
			res=nums[i+1]
			break
		}
	}
	return res
}

func main()  {
	//var nums []int=[]int{4,5,6,7,0,1,2}
	var nums []int=[]int{3,5,1}
	//fmt.Println(search(nums,9))
	//fmt.Println(findMin(nums))
	fmt.Println(searchInSortedArray(nums,5))
}

