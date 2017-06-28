// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/pkg/worker/worker_service.proto

/*
	Package worker is a generated protocol buffer package.

	It is generated from these files:
		server/pkg/worker/worker_service.proto

	It has these top-level messages:
		Input
		ProcessRequest
		ProcessResponse
		CancelRequest
		CancelResponse
*/
package worker

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"
import pps "github.com/pachyderm/pachyderm/src/client/pps"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf "github.com/gogo/protobuf/types"

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

type Input struct {
	FileInfo     *pfs.FileInfo `protobuf:"bytes,1,opt,name=file_info,json=fileInfo" json:"file_info,omitempty"`
	ParentCommit *pfs.Commit   `protobuf:"bytes,5,opt,name=parent_commit,json=parentCommit" json:"parent_commit,omitempty"`
	Name         string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lazy         bool          `protobuf:"varint,3,opt,name=lazy,proto3" json:"lazy,omitempty"`
	Branch       string        `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (m *Input) Reset()                    { *m = Input{} }
func (m *Input) String() string            { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()               {}
func (*Input) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{0} }

func (m *Input) GetFileInfo() *pfs.FileInfo {
	if m != nil {
		return m.FileInfo
	}
	return nil
}

func (m *Input) GetParentCommit() *pfs.Commit {
	if m != nil {
		return m.ParentCommit
	}
	return nil
}

func (m *Input) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Input) GetLazy() bool {
	if m != nil {
		return m.Lazy
	}
	return false
}

func (m *Input) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

type ProcessRequest struct {
	// ID of the job for which we're processing 'data'. This is attached to logs
	// generated while processing 'data', so that they can be searched.
	JobID string `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// The datum to process
	Data []*Input `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	// The tag corresponding to the previous parent's run of this datum, used for
	// incremental jobs, may be nil.
	ParentOutput *pfs.Tag `protobuf:"bytes,3,opt,name=parent_output,json=parentOutput" json:"parent_output,omitempty"`
}

func (m *ProcessRequest) Reset()                    { *m = ProcessRequest{} }
func (m *ProcessRequest) String() string            { return proto.CompactTextString(m) }
func (*ProcessRequest) ProtoMessage()               {}
func (*ProcessRequest) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{1} }

func (m *ProcessRequest) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *ProcessRequest) GetData() []*Input {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ProcessRequest) GetParentOutput() *pfs.Tag {
	if m != nil {
		return m.ParentOutput
	}
	return nil
}

// ProcessResponse contains a tag, only if the processing was successful.
type ProcessResponse struct {
	Tag      *pfs.Tag `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	StatsTag *pfs.Tag `protobuf:"bytes,3,opt,name=stats_tag,json=statsTag" json:"stats_tag,omitempty"`
	// If true, the user program has errored
	Failed  bool              `protobuf:"varint,2,opt,name=failed,proto3" json:"failed,omitempty"`
	Skipped bool              `protobuf:"varint,5,opt,name=skipped,proto3" json:"skipped,omitempty"`
	Stats   *pps.ProcessStats `protobuf:"bytes,4,opt,name=stats" json:"stats,omitempty"`
}

func (m *ProcessResponse) Reset()                    { *m = ProcessResponse{} }
func (m *ProcessResponse) String() string            { return proto.CompactTextString(m) }
func (*ProcessResponse) ProtoMessage()               {}
func (*ProcessResponse) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{2} }

func (m *ProcessResponse) GetTag() *pfs.Tag {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *ProcessResponse) GetStatsTag() *pfs.Tag {
	if m != nil {
		return m.StatsTag
	}
	return nil
}

func (m *ProcessResponse) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

func (m *ProcessResponse) GetSkipped() bool {
	if m != nil {
		return m.Skipped
	}
	return false
}

