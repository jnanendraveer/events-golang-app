package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/gin-gonic/gin"
	"github.com/jnanendraveer/events-golang-app/api/utils/Constants"
)

func (server *Server) CreateOrderController(c *gin.Context) {

	var (
		err  error
		data string
		obj  map[string]interface{}
	)
	c.BindJSON(&obj)
	if data, err = WebEngageEvents(obj); err != nil {
		c.JSON(http.StatusPreconditionFailed, data)
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

func GetOutboundIP() string {
	// url := "https://api.ipify.org?format=text"
	// fmt.Printf("Getting IP address from  ipify\n")
	// resp, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// ip, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("ip is ........", string(ip))
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
func WebEngageEvents(obj map[string]interface{}) (string, error) {
	var (
		err   error
		token string
	)
	obj["ip"] = GetOutboundIP()
	bytes, err := json.Marshal(obj)

	token, err = jose.Sign(string(bytes), jose.HS256, []byte("KEHpqq5UWQFwHnL6OBvMr7mln6OWWP3k"),
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
	req.Header.Add("Bd-Timestamp", "20220713102207")
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
