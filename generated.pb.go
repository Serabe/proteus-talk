// Code generated by protoc-gen-gogo.
// source: github.com/Serabe/proteus-talk/github.com/Serabe/proteus-talk/generated.proto
// DO NOT EDIT!

/*
	Package gomad is a generated protocol buffer package.

	It is generated from these files:
		github.com/Serabe/proteus-talk/github.com/Serabe/proteus-talk/generated.proto

	It has these top-level messages:
		IngredientSelection
		Pizza
		AddIngredientRequest
		NutritionalReportResponse
		RemoveIngredientRequest
*/
package gomad

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var Ingredient_name = map[int32]string{
	0: "PEPPERONI",
	1: "PINEAPPLE",
	2: "HAM",
	3: "BACON",
	4: "CHICKEN",
	5: "LAMB",
	6: "NO_FOOD",
}
var Ingredient_value = map[string]int32{
	"PEPPERONI": 0,
	"PINEAPPLE": 1,
	"HAM":       2,
	"BACON":     3,
	"CHICKEN":   4,
	"LAMB":      5,
	"NO_FOOD":   6,
}

func (Ingredient) EnumDescriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *IngredientSelection) Reset()                    { *m = IngredientSelection{} }
func (m *IngredientSelection) String() string            { return proto.CompactTextString(m) }
func (*IngredientSelection) ProtoMessage()               {}
func (*IngredientSelection) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *Pizza) Reset()                    { *m = Pizza{} }
func (m *Pizza) String() string            { return proto.CompactTextString(m) }
func (*Pizza) ProtoMessage()               {}
func (*Pizza) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

type AddIngredientRequest struct {
	Arg1 *Pizza     `protobuf:"bytes,1,opt,name=arg1" json:"arg1,omitempty"`
	Arg2 Ingredient `protobuf:"varint,2,opt,name=arg2,proto3,enum=github.com.Serabe.proteustalk.Ingredient" json:"arg2,omitempty"`
	Arg3 int32      `protobuf:"varint,3,opt,name=arg3,proto3" json:"arg3,omitempty"`
}

func (m *AddIngredientRequest) Reset()                    { *m = AddIngredientRequest{} }
func (m *AddIngredientRequest) String() string            { return proto.CompactTextString(m) }
func (*AddIngredientRequest) ProtoMessage()               {}
func (*AddIngredientRequest) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *AddIngredientRequest) GetArg1() *Pizza {
	if m != nil {
		return m.Arg1
	}
	return nil
}

func (m *AddIngredientRequest) GetArg2() Ingredient {
	if m != nil {
		return m.Arg2
	}
	return Pepperoni
}

func (m *AddIngredientRequest) GetArg3() int32 {
	if m != nil {
		return m.Arg3
	}
	return 0
}

type NutritionalReportResponse struct {
	Result1 string `protobuf:"bytes,1,opt,name=result1,proto3" json:"result1,omitempty"`
}

func (m *NutritionalReportResponse) Reset()         { *m = NutritionalReportResponse{} }
func (m *NutritionalReportResponse) String() string { return proto.CompactTextString(m) }
func (*NutritionalReportResponse) ProtoMessage()    {}
func (*NutritionalReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{3}
}

func (m *NutritionalReportResponse) GetResult1() string {
	if m != nil {
		return m.Result1
	}
	return ""
}

type RemoveIngredientRequest struct {
	Arg1 *Pizza     `protobuf:"bytes,1,opt,name=arg1" json:"arg1,omitempty"`
	Arg2 Ingredient `protobuf:"varint,2,opt,name=arg2,proto3,enum=github.com.Serabe.proteustalk.Ingredient" json:"arg2,omitempty"`
	Arg3 int32      `protobuf:"varint,3,opt,name=arg3,proto3" json:"arg3,omitempty"`
}

func (m *RemoveIngredientRequest) Reset()                    { *m = RemoveIngredientRequest{} }
func (m *RemoveIngredientRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveIngredientRequest) ProtoMessage()               {}
func (*RemoveIngredientRequest) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{4} }

func (m *RemoveIngredientRequest) GetArg1() *Pizza {
	if m != nil {
		return m.Arg1
	}
	return nil
}

