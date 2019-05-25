package net

import "zinxStudy/szinx/ziface"

//
type Message struct {
	Id      uint32
	Datalen uint32
	Data    []byte
}

func NewMessage(id uint32, data []byte) ziface.IMessage {
	datalen := uint32(len(data))
	return &Message{
		id,
		datalen,
		data,
	}
}


func (m *Message) GetMsgId() uint32 {
	return m.Id
}
func (m *Message) GetMsgLen() uint32 {
	return m.Datalen
}
func (m *Message) GetMsgData() []byte {
	return m.Data
}

//setter
func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}
func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}
func (m *Message) SetMsgLen(len uint32) {
	m.Datalen = len
}
