package server

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type Session struct {
	conn net.Conn
}

// NewSession 从网络连接新建一个会话
func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

func (s *Session) Write(data []byte) error {
	buf := make([]byte, 4+len(data))
	// Header
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	// Data
	copy(buf[4:], data)

	_, err := s.conn.Write(buf)

	return err
}

// Read 从 Session 中读数据
func (s *Session) Read() ([]byte, error) {
	// 读取 Header，获取 Data 长度信息

	header := make([]byte, 4)
	if _, err := io.ReadFull(s.conn, header); err != nil {
		fmt.Println("read err", err)
		return nil, err
	}
	dataLen := binary.BigEndian.Uint32(header)

	// 按照 dataLen 读取 Data
	data := make([]byte, dataLen)
	if _, err := io.ReadFull(s.conn, data); err != nil {
		return nil, err
	}
	return data, nil
}
