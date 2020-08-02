package main

import "runtime/debug"

//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//用两个栈实现
type MinStack struct {
	min []int
	stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:make([]int,0),
		stack:make([]int,0),
	}
}


func (this *MinStack) Push(x int)  {
	min:=this.GetMin()
	if x<min{
		this.min=append(this.min,x)
	}else{
		this.min=append(this.min,min)
	}
	this.stack=append(this.stack,x)
}


func (this *MinStack) Pop()  {
	if len(this.stack)==0{
		return
	}
	this.stack=this.stack[:len(this.stack)-1]
	this.min=this.min[:len(this.min)-1]
}


func (this *MinStack) Top() int {
	if len(this.stack)==0{
		return 0
	}
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	if len(this.min)==0{
		return 1<<31
	}
	min:=this.min[len(this.min)-1]
	return min
}


type MyQueue struct {
	stack []int
	back []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack:make([]int,0),
		back:make([]int,0),
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.stack=append(this.stack,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	//如果back栈里有元素 就返回栈顶  否则 将stack栈中的元素倒到back栈中
	var res int
	if len(this.back)==0{
		for len(this.stack)>0{
			temp:=this.stack[len(this.stack)-1]
			this.stack=this.stack[:len(this.stack)-1]
			this.back=append(this.back,temp)
		}
		res=this.back[len(this.back)-1]
		this.back=this.back[:len(this.back)-1]
	}else{
		res=this.back[len(this.back)-1]
		this.back=this.back[:len(this.back)-1]
	}
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	var res int
	if len(this.back)==0{
		for len(this.stack)>0{
			temp:=this.stack[len(this.stack)-1]
			this.stack=this.stack[:len(this.stack)-1]
			this.back=append(this.back,temp)
		}
		res=this.back[len(this.back)-1]
		//this.back=this.back[:len(this.back)-1]
	}else{
		res=this.back[len(this.back)-1]
		//this.back=this.back[:len(this.back)-1]
	}
	return res
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack)==0&&len(this.back)==0
}
