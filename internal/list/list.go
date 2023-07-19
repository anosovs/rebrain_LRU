package list

import (
	"fmt"
	"log"
)

type ListInterface interface {
	// Добавление с наивысшим приоритетом в начало списка
	AddFront(name string) string
	// AddBack(name string)
	// RemoveFront() error

	// Удаляет хвостик
	RemoveBack() (removedKey string, err error)

	// Сдвигает элемент с указанным именем в начало списка
	MoveToFront(name string) error

	// Удаляет из списка по заданному имени, объединяет соседей
	RemoveByName(name string) error

	// Показывает текущее состояние списка
	Traverse() error
}

type element struct{
	name string
	next *element
}

type list struct {
	maxLen int
	len int
	head *element
}

// Инициализация списка
func InitList(n int) ListInterface {
	return &list{
		maxLen: n,
	}
}


func (l *list) AddFront(name string) string{
	var removedKey string
	var err error
	element := &element{
		name: name,
	}
	if l.head == nil {
		l.head = element
	} else {
		element.next = l.head
		l.head = element
		if l.len >= l.maxLen {
			
			removedKey, err = l.RemoveBack()
			if err != nil {
				log.Println(err)
			}
		}
	}	
	l.len ++
	return removedKey
}




func (l *list) RemoveBack() (removedKey string, err error) {
	if l.head ==  nil {
		return "", fmt.Errorf("empty list")
	}
	var prev *element
	current := l.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
		removedKey = current.name
	} else {
		l.head = nil
	}
	
	l.len--
	return removedKey, nil
}


func (l *list) MoveToFront(name string) error {
	if l.head ==  nil {
		return fmt.Errorf("empty list")
	}
	if err := l.RemoveByName(name); err != nil{
		return err
	}
	l.AddFront(name)
	return nil
}


func (l *list) RemoveByName(name string) error {
	if l.head ==  nil {
		return fmt.Errorf("empty list")
	}
	var prev *element
	current := l.head
	for current.name != name && current.next!=nil {
		prev = current
		current = current.next
	}
	if current.name == name {
		if prev != nil {
			prev.next = current.next
		} else {
			l.head = current.next
		}
		l.len--
		current = &element{}
		return nil
	} else {
		return fmt.Errorf("%s not found", name)
	}
}



func (l *list) Traverse() error {
    if l.head == nil {
        return fmt.Errorf("TranverseError: List is empty")
    }
    current := l.head
    for current != nil {
        fmt.Println(current.name)
        current = current.next
    }
	fmt.Println("---------------")
    return nil
}

// func (l *list) AddBack(name string) {
// 	element := &element{
// 		name: name,
// 	}
// 	if l.head == nil {
// 		l.head = element
// 	} else {
// 		if l.len >= l.maxLen {
// 			l.RemoveBack()
// 		}
// 		current := l.head
// 		for current.next != nil {
// 			current = current.next
// 		}
// 		current.next = element
// 	}
// 	l.len++
// }

// func (l *list) RemoveFront() error {
// 	if l.head == nil {
// 		return fmt.Errorf("empty list")
// 	}
// 	l.head = l.head.next
// 	l.len--
// 	return nil
// }