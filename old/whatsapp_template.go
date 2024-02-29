package odoo

import (
	"fmt"
)

// WhatsappTemplate represents whatsapp.template model.
type WhatsappTemplate struct {
	Active                   *Bool      `xmlrpc:"active,omptempty"`
	AllowedUserIds           *Relation  `xmlrpc:"allowed_user_ids,omptempty"`
	Body                     *String    `xmlrpc:"body,omptempty"`
	ButtonIds                *Relation  `xmlrpc:"button_ids,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	ErrorMsg                 *String    `xmlrpc:"error_msg,omptempty"`
	FooterText               *String    `xmlrpc:"footer_text,omptempty"`
	HasAction                *Bool      `xmlrpc:"has_action,omptempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omptempty"`
	HeaderAttachmentIds      *Relation  `xmlrpc:"header_attachment_ids,omptempty"`
	HeaderText               *String    `xmlrpc:"header_text,omptempty"`
	HeaderType               *Selection `xmlrpc:"header_type,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	LangCode                 *Selection `xmlrpc:"lang_code,omptempty"`
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
	MessagesCount            *Int       `xmlrpc:"messages_count,omptempty"`
	Model                    *String    `xmlrpc:"model,omptempty"`
	ModelId                  *Many2One  `xmlrpc:"model_id,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	PhoneField               *String    `xmlrpc:"phone_field,omptempty"`
	Quality                  *Selection `xmlrpc:"quality,omptempty"`
	RatingIds                *Relation  `xmlrpc:"rating_ids,omptempty"`
	ReportId                 *Many2One  `xmlrpc:"report_id,omptempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omptempty"`
	Status                   *Selection `xmlrpc:"status,omptempty"`
	TemplateName             *String    `xmlrpc:"template_name,omptempty"`
	TemplateType             *Selection `xmlrpc:"template_type,omptempty"`
	VariableIds              *Relation  `xmlrpc:"variable_ids,omptempty"`
	WaAccountId              *Many2One  `xmlrpc:"wa_account_id,omptempty"`
	WaTemplateUid            *String    `xmlrpc:"wa_template_uid,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// WhatsappTemplates represents array of whatsapp.template model.
type WhatsappTemplates []WhatsappTemplate

// WhatsappTemplateModel is the odoo model name.
const WhatsappTemplateModel = "whatsapp.template"

// Many2One convert WhatsappTemplate to *Many2One.
func (wt *WhatsappTemplate) Many2One() *Many2One {
	return NewMany2One(wt.Id.Get(), "")
}

// CreateWhatsappTemplate creates a new whatsapp.template model and returns its id.
func (c *Client) CreateWhatsappTemplate(wt *WhatsappTemplate) (int64, error) {
	ids, err := c.CreateWhatsappTemplates([]*WhatsappTemplate{wt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWhatsappTemplate creates a new whatsapp.template model and returns its id.
func (c *Client) CreateWhatsappTemplates(wts []*WhatsappTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range wts {
		vv = append(vv, v)
	}
	return c.Create(WhatsappTemplateModel, vv)
}

// UpdateWhatsappTemplate updates an existing whatsapp.template record.
func (c *Client) UpdateWhatsappTemplate(wt *WhatsappTemplate) error {
	return c.UpdateWhatsappTemplates([]int64{wt.Id.Get()}, wt)
}

// UpdateWhatsappTemplates updates existing whatsapp.template records.
// All records (represented by ids) will be updated by wt values.
func (c *Client) UpdateWhatsappTemplates(ids []int64, wt *WhatsappTemplate) error {
	return c.Update(WhatsappTemplateModel, ids, wt)
}

// DeleteWhatsappTemplate deletes an existing whatsapp.template record.
func (c *Client) DeleteWhatsappTemplate(id int64) error {
	return c.DeleteWhatsappTemplates([]int64{id})
}

// DeleteWhatsappTemplates deletes existing whatsapp.template records.
func (c *Client) DeleteWhatsappTemplates(ids []int64) error {
	return c.Delete(WhatsappTemplateModel, ids)
}

// GetWhatsappTemplate gets whatsapp.template existing record.
func (c *Client) GetWhatsappTemplate(id int64) (*WhatsappTemplate, error) {
	wts, err := c.GetWhatsappTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if wts != nil && len(*wts) > 0 {
		return &((*wts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of whatsapp.template not found", id)
}

// GetWhatsappTemplates gets whatsapp.template existing records.
func (c *Client) GetWhatsappTemplates(ids []int64) (*WhatsappTemplates, error) {
	wts := &WhatsappTemplates{}
	if err := c.Read(WhatsappTemplateModel, ids, nil, wts); err != nil {
		return nil, err
	}
	return wts, nil
}

// FindWhatsappTemplate finds whatsapp.template record by querying it with criteria.
func (c *Client) FindWhatsappTemplate(criteria *Criteria) (*WhatsappTemplate, error) {
	wts := &WhatsappTemplates{}
	if err := c.SearchRead(WhatsappTemplateModel, criteria, NewOptions().Limit(1), wts); err != nil {
		return nil, err
	}
	if wts != nil && len(*wts) > 0 {
		return &((*wts)[0]), nil
	}
	return nil, fmt.Errorf("whatsapp.template was not found with criteria %v", criteria)
}

// FindWhatsappTemplates finds whatsapp.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWhatsappTemplates(criteria *Criteria, options *Options) (*WhatsappTemplates, error) {
	wts := &WhatsappTemplates{}
	if err := c.SearchRead(WhatsappTemplateModel, criteria, options, wts); err != nil {
		return nil, err
	}
	return wts, nil
}

// FindWhatsappTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWhatsappTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WhatsappTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWhatsappTemplateId finds record id by querying it with criteria.
func (c *Client) FindWhatsappTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WhatsappTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("whatsapp.template was not found with criteria %v and options %v", criteria, options)
}
