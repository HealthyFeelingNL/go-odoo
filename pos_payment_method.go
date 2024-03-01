package odoo

import (
	"fmt"
)

// PosPaymentMethod represents pos.payment.method model.
type PosPaymentMethod struct {
	Active                     *Bool      `xmlrpc:"active,omptempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omptempty"`
	ConfigIds                  *Relation  `xmlrpc:"config_ids,omptempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                *String    `xmlrpc:"display_name,omptempty"`
	HasAnOnlinePaymentProvider *Bool      `xmlrpc:"has_an_online_payment_provider,omptempty"`
	HideUsePaymentTerminal     *Bool      `xmlrpc:"hide_use_payment_terminal,omptempty"`
	Id                         *Int       `xmlrpc:"id,omptempty"`
	Image                      *String    `xmlrpc:"image,omptempty"`
	IotDeviceId                *Many2One  `xmlrpc:"iot_device_id,omptempty"`
	IsCashCount                *Bool      `xmlrpc:"is_cash_count,omptempty"`
	IsOnlinePayment            *Bool      `xmlrpc:"is_online_payment,omptempty"`
	JournalId                  *Many2One  `xmlrpc:"journal_id,omptempty"`
	Name                       *String    `xmlrpc:"name,omptempty"`
	OnlinePaymentProviderIds   *Relation  `xmlrpc:"online_payment_provider_ids,omptempty"`
	OpenSessionIds             *Relation  `xmlrpc:"open_session_ids,omptempty"`
	OutstandingAccountId       *Many2One  `xmlrpc:"outstanding_account_id,omptempty"`
	PaymentTerminalIds         *Relation  `xmlrpc:"payment_terminal_ids,omptempty"`
	ReceivableAccountId        *Many2One  `xmlrpc:"receivable_account_id,omptempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omptempty"`
	SplitTransactions          *Bool      `xmlrpc:"split_transactions,omptempty"`
	StripeSerialNumber         *String    `xmlrpc:"stripe_serial_number,omptempty"`
	Type                       *Selection `xmlrpc:"type,omptempty"`
	UsePaymentTerminal         *Selection `xmlrpc:"use_payment_terminal,omptempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PosPaymentMethods represents array of pos.payment.method model.
type PosPaymentMethods []PosPaymentMethod

// PosPaymentMethodModel is the odoo model name.
const PosPaymentMethodModel = "pos.payment.method"

// Many2One convert PosPaymentMethod to *Many2One.
func (ppm *PosPaymentMethod) Many2One() *Many2One {
	return NewMany2One(ppm.Id.Get(), "")
}

// CreatePosPaymentMethod creates a new pos.payment.method model and returns its id.
func (c *Client) CreatePosPaymentMethod(ppm *PosPaymentMethod) (int64, error) {
	ids, err := c.CreatePosPaymentMethods([]*PosPaymentMethod{ppm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPaymentMethod creates a new pos.payment.method model and returns its id.
func (c *Client) CreatePosPaymentMethods(ppms []*PosPaymentMethod) ([]int64, error) {
	var vv []interface{}
	for _, v := range ppms {
		vv = append(vv, v)
	}
	return c.Create(PosPaymentMethodModel, vv)
}

// UpdatePosPaymentMethod updates an existing pos.payment.method record.
func (c *Client) UpdatePosPaymentMethod(ppm *PosPaymentMethod) error {
	return c.UpdatePosPaymentMethods([]int64{ppm.Id.Get()}, ppm)
}

// UpdatePosPaymentMethods updates existing pos.payment.method records.
// All records (represented by ids) will be updated by ppm values.
func (c *Client) UpdatePosPaymentMethods(ids []int64, ppm *PosPaymentMethod) error {
	return c.Update(PosPaymentMethodModel, ids, ppm)
}

// DeletePosPaymentMethod deletes an existing pos.payment.method record.
func (c *Client) DeletePosPaymentMethod(id int64) error {
	return c.DeletePosPaymentMethods([]int64{id})
}

// DeletePosPaymentMethods deletes existing pos.payment.method records.
func (c *Client) DeletePosPaymentMethods(ids []int64) error {
	return c.Delete(PosPaymentMethodModel, ids)
}

// GetPosPaymentMethod gets pos.payment.method existing record.
func (c *Client) GetPosPaymentMethod(id int64) (*PosPaymentMethod, error) {
	ppms, err := c.GetPosPaymentMethods([]int64{id})
	if err != nil {
		return nil, err
	}
	if ppms != nil && len(*ppms) > 0 {
		return &((*ppms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.payment.method not found", id)
}

// GetPosPaymentMethods gets pos.payment.method existing records.
func (c *Client) GetPosPaymentMethods(ids []int64) (*PosPaymentMethods, error) {
	ppms := &PosPaymentMethods{}
	if err := c.Read(PosPaymentMethodModel, ids, nil, ppms); err != nil {
		return nil, err
	}
	return ppms, nil
}

// FindPosPaymentMethod finds pos.payment.method record by querying it with criteria.
func (c *Client) FindPosPaymentMethod(criteria *Criteria) (*PosPaymentMethod, error) {
	ppms := &PosPaymentMethods{}
	if err := c.SearchRead(PosPaymentMethodModel, criteria, NewOptions().Limit(1), ppms); err != nil {
		return nil, err
	}
	if ppms != nil && len(*ppms) > 0 {
		return &((*ppms)[0]), nil
	}
	return nil, fmt.Errorf("pos.payment.method was not found with criteria %v", criteria)
}

// FindPosPaymentMethods finds pos.payment.method records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPaymentMethods(criteria *Criteria, options *Options) (*PosPaymentMethods, error) {
	ppms := &PosPaymentMethods{}
	if err := c.SearchRead(PosPaymentMethodModel, criteria, options, ppms); err != nil {
		return nil, err
	}
	return ppms, nil
}

// FindPosPaymentMethodIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPaymentMethodIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosPaymentMethodModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosPaymentMethodId finds record id by querying it with criteria.
func (c *Client) FindPosPaymentMethodId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPaymentMethodModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.payment.method was not found with criteria %v and options %v", criteria, options)
}
