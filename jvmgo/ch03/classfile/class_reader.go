package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}


func (self *ClassReader) readUint8() uint8  {//u1
	val := self.data[0]
	self.data = self.data[1:]
	return val
}
//ClassReader并没有使用索引记录数据位置，而是使用Go语言的reslice语法跳过已经读取的数据。ReadUint16()读取u2类型数据
func (self *ClassReader) readUint16() uint16  {//u2
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}
func (self *ClassReader) readUint32() uint32  {//u4
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}
func (self *ClassReader) readUint64() uint64  {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}
//读取uint16表 表的大小由开头的uint16数据指出
func (self *ClassReader) readUint16s() []uint16  {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}