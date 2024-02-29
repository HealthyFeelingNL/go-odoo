package odoo

import (
	"fmt"
)

// ResUsers represents res.users model.
type ResUsers struct {
	AccessesCount                      *Int       `xmlrpc:"accesses_count,omptempty"`
	ActionId                           *Many2One  `xmlrpc:"action_id,omptempty"`
	Active                             *Bool      `xmlrpc:"active,omptempty"`
	ActiveLangCount                    *Int       `xmlrpc:"active_lang_count,omptempty"`
	ActivePartner                      *Bool      `xmlrpc:"active_partner,omptempty"`
	ActivityCalendarEventId            *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline               *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration        *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon              *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                        *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                      *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                    *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                   *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                     *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                     *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AdditionalInfo                     *String    `xmlrpc:"additional_info,omptempty"`
	ApiKeyIds                          *Relation  `xmlrpc:"api_key_ids,omptempty"`
	Avatar1024                         *String    `xmlrpc:"avatar_1024,omptempty"`
	Avatar128                          *String    `xmlrpc:"avatar_128,omptempty"`
	Avatar1920                         *String    `xmlrpc:"avatar_1920,omptempty"`
	Avatar256                          *String    `xmlrpc:"avatar_256,omptempty"`
	Avatar512                          *String    `xmlrpc:"avatar_512,omptempty"`
	BankAccountCount                   *Int       `xmlrpc:"bank_account_count,omptempty"`
	BankIds                            *Relation  `xmlrpc:"bank_ids,omptempty"`
	Barcode                            *String    `xmlrpc:"barcode,omptempty"`
	CalendarLastNotifAck               *Time      `xmlrpc:"calendar_last_notif_ack,omptempty"`
	CanPublish                         *Bool      `xmlrpc:"can_publish,omptempty"`
	CategoryId                         *Relation  `xmlrpc:"category_id,omptempty"`
	ChannelIds                         *Relation  `xmlrpc:"channel_ids,omptempty"`
	ChildIds                           *Relation  `xmlrpc:"child_ids,omptempty"`
	City                               *String    `xmlrpc:"city,omptempty"`
	Color                              *Int       `xmlrpc:"color,omptempty"`
	Comment                            *String    `xmlrpc:"comment,omptempty"`
	CommercialCompanyName              *String    `xmlrpc:"commercial_company_name,omptempty"`
	CommercialPartnerId                *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompaniesCount                     *Int       `xmlrpc:"companies_count,omptempty"`
	CompanyId                          *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyIds                         *Relation  `xmlrpc:"company_ids,omptempty"`
	CompanyName                        *String    `xmlrpc:"company_name,omptempty"`
	CompanyRegistry                    *String    `xmlrpc:"company_registry,omptempty"`
	CompanyType                        *Selection `xmlrpc:"company_type,omptempty"`
	CompleteName                       *String    `xmlrpc:"complete_name,omptempty"`
	ContactAddress                     *String    `xmlrpc:"contact_address,omptempty"`
	ContactAddressComplete             *String    `xmlrpc:"contact_address_complete,omptempty"`
	ContactAddressInline               *String    `xmlrpc:"contact_address_inline,omptempty"`
	ContractIds                        *Relation  `xmlrpc:"contract_ids,omptempty"`
	CountryCode                        *String    `xmlrpc:"country_code,omptempty"`
	CountryId                          *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                          *Many2One  `xmlrpc:"create_uid,omptempty"`
	Credit                             *Float     `xmlrpc:"credit,omptempty"`
	CreditLimit                        *Float     `xmlrpc:"credit_limit,omptempty"`
	CreditToInvoice                    *Float     `xmlrpc:"credit_to_invoice,omptempty"`
	CrmTeamIds                         *Relation  `xmlrpc:"crm_team_ids,omptempty"`
	CrmTeamMemberIds                   *Relation  `xmlrpc:"crm_team_member_ids,omptempty"`
	CurrencyId                         *Many2One  `xmlrpc:"currency_id,omptempty"`
	CustomerRank                       *Int       `xmlrpc:"customer_rank,omptempty"`
	Date                               *Time      `xmlrpc:"date,omptempty"`
	DaysSalesOutstanding               *Float     `xmlrpc:"days_sales_outstanding,omptempty"`
	Debit                              *Float     `xmlrpc:"debit,omptempty"`
	DebitLimit                         *Float     `xmlrpc:"debit_limit,omptempty"`
	DisplayName                        *String    `xmlrpc:"display_name,omptempty"`
	DuplicatedBankAccountPartnersCount *Int       `xmlrpc:"duplicated_bank_account_partners_count,omptempty"`
	Email                              *String    `xmlrpc:"email,omptempty"`
	EmailFormatted                     *String    `xmlrpc:"email_formatted,omptempty"`
	EmailNormalized                    *String    `xmlrpc:"email_normalized,omptempty"`
	Employee                           *Bool      `xmlrpc:"employee,omptempty"`
	FiscalCountryCodes                 *String    `xmlrpc:"fiscal_country_codes,omptempty"`
	Function                           *String    `xmlrpc:"function,omptempty"`
	GroupsCount                        *Int       `xmlrpc:"groups_count,omptempty"`
	GroupsId                           *Relation  `xmlrpc:"groups_id,omptempty"`
	HasAccessLivechat                  *Bool      `xmlrpc:"has_access_livechat,omptempty"`
	HasMessage                         *Bool      `xmlrpc:"has_message,omptempty"`
	HasUnreconciledEntries             *Bool      `xmlrpc:"has_unreconciled_entries,omptempty"`
	HelpdeskTargetClosed               *Int       `xmlrpc:"helpdesk_target_closed,omptempty"`
	HelpdeskTargetRating               *Float     `xmlrpc:"helpdesk_target_rating,omptempty"`
	HelpdeskTargetSuccess              *Float     `xmlrpc:"helpdesk_target_success,omptempty"`
	Id                                 *Int       `xmlrpc:"id,omptempty"`
	ImStatus                           *String    `xmlrpc:"im_status,omptempty"`
	Image1024                          *String    `xmlrpc:"image_1024,omptempty"`
	Image128                           *String    `xmlrpc:"image_128,omptempty"`
	Image1920                          *String    `xmlrpc:"image_1920,omptempty"`
	Image256                           *String    `xmlrpc:"image_256,omptempty"`
	Image512                           *String    `xmlrpc:"image_512,omptempty"`
	ImageMedium                        *String    `xmlrpc:"image_medium,omptempty"`
	IndustryId                         *Many2One  `xmlrpc:"industry_id,omptempty"`
	InvoiceIds                         *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoiceWarn                        *Selection `xmlrpc:"invoice_warn,omptempty"`
	InvoiceWarnMsg                     *String    `xmlrpc:"invoice_warn_msg,omptempty"`
	IsBlacklisted                      *Bool      `xmlrpc:"is_blacklisted,omptempty"`
	IsCompany                          *Bool      `xmlrpc:"is_company,omptempty"`
	IsPublic                           *Bool      `xmlrpc:"is_public,omptempty"`
	IsPublished                        *Bool      `xmlrpc:"is_published,omptempty"`
	JournalItemCount                   *Int       `xmlrpc:"journal_item_count,omptempty"`
	Lang                               *Selection `xmlrpc:"lang,omptempty"`
	LastTimeEntriesChecked             *Time      `xmlrpc:"last_time_entries_checked,omptempty"`
	LastWebsiteSoId                    *Many2One  `xmlrpc:"last_website_so_id,omptempty"`
	LivechatLangIds                    *Relation  `xmlrpc:"livechat_lang_ids,omptempty"`
	LivechatUsername                   *String    `xmlrpc:"livechat_username,omptempty"`
	LogIds                             *Relation  `xmlrpc:"log_ids,omptempty"`
	Login                              *String    `xmlrpc:"login,omptempty"`
	LoginDate                          *Time      `xmlrpc:"login_date,omptempty"`
	LoyaltyCardCount                   *Int       `xmlrpc:"loyalty_card_count,omptempty"`
	MeetingCount                       *Int       `xmlrpc:"meeting_count,omptempty"`
	MeetingIds                         *Relation  `xmlrpc:"meeting_ids,omptempty"`
	MessageAttachmentCount             *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageBounce                      *Int       `xmlrpc:"message_bounce,omptempty"`
	MessageFollowerIds                 *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                    *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter             *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                 *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                         *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                  *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageNeedaction                  *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter           *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                  *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Mobile                             *String    `xmlrpc:"mobile,omptempty"`
	MobileBlacklisted                  *Bool      `xmlrpc:"mobile_blacklisted,omptempty"`
	MyActivityDateDeadline             *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                               *String    `xmlrpc:"name,omptempty"`
	NewPassword                        *String    `xmlrpc:"new_password,omptempty"`
	NotificationType                   *Selection `xmlrpc:"notification_type,omptempty"`
	OcnToken                           *String    `xmlrpc:"ocn_token,omptempty"`
	OdoobotFailed                      *Bool      `xmlrpc:"odoobot_failed,omptempty"`
	OdoobotState                       *Selection `xmlrpc:"odoobot_state,omptempty"`
	OpportunityCount                   *Int       `xmlrpc:"opportunity_count,omptempty"`
	OpportunityIds                     *Relation  `xmlrpc:"opportunity_ids,omptempty"`
	ParentId                           *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentName                         *String    `xmlrpc:"parent_name,omptempty"`
	PartnerGid                         *Int       `xmlrpc:"partner_gid,omptempty"`
	PartnerId                          *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerLatitude                    *Float     `xmlrpc:"partner_latitude,omptempty"`
	PartnerLongitude                   *Float     `xmlrpc:"partner_longitude,omptempty"`
	PartnerShare                       *Bool      `xmlrpc:"partner_share,omptempty"`
	Password                           *String    `xmlrpc:"password,omptempty"`
	PaymentTokenCount                  *Int       `xmlrpc:"payment_token_count,omptempty"`
	PaymentTokenIds                    *Relation  `xmlrpc:"payment_token_ids,omptempty"`
	PeppolEas                          *Selection `xmlrpc:"peppol_eas,omptempty"`
	PeppolEndpoint                     *String    `xmlrpc:"peppol_endpoint,omptempty"`
	Phone                              *String    `xmlrpc:"phone,omptempty"`
	PhoneBlacklisted                   *Bool      `xmlrpc:"phone_blacklisted,omptempty"`
	PhoneMobileSearch                  *String    `xmlrpc:"phone_mobile_search,omptempty"`
	PhoneSanitized                     *String    `xmlrpc:"phone_sanitized,omptempty"`
	PhoneSanitizedBlacklisted          *Bool      `xmlrpc:"phone_sanitized_blacklisted,omptempty"`
	PickingWarn                        *Selection `xmlrpc:"picking_warn,omptempty"`
	PickingWarnMsg                     *String    `xmlrpc:"picking_warn_msg,omptempty"`
	PropertyAccountPayableId           *Many2One  `xmlrpc:"property_account_payable_id,omptempty"`
	PropertyAccountPositionId          *Many2One  `xmlrpc:"property_account_position_id,omptempty"`
	PropertyAccountReceivableId        *Many2One  `xmlrpc:"property_account_receivable_id,omptempty"`
	PropertyDeliveryCarrierId          *Many2One  `xmlrpc:"property_delivery_carrier_id,omptempty"`
	PropertyPaymentTermId              *Many2One  `xmlrpc:"property_payment_term_id,omptempty"`
	PropertyProductPricelist           *Many2One  `xmlrpc:"property_product_pricelist,omptempty"`
	PropertyStockCustomer              *Many2One  `xmlrpc:"property_stock_customer,omptempty"`
	PropertyStockSupplier              *Many2One  `xmlrpc:"property_stock_supplier,omptempty"`
	PropertySupplierPaymentTermId      *Many2One  `xmlrpc:"property_supplier_payment_term_id,omptempty"`
	PropertyWarehouseId                *Many2One  `xmlrpc:"property_warehouse_id,omptempty"`
	RatingIds                          *Relation  `xmlrpc:"rating_ids,omptempty"`
	Ref                                *String    `xmlrpc:"ref,omptempty"`
	RefCompanyIds                      *Relation  `xmlrpc:"ref_company_ids,omptempty"`
	ResUsersSettingsId                 *Many2One  `xmlrpc:"res_users_settings_id,omptempty"`
	ResUsersSettingsIds                *Relation  `xmlrpc:"res_users_settings_ids,omptempty"`
	ResourceCalendarId                 *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceIds                        *Relation  `xmlrpc:"resource_ids,omptempty"`
	RulesCount                         *Int       `xmlrpc:"rules_count,omptempty"`
	SaleOrderCount                     *Int       `xmlrpc:"sale_order_count,omptempty"`
	SaleOrderIds                       *Relation  `xmlrpc:"sale_order_ids,omptempty"`
	SaleTeamId                         *Many2One  `xmlrpc:"sale_team_id,omptempty"`
	SaleWarn                           *Selection `xmlrpc:"sale_warn,omptempty"`
	SaleWarnMsg                        *String    `xmlrpc:"sale_warn_msg,omptempty"`
	SameCompanyRegistryPartnerId       *Many2One  `xmlrpc:"same_company_registry_partner_id,omptempty"`
	SameVatPartnerId                   *Many2One  `xmlrpc:"same_vat_partner_id,omptempty"`
	SddCount                           *Int       `xmlrpc:"sdd_count,omptempty"`
	SddMandateIds                      *Relation  `xmlrpc:"sdd_mandate_ids,omptempty"`
	Self                               *Many2One  `xmlrpc:"self,omptempty"`
	Share                              *Bool      `xmlrpc:"share,omptempty"`
	ShowCreditLimit                    *Bool      `xmlrpc:"show_credit_limit,omptempty"`
	Signature                          *String    `xmlrpc:"signature,omptempty"`
	SignupExpiration                   *Time      `xmlrpc:"signup_expiration,omptempty"`
	SignupToken                        *String    `xmlrpc:"signup_token,omptempty"`
	SignupType                         *String    `xmlrpc:"signup_type,omptempty"`
	SignupUrl                          *String    `xmlrpc:"signup_url,omptempty"`
	SignupValid                        *Bool      `xmlrpc:"signup_valid,omptempty"`
	SlaIds                             *Relation  `xmlrpc:"sla_ids,omptempty"`
	StarredMessageIds                  *Relation  `xmlrpc:"starred_message_ids,omptempty"`
	State                              *Selection `xmlrpc:"state,omptempty"`
	StateId                            *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                             *String    `xmlrpc:"street,omptempty"`
	Street2                            *String    `xmlrpc:"street2,omptempty"`
	SupplierRank                       *Int       `xmlrpc:"supplier_rank,omptempty"`
	TargetSalesDone                    *Int       `xmlrpc:"target_sales_done,omptempty"`
	TargetSalesInvoiced                *Int       `xmlrpc:"target_sales_invoiced,omptempty"`
	TargetSalesWon                     *Int       `xmlrpc:"target_sales_won,omptempty"`
	TeamId                             *Many2One  `xmlrpc:"team_id,omptempty"`
	TicketCount                        *Int       `xmlrpc:"ticket_count,omptempty"`
	Title                              *Many2One  `xmlrpc:"title,omptempty"`
	TotalInvoiced                      *Float     `xmlrpc:"total_invoiced,omptempty"`
	TotpEnabled                        *Bool      `xmlrpc:"totp_enabled,omptempty"`
	TotpSecret                         *String    `xmlrpc:"totp_secret,omptempty"`
	TotpTrustedDeviceIds               *Relation  `xmlrpc:"totp_trusted_device_ids,omptempty"`
	Trust                              *Selection `xmlrpc:"trust,omptempty"`
	Type                               *Selection `xmlrpc:"type,omptempty"`
	Tz                                 *Selection `xmlrpc:"tz,omptempty"`
	TzOffset                           *String    `xmlrpc:"tz_offset,omptempty"`
	UblCiiFormat                       *Selection `xmlrpc:"ubl_cii_format,omptempty"`
	UsePartnerCreditLimit              *Bool      `xmlrpc:"use_partner_credit_limit,omptempty"`
	UserGroupWarning                   *String    `xmlrpc:"user_group_warning,omptempty"`
	UserId                             *Many2One  `xmlrpc:"user_id,omptempty"`
	UserIds                            *Relation  `xmlrpc:"user_ids,omptempty"`
	UserLivechatUsername               *String    `xmlrpc:"user_livechat_username,omptempty"`
	Vat                                *String    `xmlrpc:"vat,omptempty"`
	VisitorIds                         *Relation  `xmlrpc:"visitor_ids,omptempty"`
	Website                            *String    `xmlrpc:"website,omptempty"`
	WebsiteId                          *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds                  *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsitePublished                   *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteUrl                         *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                           *Many2One  `xmlrpc:"write_uid,omptempty"`
	XStudioCharField6Qr1Hnpvqtj7       *String    `xmlrpc:"x_studio_char_field_6qr_1hnpvqtj7,omptempty"`
	XStudioUserId                      *String    `xmlrpc:"x_studio_user_id,omptempty"`
	Zip                                *String    `xmlrpc:"zip,omptempty"`
}

