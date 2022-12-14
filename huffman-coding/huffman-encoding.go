package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var charMap map[string]int
var charToCode map[string]string
var codeToChar map[string]string

type HeapNode struct {
	value  string
	freq   int
	index  int
	lChild *HeapNode
	rChild *HeapNode
}

// implement heap interface for creating heap
type charHeap []*HeapNode

func (h charHeap) Len() int {
	return len(h)
}

func (h charHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h charHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *charHeap) Push(x any) {
	n := len(*h)
	hn := x.(*HeapNode)
	hn.index = n
	*h = append(*h, hn)
}

func (h *charHeap) Pop() any {
	old := *h
	n := len(old)
	hn := old[n-1]
	old[n-1] = nil // avoid memory leak
	hn.index = -1  // for safety
	*h = old[0 : n-1]
	return hn
}

// update modifies the priority and value of an Item in the queue.
func (h *charHeap) update(hn *HeapNode, value string, freq int) {
	hn.value = value
	hn.freq = freq
	heap.Fix(h, hn.index)
}

// createCharMap => create char and frequency map for file bytes/ chars
func createCharMap(file []byte) {
	charMap = make(map[string]int)
	for _, val := range file {
		charMap[string(val)] = charMap[string(val)] + 1
	}
}

// initHeap => create heap for charMap characters
func initHeap() *charHeap {
	var chHeap charHeap
	var i int
	// display frequency of chars and adding items in char min heap
	for key, val := range charMap {
		hNode := &HeapNode{
			value: key,
			freq:  val,
			index: i,
		}
		i++
		chHeap = append(chHeap, hNode)
	}
	heap.Init(&chHeap)
	return &chHeap
}

// createBinaryTree => create binary tree by popping/ adding elements in heap
func createBinaryTree(chHeap *charHeap) {
	// last element remains for binary tree root
	for len(*chHeap) > 1 {
		hNode1 := heap.Pop(chHeap).(*HeapNode)
		hNode2 := heap.Pop(chHeap).(*HeapNode)
		newHNode := &HeapNode{
			freq:   hNode1.freq + hNode2.freq,
			lChild: hNode1,
			rChild: hNode2,
		}
		heap.Push(chHeap, newHNode)
	}
}

func HuffmanEncoding(path string) {
	completeFilePath := path + "sampleText.txt"
	// open and read file
	file, err := os.ReadFile(completeFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create map
	createCharMap(file)

	// initialize Heap
	chHeap := initHeap()

	// create binary tree
	createBinaryTree(chHeap)

	// create code
	rootNode := heap.Pop(chHeap).(*HeapNode)
	charToCode = make(map[string]string)
	codeToChar = make(map[string]string)
	createCodes(rootNode, "")

	// create code text
	codeText := createCodeTextForFile(file)

	// embed padding
	completeText := embedPadding(codeText)

	// get bytes and write file to path
	writeDataToFile(completeText, path)

	// decompress file
	decompressFile(path)

}

// decompress file and sve into the folder
func decompressFile(path string) {
	// getting bin file to precess
	file, err := os.ReadFile(path + "sampleText.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	var txt string
	// get bit string from bytes array
	for _, val := range file {
		bitxt := fmt.Sprintf("%08b", val)
		txt = txt + bitxt
	}

	// get embedded padding info
	embeddedPaddingInfo := txt[:8]
	// get embedded info in integer
	addedZero, err := strconv.ParseInt(embeddedPaddingInfo, 2, 8)
	if err != nil {
		fmt.Println(err)
	}
	// remove embedded prefix embedded info and padded zero
	txt = txt[8 : len(txt)-int(addedZero)]

	// get string from bits processing through char stores in map for codes
	var codeText, fileData string
	for _, val := range txt {
		codeText = codeText + string(val)
		if char, okk := codeToChar[codeText]; okk {
			fileData = fileData + char
			codeText = ""
		}
	}

	outPutFilePath := path + "decompressed.txt"
	// checking status if file exits ot not
	_, err = os.Stat(outPutFilePath)
	if err != nil {
		// if file not exists, creating file for the same
		dest, err := os.Create(outPutFilePath)
		if err != nil {
			fmt.Println(err)
			dest.Close()
			return
		}
		dest.Close()
	}
	// open file to write byte array into it
	dFile, err := os.OpenFile(outPutFilePath, os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0660)
	defer dFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// write data to the file
	dFile.WriteString(fileData)
}

// createCodeTextForFile => create string of codes for compress file
func createCodeTextForFile(file []byte) string {
	var codeText string
	for _, val := range file {
		if code, ok := charToCode[string(val)]; ok {
			codeText = codeText + code
		}
	}
	return codeText
}

// writeDataToFile => Write file for provided bit strings
func writeDataToFile(codeText, path string) {
	if codeText == "" {
		return
	}
	var byteArr []byte
	// get bytes and added into byte array - we have complete set of 8 bit bytes
	for i := 0; i < len(codeText); i += 8 {
		str := codeText[i : i+8]
		// convert into byte
		val, err := strconv.ParseUint(str, 2, 8)
		if err != nil {
			fmt.Println(err)
			return
		}
		byteArr = append(byteArr, byte(val))
	}
	outPutFilePath := path + "sampleText.bin"
	// checking status if file exits ot not
	_, err := os.Stat(outPutFilePath)
	if err != nil {
		// if file not exists, creating file for the same
		dest, err := os.Create(outPutFilePath)
		if err != nil {
			fmt.Println(err)
			dest.Close()
			return
		}
		dest.Close()
	}
	// open file to write byte aarray into it
	file, err := os.OpenFile(outPutFilePath, os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0660)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Write(byteArr)
}

// embedPadding => adding zeros in last for complete byte set of 8 bits
func embedPadding(codeText string) string {
	// checking for mod to calculate zeros to add
	zeroToAdd := len(codeText) % 8
	zero := ""
	// get zeros to add
	for i := 0; i < zeroToAdd; i++ {
		zero = zero + "0"
	}
	codeText = codeText + zero
	// get one byte to know about added zeros in last
	preBits := fmt.Sprintf("%08b", zeroToAdd)
	// add bits to the codeText
	return preBits + codeText
}

// createCodes => create char to code map or code to char map
func createCodes(root *HeapNode, prefixCode string) {
	if root == nil {
		return
	}
	// if get leaf node
	if root.value != "" {
		charToCode[root.value] = prefixCode
		codeToChar[prefixCode] = root.value
		return
	}
	// sending 0 to left child to create code
	createCodes(root.lChild, prefixCode+"0")
	// sending 1 to right child to create code
	createCodes(root.rChild, prefixCode+"1")
}
