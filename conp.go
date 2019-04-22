package conp

type Fetcher struct {
	Key   string
	Value func() interface{}
}

func (f *Fetcher) Unpack() (string, func() interface{}) {
	return f.Key, f.Value
}

type Bucket interface {
	Fetch(string, func() interface{}) (interface{}, bool)
}
