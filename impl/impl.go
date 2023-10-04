package impl

type Partitioner interface {
	Get() []byte
	Set([]byte)
}

func (dh *DataHolder) Get() []byte {
	return dh.__originalData
}

func (dh *DataHolder) Set(b []byte) {
	dh.__originalData = b
}

type DataHolder struct {
	__originalData []byte
}
