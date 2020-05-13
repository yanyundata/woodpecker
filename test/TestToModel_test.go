package test

import (
	"github.com/yanyundata/woodpecker/apiap"
	"log"
	"testing"
)

type ParamOrderSettings struct {
	AfterSalesPeriod        int `json:"afterSalesPeriod"`
	DeliveryDeadline        int `json:"deliveryDeadline"`
	DeliveryReminder        int `json:"deliveryReminder"`
	DeliveryReminderToClose int `json:"deliveryReminderToClose"`
	MiaoShaConfirmTimeLimit int `json:"miaoShaConfirmTimeLimit"`
	MiaoShaPayTimeLimit     int `json:"miaoShaPayTimeLimit"`
	OrderConfirmTimeLimit   int `json:"orderConfirmTimeLimit"`
	OrderPayTimeLimit       int `json:"orderPayTimeLimit"`
	PraisePeriod            int `json:"praisePeriod"`
}

type ResultOrderSettings struct {
	Code int `json:"code"`
	Data struct {
		AfterSalesPeriod        int `json:"afterSalesPeriod"`
		DeliveryDeadline        int `json:"deliveryDeadline"`
		DeliveryReminder        int `json:"deliveryReminder"`
		DeliveryReminderToClose int `json:"deliveryReminderToClose"`
		MiaoShaConfirmTimeLimit int `json:"miaoShaConfirmTimeLimit"`
		MiaoShaPayTimeLimit     int `json:"miaoShaPayTimeLimit"`
		OrderConfirmTimeLimit   int `json:"orderConfirmTimeLimit"`
		OrderPayTimeLimit       int `json:"orderPayTimeLimit"`
		PraisePeriod            int `json:"praisePeriod"`
	} `json:"data"`
	Msg  string `json:"msg"`
	Time int    `json:"time"`
}

func TestToModel(t *testing.T) {
	resultOrderSettings := ResultOrderSettings{}
	apiap.GetA("https://yan-yun.com:38085/lvyuan/admin/api/orderSettings", "").ToModel(&resultOrderSettings)

	log.Println(resultOrderSettings.Data.DeliveryReminder)
}