func (m *ProcessResponse) GetStats() *pps.ProcessStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type CancelRequest struct {
	JobID       string   `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	DataFilters []string `protobuf:"bytes,1,rep,name=data_filters,json=dataFilters" json:"data_filters,omitempty"`
}

func (m *CancelRequest) Reset()                    { *m = CancelRequest{} }
func (m *CancelRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelRequest) ProtoMessage()               {}
func (*CancelRequest) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{3} }

func (m *CancelRequest) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *CancelRequest) GetDataFilters() []string {
	if m != nil {
		return m.DataFilters
	}
	return nil
}

type CancelResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *CancelResponse) Reset()                    { *m = CancelResponse{} }
func (m *CancelResponse) String() string            { return proto.CompactTextString(m) }
func (*CancelResponse) ProtoMessage()               {}
func (*CancelResponse) Descriptor() ([]byte, []int) { return fileDescriptorWorkerService, []int{4} }

func (m *CancelResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Input)(nil), "worker.Input")
	proto.RegisterType((*ProcessRequest)(nil), "worker.ProcessRequest")
	proto.RegisterType((*ProcessResponse)(nil), "worker.ProcessResponse")
	proto.RegisterType((*CancelRequest)(nil), "worker.CancelRequest")
	proto.RegisterType((*CancelResponse)(nil), "worker.CancelResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Worker service

type WorkerClient interface {
	Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
	Status(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*pps.WorkerStatus, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := grpc.Invoke(ctx, "/worker.Worker/Process", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Status(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*pps.WorkerStatus, error) {
	out := new(pps.WorkerStatus)
	err := grpc.Invoke(ctx, "/worker.Worker/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := grpc.Invoke(ctx, "/worker.Worker/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Worker service

type WorkerServer interface {
	Process(context.Context, *ProcessRequest) (*ProcessResponse, error)
	Status(context.Context, *google_protobuf.Empty) (*pps.WorkerStatus, error)
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.Worker/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Process(ctx, req.(*ProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.Worker/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Status(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.Worker/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "worker.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _Worker_Process_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Worker_Status_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Worker_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/pkg/worker/worker_service.proto",
}

func (m *Input) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Input) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.FileInfo != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(m.FileInfo.Size()))
		n1, err := m.FileInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Lazy {
		dAtA[i] = 0x18
		i++
		if m.Lazy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Branch) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(len(m.Branch)))
		i += copy(dAtA[i:], m.Branch)
	}
	if m.ParentCommit != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(m.ParentCommit.Size()))
		n2, err := m.ParentCommit.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *ProcessRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, msg := range m.Data {
			dAtA[i] = 0xa
			i++
			i = encodeVarintWorkerService(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.JobID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(len(m.JobID)))
		i += copy(dAtA[i:], m.JobID)
	}
	if m.ParentOutput != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(m.ParentOutput.Size()))
		n3, err := m.ParentOutput.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *ProcessResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Tag != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(m.Tag.Size()))
		n4, err := m.Tag.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Failed {
		dAtA[i] = 0x10
		i++
		if m.Failed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.StatsTag != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(m.StatsTag.Size()))
		n5, err := m.StatsTag.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Stats != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(m.Stats.Size()))
		n6, err := m.Stats.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.Skipped {
		dAtA[i] = 0x28
		i++
		if m.Skipped {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *CancelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DataFilters) > 0 {
		for _, s := range m.DataFilters {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.JobID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWorkerService(dAtA, i, uint64(len(m.JobID)))
		i += copy(dAtA[i:], m.JobID)
	}
	return i, nil
}

func (m *CancelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelResponse) MarshalTo(dAtA []byte) (int, error) {
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

func encodeFixed64WorkerService(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32WorkerService(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintWorkerService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Input) Size() (n int) {
	var l int
	_ = l
	if m.FileInfo != nil {
		l = m.FileInfo.Size()
		n += 1 + l + sovWorkerService(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWorkerService(uint64(l))
	}
	if m.Lazy {
		n += 2
	}
	l = len(m.Branch)
	if l > 0 {
		n += 1 + l + sovWorkerService(uint64(l))
	}
	if m.ParentCommit != nil {
		l = m.ParentCommit.Size()
		n += 1 + l + sovWorkerService(uint64(l))
	}
	return n
}

func (m *ProcessRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovWorkerService(uint64(l))
		}
	}
	l = len(m.JobID)
	if l > 0 {
		n += 1 + l + sovWorkerService(uint64(l))
	}
	if m.ParentOutput != nil {
		l = m.ParentOutput.Size()
		n += 1 + l + sovWorkerService(uint64(l))
	}
	return n
}

func (m *ProcessResponse) Size() (n int) {
	var l int
	_ = l
	if m.Tag != nil {
		l = m.Tag.Size()
		n += 1 + l + sovWorkerService(uint64(l))
	}
	if m.Failed {
		n += 2
	}
	if m.StatsTag != nil {
		l = m.StatsTag.Size()
		n += 1 + l + sovWorkerService(uint64(l))
	}
	if m.Stats != nil {
		l = m.Stats.Size()
		n += 1 + l + sovWorkerService(uint64(l))
	}
	if m.Skipped {
		n += 2
	}
	return n
}

func (m *CancelRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.DataFilters) > 0 {
		for _, s := range m.DataFilters {
			l = len(s)
			n += 1 + l + sovWorkerService(uint64(l))
		}
	}
	l = len(m.JobID)
	if l > 0 {
		n += 1 + l + sovWorkerService(uint64(l))
	}
	return n
}

func (m *CancelResponse) Size() (n int) {
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovWorkerService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWorkerService(x uint64) (n int) {
	return sovWorkerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Input) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkerService
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
			return fmt.Errorf("proto: Input: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Input: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FileInfo == nil {
				m.FileInfo = &pfs.FileInfo{}
			}
			if err := m.FileInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lazy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
			m.Lazy = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Branch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Branch = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ParentCommit == nil {
				m.ParentCommit = &pfs.Commit{}
			}
			if err := m.ParentCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkerService
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
func (m *ProcessRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkerService
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
			return fmt.Errorf("proto: ProcessRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &Input{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentOutput", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ParentOutput == nil {
				m.ParentOutput = &pfs.Tag{}
			}
			if err := m.ParentOutput.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkerService
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
func (m *ProcessResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkerService
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
			return fmt.Errorf("proto: ProcessResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tag == nil {
				m.Tag = &pfs.Tag{}
			}
			if err := m.Tag.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
			m.Failed = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatsTag", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StatsTag == nil {
				m.StatsTag = &pfs.Tag{}
			}
			if err := m.StatsTag.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stats == nil {
				m.Stats = &pps.ProcessStats{}
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Skipped", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
			m.Skipped = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipWorkerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkerService
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
func (m *CancelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkerService
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
			return fmt.Errorf("proto: CancelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataFilters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataFilters = append(m.DataFilters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
				return ErrInvalidLengthWorkerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkerService
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
func (m *CancelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkerService
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
			return fmt.Errorf("proto: CancelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkerService
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
			skippy, err := skipWorkerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkerService
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
func skipWorkerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkerService
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
					return 0, ErrIntOverflowWorkerService
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
					return 0, ErrIntOverflowWorkerService
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
				return 0, ErrInvalidLengthWorkerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWorkerService
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
				next, err := skipWorkerService(dAtA[start:])
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
	ErrInvalidLengthWorkerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkerService   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("server/pkg/worker/worker_service.proto", fileDescriptorWorkerService) }

var fileDescriptorWorkerService = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0x9e, 0x59, 0x9b, 0xa5, 0xee, 0x3a, 0xc0, 0x82, 0x12, 0x15, 0xa9, 0x94, 0x48, 0x40, 0x35,
	0x89, 0x04, 0x15, 0x71, 0x40, 0xe2, 0xb4, 0xc1, 0xa4, 0x72, 0x01, 0x99, 0x4a, 0x1c, 0x23, 0x27,
	0x75, 0x42, 0xd6, 0x34, 0x36, 0xb1, 0x03, 0x1a, 0x67, 0x1e, 0x82, 0x17, 0xe0, 0xcc, 0x95, 0x47,
	0xe0, 0xc8, 0x13, 0x20, 0x54, 0x5e, 0x04, 0xf9, 0x77, 0xb2, 0x6a, 0xe3, 0xc0, 0x21, 0xca, 0xff,
	0x7f, 0x9f, 0xed, 0xef, 0xf3, 0xf7, 0x1b, 0xdf, 0x57, 0xbc, 0xfa, 0xc0, 0xab, 0x50, 0xae, 0xb2,
	0xf0, 0xa3, 0xa8, 0x56, 0xbc, 0x6a, 0x7e, 0x91, 0x21, 0xf2, 0x84, 0x07, 0xb2, 0x12, 0x5a, 0x10,
	0xc7, 0xa2, 0xa3, 0x1b, 0x49, 0x91, 0xf3, 0x52, 0x87, 0x32, 0x55, 0xe6, 0xb3, 0xec, 0x16, 0x95,
	0xca, 0x7c, 0x2d, 0x9a, 0x89, 0x4c, 0x40, 0x19, 0x9a, 0xaa, 0x41, 0x6f, 0x67, 0x42, 0x64, 0x05,
	0x0f, 0xa1, 0x8b, 0xeb, 0x34, 0xe4, 0x6b, 0xa9, 0xcf, 0x2c, 0xe9, 0x7f, 0x45, 0xb8, 0x3b, 0x2f,
	0x65, 0xad, 0xc9, 0x21, 0xee, 0xa5, 0x79, 0xc1, 0xa3, 0xbc, 0x4c, 0x85, 0x87, 0x26, 0x68, 0xda,
	0x9f, 0x0d, 0x02, 0xa3, 0x78, 0x92, 0x17, 0x7c, 0x5e, 0xa6, 0x82, 0xba, 0x69, 0x53, 0x11, 0x82,
	0x3b, 0x25, 0x5b, 0x73, 0xef, 0xca, 0x04, 0x4d, 0x7b, 0x14, 0x6a, 0x83, 0x15, 0xec, 0xd3, 0x99,
	0xb7, 0x3b, 0x41, 0x53, 0x97, 0x42, 0x4d, 0x86, 0xd8, 0x89, 0x2b, 0x56, 0x26, 0xef, 0xbc, 0x0e,
	0xac, 0x6c, 0x3a, 0xf2, 0x08, 0x0f, 0x24, 0xab, 0x78, 0xa9, 0xa3, 0x44, 0xac, 0xd7, 0xb9, 0xf6,
	0xba, 0xa0, 0xd7, 0x07, 0xbd, 0x63, 0x80, 0xe8, 0xbe, 0x5d, 0x61, 0x3b, 0xff, 0x33, 0xc2, 0x07,
	0xaf, 0x2b, 0x91, 0x70, 0xa5, 0x28, 0x7f, 0x5f, 0x73, 0xa5, 0xc9, 0x5d, 0xdc, 0x59, 0x32, 0xcd,
	0x3c, 0x34, 0xd9, 0x05, 0xaf, 0x36, 0xb0, 0x00, 0x6e, 0x43, 0x81, 0x22, 0x13, 0xec, 0x9c, 0x8a,
	0x38, 0xca, 0x97, 0xd6, 0xe9, 0x51, 0x6f, 0xf3, 0xeb, 0x4e, 0xf7, 0xa5, 0x88, 0xe7, 0xcf, 0x69,
	0xf7, 0x54, 0xc4, 0xf3, 0x25, 0x79, 0x78, 0xee, 0x44, 0xd4, 0x5a, 0xd6, 0x1a, 0xec, 0xf7, 0x67,
	0x2e, 0x38, 0x59, 0xb0, 0xac, 0xb5, 0xf1, 0x0a, 0x58, 0xff, 0x1b, 0xc2, 0x57, 0xcf, 0x6d, 0x28,
	0x29, 0x4a, 0xc5, 0xc9, 0x08, 0xef, 0x6a, 0x96, 0x35, 0x91, 0x6d, 0x37, 0x1a, 0xd0, 0x04, 0x90,
	0xb2, 0xbc, 0xe0, 0xd6, 0x80, 0x4b, 0x9b, 0x8e, 0xdc, 0xc3, 0x3d, 0xa5, 0x99, 0x56, 0x91, 0xd9,
	0x79, 0x59, 0xd2, 0x05, 0x6a, 0xc1, 0x32, 0xf2, 0x00, 0x77, 0xa1, 0x86, 0xf8, 0xfa, 0xb3, 0xeb,
	0x81, 0x99, 0x75, 0xa3, 0xff, 0xc6, 0x10, 0xd4, 0xf2, 0xc4, 0xc3, 0x7b, 0x6a, 0x95, 0x4b, 0xc9,
	0x97, 0x10, 0xa5, 0x4b, 0xdb, 0xd6, 0x5f, 0xe0, 0xc1, 0x31, 0x2b, 0x13, 0x5e, 0x6c, 0x63, 0xdb,
	0x37, 0xd9, 0x44, 0x69, 0x5e, 0x68, 0x5e, 0x29, 0x88, 0xaf, 0x47, 0xfb, 0x06, 0x3b, 0xb1, 0xd0,
	0xff, 0x63, 0xf3, 0x0f, 0xf1, 0x41, 0x7b, 0x6a, 0x93, 0x82, 0x71, 0x50, 0x27, 0xc6, 0x18, 0x24,
	0x61, 0x1c, 0xd8, 0x76, 0xf6, 0x1d, 0x61, 0xe7, 0x2d, 0xcc, 0x86, 0x3c, 0xc3, 0x7b, 0x8d, 0x7b,
	0x32, 0x6c, 0xe7, 0x75, 0x71, 0xaa, 0xa3, 0x5b, 0xff, 0xe0, 0x56, 0xc0, 0xdf, 0x21, 0x4f, 0xb0,
	0x63, 0x2e, 0x5d, 0x9b, 0xcd, 0xf6, 0x4d, 0x07, 0xed, 0x9b, 0x0e, 0x5e, 0x98, 0x37, 0x3d, 0xb2,
	0x01, 0x59, 0x31, 0xbb, 0xd4, 0xdf, 0x21, 0x4f, 0xb1, 0x63, 0xbd, 0x92, 0x9b, 0xed, 0xd9, 0x17,
	0x12, 0x19, 0x0d, 0x2f, 0xc3, 0xad, 0xe2, 0xd1, 0xb5, 0x1f, 0x9b, 0x31, 0xfa, 0xb9, 0x19, 0xa3,
	0xdf, 0x9b, 0x31, 0xfa, 0xf2, 0x67, 0xbc, 0x13, 0x3b, 0xa0, 0xf8, 0xf8, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x76, 0x3e, 0x72, 0x51, 0xc7, 0x03, 0x00, 0x00,
}
