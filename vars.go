package guardian

type Vars map[string]string

func (v Vars) Get(key string) string {
	return v[key]
}

func (v Vars) Add(key string, value string) {
	v[key] = value
}