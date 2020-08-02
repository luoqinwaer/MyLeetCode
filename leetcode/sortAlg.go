package main

//快速排序
func QuickSort(nums []int) []int{
	//找一个参考 将数组依照参照分为两部分  左边小于参考 右边大于
	quickSort(nums,0,len(nums)-1)
	return nums
}

func quickSort(nums []int, start int, end int) {
	if start<end{
		//基于分治策略
		pivot:=partition(nums,start,end)
		quickSort(nums,0,pivot-1)
		quickSort(nums,pivot+1,end)
	}
}

func partition(nums []int, start int, end int) int {
	//选取最后一个元素作为基准pivot
	p:=nums[end]
	i:=start
	for j:=start;j<end;j++{
		if nums[j]<p{
			swap(nums,i,j)
			i++
		}
	}
	swap(nums,i,end)
	return i
}

func swap(nums []int, i int, j int) {
	t:=nums[i]
	nums[i]=nums[j]
	nums[j]=t
}


