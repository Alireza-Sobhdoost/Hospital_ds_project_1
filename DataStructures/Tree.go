package DataStructures

import (
	"fmt"
	"strings"
)

// First define error types and interfaces
var (
	ErrDrugNotFound  = fmt.Errorf("drug not found")
	ErrDuplicateID   = fmt.Errorf("drug with this ID already exists")
	ErrDuplicateName = fmt.Errorf("drug with this name already exists")
)

// DrugNode represents a node in the binary search tree
type DrugNode struct {
	ID       string
	Name     string
	Price    float64
	Type     string
	Left     *DrugNode
	Right    *DrugNode
	Height   int
	Count    int
	Dose     string
	Variants *LinkedList // Changed from custom implementation to use native LinkedList
}

// DrugBST represents a binary search tree for drugs
type DrugBST struct {
	Root    *DrugNode
	Trie    *DrugTrie
	MaxHeap *MaxHeap
	MinHeap *MinHeap
}

// TrieNode represents a node in the Trie for drug name searches
type TrieNode struct {
	Character rune
	Children  *LinkedList
	IsEnd     bool
	Drug      *DrugNode
}

// DrugTrie represents a Trie for drug name searches
type DrugTrie struct {
	Root *TrieNode
}

// NewDrugBST creates a new binary search tree for drugs
func NewDrugBST() *DrugBST {
	return &DrugBST{
		Root:    nil,
		Trie:    NewDrugTrie(),
		MaxHeap: NewMaxHeap(),
		MinHeap: NewMinHeap(),
	}
}

// NewDrugTrie creates a new Trie for drug name searches
func NewDrugTrie() *DrugTrie {
	return &DrugTrie{
		Root: &TrieNode{
			Children: NewLinkedList(),
		},
	}
}

// getHeight returns the height of a node
func getHeight(node *DrugNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// getBalance returns the balance factor of a node
func getBalance(node *DrugNode) int {
	if node == nil {
		return 0
	}
	return getHeight(node.Left) - getHeight(node.Right)
}

// updateHeight updates the height of a node
func updateHeight(node *DrugNode) {
	if node == nil {
		return
	}
	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)
	if leftHeight > rightHeight {
		node.Height = leftHeight + 1
	} else {
		node.Height = rightHeight + 1
	}
}

// rotateRight performs a right rotation
func rotateRight(y *DrugNode) *DrugNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	updateHeight(y)
	updateHeight(x)

	return x
}

// rotateLeft performs a left rotation
func rotateLeft(x *DrugNode) *DrugNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	updateHeight(x)
	updateHeight(y)

	return y
}

// First define core search methods since other methods depend on them
func (bst *DrugBST) SearchByID(id string) (*DrugNode, error) {
	node := bst.searchNode(bst.Root, id)
	if node == nil {
		return nil, ErrDrugNotFound
	}
	return node, nil
}

func (bst *DrugBST) searchNode(node *DrugNode, id string) *DrugNode {
	if node == nil || node.ID == id {
		return node
	}
	if id < node.ID {
		return bst.searchNode(node.Left, id)
	}
	return bst.searchNode(node.Right, id)
}

// Then define insert/delete methods that use search
func (bst *DrugBST) Insert(id string, name string, price float64, drugType string, dose string) error {
	// Check for duplicates first
	if node, _ := bst.SearchByID(id); node != nil {
		if node.Name == name && node.Dose == dose && node.Type == drugType {
			node.Count += 1
			node.Price = price
			bst.MaxHeap.Delete(node)
			bst.MaxHeap.Insert(node)

			bst.MinHeap.Delete(node)
			bst.MinHeap.Insert(node)

			bst.Trie.Delete(node.Name)
			bst.Trie.Insert(node)
			fmt.Println(node.Count)
			fmt.Println("count added by 1")

			return nil
		}
		return ErrDuplicateID
	}

	// Create the node
	newNode := &DrugNode{
		ID:     id,
		Name:   name,
		Price:  price,
		Type:   drugType,
		Height: 1,
		Count:  1,
		Dose:   dose,
	}

	// Insert into BST
	bst.Root = bst.insertNode(bst.Root, id, name, price, drugType, dose)

	// Insert into Trie
	bst.Trie.Insert(newNode)

	// Insert into Heaps
	bst.MaxHeap.Insert(newNode)
	bst.MinHeap.Insert(newNode)

	return nil
}

