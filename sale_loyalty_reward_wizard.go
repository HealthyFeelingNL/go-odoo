package odoo

import (
	"fmt"
)

// SaleLoyaltyRewardWizard represents sale.loyalty.reward.wizard model.
type SaleLoyaltyRewardWizard struct {
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	MultiProductReward *Bool     `xmlrpc:"multi_product_reward,omptempty"`
	OrderId            *Many2One `xmlrpc:"order_id,omptempty"`
	RewardIds          *Relation `xmlrpc:"reward_ids,omptempty"`
	RewardProductIds   *Relation `xmlrpc:"reward_product_ids,omptempty"`
	SelectedProductId  *Many2One `xmlrpc:"selected_product_id,omptempty"`
	SelectedRewardId   *Many2One `xmlrpc:"selected_reward_id,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaleLoyaltyRewardWizards represents array of sale.loyalty.reward.wizard model.
type SaleLoyaltyRewardWizards []SaleLoyaltyRewardWizard

// SaleLoyaltyRewardWizardModel is the odoo model name.
const SaleLoyaltyRewardWizardModel = "sale.loyalty.reward.wizard"

// Many2One convert SaleLoyaltyRewardWizard to *Many2One.
func (slrw *SaleLoyaltyRewardWizard) Many2One() *Many2One {
	return NewMany2One(slrw.Id.Get(), "")
}

// CreateSaleLoyaltyRewardWizard creates a new sale.loyalty.reward.wizard model and returns its id.
func (c *Client) CreateSaleLoyaltyRewardWizard(slrw *SaleLoyaltyRewardWizard) (int64, error) {
	ids, err := c.CreateSaleLoyaltyRewardWizards([]*SaleLoyaltyRewardWizard{slrw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleLoyaltyRewardWizard creates a new sale.loyalty.reward.wizard model and returns its id.
func (c *Client) CreateSaleLoyaltyRewardWizards(slrws []*SaleLoyaltyRewardWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range slrws {
		vv = append(vv, v)
	}
	return c.Create(SaleLoyaltyRewardWizardModel, vv)
}

// UpdateSaleLoyaltyRewardWizard updates an existing sale.loyalty.reward.wizard record.
func (c *Client) UpdateSaleLoyaltyRewardWizard(slrw *SaleLoyaltyRewardWizard) error {
	return c.UpdateSaleLoyaltyRewardWizards([]int64{slrw.Id.Get()}, slrw)
}

// UpdateSaleLoyaltyRewardWizards updates existing sale.loyalty.reward.wizard records.
// All records (represented by ids) will be updated by slrw values.
func (c *Client) UpdateSaleLoyaltyRewardWizards(ids []int64, slrw *SaleLoyaltyRewardWizard) error {
	return c.Update(SaleLoyaltyRewardWizardModel, ids, slrw)
}

// DeleteSaleLoyaltyRewardWizard deletes an existing sale.loyalty.reward.wizard record.
func (c *Client) DeleteSaleLoyaltyRewardWizard(id int64) error {
	return c.DeleteSaleLoyaltyRewardWizards([]int64{id})
}

// DeleteSaleLoyaltyRewardWizards deletes existing sale.loyalty.reward.wizard records.
func (c *Client) DeleteSaleLoyaltyRewardWizards(ids []int64) error {
	return c.Delete(SaleLoyaltyRewardWizardModel, ids)
}

// GetSaleLoyaltyRewardWizard gets sale.loyalty.reward.wizard existing record.
func (c *Client) GetSaleLoyaltyRewardWizard(id int64) (*SaleLoyaltyRewardWizard, error) {
	slrws, err := c.GetSaleLoyaltyRewardWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if slrws != nil && len(*slrws) > 0 {
		return &((*slrws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.loyalty.reward.wizard not found", id)
}

// GetSaleLoyaltyRewardWizards gets sale.loyalty.reward.wizard existing records.
func (c *Client) GetSaleLoyaltyRewardWizards(ids []int64) (*SaleLoyaltyRewardWizards, error) {
	slrws := &SaleLoyaltyRewardWizards{}
	if err := c.Read(SaleLoyaltyRewardWizardModel, ids, nil, slrws); err != nil {
		return nil, err
	}
	return slrws, nil
}

// FindSaleLoyaltyRewardWizard finds sale.loyalty.reward.wizard record by querying it with criteria.
func (c *Client) FindSaleLoyaltyRewardWizard(criteria *Criteria) (*SaleLoyaltyRewardWizard, error) {
	slrws := &SaleLoyaltyRewardWizards{}
	if err := c.SearchRead(SaleLoyaltyRewardWizardModel, criteria, NewOptions().Limit(1), slrws); err != nil {
		return nil, err
	}
	if slrws != nil && len(*slrws) > 0 {
		return &((*slrws)[0]), nil
	}
	return nil, fmt.Errorf("sale.loyalty.reward.wizard was not found with criteria %v", criteria)
}

// FindSaleLoyaltyRewardWizards finds sale.loyalty.reward.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleLoyaltyRewardWizards(criteria *Criteria, options *Options) (*SaleLoyaltyRewardWizards, error) {
	slrws := &SaleLoyaltyRewardWizards{}
	if err := c.SearchRead(SaleLoyaltyRewardWizardModel, criteria, options, slrws); err != nil {
		return nil, err
	}
	return slrws, nil
}

// FindSaleLoyaltyRewardWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleLoyaltyRewardWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleLoyaltyRewardWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleLoyaltyRewardWizardId finds record id by querying it with criteria.
func (c *Client) FindSaleLoyaltyRewardWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleLoyaltyRewardWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.loyalty.reward.wizard was not found with criteria %v and options %v", criteria, options)
}
