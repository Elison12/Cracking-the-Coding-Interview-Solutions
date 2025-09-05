package main

import (
	"fmt"
	"hash/fnv"
)

type Table struct {
	m     int
	table [][]kv
}

type kv struct {
	Key, Value string
}

func NewTable(m int) *Table {
	return &Table{
		m:     m,
		table: make([][]kv, m),
	}
}

func (t *Table) hash(s string) int {
	h := fnv.New32()
	h.Write([]byte(s))
	return int(h.Sum32()) % t.m
}

func (t *Table) Get(key string) (string, bool) {
	i := t.hash(key)

	for j, kv := range t.table[i] {
		if key == kv.Key {
			return t.table[i][j].Value, true
		}
	}

	return "", false
}

func (t *Table) Insert(key, value string) {
	i := t.hash(key)

	for j, kv := range t.table[i] {
		if key == kv.Key {
			t.table[i][j].Value = value
			return
		}
	}

	t.table[i] = append(t.table[i], kv{
		Key:   key,
		Value: value,
	})
}

func IsPermutation(f string, s string) bool {

	if len(f) != len(s) {
		return false
	}

	table := NewTable(128)

	for _, char := range f {
		key := string(char)
		table.Insert(key, "")
	}

	for _, char := range s {
		key := string(char)
		_, count := table.Get(key)

		if !count {
			return false
		}
	}

	return true
}

func main() {
	var a = "ELisonMonteiroPassos"
	var b = "PsMnstoosnLeorsaiiEo"

	fmt.Printf("São permutações? %v\n", IsPermutation(a, b))

}
