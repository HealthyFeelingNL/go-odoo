package odoo

import (
	"fmt"
)

// MailActivityPlan represents mail.activity.plan model.
type MailActivityPlan struct { 
	Active *Bool `xmlrpc:"active,omptempty"`
	AssignationSummary *String `xmlrpc:"assignation_summary,omptempty"`
	CompanyId *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate *Time `xmlrpc:"create_date,omptempty"`
	CreateUid *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	HasUserOnDemand *Bool `xmlrpc:"has_user_on_demand,omptempty"`
	Id *Int `xmlrpc:"id,omptempty"`
	Name *String `xmlrpc:"name,omptempty"`
	ResModel *Selection `xmlrpc:"res_model,omptempty"`
	ResModelId *Many2One `xmlrpc:"res_model_id,omptempty"`
	StepsCount *Int `xmlrpc:"steps_count,omptempty"`
	TemplateIds *Relation `xmlrpc:"template_ids,omptempty"`
	WriteDate *Time `xmlrpc:"write_date,omptempty"`
	WriteUid *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailActivityPlans represents array of mail.activity.plan model.
type MailActivityPlans []MailActivityPlan

// MailActivityPlanModel is the odoo model name.
const MailActivityPlanModel = "mail.activity.plan"

// Many2One convert MailActivityPlan to *Many2One.
func (map *MailActivityPlan) Many2One() *Many2One {
	return NewMany2One(map.Id.Get(), "")
}

// CreateMailActivityPlan creates a new mail.activity.plan model and returns its id.
func (c *Client) CreateMailActivityPlan(map *MailActivityPlan) (int64, error) {
	ids, err := c.CreateMailActivityPlans([]*MailActivityPlan{ map })
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailActivityPlan creates a new mail.activity.plan model and returns its id.
func (c *Client) CreateMailActivityPlans(maps []*MailActivityPlan) ([]int64, error) {
	var vv []interface{}
	for _, v := range maps {
		vv = append(vv, v)
	}
	return c.Create(MailActivityPlanModel, vv)
}

// UpdateMailActivityPlan updates an existing mail.activity.plan record.
func (c *Client) UpdateMailActivityPlan(map *MailActivityPlan) error {
	return c.UpdateMailActivityPlans([]int64{ map.Id.Get() }, map)
}

// UpdateMailActivityPlans updates existing mail.activity.plan records.
// All records (represented by ids) will be updated by map values.
func (c *Client) UpdateMailActivityPlans(ids []int64, map *MailActivityPlan) error {
	return c.Update(MailActivityPlanModel, ids, map)
}

// DeleteMailActivityPlan deletes an existing mail.activity.plan record.
func (c *Client) DeleteMailActivityPlan(id int64) error {
	return c.DeleteMailActivityPlans([]int64{id})
}

// DeleteMailActivityPlans deletes existing mail.activity.plan records.
func (c *Client) DeleteMailActivityPlans(ids []int64) error {
	return c.Delete(MailActivityPlanModel, ids)
}

// GetMailActivityPlan gets mail.activity.plan existing record.
func (c *Client) GetMailActivityPlan(id int64) (*MailActivityPlan, error) {
	maps, err := c.GetMailActivityPlans([]int64{id})
	if err != nil {
		return nil, err
	}
	if maps != nil && len(*maps) > 0 {
		return &((*maps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.activity.plan not found", id)
}

// GetMailActivityPlans gets mail.activity.plan existing records.
func (c *Client) GetMailActivityPlans(ids []int64) (*MailActivityPlans, error) {
	maps := &MailActivityPlans{}
	if err := c.Read(MailActivityPlanModel, ids, nil, maps); err != nil {
		return nil, err
	}
	return maps, nil
}

// FindMailActivityPlan finds mail.activity.plan record by querying it with criteria.
func (c *Client) FindMailActivityPlan(criteria *Criteria) (*MailActivityPlan, error) {
	maps := &MailActivityPlans{}
	if err := c.SearchRead(MailActivityPlanModel, criteria, NewOptions().Limit(1), maps); err != nil {
		return nil, err
	}
	if maps != nil && len(*maps) > 0 {
		return &((*maps)[0]), nil
	}
	return nil, fmt.Errorf("mail.activity.plan was not found with criteria %v", criteria)
}

// FindMailActivityPlans finds mail.activity.plan records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityPlans(criteria *Criteria, options *Options) (*MailActivityPlans, error) {
	maps := &MailActivityPlans{}
	if err := c.SearchRead(MailActivityPlanModel, criteria, options, maps); err != nil {
		return nil, err
	}
	return maps, nil
}

// FindMailActivityPlanIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityPlanIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailActivityPlanModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailActivityPlanId finds record id by querying it with criteria.
func (c *Client) FindMailActivityPlanId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailActivityPlanModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.activity.plan was not found with criteria %v and options %v", criteria, options)
}
