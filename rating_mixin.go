package odoo

import (
	"fmt"
)

// RatingMixin represents rating.mixin model.
type RatingMixin struct {
	HasMessage                   *Bool      `xmlrpc:"has_message,omptempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	RatingAvg                    *Float     `xmlrpc:"rating_avg,omptempty"`
	RatingAvgText                *Selection `xmlrpc:"rating_avg_text,omptempty"`
	RatingCount                  *Int       `xmlrpc:"rating_count,omptempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omptempty"`
	RatingLastFeedback           *String    `xmlrpc:"rating_last_feedback,omptempty"`
	RatingLastImage              *String    `xmlrpc:"rating_last_image,omptempty"`
	RatingLastText               *Selection `xmlrpc:"rating_last_text,omptempty"`
	RatingLastValue              *Float     `xmlrpc:"rating_last_value,omptempty"`
	RatingPercentageSatisfaction *Float     `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omptempty"`
}

// RatingMixins represents array of rating.mixin model.
type RatingMixins []RatingMixin

// RatingMixinModel is the odoo model name.
const RatingMixinModel = "rating.mixin"

// Many2One convert RatingMixin to *Many2One.
func (rm *RatingMixin) Many2One() *Many2One {
	return NewMany2One(rm.Id.Get(), "")
}

// CreateRatingMixin creates a new rating.mixin model and returns its id.
func (c *Client) CreateRatingMixin(rm *RatingMixin) (int64, error) {
	ids, err := c.CreateRatingMixins([]*RatingMixin{rm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRatingMixin creates a new rating.mixin model and returns its id.
func (c *Client) CreateRatingMixins(rms []*RatingMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range rms {
		vv = append(vv, v)
	}
	return c.Create(RatingMixinModel, vv)
}

// UpdateRatingMixin updates an existing rating.mixin record.
func (c *Client) UpdateRatingMixin(rm *RatingMixin) error {
	return c.UpdateRatingMixins([]int64{rm.Id.Get()}, rm)
}

// UpdateRatingMixins updates existing rating.mixin records.
// All records (represented by ids) will be updated by rm values.
func (c *Client) UpdateRatingMixins(ids []int64, rm *RatingMixin) error {
	return c.Update(RatingMixinModel, ids, rm)
}

// DeleteRatingMixin deletes an existing rating.mixin record.
func (c *Client) DeleteRatingMixin(id int64) error {
	return c.DeleteRatingMixins([]int64{id})
}

// DeleteRatingMixins deletes existing rating.mixin records.
func (c *Client) DeleteRatingMixins(ids []int64) error {
	return c.Delete(RatingMixinModel, ids)
}

// GetRatingMixin gets rating.mixin existing record.
func (c *Client) GetRatingMixin(id int64) (*RatingMixin, error) {
	rms, err := c.GetRatingMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if rms != nil && len(*rms) > 0 {
		return &((*rms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of rating.mixin not found", id)
}

// GetRatingMixins gets rating.mixin existing records.
func (c *Client) GetRatingMixins(ids []int64) (*RatingMixins, error) {
	rms := &RatingMixins{}
	if err := c.Read(RatingMixinModel, ids, nil, rms); err != nil {
		return nil, err
	}
	return rms, nil
}

// FindRatingMixin finds rating.mixin record by querying it with criteria.
func (c *Client) FindRatingMixin(criteria *Criteria) (*RatingMixin, error) {
	rms := &RatingMixins{}
	if err := c.SearchRead(RatingMixinModel, criteria, NewOptions().Limit(1), rms); err != nil {
		return nil, err
	}
	if rms != nil && len(*rms) > 0 {
		return &((*rms)[0]), nil
	}
	return nil, fmt.Errorf("rating.mixin was not found with criteria %v", criteria)
}

// FindRatingMixins finds rating.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRatingMixins(criteria *Criteria, options *Options) (*RatingMixins, error) {
	rms := &RatingMixins{}
	if err := c.SearchRead(RatingMixinModel, criteria, options, rms); err != nil {
		return nil, err
	}
	return rms, nil
}

// FindRatingMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRatingMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(RatingMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindRatingMixinId finds record id by querying it with criteria.
func (c *Client) FindRatingMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RatingMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("rating.mixin was not found with criteria %v and options %v", criteria, options)
}
