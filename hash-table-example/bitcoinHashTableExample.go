package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	bk "github.com/libsv/go-bk/crypto"
	"math/big"
)

type Cell struct {
	key   string
	value string
	next  *Cell
}

type HashTable struct {
	Cells []*Cell
}

func main() {
	hashTable := NewHashTable()
	hashTable.Insert("0", "zero")
	hashTable.Insert("1", "one")
	hashTable.Insert("2", "two")
	hashTable.Insert("3", "three")
	hashTable.Insert("4", "four")
	hashTable.Insert("5", "five")
	hashTable.Insert("6", "six")
	hashTable.Insert("7", "seven")
	hashTable.Insert("8", "eight")
	hashTable.Insert("9", "nine")

	fmt.Printf("%s\n", hashTable)

}

func (hashTable *HashTable) String() string {

	var out bytes.Buffer

	for _, cell := range hashTable.Cells {

		if cell != nil {

			fmt.Fprintf(&out, "%s ", cell.value)

			if cell.next != nil {
				for nextCell := cell.next; nextCell != nil; nextCell = nextCell.next {
					fmt.Fprintf(&out, "%s ", nextCell)
				}
			}
		}
	}

	return out.String()
}

func (cell *Cell) String() string {

	var out bytes.Buffer

	for ; cell != nil; cell = cell.next {

		fmt.Fprintf(&out, "%s", cell.value)

	}

	return out.String()
}

func NewHashTable() *HashTable {
	return &HashTable{Cells: make([]*Cell, 50)}
}

func GetIndex(key string) (index int) {

	bInt := new(big.Int)
	bInt.SetBytes(bk.Sha256d([]byte(key)))

	return int(binary.BigEndian.Uint64(bInt.Bytes()) % 50)
}

func (hashTable *HashTable) Insert(key string, value string) {

	index := GetIndex(key)

	if hashTable.Cells[index] == nil {

		hashTable.Cells[index] = &Cell{key: key, value: value}
	} else {

		head := hashTable.Cells[index]

		for ; head != nil; head = head.next {

			if head.key == key {
				head.value = value
			}
			return
		}

		head.next = &Cell{key: key, value: value}
	}
}

func (hashTable *HashTable) Get(key string) (string, bool) {

	index := GetIndex(key)

	if hashTable.Cells[index] != nil {

		head := hashTable.Cells[index]

		for ; ; head = head.next {

			if head.key == key {
				return head.value, true
			}

			if head.next == nil {
				break
			}
		}

	}

	return "", false
}
