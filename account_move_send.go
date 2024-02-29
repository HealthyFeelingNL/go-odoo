package odoo

import (
	"fmt"
)

// AccountMoveSend represents account.move.send model.
type AccountMoveSend struct {
	CheckboxDownload         *Bool       `xmlrpc:"checkbox_download,omptempty"`
	CheckboxSendByPost       *Bool       `xmlrpc:"checkbox_send_by_post,omptempty"`
	CheckboxSendMail         *Bool       `xmlrpc:"checkbox_send_mail,omptempty"`
	CheckboxUblCiiLabel      *String     `xmlrpc:"checkbox_ubl_cii_label,omptempty"`
	CheckboxUblCiiXml        *Bool       `xmlrpc:"checkbox_ubl_cii_xml,omptempty"`
	CompanyId                *Many2One   `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One   `xmlrpc:"create_uid,omptempty"`
	DisplayMailComposer      *Bool       `xmlrpc:"display_mail_composer,omptempty"`
	DisplayName              *String     `xmlrpc:"display_name,omptempty"`
	EnableDownload           *Bool       `xmlrpc:"enable_download,omptempty"`
	EnableSendByPost         *Bool       `xmlrpc:"enable_send_by_post,omptempty"`
	EnableSendMail           *Bool       `xmlrpc:"enable_send_mail,omptempty"`
	EnableUblCiiXml          *Bool       `xmlrpc:"enable_ubl_cii_xml,omptempty"`
	Id                       *Int        `xmlrpc:"id,omptempty"`
	MailAttachmentsWidget    interface{} `xmlrpc:"mail_attachments_widget,omptempty"`
	MailBody                 *String     `xmlrpc:"mail_body,omptempty"`
	MailLang                 *String     `xmlrpc:"mail_lang,omptempty"`
	MailPartnerIds           *Relation   `xmlrpc:"mail_partner_ids,omptempty"`
	MailSubject              *String     `xmlrpc:"mail_subject,omptempty"`
	MailTemplateId           *Many2One   `xmlrpc:"mail_template_id,omptempty"`
	Mode                     *Selection  `xmlrpc:"mode,omptempty"`
	MoveIds                  *Relation   `xmlrpc:"move_ids,omptempty"`
	SendByPostCost           *Int        `xmlrpc:"send_by_post_cost,omptempty"`
	SendByPostReadonly       *Bool       `xmlrpc:"send_by_post_readonly,omptempty"`
	SendByPostWarningMessage *String     `xmlrpc:"send_by_post_warning_message,omptempty"`
	SendMailReadonly         *Bool       `xmlrpc:"send_mail_readonly,omptempty"`
	SendMailWarningMessage   interface{} `xmlrpc:"send_mail_warning_message,omptempty"`
	WriteDate                *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// AccountMoveSends represents array of account.move.send model.
type AccountMoveSends []AccountMoveSend

// AccountMoveSendModel is the odoo model name.
const AccountMoveSendModel = "account.move.send"

// Many2One convert AccountMoveSend to *Many2One.
func (ams *AccountMoveSend) Many2One() *Many2One {
	return NewMany2One(ams.Id.Get(), "")
}

// CreateAccountMoveSend creates a new account.move.send model and returns its id.
func (c *Client) CreateAccountMoveSend(ams *AccountMoveSend) (int64, error) {
	ids, err := c.CreateAccountMoveSends([]*AccountMoveSend{ams})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMoveSend creates a new account.move.send model and returns its id.
func (c *Client) CreateAccountMoveSends(amss []*AccountMoveSend) ([]int64, error) {
	var vv []interface{}
	for _, v := range amss {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveSendModel, vv)
}

// UpdateAccountMoveSend updates an existing account.move.send record.
func (c *Client) UpdateAccountMoveSend(ams *AccountMoveSend) error {
	return c.UpdateAccountMoveSends([]int64{ams.Id.Get()}, ams)
}

// UpdateAccountMoveSends updates existing account.move.send records.
// All records (represented by ids) will be updated by ams values.
func (c *Client) UpdateAccountMoveSends(ids []int64, ams *AccountMoveSend) error {
	return c.Update(AccountMoveSendModel, ids, ams)
}

// DeleteAccountMoveSend deletes an existing account.move.send record.
func (c *Client) DeleteAccountMoveSend(id int64) error {
	return c.DeleteAccountMoveSends([]int64{id})
}

// DeleteAccountMoveSends deletes existing account.move.send records.
func (c *Client) DeleteAccountMoveSends(ids []int64) error {
	return c.Delete(AccountMoveSendModel, ids)
}

// GetAccountMoveSend gets account.move.send existing record.
func (c *Client) GetAccountMoveSend(id int64) (*AccountMoveSend, error) {
	amss, err := c.GetAccountMoveSends([]int64{id})
	if err != nil {
		return nil, err
	}
	if amss != nil && len(*amss) > 0 {
		return &((*amss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.move.send not found", id)
}

// GetAccountMoveSends gets account.move.send existing records.
func (c *Client) GetAccountMoveSends(ids []int64) (*AccountMoveSends, error) {
	amss := &AccountMoveSends{}
	if err := c.Read(AccountMoveSendModel, ids, nil, amss); err != nil {
		return nil, err
	}
	return amss, nil
}

// FindAccountMoveSend finds account.move.send record by querying it with criteria.
func (c *Client) FindAccountMoveSend(criteria *Criteria) (*AccountMoveSend, error) {
	amss := &AccountMoveSends{}
	if err := c.SearchRead(AccountMoveSendModel, criteria, NewOptions().Limit(1), amss); err != nil {
		return nil, err
	}
	if amss != nil && len(*amss) > 0 {
		return &((*amss)[0]), nil
	}
	return nil, fmt.Errorf("account.move.send was not found with criteria %v", criteria)
}

// FindAccountMoveSends finds account.move.send records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveSends(criteria *Criteria, options *Options) (*AccountMoveSends, error) {
	amss := &AccountMoveSends{}
	if err := c.SearchRead(AccountMoveSendModel, criteria, options, amss); err != nil {
		return nil, err
	}
	return amss, nil
}

// FindAccountMoveSendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveSendIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountMoveSendModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountMoveSendId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveSendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveSendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.move.send was not found with criteria %v and options %v", criteria, options)
}