func (m *RemoveIngredientRequest) GetArg2() Ingredient {
	if m != nil {
		return m.Arg2
	}
	return Pepperoni
}

func (m *RemoveIngredientRequest) GetArg3() int32 {
	if m != nil {
		return m.Arg3
	}
	return 0
}

func init() {
	proto.RegisterType((*IngredientSelection)(nil), "github.com.Serabe.proteustalk.IngredientSelection")
	proto.RegisterType((*Pizza)(nil), "github.com.Serabe.proteustalk.Pizza")
	proto.RegisterType((*AddIngredientRequest)(nil), "github.com.Serabe.proteustalk.AddIngredientRequest")
	proto.RegisterType((*NutritionalReportResponse)(nil), "github.com.Serabe.proteustalk.NutritionalReportResponse")
	proto.RegisterType((*RemoveIngredientRequest)(nil), "github.com.Serabe.proteustalk.RemoveIngredientRequest")
	proto.RegisterEnum("github.com.Serabe.proteustalk.Ingredient", Ingredient_name, Ingredient_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ProteustalkService service

type ProteustalkServiceClient interface {
	AddIngredient(ctx context.Context, in *AddIngredientRequest, opts ...grpc.CallOption) (*Pizza, error)
	NutritionalReport(ctx context.Context, in *Pizza, opts ...grpc.CallOption) (*NutritionalReportResponse, error)
	RemoveIngredient(ctx context.Context, in *RemoveIngredientRequest, opts ...grpc.CallOption) (*Pizza, error)
}

type proteustalkServiceClient struct {
	cc *grpc.ClientConn
}

func NewProteustalkServiceClient(cc *grpc.ClientConn) ProteustalkServiceClient {
	return &proteustalkServiceClient{cc}
}

func (c *proteustalkServiceClient) AddIngredient(ctx context.Context, in *AddIngredientRequest, opts ...grpc.CallOption) (*Pizza, error) {
	out := new(Pizza)
	err := grpc.Invoke(ctx, "/github.com.Serabe.proteustalk.ProteustalkService/AddIngredient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proteustalkServiceClient) NutritionalReport(ctx context.Context, in *Pizza, opts ...grpc.CallOption) (*NutritionalReportResponse, error) {
	out := new(NutritionalReportResponse)
	err := grpc.Invoke(ctx, "/github.com.Serabe.proteustalk.ProteustalkService/NutritionalReport", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proteustalkServiceClient) RemoveIngredient(ctx context.Context, in *RemoveIngredientRequest, opts ...grpc.CallOption) (*Pizza, error) {
	out := new(Pizza)
	err := grpc.Invoke(ctx, "/github.com.Serabe.proteustalk.ProteustalkService/RemoveIngredient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProteustalkService service

type ProteustalkServiceServer interface {
	AddIngredient(context.Context, *AddIngredientRequest) (*Pizza, error)
	NutritionalReport(context.Context, *Pizza) (*NutritionalReportResponse, error)
	RemoveIngredient(context.Context, *RemoveIngredientRequest) (*Pizza, error)
}

func RegisterProteustalkServiceServer(s *grpc.Server, srv ProteustalkServiceServer) {
	s.RegisterService(&_ProteustalkService_serviceDesc, srv)
}

func _ProteustalkService_AddIngredient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddIngredientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProteustalkServiceServer).AddIngredient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.Serabe.proteustalk.ProteustalkService/AddIngredient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProteustalkServiceServer).AddIngredient(ctx, req.(*AddIngredientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProteustalkService_NutritionalReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pizza)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProteustalkServiceServer).NutritionalReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.Serabe.proteustalk.ProteustalkService/NutritionalReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProteustalkServiceServer).NutritionalReport(ctx, req.(*Pizza))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProteustalkService_RemoveIngredient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveIngredientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProteustalkServiceServer).RemoveIngredient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.Serabe.proteustalk.ProteustalkService/RemoveIngredient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProteustalkServiceServer).RemoveIngredient(ctx, req.(*RemoveIngredientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProteustalkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.Serabe.proteustalk.ProteustalkService",
	HandlerType: (*ProteustalkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddIngredient",
			Handler:    _ProteustalkService_AddIngredient_Handler,
		},
		{
			MethodName: "NutritionalReport",
			Handler:    _ProteustalkService_NutritionalReport_Handler,
		},
		{
			MethodName: "RemoveIngredient",
			Handler:    _ProteustalkService_RemoveIngredient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/Serabe/proteus-talk/github.com/Serabe/proteus-talk/generated.proto",
}

func (m *IngredientSelection) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IngredientSelection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ingredient != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Ingredient))
	}
	if m.Portions != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Portions))
	}
	return i, nil
}

