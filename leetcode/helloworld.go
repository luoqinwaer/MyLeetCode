package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}


//Definition for a binary tree node.
//type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
//}
//二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root==nil {
		return 0
	}
	//分 分别计算
	left := maxDepth(root.Left)
	right :=maxDepth(root.Right)
	//治  合并结果
	if left>right {
		return left+1
	}
	return right+1
}
//判断是否为平衡二叉树
//func isBalanced(root *TreeNode) bool {
//	if maxDepth(root)==-1{
//		return false
//	}
//	return true
//}
//func maxDepth(root *TreeNode) int {
//	if root==nil {
//		return 0
//	}
//	//分 分别计算
//	left := maxDepth(root.Left)
//	right :=maxDepth(root.Right)
//
//	if left==-1||right==-1||left-right>1||right-left>1{
//		return -1
//	}
//	//治  合并结果
//	if left>right {
//		return left+1
//	}
//	return right+1
//}

//Binary Tree Maximum Path Sum
func maxPathSum(root *TreeNode) int {
	if root==nil{
		return 0
	}
	maxLeft:=maxPathSum(root.Left)
	maxRight:=maxPathSum(root.Right)

	return threeMax(maxLeft,maxRight,maxLeft+maxRight+root.Val)
}
func threeMax(i,j,k int) int {
	var max int = 0
	if i>j{
		max=i
	} else {
		max=i
	}
	if max<k{
		max=k
	}
	return max
}

// 创建队列
//queue:=make([]int,0)
//// enqueue入队
//queue=append(queue,10)
//// dequeue出队
//v:=queue[0]
//queue=queue[1:]
//// 长度0为空
//len(queue)==0

func initBinary(source []int ) *TreeNode{
	// 创建队列  用以存放二叉树节点 从中依次取出赋予子字节点值
	queue:=make([]TreeNode,0)
	//处理根节点 用以返回
	var root TreeNode
	root.Val=source[0]
	root.Right=nil
	root.Left=nil
	queue=append(queue,root)

	//动态申请size大小的指针数组
	//TreeNode **nodes = new TreeNode*[len(source)];
	for i := 1; i < len(source); i++ {
		//var current int =1 //当前
		if len(queue)!=0 && i<len(source){
			temp:=queue[0]	//取出队头
			queue=queue[1:]

			//构造左儿子 右儿子
			var leftChild TreeNode
			leftChild.Left=nil
			leftChild.Right=nil
			leftChild.Val=source[i]

			var rightChild TreeNode
			rightChild.Right=nil
			rightChild.Left=nil
			rightChild.Val=source[i+1]
			i++

			temp.Left=&leftChild
			temp.Right=&rightChild

			queue=append(queue,leftChild)
			queue=append(queue,rightChild)
		}
	}

	return &root
}
func main(){
	var array1 = []int{1,2,3,4,5}
	//initBinary(array1)

	//s :="hello wo :="hello"
	//	//needle rld"
	//haysta:="ll"

	//fmt.Printf(string(s[2]))
	//fmt.Println(strStr(haystack,needle ))
	fmt.Println(maxDepth(initBinary(array1)))
}
