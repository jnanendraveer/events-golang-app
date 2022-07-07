package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jnanendraveer/transactions-golang-app/api/utils/CommonFunction/Constants"
)

type CustomerPendingOrders struct {
}

func (RD *CustomerPendingOrders) TableName() string {
	return "fitpass_customer_pending_orders"
}

// func (){

// }
func WebEngageEvents(RequestData interface{}) []byte {
	url := "https://api.in.webengage.com/v1/accounts/in~11b56432b/events"
	method := "POST"
	webEngageData, _ := json.Marshal(RequestData)
	payload := strings.NewReader(string(webEngageData))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	req.Header.Add("Content-Type", Constants.CONTENT_TYPE_JSON)
	req.Header.Add("Authorization", Constants.WEBENGAGE_API_KEY)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return body
}
