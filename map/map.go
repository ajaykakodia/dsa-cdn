package main

import (
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

func (m *HMap) Insert(key, value string) {
	hCode := getHash(key)

	indexToStore := hCode % m.bucketSize

	head := m.bucketArr[indexToStore]

	for head != nil {
		if head.key == key {
			head.value = value
			return
		}
	}

	newMapNode := &MapNode{
		key:   key,
		value: value,
		next:  m.bucketArr[indexToStore],
	}

	m.bucketArr[indexToStore] = newMapNode
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
