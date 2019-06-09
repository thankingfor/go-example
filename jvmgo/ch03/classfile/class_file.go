package classfile

import "fmt"

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	//magic      uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

//读取class文件
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader) //检验魔数
	self.readAndCheckVersion(reader) //检验版本号
	self.constantPool = readConstantPool(reader)  //常量池
	self.accessFlags = reader.readUint16() //读取访问标记
	self.thisClass = reader.readUint16() //读取这个class的常量池索引
	self.superClass = reader.readUint16() //读取父class的常量池索引
	self.interfaces = reader.readUint16s() //读取接口索引表
	self.fields = readMembers(reader, self.constantPool) //读取字段
	self.methods = readMembers(reader, self.constantPool) //读取方法
	self.attributes = readAttributes(reader, self.constantPool) //
}

//魔数是固定的  class文件的魔数为 0xCAFEBABE  例如  PDF 为 %PDF  ZIP 为 PK
//JAVA虚拟机规定，如果class文件不符合要求 jvm要抛出java.lang.ClassFormatError异常。目前还无法抛出异常 所以用panic()方法终止程序执行
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

//版本号 Oracle的实现是完全兼容后面的版本  所以Java8  支持 45-52的class文件。
//如果碰见其他版本号 暂时先调用panic方法终止程序 抛出java.lang.UnsupportedClassVersionError
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()  //最小版本
	self.majorVersion = reader.readUint16()  //最大版本
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
