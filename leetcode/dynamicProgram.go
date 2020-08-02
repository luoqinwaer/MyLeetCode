package main

import "fmt"

//动态规划 专题
//子问题如果有交集  可以考虑采用动态规划来求解

//e.g. 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
/*
[
   [2],
  [3,4],
 [6,5,7],
[4,1,8,3]
]
*/
//动态规划，自底向上
func minimumTotalBottomToTop(triangle [][]int) int {
	if len(triangle)==0||len(triangle[0])==0{
		return 0
	}
	//以 f[i][j] 表示从i，j出发 到达最后一层的最短路径
	var l=len(triangle)
	var f=make([][]int,l) //定义切片 保存结果
	//初始化  ？？？  干什么用
	for i:=0;i<l;i++{
		for j:=0;j<len(triangle[i]);j++{
			if f[i]==nil{
				f[i]=make([]int, len(triangle[i]))
			}
			f[i][j]=triangle[i][j]  //什么作用？ 不就是相当于复制了一下么
									//因为需要存在一个副本保存之前计算的结果
		}
	}
	//求解
	for i:=len(triangle)-2;i>=0;i--{
		for j:=0;j<len(triangle[i]);j++{
			f[i][j]=min(f[i+1][j],f[i+1][j+1])+triangle[i][j]
		}
	}
	//答案
	return f[0][0]

}
func min(a, b int) int{
	if a>b{
		return b
	}
	return a
}

func minimumTotalTopToBottom(triangle [][]int) int {
	if len(triangle)==0||len(triangle[0])==0{
		return 0
	}
	//以 f[i][j] 表示从i，j出发 到达最后一层的最短路径
	var l=len(triangle)
	var f=make([][]int,l) //定义切片 保存结果
	//初始化  ？？？  干什么用
	for i:=0;i<l;i++{
		for j:=0;j<len(triangle[0]);j++{
			if f[i]==nil{
				f[i]=make([]int, len(triangle[i]))
			}
			f[i][j]=triangle[i][j]  //什么作用？ 不就是相当于复制了一下么
			//因为需要存在一个副本保存之前计算的结果
		}
	}
	//求解
	for i:=1;i<l;i++{
		for j:=0;j<len(triangle[i]);j++{
			//分类 上一层没有左边值  上一层没有右边值
			if j-1<0{ //没有左边值
				f[i][j] =f[i-1][j]+triangle[i][j]
			}else if j>=len(triangle[i]){
				f[i][j]=f[i-1][j-1] + triangle[i][j]
			}else{
				f[i][j]=min(f[i-1][j],f[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	result:=f[l-1][0]  //在左后一排中找最小值
	for i:=1;i<len(f[l-1]);i++{
		result=min(result,f[l-1][i])
	}
	return result
}

//Minimum Path Sum
/*
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
*/
func minPathSum(grid [][]int) int {
	//动态规划 自底向上
	if len(grid)==0||len(grid[0])==0{
		return 0
	}
	//l :=len(grid)
	//这里可以初始化的原因是因为 这里的结果只能来自一个方向
	for i:=1;i<len(grid);i++{
		grid[i][0]=grid[i][0]+grid[i-1][0]
	}
	for j:=1;j<len(grid[0]);j++{
		grid[0][j]=grid[0][j]+grid[0][j-1]
	}
	//求解
	for i:=1;i<len(grid);i++{
		for j:=1;j<len(grid[i]);j++{
			grid[i][j]=min(grid[i][j-1],grid[i-1][j])+grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

//Unique Paths
func uniquePaths(m int, n int) int {
	if m==0||n==0{
		return 0
	}
	f:=make([][]int, m)
	for i:=0;i<m;i++{
		f[i]=make([]int,n)
	}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if i==0||j==0{
				f[i][j]=1
			}
		}
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			f[i][j]=f[i-1][j]+f[i][j-1]
		}
	}

	return f[m-1][n-1]
}

//Unique Paths II
//  需要根据附近是否有阻碍 调整一下 推导公式
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid)==0||len(obstacleGrid[0])==0{
		return 0
	}
	row:=len(obstacleGrid)
	col:=len(obstacleGrid[0])
	f:=make([][]int,row)
	//初始化空值
	for i := 0; i < row; i++ {
		f[i]=make([]int,col)
	}
	//初始化最上边行最左边列值
	for i:=0;i<row;i++{
		if obstacleGrid[i][0]==0{
			f[i][0]=1
		}else{
			break
		}
	}
	for j:=0;j<col;j++{
		if obstacleGrid[0][j]==0{
			f[0][j]=1
		}else{
			break
		}
	}
	//求解  分情况讨论
	for i:=1;i<row;i++{
		for j:=1;j<col;j++{
			if obstacleGrid[i][j]!=0{
				f[i][j]=0
			}else if obstacleGrid[i-1][j]!=0{
				f[i][j]=f[i][j-1]
			}else if obstacleGrid[i][j-1]!=0{
				f[i][j]=f[i-1][j]
			}else{
				f[i][j]=f[i][j-1]+f[i-1][j]
			}
		}
	}
	//结果
	return f[row-1][col-1]
}

//Climbing Stairs
func climbStairs(n int) int {
	var temp []int=[]int{0,1,2}
	if n<3{
		return temp[n]
	}
	f:=make([]int,n+1)
	f[0]=0
	f[1]=1
	f[2]=2
	for i:=3;i<n+1;i++{
		f[i]=f[i-1]+f[i-2]
	}
	return f[n]
}

//Jump Game
func canJump(nums []int) bool {
	if len(nums)==1{
		return true
	}
	if len(nums)==2{
		return nums[0]>=1
	}
	length:=len(nums)
	f:=make([]int,length-1)
	f[0]=nums[0]
	for i:=1;i<length-1;i++{
		if f[i-1]>0{ //前一步可达
			f[i]=i+nums[i]
			if f[i]>=length-1{
				return true
			}
		}else{
			f[i]=0
		}
	}
	return false
}

func main(){
	//var nums []int=[]int{2,3,1,1,4}
	var nums []int=[]int{2,0,0}
	//var f [][]int=[][]int{
	//	{0,0,0},
	//	{0,1,0},
	//	{0,0,0}}
	//fmt.Println(uniquePaths(7,3))
	//fmt.Println(uniquePathsWithObstacles(f))
	//fmt.Println(climbStairs(4))
	fmt.Println(canJump(nums))
}