// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jobs/jobspb/schedule.proto

package jobspb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// WaitBehavior describes how to handle previously  started
// jobs that have not completed yet.
type ScheduleDetails_WaitBehavior int32

const (
	// Wait for the previous run to complete
	// before starting the next one.
	ScheduleDetails_WAIT ScheduleDetails_WaitBehavior = 0
	// Do not wait for the previous run to complete.
	ScheduleDetails_NO_WAIT ScheduleDetails_WaitBehavior = 1
	// If the previous run is still running, skip this run
	// and advance schedule to the next recurrence.
	ScheduleDetails_SKIP ScheduleDetails_WaitBehavior = 2
)

var ScheduleDetails_WaitBehavior_name = map[int32]string{
	0: "WAIT",
	1: "NO_WAIT",
	2: "SKIP",
}
var ScheduleDetails_WaitBehavior_value = map[string]int32{
	"WAIT":    0,
	"NO_WAIT": 1,
	"SKIP":    2,
}

func (x ScheduleDetails_WaitBehavior) String() string {
	return proto.EnumName(ScheduleDetails_WaitBehavior_name, int32(x))
}
func (ScheduleDetails_WaitBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_schedule_3c801e1f423ca12e, []int{0, 0}
}

// ErrorHandlingBehavior describes how to handle failed job runs.
type ScheduleDetails_ErrorHandlingBehavior int32

const (
	// By default, failed jobs will run again, based on their schedule.
	ScheduleDetails_RETRY_SCHED ScheduleDetails_ErrorHandlingBehavior = 0
	// Retry failed jobs soon.
	ScheduleDetails_RETRY_SOON ScheduleDetails_ErrorHandlingBehavior = 1
	// Stop running this schedule
	ScheduleDetails_PAUSE_SCHED ScheduleDetails_ErrorHandlingBehavior = 2
)

var ScheduleDetails_ErrorHandlingBehavior_name = map[int32]string{
	0: "RETRY_SCHED",
	1: "RETRY_SOON",
	2: "PAUSE_SCHED",
}
var ScheduleDetails_ErrorHandlingBehavior_value = map[string]int32{
	"RETRY_SCHED": 0,
	"RETRY_SOON":  1,
	"PAUSE_SCHED": 2,
}

func (x ScheduleDetails_ErrorHandlingBehavior) String() string {
	return proto.EnumName(ScheduleDetails_ErrorHandlingBehavior_name, int32(x))
}
func (ScheduleDetails_ErrorHandlingBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_schedule_3c801e1f423ca12e, []int{0, 1}
}

// ScheduleDetails describes how to schedule and execute the job.
type ScheduleDetails struct {
	// How to handle running jobs.
	Wait ScheduleDetails_WaitBehavior `protobuf:"varint,1,opt,name=wait,proto3,enum=cockroach.jobs.jobspb.ScheduleDetails_WaitBehavior" json:"wait,omitempty"`
	// How to handle failed jobs.
	OnError ScheduleDetails_ErrorHandlingBehavior `protobuf:"varint,2,opt,name=on_error,json=onError,proto3,enum=cockroach.jobs.jobspb.ScheduleDetails_ErrorHandlingBehavior" json:"on_error,omitempty"`
}

func (m *ScheduleDetails) Reset()         { *m = ScheduleDetails{} }
func (m *ScheduleDetails) String() string { return proto.CompactTextString(m) }
func (*ScheduleDetails) ProtoMessage()    {}
func (*ScheduleDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_3c801e1f423ca12e, []int{0}
}
func (m *ScheduleDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ScheduleDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleDetails.Merge(dst, src)
}
func (m *ScheduleDetails) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleDetails proto.InternalMessageInfo

// ExecutionArguments describes data needed to execute scheduled jobs.
type ExecutionArguments struct {
	Args *types.Any `protobuf:"bytes,1,opt,name=args,proto3" json:"args,omitempty"`
}

func (m *ExecutionArguments) Reset()         { *m = ExecutionArguments{} }
func (m *ExecutionArguments) String() string { return proto.CompactTextString(m) }
func (*ExecutionArguments) ProtoMessage()    {}
func (*ExecutionArguments) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_3c801e1f423ca12e, []int{1}
}
func (m *ExecutionArguments) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecutionArguments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ExecutionArguments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionArguments.Merge(dst, src)
}
func (m *ExecutionArguments) XXX_Size() int {
	return m.Size()
}
func (m *ExecutionArguments) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionArguments.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionArguments proto.InternalMessageInfo

