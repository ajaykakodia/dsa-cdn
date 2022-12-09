package main

import (
	"fmt"
	"hash/fnv"
)

type MapNode struct {
	key   string
	value string
	next  *MapNode
}

type HMap struct {
	bucketArr  []*MapNode
	count      int
	bucketSize int
}

func NewMap(size int) *HMap {
	return &HMap{
		bucketArr:  make([]*MapNode, size),
		bucketSize: size,
	}
}

func (m *HMap) IsEmpty() bool {
	return m.count == 0
}

func (m *HMap) Size() int {
	return m.count
}

func (m *HMap) Insert(key, value string) {
	hCode := getHash(key)

	indexToStore := hCode % m.bucketSize

	head := m.bucketArr[indexToStore]

	for head != nil {
		if head.key == key {
			head.value = value
			return
		}
		head = head.next
	}

	newMapNode := &MapNode{
		key:   key,
		value: value,
		next:  m.bucketArr[indexToStore],
	}

	m.bucketArr[indexToStore] = newMapNode

	m.count = m.count + 1

	loadFactor := float64(m.count) / float64(m.bucketSize)
	fmt.Println("LoadFactor: ", loadFactor)
	if loadFactor >= 0.8 {
		distributeLoad(m)
	}
}

func distributeLoad(m *HMap) {
	dataArray := m.bucketArr
	m.bucketSize = 2 * m.bucketSize
	m.bucketArr = make([]*MapNode, m.bucketSize)
	m.count = 0
	for _, val := range dataArray {
		if val == nil {
			continue
		}
		head := val
		for head != nil {
			m.Insert(head.key, head.value)
			head = head.next
		}
	}
}

func (m *HMap) Find(key string) (string, bool) {
	hCode := getHash(key)
	indexToStore := hCode % m.bucketSize

	if m.bucketArr[indexToStore] == nil {
		return "", false
	}

	head := m.bucketArr[indexToStore]

	for head != nil {
		if head.key == key {
			return head.value, true
		}
		head = head.next
	}
	return "", false
}

func (m *HMap) Delete(key string) bool {
	hCode := getHash(key)
	indexToStore := hCode % m.bucketSize
	head := m.bucketArr[indexToStore]
	var prevEle, currentHead *MapNode
	currentHead = head
	for currentHead != nil {
		if currentHead.key == key {
			if prevEle == nil {
				m.bucketArr[indexToStore] = nil
			} else {
				prevEle.next = currentHead.next
			}
			m.count = m.count - 1
			return true
		}
		prevEle = currentHead
		currentHead = currentHead.next
	}
	return false
}

func getHash(key string) int {
	hsh := fnv.New32()
	hsh.Write([]byte(key))
	return int(hsh.Sum32())
}
