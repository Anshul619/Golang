package main

func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

type List[T any] struct {
    head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}