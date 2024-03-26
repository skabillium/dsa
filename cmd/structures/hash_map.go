package structures

type HashMap struct {
	Length int
	arr    []any
}

func NewHashMap() *HashMap {
	return &HashMap{arr: []any{}, Length: 0}
}

func (m *HashMap) Get(key string) any { return 0 }

func (m *HashMap) Set(key string, value any) {}

func (m *HashMap) Has(key string) bool { return false }
