package main

import "fmt"

// 实现 SmallestInfiniteSet 类：

// SmallestInfiniteSet() 初始化 SmallestInfiniteSet 对象以包含 所有 正整数。
// int popSmallest() 移除 并返回该无限集中的最小整数。
// void addBack(int num) 如果正整数 num 不 存在于无限集中，则将一个 num 添加 到该无限集中。

// 用个map来存储已经删除的数字，如果有已经删除的数字被添加，则从map里剔除

type SmallestInfiniteSet struct {
	min        int
	deletedMap map[int]struct{}
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		1, make(map[int]struct{}, 1000),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	result := this.min
	this.deletedMap[this.min] = struct{}{}
	for {
		_, ok := this.deletedMap[this.min]
		if !ok {
			break
		}
		this.min++
	}
	return result
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.deletedMap[num]; ok {
		if this.min > num {
			this.min = num
		}
		delete(this.deletedMap, num)
	}
}

func main() {
	obj := Constructor()
	var param_1 int
	// param_1 := obj.PopSmallest()
	// fmt.Println(param_1)
	obj.AddBack(2)
	param_1 = obj.PopSmallest()
	param_1 = obj.PopSmallest()
	param_1 = obj.PopSmallest()
	obj.AddBack(1)
	param_1 = obj.PopSmallest()
	param_1 = obj.PopSmallest()
	param_1 = obj.PopSmallest()
	// obj.AddBack(5)
	// param_1 = obj.PopSmallest()
	fmt.Println(param_1)
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
