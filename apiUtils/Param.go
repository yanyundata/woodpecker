package apiUtils

type Param struct {
	Data map[string]string
}

func (p Param) applyA(url string) string {
	if len(p.Data) > 0 {
		url = url + "?"

		isFirst := true
		for k, v := range p.Data {
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