// Delete and reinsert the node in the trie to update its information

// insertNode inserts a new node into the BST
func (bst *DrugBST) insertNode(node *DrugNode, id string, name string, price float64, drugType string, dose string) *DrugNode {
	if node == nil {
		return &DrugNode{
			ID:     id,
			Name:   name,
			Price:  price,
			Type:   drugType,
			Height: 1,
			Count:  1,
			Dose:   dose,
		}
	}

	if id < node.ID {
		node.Left = bst.insertNode(node.Left, id, name, price, drugType, dose)
	} else if id > node.ID {
		node.Right = bst.insertNode(node.Right, id, name, price, drugType, dose)
	} else {
		return node // Duplicate IDs not allowed
	}

	updateHeight(node)
	balance := getBalance(node)

	// Left Left Case
	if balance > 1 && id < node.Left.ID {
		return rotateRight(node)
	}

	// Right Right Case
	if balance < -1 && id > node.Right.ID {
		return rotateLeft(node)
	}

	// Left Right Case
	if balance > 1 && id > node.Left.ID {
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}

	// Right Left Case
	if balance < -1 && id < node.Right.ID {
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}

	return node
}

func (bst *DrugBST) Delete(id string) error {
	// First check if node exists
	node, err := bst.SearchByID(id)
	if err != nil {
		return err
	}

	// Remove from Heaps
	bst.MaxHeap.Delete(node)
	bst.MinHeap.Delete(node)

	// Remove from Trie
	bst.Trie.Delete(node.Name)

	// Then proceed with deletion
	bst.Root = bst.deleteNode(bst.Root, id)
	return nil
}

func (bst *DrugBST) deleteNode(node *DrugNode, id string) *DrugNode {
	if node == nil {
		return nil
	}

	if id < node.ID {
		node.Left = bst.deleteNode(node.Left, id)
	} else if id > node.ID {
		node.Right = bst.deleteNode(node.Right, id)
	} else {
		// Node to delete found
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Node has two children
		successor := bst.findMin(node.Right)
		node.ID = successor.ID
		node.Name = successor.Name
		node.Price = successor.Price
		node.Type = successor.Type
		node.Right = bst.deleteNode(node.Right, successor.ID)
	}

	updateHeight(node)
	return bst.balance(node)
}

// Insert adds a drug name to the Trie
func (t *DrugTrie) Insert(drug *DrugNode) {
	current := t.Root
	name := strings.ToLower(drug.Name)

	for _, char := range name {
		// charStr := string(char)
		found := false

		for node := current.Children.Head; node != nil; node = node.Next {
			childNode, ok := node.Data.(*TrieNode)
			if ok && childNode.Character == char {
				current = childNode
				found = true
				break
			}
		}

		if !found {
			newNode := &TrieNode{
				Character: char,
				Children:  NewLinkedList(),
			}
			current.Children.AddToEnd(newNode)
			current = newNode
		}
	}

	current.IsEnd = true
	current.Drug = drug
}

// SearchByName returns all drugs that match the given prefix
func (t *DrugTrie) SearchByName(prefix string) (*LinkedList, error) {
	prefix = strings.ToLower(prefix)
	current := t.Root
	results := NewLinkedList()

	for _, char := range prefix {
		found := false
		for node := current.Children.Head; node != nil; node = node.Next {
			childNode, ok := node.Data.(*TrieNode)
			if ok && childNode.Character == char {
				current = childNode
				found = true
				break
			}
		}
		if !found {
			return results, nil
		}
	}

	t.collectDrugsWithPrefix(current, results)
	return results, nil
}

// Helper method to collect all drugs with a given prefix
func (t *DrugTrie) collectDrugsWithPrefix(node *TrieNode, results *LinkedList) {
	if node.IsEnd && node.Drug != nil {
		results.AddToEnd(node.Drug)
	}

	for node := node.Children.Head; node != nil; node = node.Next {
		if childNode, ok := node.Data.(*TrieNode); ok {
			t.collectDrugsWithPrefix(childNode, results)
		}
	}
}

// AutoComplete returns all drugs that match the given prefix
func (t *DrugTrie) AutoComplete(prefix string) (*LinkedList, error) {
	results := t.autoCompleteHelper(prefix)
	if results.IsEmpty() {
		return nil, ErrDrugNotFound
	}
	return results, nil
}

