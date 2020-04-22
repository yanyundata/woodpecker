package apiUtils

type Param struct {
	data map[string]string
}

func (p Param) applyA(url string) string {
	if len(p.data) > 0 {
		url = url + "?"

		isFirst := true
		for k, v := range p.data {
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
