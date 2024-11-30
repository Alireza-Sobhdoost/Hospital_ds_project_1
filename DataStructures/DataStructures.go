package DataStructures

import (
	"fmt"
	// "project_1/Entities"
)

type Node struct {
	Data interface{} 
	Next *Node       
	Prev *Node       
	Down *Node       
}

// this LinkedList here represents the General two way (dublle) linked list
type LinkedList struct {
	Head *Node
	Tail *Node 
}



// after initiation of our linked list Ds its time to implemet its methods

func (list *LinkedList) AddToEnd(data interface{}) {
	newNode := &Node{Data: data}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.Next = newNode
		newNode.Prev = list.Tail
		list.Tail = newNode
	}
}

func (list *LinkedList) AddToStart(data interface{}) {
	newNode := &Node{Data: data}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
	}
}

func (ll *LinkedList) Find_by_index(index int , len int)(*Node) {
	if index < 0 {
		fmt.Println("Invalid index")
		return nil
	} else if index > len {
		fmt.Println("Invalid index2")

		return nil
	} 
    currentNode := ll.Head
	count := 0
	for count < index {

		currentNode = currentNode.Next
		fmt.Println("currentNode" , currentNode)
		count += 1
	}
	return currentNode  // Return node if found

}



func (node *Node) AddDown(data interface{}) {
	newNode := &Node{Data: data}
	if node.Down == nil {
		node.Down = newNode
	} else {
		current := node.Down
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
		newNode.Prev = current
	}
}



func (list *LinkedList) Remove(data interface{}) {
	if list.Head == nil {
		return
	}
	current := list.Head
	for current != nil {
		if current.Data == data {
		
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				list.Head = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				list.Tail = current.Prev 
			}
			break
		}
		current = current.Next
	}
}

func (list *LinkedList) Display() {
	current := list.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Data)
		down := current.Down
		for down != nil {
			fmt.Printf("[%v] -> ", down.Data)
			down = down.Next
		}
		fmt.Print("nil  ")
		current = current.Next
	}
	fmt.Println("nil")
}


// implementing stack DS

type Stack struct {
	items []interface{}
}

// after initiation of our stack Ds its time to implemet its methods

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	lastIndex := len(s.items) - 1
	element := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return element, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func NewStack() *Stack {
	return &Stack{
		items : [] interface{}{},
	}
}

// implementing queue DS

type Queue struct {
	items []interface{}
}

// after initiation of our queue Ds its time to implemet its methods

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	element := q.items[0]
	q.items = q.items[1:]
	return element, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if len(q.items) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}


type PriorityQueue struct {
	Heap []interface{}
	less func(a, b interface{}) bool // Comparison function
}

// NewPriorityQueue creates a new PriorityQueue
func NewPriorityQueue(lessFunc func(a, b interface{}) bool) *PriorityQueue {
	return &PriorityQueue{
		Heap: []interface{}{},
		less: lessFunc,
	}
}

// Push adds an element to the priority queue
func (pq *PriorityQueue) Push(Value interface{}) {
	pq.Heap = append(pq.Heap, Value)
	pq.upHeap(len(pq.Heap) - 1)
}

// Pop removes and returns the smallest element (root) from the priority queue
func (pq *PriorityQueue) Pop() (interface{}, error) {
	if len(pq.Heap) == 0 {
		return nil, fmt.Errorf("priority queue is empty")
	}

	// Swap the root with the last element and remove the last element
	root := pq.Heap[0]
	pq.Heap[0] = pq.Heap[len(pq.Heap)-1]
	pq.Heap = pq.Heap[:len(pq.Heap)-1]

	// Restore the Heap property
	pq.downHeap(0)

	return root, nil
}

// Peek returns the smallest element without removing it
func (pq *PriorityQueue) Peek() (interface{}, error) {
	if len(pq.Heap) == 0 {
		return nil, fmt.Errorf("priority queue is empty")
	}
	return pq.Heap[0], nil
}

// IsEmpty checks if the priority queue is empty
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.Heap) == 0
}

// upHeap restores the Heap property by moving the element at index up
func (pq *PriorityQueue) upHeap(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if !pq.less(pq.Heap[index], pq.Heap[parent]) {
			break
		}
		pq.Heap[index], pq.Heap[parent] = pq.Heap[parent], pq.Heap[index]
		index = parent
	}
}

// downHeap restores the Heap property by moving the element at index down
func (pq *PriorityQueue) downHeap(index int) {
	lastIndex := len(pq.Heap) - 1
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild <= lastIndex && pq.less(pq.Heap[leftChild], pq.Heap[smallest]) {
			smallest = leftChild
		}
		if rightChild <= lastIndex && pq.less(pq.Heap[rightChild], pq.Heap[smallest]) {
			smallest = rightChild
		}
		if smallest == index {
			break
		}
		pq.Heap[index], pq.Heap[smallest] = pq.Heap[smallest], pq.Heap[index]
		index = smallest
	}
}

