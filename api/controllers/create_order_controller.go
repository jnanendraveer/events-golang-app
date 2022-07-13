package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/gin-gonic/gin"
	"github.com/jnanendraveer/events-golang-app/api/utils/Constants"
)

func (server *Server) CreateOrderController(c *gin.Context) {
	var (
		err  error
		data string
	)

	if data, err = WebEngageEvents(); err != nil {
		c.JSON(http.StatusPreconditionFailed, data)
		return
	}
	c.JSON(http.StatusOK, data)
	return
}
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
func WebEngageEvents() (string, error) {
	var (
		err   error
		token string
	)
	fmt.Println(GetOutboundIP())
	token, err = jose.Sign(fmt.Sprintf(`{
		"mercid":"FITPAS2UAT",
		"orderid":"TSSGF43214F",
		"amount":"300.00",
		"order_date":"2020-08-17T15:19:00+0530",
		"currency":"356",
		"ru":"https://www.example.com/merchant/api/pgresponse",
		"additional_info":{
		"additional_info1":"Details1",
		"additional_info2":"Details2"
		},
		"itemcode":"DIRECT",
		"device":{
		"init_channel":"internet",
		"ip": "%s",
		"user_agent":"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:51.0) Gecko/20100101 Firefox/51.0",
		"accept_header":"text/html"
		}
		}`, GetOutboundIP()), jose.HS256, []byte("KEHpqq5UWQFwHnL6OBvMr7mln6OWWP3k"),
		jose.Header("clientid", "fitpas2uat"),
		jose.Header("alg", "HS256"), jose.Header("kid", "HMAC"))

	url := Constants.BILLDESK_CREATE_ORDER_URL
	method := "POST"
	// webEngageData, _ := json.Marshal(RequestData)
	fmt.Printf(token)
	payload := strings.NewReader(token)

	client := &http.Client{}
	var req *http.Request
	req, err = http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/jose")
	req.Header.Add("Bd-Timestamp", "20200712102207")
	req.Header.Add("Accept", "application/jose")
	req.Header.Add("Bd-Traceid", "TSSGF43214F")
	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	token, _, err = jose.Decode(string(body), "KEHpqq5UWQFwHnL6OBvMr7mln6OWWP3k")

	if err == nil {
		//go use token
		fmt.Printf("\ndecoded payload = %v\n%v", token, err)
	}
	return string(body), err
}
