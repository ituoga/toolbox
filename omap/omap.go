package omap

// Map is a simple ordered map implementation in Go.
// It maintains the order of keys as they are added, allowing iteration in insertion order.
// It supports basic operations like Set, Get, Delete, Keys, Values, Len, Clear, and ForEach.
// The type parameter T allows the map to hold values of any type.
// Example usage:
//
//	m := omap.New[int]()
//	m.Set("one", 1)
//	m.Set("two", 2)
//	fmt.Println(m.Get("one")) // Output: 1, true
//	fmt.Println(m.Keys())     // Output: ["one", "two"]
//	fmt.Println(m.Values())   // Output: [1, 2]
//	m.ForEach(func(key string, value int) {
//	    fmt.Printf("%s: %d\n", key, value)
//	})
//	m.Delete("one")
//	fmt.Println(m.Keys())     // Output: ["two"]
//	fmt.Println(m.Len())      // Output: 1
//	m.Clear()
//	fmt.Println(m.Len())      // Output: 0
//	m.Set("three", 3)
//	fmt.Println(m.Keys())     // Output: ["three"]
//	fmt.Println(m.Values())   // Output: [3]

type Map[T any] struct {
	data map[string]T
	keys []string
}

func New[T any]() *Map[T] {
	return &Map[T]{
		data: make(map[string]T),
	}
}
func (m *Map[T]) Set(key string, value T) {
	if _, exists := m.data[key]; !exists {
		m.keys = append(m.keys, key)
	}
	m.data[key] = value
}

func (m *Map[T]) Get(key string) (T, bool) {
	value, exists := m.data[key]
	return value, exists
}

func (m *Map[T]) Delete(key string) {
	if _, exists := m.data[key]; exists {
		delete(m.data, key)
		for i, k := range m.keys {
			if k == key {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
	}
}

func (m *Map[T]) Keys() []string {
	return m.keys
}

func (m *Map[T]) Values() []T {
	values := make([]T, 0, len(m.data))
	for _, key := range m.keys {
		values = append(values, m.data[key])
	}
	return values
}

func (m *Map[T]) Len() int {
	return len(m.data)
}

func (m *Map[T]) Clear() {
	m.data = make(map[string]T)
	m.keys = []string{}
}

func (m *Map[T]) ForEach(fn func(key string, value T)) {
	for _, key := range m.keys {
		value := m.data[key]
		fn(key, value)
	}
}
