package odoo

import (
	"fmt"
)

// LoyaltyCard represents loyalty.card model.
type LoyaltyCard struct {
	Code                     *String    `xmlrpc:"code,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	ExpirationDate           *Time      `xmlrpc:"expiration_date,omptempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	OrderId                  *Many2One  `xmlrpc:"order_id,omptempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omptempty"`
	PointName                *String    `xmlrpc:"point_name,omptempty"`
	Points                   *Float     `xmlrpc:"points,omptempty"`
	PointsDisplay            *String    `xmlrpc:"points_display,omptempty"`
	ProgramId                *Many2One  `xmlrpc:"program_id,omptempty"`
	ProgramType              *Selection `xmlrpc:"program_type,omptempty"`
	RatingIds                *Relation  `xmlrpc:"rating_ids,omptempty"`
	UseCount                 *Int       `xmlrpc:"use_count,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// LoyaltyCards represents array of loyalty.card model.
type LoyaltyCards []LoyaltyCard

// LoyaltyCardModel is the odoo model name.
const LoyaltyCardModel = "loyalty.card"

// Many2One convert LoyaltyCard to *Many2One.
func (lc *LoyaltyCard) Many2One() *Many2One {
	return NewMany2One(lc.Id.Get(), "")
}

// CreateLoyaltyCard creates a new loyalty.card model and returns its id.
func (c *Client) CreateLoyaltyCard(lc *LoyaltyCard) (int64, error) {
	ids, err := c.CreateLoyaltyCards([]*LoyaltyCard{lc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLoyaltyCard creates a new loyalty.card model and returns its id.
func (c *Client) CreateLoyaltyCards(lcs []*LoyaltyCard) ([]int64, error) {
	var vv []interface{}
	for _, v := range lcs {
		vv = append(vv, v)
	}
	return c.Create(LoyaltyCardModel, vv)
}

// UpdateLoyaltyCard updates an existing loyalty.card record.
func (c *Client) UpdateLoyaltyCard(lc *LoyaltyCard) error {
	return c.UpdateLoyaltyCards([]int64{lc.Id.Get()}, lc)
}

// UpdateLoyaltyCards updates existing loyalty.card records.
// All records (represented by ids) will be updated by lc values.
func (c *Client) UpdateLoyaltyCards(ids []int64, lc *LoyaltyCard) error {
	return c.Update(LoyaltyCardModel, ids, lc)
}

// DeleteLoyaltyCard deletes an existing loyalty.card record.
func (c *Client) DeleteLoyaltyCard(id int64) error {
	return c.DeleteLoyaltyCards([]int64{id})
}

// DeleteLoyaltyCards deletes existing loyalty.card records.
func (c *Client) DeleteLoyaltyCards(ids []int64) error {
	return c.Delete(LoyaltyCardModel, ids)
}

// GetLoyaltyCard gets loyalty.card existing record.
func (c *Client) GetLoyaltyCard(id int64) (*LoyaltyCard, error) {
	lcs, err := c.GetLoyaltyCards([]int64{id})
	if err != nil {
		return nil, err
	}
	if lcs != nil && len(*lcs) > 0 {
		return &((*lcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of loyalty.card not found", id)
}

// GetLoyaltyCards gets loyalty.card existing records.
func (c *Client) GetLoyaltyCards(ids []int64) (*LoyaltyCards, error) {
	lcs := &LoyaltyCards{}
	if err := c.Read(LoyaltyCardModel, ids, nil, lcs); err != nil {
		return nil, err
	}
	return lcs, nil
}

// FindLoyaltyCard finds loyalty.card record by querying it with criteria.
func (c *Client) FindLoyaltyCard(criteria *Criteria) (*LoyaltyCard, error) {
	lcs := &LoyaltyCards{}
	if err := c.SearchRead(LoyaltyCardModel, criteria, NewOptions().Limit(1), lcs); err != nil {
		return nil, err
	}
	if lcs != nil && len(*lcs) > 0 {
		return &((*lcs)[0]), nil
	}
	return nil, fmt.Errorf("loyalty.card was not found with criteria %v", criteria)
}

// FindLoyaltyCards finds loyalty.card records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyCards(criteria *Criteria, options *Options) (*LoyaltyCards, error) {
	lcs := &LoyaltyCards{}
	if err := c.SearchRead(LoyaltyCardModel, criteria, options, lcs); err != nil {
		return nil, err
	}
	return lcs, nil
}

// FindLoyaltyCardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyCardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LoyaltyCardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLoyaltyCardId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyCardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyCardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("loyalty.card was not found with criteria %v and options %v", criteria, options)
}
