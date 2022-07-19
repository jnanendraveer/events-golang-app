package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/gin-gonic/gin"
)

func (server *Server) BillDeskController(c *gin.Context) {

	var (
		err  error
		data map[string]interface{}
	)
	if data, err = BillDeskJwtTokenGenerate(); err != nil {
		c.JSON(http.StatusPreconditionFailed, data)
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

func GetOutboundIP() string {
	url := "https://api.ipify.org?format=text" // we are using a pulib IP API, we're using ipify here, below are some others
	fmt.Printf("Getting IP address from  ipify ...\n")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("My IP is:%s\n", string(ip))
	return string(ip)
}
func BillDeskJwtTokenGenerate() (map[string]interface{}, error) {
	var (
		err error
		// token string
	)
	var obj map[string]interface{}
	str := []byte(`{"mercid":"FITPAS2UAT","orderid":"DSCGD43212F","amount":"300.00","order_date":"2022-07-19T16:48:00+0530","currency":"356","ru":"https://www.example.com/merchant/api/pgresponse","additional_info":{"additional_info1":"Details1","additional_info2":"Details2"},"itemcode":"DIRECT","device":{"init_channel":"internet","user_agent":"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:51.0) Gecko/20100101 Firefox/51.0","accept_header":"text/html"}}`)
	json.Unmarshal(str, &obj)
	obj["ip"] = GetOutboundIP()

	bytes, err := json.Marshal(obj)

	token, err := jose.Sign(string(bytes), jose.HS256, []byte("KEHpqq5UWQFwHnL6OBvMr7mln6OWWP3k"),
		jose.Header("clientid", "fitpas2uat"),
		jose.Header("alg", "HS256"), jose.Header("kid", "HMAC"))
	obj["token"] = token
	return obj, err
}
