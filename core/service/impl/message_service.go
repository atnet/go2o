/**
 * Copyright 2015 @ 56x.net.
 * name : message_service.go
 * author : jarryliu
 * date : 2016-06-11 20:51
 * description :
 * history :
 */
package impl

import (
	"context"
	"fmt"

	mss "github.com/ixre/go2o/core/domain/interface/message"
	"github.com/ixre/go2o/core/dto"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/gof/types"
)

var _ proto.MessageServiceServer = new(messageService)

type messageService struct {
	_repo mss.IMessageRepo
	serviceUtil
	proto.UnimplementedMessageServiceServer
}

func NewMessageService(repo mss.IMessageRepo) proto.MessageServiceServer {
	return &messageService{
		_repo: repo,
	}
}

func (m *messageService) SaveNotifyTemplate(_ context.Context, r *proto.SaveNotifyTemplateRequest) (*proto.TxResult, error) {
	mgr := m._repo.NotifyManager()
	err := mgr.SaveNotifyTemplate(&mss.NotifyTemplate{
		Id:      int(r.Id),
		TplCode: r.TplCode,
		TplType: int(r.TplType),
		TplName: r.TplName,
		Content: r.Content,
		Labels:  r.Labels,
		SpCode:  r.SpCode,
		SpTid:   r.SpTid,
	})
	return m.errorV2(err), nil
}

// 获取通知项配置
func (m *messageService) GetNotifyItem(_ context.Context, key *proto.String) (*proto.SNotifyItem, error) {
	it := m._repo.NotifyManager().GetNotifyItem(key.Value)
	return &proto.SNotifyItem{
		Key:        it.Key,
		NotifyBy:   int32(it.NotifyBy),
		ReadonlyBy: it.ReadonlyBy,
		TplId:      int32(it.TplId),
		Content:    it.Content,
		Tags:       it.Tags,
	}, nil
}

func (m *messageService) GetAllNotifyItem(_ context.Context, empty *proto.Empty) (*proto.NotifyItemListResponse, error) {
	list := m._repo.NotifyManager().GetAllNotifyItem()
	arr := make([]*proto.SNotifyItem, len(list))
	for i, v := range list {
		arr[i] = m.parseNotifyItemDto(v)
	}
	return &proto.NotifyItemListResponse{
		Value: arr,
	}, nil
}

// 发送短信
func (m *messageService) SendPhoneMessage(_ context.Context, r *proto.SendMessageRequest) (*proto.TxResult, error) {
	mg := m._repo.NotifyManager()
	err := mg.SendPhoneMessage(r.Account, mss.PhoneMessage(r.Message), r.Data, r.TemplateId)
	if err != nil {
		return m.errorV2(err), nil
	}
	return m.successV2(nil), nil
}

// 保存通知项设置
func (m *messageService) SaveNotifyItem(_ context.Context, item *proto.SNotifyItem) (*proto.Result, error) {
	v := m.parseNotifyItem(item)
	err := m._repo.NotifyManager().SaveNotifyItem(v)
	return m.error(err), nil
}

// 获取邮件模版
func (m *messageService) GetMailTemplate(_ context.Context, id *proto.Int64) (*proto.SMailTemplate, error) {
	v := m._repo.GetProvider().GetMailTemplate(int32(id.Value))
	if v != nil {
		return m.parseMailTemplateDto(v), nil
	}
	return nil, fmt.Errorf("no such mail template")
}

// 保存邮件模板
func (m *messageService) SaveMailTemplate(_ context.Context, src *proto.SMailTemplate) (*proto.Result, error) {
	v := m.parseMailTemplate(src)
	_, err := m._repo.GetProvider().SaveMailTemplate(v)
	return m.error(err), nil
}

// 获取邮件模板
func (m *messageService) GetMailTemplates(_ context.Context, empty *proto.Empty) (*proto.MailTemplateListResponse, error) {
	list := m._repo.GetProvider().GetMailTemplates()
	arr := make([]*proto.SMailTemplate, len(list))
	for i, v := range list {
		arr[i] = m.parseMailTemplateDto(v)
	}
	return &proto.MailTemplateListResponse{
		Value: arr,
	}, nil
}

// 删除邮件模板
func (m *messageService) DeleteMailTemplate(_ context.Context, id *proto.Int64) (*proto.Result, error) {
	err := m._repo.GetProvider().DeleteMailTemplate(int32(id.Value))
	return m.error(err), nil
}

// 发送站内信
func (m *messageService) SendSiteMessage(_ context.Context, r *proto.SendSiteMessageRequest) (*proto.Result, error) {
	v := &mss.Message{
		Id: 0,
		// 消息类型
		Type: mss.TypeSiteMessage,
		// 消息用途
		UseFor: mss.UseForNotify,
		// 发送人角色
		SenderRole: mss.RoleSystem,
		// 发送人编号
		SenderId: int32(r.SenderId),
		To: []mss.User{
			{Id: int32(r.ReceiverId), Role: int(r.ReceiverType)},
		},
		// 发送的用户角色
		ToRole: -1,
		// 全系统接收
		AllUser: -1,
		// 是否只能阅读
		Readonly: 1,
	}
	var err error
	im := m._repo.MessageManager().CreateMessage(v, &mss.SiteMessage{
		Subject: r.Msg.Subject,
		Message: r.Msg.Message,
	})
	if _, err = im.Save(); err == nil {
		err = im.Send(nil)
	}
	return m.error(err), nil
}

// 获取邮件绑定
func (m *messageService) GetConfig() mss.Config {
	return m._repo.GetProvider().GetConfig()
}

