package utils

type PathParam map[string]string

func (pp PathParam) Set(name, value string) PathParam {
	pp[name] = value
	return pp
}

func (pp PathParam) applyA(url string) string {
	if len(pp) > 0 {
		url = url + "?"

		isFirst := true
		for k, v := range pp {
			if isFirst {
				url = url + k + "=" + v
				isFirst = false
			} else {
				url = url + "&" + k + "=" + v
			}
		}
	}

	return url
}

func (pp PathParam) applyB(url string) string {
	if len(pp) > 0 {
		for _, v := range pp {
			url = url + "/" + v
		}
	}

	return url
}
