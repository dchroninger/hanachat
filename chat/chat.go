package chat

import (
	"time"

	"github.com/dchroninger/hanachat/user"
	"github.com/google/uuid"
)

type ChatRoom struct {
	ChatRoomID          ChatRoomID
	ChatRoomName        ChatRoomName
	ChatRoomDescription ChatRoomDescription
	ChatRoomUsers       ChatRoomUsers
	ChatRoomMessages    ChatRoomMessages
}

func New(i, n, d string) *ChatRoom {
	return &ChatRoom{
		ChatRoomID:          ChatRoomID(i),
		ChatRoomName:        ChatRoomName(n),
		ChatRoomDescription: ChatRoomDescription(d),
	}
}

func (c *ChatRoom) AddUser(u user.User) {
	c.ChatRoomUsers[u.UserID] = u
}

func (c *ChatRoom) RemoveUser(u user.User) {
	delete(c.ChatRoomUsers, u.UserID)
}

func (c *ChatRoom) AddMessage(m ChatRoomMessage) {
	c.ChatRoomMessages = append(c.ChatRoomMessages, m)
}

type ChatRoomID string

func NewChatRoomID() ChatRoomID {
	return ChatRoomID(uuid.NewString())
}

func (id ChatRoomID) String() string {
	return string(id)
}

type ChatRoomName string

func (n ChatRoomName) String() string {
	return string(n)
}

type ChatRoomDescription string

func (d ChatRoomDescription) String() string {
	return string(d)
}

type ChatRoomUsers map[user.UserID]user.User

type ChatRoomMessages []ChatRoomMessage

type ChatRoomMessage struct {
	ID      ChatRoomMessageID
	Author  user.User
	Message ChatRoomMessageText
	Time    time.Time
}

type ChatRoomMessageID string

type ChatRoomMessageText string

func (m ChatRoomMessageText) String() string {
	return string(m)
}
