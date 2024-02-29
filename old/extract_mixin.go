package odoo

import (
	"fmt"
)

// ExtractMixin represents extract.mixin model.
type ExtractMixin struct {
	ExtractCanShowSendButton *Bool      `xmlrpc:"extract_can_show_send_button,omptempty"`
	ExtractDocumentUuid      *String    `xmlrpc:"extract_document_uuid,omptempty"`
	ExtractErrorMessage      *String    `xmlrpc:"extract_error_message,omptempty"`
	ExtractState             *Selection `xmlrpc:"extract_state,omptempty"`
	ExtractStateProcessed    *Bool      `xmlrpc:"extract_state_processed,omptempty"`
	ExtractStatus            *String    `xmlrpc:"extract_status,omptempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omptempty"`
	IsInExtractableState     *Bool      `xmlrpc:"is_in_extractable_state,omptempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	RatingIds                *Relation  `xmlrpc:"rating_ids,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
}

// ExtractMixins represents array of extract.mixin model.
type ExtractMixins []ExtractMixin

// ExtractMixinModel is the odoo model name.
const ExtractMixinModel = "extract.mixin"

// Many2One convert ExtractMixin to *Many2One.
func (em *ExtractMixin) Many2One() *Many2One {
	return NewMany2One(em.Id.Get(), "")
}

// CreateExtractMixin creates a new extract.mixin model and returns its id.
func (c *Client) CreateExtractMixin(em *ExtractMixin) (int64, error) {
	ids, err := c.CreateExtractMixins([]*ExtractMixin{em})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateExtractMixin creates a new extract.mixin model and returns its id.
func (c *Client) CreateExtractMixins(ems []*ExtractMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range ems {
		vv = append(vv, v)
	}
	return c.Create(ExtractMixinModel, vv)
}

// UpdateExtractMixin updates an existing extract.mixin record.
func (c *Client) UpdateExtractMixin(em *ExtractMixin) error {
	return c.UpdateExtractMixins([]int64{em.Id.Get()}, em)
}

// UpdateExtractMixins updates existing extract.mixin records.
// All records (represented by ids) will be updated by em values.
func (c *Client) UpdateExtractMixins(ids []int64, em *ExtractMixin) error {
	return c.Update(ExtractMixinModel, ids, em)
}

// DeleteExtractMixin deletes an existing extract.mixin record.
func (c *Client) DeleteExtractMixin(id int64) error {
	return c.DeleteExtractMixins([]int64{id})
}

// DeleteExtractMixins deletes existing extract.mixin records.
func (c *Client) DeleteExtractMixins(ids []int64) error {
	return c.Delete(ExtractMixinModel, ids)
}

// GetExtractMixin gets extract.mixin existing record.
func (c *Client) GetExtractMixin(id int64) (*ExtractMixin, error) {
	ems, err := c.GetExtractMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if ems != nil && len(*ems) > 0 {
		return &((*ems)[0]), nil
	}
	return nil, fmt.Errorf("id %v of extract.mixin not found", id)
}

// GetExtractMixins gets extract.mixin existing records.
func (c *Client) GetExtractMixins(ids []int64) (*ExtractMixins, error) {
	ems := &ExtractMixins{}
	if err := c.Read(ExtractMixinModel, ids, nil, ems); err != nil {
		return nil, err
	}
	return ems, nil
}

// FindExtractMixin finds extract.mixin record by querying it with criteria.
func (c *Client) FindExtractMixin(criteria *Criteria) (*ExtractMixin, error) {
	ems := &ExtractMixins{}
	if err := c.SearchRead(ExtractMixinModel, criteria, NewOptions().Limit(1), ems); err != nil {
		return nil, err
	}
	if ems != nil && len(*ems) > 0 {
		return &((*ems)[0]), nil
	}
	return nil, fmt.Errorf("extract.mixin was not found with criteria %v", criteria)
}

// FindExtractMixins finds extract.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindExtractMixins(criteria *Criteria, options *Options) (*ExtractMixins, error) {
	ems := &ExtractMixins{}
	if err := c.SearchRead(ExtractMixinModel, criteria, options, ems); err != nil {
		return nil, err
	}
	return ems, nil
}

// FindExtractMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindExtractMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ExtractMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindExtractMixinId finds record id by querying it with criteria.
func (c *Client) FindExtractMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ExtractMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("extract.mixin was not found with criteria %v and options %v", criteria, options)
}
