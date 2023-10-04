package impl

type SimplifiedYamlModel interface {
	UnfoldYaml() []byte
	Set([]byte, interface{})
}

func (dh *DataHolder) UnfoldYaml() []byte {
	return merge(dh.__originalData, dh.__var)
}

func merge(originalData []byte, i interface{}) []byte {
	return nil
}

func (dh *DataHolder) Set(b []byte, i interface{}) {
	dh.__originalData = b
	dh.__var = i
}

type DataHolder struct {
	__originalData []byte
	__var          interface{}
}