// Helper method to collect all completions for a given prefix
func (t *DrugTrie) autoCompleteHelper(prefix string) *LinkedList {
	prefix = strings.ToLower(prefix)
	current := t.Root
	results := NewLinkedList()

	for _, char := range prefix {
		found := false
		for node := current.Children.Head; node != nil; node = node.Next {
			childNode, ok := node.Data.(*TrieNode)
			if ok && childNode.Character == char {
				current = childNode
				found = true
				break
			}
		}
		if !found {
			return results
		}
	}

	t.collectCompletions(current, results)
	return results
}

// Helper method to collect all completions from a given node
func (t *DrugTrie) collectCompletions(node *TrieNode, results *LinkedList) {
	if node.IsEnd {
		results.AddToEnd(node.Drug)
	}

	for node := node.Children.Head; node != nil; node = node.Next {
		if childNode, ok := node.Data.(*TrieNode); ok {
			t.collectCompletions(childNode, results)
		}
	}
}

// Delete removes a drug from the Trie
func (t *DrugTrie) Delete(name string) {
	name = strings.ToLower(name)
	t.deleteHelper(t.Root, name, 0)
}

func (t *DrugTrie) deleteHelper(node *TrieNode, name string, depth int) bool {
	if node == nil {
		return false
	}

	if depth == len(name) {
		if node.IsEnd {
			node.IsEnd = false
			node.Drug = nil
			return node.Children.IsEmpty()
		}
		return false
	}

	char := rune(name[depth])
	for child := node.Children.Head; child != nil; child = child.Next {
		childNode, ok := child.Data.(*TrieNode)
		if ok && childNode.Character == char {
			shouldDelete := t.deleteHelper(childNode, name, depth+1)
			if shouldDelete {
				node.Children.Delete(childNode)
				return node.Children.IsEmpty() && !node.IsEnd
			}
			break
		}
	}
	return false
}

// InOrderTraversal performs in-order traversal of the BST
func (bst *DrugBST) InOrderTraversal() []*DrugNode {
	var drugs []*DrugNode
	bst.inOrder(bst.Root, &drugs)
	return drugs
}

func (bst *DrugBST) inOrder(node *DrugNode, drugs *[]*DrugNode) {
	if node != nil {
		bst.inOrder(node.Left, drugs)
		*drugs = append(*drugs, node)
		bst.inOrder(node.Right, drugs)
	}
}

// InOrderTraversalByID performs in-order traversal of the BST and returns all drugs sorted by their IDs
func (bst *DrugBST) InOrderTraversalByID() []*DrugNode {
	var drugs []*DrugNode
	bst.inOrderByID(bst.Root, &drugs)
	return drugs
}

func (bst *DrugBST) inOrderByID(node *DrugNode, drugs *[]*DrugNode) {
	if node != nil {
		bst.inOrderByID(node.Left, drugs)
		*drugs = append(*drugs, node)
		bst.inOrderByID(node.Right, drugs)
	}
}

// GetCheapestDrug returns the cheapest drug in the BST
func (bst *DrugBST) GetCheapestDrug() *DrugNode {
	if bst.Root == nil {
		return nil
	}
	return bst.findMin(bst.Root)
}

func (bst *DrugBST) findMin(node *DrugNode) *DrugNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// GetMostExpensiveDrug returns the most expensive drug in the BST
func (bst *DrugBST) GetMostExpensiveDrug() *DrugNode {
	if bst.Root == nil {
		return nil
	}
	return bst.findMax(bst.Root)
}

func (bst *DrugBST) findMax(node *DrugNode) *DrugNode {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current
}

// MaxHeapNode represents a node in the max heap
type MaxHeapNode struct {
	Drug *DrugNode
	Next *MaxHeapNode
}

// MinHeapNode represents a node in the min heap
type MinHeapNode struct {
	Drug *DrugNode
	Next *MinHeapNode
}

// MaxHeap represents a max heap data structure
type MaxHeap struct {
	Root *MaxHeapNode
	Size int
}

// MinHeap represents a min heap data structure
type MinHeap struct {
	Root *MinHeapNode
	Size int
}

// NewMaxHeap creates a new max heap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{nil, 0}
}

// NewMinHeap creates a new min heap
func NewMinHeap() *MinHeap {
	return &MinHeap{nil, 0}
}

