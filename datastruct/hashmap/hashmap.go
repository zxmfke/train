package hashmap

type HSMap struct {
	m        HSMapValue
	hashFunc HashFun
}

type HashFun = func(string) string

func HF(f func(string) string) HashFun {
	return HashFun(f)
}

func NewHSMap(hf HashFun, hsmv HSMapValue) *HSMap {
	hsm := new(HSMap)
	hsm.m = hsmv

	hsm.hashFunc = defaultHashFunc
	if hf != nil {
		hsm.hashFunc = hf
	}

	return hsm
}

func defaultHashFunc(key string) string {
	return key
}

func (h *HSMap) Set(key string, value interface{}) {
	hsKey := h.hashFunc(key)
	h.m.Set(hsKey, value)
}

func (h *HSMap) Get(key string) (interface{}, bool) {
	return h.m.Get(h.hashFunc(key))
}

func (h *HSMap) Delete(key string) {
	h.m.Delete(h.hashFunc(key))
}
