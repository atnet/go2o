package impl

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ixre/go2o/core/domain/interface/chat"
	"github.com/ixre/go2o/core/domain/interface/member"
	"github.com/ixre/go2o/core/service/proto"
)

var _ proto.ChatServiceServer = new(chatServiceImpl)

type chatServiceImpl struct {
	repo       chat.IChatRepository
	memberRepo member.IMemberRepo
	proto.UnimplementedChatServiceServer
	serviceUtil
}

func NewChatService(repo chat.IChatRepository, memberRepo member.IMemberRepo) proto.ChatServiceServer {
	return &chatServiceImpl{
		repo:       repo,
		memberRepo: memberRepo,
	}
}

// GetConversation implements proto.ChatServiceServer.
func (c *chatServiceImpl) GetConversation(_ context.Context, req *proto.ChatConversationRequest) (*proto.ChatConversationResponse, error) {
	iu := c.repo.GetChatUser(int(req.Uid))
	ic, err := iu.BuildConversation(int(req.Rid), chat.ChatType(req.ChatType), req.OutOrderNo)
	if err != nil {
		return &proto.ChatConversationResponse{
			Code:    1,
			Message: err.Error(),
		}, nil
	}
	v := ic.Get()

	ret := &proto.ChatConversationResponse{
		ConvId:       int64(ic.GetDomainId()),
		Key:          v.Key,
		Sid:          int64(v.Sid),
		Rid:          int64(v.Rid),
		ChatType:     int32(v.ChatType),
		LastMsg:      v.LastMsg,
		LastChatTime: int64(v.LastChatTime),
	}
	if v.ChatType == chat.ChatTypeNormal {
		// 绑定会员用户代码(聊天对方用户)
		im := c.memberRepo.GetMember(req.Rid)
		if im != nil {
			iv := im.GetValue()
			ret.ChatProfile = &proto.SChatUserProfile{
				ProfilePhoto: iv.ProfilePhoto,
				UserCode:     iv.UserCode,
				Nickname:     iv.Nickname,
				UserRole:     int32(iv.RoleFlag),
			}
		}
	}
	return ret, nil
}

// BindOutOrderNo 绑定外部订单号
func (c *chatServiceImpl) BindOutOrderNo(_ context.Context, req *proto.BindChatOutOrderNoRequest) (*proto.TxResult, error) {
	iu := c.repo.GetChatUser(int(req.Uid))
	ic := iu.GetConversation(int(req.ConvId))
	if ic == nil {
		return c.errorV2(errors.New("no such conversation")), nil
	}
	err := ic.BindOutOrderNo(req.OutOrderNo)
	return c.errorV2(err), nil
}

// DeleteMsg implements proto.ChatServiceServer.
func (c *chatServiceImpl) DeleteMsg(_ context.Context, req *proto.MsgIdRequest) (*proto.TxResult, error) {
	var err error
	iu := c.repo.GetChatUser(int(req.Uid))
	ic := iu.GetConversation(int(req.ConvId))
	if ic == nil {
		err = errors.New("no such conversation")
	} else {
		err = ic.DeleteMsg(int(req.MsgId))
	}
	return c.errorV2(err), nil
}

// DestroyConversation implements proto.ChatServiceServer.
func (c *chatServiceImpl) DestroyConversation(_ context.Context, req *proto.ConversationIdRequest) (*proto.TxResult, error) {
	var err error
	iu := c.repo.GetChatUser(int(req.Uid))
	ic := iu.GetConversation(int(req.ConvId))
	if ic == nil {
		err = errors.New("no such conversation")
	} else {
		err = ic.Destroy()
	}
	return c.errorV2(err), nil
}

// FetchHistoryMsgList implements proto.ChatServiceServer.
func (c *chatServiceImpl) FetchMsgList(_ context.Context, req *proto.FetchMsgRequest) (*proto.FetchMsgResponse, error) {
	var err error
	iu := c.repo.GetChatUser(int(req.Uid))
	ic := iu.GetConversation(int(req.ConvId))
	if ic == nil {
		return &proto.FetchMsgResponse{}, err
	}
	var list []*chat.ChatMsg
	if req.HistoryMsg {
		list = ic.FetchHistoryMsgList(int(req.LastTime), int(req.Size))
	} else {
		list = ic.FetchMsgList(int(req.LastTime), int(req.Size))
	}
	arr := make([]*proto.SMsg, len(list))
	for i, v := range list {
		arr[i] = c.parseChatMsg(v)
	}
	return &proto.FetchMsgResponse{
		MsgList: arr,
		IsOver:  false,
	}, nil
}

func (c *chatServiceImpl) parseChatMsg(v *chat.ChatMsg) *proto.SMsg {
	var mp map[string]string
	json.Unmarshal([]byte(v.Extra), &mp)
	return &proto.SMsg{
		MsgId:      int64(v.Id),
		Sid:        int64(v.Sid),
		MsgType:    int32(v.MsgType),
		MsgFlag:    int32(v.MsgFlag),
		Content:    v.Content,
		Extra:      mp,
		CreateTime: int64(v.CreateTime),
	}
}

// GetMsg implements proto.ChatServiceServer.
func (c *chatServiceImpl) GetMsg(_ context.Context, req *proto.MsgIdRequest) (*proto.SMsg, error) {
	var err error
	iu := c.repo.GetChatUser(int(req.Uid))
	ic := iu.GetConversation(int(req.ConvId))
	if ic == nil {
		return &proto.SMsg{}, err
	}
	msg := ic.GetMsg(int(req.MsgId))
	return c.parseChatMsg(msg), nil
}

// RevertMsg implements proto.ChatServiceServer.
func (c *chatServiceImpl) RevertMsg(_ context.Context, req *proto.MsgIdRequest) (*proto.TxResult, error) {
	iu := c.repo.GetChatUser(int(req.Uid))
	ic := iu.GetConversation(int(req.ConvId))
	var err error
	if ic == nil {
		err = errors.New("no such conversation")
	} else {
		err = ic.RevertMsg(int(req.MsgId))
	}
	return c.errorV2(err), nil
}

// Send implements proto.ChatServiceServer.
func (c *chatServiceImpl) Send(_ context.Context, req *proto.SendMsgRequest) (*proto.SendMsgResponse, error) {
	iu := c.repo.GetChatUser(int(req.Uid))
	ic := iu.GetConversation(int(req.ConvId))
	var err error
	if ic != nil {
		id, err1 := ic.Send(&chat.MsgBody{
			MsgType: int(req.MsgType),
			Content: req.Content,
			Extra:   req.Extra,
		})
		if err1 == nil {
			return &proto.SendMsgResponse{
				MsgId: int64(id),
			}, nil
		}
		err = err1
	}
	return &proto.SendMsgResponse{
		Code:    1,
		Messsge: err.Error(),
	}, nil
}

// UpdateMsgAttrs implements proto.ChatServiceServer.
func (c *chatServiceImpl) UpdateMsgAttrs(_ context.Context, req *proto.UpdateMsgAttrRequest) (*proto.TxResult, error) {
	iu := c.repo.GetChatUser(int(req.Uid))
	ic := iu.GetConversation(int(req.ConvId))
	var err error
	if ic == nil {
		err = errors.New("no such conversation")
	} else {
		err = ic.UpdateMsgAttrs(int(req.MsgId), req.Attr)
	}
	return c.errorV2(err), nil
}