// Insert into max heap
// Insert into max heap
func (h *MaxHeap) Insert(drug *DrugNode) {
	node := &MaxHeapNode{Drug: drug}
	if h.Root == nil || h.Root.Drug.Price <= drug.Price {
		// Insert at the root if the list is empty or the new node has the highest price
		node.Next = h.Root
		h.Root = node
	} else {
		// Traverse to find the correct position
		current := h.Root
		for current.Next != nil && current.Next.Drug.Price > drug.Price {
			current = current.Next
		}
		// Insert the new node in the correct position
		node.Next = current.Next
		current.Next = node
	}
	h.Size++
}

// Insert into min heap
func (h *MinHeap) Insert(drug *DrugNode) {
	node := &MinHeapNode{Drug: drug}
	if h.Root == nil || h.Root.Drug.Price >= drug.Price {
		// Insert at the root if the list is empty or the new node has the lowest price
		node.Next = h.Root
		h.Root = node
	} else {
		// Traverse to find the correct position
		current := h.Root
		for current.Next != nil && current.Next.Drug.Price < drug.Price {
			current = current.Next
		}
		// Insert the new node in the correct position
		node.Next = current.Next
		current.Next = node
	}
	h.Size++
}

// ExtractMax removes and returns the maximum element
func (h *MaxHeap) ExtractMax() *DrugNode {
	if h.Root == nil {
		return nil
	}
	max := h.Root
	return max.Drug
}

// ExtractMin removes and returns the minimum element
func (h *MinHeap) ExtractMin() *DrugNode {
	if h.Root == nil {
		return nil
	}
	min := h.Root
	return min.Drug
}
// Delete removes a specific node from the max heap
func (h *MaxHeap) Delete(drug *DrugNode) {
	if h.Root == nil {
		fmt.Println("Heap is empty. Cannot delete.")
		return
	}

	// If the root is the node to delete
	if h.Root.Drug.ID == drug.ID {
		h.Root = h.Root.Next
		h.Size--
		fmt.Printf("Deleted node with price %d from MaxHeap\n", drug.Price)
		return
	}

	// Traverse to find the node to delete
	current := h.Root
	for current.Next != nil && current.Next.Drug.ID != drug.ID  {
		current = current.Next
	}

	// If the node was found, remove it
	if current.Next != nil {
		current.Next = current.Next.Next
		h.Size--
		fmt.Printf("Deleted node with price %d from MaxHeap\n", drug.Price)
	} else {
		fmt.Println("Node not found in MaxHeap.")
	}
}

// Delete removes a specific node from the min heap
func (h *MinHeap) Delete(drug *DrugNode) {
	if h.Root == nil {
		fmt.Println("Heap is empty. Cannot delete.")
		return
	}

	// If the root is the node to delete
	if h.Root.Drug.ID  == drug.ID  {
		h.Root = h.Root.Next
		h.Size--
		fmt.Printf("Deleted node with price %d from MinHeap\n", drug.Price)
		return
	}

	// Traverse to find the node to delete
	current := h.Root
	for current.Next != nil && current.Next.Drug.ID  != drug.ID  {
		current = current.Next
	}

	// If the node was found, remove it
	if current.Next != nil {
		current.Next = current.Next.Next
		h.Size--
		fmt.Printf("Deleted node with price %d from MinHeap\n", drug.Price)
	} else {
		fmt.Println("Node not found in MinHeap.")
	}
}

// AddVariant adds a variant to a drug's variant list
func (node *DrugNode) AddVariant(variant *DrugNode) {
	if node.Variants == nil {
		node.Variants = NewLinkedList()
	}
	node.Variants.AddToEnd(variant)
}

// Helper methods
func (bst *DrugBST) balance(node *DrugNode) *DrugNode {
	balance := getBalance(node)

	// Left heavy
	if balance > 1 {
		if getBalance(node.Left) < 0 {
			node.Left = rotateLeft(node.Left)
		}
		return rotateRight(node)
	}

	// Right heavy
	if balance < -1 {
		if getBalance(node.Right) > 0 {
			node.Right = rotateRight(node.Right)
		}
		return rotateLeft(node)
	}

	return node
}

// func (t *DrugTrie) collectDrugs(node *TrieNode, results *[]*DrugNode) {
// 	if node.IsEnd {
// 		*results = append(*results, node.Drug)
// 	}

