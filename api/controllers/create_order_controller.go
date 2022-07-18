package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/gin-gonic/gin"
)

type PendingOrder struct {
	PendingOrderId uint64          `gorm:"primary_key;auto_increment" json:"pending_order_id"`
	UserId         int64           `json:"user_id"`
	UserName       string          `json:"user_name"`
	MobileNumber   int64           `json:"mobile_number"`
	IsSubscription bool            `json:"is_subscription"`
	DeviceDetails  json.RawMessage `json:"device_details"`
	CreateTime     time.Time       `sql:"default:CURRENT_TIMESTAMP" json:"create_time"`
	CreatedBy      string          `gorm:"size:100" json:"created_by"`
}

func (RD *PendingOrder) TableName() string {
	return "ftp_pending_orders"
}
func (server *Server) CreateOrderController(c *gin.Context) {

	var (
		err  error
		data map[string]interface{}
	)

	// c.BindJSON(&obj)
	if data, err = WebEngageEvents(); err != nil {
		c.JSON(http.StatusPreconditionFailed, data)
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

// func (RD *PendingOrder) SavePendingOrders(db *gorm.DB) (*PendingOrder, error) {
// 	tx := db.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	if err := tx.Error; err != nil {
// 		fmt.Println(err)
// 		return &PendingOrder{}, err
// 	}
// 	var data []PendingOrder
// 	for i := 0; i < 5000; i++ {
// 		fmt.Println(i)
// 		data = append(data, *RD)
// 	}
// 	if err := tx.Debug().Create(&data).Error; err != nil {
// 		tx.Rollback()
// 		fmt.Println(err)
// 		return &PendingOrder{}, err
// 	}
// 	// tx.SavePoint("fitpass_payments_link")
// 	tx.Commit()
// 	return RD, tx.Commit().Error
// }

func GetOutboundIP() string {
	url := "https://api.ipify.org?format=text" // we are using a pulib IP API, we're using ipify here, below are some others
	// https://www.ipify.org
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api
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
func WebEngageEvents() (map[string]interface{}, error) {
	var (
		err error
		// token string
	)
	var obj map[string]interface{}
	str := []byte(`{"mercid":"FITPAS2UAT","orderid":"TSSGF43214F","amount":"300.00","order_date":"2022-08-17T15:19:00+0530","currency":"356","ru":"https://www.example.com/merchant/api/pgresponse","additional_info":{"additional_info1":"Details1","additional_info2":"Details2"},"itemcode":"DIRECT","device":{"init_channel":"internet","user_agent":"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:51.0) Gecko/20100101 Firefox/51.0","accept_header":"text/html"}}`)
	json.Unmarshal(str, &obj)
	obj["ip"] = GetOutboundIP()

	bytes, err := json.Marshal(obj)

	token, err := jose.Sign(string(bytes), jose.HS256, []byte("KEHpqq5UWQFwHnL6OBvMr7mln6OWWP3k"),
		jose.Header("clientid", "fitpas2uat"),
		jose.Header("alg", "HS256"), jose.Header("kid", "HMAC"))

	fmt.Printf(token, obj["ip"])
	obj["token"] = token
	// url := Constants.BILLDESK_CREATE_ORDER_URL
	// method := "POST"

	// payload := strings.NewReader(token)

	// client := &http.Client{}
	// var req *http.Request
	// req, err = http.NewRequest(method, url, payload)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", err
	// }
	// req.Header.Add("Content-Type", "application/jose")
	// req.Header.Add("Bd-Timestamp", "20200712102207")
	// req.Header.Add("Accept", "application/jose")
	// req.Header.Add("Bd-Traceid", "TSSGF43214F")
	// var res *http.Response
	// res, err = client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", err
	// }
	// defer res.Body.Close()
	// var body []byte
	// body, err = ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", err
	// }
	// token, _, err = jose.Decode(string(body), "KEHpqq5UWQFwHnL6OBvMr7mln6OWWP3k")

	// if err == nil {
	// 	//go use token
	// 	fmt.Printf("\ndecoded payload = %v\n%v", token, err)
	// }

	return obj, err
}
