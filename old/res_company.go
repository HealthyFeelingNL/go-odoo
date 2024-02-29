package odoo

import (
	"fmt"
)

// ResCompany represents res.company model.
type ResCompany struct {
	AccountCashBasisBaseAccountId               *Many2One  `xmlrpc:"account_cash_basis_base_account_id,omptempty"`
	AccountDefaultPosReceivableAccountId        *Many2One  `xmlrpc:"account_default_pos_receivable_account_id,omptempty"`
	AccountDiscountExpenseAllocationId          *Many2One  `xmlrpc:"account_discount_expense_allocation_id,omptempty"`
	AccountDiscountIncomeAllocationId           *Many2One  `xmlrpc:"account_discount_income_allocation_id,omptempty"`
	AccountEnabledTaxCountryIds                 *Relation  `xmlrpc:"account_enabled_tax_country_ids,omptempty"`
	AccountFiscalCountryId                      *Many2One  `xmlrpc:"account_fiscal_country_id,omptempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omptempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omptempty"`
	AccountJournalPaymentCreditAccountId        *Many2One  `xmlrpc:"account_journal_payment_credit_account_id,omptempty"`
	AccountJournalPaymentDebitAccountId         *Many2One  `xmlrpc:"account_journal_payment_debit_account_id,omptempty"`
	AccountJournalSuspenseAccountId             *Many2One  `xmlrpc:"account_journal_suspense_account_id,omptempty"`
	AccountOpeningDate                          *Time      `xmlrpc:"account_opening_date,omptempty"`
	AccountOpeningJournalId                     *Many2One  `xmlrpc:"account_opening_journal_id,omptempty"`
	AccountOpeningMoveId                        *Many2One  `xmlrpc:"account_opening_move_id,omptempty"`
	AccountPurchaseTaxId                        *Many2One  `xmlrpc:"account_purchase_tax_id,omptempty"`
	AccountSaleTaxId                            *Many2One  `xmlrpc:"account_sale_tax_id,omptempty"`
	AccountStorno                               *Bool      `xmlrpc:"account_storno,omptempty"`
	AccountUseCreditLimit                       *Bool      `xmlrpc:"account_use_credit_limit,omptempty"`
	Active                                      *Bool      `xmlrpc:"active,omptempty"`
	AliasDomainId                               *Many2One  `xmlrpc:"alias_domain_id,omptempty"`
	AliasDomainName                             *String    `xmlrpc:"alias_domain_name,omptempty"`
	AllChildIds                                 *Relation  `xmlrpc:"all_child_ids,omptempty"`
	AngloSaxonAccounting                        *Bool      `xmlrpc:"anglo_saxon_accounting,omptempty"`
	AnnualInventoryDay                          *Int       `xmlrpc:"annual_inventory_day,omptempty"`
	AnnualInventoryMonth                        *Selection `xmlrpc:"annual_inventory_month,omptempty"`
	AutomaticEntryDefaultJournalId              *Many2One  `xmlrpc:"automatic_entry_default_journal_id,omptempty"`
	BackgroundImage                             *String    `xmlrpc:"background_image,omptempty"`
	BankAccountCodePrefix                       *String    `xmlrpc:"bank_account_code_prefix,omptempty"`
	BankIds                                     *Relation  `xmlrpc:"bank_ids,omptempty"`
	BankJournalIds                              *Relation  `xmlrpc:"bank_journal_ids,omptempty"`
	BounceEmail                                 *String    `xmlrpc:"bounce_email,omptempty"`
	BounceFormatted                             *String    `xmlrpc:"bounce_formatted,omptempty"`
	CashAccountCodePrefix                       *String    `xmlrpc:"cash_account_code_prefix,omptempty"`
	CatchallEmail                               *String    `xmlrpc:"catchall_email,omptempty"`
	CatchallFormatted                           *String    `xmlrpc:"catchall_formatted,omptempty"`
	ChartTemplate                               *Selection `xmlrpc:"chart_template,omptempty"`
	ChildIds                                    *Relation  `xmlrpc:"child_ids,omptempty"`
	City                                        *String    `xmlrpc:"city,omptempty"`
	Color                                       *Int       `xmlrpc:"color,omptempty"`
	CompanyDetails                              *String    `xmlrpc:"company_details,omptempty"`
	CompanyRegistry                             *String    `xmlrpc:"company_registry,omptempty"`
	CountryCode                                 *String    `xmlrpc:"country_code,omptempty"`
	CountryId                                   *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyExchangeJournalId                   *Many2One  `xmlrpc:"currency_exchange_journal_id,omptempty"`
	CurrencyId                                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrencyIntervalUnit                        *Selection `xmlrpc:"currency_interval_unit,omptempty"`
	CurrencyNextExecutionDate                   *Time      `xmlrpc:"currency_next_execution_date,omptempty"`
	CurrencyProvider                            *Selection `xmlrpc:"currency_provider,omptempty"`
	DefaultCashDifferenceExpenseAccountId       *Many2One  `xmlrpc:"default_cash_difference_expense_account_id,omptempty"`
	DefaultCashDifferenceIncomeAccountId        *Many2One  `xmlrpc:"default_cash_difference_income_account_id,omptempty"`
	DefaultFromEmail                            *String    `xmlrpc:"default_from_email,omptempty"`
	DisplayInvoiceAmountTotalWords              *Bool      `xmlrpc:"display_invoice_amount_total_words,omptempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omptempty"`
	Email                                       *String    `xmlrpc:"email,omptempty"`
	EmailFormatted                              *String    `xmlrpc:"email_formatted,omptempty"`
	EmailPrimaryColor                           *String    `xmlrpc:"email_primary_color,omptempty"`
	EmailSecondaryColor                         *String    `xmlrpc:"email_secondary_color,omptempty"`
	ExpectsChartOfAccounts                      *Bool      `xmlrpc:"expects_chart_of_accounts,omptempty"`
	ExpenseAccrualAccountId                     *Many2One  `xmlrpc:"expense_accrual_account_id,omptempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omptempty"`
	ExternalReportLayoutId                      *Many2One  `xmlrpc:"external_report_layout_id,omptempty"`
	ExtractInInvoiceDigitalizationMode          *Selection `xmlrpc:"extract_in_invoice_digitalization_mode,omptempty"`
	ExtractOutInvoiceDigitalizationMode         *Selection `xmlrpc:"extract_out_invoice_digitalization_mode,omptempty"`
	ExtractSingleLinePerTax                     *Bool      `xmlrpc:"extract_single_line_per_tax,omptempty"`
	FiscalPositionIds                           *Relation  `xmlrpc:"fiscal_position_ids,omptempty"`
	FiscalyearLastDay                           *Int       `xmlrpc:"fiscalyear_last_day,omptempty"`
	FiscalyearLastMonth                         *Selection `xmlrpc:"fiscalyear_last_month,omptempty"`
	FiscalyearLockDate                          *Time      `xmlrpc:"fiscalyear_lock_date,omptempty"`
	Font                                        *Selection `xmlrpc:"font,omptempty"`
	HasMessage                                  *Bool      `xmlrpc:"has_message,omptempty"`
	HasReceivedWarningStockSms                  *Bool      `xmlrpc:"has_received_warning_stock_sms,omptempty"`
	IapEnrichAutoDone                           *Bool      `xmlrpc:"iap_enrich_auto_done,omptempty"`
	Id                                          *Int       `xmlrpc:"id,omptempty"`
	IncomeCurrencyExchangeAccountId             *Many2One  `xmlrpc:"income_currency_exchange_account_id,omptempty"`
	IncotermId                                  *Many2One  `xmlrpc:"incoterm_id,omptempty"`
	InternalTransitLocationId                   *Many2One  `xmlrpc:"internal_transit_location_id,omptempty"`
	InvoiceIsDownload                           *Bool      `xmlrpc:"invoice_is_download,omptempty"`
	InvoiceIsEmail                              *Bool      `xmlrpc:"invoice_is_email,omptempty"`
	InvoiceIsSnailmail                          *Bool      `xmlrpc:"invoice_is_snailmail,omptempty"`
	InvoiceIsUblCii                             *Bool      `xmlrpc:"invoice_is_ubl_cii,omptempty"`
	InvoiceTerms                                *String    `xmlrpc:"invoice_terms,omptempty"`
	InvoiceTermsHtml                            *String    `xmlrpc:"invoice_terms_html,omptempty"`
	IsCompanyDetailsEmpty                       *Bool      `xmlrpc:"is_company_details_empty,omptempty"`
	LayoutBackground                            *Selection `xmlrpc:"layout_background,omptempty"`
	LayoutBackgroundImage                       *String    `xmlrpc:"layout_background_image,omptempty"`
	Logo                                        *String    `xmlrpc:"logo,omptempty"`
	LogoWeb                                     *String    `xmlrpc:"logo_web,omptempty"`
	MaxTaxLockDate                              *Time      `xmlrpc:"max_tax_lock_date,omptempty"`
	MessageAttachmentCount                      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds                          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter                      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageNeedaction                           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter                    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Mobile                                      *String    `xmlrpc:"mobile,omptempty"`
	MultiVatForeignCountryIds                   *Relation  `xmlrpc:"multi_vat_foreign_country_ids,omptempty"`
	Name                                        *String    `xmlrpc:"name,omptempty"`
	NomenclatureId                              *Many2One  `xmlrpc:"nomenclature_id,omptempty"`
	PaperformatId                               *Many2One  `xmlrpc:"paperformat_id,omptempty"`
	ParentId                                    *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentIds                                   *Relation  `xmlrpc:"parent_ids,omptempty"`
	ParentPath                                  *String    `xmlrpc:"parent_path,omptempty"`
	PartnerGid                                  *Int       `xmlrpc:"partner_gid,omptempty"`
	PartnerId                                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	PaymentOnboardingPaymentMethod              *Selection `xmlrpc:"payment_onboarding_payment_method,omptempty"`
	PeriodLockDate                              *Time      `xmlrpc:"period_lock_date,omptempty"`
	Phone                                       *String    `xmlrpc:"phone,omptempty"`
	PortalConfirmationPay                       *Bool      `xmlrpc:"portal_confirmation_pay,omptempty"`
	PortalConfirmationSign                      *Bool      `xmlrpc:"portal_confirmation_sign,omptempty"`
	PrepaymentPercent                           *Float     `xmlrpc:"prepayment_percent,omptempty"`
	PrimaryColor                                *String    `xmlrpc:"primary_color,omptempty"`
	QrCode                                      *Bool      `xmlrpc:"qr_code,omptempty"`
	QuickEditMode                               *Selection `xmlrpc:"quick_edit_mode,omptempty"`
	QuotationValidityDays                       *Int       `xmlrpc:"quotation_validity_days,omptempty"`
	RatingIds                                   *Relation  `xmlrpc:"rating_ids,omptempty"`
	ReportFooter                                *String    `xmlrpc:"report_footer,omptempty"`
	ReportHeader                                *String    `xmlrpc:"report_header,omptempty"`
	ResourceCalendarId                          *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceCalendarIds                         *Relation  `xmlrpc:"resource_calendar_ids,omptempty"`
	RevenueAccrualAccountId                     *Many2One  `xmlrpc:"revenue_accrual_account_id,omptempty"`
	RootId                                      *Many2One  `xmlrpc:"root_id,omptempty"`
	SaleDiscountProductId                       *Many2One  `xmlrpc:"sale_discount_product_id,omptempty"`
	SaleDownPaymentProductId                    *Many2One  `xmlrpc:"sale_down_payment_product_id,omptempty"`
	SaleFooter                                  *String    `xmlrpc:"sale_footer,omptempty"`
	SaleFooterName                              *String    `xmlrpc:"sale_footer_name,omptempty"`
	SaleHeader                                  *String    `xmlrpc:"sale_header,omptempty"`
	SaleHeaderName                              *String    `xmlrpc:"sale_header_name,omptempty"`
	SaleOnboardingPaymentMethod                 *Selection `xmlrpc:"sale_onboarding_payment_method,omptempty"`
	SaleOrderTemplateId                         *Many2One  `xmlrpc:"sale_order_template_id,omptempty"`
	SddCreditorIdentifier                       *String    `xmlrpc:"sdd_creditor_identifier,omptempty"`
	SecondaryColor                              *String    `xmlrpc:"secondary_color,omptempty"`
	SecurityLead                                *Float     `xmlrpc:"security_lead,omptempty"`
	Sequence                                    *Int       `xmlrpc:"sequence,omptempty"`
	SnailmailColor                              *Bool      `xmlrpc:"snailmail_color,omptempty"`
	SnailmailCover                              *Bool      `xmlrpc:"snailmail_cover,omptempty"`
	SnailmailDuplex                             *Bool      `xmlrpc:"snailmail_duplex,omptempty"`
	SocialFacebook                              *String    `xmlrpc:"social_facebook,omptempty"`
	SocialGithub                                *String    `xmlrpc:"social_github,omptempty"`
	SocialInstagram                             *String    `xmlrpc:"social_instagram,omptempty"`
	SocialLinkedin                              *String    `xmlrpc:"social_linkedin,omptempty"`
	SocialTiktok                                *String    `xmlrpc:"social_tiktok,omptempty"`
	SocialTwitter                               *String    `xmlrpc:"social_twitter,omptempty"`
	SocialYoutube                               *String    `xmlrpc:"social_youtube,omptempty"`
	StateId                                     *Many2One  `xmlrpc:"state_id,omptempty"`
	StockMailConfirmationTemplateId             *Many2One  `xmlrpc:"stock_mail_confirmation_template_id,omptempty"`
	StockMoveEmailValidation                    *Bool      `xmlrpc:"stock_move_email_validation,omptempty"`
	StockMoveSmsValidation                      *Bool      `xmlrpc:"stock_move_sms_validation,omptempty"`
	StockSmsConfirmationTemplateId              *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omptempty"`
	Street                                      *String    `xmlrpc:"street,omptempty"`
	Street2                                     *String    `xmlrpc:"street2,omptempty"`
	TaxCalculationRoundingMethod                *Selection `xmlrpc:"tax_calculation_rounding_method,omptempty"`
	TaxCashBasisJournalId                       *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omptempty"`
	TaxExigibility                              *Bool      `xmlrpc:"tax_exigibility,omptempty"`
	TaxLockDate                                 *Time      `xmlrpc:"tax_lock_date,omptempty"`
	TermsType                                   *Selection `xmlrpc:"terms_type,omptempty"`
	TransferAccountCodePrefix                   *String    `xmlrpc:"transfer_account_code_prefix,omptempty"`
	TransferAccountId                           *Many2One  `xmlrpc:"transfer_account_id,omptempty"`
	UserIds                                     *Relation  `xmlrpc:"user_ids,omptempty"`
	UsesDefaultLogo                             *Bool      `xmlrpc:"uses_default_logo,omptempty"`
	Vat                                         *String    `xmlrpc:"vat,omptempty"`
	Website                                     *String    `xmlrpc:"website,omptempty"`
	WebsiteId                                   *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds                           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omptempty"`
	Zip                                         *String    `xmlrpc:"zip,omptempty"`
}

