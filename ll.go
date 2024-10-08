package toolbox

// Node represents a generic node in a linked list
type Node[T any] struct {
	Value T        `json:"el"`
	Next  *Node[T] `json:"next"`
}

// LinkedList is a generic linked list
type LinkedList[T any] struct {
	Head *Node[T] `json:"next"`
	Last *Node[T] `json:"-"` // Pointer to the last node
}

// Add adds a new node to the linked list
func (ll *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{Value: value}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Last = newNode
	} else {
		ll.Last.Next = newNode
		ll.Last = newNode
	}
}

// RemoveById removes a node from the list based on a function to match the ID
func (ll *LinkedList[T]) RemoveById(matchFunc func(T) bool) {
	if ll.Head == nil {
		return
	}

	// If the head matches, remove it
	if matchFunc(ll.Head.Value) {
		ll.Head = ll.Head.Next
		if ll.Head == nil {
			ll.Last = nil // List is now empty
		}
		return
	}

	// Traverse the list to find the node to remove
	current := ll.Head
	for current.Next != nil {
		if matchFunc(current.Next.Value) {
			// Check if it's the last node
			if current.Next.Next == nil {
				current.Next = nil
				ll.Last = current // Update Last pointer
			} else {
				current.Next = current.Next.Next
			}
			return
		}
		current = current.Next
	}
}

// Traverse is a utility function to apply a callback function to each node
func (ll *LinkedList[T]) Traverse(callback func(T)) {
	current := ll.Head
	for current != nil {
		callback(current.Value)
		current = current.Next
	}
}

func (ll *LinkedList[T]) AsSlice() []T {
	var roles []T
	ll.Traverse(func(role T) {
		roles = append(roles, role)
	})
	return roles
}

func (ll *LinkedList[T]) Reset() {
	l := &LinkedList[T]{}
	*ll = *l
}

