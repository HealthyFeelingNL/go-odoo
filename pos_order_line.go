package odoo

import (
	"fmt"
)

// PosOrderLine represents pos.order.line model.
type PosOrderLine struct {
	AttributeValueIds         *Relation `xmlrpc:"attribute_value_ids,omptempty"`
	ComboId                   *Many2One `xmlrpc:"combo_id,omptempty"`
	ComboLineIds              *Relation `xmlrpc:"combo_line_ids,omptempty"`
	ComboParentId             *Many2One `xmlrpc:"combo_parent_id,omptempty"`
	CompanyId                 *Many2One `xmlrpc:"company_id,omptempty"`
	CouponId                  *Many2One `xmlrpc:"coupon_id,omptempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId                *Many2One `xmlrpc:"currency_id,omptempty"`
	CustomAttributeValueIds   *Relation `xmlrpc:"custom_attribute_value_ids,omptempty"`
	CustomerNote              *String   `xmlrpc:"customer_note,omptempty"`
	Discount                  *Float    `xmlrpc:"discount,omptempty"`
	DisplayName               *String   `xmlrpc:"display_name,omptempty"`
	DownPaymentDetails        *String   `xmlrpc:"down_payment_details,omptempty"`
	FullProductName           *String   `xmlrpc:"full_product_name,omptempty"`
	Id                        *Int      `xmlrpc:"id,omptempty"`
	IsRewardLine              *Bool     `xmlrpc:"is_reward_line,omptempty"`
	IsTotalCostComputed       *Bool     `xmlrpc:"is_total_cost_computed,omptempty"`
	Margin                    *Float    `xmlrpc:"margin,omptempty"`
	MarginPercent             *Float    `xmlrpc:"margin_percent,omptempty"`
	Name                      *String   `xmlrpc:"name,omptempty"`
	Note                      *String   `xmlrpc:"note,omptempty"`
	Notice                    *String   `xmlrpc:"notice,omptempty"`
	OrderId                   *Many2One `xmlrpc:"order_id,omptempty"`
	PackLotIds                *Relation `xmlrpc:"pack_lot_ids,omptempty"`
	PointsCost                *Float    `xmlrpc:"points_cost,omptempty"`
	PriceExtra                *Float    `xmlrpc:"price_extra,omptempty"`
	PriceSubtotal             *Float    `xmlrpc:"price_subtotal,omptempty"`
	PriceSubtotalIncl         *Float    `xmlrpc:"price_subtotal_incl,omptempty"`
	PriceUnit                 *Float    `xmlrpc:"price_unit,omptempty"`
	ProductId                 *Many2One `xmlrpc:"product_id,omptempty"`
	ProductUomId              *Many2One `xmlrpc:"product_uom_id,omptempty"`
	Qty                       *Float    `xmlrpc:"qty,omptempty"`
	RefundOrderlineIds        *Relation `xmlrpc:"refund_orderline_ids,omptempty"`
	RefundedOrderlineId       *Many2One `xmlrpc:"refunded_orderline_id,omptempty"`
	RefundedQty               *Float    `xmlrpc:"refunded_qty,omptempty"`
	RewardId                  *Many2One `xmlrpc:"reward_id,omptempty"`
	RewardIdentifierCode      *String   `xmlrpc:"reward_identifier_code,omptempty"`
	SaleOrderLineId           *Many2One `xmlrpc:"sale_order_line_id,omptempty"`
	SaleOrderOriginId         *Many2One `xmlrpc:"sale_order_origin_id,omptempty"`
	SkipChange                *Bool     `xmlrpc:"skip_change,omptempty"`
	TaxIds                    *Relation `xmlrpc:"tax_ids,omptempty"`
	TaxIdsAfterFiscalPosition *Relation `xmlrpc:"tax_ids_after_fiscal_position,omptempty"`
	TotalCost                 *Float    `xmlrpc:"total_cost,omptempty"`
	Uuid                      *String   `xmlrpc:"uuid,omptempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosOrderLines represents array of pos.order.line model.
type PosOrderLines []PosOrderLine

// PosOrderLineModel is the odoo model name.
const PosOrderLineModel = "pos.order.line"

// Many2One convert PosOrderLine to *Many2One.
func (pol *PosOrderLine) Many2One() *Many2One {
	return NewMany2One(pol.Id.Get(), "")
}

// CreatePosOrderLine creates a new pos.order.line model and returns its id.
func (c *Client) CreatePosOrderLine(pol *PosOrderLine) (int64, error) {
	ids, err := c.CreatePosOrderLines([]*PosOrderLine{pol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosOrderLine creates a new pos.order.line model and returns its id.
func (c *Client) CreatePosOrderLines(pols []*PosOrderLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range pols {
		vv = append(vv, v)
	}
	return c.Create(PosOrderLineModel, vv)
}

// UpdatePosOrderLine updates an existing pos.order.line record.
func (c *Client) UpdatePosOrderLine(pol *PosOrderLine) error {
	return c.UpdatePosOrderLines([]int64{pol.Id.Get()}, pol)
}

// UpdatePosOrderLines updates existing pos.order.line records.
// All records (represented by ids) will be updated by pol values.
func (c *Client) UpdatePosOrderLines(ids []int64, pol *PosOrderLine) error {
	return c.Update(PosOrderLineModel, ids, pol)
}

// DeletePosOrderLine deletes an existing pos.order.line record.
func (c *Client) DeletePosOrderLine(id int64) error {
	return c.DeletePosOrderLines([]int64{id})
}

// DeletePosOrderLines deletes existing pos.order.line records.
func (c *Client) DeletePosOrderLines(ids []int64) error {
	return c.Delete(PosOrderLineModel, ids)
}

// GetPosOrderLine gets pos.order.line existing record.
func (c *Client) GetPosOrderLine(id int64) (*PosOrderLine, error) {
	pols, err := c.GetPosOrderLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if pols != nil && len(*pols) > 0 {
		return &((*pols)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.order.line not found", id)
}

// GetPosOrderLines gets pos.order.line existing records.
func (c *Client) GetPosOrderLines(ids []int64) (*PosOrderLines, error) {
	pols := &PosOrderLines{}
	if err := c.Read(PosOrderLineModel, ids, nil, pols); err != nil {
		return nil, err
	}
	return pols, nil
}

// FindPosOrderLine finds pos.order.line record by querying it with criteria.
func (c *Client) FindPosOrderLine(criteria *Criteria) (*PosOrderLine, error) {
	pols := &PosOrderLines{}
	if err := c.SearchRead(PosOrderLineModel, criteria, NewOptions().Limit(1), pols); err != nil {
		return nil, err
	}
	if pols != nil && len(*pols) > 0 {
		return &((*pols)[0]), nil
	}
	return nil, fmt.Errorf("pos.order.line was not found with criteria %v", criteria)
}

// FindPosOrderLines finds pos.order.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosOrderLines(criteria *Criteria, options *Options) (*PosOrderLines, error) {
	pols := &PosOrderLines{}
	if err := c.SearchRead(PosOrderLineModel, criteria, options, pols); err != nil {
		return nil, err
	}
	return pols, nil
}

// FindPosOrderLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosOrderLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosOrderLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosOrderLineId finds record id by querying it with criteria.
func (c *Client) FindPosOrderLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosOrderLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.order.line was not found with criteria %v and options %v", criteria, options)
}
