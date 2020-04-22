package apiUtils

type Param struct {
	param map[string]string
}

func (p Param) applyA(url string) string {
	if len(p.param) > 0 {
		url = url + "?"

		isFirst := true
		for k, v := range p.param {
			if isFirst {
				url = url + k + "=" + v
				isFirst = false
			} else {
				url = "&" + url + k + "=" + v
			}
		}
	}

	return url
}
