package odoo

import (
	"fmt"
)

// AccountPayment represents account.payment model.
type AccountPayment struct {
	AccessToken                           *String     `xmlrpc:"access_token,omptempty"`
	AccessUrl                             *String     `xmlrpc:"access_url,omptempty"`
	AccessWarning                         *String     `xmlrpc:"access_warning,omptempty"`
	ActivityCalendarEventId               *Many2One   `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline                  *Time       `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration           *Selection  `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon                 *String     `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                           *Relation   `xmlrpc:"activity_ids,omptempty"`
	ActivityState                         *Selection  `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                       *String     `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                      *String     `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                        *Many2One   `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                        *Many2One   `xmlrpc:"activity_user_id,omptempty"`
	AlwaysTaxExigible                     *Bool       `xmlrpc:"always_tax_exigible,omptempty"`
	Amount                                *Float      `xmlrpc:"amount,omptempty"`
	AmountAvailableForRefund              *Float      `xmlrpc:"amount_available_for_refund,omptempty"`
	AmountCompanyCurrencySigned           *Float      `xmlrpc:"amount_company_currency_signed,omptempty"`
	AmountPaid                            *Float      `xmlrpc:"amount_paid,omptempty"`
	AmountResidual                        *Float      `xmlrpc:"amount_residual,omptempty"`
	AmountResidualSigned                  *Float      `xmlrpc:"amount_residual_signed,omptempty"`
	AmountSigned                          *Float      `xmlrpc:"amount_signed,omptempty"`
	AmountTax                             *Float      `xmlrpc:"amount_tax,omptempty"`
	AmountTaxSigned                       *Float      `xmlrpc:"amount_tax_signed,omptempty"`
	AmountTotal                           *Float      `xmlrpc:"amount_total,omptempty"`
	AmountTotalInCurrencySigned           *Float      `xmlrpc:"amount_total_in_currency_signed,omptempty"`
	AmountTotalSigned                     *Float      `xmlrpc:"amount_total_signed,omptempty"`
	AmountTotalWords                      *String     `xmlrpc:"amount_total_words,omptempty"`
	AmountUntaxed                         *Float      `xmlrpc:"amount_untaxed,omptempty"`
	AmountUntaxedSigned                   *Float      `xmlrpc:"amount_untaxed_signed,omptempty"`
	AttachmentIds                         *Relation   `xmlrpc:"attachment_ids,omptempty"`
	AuthorizedTransactionIds              *Relation   `xmlrpc:"authorized_transaction_ids,omptempty"`
	AutoPost                              *Selection  `xmlrpc:"auto_post,omptempty"`
	AutoPostOriginId                      *Many2One   `xmlrpc:"auto_post_origin_id,omptempty"`
	AutoPostUntil                         *Time       `xmlrpc:"auto_post_until,omptempty"`
	AvailableJournalIds                   *Relation   `xmlrpc:"available_journal_ids,omptempty"`
	AvailablePartnerBankIds               *Relation   `xmlrpc:"available_partner_bank_ids,omptempty"`
	AvailablePaymentMethodLineIds         *Relation   `xmlrpc:"available_payment_method_line_ids,omptempty"`
	BankPartnerId                         *Many2One   `xmlrpc:"bank_partner_id,omptempty"`
	BatchPaymentId                        *Many2One   `xmlrpc:"batch_payment_id,omptempty"`
	CampaignId                            *Many2One   `xmlrpc:"campaign_id,omptempty"`
	CommercialPartnerId                   *Many2One   `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyCurrencyId                     *Many2One   `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                             *Many2One   `xmlrpc:"company_id,omptempty"`
	CountryCode                           *String     `xmlrpc:"country_code,omptempty"`
	CreateDate                            *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                             *Many2One   `xmlrpc:"create_uid,omptempty"`
	CurrencyId                            *Many2One   `xmlrpc:"currency_id,omptempty"`
	Date                                  *Time       `xmlrpc:"date,omptempty"`
	DeliveryDate                          *Time       `xmlrpc:"delivery_date,omptempty"`
	DestinationAccountId                  *Many2One   `xmlrpc:"destination_account_id,omptempty"`
	DestinationJournalId                  *Many2One   `xmlrpc:"destination_journal_id,omptempty"`
	DirectionSign                         *Int        `xmlrpc:"direction_sign,omptempty"`
	DisplayInactiveCurrencyWarning        *Bool       `xmlrpc:"display_inactive_currency_warning,omptempty"`
	DisplayName                           *String     `xmlrpc:"display_name,omptempty"`
	DisplayQrCode                         *Bool       `xmlrpc:"display_qr_code,omptempty"`
	DuplicatedRefIds                      *Relation   `xmlrpc:"duplicated_ref_ids,omptempty"`
	ExtractAttachmentId                   *Many2One   `xmlrpc:"extract_attachment_id,omptempty"`
	ExtractCanShowBanners                 *Bool       `xmlrpc:"extract_can_show_banners,omptempty"`
	ExtractCanShowSendButton              *Bool       `xmlrpc:"extract_can_show_send_button,omptempty"`
	ExtractDetectedLayout                 *Int        `xmlrpc:"extract_detected_layout,omptempty"`
	ExtractDocumentUuid                   *String     `xmlrpc:"extract_document_uuid,omptempty"`
	ExtractErrorMessage                   *String     `xmlrpc:"extract_error_message,omptempty"`
	ExtractPartnerName                    *String     `xmlrpc:"extract_partner_name,omptempty"`
	ExtractState                          *Selection  `xmlrpc:"extract_state,omptempty"`
	ExtractStateProcessed                 *Bool       `xmlrpc:"extract_state_processed,omptempty"`
	ExtractStatus                         *String     `xmlrpc:"extract_status,omptempty"`
	ExtractWordIds                        *Relation   `xmlrpc:"extract_word_ids,omptempty"`
	FiscalPositionId                      *Many2One   `xmlrpc:"fiscal_position_id,omptempty"`
	HasMessage                            *Bool       `xmlrpc:"has_message,omptempty"`
	HasReconciledEntries                  *Bool       `xmlrpc:"has_reconciled_entries,omptempty"`
	HidePostButton                        *Bool       `xmlrpc:"hide_post_button,omptempty"`
	HighestName                           *String     `xmlrpc:"highest_name,omptempty"`
	Id                                    *Int        `xmlrpc:"id,omptempty"`
	InalterableHash                       *String     `xmlrpc:"inalterable_hash,omptempty"`
	IncotermLocation                      *String     `xmlrpc:"incoterm_location,omptempty"`
	InvoiceCashRoundingId                 *Many2One   `xmlrpc:"invoice_cash_rounding_id,omptempty"`
	InvoiceDate                           *Time       `xmlrpc:"invoice_date,omptempty"`
	InvoiceDateDue                        *Time       `xmlrpc:"invoice_date_due,omptempty"`
	InvoiceFilterTypeDomain               *String     `xmlrpc:"invoice_filter_type_domain,omptempty"`
	InvoiceHasOutstanding                 *Bool       `xmlrpc:"invoice_has_outstanding,omptempty"`
	InvoiceIncotermId                     *Many2One   `xmlrpc:"invoice_incoterm_id,omptempty"`
	InvoiceLineIds                        *Relation   `xmlrpc:"invoice_line_ids,omptempty"`
	InvoiceOrigin                         *String     `xmlrpc:"invoice_origin,omptempty"`
	InvoiceOutstandingCreditsDebitsWidget *String     `xmlrpc:"invoice_outstanding_credits_debits_widget,omptempty"`
	InvoicePartnerDisplayName             *String     `xmlrpc:"invoice_partner_display_name,omptempty"`
	InvoicePaymentTermId                  *Many2One   `xmlrpc:"invoice_payment_term_id,omptempty"`
	InvoicePaymentsWidget                 *String     `xmlrpc:"invoice_payments_widget,omptempty"`
	InvoicePdfReportFile                  *String     `xmlrpc:"invoice_pdf_report_file,omptempty"`
	InvoicePdfReportId                    *Many2One   `xmlrpc:"invoice_pdf_report_id,omptempty"`
	InvoiceSourceEmail                    *String     `xmlrpc:"invoice_source_email,omptempty"`
	InvoiceUserId                         *Many2One   `xmlrpc:"invoice_user_id,omptempty"`
	InvoiceVendorBillId                   *Many2One   `xmlrpc:"invoice_vendor_bill_id,omptempty"`
	IsBeingSent                           *Bool       `xmlrpc:"is_being_sent,omptempty"`
	IsDonation                            *Bool       `xmlrpc:"is_donation,omptempty"`
	IsInExtractableState                  *Bool       `xmlrpc:"is_in_extractable_state,omptempty"`
	IsInternalTransfer                    *Bool       `xmlrpc:"is_internal_transfer,omptempty"`
	IsMatched                             *Bool       `xmlrpc:"is_matched,omptempty"`
	IsMoveSent                            *Bool       `xmlrpc:"is_move_sent,omptempty"`
	IsReconciled                          *Bool       `xmlrpc:"is_reconciled,omptempty"`
	IsStorno                              *Bool       `xmlrpc:"is_storno,omptempty"`
	JournalId                             *Many2One   `xmlrpc:"journal_id,omptempty"`
	LineIds                               *Relation   `xmlrpc:"line_ids,omptempty"`
	MadeSequenceHole                      *Bool       `xmlrpc:"made_sequence_hole,omptempty"`
	MediumId                              *Many2One   `xmlrpc:"medium_id,omptempty"`
	MessageAttachmentCount                *Int        `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds                    *Relation   `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                       *Bool       `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter                *Int        `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                    *Bool       `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                            *Relation   `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                     *Bool       `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId               *Many2One   `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                     *Bool       `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter              *Int        `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                     *Relation   `xmlrpc:"message_partner_ids,omptempty"`
	MoveId                                *Many2One   `xmlrpc:"move_id,omptempty"`
	MoveType                              *Selection  `xmlrpc:"move_type,omptempty"`
	MyActivityDateDeadline                *Time       `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                                  *String     `xmlrpc:"name,omptempty"`
	Narration                             *String     `xmlrpc:"narration,omptempty"`
	NeedCancelRequest                     *Bool       `xmlrpc:"need_cancel_request,omptempty"`
	NeededTerms                           *String     `xmlrpc:"needed_terms,omptempty"`
	NeededTermsDirty                      *Bool       `xmlrpc:"needed_terms_dirty,omptempty"`
	OutstandingAccountId                  *Many2One   `xmlrpc:"outstanding_account_id,omptempty"`
	PairedInternalTransferPaymentId       *Many2One   `xmlrpc:"paired_internal_transfer_payment_id,omptempty"`
	PartnerBankId                         *Many2One   `xmlrpc:"partner_bank_id,omptempty"`
	PartnerCredit                         *Float      `xmlrpc:"partner_credit,omptempty"`
	PartnerCreditWarning                  *String     `xmlrpc:"partner_credit_warning,omptempty"`
	PartnerId                             *Many2One   `xmlrpc:"partner_id,omptempty"`
	PartnerShippingId                     *Many2One   `xmlrpc:"partner_shipping_id,omptempty"`
	PartnerType                           *Selection  `xmlrpc:"partner_type,omptempty"`
	PaymentId                             *Many2One   `xmlrpc:"payment_id,omptempty"`
	PaymentIds                            *Relation   `xmlrpc:"payment_ids,omptempty"`
	PaymentMethodCode                     *String     `xmlrpc:"payment_method_code,omptempty"`
	PaymentMethodId                       *Many2One   `xmlrpc:"payment_method_id,omptempty"`
	PaymentMethodLineId                   *Many2One   `xmlrpc:"payment_method_line_id,omptempty"`
	PaymentMethodName                     *String     `xmlrpc:"payment_method_name,omptempty"`
	PaymentReference                      *String     `xmlrpc:"payment_reference,omptempty"`
	PaymentState                          *Selection  `xmlrpc:"payment_state,omptempty"`
	PaymentTermDetails                    *String     `xmlrpc:"payment_term_details,omptempty"`
	PaymentTokenId                        *Many2One   `xmlrpc:"payment_token_id,omptempty"`
	PaymentTransactionId                  *Many2One   `xmlrpc:"payment_transaction_id,omptempty"`
	PaymentType                           *Selection  `xmlrpc:"payment_type,omptempty"`
	PostedBefore                          *Bool       `xmlrpc:"posted_before,omptempty"`
	QrCode                                *String     `xmlrpc:"qr_code,omptempty"`
	QrCodeMethod                          *Selection  `xmlrpc:"qr_code_method,omptempty"`
	QuickEditMode                         *Bool       `xmlrpc:"quick_edit_mode,omptempty"`
	QuickEditTotalAmount                  *Float      `xmlrpc:"quick_edit_total_amount,omptempty"`
	QuickEncodingVals                     *String     `xmlrpc:"quick_encoding_vals,omptempty"`
	RatingIds                             *Relation   `xmlrpc:"rating_ids,omptempty"`
	ReconciledBillIds                     *Relation   `xmlrpc:"reconciled_bill_ids,omptempty"`
	ReconciledBillsCount                  *Int        `xmlrpc:"reconciled_bills_count,omptempty"`
	ReconciledInvoiceIds                  *Relation   `xmlrpc:"reconciled_invoice_ids,omptempty"`
	ReconciledInvoicesCount               *Int        `xmlrpc:"reconciled_invoices_count,omptempty"`
	ReconciledInvoicesType                *Selection  `xmlrpc:"reconciled_invoices_type,omptempty"`
	ReconciledStatementLineIds            *Relation   `xmlrpc:"reconciled_statement_line_ids,omptempty"`
	ReconciledStatementLinesCount         *Int        `xmlrpc:"reconciled_statement_lines_count,omptempty"`
	Ref                                   *String     `xmlrpc:"ref,omptempty"`
	RefundsCount                          *Int        `xmlrpc:"refunds_count,omptempty"`
	RequirePartnerBankAccount             *Bool       `xmlrpc:"require_partner_bank_account,omptempty"`
	RestrictModeHashTable                 *Bool       `xmlrpc:"restrict_mode_hash_table,omptempty"`
	ReversalMoveId                        *Relation   `xmlrpc:"reversal_move_id,omptempty"`
	ReversedEntryId                       *Many2One   `xmlrpc:"reversed_entry_id,omptempty"`
	SaleOrderCount                        *Int        `xmlrpc:"sale_order_count,omptempty"`
	SddHasUsableMandate                   *Bool       `xmlrpc:"sdd_has_usable_mandate,omptempty"`
	SddMandateId                          *Many2One   `xmlrpc:"sdd_mandate_id,omptempty"`
	SddMandateScheme                      *Selection  `xmlrpc:"sdd_mandate_scheme,omptempty"`
	SddMandateUsable                      *Bool       `xmlrpc:"sdd_mandate_usable,omptempty"`
	SecureSequenceNumber                  *Int        `xmlrpc:"secure_sequence_number,omptempty"`
	SendAndPrintValues                    interface{} `xmlrpc:"send_and_print_values,omptempty"`
	SequenceNumber                        *Int        `xmlrpc:"sequence_number,omptempty"`
	SequencePrefix                        *String     `xmlrpc:"sequence_prefix,omptempty"`
	ShowDeliveryDate                      *Bool       `xmlrpc:"show_delivery_date,omptempty"`
	ShowDiscountDetails                   *Bool       `xmlrpc:"show_discount_details,omptempty"`
	ShowNameWarning                       *Bool       `xmlrpc:"show_name_warning,omptempty"`
	ShowPartnerBankAccount                *Bool       `xmlrpc:"show_partner_bank_account,omptempty"`
	ShowPaymentTermDetails                *Bool       `xmlrpc:"show_payment_term_details,omptempty"`
	ShowResetToDraftButton                *Bool       `xmlrpc:"show_reset_to_draft_button,omptempty"`
	SourceId                              *Many2One   `xmlrpc:"source_id,omptempty"`
	SourcePaymentId                       *Many2One   `xmlrpc:"source_payment_id,omptempty"`
	State                                 *Selection  `xmlrpc:"state,omptempty"`
	StatementId                           *Many2One   `xmlrpc:"statement_id,omptempty"`
	StatementLineId                       *Many2One   `xmlrpc:"statement_line_id,omptempty"`
	StatementLineIds                      *Relation   `xmlrpc:"statement_line_ids,omptempty"`
	StockMoveId                           *Many2One   `xmlrpc:"stock_move_id,omptempty"`
	StockValuationLayerIds                *Relation   `xmlrpc:"stock_valuation_layer_ids,omptempty"`
	StringToHash                          *String     `xmlrpc:"string_to_hash,omptempty"`
	SuitableJournalIds                    *Relation   `xmlrpc:"suitable_journal_ids,omptempty"`
	SuitablePaymentTokenIds               *Relation   `xmlrpc:"suitable_payment_token_ids,omptempty"`
	TaxCalculationRoundingMethod          *Selection  `xmlrpc:"tax_calculation_rounding_method,omptempty"`
	TaxCashBasisCreatedMoveIds            *Relation   `xmlrpc:"tax_cash_basis_created_move_ids,omptempty"`
	TaxCashBasisOriginMoveId              *Many2One   `xmlrpc:"tax_cash_basis_origin_move_id,omptempty"`
	TaxCashBasisRecId                     *Many2One   `xmlrpc:"tax_cash_basis_rec_id,omptempty"`
	TaxCountryCode                        *String     `xmlrpc:"tax_country_code,omptempty"`
	TaxCountryId                          *Many2One   `xmlrpc:"tax_country_id,omptempty"`
	TaxLockDateMessage                    *String     `xmlrpc:"tax_lock_date_message,omptempty"`
	TaxTotals                             *String     `xmlrpc:"tax_totals,omptempty"`
	TeamId                                *Many2One   `xmlrpc:"team_id,omptempty"`
	ToCheck                               *Bool       `xmlrpc:"to_check,omptempty"`
	TransactionIds                        *Relation   `xmlrpc:"transaction_ids,omptempty"`
	TypeName                              *String     `xmlrpc:"type_name,omptempty"`
	UblCiiXmlFile                         *String     `xmlrpc:"ubl_cii_xml_file,omptempty"`
	UblCiiXmlId                           *Many2One   `xmlrpc:"ubl_cii_xml_id,omptempty"`
	UseElectronicPaymentMethod            *Bool       `xmlrpc:"use_electronic_payment_method,omptempty"`
	UserId                                *Many2One   `xmlrpc:"user_id,omptempty"`
	WebsiteId                             *Many2One   `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds                     *Relation   `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                             *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                              *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// AccountPayments represents array of account.payment model.
type AccountPayments []AccountPayment

// AccountPaymentModel is the odoo model name.
const AccountPaymentModel = "account.payment"

// Many2One convert AccountPayment to *Many2One.
func (ap *AccountPayment) Many2One() *Many2One {
	return NewMany2One(ap.Id.Get(), "")
}

// CreateAccountPayment creates a new account.payment model and returns its id.
func (c *Client) CreateAccountPayment(ap *AccountPayment) (int64, error) {
	ids, err := c.CreateAccountPayments([]*AccountPayment{ap})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPayment creates a new account.payment model and returns its id.
func (c *Client) CreateAccountPayments(aps []*AccountPayment) ([]int64, error) {
	var vv []interface{}
	for _, v := range aps {
		vv = append(vv, v)
	}
	return c.Create(AccountPaymentModel, vv)
}

// UpdateAccountPayment updates an existing account.payment record.
func (c *Client) UpdateAccountPayment(ap *AccountPayment) error {
	return c.UpdateAccountPayments([]int64{ap.Id.Get()}, ap)
}

// UpdateAccountPayments updates existing account.payment records.
// All records (represented by ids) will be updated by ap values.
func (c *Client) UpdateAccountPayments(ids []int64, ap *AccountPayment) error {
	return c.Update(AccountPaymentModel, ids, ap)
}

// DeleteAccountPayment deletes an existing account.payment record.
func (c *Client) DeleteAccountPayment(id int64) error {
	return c.DeleteAccountPayments([]int64{id})
}

// DeleteAccountPayments deletes existing account.payment records.
func (c *Client) DeleteAccountPayments(ids []int64) error {
	return c.Delete(AccountPaymentModel, ids)
}

// GetAccountPayment gets account.payment existing record.
func (c *Client) GetAccountPayment(id int64) (*AccountPayment, error) {
	aps, err := c.GetAccountPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	if aps != nil && len(*aps) > 0 {
		return &((*aps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.payment not found", id)
}

// GetAccountPayments gets account.payment existing records.
func (c *Client) GetAccountPayments(ids []int64) (*AccountPayments, error) {
	aps := &AccountPayments{}
	if err := c.Read(AccountPaymentModel, ids, nil, aps); err != nil {
		return nil, err
	}
	return aps, nil
}

// FindAccountPayment finds account.payment record by querying it with criteria.
func (c *Client) FindAccountPayment(criteria *Criteria) (*AccountPayment, error) {
	aps := &AccountPayments{}
	if err := c.SearchRead(AccountPaymentModel, criteria, NewOptions().Limit(1), aps); err != nil {
		return nil, err
	}
	if aps != nil && len(*aps) > 0 {
		return &((*aps)[0]), nil
	}
	return nil, fmt.Errorf("account.payment was not found with criteria %v", criteria)
}

// FindAccountPayments finds account.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPayments(criteria *Criteria, options *Options) (*AccountPayments, error) {
	aps := &AccountPayments{}
	if err := c.SearchRead(AccountPaymentModel, criteria, options, aps); err != nil {
		return nil, err
	}
	return aps, nil
}

// FindAccountPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountPaymentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountPaymentId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.payment was not found with criteria %v and options %v", criteria, options)
}
