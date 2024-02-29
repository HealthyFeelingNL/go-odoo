package odoo

import (
	"fmt"
)

// ProductPricelist represents product.pricelist model.
type ProductPricelist struct {
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	Code                        *String    `xmlrpc:"code,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryGroupIds             *Relation  `xmlrpc:"country_group_ids,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	DiscountPolicy              *Selection `xmlrpc:"discount_policy,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	ItemIds                     *Relation  `xmlrpc:"item_ids,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omptempty"`
	Selectable                  *Bool      `xmlrpc:"selectable,omptempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omptempty"`
	WebsiteId                   *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProductPricelists represents array of product.pricelist model.
type ProductPricelists []ProductPricelist

// ProductPricelistModel is the odoo model name.
const ProductPricelistModel = "product.pricelist"

// Many2One convert ProductPricelist to *Many2One.
func (pp *ProductPricelist) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductPricelist creates a new product.pricelist model and returns its id.
func (c *Client) CreateProductPricelist(pp *ProductPricelist) (int64, error) {
	ids, err := c.CreateProductPricelists([]*ProductPricelist{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductPricelist creates a new product.pricelist model and returns its id.
func (c *Client) CreateProductPricelists(pps []*ProductPricelist) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProductPricelistModel, vv)
}

// UpdateProductPricelist updates an existing product.pricelist record.
func (c *Client) UpdateProductPricelist(pp *ProductPricelist) error {
	return c.UpdateProductPricelists([]int64{pp.Id.Get()}, pp)
}

// UpdateProductPricelists updates existing product.pricelist records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductPricelists(ids []int64, pp *ProductPricelist) error {
	return c.Update(ProductPricelistModel, ids, pp)
}

// DeleteProductPricelist deletes an existing product.pricelist record.
func (c *Client) DeleteProductPricelist(id int64) error {
	return c.DeleteProductPricelists([]int64{id})
}

// DeleteProductPricelists deletes existing product.pricelist records.
func (c *Client) DeleteProductPricelists(ids []int64) error {
	return c.Delete(ProductPricelistModel, ids)
}

// GetProductPricelist gets product.pricelist existing record.
func (c *Client) GetProductPricelist(id int64) (*ProductPricelist, error) {
	pps, err := c.GetProductPricelists([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.pricelist not found", id)
}

// GetProductPricelists gets product.pricelist existing records.
func (c *Client) GetProductPricelists(ids []int64) (*ProductPricelists, error) {
	pps := &ProductPricelists{}
	if err := c.Read(ProductPricelistModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPricelist finds product.pricelist record by querying it with criteria.
func (c *Client) FindProductPricelist(criteria *Criteria) (*ProductPricelist, error) {
	pps := &ProductPricelists{}
	if err := c.SearchRead(ProductPricelistModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("product.pricelist was not found with criteria %v", criteria)
}

// FindProductPricelists finds product.pricelist records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricelists(criteria *Criteria, options *Options) (*ProductPricelists, error) {
	pps := &ProductPricelists{}
	if err := c.SearchRead(ProductPricelistModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPricelistIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricelistIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductPricelistModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductPricelistId finds record id by querying it with criteria.
func (c *Client) FindProductPricelistId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPricelistModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.pricelist was not found with criteria %v and options %v", criteria, options)
}
