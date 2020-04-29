package api

type API struct {
	method   string
	url      string
	urlParam string
	body     interface{}
}

func (api API) NewTypeAB(method string, url string, urlParam string) {
	api.method = method
	api.url = url
	api.urlParam = urlParam
}

func (api API) NewTypeC(method string, url string, body interface{}) {
	api.method = method
	api.url = url
	api.body = body
}
