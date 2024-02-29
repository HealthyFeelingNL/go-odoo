package odoo

import (
	"fmt"
)

// MailThreadMainAttachment represents mail.thread.main.attachment model.
type MailThreadMainAttachment struct {
	HasMessage               *Bool     `xmlrpc:"has_message,omptempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	RatingIds                *Relation `xmlrpc:"rating_ids,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
}

// MailThreadMainAttachments represents array of mail.thread.main.attachment model.
type MailThreadMainAttachments []MailThreadMainAttachment

// MailThreadMainAttachmentModel is the odoo model name.
const MailThreadMainAttachmentModel = "mail.thread.main.attachment"

// Many2One convert MailThreadMainAttachment to *Many2One.
func (mtma *MailThreadMainAttachment) Many2One() *Many2One {
	return NewMany2One(mtma.Id.Get(), "")
}

// CreateMailThreadMainAttachment creates a new mail.thread.main.attachment model and returns its id.
func (c *Client) CreateMailThreadMainAttachment(mtma *MailThreadMainAttachment) (int64, error) {
	ids, err := c.CreateMailThreadMainAttachments([]*MailThreadMainAttachment{mtma})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailThreadMainAttachment creates a new mail.thread.main.attachment model and returns its id.
func (c *Client) CreateMailThreadMainAttachments(mtmas []*MailThreadMainAttachment) ([]int64, error) {
	var vv []interface{}
	for _, v := range mtmas {
		vv = append(vv, v)
	}
	return c.Create(MailThreadMainAttachmentModel, vv)
}

// UpdateMailThreadMainAttachment updates an existing mail.thread.main.attachment record.
func (c *Client) UpdateMailThreadMainAttachment(mtma *MailThreadMainAttachment) error {
	return c.UpdateMailThreadMainAttachments([]int64{mtma.Id.Get()}, mtma)
}

// UpdateMailThreadMainAttachments updates existing mail.thread.main.attachment records.
// All records (represented by ids) will be updated by mtma values.
func (c *Client) UpdateMailThreadMainAttachments(ids []int64, mtma *MailThreadMainAttachment) error {
	return c.Update(MailThreadMainAttachmentModel, ids, mtma)
}

// DeleteMailThreadMainAttachment deletes an existing mail.thread.main.attachment record.
func (c *Client) DeleteMailThreadMainAttachment(id int64) error {
	return c.DeleteMailThreadMainAttachments([]int64{id})
}

// DeleteMailThreadMainAttachments deletes existing mail.thread.main.attachment records.
func (c *Client) DeleteMailThreadMainAttachments(ids []int64) error {
	return c.Delete(MailThreadMainAttachmentModel, ids)
}

// GetMailThreadMainAttachment gets mail.thread.main.attachment existing record.
func (c *Client) GetMailThreadMainAttachment(id int64) (*MailThreadMainAttachment, error) {
	mtmas, err := c.GetMailThreadMainAttachments([]int64{id})
	if err != nil {
		return nil, err
	}
	if mtmas != nil && len(*mtmas) > 0 {
		return &((*mtmas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.thread.main.attachment not found", id)
}

// GetMailThreadMainAttachments gets mail.thread.main.attachment existing records.
func (c *Client) GetMailThreadMainAttachments(ids []int64) (*MailThreadMainAttachments, error) {
	mtmas := &MailThreadMainAttachments{}
	if err := c.Read(MailThreadMainAttachmentModel, ids, nil, mtmas); err != nil {
		return nil, err
	}
	return mtmas, nil
}

// FindMailThreadMainAttachment finds mail.thread.main.attachment record by querying it with criteria.
func (c *Client) FindMailThreadMainAttachment(criteria *Criteria) (*MailThreadMainAttachment, error) {
	mtmas := &MailThreadMainAttachments{}
	if err := c.SearchRead(MailThreadMainAttachmentModel, criteria, NewOptions().Limit(1), mtmas); err != nil {
		return nil, err
	}
	if mtmas != nil && len(*mtmas) > 0 {
		return &((*mtmas)[0]), nil
	}
	return nil, fmt.Errorf("mail.thread.main.attachment was not found with criteria %v", criteria)
}

// FindMailThreadMainAttachments finds mail.thread.main.attachment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailThreadMainAttachments(criteria *Criteria, options *Options) (*MailThreadMainAttachments, error) {
	mtmas := &MailThreadMainAttachments{}
	if err := c.SearchRead(MailThreadMainAttachmentModel, criteria, options, mtmas); err != nil {
		return nil, err
	}
	return mtmas, nil
}

// FindMailThreadMainAttachmentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailThreadMainAttachmentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailThreadMainAttachmentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailThreadMainAttachmentId finds record id by querying it with criteria.
func (c *Client) FindMailThreadMainAttachmentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailThreadMainAttachmentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.thread.main.attachment was not found with criteria %v and options %v", criteria, options)
}
