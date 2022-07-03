package test

import (
	"fmt"
	"gorm.io/gorm"
	"mall/global/dao"
	"mall/global/dao/domain"
	"mall/global/dao/model"
	"testing"
)

func TestFF(t *testing.T) {
	var count = 0
	var orderSetting model.OmsOrderSetting
	dao.DB.First(&orderSetting)
	timeOutOrders := GetTimeOutOrders(orderSetting.NormalOrderOvertime)
	fmt.Println(len(timeOutOrders))
	fmt.Println(count)
}
func GetTimeOutOrders(time int) (timeOutOrders []domain.OmsOrderDetail) {
	order := &model.OmsOrder{}
	dao.DB.Table(order.TableName() + " o").Select("o.id," +
		"o.order_sn," +
		"o.coupon_id," +
		"o.integration," +
		"o.member_id," +
		"o.use_integration," +
		"ot.id ot_id," +
		"ot.product_name ot_product_name," +
		"ot.product_sku_id ot_product_sku_id," +
		"ot.product_sku_code ot_product_sku_code," +
		"ot.product_quantity ot_product_quantity").
		Joins("left join oms_order_item ot on o.id=ot.order_id").
		//Where("o.status =0 and "+
		//	"date_add(NOW(),INTERVAL INTERVAL -? MINUTE ", time).
		Scan(&timeOutOrders)
	sql := dao.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Table(order.TableName() + " o").Select("o.id," +
			"o.order_sn," +
			"o.coupon_id," +
			"o.integration," +
			"o.member_id," +
			"o.use_integration," +
			"ot.id ot_id," +
			"ot.product_name ot_product_name," +
			"ot.product_sku_id ot_product_sku_id," +
			"ot.product_sku_code ot_product_sku_code," +
			"ot.product_quantity ot_product_quantity").
			Joins("left join oms_order_item ot on o.id=ot.order_id").Scan(&timeOutOrders)
	})
	fmt.Println(sql)
	return
}
