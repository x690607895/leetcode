package main

import (
	"log"
)

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

// 实现 MinStack 类:

// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。

// 示例 1:

// 输入：
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]

// 输出：
// [null,null,null,null,-3,null,0,-2]

// 解释：
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnkq37/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	obj := Constructor()
	obj.Push(2147483646)
	obj.Push(2147483646)
	obj.Push(2147483647)
	log.Println(obj.Top())
	obj.Pop()
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.GetMin())
	obj.Pop()
	obj.Push(2147483647)
	log.Println(obj.Top())
	log.Println(obj.GetMin())
	obj.Push(-2147483648)
	log.Println(obj.Top())
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.GetMin())
	log.Println(obj)
}

type MinStack struct {
	data     []int
	num      int
	top, min *int
}

func Constructor() MinStack {
	return MinStack{[]int(nil), 0, nil, nil}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	this.top = &this.data[this.num]
	if this.num == 0 || (this.min != nil && *this.min > val) {
		this.min = &this.data[this.num]
	}
	this.num++
}

func (this *MinStack) Pop() {
	if this.num <= 0 {
		return
	}
	this.num--
	if this.num == 0 {
		this.data = nil
		this.top = nil
		this.min = nil
	} else {
		if this.min != nil && this.data[this.num] == *this.min {
			this.min = nil
		}
		this.data = this.data[:this.num]
		this.top = &this.data[this.num-1]
	}
}

func (this *MinStack) Top() int {
	if this.top == nil {
		return 0
	}
	return *this.top
}

func (this *MinStack) GetMin() int {
	if this.min == nil {
		if this.num <= 0 {
			return 0
		}
		for k, v := range this.data {
			if this.min == nil || *this.min > v {
				this.min = &this.data[k]
			}
		}
	}
	return *this.min
}
