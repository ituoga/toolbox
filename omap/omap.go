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

type Map[K comparable, T any] struct {
	data map[K]T
	keys []K
}

func New[K comparable, T any]() *Map[K, T] {
	return &Map[K, T]{
		data: make(map[K]T),
	}
}
func (m *Map[K, T]) Set(key K, value T) {
	if _, exists := m.data[key]; !exists {
		m.keys = append(m.keys, key)
	}
	m.data[key] = value
}

func (m *Map[K, T]) Get(key K) (T, bool) {
	value, exists := m.data[key]
	return value, exists
}

func (m *Map[K, T]) Delete(key K) {
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

func (m *Map[K, T]) Keys() []K {
	return m.keys
}

func (m *Map[K, T]) Values() []T {
	values := make([]T, 0, len(m.data))
	for _, key := range m.keys {
		values = append(values, m.data[key])
	}
	return values
}

func (m *Map[K, T]) Len() int {
	return len(m.data)
}

func (m *Map[K, T]) Clear() {
	m.data = make(map[K]T)
	m.keys = []K{}
}

func (m *Map[K, T]) ForEach(fn func(key K, value T)) {
	for _, key := range m.keys {
		value := m.data[key]
		fn(key, value)
	}
}