// ScheduleChangeInfo describes the reasons for schedule changes.
type ScheduleChangeInfo struct {
	Changes []ScheduleChangeInfo_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes"`
}

func (m *ScheduleChangeInfo) Reset()         { *m = ScheduleChangeInfo{} }
func (m *ScheduleChangeInfo) String() string { return proto.CompactTextString(m) }
func (*ScheduleChangeInfo) ProtoMessage()    {}
func (*ScheduleChangeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_3c801e1f423ca12e, []int{2}
}
func (m *ScheduleChangeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleChangeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ScheduleChangeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleChangeInfo.Merge(dst, src)
}
func (m *ScheduleChangeInfo) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleChangeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleChangeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleChangeInfo proto.InternalMessageInfo

type ScheduleChangeInfo_Change struct {
	Time   int64  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (m *ScheduleChangeInfo_Change) Reset()         { *m = ScheduleChangeInfo_Change{} }
func (m *ScheduleChangeInfo_Change) String() string { return proto.CompactTextString(m) }
func (*ScheduleChangeInfo_Change) ProtoMessage()    {}
func (*ScheduleChangeInfo_Change) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_3c801e1f423ca12e, []int{2, 0}
}
func (m *ScheduleChangeInfo_Change) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleChangeInfo_Change) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ScheduleChangeInfo_Change) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleChangeInfo_Change.Merge(dst, src)
}
func (m *ScheduleChangeInfo_Change) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleChangeInfo_Change) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleChangeInfo_Change.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleChangeInfo_Change proto.InternalMessageInfo

// Message representing sql statement to execute.
type SqlStatementExecutionArg struct {
	Statement string `protobuf:"bytes,1,opt,name=statement,proto3" json:"statement,omitempty"`
}

func (m *SqlStatementExecutionArg) Reset()         { *m = SqlStatementExecutionArg{} }
func (m *SqlStatementExecutionArg) String() string { return proto.CompactTextString(m) }
func (*SqlStatementExecutionArg) ProtoMessage()    {}
func (*SqlStatementExecutionArg) Descriptor() ([]byte, []int) {
	return fileDescriptor_schedule_3c801e1f423ca12e, []int{3}
}
func (m *SqlStatementExecutionArg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SqlStatementExecutionArg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *SqlStatementExecutionArg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqlStatementExecutionArg.Merge(dst, src)
}
func (m *SqlStatementExecutionArg) XXX_Size() int {
	return m.Size()
}
func (m *SqlStatementExecutionArg) XXX_DiscardUnknown() {
	xxx_messageInfo_SqlStatementExecutionArg.DiscardUnknown(m)
}

var xxx_messageInfo_SqlStatementExecutionArg proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ScheduleDetails)(nil), "cockroach.jobs.jobspb.ScheduleDetails")
	proto.RegisterType((*ExecutionArguments)(nil), "cockroach.jobs.jobspb.ExecutionArguments")
	proto.RegisterType((*ScheduleChangeInfo)(nil), "cockroach.jobs.jobspb.ScheduleChangeInfo")
	proto.RegisterType((*ScheduleChangeInfo_Change)(nil), "cockroach.jobs.jobspb.ScheduleChangeInfo.Change")
	proto.RegisterType((*SqlStatementExecutionArg)(nil), "cockroach.jobs.jobspb.SqlStatementExecutionArg")
	proto.RegisterEnum("cockroach.jobs.jobspb.ScheduleDetails_WaitBehavior", ScheduleDetails_WaitBehavior_name, ScheduleDetails_WaitBehavior_value)
	proto.RegisterEnum("cockroach.jobs.jobspb.ScheduleDetails_ErrorHandlingBehavior", ScheduleDetails_ErrorHandlingBehavior_name, ScheduleDetails_ErrorHandlingBehavior_value)
}
func (m *ScheduleDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Wait != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(m.Wait))
	}
	if m.OnError != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(m.OnError))
	}
	return i, nil
}

func (m *ExecutionArguments) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecutionArguments) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Args != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(m.Args.Size()))
		n1, err := m.Args.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *ScheduleChangeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleChangeInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for _, msg := range m.Changes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSchedule(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ScheduleChangeInfo_Change) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleChangeInfo_Change) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Time != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(m.Time))
	}
	if len(m.Reason) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Reason)))
		i += copy(dAtA[i:], m.Reason)
	}
	return i, nil
}