// 保存邮件
func (m *messageService) SaveConfig(conf *mss.Config) error {
	return m._repo.GetProvider().SaveConfig(conf)
}

// 可通过外部添加
func (m *messageService) RegisterNotifyItem(key string, item *mss.NotifyItem) {
	mss.RegisterNotifyItem(key, item)
}

// todo: 考虑弄一个,确定后再发送.这样可以先在系统,然后才发送
// 发送站内通知信息,
// toRole: 为-1时发送给所有用户
// sendNow: 是否马上发送
func (m *messageService) SendSiteNotifyMessage(senderId int32, toRole int,
	msg *mss.SiteMessage, sendNow bool) error {
	v := &mss.Message{
		Id: 0,
		// 消息类型
		Type: mss.TypeSiteMessage,
		// 消息用途
		UseFor: mss.UseForNotify,
		// 发送人角色
		SenderRole: mss.RoleSystem,
		// 发送人编号
		SenderId: senderId,
		// 发送的用户角色
		ToRole: -1,
		// 全系统接收
		AllUser: -1,
		// 是否只能阅读
		Readonly: 1,
	}

	if toRole <= 0 {
		v.AllUser = 1
	} else {
		v.ToRole = toRole
	}
	var err error
	im := m._repo.MessageManager().CreateMessage(v, msg)
	if _, err = im.Save(); err == nil {
		err = im.Send(nil)
	}
	return err
}

// 对会用户发送站内信
func (m *messageService) SendSiteMessageToUser(senderId int32, toRole int, toUser int64,
	msg *mss.SiteMessage, sendNow bool) error {
	v := &mss.Message{
		Id: 0,
		// 消息类型
		Type: mss.TypeSiteMessage,
		// 消息用途
		UseFor: mss.UseForNotify,
		// 发送人角色
		SenderRole: mss.RoleSystem,
		// 发送人编号
		SenderId: senderId,
		To: []mss.User{
			{Id: int32(toUser), Role: toRole},
		},
		// 发送的用户角色
		ToRole: -1,
		// 全系统接收
		AllUser: -1,
		// 是否只能阅读
		Readonly: 1,
	}
	var err error
	im := m._repo.MessageManager().CreateMessage(v, msg)
	if _, err = im.Save(); err == nil {
		err = im.Send(nil)
	}
	return err
}

// 获取站内信
func (m *messageService) GetSiteMessage(id, toUserId int64, toRole int) *dto.SiteMessage {
	//todo: id int64
	msg := m._repo.MessageManager().GetMessage(int32(id))
	if msg != nil && msg.CheckPerm(int32(toUserId), toRole) {
		val := msg.GetValue()
		dto := &dto.SiteMessage{
			Id:           val.Id,
			Type:         val.Type,
			UseFor:       val.UseFor,
			SenderUserId: 0,
			SenderName:   "系统",
			Readonly:     val.Readonly,
			CreateTime:   val.CreateTime,
			ToId:         int32(toUserId),
			ToRole:       toRole,
		}

		switch msg.Type() {
		case mss.TypeSMS:
			dto.Data = msg.(mss.IPhoneMessage).Value()
		case mss.TypeEmail:
			dto.Data = msg.(mss.IMailMessage).Value()
		case mss.TypeSiteMessage:
			dto.Data = msg.(mss.ISiteMessage).Value()
		}

		if msg.SpecialTo() {
			if to := msg.GetTo(int32(toUserId), toRole); to != nil {
				dto.HasRead = to.HasRead
				dto.ReadTime = to.ReadTime
			}
		}
		return dto
	}
	return nil
}

// 获取聊天会话编号
func (m *messageService) GetChatSessionId(senderRole int, senderId int32, toRole int, toId int32) int32 {
	return m._repo.MessageManager().GetChatSessionId(senderRole, senderId, toRole, toId)
}

// 创建聊天会话
func (m *messageService) CreateChatSession(senderRole int, senderId int32, toRole int, toId int32) (mss.Message, error) {
	return m._repo.MessageManager().CreateChatSession(senderRole, senderId, toRole, toId)
}

func (m *messageService) parseNotifyItemDto(v mss.NotifyItem) *proto.SNotifyItem {
	return &proto.SNotifyItem{
		Key:        v.Key,
		NotifyBy:   int32(v.NotifyBy),
		ReadonlyBy: v.ReadonlyBy,
		TplId:      int32(v.TplId),
		Content:    v.Content,
		Tags:       v.Tags,
	}
}

func (m *messageService) parseNotifyItem(v *proto.SNotifyItem) *mss.NotifyItem {
	return &mss.NotifyItem{
		Key:        v.Key,
		NotifyBy:   int(v.NotifyBy),
		ReadonlyBy: v.ReadonlyBy,
		TplId:      int(v.TplId),
		Content:    v.Content,
		Tags:       v.Tags,
	}
}

func (m *messageService) parseMailTemplateDto(v *mss.MailTemplate) *proto.SMailTemplate {
	return &proto.SMailTemplate{
		Id:         int64(v.Id),
		MerchantId: v.MerchantId,
		Name:       v.Name,
		Subject:    v.Subject,
		Body:       v.Body,
		Enabled:    v.Enabled == 1,
	}
}

func (m *messageService) parseMailTemplate(v *proto.SMailTemplate) *mss.MailTemplate {
	return &mss.MailTemplate{
		Id:         int32(v.Id),
		MerchantId: v.MerchantId,
		Name:       v.Name,
		Subject:    v.Subject,
		Body:       v.Body,
		Enabled:    types.ElseInt(v.Enabled, 1, 0),
	}
}
