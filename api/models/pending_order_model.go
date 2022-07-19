package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

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

func (RD *PendingOrder) FillRemainningPendingOrder() {
	RD.CreateTime = CommonFunction.CurrentTime()
	RD.CreatedBy = Constants.SELF_CUSTOMER
}

func (RD *PendingOrder) Validate(action string) error {
	switch strings.ToLower(action) {
	default:
		if RD.UserName == "" {
			return errors.New("Required User Name")
		}
		if RD.MobileNumber == 0 {
			return errors.New("Required User Name")
		}
		// if RD.UserName == "" {
		// 	return errors.New("Required User Name")
		// }
		// if RD.UserName == "" {
		// 	return errors.New("Required User Name")
		// }
		// if RD.UserName == "" {
		// 	return errors.New("Required User Name")
		// }
		// if RD.UserName == "" {
		// 	return errors.New("Required User Name")
		// }
		// if RD.UserName == "" {
		// 	return errors.New("Required User Name")
		// }
		return nil
	}
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
	var data []PendingOrder
	for i := 0; i < 5000; i++ {
		fmt.Println(i)
		data = append(data, *RD)
	}
	if err := tx.Debug().Create(&data).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		return &PendingOrder{}, err
	}
	// tx.SavePoint("fitpass_payments_link")
	tx.Commit()
	return RD, tx.Commit().Error
}
