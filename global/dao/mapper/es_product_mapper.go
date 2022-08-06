package mapper

import (
	"fmt"
	"gorm.io/gorm"
	"mall/global/dao/document"
	"mall/global/dao/domain"
	"mall/global/dao/model"
)

type esProductMapper struct {
	db *gorm.DB
}

func NewEsProductMapper(db *gorm.DB) EsProductMapper {
	return &esProductMapper{db: db}
}
func (m *esProductMapper) GetAllEsProductList(id int64) (esProducts []document.EsProduct,
	err error) {

	product := model.PmsProduct{}

	sel := "p.id id," +
		"p.product_sn productSn," +
		"p.brand_id brandId," +
		"p.brand_name brandName," +
		"p.product_category_id productCategoryId," +
		"p.product_category_name productCategoryName," +
		"p.pic pic," +
		"p.name name," +
		"p.sub_title subTitle," +
		"p.price price," +
		"p.sale sale," +
		"p.new_status newStatus," +
		" p.recommand_status recommandStatus," +
		" p.stock stock," +
		"p.promotion_type promotionType," +
		"p.keywords keywords," +
		" p.sort sort," +
		"pav.id attr_id," +
		"pav.value attr_value," +
		"pav.product_attribute_id attr_product_attribute_id," +
		"pa.type attr_type," +
		"pa.name attr_name"
	whe := " delete_status = 0 and publish_status = 1"
	if id != 0 {
		whe = whe + fmt.Sprintf(" and p.id = %d", id)
	}
	rows, err := m.db.Table(product.TableName() + " p").Select(sel).
		Joins("  left join pms_product_attribute_value pav on p.id = pav.product_id").
		Joins(" left join pms_product_attribute pa on pav.product_attribute_id= pa.id").
		Where(whe).
		Rows()
	if err != nil {
		return
	}

	if rows.Next() {
		e := document.EsProduct{}
		att := domain.EsProductAttributeValue{}
		err = rows.Scan(&e.Id, &e.ProductSn, &e.BrandId, &e.BrandName, &e.ProductCategoryId, &e.ProductCategoryName,
			&e.Pic, &e.Name, &e.SubTitle, &e.Price, &e.Sale, &e.NewStatus, &e.RecommandStatus, &e.Stock, &e.PromotionType,
			&e.Keywords, &e.Sort, &att.Id, &att.Value, &att.ProductAttributeID, &att.Type, &att.Name)
		e.AttrValueList = append(e.AttrValueList, att)
		esProducts = append(esProducts, e)
		if err != nil {
			return
		}
	} else {
		return
	}
	i := 0
	for rows.Next() {
		e := document.EsProduct{}
		att := domain.EsProductAttributeValue{}
		err = rows.Scan(&e.Id, &e.ProductSn, &e.BrandId, &e.BrandName, &e.ProductCategoryId, &e.ProductCategoryName,
			&e.Pic, &e.Name, &e.SubTitle, &e.Price, &e.Sale, &e.NewStatus, &e.RecommandStatus, &e.Stock, &e.PromotionType,
			&e.Keywords, &e.Sort, &att.Id, &att.Value, &att.ProductAttributeID, &att.Type, &att.Name)
		if err != nil {
			return
		}

		if esProducts[i].Id == e.Id {
			e = esProducts[i]
			e.AttrValueList = append(e.AttrValueList, att)
			esProducts[i] = e
		} else {
			e.AttrValueList = append(e.AttrValueList, att)
			esProducts = append(esProducts, e)
			i++
		}
	}
	return
}