// ResCompanys represents array of res.company model.
type ResCompanys []ResCompany

// ResCompanyModel is the odoo model name.
const ResCompanyModel = "res.company"

// Many2One convert ResCompany to *Many2One.
func (rc *ResCompany) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompany(rc *ResCompany) (int64, error) {
	ids, err := c.CreateResCompanys([]*ResCompany{rc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompanys(rcs []*ResCompany) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcs {
		vv = append(vv, v)
	}
	return c.Create(ResCompanyModel, vv)
}

// UpdateResCompany updates an existing res.company record.
func (c *Client) UpdateResCompany(rc *ResCompany) error {
	return c.UpdateResCompanys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCompanys updates existing res.company records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCompanys(ids []int64, rc *ResCompany) error {
	return c.Update(ResCompanyModel, ids, rc)
}

// DeleteResCompany deletes an existing res.company record.
func (c *Client) DeleteResCompany(id int64) error {
	return c.DeleteResCompanys([]int64{id})
}

// DeleteResCompanys deletes existing res.company records.
func (c *Client) DeleteResCompanys(ids []int64) error {
	return c.Delete(ResCompanyModel, ids)
}

// GetResCompany gets res.company existing record.
func (c *Client) GetResCompany(id int64) (*ResCompany, error) {
	rcs, err := c.GetResCompanys([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.company not found", id)
}

// GetResCompanys gets res.company existing records.
func (c *Client) GetResCompanys(ids []int64) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.Read(ResCompanyModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompany finds res.company record by querying it with criteria.
func (c *Client) FindResCompany(criteria *Criteria) (*ResCompany, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("res.company was not found with criteria %v", criteria)
}

// FindResCompanys finds res.company records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanys(criteria *Criteria, options *Options) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompanyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResCompanyId finds record id by querying it with criteria.
func (c *Client) FindResCompanyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.company was not found with criteria %v and options %v", criteria, options)
}