// 	curr := node.Children.Head
// 	for curr != nil {
// 		if childNode, ok := curr.Data.(*TrieNode); ok {
// 			t.collectDrugs(childNode, results)
// 		}
// 		curr = curr.Next
// 	}
// }

// SuggestSimilarDrugs returns drug names similar to the given name
func (bst *DrugBST) SuggestSimilarDrugs(searchName string, maxDistance int) ([]*DrugNode, error) {
	suggestions := bst.suggestSimilarDrugsHelper(searchName, maxDistance)
	if len(suggestions) == 0 {
		return nil, ErrDrugNotFound
	}
	return suggestions, nil
}

func (bst *DrugBST) suggestSimilarDrugsHelper(searchName string, maxDistance int) []*DrugNode {
	suggestions := []*DrugNode{}
	allDrugs := bst.InOrderTraversal()

	searchName = strings.ToLower(searchName)
	for _, drug := range allDrugs {
		distance := levenshteinDistance(searchName, strings.ToLower(drug.Name))
		if distance <= maxDistance {
			suggestions = append(suggestions, drug)
		}
	}

	return suggestions
}

// levenshteinDistance calculates the edit distance between two strings
func levenshteinDistance(s1, s2 string) int {
	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}

	// Create matrix
	matrix := make([][]int, len(s1)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(s2)+1)
	}

	// Initialize first row and column
	for i := 0; i <= len(s1); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(s2); j++ {
		matrix[0][j] = j
	}

	// Fill in the rest of the matrix
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = min(
					matrix[i-1][j]+1,   // deletion
					matrix[i][j-1]+1,   // insertion
					matrix[i-1][j-1]+1, // substitution
				)
			}
		}
	}

	return matrix[len(s1)][len(s2)]
}

func min(numbers ...int) int {
	result := numbers[0]
	for _, num := range numbers {
		if num < result {
			result = num
		}
	}
	return result
}

// Modify SearchByType to return error
func (bst *DrugBST) SearchByType(drugType string) (*LinkedList, error) {
	result := NewLinkedList()
	bst.searchByTypeHelper(bst.Root, drugType, result)

	if result.IsEmpty() {
		return nil, ErrDrugNotFound
	}
	return result, nil
}

// Helper function for SearchByType
func (bst *DrugBST) searchByTypeHelper(node *DrugNode, drugType string, list *LinkedList) {
	if node == nil {
		return
	}

	// In-order traversal
	bst.searchByTypeHelper(node.Left, drugType, list)

	// If current node matches the type, add it to the list
	if strings.ToLower(node.Type) == strings.ToLower(drugType) {
		list.AddToEnd(node)
	}

	bst.searchByTypeHelper(node.Right, drugType, list)
}

// SearchByPriceRange returns all drugs with prices between minPrice and maxPrice
func (bst *DrugBST) SearchByPriceRange(minPrice, maxPrice float64) (*LinkedList, error) {
	result := NewLinkedList()
	bst.searchByPriceRangeHelper(bst.MinHeap.Root, minPrice, maxPrice, result)

	if result.IsEmpty() {
		return nil, ErrDrugNotFound
	}
	return result, nil
}

// Helper function for SearchByPriceRange
func (bst *DrugBST) searchByPriceRangeHelper(node *MinHeapNode, minPrice, maxPrice float64, list *LinkedList) {
	for node != nil {
		if node.Drug.Price >= minPrice && node.Drug.Price <= maxPrice {
			list.AddToEnd(node.Drug)
		}
		node = node.Next
	}
}

// CountAllDrugs returns the total number of drugs in the BST
func (bst *DrugBST) CountAllDrugs() int {
	return bst.countDrugs(bst.Root)
}

func (bst *DrugBST) countDrugs(node *DrugNode) int {
	if node == nil {
		return 0
	}
	return node.Count + bst.countDrugs(node.Left) + bst.countDrugs(node.Right)
}

// GetBSTDepth returns the depth of the BST
func (bst *DrugBST) GetBSTDepth() int {
	return bst.getDepth(bst.Root)
}

func (bst *DrugBST) getDepth(node *DrugNode) int {
	if node == nil {
		return 0
	}
	leftDepth := bst.getDepth(node.Left)
	rightDepth := bst.getDepth(node.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// More methods to be implemented...
