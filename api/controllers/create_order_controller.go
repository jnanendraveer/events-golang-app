package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/gin-gonic/gin"
	"github.com/jnanendraveer/events-golang-app/api/utils/CommonFunction"
	"github.com/jnanendraveer/events-golang-app/api/utils/Constants"
	"gorm.io/gorm"
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
		data string
		// obj    json.RawMessage
		obj    map[string]interface{}
		orders PendingOrder
	)
	c.BindJSON(&obj)
	// orders.DeviceDetails = obj
	orders.CreateTime = CommonFunction.CurrentTime()
	orders.CreatedBy = Constants.SELF_CUSTOMER
	orders.IsSubscription = true
	// orders.SavePendingOrders(server.DB)
	if data, err = WebEngageEvents(obj); err != nil {
		c.JSON(http.StatusPreconditionFailed, data)
		return
	}
	c.JSON(http.StatusOK, data)
	return
}

func (RD *PendingOrder) SavePendingOrders(db *gorm.DB) (*PendingOrder, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return &PendingOrder{}, err
	}
	data := []PendingOrder{}
	for i := 0; i < 5000; i++ {
		data = append(data, *RD)
	}

	if err := tx.Debug().Create(&data).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		return &PendingOrder{}, err
	}
	// tx.SavePoint("fitpass_payments_link")
	return RD, tx.Commit().Error
}

func GetOutboundIP() string {
	url := "https://api.ipify.org?format=text"
	fmt.Printf("Getting IP address from  ipify\n")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("ip is ........", string(ip))
	return string(ip)
}
func WebEngageEvents(obj map[string]interface{}) (string, error) {
	var (
		err   error
		token string
	)
	fmt.Println(GetOutboundIP())
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
