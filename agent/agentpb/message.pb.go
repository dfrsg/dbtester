// Code generated by protoc-gen-gogo.
// source: agent/agentpb/message.proto
// DO NOT EDIT!

/*
	Package agentpb is a generated protocol buffer package.

	It is generated from these files:
		agent/agentpb/message.proto

	It has these top-level messages:
		Request
		Response
*/
package agentpb

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

type Request_Operation int32

const (
	Request_Start     Request_Operation = 0
	Request_Stop      Request_Operation = 1
	Request_Heartbeat Request_Operation = 2
)

var Request_Operation_name = map[int32]string{
	0: "Start",
	1: "Stop",
	2: "Heartbeat",
}
var Request_Operation_value = map[string]int32{
	"Start":     0,
	"Stop":      1,
	"Heartbeat": 2,
}

func (x Request_Operation) String() string {
	return proto.EnumName(Request_Operation_name, int32(x))
}
func (Request_Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0, 0} }

type Request_Database int32

const (
	Request_etcdv2    Request_Database = 0
	Request_etcdv3    Request_Database = 1
	Request_ZooKeeper Request_Database = 2
	Request_Consul    Request_Database = 3
	Request_zetcd     Request_Database = 4
	Request_cetcd     Request_Database = 5
)

var Request_Database_name = map[int32]string{
	0: "etcdv2",
	1: "etcdv3",
	2: "ZooKeeper",
	3: "Consul",
	4: "zetcd",
	5: "cetcd",
}
var Request_Database_value = map[string]int32{
	"etcdv2":    0,
	"etcdv3":    1,
	"ZooKeeper": 2,
	"Consul":    3,
	"zetcd":     4,
	"cetcd":     5,
}

func (x Request_Database) String() string {
	return proto.EnumName(Request_Database_name, int32(x))
}
func (Request_Database) EnumDescriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0, 1} }

