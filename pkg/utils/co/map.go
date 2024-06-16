package co

type Map[K comparable, V any] struct {
	data map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V),
	}
}

func (m *Map[K, V]) Set(k K, v V) {
	m.data[k] = v
}

func (m *Map[K, V]) Get(k K) (v V, ok bool) {
	v, ok = m.data[k]
	return
}

func (m *Map[K, V]) Del(k K) {
	delete(m.data, k)
}

func (m *Map[K, V]) Len() int {
	return len(m.data)
}

func (m *Map[K, V]) Each(f func(K, V) error) error {
	for k, v := range m.data {
		if e := f(k, v); e != nil {
			return e
		}
	}
	return nil
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, m.Len())
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, m.Len())
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}
