package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jnanendraveer/events-golang-app/api/utils/CommonFunction"
	"github.com/jnanendraveer/events-golang-app/api/utils/Constants"
	"gorm.io/gorm"
)

type CustomerPendingOrders struct {
	// user_membership_id, order_details, create_time, update_time, created_by, product_type, transaction_id, email_address, customer_name, mobile_number, other_details, team_zone_name, is_follow_up, follow_status, follow_up_datetime, follow_up_remarks, agent_user_id, fitcard_agents_id, is_order_confirm, is_renew, corporate_id, redeem_points, membership_price, fitprime_price, after_discount_price, service_tax, coupon_id, coupon_code, membership_plan_id, total_order_price, coupon_discount, referral_channel, start_date, end_date, device_name, app_version, device_type, is_subscription, payment_refund_remark, transaction_error_response, payment_status
	CustomerPendingId uint64    `gorm:"primary_key;auto_increment" json:"customer_pending_id" `
	OrderDetails      string    `json:"order_details"`
	CreateTime        time.Time `sql:"default:CURRENT_TIMESTAMP" json:"create_time"`
	CreatedBy         string    `gorm:"size:100" json:"created_by"`
}

func (RD *CustomerPendingOrders) TableName() string {
	return "fitpass_customer_pending_orders"
}

func (RD *CustomerPendingOrders) FillRemainningCustomerPendingOrder(requestBody map[string]interface{}) {
	orderDetails, _ := json.Marshal(requestBody)
	RD.CreateTime = CommonFunction.CurrentTime()
	RD.CreatedBy = Constants.SELF_CUSTOMER
	RD.OrderDetails = string(orderDetails)
}

func (RD *CustomerPendingOrders) SaveCustomerPendingOrders(db *gorm.DB) (*CustomerPendingOrders, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return &CustomerPendingOrders{}, err
	}
	if err := tx.Debug().Create(&RD).Error; err != nil {
		tx.Rollback()
		return &CustomerPendingOrders{}, err
	}
	// tx.SavePoint("fitpass_payments_link")
	return RD, tx.Commit().Error
}

func WebEngageEvents(RequestData interface{}) []byte {
	url := Constants.WEBENGAGE_EVENT_URL
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
