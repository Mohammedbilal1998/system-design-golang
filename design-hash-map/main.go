package main

import (
	"fmt"
	"log"
)

type keyValue struct{
	key int
	value interface{}
}

type HashMap struct{
	size int
	table [][]keyValue
}

func NewHashMap(size int)*HashMap{
	return &HashMap{
		size: size,
		table: make([][]keyValue, size),
	}
}

func (h *HashMap) hash(key int) int{
	return key % h.size
}

func (h *HashMap) get(key int) (keyValue, error){
	hashIndex := h.hash(key)
	fmt.Println(h.table[hashIndex])

	for _, item := range h.table[hashIndex]{
		if item.key == key{
			return item, nil
		}
	}
	return keyValue{}, fmt.Errorf("key value pair not found")
}

func (h *HashMap) set(key int, value interface{}){
	hashIndex := h.hash(key)
	for _, item := range h.table[hashIndex]{
		if item.key == key{
			item.value = value
			return
		}
	}
	h.table[hashIndex] = append(h.table[hashIndex], keyValue{key, value})
	return
}

func (h *HashMap) remove(key int)error{
	hashIndex := h.hash(key)

	for i, item := range h.table[hashIndex]{
		if item.key == key{
			h.table[hashIndex] = append(h.table[hashIndex][:i], h.table[hashIndex][i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("key value pair doesnot exist")
}

func main(){
	hash_map := NewHashMap(10)
	fmt.Println(hash_map.table)
	//insert
	for i:=0; i<100; i++{
		hash_map.set(i, fmt.Sprintf("bilal-%d", i))
	}
	fmt.Println(hash_map.table)

	hash_map.remove(99)
	fmt.Println(hash_map.table)
	data, err := hash_map.get(99)
	if err != nil{
		log.Printf("error occured: %w", err)
	}
	fmt.Println("get result: ", data)

}