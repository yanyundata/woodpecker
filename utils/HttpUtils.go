package utils

import (
	"github.com/go-resty/resty/v2"
)

func GetUrl(url string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(url)

	return Result{resp, err}
}

func GetParam(url string, param string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetQueryString(param).
		Get(url)

	return Result{resp, err}
}

func PostParam(url string, param string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetQueryString(param).
		Post(url)

	return Result{resp, err}
}

func PostJson(url string, object interface{}) Result {
	client := resty.New()
	resp, err := client.R().
		SetBody(object).
		Post(url)

	return Result{resp, err}
}

func PutParam(url string, param string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetQueryString(param).
		Put(url)

	return Result{resp, err}
}

func PutUrl(url string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Put(url)

	return Result{resp, err}
}

func PutJson(url string, object interface{}) Result {
	client := resty.New()
	resp, err := client.R().
		SetBody(object).
		Put(url)

	return Result{resp, err}
}

func DeleteParam(url string, param string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetQueryString(param).
		Delete(url)

	return Result{resp, err}
}

func DeleteUrl(url string) Result {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Delete(url)

	return Result{resp, err}
}
