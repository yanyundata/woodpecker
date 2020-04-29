package apiap

func call(model Model) Result {
	switch model.Method {
	case "A":
	case "B":
		return aB(model.Method, model.Url, model.UrlParam)
	case "C":
		return c(model.Method, model.Url, model.Body)
	}

	return Result{}
}
