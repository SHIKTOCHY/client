// Auto-generated by avdl-compiler v1.3.13 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/notify.avdl

package chat1

import (
	"errors"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type ChatActivityType int

const (
	ChatActivityType_RESERVED         ChatActivityType = 0
	ChatActivityType_INCOMING_MESSAGE ChatActivityType = 1
	ChatActivityType_READ_MESSAGE     ChatActivityType = 2
	ChatActivityType_NEW_CONVERSATION ChatActivityType = 3
	ChatActivityType_SET_STATUS       ChatActivityType = 4
	ChatActivityType_FAILED_MESSAGE   ChatActivityType = 5
)

var ChatActivityTypeMap = map[string]ChatActivityType{
	"RESERVED":         0,
	"INCOMING_MESSAGE": 1,
	"READ_MESSAGE":     2,
	"NEW_CONVERSATION": 3,
	"SET_STATUS":       4,
	"FAILED_MESSAGE":   5,
}

var ChatActivityTypeRevMap = map[ChatActivityType]string{
	0: "RESERVED",
	1: "INCOMING_MESSAGE",
	2: "READ_MESSAGE",
	3: "NEW_CONVERSATION",
	4: "SET_STATUS",
	5: "FAILED_MESSAGE",
}

func (e ChatActivityType) String() string {
	if v, ok := ChatActivityTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type IncomingMessage struct {
	Message    MessageUnboxed     `codec:"message" json:"message"`
	ConvID     ConversationID     `codec:"convID" json:"convID"`
	Conv       *ConversationLocal `codec:"conv,omitempty" json:"conv,omitempty"`
	Pagination *Pagination        `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type ReadMessageInfo struct {
	ConvID ConversationID     `codec:"convID" json:"convID"`
	MsgID  MessageID          `codec:"msgID" json:"msgID"`
	Conv   *ConversationLocal `codec:"conv,omitempty" json:"conv,omitempty"`
}

type NewConversationInfo struct {
	Conv ConversationLocal `codec:"conv" json:"conv"`
}

type SetStatusInfo struct {
	ConvID ConversationID     `codec:"convID" json:"convID"`
	Status ConversationStatus `codec:"status" json:"status"`
	Conv   *ConversationLocal `codec:"conv,omitempty" json:"conv,omitempty"`
}

type FailedMessageInfo struct {
	OutboxRecords []OutboxRecord `codec:"outboxRecords" json:"outboxRecords"`
}

type ChatActivity struct {
	ActivityType__    ChatActivityType     `codec:"activityType" json:"activityType"`
	IncomingMessage__ *IncomingMessage     `codec:"incomingMessage,omitempty" json:"incomingMessage,omitempty"`
	ReadMessage__     *ReadMessageInfo     `codec:"readMessage,omitempty" json:"readMessage,omitempty"`
	NewConversation__ *NewConversationInfo `codec:"newConversation,omitempty" json:"newConversation,omitempty"`
	SetStatus__       *SetStatusInfo       `codec:"setStatus,omitempty" json:"setStatus,omitempty"`
	FailedMessage__   *FailedMessageInfo   `codec:"failedMessage,omitempty" json:"failedMessage,omitempty"`
}

func (o *ChatActivity) ActivityType() (ret ChatActivityType, err error) {
	switch o.ActivityType__ {
	case ChatActivityType_INCOMING_MESSAGE:
		if o.IncomingMessage__ == nil {
			err = errors.New("unexpected nil value for IncomingMessage__")
			return ret, err
		}
	case ChatActivityType_READ_MESSAGE:
		if o.ReadMessage__ == nil {
			err = errors.New("unexpected nil value for ReadMessage__")
			return ret, err
		}
	case ChatActivityType_NEW_CONVERSATION:
		if o.NewConversation__ == nil {
			err = errors.New("unexpected nil value for NewConversation__")
			return ret, err
		}
	case ChatActivityType_SET_STATUS:
		if o.SetStatus__ == nil {
			err = errors.New("unexpected nil value for SetStatus__")
			return ret, err
		}
	case ChatActivityType_FAILED_MESSAGE:
		if o.FailedMessage__ == nil {
			err = errors.New("unexpected nil value for FailedMessage__")
			return ret, err
		}
	}
	return o.ActivityType__, nil
}

func (o ChatActivity) IncomingMessage() (res IncomingMessage) {
	if o.ActivityType__ != ChatActivityType_INCOMING_MESSAGE {
		panic("wrong case accessed")
	}
	if o.IncomingMessage__ == nil {
		return
	}
	return *o.IncomingMessage__
}

func (o ChatActivity) ReadMessage() (res ReadMessageInfo) {
	if o.ActivityType__ != ChatActivityType_READ_MESSAGE {
		panic("wrong case accessed")
	}
	if o.ReadMessage__ == nil {
		return
	}
	return *o.ReadMessage__
}

func (o ChatActivity) NewConversation() (res NewConversationInfo) {
	if o.ActivityType__ != ChatActivityType_NEW_CONVERSATION {
		panic("wrong case accessed")
	}
	if o.NewConversation__ == nil {
		return
	}
	return *o.NewConversation__
}

func (o ChatActivity) SetStatus() (res SetStatusInfo) {
	if o.ActivityType__ != ChatActivityType_SET_STATUS {
		panic("wrong case accessed")
	}
	if o.SetStatus__ == nil {
		return
	}
	return *o.SetStatus__
}

func (o ChatActivity) FailedMessage() (res FailedMessageInfo) {
	if o.ActivityType__ != ChatActivityType_FAILED_MESSAGE {
		panic("wrong case accessed")
	}
	if o.FailedMessage__ == nil {
		return
	}
	return *o.FailedMessage__
}

func NewChatActivityWithIncomingMessage(v IncomingMessage) ChatActivity {
	return ChatActivity{
		ActivityType__:    ChatActivityType_INCOMING_MESSAGE,
		IncomingMessage__: &v,
	}
}

func NewChatActivityWithReadMessage(v ReadMessageInfo) ChatActivity {
	return ChatActivity{
		ActivityType__: ChatActivityType_READ_MESSAGE,
		ReadMessage__:  &v,
	}
}

func NewChatActivityWithNewConversation(v NewConversationInfo) ChatActivity {
	return ChatActivity{
		ActivityType__:    ChatActivityType_NEW_CONVERSATION,
		NewConversation__: &v,
	}
}

func NewChatActivityWithSetStatus(v SetStatusInfo) ChatActivity {
	return ChatActivity{
		ActivityType__: ChatActivityType_SET_STATUS,
		SetStatus__:    &v,
	}
}

func NewChatActivityWithFailedMessage(v FailedMessageInfo) ChatActivity {
	return ChatActivity{
		ActivityType__:  ChatActivityType_FAILED_MESSAGE,
		FailedMessage__: &v,
	}
}

type NewChatActivityArg struct {
	Uid      keybase1.UID `codec:"uid" json:"uid"`
	Activity ChatActivity `codec:"activity" json:"activity"`
}

type ChatIdentifyUpdateArg struct {
	Update keybase1.CanonicalTLFNameAndIDWithBreaks `codec:"update" json:"update"`
}

type ChatTLFFinalizeArg struct {
	Uid          keybase1.UID             `codec:"uid" json:"uid"`
	ConvID       ConversationID           `codec:"convID" json:"convID"`
	FinalizeInfo ConversationFinalizeInfo `codec:"finalizeInfo" json:"finalizeInfo"`
	Conv         *ConversationLocal       `codec:"conv,omitempty" json:"conv,omitempty"`
}

type ChatTLFResolveArg struct {
	Uid         keybase1.UID            `codec:"uid" json:"uid"`
	ConvID      ConversationID          `codec:"convID" json:"convID"`
	ResolveInfo ConversationResolveInfo `codec:"resolveInfo" json:"resolveInfo"`
}

type ChatInboxStaleArg struct {
	Uid keybase1.UID `codec:"uid" json:"uid"`
}

type ChatThreadsStaleArg struct {
	Uid     keybase1.UID     `codec:"uid" json:"uid"`
	ConvIDs []ConversationID `codec:"convIDs" json:"convIDs"`
}

type NotifyChatInterface interface {
	NewChatActivity(context.Context, NewChatActivityArg) error
	ChatIdentifyUpdate(context.Context, keybase1.CanonicalTLFNameAndIDWithBreaks) error
	ChatTLFFinalize(context.Context, ChatTLFFinalizeArg) error
	ChatTLFResolve(context.Context, ChatTLFResolveArg) error
	ChatInboxStale(context.Context, keybase1.UID) error
	ChatThreadsStale(context.Context, ChatThreadsStaleArg) error
}

func NotifyChatProtocol(i NotifyChatInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "chat.1.NotifyChat",
		Methods: map[string]rpc.ServeHandlerDescription{
			"NewChatActivity": {
				MakeArg: func() interface{} {
					ret := make([]NewChatActivityArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NewChatActivityArg)
					if !ok {
						err = rpc.NewTypeError((*[]NewChatActivityArg)(nil), args)
						return
					}
					err = i.NewChatActivity(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatIdentifyUpdate": {
				MakeArg: func() interface{} {
					ret := make([]ChatIdentifyUpdateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatIdentifyUpdateArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatIdentifyUpdateArg)(nil), args)
						return
					}
					err = i.ChatIdentifyUpdate(ctx, (*typedArgs)[0].Update)
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatTLFFinalize": {
				MakeArg: func() interface{} {
					ret := make([]ChatTLFFinalizeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatTLFFinalizeArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatTLFFinalizeArg)(nil), args)
						return
					}
					err = i.ChatTLFFinalize(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatTLFResolve": {
				MakeArg: func() interface{} {
					ret := make([]ChatTLFResolveArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatTLFResolveArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatTLFResolveArg)(nil), args)
						return
					}
					err = i.ChatTLFResolve(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatInboxStale": {
				MakeArg: func() interface{} {
					ret := make([]ChatInboxStaleArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatInboxStaleArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatInboxStaleArg)(nil), args)
						return
					}
					err = i.ChatInboxStale(ctx, (*typedArgs)[0].Uid)
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"ChatThreadsStale": {
				MakeArg: func() interface{} {
					ret := make([]ChatThreadsStaleArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatThreadsStaleArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatThreadsStaleArg)(nil), args)
						return
					}
					err = i.ChatThreadsStale(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
		},
	}
}

type NotifyChatClient struct {
	Cli rpc.GenericClient
}

func (c NotifyChatClient) NewChatActivity(ctx context.Context, __arg NewChatActivityArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.NewChatActivity", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatIdentifyUpdate(ctx context.Context, update keybase1.CanonicalTLFNameAndIDWithBreaks) (err error) {
	__arg := ChatIdentifyUpdateArg{Update: update}
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatIdentifyUpdate", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatTLFFinalize(ctx context.Context, __arg ChatTLFFinalizeArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatTLFFinalize", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatTLFResolve(ctx context.Context, __arg ChatTLFResolveArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatTLFResolve", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatInboxStale(ctx context.Context, uid keybase1.UID) (err error) {
	__arg := ChatInboxStaleArg{Uid: uid}
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatInboxStale", []interface{}{__arg})
	return
}

func (c NotifyChatClient) ChatThreadsStale(ctx context.Context, __arg ChatThreadsStaleArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.NotifyChat.ChatThreadsStale", []interface{}{__arg})
	return
}