func (m *Pizza) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pizza) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.ID))
	}
	if len(m.Ingredients) > 0 {
		for _, msg := range m.Ingredients {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.ProtoSize()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *AddIngredientRequest) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddIngredientRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Arg1 != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Arg1.ProtoSize()))
		n1, err := m.Arg1.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Arg2 != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Arg2))
	}
	if m.Arg3 != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Arg3))
	}
	return i, nil
}

func (m *NutritionalReportResponse) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NutritionalReportResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Result1) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Result1)))
		i += copy(dAtA[i:], m.Result1)
	}
	return i, nil
}

func (m *RemoveIngredientRequest) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveIngredientRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Arg1 != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Arg1.ProtoSize()))
		n2, err := m.Arg1.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Arg2 != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Arg2))
	}
	if m.Arg3 != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Arg3))
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *IngredientSelection) ProtoSize() (n int) {
	var l int
	_ = l
	if m.Ingredient != 0 {
		n += 1 + sovGenerated(uint64(m.Ingredient))
	}
	if m.Portions != 0 {
		n += 1 + sovGenerated(uint64(m.Portions))
	}
	return n
}

func (m *Pizza) ProtoSize() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovGenerated(uint64(m.ID))
	}
	if len(m.Ingredients) > 0 {
		for _, e := range m.Ingredients {
			l = e.ProtoSize()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *AddIngredientRequest) ProtoSize() (n int) {
	var l int
	_ = l
	if m.Arg1 != nil {
		l = m.Arg1.ProtoSize()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Arg2 != 0 {
		n += 1 + sovGenerated(uint64(m.Arg2))
	}
	if m.Arg3 != 0 {
		n += 1 + sovGenerated(uint64(m.Arg3))
	}
	return n
}

func (m *NutritionalReportResponse) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.Result1)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *RemoveIngredientRequest) ProtoSize() (n int) {
	var l int
	_ = l
	if m.Arg1 != nil {
		l = m.Arg1.ProtoSize()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Arg2 != 0 {
		n += 1 + sovGenerated(uint64(m.Arg2))
	}
	if m.Arg3 != 0 {
		n += 1 + sovGenerated(uint64(m.Arg3))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IngredientSelection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IngredientSelection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IngredientSelection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ingredient", wireType)
			}
			m.Ingredient = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ingredient |= (Ingredient(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Portions", wireType)
			}
			m.Portions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Portions |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *Pizza) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: Pizza: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pizza: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ingredients", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ingredients = append(m.Ingredients, &IngredientSelection{})
			if err := m.Ingredients[len(m.Ingredients)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *AddIngredientRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: AddIngredientRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddIngredientRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arg1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Arg1 == nil {
				m.Arg1 = &Pizza{}
			}
			if err := m.Arg1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arg2", wireType)
			}
			m.Arg2 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Arg2 |= (Ingredient(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arg3", wireType)
			}
			m.Arg3 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Arg3 |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *NutritionalReportResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: NutritionalReportResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NutritionalReportResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *RemoveIngredientRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: RemoveIngredientRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveIngredientRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arg1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Arg1 == nil {
				m.Arg1 = &Pizza{}
			}
			if err := m.Arg1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arg2", wireType)
			}
			m.Arg2 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Arg2 |= (Ingredient(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arg3", wireType)
			}
			m.Arg3 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Arg3 |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/Serabe/proteus-talk/github.com/Serabe/proteus-talk/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xb5, 0x93, 0x3a, 0x69, 0x6e, 0xbe, 0x7e, 0x32, 0x43, 0x45, 0x83, 0x01, 0xd7, 0x8a, 0x58,
	0x14, 0xa4, 0xa6, 0x22, 0x15, 0xa8, 0x42, 0x62, 0x91, 0xa4, 0xa9, 0x1a, 0xd1, 0x3a, 0x96, 0xcb,
	0x8a, 0x0d, 0x72, 0xe2, 0x8b, 0x3b, 0x34, 0xf1, 0x18, 0xff, 0x74, 0xd1, 0x0d, 0x3b, 0x54, 0xf9,
	0x09, 0xd8, 0x58, 0x2a, 0x6a, 0x85, 0x78, 0x0c, 0x36, 0x48, 0x2c, 0x79, 0x02, 0x84, 0xd2, 0x05,
	0x5b, 0x1e, 0x01, 0x79, 0x5a, 0x92, 0xf0, 0x53, 0x12, 0x76, 0xec, 0xe6, 0xce, 0x9c, 0x73, 0xee,
	0x3d, 0xd7, 0x47, 0x86, 0x6d, 0x87, 0x86, 0xbb, 0x51, 0xa7, 0xd2, 0x65, 0xfd, 0x95, 0x1d, 0xf4,
	0xad, 0x0e, 0xae, 0x78, 0x3e, 0x0b, 0x31, 0x0a, 0x96, 0x43, 0xab, 0xb7, 0xb7, 0x32, 0xe9, 0x19,
	0x5d, 0xf4, 0xad, 0x10, 0xed, 0x4a, 0x7a, 0xcd, 0xc8, 0x8d, 0x11, 0xbe, 0x72, 0x86, 0xaf, 0x9c,
	0xe3, 0x53, 0xb8, 0xb2, 0x3c, 0x26, 0xe7, 0x30, 0x87, 0x71, 0x31, 0xd6, 0x89, 0x9e, 0xf2, 0x8a,
	0x17, 0xfc, 0x74, 0xa6, 0x56, 0x7e, 0x29, 0xc2, 0xe5, 0x96, 0xeb, 0xf8, 0x68, 0x53, 0x74, 0xc3,
	0x1d, 0xec, 0x61, 0x37, 0xa4, 0xcc, 0x25, 0x2d, 0x00, 0x3a, 0xbc, 0x2e, 0x89, 0x9a, 0xb8, 0xf4,
	0x7f, 0xf5, 0x56, 0xe5, 0x8f, 0xad, 0x2b, 0x23, 0x1d, 0x73, 0x8c, 0x4c, 0x14, 0x98, 0xf5, 0x98,
	0x9f, 0xaa, 0x06, 0xa5, 0x8c, 0x26, 0x2e, 0x49, 0xe6, 0xb0, 0xbe, 0x3f, 0x7b, 0x78, 0xb4, 0x28,
	0x7c, 0x7d, 0xbd, 0x28, 0x94, 0x5f, 0x80, 0x64, 0xd0, 0x83, 0x03, 0x8b, 0x5c, 0x81, 0x0c, 0xb5,
	0x79, 0x47, 0xa9, 0x9e, 0x1b, 0x7c, 0x5a, 0xcc, 0xb4, 0xd6, 0xcd, 0x0c, 0xb5, 0xc9, 0x23, 0x28,
	0x8e, 0x44, 0x53, 0xa5, 0xec, 0x52, 0xb1, 0x5a, 0x9d, 0x7a, 0xa4, 0xa1, 0x35, 0x73, 0x5c, 0x66,
	0x6c, 0x80, 0x63, 0x11, 0xe6, 0x6b, 0xb6, 0x3d, 0x66, 0x02, 0x9f, 0x47, 0x18, 0x84, 0x64, 0x0d,
	0x66, 0x2c, 0xdf, 0xb9, 0xc3, 0x47, 0x2a, 0x56, 0x6f, 0x4e, 0xe8, 0xc8, 0x4d, 0x98, 0x9c, 0x41,
	0x1e, 0x70, 0x66, 0x95, 0xbb, 0xfe, 0xab, 0xf5, 0x71, 0x1a, 0x21, 0x9c, 0xbe, 0x5a, 0xca, 0xf2,
	0xa5, 0xf1, 0x73, 0xf9, 0x2e, 0x5c, 0xd5, 0xa3, 0xd0, 0xa7, 0xa9, 0x13, 0xab, 0x67, 0x62, 0xba,
	0x49, 0x13, 0x03, 0x8f, 0xb9, 0x01, 0x92, 0x12, 0xe4, 0x7d, 0x0c, 0xa2, 0x5e, 0x78, 0x36, 0x6c,
	0xc1, 0xfc, 0x5e, 0x96, 0xdf, 0x88, 0xb0, 0x60, 0x62, 0x9f, 0xed, 0xe3, 0xbf, 0xed, 0xef, 0xf6,
	0x7b, 0x11, 0x60, 0x04, 0x24, 0xd7, 0xa1, 0x60, 0x34, 0x0d, 0xa3, 0x69, 0xb6, 0xf5, 0x96, 0x2c,
	0x28, 0x73, 0x71, 0xa2, 0x15, 0x0c, 0xf4, 0x3c, 0xf4, 0x99, 0x4b, 0xf9, 0x6b, 0x4b, 0x6f, 0xd6,
	0x0c, 0x63, 0xab, 0x29, 0x8b, 0xe7, 0xaf, 0xd4, 0x45, 0xcb, 0xf3, 0x7a, 0x48, 0x64, 0xc8, 0x6e,
	0xd6, 0xb6, 0xe5, 0x8c, 0x92, 0x8f, 0x13, 0x2d, 0xbb, 0x69, 0xf5, 0xc9, 0x3c, 0x48, 0xf5, 0x5a,
	0xa3, 0xad, 0xcb, 0x59, 0xa5, 0x10, 0x27, 0x9a, 0x54, 0xb7, 0xba, 0xcc, 0x4d, 0xb7, 0xd6, 0xd8,
	0x6c, 0x35, 0x1e, 0x36, 0x75, 0x79, 0x46, 0x29, 0xc6, 0x89, 0x96, 0x6f, 0xec, 0xd2, 0xee, 0x1e,
	0xba, 0xe9, 0x80, 0x5b, 0xb5, 0xed, 0xba, 0x2c, 0x29, 0xb3, 0x71, 0xa2, 0xcd, 0x6c, 0x59, 0xfd,
	0x0e, 0x59, 0x80, 0xbc, 0xde, 0x7e, 0xb2, 0xd1, 0x6e, 0xaf, 0xcb, 0x39, 0x05, 0xe2, 0x44, 0xcb,
	0xe9, 0x6c, 0x83, 0x31, 0x5b, 0xf9, 0xef, 0xf0, 0x58, 0x15, 0xde, 0x9e, 0xa8, 0xc2, 0xbb, 0x13,
	0x55, 0xa8, 0x7e, 0xc9, 0x00, 0x31, 0x46, 0xe6, 0x77, 0xd0, 0xdf, 0xa7, 0x5d, 0x24, 0xcf, 0x60,
	0xee, 0x87, 0x8c, 0x91, 0xd5, 0x09, 0x4b, 0xfb, 0x5d, 0x22, 0x95, 0xa9, 0xbe, 0x11, 0x09, 0xe0,
	0xd2, 0x2f, 0x51, 0x21, 0x53, 0x51, 0x95, 0xb5, 0x09, 0xa8, 0x8b, 0x23, 0xe8, 0x81, 0xfc, 0x73,
	0xce, 0xc8, 0xbd, 0x09, 0x6a, 0x17, 0x04, 0x73, 0x3a, 0x9b, 0xf5, 0x6b, 0x1f, 0x06, 0xaa, 0xf8,
	0x71, 0xa0, 0x8a, 0x9f, 0x07, 0xaa, 0xf0, 0xea, 0x54, 0x15, 0x8e, 0x4e, 0x55, 0xf1, 0xb1, 0xe4,
	0xb0, 0xbe, 0x65, 0x77, 0x72, 0xfc, 0x2f, 0xb7, 0xfa, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x72, 0x3b,
	0x3a, 0xec, 0x84, 0x05, 0x00, 0x00,
}
