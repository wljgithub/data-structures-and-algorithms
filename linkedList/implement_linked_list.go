package main

import (
	"errors"
	"fmt"
)

/*
链表实现过程中需要注意的细节：

1.
	如果希望遍历到最后一个元素，用：
	for current.next!= nil
	如果希望遍历到倒数第二个元素，用：
	var prev *ele
	for current.next != nil{
		prev = current
}

2. 如果添加和删除的时候，要考虑头节点为空，当前节点为尾节点
*/

type ele struct {
	name string
	next *ele
}

// 这种实现方式有些巧妙，用一个struct来包住linkedlist
type singleList struct {
	len  int
	head *ele
}

func (this *singleList) AddFront(s string) {
	e := &ele{name: s}

	if this.head == nil {
		this.head = e
	} else {
		e.next = this.head
		this.head = e
	}
	this.len++
}

func (this *singleList) Size() int {
	return this.len
}

func (this *singleList) Traverse() error {
	if this.head == nil {
		return errors.New("list is empty")
	}
	current := this.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}
	return nil
}

func (this *singleList) RemoveFront() error {
	if this.head == nil {
		return errors.New("list is empty")
	}
	this.head = this.head.next
	this.len--
	return nil
}

func (this *singleList) RemoveBack() error {
	if this.head == nil {
		return errors.New("list is epmty")
	}
	if this.head.next == nil {
		this.head = nil
	} else {
		var prev *ele
		var current = this.head
		for current.next != nil {
			prev = current
			current = current.next
		}
		prev.next = nil
	}
	this.len--
	return nil
}

func initList() *singleList {
	return &singleList{}
}

//func (s *singleList) AddFront(name string) {
//	ele := &ele{name: name}
//
//	if s.head == nil {
//		s.head = ele
//	} else {
//		ele.next = s.head
//		s.head = ele
//	}
//	s.len++
//	return
//}

//func (s *singleList) AddBack(name string) {
//	ele := &ele{name: name}
//	if s.head == nil {
//		s.head = ele
//	} else {
//		current := s.head
//		for current.next != nil {
//			current = current.next
//		}
//		current.next = ele
//	}
//	s.len++
//	return
//}

//func (s *singleList) RemoveFront() error {
//	if s.head == nil {
//		return fmt.Errorf("list is empty")
//	}
//	s.head = s.head.next
//	s.len--
//	return nil
//}

//func (s *singleList) RemoveBack() error {
//	if s.head == nil {
//		return fmt.Errorf("removeBack:list is empty")
//	}
//	var prev *ele
//	current := s.head
//	for current.next != nil {
//		prev = current
//		current = current.next
//	}
//
//	if prev != nil {
//		prev.next = nil
//	} else {
//		s.head = nil
//	}
//
//	s.len--
//	return nil
//}

//func (s *singleList) Front() (string, error) {
//	if s.head == nil {
//		return "", fmt.Errorf("Single List is empty")
//	}
//	return s.head.name, nil
//}

//func (s *singleList) Size() int {
//	return s.len
//}

//func (s *singleList) Traverse() error {
//	if s.head == nil {
//		return fmt.Errorf("Traverse Error: List is empty")
//	}
//	current := s.head
//	for current != nil {
//		fmt.Println(current.name)
//		current = current.next
//	}
//	return nil
//}

func main() {
	singleList := initList()
	fmt.Println("AddFront: A")
	singleList.AddFront("A")

	singleList.AddFront("B")
	fmt.Println("AddFront: B")

	singleList.AddFront("C")
	fmt.Println("AddFront: C")

	fmt.Println("Size: ", singleList.Size())

	err := singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("RemoveFront")
	err = singleList.RemoveFront()
	if err != nil {
		fmt.Println("RemoveFront Error:", err.Error())
	}

	fmt.Println("RemoveBack")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Println("RemoveBack Error: ", err.Error())
	}

	fmt.Println("RemoveBack")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Println("RemoveBack Error: ", err.Error())
	}

	fmt.Println("RemoveBack")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Println("RemoveBack Error: ", err.Error())
	}

	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Size: ", singleList.Size())
}
