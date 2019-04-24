package classfile

type MemberInfo struct {
	cp				ConstantPool
	accessFlags		uint16
	nameIndex 		uint16
	descriptorIndex	uint16
	attributes		[]AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMembers(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo {
		cp:				cp,
		accessFlags:	reader.readUint16(),
		nameIndex:		reader.readUint16(),
		descriptorIndex:reader.readUint16(),
		attributes:		readAttributes(reader, cp),
	}
}

func (mi *MemberInfo) Name() string {
	return mi.cp.getUtf8(mi.nameIndex)
}

func (mi* MemberInfo) Descriptor() string {
	return mi.cp.getUtf8(mi.descriptorIndex)
}