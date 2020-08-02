package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func evalRPN(tokens []string) int {
	if len(tokens)==0{
		return 0
	}
	stack:=make([]int,0)
	for i:=0;i<len(tokens);i++{
		switch tokens[i] {
		case "+","-","*","/":
			if len(stack)<2{
				return -1 //表示出错
			}
			b:=stack[len(stack)-1]
			a:=stack[len(stack)-2]
			stack=stack[:len(stack)-2]
			var result int
			switch tokens[i]{
			case "+":
				result=a+b
			case "-":
				result=a-b
			case "*":
				result=a*b
			case "/":
				result=a/b
			}
			stack=append(stack,result)
		default:
			val,_:=strconv.Atoi(tokens[i])
			stack=append(stack,val)
		}
	}
	return stack[0]
}
type TreeNode struct {
		Val int
	    Left *TreeNode
	    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result:=make([]int,0)
	if root==nil{
		return result
	}
	stack:=make([] *TreeNode,0)
	for len(stack)>0||root!=nil{
		for root!=nil{
			stack=append(stack,root)
			root=root.Left
		}
		//弹出
		val:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		result=append(result,val.Val)
		root=val.Right
	}
	return result
}

func removeElement(nums []int, val int) int {
	//双指针法
	var j int =0
	for i:=0;i<len(nums);i++{
		if nums[i]!=val{
			nums[j]=nums[i]
			j++
		}
	}
	return j
}

func isValid(s string) bool {
	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune

	for _, char := range s {
		fmt.Println(reflect.TypeOf(char))
		if char == '(' || char == '{' || char == '[' {
			// 入栈
			stack = append(stack, char)
			// 循环中，stack不能为空
		} else if len(stack) > 0 && brackets[char] == stack[len(stack) - 1] {
			// 栈中有数据，且此元素与栈尾元素相同
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}

	// 循环结束，栈中还有数据则 false
	return len(stack) == 0
}

func searchInsert(nums []int, target int) int {
	var i int=0
	var res int =0
	for i<len(nums){
		if nums[i]<target{
			i++
		}else{
			res = i
			break
		}
	}
	//遍历完 仍未找到
	if nums[len(nums)-1] < target{
		res=len(nums)
	}
	return res
}

//Number of Islands
//思路 只要我们遇到为‘1’的点，说明这一块一定是一个岛屿，于是res++，并且我们此时就将与其相连的所有为‘1’的点都改为‘0’，说明是已经计算在内的岛屿。
func numIslands(grid [][]byte) int {
	var count int
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j]=='1'&&dfs(grid,i,j)>=1{
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i,j int) int{
	if i<0||i>len(grid)||j<0||j>len(grid[0]){
		return 0 //范围之外  返回0
	}
	if grid[i][j]=='1'{
		//岛屿周围的陆地都标记为1 表示是同一块岛屿
		grid[i][j]=0
		return dfs(grid,i-1,j)+dfs(grid,i+1,j)+dfs(grid,i,j-1)+dfs(grid,i,j+1)+1
	}
	return 0
}

func levelOrder(root *TreeNode) [][]int {
	result:=make([][]int,0)
	if root==nil{
		return result
	}
	quene:=make([]*TreeNode,0)
	quene=append(quene,root)
	for len(quene)>0{
		list:=make([]int,0)
		l:=len(quene)
		for i:=0;i<l;i++{
			//pop
			level:=quene[0]
			quene=quene[1:]
			list=append(list,level.Val)
			if level.Left!=nil{
				quene=append(quene,level.Left)
			}
			if level.Right!=nil{
				quene=append(quene,level.Right)
			}
		}
		result=append(result,list)
	}
	return result
}

//虽然写出来了 ，但是不优雅
func spiralOrder(matrix [][]int) []int {
	 //列数 len(matrix[0]) 2  //行数 len(matrix)  3

	//形成螺旋的顺序是一样的，问题也就简单了
	res:=make([]int,0)
	if len(matrix)==0{
		return res
	}
	var spiralFloor int //旋转的层数
	var spiralCount int //旋转了多少次
	totalCount:=len(matrix[0])*len(matrix) //元素的总数目

	for spiralCount<totalCount{
		for i:=spiralFloor;i<len(matrix[0])-spiralFloor;i++{
			res=append(res,matrix[spiralFloor][i])

			spiralCount++
			if spiralCount>=totalCount {
				break
			}
		}
		if spiralCount>=totalCount {
			break
		}
		for j:=spiralFloor+1;j<len(matrix)-spiralFloor;j++{
			res=append(res,matrix[j][len(matrix[0])-spiralFloor-1])
			spiralCount++
			if spiralCount>=totalCount {
				break
			}
		}
		if spiralCount>=totalCount {
			break
		}
		for k:=len(matrix[0])-spiralFloor-1-1;k>=spiralFloor;k--{
			res=append(res,matrix[len(matrix)-spiralFloor-1][k])
			spiralCount++
			if spiralCount>=totalCount {
				break
			}
		}
		if spiralCount>=totalCount {
			break
		}
		for z:=len(matrix)-spiralFloor-1-1;z>=spiralFloor+1;z--{
			res=append(res,matrix[z][spiralFloor])
			spiralCount++
			if spiralCount>=totalCount {
				break
			}
		}
		if spiralCount>=totalCount {
			break
		}
		spiralFloor++
	}
	return res

}

func main(){
	//str:=[]string{"4", "13", "5", "/", "+"}
	//var s string="()[]{}"
	//isValid(s)
	//fmt.Print(evalRPN(str))

	//var num []int=[]int{1,3,5,6}
	//fmt.Println(searchInsert(num,0))
	var nums [][]int=[][]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12}}
	fmt.Println(spiralOrder(nums))
}