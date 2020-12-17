package websockets

const UserChat = 1
const GroupChat = 2

type SenderMessage struct {
	SenderId int     `json:"sender_id"`
	Message
}
type ReceiverMessage struct {
	ReceiverId int     `json:"receiver_id"`
	Message
}

type Message struct {
	MessageType int    `json:"message_type"`
	Data        string `json:"data"`
}

