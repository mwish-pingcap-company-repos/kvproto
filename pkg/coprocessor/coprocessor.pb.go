// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coprocessor.proto

package coprocessor

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	_ "github.com/gogo/protobuf/gogoproto"

	errorpb "github.com/pingcap/kvproto/pkg/errorpb"

	kvrpcpb "github.com/pingcap/kvproto/pkg/kvrpcpb"

	metapb "github.com/pingcap/kvproto/pkg/metapb"

	github_com_pingcap_kvproto_pkg_sharedbytes "github.com/pingcap/kvproto/pkg/sharedbytes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Context              *kvrpcpb.Context   `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Tp                   int64              `protobuf:"varint,2,opt,name=tp,proto3" json:"tp,omitempty"`
	Data                 []byte             `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Ranges               []*metapb.KeyRange `protobuf:"bytes,4,rep,name=ranges" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_coprocessor_6a050c2ff227e345, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetContext() *kvrpcpb.Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Request) GetTp() int64 {
	if m != nil {
		return m.Tp
	}
	return 0
}

func (m *Request) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Request) GetRanges() []*metapb.KeyRange {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type Response struct {
	Data                 github_com_pingcap_kvproto_pkg_sharedbytes.SharedBytes `protobuf:"bytes,1,opt,name=data,proto3,customtype=github.com/pingcap/kvproto/pkg/sharedbytes.SharedBytes" json:"data"`
	RegionError          *errorpb.Error                                         `protobuf:"bytes,2,opt,name=region_error,json=regionError" json:"region_error,omitempty"`
	Locked               *kvrpcpb.LockInfo                                      `protobuf:"bytes,3,opt,name=locked" json:"locked,omitempty"`
	OtherError           string                                                 `protobuf:"bytes,4,opt,name=other_error,json=otherError,proto3" json:"other_error,omitempty"`
	Range                *metapb.KeyRange                                       `protobuf:"bytes,5,opt,name=range" json:"range,omitempty"`
	ExecDetails          *kvrpcpb.ExecDetails                                   `protobuf:"bytes,6,opt,name=exec_details,json=execDetails" json:"exec_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                               `json:"-"`
	XXX_unrecognized     []byte                                                 `json:"-"`
	XXX_sizecache        int32                                                  `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_coprocessor_6a050c2ff227e345, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetRegionError() *errorpb.Error {
	if m != nil {
		return m.RegionError
	}
	return nil
}

func (m *Response) GetLocked() *kvrpcpb.LockInfo {
	if m != nil {
		return m.Locked
	}
	return nil
}

func (m *Response) GetOtherError() string {
	if m != nil {
		return m.OtherError
	}
	return ""
}

func (m *Response) GetRange() *metapb.KeyRange {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *Response) GetExecDetails() *kvrpcpb.ExecDetails {
	if m != nil {
		return m.ExecDetails
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "coprocessor.Request")
	proto.RegisterType((*Response)(nil), "coprocessor.Response")
}
func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Context != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.Context.Size()))
		n1, err := m.Context.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Tp != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.Tp))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if len(m.Ranges) > 0 {
		for _, msg := range m.Ranges {
			dAtA[i] = 0x22
			i++
			i = encodeVarintCoprocessor(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintCoprocessor(dAtA, i, uint64(m.Data.Size()))
	n2, err := m.Data.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.RegionError != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.RegionError.Size()))
		n3, err := m.RegionError.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Locked != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.Locked.Size()))
		n4, err := m.Locked.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.OtherError) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(len(m.OtherError)))
		i += copy(dAtA[i:], m.OtherError)
	}
	if m.Range != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.Range.Size()))
		n5, err := m.Range.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.ExecDetails != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.ExecDetails.Size()))
		n6, err := m.ExecDetails.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintCoprocessor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.Context != nil {
		l = m.Context.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.Tp != 0 {
		n += 1 + sovCoprocessor(uint64(m.Tp))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if len(m.Ranges) > 0 {
		for _, e := range m.Ranges {
			l = e.Size()
			n += 1 + l + sovCoprocessor(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	l = m.Data.Size()
	n += 1 + l + sovCoprocessor(uint64(l))
	if m.RegionError != nil {
		l = m.RegionError.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.Locked != nil {
		l = m.Locked.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	l = len(m.OtherError)
	if l > 0 {
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.Range != nil {
		l = m.Range.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.ExecDetails != nil {
		l = m.ExecDetails.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCoprocessor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCoprocessor(x uint64) (n int) {
	return sovCoprocessor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoprocessor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Context == nil {
				m.Context = &kvrpcpb.Context{}
			}
			if err := m.Context.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ranges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ranges = append(m.Ranges, &metapb.KeyRange{})
			if err := m.Ranges[len(m.Ranges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoprocessor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoprocessor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoprocessor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionError", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RegionError == nil {
				m.RegionError = &errorpb.Error{}
			}
			if err := m.RegionError.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locked", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Locked == nil {
				m.Locked = &kvrpcpb.LockInfo{}
			}
			if err := m.Locked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OtherError = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Range", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Range == nil {
				m.Range = &metapb.KeyRange{}
			}
			if err := m.Range.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExecDetails == nil {
				m.ExecDetails = &kvrpcpb.ExecDetails{}
			}
			if err := m.ExecDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoprocessor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoprocessor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCoprocessor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoprocessor
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
					return 0, ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCoprocessor
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCoprocessor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCoprocessor
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCoprocessor(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCoprocessor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoprocessor   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("coprocessor.proto", fileDescriptor_coprocessor_6a050c2ff227e345) }

var fileDescriptor_coprocessor_6a050c2ff227e345 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0xb1, 0x6e, 0xd4, 0x40,
	0x10, 0xcd, 0xde, 0x5d, 0x2e, 0x61, 0xed, 0x84, 0x64, 0x95, 0xc2, 0x4a, 0xe1, 0xb3, 0x52, 0x44,
	0x86, 0xc2, 0x16, 0x87, 0x04, 0x1d, 0xc5, 0x41, 0x0a, 0x04, 0xd5, 0xf2, 0x01, 0x91, 0xbd, 0x1e,
	0xf6, 0x2c, 0x27, 0x9e, 0x65, 0x77, 0xef, 0x74, 0xa9, 0xf9, 0x09, 0x3e, 0x81, 0xaf, 0xa0, 0x4e,
	0x49, 0x89, 0x28, 0x22, 0x74, 0xfc, 0x08, 0xf2, 0x7a, 0x7d, 0x20, 0xaa, 0x7d, 0xf3, 0xde, 0xec,
	0x9b, 0x99, 0x47, 0x4f, 0x05, 0x2a, 0x8d, 0x02, 0x8c, 0x41, 0x9d, 0x29, 0x8d, 0x16, 0x59, 0xf0,
	0x0f, 0x75, 0x7e, 0x04, 0x5a, 0xa3, 0x56, 0x65, 0xaf, 0x9d, 0x1f, 0x35, 0x6b, 0xad, 0xc4, 0xae,
	0x3c, 0x93, 0x28, 0xd1, 0xc1, 0xbc, 0x43, 0x9e, 0x7d, 0xac, 0x57, 0xc6, 0x3a, 0xe8, 0x89, 0xf0,
	0x16, 0x6c, 0x31, 0x7c, 0xba, 0xf8, 0x4c, 0xe8, 0x01, 0x87, 0x4f, 0x2b, 0x30, 0x96, 0x3d, 0xa5,
	0x07, 0x02, 0x5b, 0x0b, 0x1b, 0x1b, 0x91, 0x84, 0xa4, 0xc1, 0xfc, 0x24, 0x1b, 0x26, 0xbc, 0xee,
	0x79, 0x3e, 0x34, 0xb0, 0x63, 0x3a, 0xb2, 0x2a, 0x1a, 0x25, 0x24, 0x1d, 0xf3, 0x91, 0x55, 0x8c,
	0xd1, 0x49, 0x55, 0xd8, 0x22, 0x1a, 0x27, 0x24, 0x0d, 0xb9, 0xc3, 0x2c, 0xa5, 0x53, 0x5d, 0xb4,
	0x12, 0x4c, 0x34, 0x49, 0xc6, 0xce, 0xce, 0x8f, 0x7e, 0x07, 0x77, 0xbc, 0x13, 0xb8, 0xd7, 0x2f,
	0xbe, 0x8d, 0xe8, 0x21, 0x07, 0xa3, 0xb0, 0x35, 0xc0, 0xb8, 0xb7, 0xea, 0x76, 0x08, 0x17, 0xaf,
	0xee, 0x1f, 0x66, 0x7b, 0x3f, 0x1f, 0x66, 0x2f, 0x64, 0x6d, 0x97, 0xab, 0x32, 0x13, 0x78, 0x9b,
	0xab, 0xba, 0x95, 0xa2, 0x50, 0x79, 0xb3, 0xee, 0xcf, 0x55, 0x8d, 0xcc, 0xcd, 0xb2, 0xd0, 0x50,
	0x95, 0x77, 0x16, 0x4c, 0xf6, 0xc1, 0xe1, 0x45, 0x87, 0xfd, 0x2a, 0xcf, 0x68, 0xa8, 0x41, 0xd6,
	0xd8, 0x5e, 0xbb, 0x08, 0xdd, 0xe2, 0xc1, 0xfc, 0x38, 0x1b, 0x02, 0xbd, 0xea, 0x5e, 0x1e, 0xf4,
	0x3d, 0xae, 0x60, 0x4f, 0xe8, 0xf4, 0x06, 0x45, 0x03, 0x95, 0xbb, 0x29, 0x98, 0x9f, 0xee, 0xc2,
	0x78, 0x8f, 0xa2, 0x79, 0xdb, 0x7e, 0x44, 0xee, 0x1b, 0xd8, 0x8c, 0x06, 0x68, 0x97, 0xa0, 0xbd,
	0xf9, 0x24, 0x21, 0xe9, 0x23, 0x4e, 0x1d, 0xd5, 0x7b, 0x5d, 0xd2, 0x7d, 0x77, 0x69, 0xb4, 0xef,
	0x73, 0xfd, 0x3f, 0x88, 0x5e, 0x66, 0x2f, 0x69, 0x08, 0x1b, 0x10, 0xd7, 0x15, 0xd8, 0xa2, 0xbe,
	0x31, 0xd1, 0xd4, 0xb5, 0x9f, 0xed, 0x26, 0x5f, 0x6d, 0x40, 0xbc, 0xe9, 0x35, 0x1e, 0xc0, 0xdf,
	0x62, 0x71, 0xf9, 0xe3, 0xeb, 0x21, 0xb9, 0xdf, 0xc6, 0xe4, 0xfb, 0x36, 0x26, 0xbf, 0xb6, 0x31,
	0xf9, 0xf2, 0x3b, 0xde, 0xa3, 0x27, 0xa8, 0x65, 0x66, 0xeb, 0x66, 0x9d, 0xf9, 0xa4, 0xca, 0xa9,
	0x7b, 0x9e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xc5, 0x6a, 0xde, 0x6a, 0x02, 0x00, 0x00,
}
