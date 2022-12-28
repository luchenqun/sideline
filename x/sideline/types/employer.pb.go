// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sideline/sideline/employer.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Employer struct {
	Index     string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Introduce string   `protobuf:"bytes,3,opt,name=introduce,proto3" json:"introduce,omitempty"`
	Email     string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Avatar    string   `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Address   string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	TaskCount uint64   `protobuf:"varint,7,opt,name=taskCount,proto3" json:"taskCount,omitempty"`
	TaskFail  uint64   `protobuf:"varint,8,opt,name=taskFail,proto3" json:"taskFail,omitempty"`
	Feedbacks []string `protobuf:"bytes,9,rep,name=feedbacks,proto3" json:"feedbacks,omitempty"`
}

func (m *Employer) Reset()         { *m = Employer{} }
func (m *Employer) String() string { return proto.CompactTextString(m) }
func (*Employer) ProtoMessage()    {}
func (*Employer) Descriptor() ([]byte, []int) {
	return fileDescriptor_30760f168d2f15ab, []int{0}
}
func (m *Employer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Employer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Employer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Employer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Employer.Merge(m, src)
}
func (m *Employer) XXX_Size() int {
	return m.Size()
}
func (m *Employer) XXX_DiscardUnknown() {
	xxx_messageInfo_Employer.DiscardUnknown(m)
}

var xxx_messageInfo_Employer proto.InternalMessageInfo

func (m *Employer) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Employer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Employer) GetIntroduce() string {
	if m != nil {
		return m.Introduce
	}
	return ""
}

func (m *Employer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Employer) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Employer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Employer) GetTaskCount() uint64 {
	if m != nil {
		return m.TaskCount
	}
	return 0
}

func (m *Employer) GetTaskFail() uint64 {
	if m != nil {
		return m.TaskFail
	}
	return 0
}

func (m *Employer) GetFeedbacks() []string {
	if m != nil {
		return m.Feedbacks
	}
	return nil
}

func init() {
	proto.RegisterType((*Employer)(nil), "sideline.sideline.Employer")
}

func init() { proto.RegisterFile("sideline/sideline/employer.proto", fileDescriptor_30760f168d2f15ab) }

var fileDescriptor_30760f168d2f15ab = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x63, 0x9a, 0xa6, 0x89, 0x77, 0x58, 0x08, 0x0d, 0x08, 0x59, 0x11, 0xab, 0xae, 0xca,
	0xa2, 0x37, 0x00, 0xc1, 0x01, 0xba, 0x64, 0xe7, 0xd6, 0x83, 0x64, 0x35, 0xb1, 0x23, 0xdb, 0x45,
	0xed, 0x2d, 0x38, 0x16, 0xcb, 0x2e, 0x59, 0xa2, 0xe4, 0x06, 0x9c, 0x00, 0xd9, 0xcd, 0xcf, 0xee,
	0x7d, 0xef, 0xd9, 0x33, 0xa3, 0x47, 0x4b, 0xa7, 0x24, 0x56, 0x4a, 0xe3, 0xd3, 0x28, 0xb0, 0x6e,
	0x2a, 0x73, 0x42, 0xbb, 0x6a, 0xac, 0xf1, 0x86, 0x5d, 0x0f, 0xc1, 0x6a, 0x10, 0x8f, 0x7f, 0x84,
	0xe6, 0xaf, 0xfd, 0x2b, 0x76, 0x43, 0xe7, 0x4a, 0x4b, 0x3c, 0x02, 0x29, 0xc9, 0xb2, 0xd8, 0x5c,
	0x80, 0x31, 0x9a, 0x6a, 0x51, 0x23, 0x5c, 0x45, 0x33, 0x6a, 0xf6, 0x40, 0x0b, 0xa5, 0xbd, 0x35,
	0xf2, 0xb0, 0x43, 0x98, 0xc5, 0x60, 0x32, 0xc2, 0x1c, 0xac, 0x85, 0xaa, 0x20, 0xbd, 0xcc, 0x89,
	0xc0, 0x6e, 0x69, 0x26, 0x3e, 0x85, 0x17, 0x16, 0xe6, 0xd1, 0xee, 0x89, 0x01, 0x5d, 0x08, 0x29,
	0x2d, 0x3a, 0x07, 0x59, 0x0c, 0x06, 0x0c, 0x5b, 0xbc, 0x70, 0xfb, 0x17, 0x73, 0xd0, 0x1e, 0x16,
	0x25, 0x59, 0xa6, 0x9b, 0xc9, 0x60, 0xf7, 0x34, 0x0f, 0xf0, 0x16, 0x16, 0xe5, 0x31, 0x1c, 0x39,
	0xfc, 0xfc, 0x40, 0x94, 0x5b, 0xb1, 0xdb, 0x3b, 0x28, 0xca, 0x59, 0xb8, 0x6f, 0x34, 0x9e, 0xd7,
	0xdf, 0x2d, 0x27, 0xe7, 0x96, 0x93, 0xdf, 0x96, 0x93, 0xaf, 0x8e, 0x27, 0xe7, 0x8e, 0x27, 0x3f,
	0x1d, 0x4f, 0xde, 0xef, 0xc6, 0xea, 0x8e, 0x53, 0x8b, 0xfe, 0xd4, 0xa0, 0xdb, 0x66, 0xb1, 0xc3,
	0xf5, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x67, 0x5f, 0xd6, 0x67, 0x01, 0x00, 0x00,
}

func (m *Employer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Employer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Employer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Feedbacks) > 0 {
		for iNdEx := len(m.Feedbacks) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Feedbacks[iNdEx])
			copy(dAtA[i:], m.Feedbacks[iNdEx])
			i = encodeVarintEmployer(dAtA, i, uint64(len(m.Feedbacks[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.TaskFail != 0 {
		i = encodeVarintEmployer(dAtA, i, uint64(m.TaskFail))
		i--
		dAtA[i] = 0x40
	}
	if m.TaskCount != 0 {
		i = encodeVarintEmployer(dAtA, i, uint64(m.TaskCount))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEmployer(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Avatar) > 0 {
		i -= len(m.Avatar)
		copy(dAtA[i:], m.Avatar)
		i = encodeVarintEmployer(dAtA, i, uint64(len(m.Avatar)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintEmployer(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Introduce) > 0 {
		i -= len(m.Introduce)
		copy(dAtA[i:], m.Introduce)
		i = encodeVarintEmployer(dAtA, i, uint64(len(m.Introduce)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEmployer(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintEmployer(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEmployer(dAtA []byte, offset int, v uint64) int {
	offset -= sovEmployer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Employer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovEmployer(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEmployer(uint64(l))
	}
	l = len(m.Introduce)
	if l > 0 {
		n += 1 + l + sovEmployer(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovEmployer(uint64(l))
	}
	l = len(m.Avatar)
	if l > 0 {
		n += 1 + l + sovEmployer(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEmployer(uint64(l))
	}
	if m.TaskCount != 0 {
		n += 1 + sovEmployer(uint64(m.TaskCount))
	}
	if m.TaskFail != 0 {
		n += 1 + sovEmployer(uint64(m.TaskFail))
	}
	if len(m.Feedbacks) > 0 {
		for _, s := range m.Feedbacks {
			l = len(s)
			n += 1 + l + sovEmployer(uint64(l))
		}
	}
	return n
}

func sovEmployer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEmployer(x uint64) (n int) {
	return sovEmployer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Employer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEmployer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Employer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Employer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmployer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmployer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmployer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmployer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Introduce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmployer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmployer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Introduce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmployer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmployer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avatar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmployer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmployer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Avatar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmployer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmployer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskCount", wireType)
			}
			m.TaskCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaskCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskFail", wireType)
			}
			m.TaskFail = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaskFail |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feedbacks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEmployer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEmployer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEmployer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feedbacks = append(m.Feedbacks, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEmployer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEmployer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEmployer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEmployer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEmployer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEmployer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEmployer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEmployer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEmployer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEmployer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEmployer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEmployer = fmt.Errorf("proto: unexpected end of group")
)