// Define the structure for a Hash Map
type HashMap struct {
	Buckets [][]KeyValue
	size    int
	count   int // To keep track of number of elements in the map
}

// Define a structure for the key-Value pair
type KeyValue struct {
	key   string
	Value interface{}
}

// Create a new HashMap
func NewHashMap(size int) *HashMap {
	return &HashMap{
		Buckets: make([][]KeyValue, size),
		size:    size,
		count:   0,
	}
}

// Hash function to compute the index
func (hm *HashMap) hash(key string) int {
	hashValue := 0
	for i := 0; i < len(key); i++ {
		hashValue += int(key[i])
	}
	return hashValue % hm.size
}

// Resize the hash map (double the size)
// Resize the hash map (double the size)
func (hm *HashMap) resize() {
	// Double the size of the hash map
	newSize := hm.size * 2
	newBuckets := make([][]KeyValue, newSize)

	// Rehash and insert existing elements into the new bucket array
	for _, bucket := range hm.Buckets {
		for _, kv := range bucket {
			// Apply a hash function to the key (instead of trying kv.key % newSize)
			index := hm.hash(kv.key) % newSize
			newBuckets[index] = append(newBuckets[index], kv)
		}
	}

	// Update the hash map with the resized Buckets
	hm.Buckets = newBuckets
	hm.size = newSize
}


// Insert a key-Value pair
func (hm *HashMap) Insert(key string, Value interface{}) {
	// Check if resizing is needed
	if float64(hm.count)/float64(hm.size) > 0.80 {
		hm.resize()
	}

	// Insert the key-Value pair
	index := hm.hash(key)
	hm.Buckets[index] = append(hm.Buckets[index], KeyValue{key, Value})
	hm.count++
}

// Retrieve a Value for a given key
func (hm *HashMap) Get(key string) (interface{}, bool) {
	index := hm.hash(key)
	for _, kv := range hm.Buckets[index] {
		if kv.key == key {
			return kv.Value, true
		}
	}
	return 0, false // Key not found
}

func (hm *HashMap) GetByID(key string) (interface{}, bool) {
	// Check if the key is "Doctors" for 2D HashMap logic
	if key == "Doctors" {
		// Get the doctors map from the outer HashMap
		index := hm.hash(key)
		for _, kv := range hm.Buckets[index] {
			if kv.key == key {
				// Type assertion to get the inner doctors map
				if innerMap, ok := kv.Value.(*HashMap); ok {
					return innerMap, true
				}
			}
		}
	}

	// If key is a National ID, look for it directly in the "Doctors" inner map
	for _, bucket := range hm.Buckets {
		for _, kv := range bucket {
			if innerMap, ok := kv.Value.(*HashMap); ok {
				// Search for the doctor by National ID in the inner map
				index := innerMap.hash(key)
				for _, innerKV := range innerMap.Buckets[index] {
					if innerKV.key == key {
						return innerKV.Value, true
					}
				}
			}
		}
	}

	// If the key is not found
	return nil, false
}

func (hm *HashMap) GetRecursive(key string) (interface{}, bool) {
	// Iterate through the Buckets of the current HashMap
	for _, bucket := range hm.Buckets {
		for _, kv := range bucket {
			// If the key matches, return the Value
			if kv.key == key {
				return kv.Value, true
			}

			// If the Value is another HashMap, search recursively
			if innerMap, ok := kv.Value.(*HashMap); ok {
				result, found := innerMap.GetRecursive(key)
				if found {
					return result, true
				}
			}
		}
	}

	// If not found, return false
	return nil, false
}

// Delete a key-Value pair
func (hm *HashMap) Delete(key string) {
	index := hm.hash(key)
	for i, kv := range hm.Buckets[index] {
		if kv.key == key {
			hm.Buckets[index] = append(hm.Buckets[index][:i], hm.Buckets[index][i+1:]...)
			hm.count--
			return
		}
	}
}

// Display the entire hash map
func (hm *HashMap) Display() {
	for i, bucket := range hm.Buckets {
		if len(bucket) > 0 {
			fmt.Printf("Bucket %d: ", i)
			for _, kv := range bucket {
				fmt.Printf("[%s: %d] ", kv.key, kv.Value)
			}
			fmt.Println()
		}
	}
}



// DataBase := DataStructures.NewHashMap(100)
// user , err := Auth.Signup("1", "John", "Doe", "password", "Patient", 20)
// if err != nil {
// 	log.Fatal(err)
// }
// p1 := user.(*Entities.Patient)
// fmt.Println(p1.FirstName)
// DataBase.Insert(p1.ID, p1)
// if err != nil {
// user2 , role ,err2 := Auth.Login(*DataBase, "1", "password")
// 	log.Fatal(err2)
// }