package classfile

type ConstantUtf8Info struct {
	str string
}

func (cui *ConstantUtf8Info) readInfo(reader *ClassReader) {
	cui.str = decodeMUTF8(reader.readBytes(uint32(reader.readUint16())))
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}