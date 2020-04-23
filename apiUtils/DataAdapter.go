package apiUtils

type DataAdapter struct {
	data interface{}
}

func (d DataAdapter) ToString() string {
	return d.data.(string)
}

func (d DataAdapter) ToFloat() float64 {
	return d.data.(float64)
}
