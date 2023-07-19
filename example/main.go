package main

import (
	"fmt"
	"strconv"

	"github.com/anosovs/rebrain_LRU/lrucache"
)

func main(){
	lru := lrucache.NewLRUCache(5)

	for i:=0; i<20; i++ {
		if i==19 {
			// На последней итерации обращамся к "хвостику", для поднятия его приоритета и удаляем 15 элемент, а не 14
			// Заодно смотрим на наличие старого ключа
			if value, ok := lru.Get("key14"); ok {
				fmt.Println(value)
			} else {
				fmt.Println("Not found key14")
			}
			
			if value, ok := lru.Get("key0"); ok {
				fmt.Println(value)
			} else {
				fmt.Println("Not found key0")
			}			
		}
		lru.Add("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
		
	}

}
