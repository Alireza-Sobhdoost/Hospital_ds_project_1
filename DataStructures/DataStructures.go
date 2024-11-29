package DataStructures

import "fmt"

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
	heap []interface{}
	less func(a, b interface{}) bool // Comparison function
}

// NewPriorityQueue creates a new PriorityQueue
func NewPriorityQueue(lessFunc func(a, b interface{}) bool) *PriorityQueue {
	return &PriorityQueue{
		heap: []interface{}{},
		less: lessFunc,
	}
}

// Push adds an element to the priority queue
func (pq *PriorityQueue) Push(value interface{}) {
	pq.heap = append(pq.heap, value)
	pq.upHeap(len(pq.heap) - 1)
}

// Pop removes and returns the smallest element (root) from the priority queue
func (pq *PriorityQueue) Pop() (interface{}, error) {
	if len(pq.heap) == 0 {
		return nil, fmt.Errorf("priority queue is empty")
	}

	// Swap the root with the last element and remove the last element
	root := pq.heap[0]
	pq.heap[0] = pq.heap[len(pq.heap)-1]
	pq.heap = pq.heap[:len(pq.heap)-1]

	// Restore the heap property
	pq.downHeap(0)

	return root, nil
}

// Peek returns the smallest element without removing it
func (pq *PriorityQueue) Peek() (interface{}, error) {
	if len(pq.heap) == 0 {
		return nil, fmt.Errorf("priority queue is empty")
	}
	return pq.heap[0], nil
}

// IsEmpty checks if the priority queue is empty
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.heap) == 0
}

// upHeap restores the heap property by moving the element at index up
func (pq *PriorityQueue) upHeap(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if !pq.less(pq.heap[index], pq.heap[parent]) {
			break
		}
		pq.heap[index], pq.heap[parent] = pq.heap[parent], pq.heap[index]
		index = parent
	}
}

// downHeap restores the heap property by moving the element at index down
func (pq *PriorityQueue) downHeap(index int) {
	lastIndex := len(pq.heap) - 1
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild <= lastIndex && pq.less(pq.heap[leftChild], pq.heap[smallest]) {
			smallest = leftChild
		}
		if rightChild <= lastIndex && pq.less(pq.heap[rightChild], pq.heap[smallest]) {
			smallest = rightChild
		}
		if smallest == index {
			break
		}
		pq.heap[index], pq.heap[smallest] = pq.heap[smallest], pq.heap[index]
		index = smallest
	}
}

// Define the structure for a Hash Map
type HashMap struct {
	buckets [][]KeyValue
	size    int
	count   int // To keep track of number of elements in the map
}

// Define a structure for the key-value pair
type KeyValue struct {
	key   string
	value interface{}
}

// Create a new HashMap
func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([][]KeyValue, size),
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
	for _, bucket := range hm.buckets {
		for _, kv := range bucket {
			// Apply a hash function to the key (instead of trying kv.key % newSize)
			index := hm.hash(kv.key) % newSize
			newBuckets[index] = append(newBuckets[index], kv)
		}
	}

	// Update the hash map with the resized buckets
	hm.buckets = newBuckets
	hm.size = newSize
}


// Insert a key-value pair
func (hm *HashMap) Insert(key string, value interface{}) {
	// Check if resizing is needed
	if float64(hm.count)/float64(hm.size) > 0.80 {
		hm.resize()
	}

	// Insert the key-value pair
	index := hm.hash(key)
	hm.buckets[index] = append(hm.buckets[index], KeyValue{key, value})
	hm.count++
}

// Retrieve a value for a given key
func (hm *HashMap) Get(key string) (interface{}, bool) {
	index := hm.hash(key)
	for _, kv := range hm.buckets[index] {
		if kv.key == key {
			return kv.value, true
		}
	}
	return 0, false // Key not found
}

func (hm *HashMap) GetByID(key string) (interface{}, bool) {
	// Check if the key is "Doctors" for 2D HashMap logic
	if key == "Doctors" {
		// Get the doctors map from the outer HashMap
		index := hm.hash(key)
		for _, kv := range hm.buckets[index] {
			if kv.key == key {
				// Type assertion to get the inner doctors map
				if innerMap, ok := kv.value.(*HashMap); ok {
					return innerMap, true
				}
			}
		}
	}

	// If key is a National ID, look for it directly in the "Doctors" inner map
	for _, bucket := range hm.buckets {
		for _, kv := range bucket {
			if innerMap, ok := kv.value.(*HashMap); ok {
				// Search for the doctor by National ID in the inner map
				index := innerMap.hash(key)
				for _, innerKV := range innerMap.buckets[index] {
					if innerKV.key == key {
						return innerKV.value, true
					}
				}
			}
		}
	}

	// If the key is not found
	return nil, false
}
// Delete a key-value pair
func (hm *HashMap) Delete(key string) {
	index := hm.hash(key)
	for i, kv := range hm.buckets[index] {
		if kv.key == key {
			hm.buckets[index] = append(hm.buckets[index][:i], hm.buckets[index][i+1:]...)
			hm.count--
			return
		}
	}
}

// Display the entire hash map
func (hm *HashMap) Display() {
	for i, bucket := range hm.buckets {
		if len(bucket) > 0 {
			fmt.Printf("Bucket %d: ", i)
			for _, kv := range bucket {
				fmt.Printf("[%s: %d] ", kv.key, kv.value)
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