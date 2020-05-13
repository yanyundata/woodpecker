package apiap

import (
	"github.com/go-resty/resty/v2"
)

func GetA(url string, param string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetQueryString(param).
		Get(url)

	return Result{resp, err}
}

func GetB(url string, param string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(url + param)

	return Result{resp, err}
}

func PostA(url string, param string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetQueryString(param).
		Post(url)

	return Result{resp, err}
}

func PostC(url string, object interface{}) Result {
	client := resty.New()
	resp, err := client.R().
		SetBody(object).
		Post(url)

	return Result{resp, err}
}

func PutA(url string, param string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetQueryString(param).
		Put(url)

	return Result{resp, err}
}

func PutB(url string, param string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Put(url + param)

	return Result{resp, err}
}

func PutC(url string, object interface{}) Result {
	client := resty.New()
	resp, err := client.R().
		SetBody(object).
		Put(url)

	return Result{resp, err}
}

func DeleteA(url string, param string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetQueryString(param).
		Delete(url)

	return Result{resp, err}
}

func DeleteB(url string, param string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Delete(url + param)

	return Result{resp, err}
}
