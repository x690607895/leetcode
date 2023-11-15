package main

import "fmt"

type RandomizedSet struct {
	data map[int]struct{}
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]struct{})}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.data[val]; ok {
		return false
	}
	this.data[val] = struct{}{}
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.data[val]; !ok {
		return false
	}
	delete(this.data, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	for k := range this.data {
		return k
	}
	return 0
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Insert(3))
	fmt.Println(obj.Remove(2))
	fmt.Println(obj.GetRandom())
}
