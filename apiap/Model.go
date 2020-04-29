package apiap

type Model struct {
	TYPE     string
	Method   string
	Url      string
	UrlParam string
	Body     interface{}
}

func typeAB(methodType string, method string, url string, urlParam string) Model {
	model := Model{}
	model.TYPE = methodType
	model.Method = method
	model.Url = url
	model.UrlParam = urlParam

	return model
}

func TypeA(method string, url string, urlParam string) Model {
	return typeAB("A", method, url, urlParam)
}

func TypeB(method string, url string, urlParam string) Model {
	return typeAB("B", method, url, urlParam)
}

func TypeC(method string, url string, body interface{}) Model {
	model := Model{}
	model.TYPE = "C"
	model.Method = method
	model.Url = url
	model.Body = body

	return model
}
