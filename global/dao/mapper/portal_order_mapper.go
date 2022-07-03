package mapper

import (
	"mall/global/dao"
	"mall/global/dao/domain"
	"mall/global/dao/model"
)

func GetTimeOutOrders(time int) (timeOutOrders []domain.OmsOrderDetail) {
	order := &model.OmsOrder{}
	dao.DB.Table(order.TableName()+" 0").Select("o.id,"+
		"o.order_sn,"+
		"o.coupon_id,"+
		"o.integration,"+
		"o.member_id,"+
		"o.use_integration,"+
		"or.id ot_id,"+
		"ot.product_name ot_product_name,"+
		"ot.product_sku_id ot_product_sku_id,"+
		"ot.product_sku_coded ot_product_sku_code"+
		"ot.product_quantity ot_product_quantity").
		Joins("left join oms_order_item ot on o.id=ot.order_id").
		Where("o.status =0 and "+
			"date_add(NOW(),INTERVAL INTERVAL -? MINUTE ", time).
		Scan(&timeOutOrders)
	return
}
