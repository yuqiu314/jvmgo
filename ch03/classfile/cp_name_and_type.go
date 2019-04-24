package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex		uint16
	descriptorIndex	uint16
}

func (cni *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	cni.nameIndex = reader.readUint16()
	cni.descriptorIndex = reader.readUint16()
}

