package odoo

import (
	"fmt"
)

// DiscussChannel represents discuss.channel model.
type DiscussChannel struct {
	Active                       *Bool      `xmlrpc:"active,omptempty"`
	AllowPublicUpload            *Bool      `xmlrpc:"allow_public_upload,omptempty"`
	AnonymousName                *String    `xmlrpc:"anonymous_name,omptempty"`
	Avatar128                    *String    `xmlrpc:"avatar_128,omptempty"`
	ChannelMemberIds             *Relation  `xmlrpc:"channel_member_ids,omptempty"`
	ChannelPartnerIds            *Relation  `xmlrpc:"channel_partner_ids,omptempty"`
	ChannelType                  *Selection `xmlrpc:"channel_type,omptempty"`
	ChatbotCurrentStepId         *Many2One  `xmlrpc:"chatbot_current_step_id,omptempty"`
	ChatbotMessageIds            *Relation  `xmlrpc:"chatbot_message_ids,omptempty"`
	CountryId                    *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	DefaultDisplayMode           *Selection `xmlrpc:"default_display_mode,omptempty"`
	Description                  *String    `xmlrpc:"description,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	Duration                     *Float     `xmlrpc:"duration,omptempty"`
	GroupIds                     *Relation  `xmlrpc:"group_ids,omptempty"`
	GroupPublicId                *Many2One  `xmlrpc:"group_public_id,omptempty"`
	HasMessage                   *Bool      `xmlrpc:"has_message,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	Image128                     *String    `xmlrpc:"image_128,omptempty"`
	InvitationUrl                *String    `xmlrpc:"invitation_url,omptempty"`
	IsChat                       *Bool      `xmlrpc:"is_chat,omptempty"`
	IsEditable                   *Bool      `xmlrpc:"is_editable,omptempty"`
	IsMember                     *Bool      `xmlrpc:"is_member,omptempty"`
	LivechatActive               *Bool      `xmlrpc:"livechat_active,omptempty"`
	LivechatChannelId            *Many2One  `xmlrpc:"livechat_channel_id,omptempty"`
	LivechatOperatorId           *Many2One  `xmlrpc:"livechat_operator_id,omptempty"`
	MemberCount                  *Int       `xmlrpc:"member_count,omptempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	PinnedMessageIds             *Relation  `xmlrpc:"pinned_message_ids,omptempty"`
	RatingAvg                    *Float     `xmlrpc:"rating_avg,omptempty"`
	RatingAvgText                *Selection `xmlrpc:"rating_avg_text,omptempty"`
	RatingCount                  *Int       `xmlrpc:"rating_count,omptempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omptempty"`
	RatingLastFeedback           *String    `xmlrpc:"rating_last_feedback,omptempty"`
	RatingLastImage              *String    `xmlrpc:"rating_last_image,omptempty"`
	RatingLastText               *Selection `xmlrpc:"rating_last_text,omptempty"`
	RatingLastValue              *Float     `xmlrpc:"rating_last_value,omptempty"`
	RatingPercentageSatisfaction *Float     `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	RtcSessionIds                *Relation  `xmlrpc:"rtc_session_ids,omptempty"`
	SfuChannelUuid               *String    `xmlrpc:"sfu_channel_uuid,omptempty"`
	SfuServerUrl                 *String    `xmlrpc:"sfu_server_url,omptempty"`
	Uuid                         *String    `xmlrpc:"uuid,omptempty"`
	WaAccountId                  *Many2One  `xmlrpc:"wa_account_id,omptempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WhatsappChannelValidUntil    *Time      `xmlrpc:"whatsapp_channel_valid_until,omptempty"`
	WhatsappMailMessageId        *Many2One  `xmlrpc:"whatsapp_mail_message_id,omptempty"`
	WhatsappNumber               *String    `xmlrpc:"whatsapp_number,omptempty"`
	WhatsappPartnerId            *Many2One  `xmlrpc:"whatsapp_partner_id,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DiscussChannels represents array of discuss.channel model.
type DiscussChannels []DiscussChannel

// DiscussChannelModel is the odoo model name.
const DiscussChannelModel = "discuss.channel"

// Many2One convert DiscussChannel to *Many2One.
func (dc *DiscussChannel) Many2One() *Many2One {
	return NewMany2One(dc.Id.Get(), "")
}

// CreateDiscussChannel creates a new discuss.channel model and returns its id.
func (c *Client) CreateDiscussChannel(dc *DiscussChannel) (int64, error) {
	ids, err := c.CreateDiscussChannels([]*DiscussChannel{dc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDiscussChannel creates a new discuss.channel model and returns its id.
func (c *Client) CreateDiscussChannels(dcs []*DiscussChannel) ([]int64, error) {
	var vv []interface{}
	for _, v := range dcs {
		vv = append(vv, v)
	}
	return c.Create(DiscussChannelModel, vv)
}

// UpdateDiscussChannel updates an existing discuss.channel record.
func (c *Client) UpdateDiscussChannel(dc *DiscussChannel) error {
	return c.UpdateDiscussChannels([]int64{dc.Id.Get()}, dc)
}

// UpdateDiscussChannels updates existing discuss.channel records.
// All records (represented by ids) will be updated by dc values.
func (c *Client) UpdateDiscussChannels(ids []int64, dc *DiscussChannel) error {
	return c.Update(DiscussChannelModel, ids, dc)
}

// DeleteDiscussChannel deletes an existing discuss.channel record.
func (c *Client) DeleteDiscussChannel(id int64) error {
	return c.DeleteDiscussChannels([]int64{id})
}

// DeleteDiscussChannels deletes existing discuss.channel records.
func (c *Client) DeleteDiscussChannels(ids []int64) error {
	return c.Delete(DiscussChannelModel, ids)
}

// GetDiscussChannel gets discuss.channel existing record.
func (c *Client) GetDiscussChannel(id int64) (*DiscussChannel, error) {
	dcs, err := c.GetDiscussChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	if dcs != nil && len(*dcs) > 0 {
		return &((*dcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of discuss.channel not found", id)
}

// GetDiscussChannels gets discuss.channel existing records.
func (c *Client) GetDiscussChannels(ids []int64) (*DiscussChannels, error) {
	dcs := &DiscussChannels{}
	if err := c.Read(DiscussChannelModel, ids, nil, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDiscussChannel finds discuss.channel record by querying it with criteria.
func (c *Client) FindDiscussChannel(criteria *Criteria) (*DiscussChannel, error) {
	dcs := &DiscussChannels{}
	if err := c.SearchRead(DiscussChannelModel, criteria, NewOptions().Limit(1), dcs); err != nil {
		return nil, err
	}
	if dcs != nil && len(*dcs) > 0 {
		return &((*dcs)[0]), nil
	}
	return nil, fmt.Errorf("discuss.channel was not found with criteria %v", criteria)
}

// FindDiscussChannels finds discuss.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussChannels(criteria *Criteria, options *Options) (*DiscussChannels, error) {
	dcs := &DiscussChannels{}
	if err := c.SearchRead(DiscussChannelModel, criteria, options, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDiscussChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DiscussChannelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDiscussChannelId finds record id by querying it with criteria.
func (c *Client) FindDiscussChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DiscussChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("discuss.channel was not found with criteria %v and options %v", criteria, options)
}
