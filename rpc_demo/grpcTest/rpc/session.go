package rpc

import (
	"encoding/binary"
	"io"
	"net"
)

// 自己实现一个rpc连接的回话，方便理解

// 会话连接的结构体
type Session struct {
	conn net.Conn
}

// 创建新连接
func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

// 向连接中写数据
func (s *Session) Write(data []byte) error {
	// 定义数据传输格式： handler uint32 + data []byte
	// 4字节+ 数据长度的切片
	buf := make([]byte, 4+len(data))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data))) //在写入 handler数据记录数据长度
	// 写入数据
	copy(buf[4:], data)
	_, err := s.conn.Write(buf) // 向连接中写入数据
	if err != nil {
		return err
	}
	return nil
}

// 从连接里读数据
func (s *Session) Read() ([]byte, error) {
	// 读取头部长度
	header := make([]byte, 4)
	// 按头部长度，读取头部数据
	_, err := io.ReadFull(s.conn, header)
	if err != nil {
		return nil, err
	}
	// 读取数据长度
	dataLen := binary.BigEndian.Uint32(header)
	// 按照数据长度去读取数据
	data := make([]byte, dataLen)
	_, err = io.ReadFull(s.conn, data) // 从连接里读取数据
	if err != nil {
		return nil, err
	}
	return data, nil
}
