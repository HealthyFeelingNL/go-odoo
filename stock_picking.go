package odoo

import (
	"fmt"
)

// StockPicking represents stock.picking model.
type StockPicking struct {
	ActivityCalendarEventId     *Many2One   `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omptempty"`
	BackorderId                 *Many2One   `xmlrpc:"backorder_id,omptempty"`
	BackorderIds                *Relation   `xmlrpc:"backorder_ids,omptempty"`
	CarrierId                   *Many2One   `xmlrpc:"carrier_id,omptempty"`
	CarrierPrice                *Float      `xmlrpc:"carrier_price,omptempty"`
	CarrierTrackingRef          *String     `xmlrpc:"carrier_tracking_ref,omptempty"`
	CarrierTrackingUrl          *String     `xmlrpc:"carrier_tracking_url,omptempty"`
	CompanyId                   *Many2One   `xmlrpc:"company_id,omptempty"`
	CountryCode                 *String     `xmlrpc:"country_code,omptempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omptempty"`
	Date                        *Time       `xmlrpc:"date,omptempty"`
	DateDeadline                *Time       `xmlrpc:"date_deadline,omptempty"`
	DateDone                    *Time       `xmlrpc:"date_done,omptempty"`
	DelayAlertDate              *Time       `xmlrpc:"delay_alert_date,omptempty"`
	DeliveryType                *Selection  `xmlrpc:"delivery_type,omptempty"`
	DestinationCountryCode      *String     `xmlrpc:"destination_country_code,omptempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omptempty"`
	GroupId                     *Many2One   `xmlrpc:"group_id,omptempty"`
	HasDeadlineIssue            *Bool       `xmlrpc:"has_deadline_issue,omptempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omptempty"`
	HasPackages                 *Bool       `xmlrpc:"has_packages,omptempty"`
	HasScrapMove                *Bool       `xmlrpc:"has_scrap_move,omptempty"`
	HasTracking                 *Bool       `xmlrpc:"has_tracking,omptempty"`
	HidePickingType             *Bool       `xmlrpc:"hide_picking_type,omptempty"`
	Id                          *Int        `xmlrpc:"id,omptempty"`
	IsLocked                    *Bool       `xmlrpc:"is_locked,omptempty"`
	IsReturnPicking             *Bool       `xmlrpc:"is_return_picking,omptempty"`
	IsSigned                    *Bool       `xmlrpc:"is_signed,omptempty"`
	JsonPopover                 *String     `xmlrpc:"json_popover,omptempty"`
	LocationDestId              *Many2One   `xmlrpc:"location_dest_id,omptempty"`
	LocationId                  *Many2One   `xmlrpc:"location_id,omptempty"`
	LotId                       *Many2One   `xmlrpc:"lot_id,omptempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omptempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omptempty"`
	MoveIds                     *Relation   `xmlrpc:"move_ids,omptempty"`
	MoveIdsWithoutPackage       *Relation   `xmlrpc:"move_ids_without_package,omptempty"`
	MoveLineExist               *Bool       `xmlrpc:"move_line_exist,omptempty"`
	MoveLineIds                 *Relation   `xmlrpc:"move_line_ids,omptempty"`
	MoveLineIdsWithoutPackage   *Relation   `xmlrpc:"move_line_ids_without_package,omptempty"`
	MoveType                    *Selection  `xmlrpc:"move_type,omptempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String     `xmlrpc:"name,omptempty"`
	Note                        *String     `xmlrpc:"note,omptempty"`
	Origin                      *String     `xmlrpc:"origin,omptempty"`
	OwnerId                     *Many2One   `xmlrpc:"owner_id,omptempty"`
	PackageIds                  *Relation   `xmlrpc:"package_ids,omptempty"`
	PackageLevelIds             *Relation   `xmlrpc:"package_level_ids,omptempty"`
	PackageLevelIdsDetails      *Relation   `xmlrpc:"package_level_ids_details,omptempty"`
	PartnerId                   *Many2One   `xmlrpc:"partner_id,omptempty"`
	PickingProperties           interface{} `xmlrpc:"picking_properties,omptempty"`
	PickingTypeCode             *Selection  `xmlrpc:"picking_type_code,omptempty"`
	PickingTypeEntirePacks      *Bool       `xmlrpc:"picking_type_entire_packs,omptempty"`
	PickingTypeId               *Many2One   `xmlrpc:"picking_type_id,omptempty"`
	Printed                     *Bool       `xmlrpc:"printed,omptempty"`
	Priority                    *Selection  `xmlrpc:"priority,omptempty"`
	ProductId                   *Many2One   `xmlrpc:"product_id,omptempty"`
	ProductsAvailability        *String     `xmlrpc:"products_availability,omptempty"`
	ProductsAvailabilityState   *Selection  `xmlrpc:"products_availability_state,omptempty"`
	RatingIds                   *Relation   `xmlrpc:"rating_ids,omptempty"`
	ReturnCount                 *Int        `xmlrpc:"return_count,omptempty"`
	ReturnId                    *Many2One   `xmlrpc:"return_id,omptempty"`
	ReturnIds                   *Relation   `xmlrpc:"return_ids,omptempty"`
	ReturnLabelIds              *Relation   `xmlrpc:"return_label_ids,omptempty"`
	SaleId                      *Many2One   `xmlrpc:"sale_id,omptempty"`
	ScheduledDate               *Time       `xmlrpc:"scheduled_date,omptempty"`
	ShippingWeight              *Float      `xmlrpc:"shipping_weight,omptempty"`
	ShowAllocation              *Bool       `xmlrpc:"show_allocation,omptempty"`
	ShowCheckAvailability       *Bool       `xmlrpc:"show_check_availability,omptempty"`
	ShowClearQtyButton          *Bool       `xmlrpc:"show_clear_qty_button,omptempty"`
	ShowLotsText                *Bool       `xmlrpc:"show_lots_text,omptempty"`
	ShowOperations              *Bool       `xmlrpc:"show_operations,omptempty"`
	ShowReserved                *Bool       `xmlrpc:"show_reserved,omptempty"`
	ShowSetQtyButton            *Bool       `xmlrpc:"show_set_qty_button,omptempty"`
	Signature                   *String     `xmlrpc:"signature,omptempty"`
	State                       *Selection  `xmlrpc:"state,omptempty"`
	UseCreateLots               *Bool       `xmlrpc:"use_create_lots,omptempty"`
	UseExistingLots             *Bool       `xmlrpc:"use_existing_lots,omptempty"`
	UserId                      *Many2One   `xmlrpc:"user_id,omptempty"`
	WebsiteId                   *Many2One   `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omptempty"`
	Weight                      *Float      `xmlrpc:"weight,omptempty"`
	WeightBulk                  *Float      `xmlrpc:"weight_bulk,omptempty"`
	WeightUomName               *String     `xmlrpc:"weight_uom_name,omptempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// StockPickings represents array of stock.picking model.
type StockPickings []StockPicking

// StockPickingModel is the odoo model name.
const StockPickingModel = "stock.picking"

// Many2One convert StockPicking to *Many2One.
func (sp *StockPicking) Many2One() *Many2One {
	return NewMany2One(sp.Id.Get(), "")
}

// CreateStockPicking creates a new stock.picking model and returns its id.
func (c *Client) CreateStockPicking(sp *StockPicking) (int64, error) {
	ids, err := c.CreateStockPickings([]*StockPicking{sp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPicking creates a new stock.picking model and returns its id.
func (c *Client) CreateStockPickings(sps []*StockPicking) ([]int64, error) {
	var vv []interface{}
	for _, v := range sps {
		vv = append(vv, v)
	}
	return c.Create(StockPickingModel, vv)
}

// UpdateStockPicking updates an existing stock.picking record.
func (c *Client) UpdateStockPicking(sp *StockPicking) error {
	return c.UpdateStockPickings([]int64{sp.Id.Get()}, sp)
}

// UpdateStockPickings updates existing stock.picking records.
// All records (represented by ids) will be updated by sp values.
func (c *Client) UpdateStockPickings(ids []int64, sp *StockPicking) error {
	return c.Update(StockPickingModel, ids, sp)
}

// DeleteStockPicking deletes an existing stock.picking record.
func (c *Client) DeleteStockPicking(id int64) error {
	return c.DeleteStockPickings([]int64{id})
}

// DeleteStockPickings deletes existing stock.picking records.
func (c *Client) DeleteStockPickings(ids []int64) error {
	return c.Delete(StockPickingModel, ids)
}

// GetStockPicking gets stock.picking existing record.
func (c *Client) GetStockPicking(id int64) (*StockPicking, error) {
	sps, err := c.GetStockPickings([]int64{id})
	if err != nil {
		return nil, err
	}
	if sps != nil && len(*sps) > 0 {
		return &((*sps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.picking not found", id)
}

// GetStockPickings gets stock.picking existing records.
func (c *Client) GetStockPickings(ids []int64) (*StockPickings, error) {
	sps := &StockPickings{}
	if err := c.Read(StockPickingModel, ids, nil, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPicking finds stock.picking record by querying it with criteria.
func (c *Client) FindStockPicking(criteria *Criteria) (*StockPicking, error) {
	sps := &StockPickings{}
	if err := c.SearchRead(StockPickingModel, criteria, NewOptions().Limit(1), sps); err != nil {
		return nil, err
	}
	if sps != nil && len(*sps) > 0 {
		return &((*sps)[0]), nil
	}
	return nil, fmt.Errorf("stock.picking was not found with criteria %v", criteria)
}

// FindStockPickings finds stock.picking records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickings(criteria *Criteria, options *Options) (*StockPickings, error) {
	sps := &StockPickings{}
	if err := c.SearchRead(StockPickingModel, criteria, options, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPickingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockPickingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockPickingId finds record id by querying it with criteria.
func (c *Client) FindStockPickingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPickingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.picking was not found with criteria %v and options %v", criteria, options)
}
