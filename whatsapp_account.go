package odoo

import (
	"fmt"
)

// WhatsappAccount represents whatsapp.account model.
type WhatsappAccount struct {
	AccountUid               *String   `xmlrpc:"account_uid,omptempty"`
	Active                   *Bool     `xmlrpc:"active,omptempty"`
	AllowedCompanyIds        *Relation `xmlrpc:"allowed_company_ids,omptempty"`
	AppSecret                *String   `xmlrpc:"app_secret,omptempty"`
	AppUid                   *String   `xmlrpc:"app_uid,omptempty"`
	CallbackUrl              *String   `xmlrpc:"callback_url,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	NotifyUserIds            *Relation `xmlrpc:"notify_user_ids,omptempty"`
	PhoneUid                 *String   `xmlrpc:"phone_uid,omptempty"`
	RatingIds                *Relation `xmlrpc:"rating_ids,omptempty"`
	TemplatesCount           *Int      `xmlrpc:"templates_count,omptempty"`
	Token                    *String   `xmlrpc:"token,omptempty"`
	WebhookVerifyToken       *String   `xmlrpc:"webhook_verify_token,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// WhatsappAccounts represents array of whatsapp.account model.
type WhatsappAccounts []WhatsappAccount

// WhatsappAccountModel is the odoo model name.
const WhatsappAccountModel = "whatsapp.account"

// Many2One convert WhatsappAccount to *Many2One.
func (wa *WhatsappAccount) Many2One() *Many2One {
	return NewMany2One(wa.Id.Get(), "")
}

// CreateWhatsappAccount creates a new whatsapp.account model and returns its id.
func (c *Client) CreateWhatsappAccount(wa *WhatsappAccount) (int64, error) {
	ids, err := c.CreateWhatsappAccounts([]*WhatsappAccount{wa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWhatsappAccount creates a new whatsapp.account model and returns its id.
func (c *Client) CreateWhatsappAccounts(was []*WhatsappAccount) ([]int64, error) {
	var vv []interface{}
	for _, v := range was {
		vv = append(vv, v)
	}
	return c.Create(WhatsappAccountModel, vv)
}

// UpdateWhatsappAccount updates an existing whatsapp.account record.
func (c *Client) UpdateWhatsappAccount(wa *WhatsappAccount) error {
	return c.UpdateWhatsappAccounts([]int64{wa.Id.Get()}, wa)
}

// UpdateWhatsappAccounts updates existing whatsapp.account records.
// All records (represented by ids) will be updated by wa values.
func (c *Client) UpdateWhatsappAccounts(ids []int64, wa *WhatsappAccount) error {
	return c.Update(WhatsappAccountModel, ids, wa)
}

// DeleteWhatsappAccount deletes an existing whatsapp.account record.
func (c *Client) DeleteWhatsappAccount(id int64) error {
	return c.DeleteWhatsappAccounts([]int64{id})
}

// DeleteWhatsappAccounts deletes existing whatsapp.account records.
func (c *Client) DeleteWhatsappAccounts(ids []int64) error {
	return c.Delete(WhatsappAccountModel, ids)
}

// GetWhatsappAccount gets whatsapp.account existing record.
func (c *Client) GetWhatsappAccount(id int64) (*WhatsappAccount, error) {
	was, err := c.GetWhatsappAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	if was != nil && len(*was) > 0 {
		return &((*was)[0]), nil
	}
	return nil, fmt.Errorf("id %v of whatsapp.account not found", id)
}

// GetWhatsappAccounts gets whatsapp.account existing records.
func (c *Client) GetWhatsappAccounts(ids []int64) (*WhatsappAccounts, error) {
	was := &WhatsappAccounts{}
	if err := c.Read(WhatsappAccountModel, ids, nil, was); err != nil {
		return nil, err
	}
	return was, nil
}

// FindWhatsappAccount finds whatsapp.account record by querying it with criteria.
func (c *Client) FindWhatsappAccount(criteria *Criteria) (*WhatsappAccount, error) {
	was := &WhatsappAccounts{}
	if err := c.SearchRead(WhatsappAccountModel, criteria, NewOptions().Limit(1), was); err != nil {
		return nil, err
	}
	if was != nil && len(*was) > 0 {
		return &((*was)[0]), nil
	}
	return nil, fmt.Errorf("whatsapp.account was not found with criteria %v", criteria)
}

// FindWhatsappAccounts finds whatsapp.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWhatsappAccounts(criteria *Criteria, options *Options) (*WhatsappAccounts, error) {
	was := &WhatsappAccounts{}
	if err := c.SearchRead(WhatsappAccountModel, criteria, options, was); err != nil {
		return nil, err
	}
	return was, nil
}

// FindWhatsappAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWhatsappAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WhatsappAccountModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWhatsappAccountId finds record id by querying it with criteria.
func (c *Client) FindWhatsappAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WhatsappAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("whatsapp.account was not found with criteria %v and options %v", criteria, options)
}