type Request struct {
	Operation    Request_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=agentpb.Request_Operation" json:"operation,omitempty"`
	Database     Request_Database  `protobuf:"varint,2,opt,name=database,proto3,enum=agentpb.Request_Database" json:"database,omitempty"`
	PeerIPString string            `protobuf:"bytes,3,opt,name=peerIPString,proto3" json:"peerIPString,omitempty"`
	// ServerIPIndex is the index in peerIPs that points to the
	// corresponding remote IP.
	ServerIndex uint32 `protobuf:"varint,4,opt,name=serverIndex,proto3" json:"serverIndex,omitempty"`
	// TestName prefixes all logs to be generated in agent.
	TestName string `protobuf:"bytes,5,opt,name=testName,proto3" json:"testName,omitempty"`
	// GoogleCloudProjectName is the project name to use
	// to upload logs.
	GoogleCloudProjectName string `protobuf:"bytes,6,opt,name=googleCloudProjectName,proto3" json:"googleCloudProjectName,omitempty"`
	// GoogleCloudStorageKey is the key to be used to upload
	// data and logs to Google Cloud Storage and others.
	GoogleCloudStorageKey string `protobuf:"bytes,7,opt,name=googleCloudStorageKey,proto3" json:"googleCloudStorageKey,omitempty"`
	// GoogleCloudStorageBucketName is the bucket name to store all data and logs.
	GoogleCloudStorageBucketName string `protobuf:"bytes,8,opt,name=googleCloudStorageBucketName,proto3" json:"googleCloudStorageBucketName,omitempty"`
	// GoogleCloudStorageSubDirectory is the sub-directory name to store data.
	GoogleCloudStorageSubDirectory string `protobuf:"bytes,9,opt,name=googleCloudStorageSubDirectory,proto3" json:"googleCloudStorageSubDirectory,omitempty"`
	// ZookeeperMyID is myid that needs to be stored as a file in the remote machine.
	ZookeeperMyID uint32 `protobuf:"varint,10,opt,name=zookeeperMyID,proto3" json:"zookeeperMyID,omitempty"`
	// EtcdSnapCount is 100,000 by default.
	EtcdSnapCount int64 `protobuf:"varint,11,opt,name=etcdSnapCount,proto3" json:"etcdSnapCount,omitempty"`
	// ZookeeperSnapCount is 100,000 by default.
	ZookeeperSnapCount int64 `protobuf:"varint,12,opt,name=zookeeperSnapCount,proto3" json:"zookeeperSnapCount,omitempty"`
	// ZookeeperMaxClientCnxns limits the number of concurrent connections
	// (at the socket level) that a single client, identified by IP address.
	ZookeeperMaxClientCnxns int64 `protobuf:"varint,13,opt,name=zookeeperMaxClientCnxns,proto3" json:"zookeeperMaxClientCnxns,omitempty"`
	// ClientNum is current number of clients.
	ClientNum int64 `protobuf:"varint,14,opt,name=clientNum,proto3" json:"clientNum,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

type Response struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "agentpb.Request")
	proto.RegisterType((*Response)(nil), "agentpb.Response")
	proto.RegisterEnum("agentpb.Request_Operation", Request_Operation_name, Request_Operation_value)
	proto.RegisterEnum("agentpb.Request_Database", Request_Database_name, Request_Database_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Transporter service

type TransporterClient interface {
	Transfer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type transporterClient struct {
	cc *grpc.ClientConn
}

func NewTransporterClient(cc *grpc.ClientConn) TransporterClient {
	return &transporterClient{cc}
}

func (c *transporterClient) Transfer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/agentpb.Transporter/Transfer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Transporter service

type TransporterServer interface {
	Transfer(context.Context, *Request) (*Response, error)
}

func RegisterTransporterServer(s *grpc.Server, srv TransporterServer) {
	s.RegisterService(&_Transporter_serviceDesc, srv)
}

func _Transporter_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransporterServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Transporter/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransporterServer).Transfer(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agentpb.Transporter",
	HandlerType: (*TransporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _Transporter_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent/agentpb/message.proto",
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
	if m.Operation != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Operation))
	}
	if m.Database != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Database))
	}
	if len(m.PeerIPString) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.PeerIPString)))
		i += copy(dAtA[i:], m.PeerIPString)
	}
	if m.ServerIndex != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ServerIndex))
	}
	if len(m.TestName) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.TestName)))
		i += copy(dAtA[i:], m.TestName)
	}
	if len(m.GoogleCloudProjectName) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.GoogleCloudProjectName)))
		i += copy(dAtA[i:], m.GoogleCloudProjectName)
	}
	if len(m.GoogleCloudStorageKey) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.GoogleCloudStorageKey)))
		i += copy(dAtA[i:], m.GoogleCloudStorageKey)
	}
	if len(m.GoogleCloudStorageBucketName) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.GoogleCloudStorageBucketName)))
		i += copy(dAtA[i:], m.GoogleCloudStorageBucketName)
	}
	if len(m.GoogleCloudStorageSubDirectory) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.GoogleCloudStorageSubDirectory)))
		i += copy(dAtA[i:], m.GoogleCloudStorageSubDirectory)
	}
	if m.ZookeeperMyID != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ZookeeperMyID))
	}
	if m.EtcdSnapCount != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.EtcdSnapCount))
	}
	if m.ZookeeperSnapCount != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ZookeeperSnapCount))
	}
	if m.ZookeeperMaxClientCnxns != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ZookeeperMaxClientCnxns))
	}
	if m.ClientNum != 0 {
		dAtA[i] = 0x70
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ClientNum))
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
	if m.Success {
		dAtA[i] = 0x8
		i++
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Message(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Message(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
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
	if m.Operation != 0 {
		n += 1 + sovMessage(uint64(m.Operation))
	}
	if m.Database != 0 {
		n += 1 + sovMessage(uint64(m.Database))
	}
	l = len(m.PeerIPString)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ServerIndex != 0 {
		n += 1 + sovMessage(uint64(m.ServerIndex))
	}
	l = len(m.TestName)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.GoogleCloudProjectName)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.GoogleCloudStorageKey)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.GoogleCloudStorageBucketName)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.GoogleCloudStorageSubDirectory)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ZookeeperMyID != 0 {
		n += 1 + sovMessage(uint64(m.ZookeeperMyID))
	}
	if m.EtcdSnapCount != 0 {
		n += 1 + sovMessage(uint64(m.EtcdSnapCount))
	}
	if m.ZookeeperSnapCount != 0 {
		n += 1 + sovMessage(uint64(m.ZookeeperSnapCount))
	}
	if m.ZookeeperMaxClientCnxns != 0 {
		n += 1 + sovMessage(uint64(m.ZookeeperMaxClientCnxns))
	}
	if m.ClientNum != 0 {
		n += 1 + sovMessage(uint64(m.ClientNum))
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			m.Operation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operation |= (Request_Operation(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Database", wireType)
			}
			m.Database = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Database |= (Request_Database(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerIPString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerIPString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerIndex", wireType)
			}
			m.ServerIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerIndex |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TestName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleCloudProjectName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleCloudProjectName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleCloudStorageKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleCloudStorageKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleCloudStorageBucketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleCloudStorageBucketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleCloudStorageSubDirectory", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleCloudStorageSubDirectory = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZookeeperMyID", wireType)
			}
			m.ZookeeperMyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ZookeeperMyID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdSnapCount", wireType)
			}
			m.EtcdSnapCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EtcdSnapCount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZookeeperSnapCount", wireType)
			}
			m.ZookeeperSnapCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ZookeeperSnapCount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZookeeperMaxClientCnxns", wireType)
			}
			m.ZookeeperMaxClientCnxns = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ZookeeperMaxClientCnxns |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientNum", wireType)
			}
			m.ClientNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientNum |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
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
				next, err := skipMessage(dAtA[start:])
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
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("agent/agentpb/message.proto", fileDescriptorMessage) }

var fileDescriptorMessage = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x93, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xe3, 0x5e, 0x12, 0xfb, 0xb4, 0xa9, 0xcc, 0x88, 0xcb, 0x10, 0x2a, 0x2b, 0xb2, 0xba,
	0xc8, 0x06, 0x47, 0xb4, 0x80, 0xba, 0x44, 0x49, 0x84, 0x88, 0x2a, 0x4a, 0x65, 0xb3, 0x62, 0x37,
	0x76, 0x4e, 0x4d, 0x68, 0x32, 0x63, 0x66, 0xc6, 0x55, 0xd2, 0x27, 0xe1, 0x31, 0x78, 0x8c, 0x2e,
	0x79, 0x04, 0x08, 0x2f, 0x82, 0x3c, 0xb9, 0x93, 0x16, 0x36, 0xd6, 0x39, 0xff, 0xff, 0xfd, 0x67,
	0x6c, 0x6b, 0x0e, 0x3c, 0x63, 0x29, 0x72, 0xdd, 0x34, 0xcf, 0x2c, 0x6e, 0x0e, 0x51, 0x29, 0x96,
	0x62, 0x90, 0x49, 0xa1, 0x05, 0xa9, 0xcc, 0xe4, 0xda, 0xf3, 0xb4, 0xaf, 0x3f, 0xe7, 0x71, 0x90,
	0x88, 0x61, 0x33, 0x15, 0xa9, 0x68, 0x1a, 0x3f, 0xce, 0x2f, 0x4d, 0x67, 0x1a, 0x53, 0x4d, 0x73,
	0xfe, 0xf7, 0x32, 0x54, 0x42, 0xfc, 0x9a, 0xa3, 0xd2, 0xe4, 0x14, 0x1c, 0x91, 0xa1, 0x64, 0xba,
	0x2f, 0x38, 0xb5, 0xea, 0x56, 0xe3, 0xe0, 0xb8, 0x16, 0xcc, 0xe6, 0x06, 0x33, 0x28, 0xf8, 0x30,
	0x27, 0xc2, 0x25, 0x4c, 0x5e, 0x81, 0xdd, 0x63, 0x9a, 0xc5, 0x4c, 0x21, 0xdd, 0x32, 0xc1, 0xa7,
	0x1b, 0xc1, 0xce, 0x0c, 0x08, 0x17, 0x28, 0xf1, 0x61, 0x3f, 0x43, 0x94, 0xdd, 0x8b, 0x48, 0xcb,
	0x3e, 0x4f, 0xe9, 0x76, 0xdd, 0x6a, 0x38, 0xe1, 0x9a, 0x46, 0xea, 0xb0, 0xa7, 0x50, 0x5e, 0xa3,
	0xec, 0xf2, 0x1e, 0x8e, 0xe8, 0x4e, 0xdd, 0x6a, 0x54, 0xc3, 0x55, 0x89, 0xd4, 0xc0, 0xd6, 0xa8,
	0xf4, 0x39, 0x1b, 0x22, 0xdd, 0x35, 0x13, 0x16, 0x3d, 0x79, 0x0d, 0x8f, 0x53, 0x21, 0xd2, 0x01,
	0xb6, 0x07, 0x22, 0xef, 0x5d, 0x48, 0xf1, 0x05, 0x93, 0x29, 0x59, 0x36, 0xe4, 0x3d, 0x2e, 0x79,
	0x09, 0x8f, 0x56, 0x9c, 0x48, 0x0b, 0xc9, 0x52, 0x3c, 0xc3, 0x31, 0xad, 0x98, 0xd8, 0xdd, 0x26,
	0x69, 0xc1, 0xe1, 0xa6, 0xd1, 0xca, 0x93, 0x2b, 0x9c, 0x9e, 0x69, 0x9b, 0xf0, 0x3f, 0x19, 0xf2,
	0x16, 0xbc, 0x4d, 0x3f, 0xca, 0xe3, 0x4e, 0x5f, 0x62, 0xa2, 0x85, 0x1c, 0x53, 0xc7, 0x4c, 0xf9,
	0x0f, 0x45, 0x8e, 0xa0, 0x7a, 0x23, 0xc4, 0x15, 0x62, 0x86, 0xf2, 0xfd, 0xb8, 0xdb, 0xa1, 0x60,
	0xfe, 0xdc, 0xba, 0x58, 0x50, 0xa8, 0x93, 0x5e, 0xc4, 0x59, 0xd6, 0x16, 0x39, 0xd7, 0x74, 0xaf,
	0x6e, 0x35, 0xb6, 0xc3, 0x75, 0x91, 0x04, 0x40, 0x16, 0xb1, 0x25, 0xba, 0x6f, 0xd0, 0x3b, 0x1c,
	0x72, 0x0a, 0x4f, 0x96, 0xc7, 0xb0, 0x51, 0x7b, 0xd0, 0x47, 0xae, 0xdb, 0x7c, 0xc4, 0x15, 0xad,
	0x9a, 0xd0, 0x7d, 0x36, 0x39, 0x04, 0x27, 0x31, 0xed, 0x79, 0x3e, 0xa4, 0x07, 0x86, 0x5d, 0x0a,
	0x7e, 0x13, 0x9c, 0xc5, 0xf5, 0x23, 0x0e, 0xec, 0x46, 0x9a, 0x49, 0xed, 0x96, 0x88, 0x0d, 0x3b,
	0x91, 0x16, 0x99, 0x6b, 0x91, 0x2a, 0x38, 0xef, 0x90, 0x49, 0x1d, 0x23, 0xd3, 0xee, 0x96, 0x1f,
	0x81, 0x3d, 0xbf, 0x76, 0x04, 0xa0, 0x5c, 0x7c, 0xd5, 0xf5, 0xb1, 0x5b, 0x5a, 0xd4, 0x27, 0xd3,
	0xc8, 0x27, 0x21, 0xce, 0xcc, 0xdb, 0xb8, 0x5b, 0x85, 0xd5, 0x16, 0x5c, 0xe5, 0x03, 0x77, 0xbb,
	0x38, 0xe2, 0xa6, 0xe0, 0xdc, 0x9d, 0xa2, 0x4c, 0x4c, 0xb9, 0xeb, 0x1f, 0x81, 0x1d, 0xa2, 0xca,
	0x04, 0x57, 0x48, 0x28, 0x54, 0x54, 0x9e, 0x24, 0xa8, 0x94, 0x59, 0x18, 0x3b, 0x9c, 0xb7, 0xc7,
	0x6f, 0x60, 0xef, 0xa3, 0x64, 0x5c, 0x65, 0x42, 0x6a, 0x94, 0xe4, 0x05, 0xd8, 0xa6, 0xbd, 0x44,
	0x49, 0xdc, 0xbf, 0x77, 0xa3, 0xf6, 0x60, 0x45, 0x99, 0x4e, 0xf6, 0x4b, 0xad, 0x87, 0xb7, 0xbf,
	0xbc, 0xd2, 0xed, 0xc4, 0xb3, 0x7e, 0x4c, 0x3c, 0xeb, 0xe7, 0xc4, 0xb3, 0xbe, 0xfd, 0xf6, 0x4a,
	0x71, 0xd9, 0xec, 0xed, 0xc9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0xd4, 0x32, 0x95, 0x0e,
	0x04, 0x00, 0x00,
}