func (m *SqlStatementExecutionArg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SqlStatementExecutionArg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Statement) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchedule(dAtA, i, uint64(len(m.Statement)))
		i += copy(dAtA[i:], m.Statement)
	}
	return i, nil
}

func encodeVarintSchedule(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ScheduleDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Wait != 0 {
		n += 1 + sovSchedule(uint64(m.Wait))
	}
	if m.OnError != 0 {
		n += 1 + sovSchedule(uint64(m.OnError))
	}
	return n
}

func (m *ExecutionArguments) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Args != nil {
		l = m.Args.Size()
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func (m *ScheduleChangeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for _, e := range m.Changes {
			l = e.Size()
			n += 1 + l + sovSchedule(uint64(l))
		}
	}
	return n
}

func (m *ScheduleChangeInfo_Change) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Time != 0 {
		n += 1 + sovSchedule(uint64(m.Time))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func (m *SqlStatementExecutionArg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Statement)
	if l > 0 {
		n += 1 + l + sovSchedule(uint64(l))
	}
	return n
}

func sovSchedule(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSchedule(x uint64) (n int) {
	return sovSchedule(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ScheduleDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: ScheduleDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wait", wireType)
			}
			m.Wait = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Wait |= (ScheduleDetails_WaitBehavior(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnError", wireType)
			}
			m.OnError = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OnError |= (ScheduleDetails_ErrorHandlingBehavior(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchedule
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
func (m *ExecutionArguments) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: ExecutionArguments: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecutionArguments: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Args == nil {
				m.Args = &types.Any{}
			}
			if err := m.Args.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchedule
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
func (m *ScheduleChangeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: ScheduleChangeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleChangeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Changes = append(m.Changes, ScheduleChangeInfo_Change{})
			if err := m.Changes[len(m.Changes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchedule
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
func (m *ScheduleChangeInfo_Change) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: Change: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Change: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchedule
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
func (m *SqlStatementExecutionArg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchedule
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
			return fmt.Errorf("proto: SqlStatementExecutionArg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SqlStatementExecutionArg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchedule
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
				return ErrInvalidLengthSchedule
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statement = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchedule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchedule
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
func skipSchedule(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchedule
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
					return 0, ErrIntOverflowSchedule
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
					return 0, ErrIntOverflowSchedule
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
				return 0, ErrInvalidLengthSchedule
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSchedule
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
				next, err := skipSchedule(dAtA[start:])
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
	ErrInvalidLengthSchedule = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchedule   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("jobs/jobspb/schedule.proto", fileDescriptor_schedule_3c801e1f423ca12e)
}

var fileDescriptor_schedule_3c801e1f423ca12e = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbd, 0xa9, 0x95, 0x34, 0x13, 0xd4, 0x46, 0xab, 0x16, 0x85, 0x08, 0x19, 0xe4, 0x53,
	0x4e, 0x6b, 0x94, 0x72, 0xe0, 0x80, 0x90, 0x92, 0xd6, 0xa2, 0x11, 0x52, 0x13, 0xd9, 0x45, 0x11,
	0x5c, 0xa2, 0x8d, 0xbb, 0x75, 0x0c, 0xee, 0x4e, 0x59, 0x6f, 0x80, 0xbe, 0x05, 0x0f, 0xc0, 0x3b,
	0xf0, 0x1a, 0x39, 0xf6, 0xd8, 0x13, 0x82, 0xe4, 0x45, 0x90, 0xd7, 0x31, 0xad, 0x50, 0x24, 0x7a,
	0xb1, 0xe6, 0xcf, 0x37, 0xbf, 0xf1, 0x7c, 0x0b, 0xed, 0x0f, 0x38, 0xcd, 0xbc, 0xfc, 0x73, 0x39,
	0xf5, 0xb2, 0x68, 0x26, 0xce, 0xe6, 0xa9, 0x60, 0x97, 0x0a, 0x35, 0xd2, 0xfd, 0x08, 0xa3, 0x8f,
	0x0a, 0x79, 0x34, 0x63, 0xb9, 0x80, 0x15, 0xaa, 0xf6, 0x5e, 0x8c, 0x31, 0x1a, 0x85, 0x97, 0x47,
	0x85, 0xb8, 0xfd, 0x28, 0x46, 0x8c, 0x53, 0xe1, 0x99, 0x6c, 0x3a, 0x3f, 0xf7, 0xb8, 0xbc, 0x2a,
	0x5a, 0xee, 0x8f, 0x0a, 0xec, 0x86, 0x6b, 0xf4, 0x91, 0xd0, 0x3c, 0x49, 0x33, 0xfa, 0x1a, 0xec,
	0x2f, 0x3c, 0xd1, 0x2d, 0xf2, 0x94, 0x74, 0x76, 0xba, 0x07, 0x6c, 0xe3, 0x2a, 0xf6, 0xcf, 0x14,
	0x1b, 0xf3, 0x44, 0xf7, 0xc5, 0x8c, 0x7f, 0x4e, 0x50, 0x05, 0x06, 0x40, 0xc7, 0xb0, 0x8d, 0x72,
	0x22, 0x94, 0x42, 0xd5, 0xaa, 0x18, 0xd8, 0xcb, 0x7b, 0xc2, 0xfc, 0x7c, 0xe6, 0x98, 0xcb, 0xb3,
	0x34, 0x91, 0xf1, 0x5f, 0x6a, 0x0d, 0xa5, 0x69, 0xb8, 0x1e, 0x3c, 0xb8, 0xbb, 0x8e, 0x6e, 0x83,
	0x3d, 0xee, 0x0d, 0x4e, 0x9b, 0x16, 0x6d, 0x40, 0xed, 0x64, 0x38, 0x31, 0x09, 0xc9, 0xcb, 0xe1,
	0x9b, 0xc1, 0xa8, 0x59, 0x71, 0x07, 0xb0, 0xbf, 0x11, 0x49, 0x77, 0xa1, 0x11, 0xf8, 0xa7, 0xc1,
	0xbb, 0x49, 0x78, 0x78, 0xec, 0x1f, 0x35, 0x2d, 0xba, 0x03, 0xb0, 0x2e, 0x0c, 0x87, 0x27, 0x4d,
	0x92, 0x0b, 0x46, 0xbd, 0xb7, 0xa1, 0xbf, 0x16, 0x54, 0xdc, 0x57, 0x40, 0xfd, 0xaf, 0x22, 0x9a,
	0xeb, 0x04, 0x65, 0x4f, 0xc5, 0xf3, 0x0b, 0x21, 0x75, 0x46, 0x3b, 0x60, 0x73, 0x15, 0x67, 0xc6,
	0xb3, 0x46, 0x77, 0x8f, 0x15, 0x8e, 0xb3, 0xd2, 0x71, 0xd6, 0x93, 0x57, 0x81, 0x51, 0xb8, 0xdf,
	0x09, 0xd0, 0xf2, 0xdc, 0xc3, 0x19, 0x97, 0xb1, 0x18, 0xc8, 0x73, 0xa4, 0x23, 0xa8, 0x45, 0x26,
	0xcb, 0x19, 0x5b, 0x9d, 0x46, 0xf7, 0xd9, 0x7f, 0xac, 0xba, 0x9d, 0x65, 0x45, 0xd8, 0xb7, 0x17,
	0x3f, 0x9f, 0x58, 0x41, 0x89, 0x69, 0x3f, 0x87, 0x6a, 0xd1, 0xa0, 0x14, 0x6c, 0x9d, 0x5c, 0x08,
	0xf3, 0x73, 0x5b, 0x81, 0x89, 0xe9, 0x43, 0xa8, 0x2a, 0xc1, 0x33, 0x94, 0xe6, 0x65, 0xea, 0xc1,
	0x3a, 0x73, 0x5f, 0x40, 0x2b, 0xfc, 0x94, 0x86, 0x9a, 0x6b, 0x91, 0x5f, 0x76, 0xf7, 0x54, 0xfa,
	0x18, 0xea, 0x59, 0xd9, 0x30, 0xb0, 0x7a, 0x70, 0x5b, 0xe8, 0x77, 0x16, 0xbf, 0x1d, 0x6b, 0xb1,
	0x74, 0xc8, 0xf5, 0xd2, 0x21, 0x37, 0x4b, 0x87, 0xfc, 0x5a, 0x3a, 0xe4, 0xdb, 0xca, 0xb1, 0xae,
	0x57, 0x8e, 0x75, 0xb3, 0x72, 0xac, 0xf7, 0xd5, 0xe2, 0x84, 0x69, 0xd5, 0xd8, 0x72, 0xf0, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x96, 0xb7, 0x42, 0xe1, 0x02, 0x00, 0x00,
}
