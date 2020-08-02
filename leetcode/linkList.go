package main

import (
	"fmt"
)


type ListNode struct {
	    Val int
	    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current!=nil{
		//这里要用for  因为元素可能不止重复依次  需要全部删除完 再移动到下一个
		for current.Next!=nil&&current.Val==current.Next.Val{
			current.Next=current.Next.Next
		}
		current=current.Next
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head==nil{
		return head
	}
	//由于链表头节点可能被删除，所以这里设置辅助节点
	dummy:=&ListNode{Val:0}
	dummy.Next=head
	head=dummy

	var rmVal int //记录重复值
	for head.Next!=nil&&head.Next.Next!=nil{
		if head.Next.Val==head.Next.Next.Val{
			rmVal=head.Next.Val
			for head.Next!=nil&&head.Next.Val==rmVal{
				head.Next=head.Next.Next
			}
		}else{
			head=head.Next
		}
	}
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head!=nil{
		//从链表上依次取下节点，然后链接到prev节点
		temp:=head.Next
		head.Next=prev
		//指针移动
		prev=head
		head=temp
	}
	return prev
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//先定位位置 然后依次取下节点 形成局部的链表翻转  再与头尾拼接
	if head==nil || m>=n{ //无需翻转
		return head
	}
	//因为头部可能变化 所以使用dummy node
	dummy := &ListNode{Val: 0}
	dummy.Next=head
	head=dummy

	startPrev:=head
	var i int = 0
	for i<m-1{
		i++
		startPrev=startPrev.Next
	}

	start:=startPrev.Next //用作翻转之后链接
	//fmt.Println(start.Val)
	para:=start
	step:=n-m
	var j int = 0
	var prev *ListNode
	for j<=step{
		j++
		temp:=para.Next
		para.Next=prev
		prev=para
		para=temp
		fmt.Println(para.Val,prev.Val)
	}
	//链接
	//prev.Next=startPrev
	startPrev.Next=prev
	start.Next=para

	return dummy.Next
}

//func partition(head *ListNode, x int) *ListNode {
//	if head==nil{
//		return head
//	}
//	//头节点可能被修改 因此采用dummy 节点
//	dummy:=&ListNode{Val: 0}
//	dummy.Next=head
//	head=dummy
//	//思路 构造一个新的链表存储大于x值的元素 最后与原链表拼接
//	dummyLarge:=&ListNode{Val: 0}
//
//	largeHeadCopy:=dummyLarge
//	for head.Next!=nil{
//		if head.Next.Val < x{
//			head=head.Next
//		}else{
//			temp:=head.Next
//			head.Next=head.Next.Next
//
//			dummyLarge.Next=temp
//			dummyLarge=temp
//		}
//	}
//
//	head.Next=largeHeadCopy.Next
//	dummyLarge.Next=nil
//	return dummy.Next
//}

//Sort a linked list in O(n log n) time using constant space complexity.
//很懵  想到了参考数组的快速排序方式 但是不知道如何实现
func sortList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	middle:=findMiddle(head)
	tail:=middle.Next
	middle.Next=nil
	right:=sortList(tail)
	left:=sortList(head)
	result:=mergeTwoList(left, right)
	return result
}

func findMiddle(head *ListNode) *ListNode{
	//用两个快慢指针跑 快指针跑到末尾的时候  慢指针跑道的就是中点
	slow:=head
	fast:=head
	for fast.Next!=nil&&fast.Next.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
	}
	return slow
}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode{
	dummy:=&ListNode{Val: 0}
	head:=dummy
	for l1!=nil&&l2!=nil{
		if l1.Val<l2.Val{
			head.Next=l1
			l1=l1.Next
		}else{
			head.Next=l2
			l2=l2.Next
		}
		head=head.Next
	}
	//如果l1还有剩余
	for l1!=nil{
		head.Next=l1
		l1=l1.Next
		head=head.Next
	}
	//如果l2还有剩余
	for l2!=nil{
		head.Next=l2
		l2=l2.Next
		head=head.Next
	}
	return dummy.Next
}

//Given a singly linked list L: L0→L1→…→Ln-1→Ln,
//reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
func reorderList(head *ListNode)  {
	if head==nil||head.Next==nil||head.Next.Next==nil{
		return
	}
	//思路 先找到中点 再翻转后一半的链表 再依次与前一半链表拼接
	middle:=findMiddle(head)
	tail:=middle.Next
	middle.Next=nil
	//翻转
	tailReverse:=reverseList(tail)

	for tailReverse!=nil{
		tempTail:=tailReverse.Next
		tempHead:=head.Next

		tailReverse.Next=head.Next
		head.Next=tailReverse

		tailReverse=tempTail
		head=tempHead
	}
}

//Given a linked list, determine if it has a cycle in it.
func hasCycle(head *ListNode) bool {
	if head==nil||head.Next==nil{
		return false
	}
	//思路  用快慢两个指针跑 如果有环 快慢指针肯定会重合
	fast:=head.Next
	slow:=head

	for fast!=nil && fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow{
			return true
		}
	}
	return  false
}

//Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
// 快慢指针，快慢相遇之后，慢指针回到头，快慢指针步调一致一起移动，相遇点即为入环点
func detectCycle(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return nil
	}
	//思路  用快慢两个指针跑 如果有环 快慢指针肯定会重合
	fast:=head.Next
	slow:=head

	for fast!=nil && fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow{
			fast=fast.Next
			slow=head
			for slow!=fast{
				fast=fast.Next
				slow=slow.Next
			}
			return slow
		}
	}
	return  nil
}

//构造链表
func makeListNode(nums []int) *ListNode {

	if len(nums) == 0{
		return nil
	}

	res := &ListNode{
		Val:nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val:nums[i],}
		temp = temp.Next
	}

	return  res
}

func main()  {
	one := makeListNode([]int{1, 2, 3, 4, 5})
	reorderList(one)
	for one != nil {
		fmt.Println(one.Val)
		one = one.Next
	}
	//middle:=findMiddle(one)
	//fmt.Println(middle.Val)
	//two:=sortList(one)
	//for two != nil {
	//	fmt.Println(two.Val)
	//	two = two.Next
	//}
}