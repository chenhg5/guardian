package guardian

type Vars map[string]interface{}

func (v Vars) Get(key string) interface{} {
	return v[key]
}

func (v Vars) Add(key string, value interface{}) {
	v[key] = value
}