// ResUserss represents array of res.users model.
type ResUserss []ResUsers

// ResUsersModel is the odoo model name.
const ResUsersModel = "res.users"

// Many2One convert ResUsers to *Many2One.
func (ru *ResUsers) Many2One() *Many2One {
	return NewMany2One(ru.Id.Get(), "")
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) CreateResUsers(ru *ResUsers) (int64, error) {
	ids, err := c.CreateResUserss([]*ResUsers{ru})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) CreateResUserss(rus []*ResUsers) ([]int64, error) {
	var vv []interface{}
	for _, v := range rus {
		vv = append(vv, v)
	}
	return c.Create(ResUsersModel, vv)
}

// UpdateResUsers updates an existing res.users record.
func (c *Client) UpdateResUsers(ru *ResUsers) error {
	return c.UpdateResUserss([]int64{ru.Id.Get()}, ru)
}

// UpdateResUserss updates existing res.users records.
// All records (represented by ids) will be updated by ru values.
func (c *Client) UpdateResUserss(ids []int64, ru *ResUsers) error {
	return c.Update(ResUsersModel, ids, ru)
}

// DeleteResUsers deletes an existing res.users record.
func (c *Client) DeleteResUsers(id int64) error {
	return c.DeleteResUserss([]int64{id})
}

// DeleteResUserss deletes existing res.users records.
func (c *Client) DeleteResUserss(ids []int64) error {
	return c.Delete(ResUsersModel, ids)
}

// GetResUsers gets res.users existing record.
func (c *Client) GetResUsers(id int64) (*ResUsers, error) {
	rus, err := c.GetResUserss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rus != nil && len(*rus) > 0 {
		return &((*rus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users not found", id)
}

// GetResUserss gets res.users existing records.
func (c *Client) GetResUserss(ids []int64) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.Read(ResUsersModel, ids, nil, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsers finds res.users record by querying it with criteria.
func (c *Client) FindResUsers(criteria *Criteria) (*ResUsers, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, NewOptions().Limit(1), rus); err != nil {
		return nil, err
	}
	if rus != nil && len(*rus) > 0 {
		return &((*rus)[0]), nil
	}
	return nil, fmt.Errorf("res.users was not found with criteria %v", criteria)
}

// FindResUserss finds res.users records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUserss(criteria *Criteria, options *Options) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, options, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersId finds record id by querying it with criteria.
func (c *Client) FindResUsersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users was not found with criteria %v and options %v", criteria, options)
}
