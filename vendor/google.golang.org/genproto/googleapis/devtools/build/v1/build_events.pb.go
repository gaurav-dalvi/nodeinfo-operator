// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/build/v1/build_events.proto

package build

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The type of console output stream.
type ConsoleOutputStream int32

const (
	// Unspecified or unknown.
	ConsoleOutputStream_UNKNOWN ConsoleOutputStream = 0
	// Normal output stream.
	ConsoleOutputStream_STDOUT ConsoleOutputStream = 1
	// Error output stream.
	ConsoleOutputStream_STDERR ConsoleOutputStream = 2
)

var ConsoleOutputStream_name = map[int32]string{
	0: "UNKNOWN",
	1: "STDOUT",
	2: "STDERR",
}

var ConsoleOutputStream_value = map[string]int32{
	"UNKNOWN": 0,
	"STDOUT":  1,
	"STDERR":  2,
}

func (x ConsoleOutputStream) String() string {
	return proto.EnumName(ConsoleOutputStream_name, int32(x))
}

func (ConsoleOutputStream) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b1e5c65e074f08f4, []int{0}
}

// How did the event stream finish.
type BuildEvent_BuildComponentStreamFinished_FinishType int32

const (
	// Unknown or unspecified; callers should never set this value.
	BuildEvent_BuildComponentStreamFinished_FINISH_TYPE_UNSPECIFIED BuildEvent_BuildComponentStreamFinished_FinishType = 0
	// Set by the event publisher to indicate a build event stream is
	// finished.
	BuildEvent_BuildComponentStreamFinished_FINISHED BuildEvent_BuildComponentStreamFinished_FinishType = 1
	// Set by the WatchBuild RPC server when the publisher of a build event
	// stream stops publishing events without publishing a
	// BuildComponentStreamFinished event whose type equals FINISHED.
	BuildEvent_BuildComponentStreamFinished_EXPIRED BuildEvent_BuildComponentStreamFinished_FinishType = 2
)

var BuildEvent_BuildComponentStreamFinished_FinishType_name = map[int32]string{
	0: "FINISH_TYPE_UNSPECIFIED",
	1: "FINISHED",
	2: "EXPIRED",
}

var BuildEvent_BuildComponentStreamFinished_FinishType_value = map[string]int32{
	"FINISH_TYPE_UNSPECIFIED": 0,
	"FINISHED":                1,
	"EXPIRED":                 2,
}

func (x BuildEvent_BuildComponentStreamFinished_FinishType) String() string {
	return proto.EnumName(BuildEvent_BuildComponentStreamFinished_FinishType_name, int32(x))
}

func (BuildEvent_BuildComponentStreamFinished_FinishType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b1e5c65e074f08f4, []int{0, 5, 0}
}

// Which build component generates this event stream. Each build component
// may generate one event stream.
type StreamId_BuildComponent int32

const (
	// Unknown or unspecified; callers should never set this value.
	StreamId_UNKNOWN_COMPONENT StreamId_BuildComponent = 0
	// A component that coordinates builds.
	StreamId_CONTROLLER StreamId_BuildComponent = 1
	// A component that runs executables needed to complete a build.
	StreamId_WORKER StreamId_BuildComponent = 2
	// A component that builds something.
	StreamId_TOOL StreamId_BuildComponent = 3
)

var StreamId_BuildComponent_name = map[int32]string{
	0: "UNKNOWN_COMPONENT",
	1: "CONTROLLER",
	2: "WORKER",
	3: "TOOL",
}

var StreamId_BuildComponent_value = map[string]int32{
	"UNKNOWN_COMPONENT": 0,
	"CONTROLLER":        1,
	"WORKER":            2,
	"TOOL":              3,
}

func (x StreamId_BuildComponent) String() string {
	return proto.EnumName(StreamId_BuildComponent_name, int32(x))
}

func (StreamId_BuildComponent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b1e5c65e074f08f4, []int{1, 0}
}

// An event representing some state change that occurred in the build. This
// message does not include field for uniquely identifying an event.
type BuildEvent struct {
	// The timestamp of this event.
	EventTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// //////////////////////////////////////////////////////////////////////////
	// Events that indicate a state change of a build request in the build
	// queue.
	//
	// Types that are valid to be assigned to Event:
	//	*BuildEvent_InvocationAttemptStarted_
	//	*BuildEvent_InvocationAttemptFinished_
	//	*BuildEvent_BuildEnqueued_
	//	*BuildEvent_BuildFinished_
	//	*BuildEvent_ConsoleOutput_
	//	*BuildEvent_ComponentStreamFinished
	//	*BuildEvent_BazelEvent
	//	*BuildEvent_BuildExecutionEvent
	//	*BuildEvent_SourceFetchEvent
	Event                isBuildEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BuildEvent) Reset()         { *m = BuildEvent{} }
func (m *BuildEvent) String() string { return proto.CompactTextString(m) }
func (*BuildEvent) ProtoMessage()    {}
func (*BuildEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e5c65e074f08f4, []int{0}
}

func (m *BuildEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEvent.Unmarshal(m, b)
}
func (m *BuildEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEvent.Marshal(b, m, deterministic)
}
func (m *BuildEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEvent.Merge(m, src)
}
func (m *BuildEvent) XXX_Size() int {
	return xxx_messageInfo_BuildEvent.Size(m)
}
func (m *BuildEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEvent proto.InternalMessageInfo

func (m *BuildEvent) GetEventTime() *timestamp.Timestamp {
	if m != nil {
		return m.EventTime
	}
	return nil
}

type isBuildEvent_Event interface {
	isBuildEvent_Event()
}

type BuildEvent_InvocationAttemptStarted_ struct {
	InvocationAttemptStarted *BuildEvent_InvocationAttemptStarted `protobuf:"bytes,51,opt,name=invocation_attempt_started,json=invocationAttemptStarted,proto3,oneof"`
}

type BuildEvent_InvocationAttemptFinished_ struct {
	InvocationAttemptFinished *BuildEvent_InvocationAttemptFinished `protobuf:"bytes,52,opt,name=invocation_attempt_finished,json=invocationAttemptFinished,proto3,oneof"`
}

type BuildEvent_BuildEnqueued_ struct {
	BuildEnqueued *BuildEvent_BuildEnqueued `protobuf:"bytes,53,opt,name=build_enqueued,json=buildEnqueued,proto3,oneof"`
}

type BuildEvent_BuildFinished_ struct {
	BuildFinished *BuildEvent_BuildFinished `protobuf:"bytes,55,opt,name=build_finished,json=buildFinished,proto3,oneof"`
}

type BuildEvent_ConsoleOutput_ struct {
	ConsoleOutput *BuildEvent_ConsoleOutput `protobuf:"bytes,56,opt,name=console_output,json=consoleOutput,proto3,oneof"`
}

type BuildEvent_ComponentStreamFinished struct {
	ComponentStreamFinished *BuildEvent_BuildComponentStreamFinished `protobuf:"bytes,59,opt,name=component_stream_finished,json=componentStreamFinished,proto3,oneof"`
}

type BuildEvent_BazelEvent struct {
	BazelEvent *any.Any `protobuf:"bytes,60,opt,name=bazel_event,json=bazelEvent,proto3,oneof"`
}

type BuildEvent_BuildExecutionEvent struct {
	BuildExecutionEvent *any.Any `protobuf:"bytes,61,opt,name=build_execution_event,json=buildExecutionEvent,proto3,oneof"`
}

type BuildEvent_SourceFetchEvent struct {
	SourceFetchEvent *any.Any `protobuf:"bytes,62,opt,name=source_fetch_event,json=sourceFetchEvent,proto3,oneof"`
}

func (*BuildEvent_InvocationAttemptStarted_) isBuildEvent_Event() {}

func (*BuildEvent_InvocationAttemptFinished_) isBuildEvent_Event() {}

func (*BuildEvent_BuildEnqueued_) isBuildEvent_Event() {}

func (*BuildEvent_BuildFinished_) isBuildEvent_Event() {}

func (*BuildEvent_ConsoleOutput_) isBuildEvent_Event() {}

func (*BuildEvent_ComponentStreamFinished) isBuildEvent_Event() {}

func (*BuildEvent_BazelEvent) isBuildEvent_Event() {}

func (*BuildEvent_BuildExecutionEvent) isBuildEvent_Event() {}

func (*BuildEvent_SourceFetchEvent) isBuildEvent_Event() {}

func (m *BuildEvent) GetEvent() isBuildEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *BuildEvent) GetInvocationAttemptStarted() *BuildEvent_InvocationAttemptStarted {
	if x, ok := m.GetEvent().(*BuildEvent_InvocationAttemptStarted_); ok {
		return x.InvocationAttemptStarted
	}
	return nil
}

func (m *BuildEvent) GetInvocationAttemptFinished() *BuildEvent_InvocationAttemptFinished {
	if x, ok := m.GetEvent().(*BuildEvent_InvocationAttemptFinished_); ok {
		return x.InvocationAttemptFinished
	}
	return nil
}

func (m *BuildEvent) GetBuildEnqueued() *BuildEvent_BuildEnqueued {
	if x, ok := m.GetEvent().(*BuildEvent_BuildEnqueued_); ok {
		return x.BuildEnqueued
	}
	return nil
}

func (m *BuildEvent) GetBuildFinished() *BuildEvent_BuildFinished {
	if x, ok := m.GetEvent().(*BuildEvent_BuildFinished_); ok {
		return x.BuildFinished
	}
	return nil
}

func (m *BuildEvent) GetConsoleOutput() *BuildEvent_ConsoleOutput {
	if x, ok := m.GetEvent().(*BuildEvent_ConsoleOutput_); ok {
		return x.ConsoleOutput
	}
	return nil
}

func (m *BuildEvent) GetComponentStreamFinished() *BuildEvent_BuildComponentStreamFinished {
	if x, ok := m.GetEvent().(*BuildEvent_ComponentStreamFinished); ok {
		return x.ComponentStreamFinished
	}
	return nil
}

func (m *BuildEvent) GetBazelEvent() *any.Any {
	if x, ok := m.GetEvent().(*BuildEvent_BazelEvent); ok {
		return x.BazelEvent
	}
	return nil
}

func (m *BuildEvent) GetBuildExecutionEvent() *any.Any {
	if x, ok := m.GetEvent().(*BuildEvent_BuildExecutionEvent); ok {
		return x.BuildExecutionEvent
	}
	return nil
}

func (m *BuildEvent) GetSourceFetchEvent() *any.Any {
	if x, ok := m.GetEvent().(*BuildEvent_SourceFetchEvent); ok {
		return x.SourceFetchEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BuildEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BuildEvent_InvocationAttemptStarted_)(nil),
		(*BuildEvent_InvocationAttemptFinished_)(nil),
		(*BuildEvent_BuildEnqueued_)(nil),
		(*BuildEvent_BuildFinished_)(nil),
		(*BuildEvent_ConsoleOutput_)(nil),
		(*BuildEvent_ComponentStreamFinished)(nil),
		(*BuildEvent_BazelEvent)(nil),
		(*BuildEvent_BuildExecutionEvent)(nil),
		(*BuildEvent_SourceFetchEvent)(nil),
	}
}

// Notification that the build system has attempted to run the build tool.
type BuildEvent_InvocationAttemptStarted struct {
	// The number of the invocation attempt, starting at 1 and increasing by 1
	// for each new attempt. Can be used to determine if there is a later
	// invocation attempt replacing the current one a client is processing.
	AttemptNumber int64 `protobuf:"varint,1,opt,name=attempt_number,json=attemptNumber,proto3" json:"attempt_number,omitempty"`
	// Additional details about the invocation.
	Details              *any.Any `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildEvent_InvocationAttemptStarted) Reset()         { *m = BuildEvent_InvocationAttemptStarted{} }
func (m *BuildEvent_InvocationAttemptStarted) String() string { return proto.CompactTextString(m) }
func (*BuildEvent_InvocationAttemptStarted) ProtoMessage()    {}
func (*BuildEvent_InvocationAttemptStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e5c65e074f08f4, []int{0, 0}
}

func (m *BuildEvent_InvocationAttemptStarted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEvent_InvocationAttemptStarted.Unmarshal(m, b)
}
func (m *BuildEvent_InvocationAttemptStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEvent_InvocationAttemptStarted.Marshal(b, m, deterministic)
}
func (m *BuildEvent_InvocationAttemptStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEvent_InvocationAttemptStarted.Merge(m, src)
}
func (m *BuildEvent_InvocationAttemptStarted) XXX_Size() int {
	return xxx_messageInfo_BuildEvent_InvocationAttemptStarted.Size(m)
}
func (m *BuildEvent_InvocationAttemptStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEvent_InvocationAttemptStarted.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEvent_InvocationAttemptStarted proto.InternalMessageInfo

func (m *BuildEvent_InvocationAttemptStarted) GetAttemptNumber() int64 {
	if m != nil {
		return m.AttemptNumber
	}
	return 0
}

func (m *BuildEvent_InvocationAttemptStarted) GetDetails() *any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

// Notification that an invocation attempt has finished.
type BuildEvent_InvocationAttemptFinished struct {
	// Final status of the invocation.
	InvocationStatus     *BuildStatus `protobuf:"bytes,3,opt,name=invocation_status,json=invocationStatus,proto3" json:"invocation_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BuildEvent_InvocationAttemptFinished) Reset()         { *m = BuildEvent_InvocationAttemptFinished{} }
func (m *BuildEvent_InvocationAttemptFinished) String() string { return proto.CompactTextString(m) }
func (*BuildEvent_InvocationAttemptFinished) ProtoMessage()    {}
func (*BuildEvent_InvocationAttemptFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e5c65e074f08f4, []int{0, 1}
}

func (m *BuildEvent_InvocationAttemptFinished) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEvent_InvocationAttemptFinished.Unmarshal(m, b)
}
func (m *BuildEvent_InvocationAttemptFinished) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEvent_InvocationAttemptFinished.Marshal(b, m, deterministic)
}
func (m *BuildEvent_InvocationAttemptFinished) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEvent_InvocationAttemptFinished.Merge(m, src)
}
func (m *BuildEvent_InvocationAttemptFinished) XXX_Size() int {
	return xxx_messageInfo_BuildEvent_InvocationAttemptFinished.Size(m)
}
func (m *BuildEvent_InvocationAttemptFinished) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEvent_InvocationAttemptFinished.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEvent_InvocationAttemptFinished proto.InternalMessageInfo

func (m *BuildEvent_InvocationAttemptFinished) GetInvocationStatus() *BuildStatus {
	if m != nil {
		return m.InvocationStatus
	}
	return nil
}

// Notification that the build request is enqueued.
type BuildEvent_BuildEnqueued struct {
	// Additional details about the Build.
	Details              *any.Any `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildEvent_BuildEnqueued) Reset()         { *m = BuildEvent_BuildEnqueued{} }
func (m *BuildEvent_BuildEnqueued) String() string { return proto.CompactTextString(m) }
func (*BuildEvent_BuildEnqueued) ProtoMessage()    {}
func (*BuildEvent_BuildEnqueued) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e5c65e074f08f4, []int{0, 2}
}

func (m *BuildEvent_BuildEnqueued) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEvent_BuildEnqueued.Unmarshal(m, b)
}
func (m *BuildEvent_BuildEnqueued) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEvent_BuildEnqueued.Marshal(b, m, deterministic)
}
func (m *BuildEvent_BuildEnqueued) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEvent_BuildEnqueued.Merge(m, src)
}
func (m *BuildEvent_BuildEnqueued) XXX_Size() int {
	return xxx_messageInfo_BuildEvent_BuildEnqueued.Size(m)
}
func (m *BuildEvent_BuildEnqueued) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEvent_BuildEnqueued.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEvent_BuildEnqueued proto.InternalMessageInfo

func (m *BuildEvent_BuildEnqueued) GetDetails() *any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

// Notification that the build request has finished, and no further
// invocations will occur.  Note that this applies to the entire Build.
// Individual invocations trigger InvocationFinished when they finish.
type BuildEvent_BuildFinished struct {
	// Final status of the build.
	Status               *BuildStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BuildEvent_BuildFinished) Reset()         { *m = BuildEvent_BuildFinished{} }
func (m *BuildEvent_BuildFinished) String() string { return proto.CompactTextString(m) }
func (*BuildEvent_BuildFinished) ProtoMessage()    {}
func (*BuildEvent_BuildFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e5c65e074f08f4, []int{0, 3}
}

func (m *BuildEvent_BuildFinished) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEvent_BuildFinished.Unmarshal(m, b)
}
func (m *BuildEvent_BuildFinished) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEvent_BuildFinished.Marshal(b, m, deterministic)
}
func (m *BuildEvent_BuildFinished) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEvent_BuildFinished.Merge(m, src)
}
func (m *BuildEvent_BuildFinished) XXX_Size() int {
	return xxx_messageInfo_BuildEvent_BuildFinished.Size(m)
}
func (m *BuildEvent_BuildFinished) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEvent_BuildFinished.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEvent_BuildFinished proto.InternalMessageInfo

func (m *BuildEvent_BuildFinished) GetStatus() *BuildStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Textual output written to standard output or standard error.
type BuildEvent_ConsoleOutput struct {
	// The output stream type.
	Type ConsoleOutputStream `protobuf:"varint,1,opt,name=type,proto3,enum=google.devtools.build.v1.ConsoleOutputStream" json:"type,omitempty"`
	// The output stream content.
	//
	// Types that are valid to be assigned to Output:
	//	*BuildEvent_ConsoleOutput_TextOutput
	//	*BuildEvent_ConsoleOutput_BinaryOutput
	Output               isBuildEvent_ConsoleOutput_Output `protobuf_oneof:"output"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *BuildEvent_ConsoleOutput) Reset()         { *m = BuildEvent_ConsoleOutput{} }
func (m *BuildEvent_ConsoleOutput) String() string { return proto.CompactTextString(m) }
func (*BuildEvent_ConsoleOutput) ProtoMessage()    {}
func (*BuildEvent_ConsoleOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e5c65e074f08f4, []int{0, 4}
}

func (m *BuildEvent_ConsoleOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEvent_ConsoleOutput.Unmarshal(m, b)
}
func (m *BuildEvent_ConsoleOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEvent_ConsoleOutput.Marshal(b, m, deterministic)
}
func (m *BuildEvent_ConsoleOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEvent_ConsoleOutput.Merge(m, src)
}
func (m *BuildEvent_ConsoleOutput) XXX_Size() int {
	return xxx_messageInfo_BuildEvent_ConsoleOutput.Size(m)
}
func (m *BuildEvent_ConsoleOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEvent_ConsoleOutput.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEvent_ConsoleOutput proto.InternalMessageInfo

func (m *BuildEvent_ConsoleOutput) GetType() ConsoleOutputStream {
	if m != nil {
		return m.Type
	}
	return ConsoleOutputStream_UNKNOWN
}

type isBuildEvent_ConsoleOutput_Output interface {
	isBuildEvent_ConsoleOutput_Output()
}

type BuildEvent_ConsoleOutput_TextOutput struct {
	TextOutput string `protobuf:"bytes,2,opt,name=text_output,json=textOutput,proto3,oneof"`
}

type BuildEvent_ConsoleOutput_BinaryOutput struct {
	BinaryOutput []byte `protobuf:"bytes,3,opt,name=binary_output,json=binaryOutput,proto3,oneof"`
}

func (*BuildEvent_ConsoleOutput_TextOutput) isBuildEvent_ConsoleOutput_Output() {}

func (*BuildEvent_ConsoleOutput_BinaryOutput) isBuildEvent_ConsoleOutput_Output() {}

func (m *BuildEvent_ConsoleOutput) GetOutput() isBuildEvent_ConsoleOutput_Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *BuildEvent_ConsoleOutput) GetTextOutput() string {
	if x, ok := m.GetOutput().(*BuildEvent_ConsoleOutput_TextOutput); ok {
		return x.TextOutput
	}
	return ""
}

func (m *BuildEvent_ConsoleOutput) GetBinaryOutput() []byte {
	if x, ok := m.GetOutput().(*BuildEvent_ConsoleOutput_BinaryOutput); ok {
		return x.BinaryOutput
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BuildEvent_ConsoleOutput) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BuildEvent_ConsoleOutput_TextOutput)(nil),
		(*BuildEvent_ConsoleOutput_BinaryOutput)(nil),
	}
}

// Notification of the end of a build event stream published by a build
// component other than CONTROLLER (See StreamId.BuildComponents).
type BuildEvent_BuildComponentStreamFinished struct {
	// How the event stream finished.
	Type                 BuildEvent_BuildComponentStreamFinished_FinishType `protobuf:"varint,1,opt,name=type,proto3,enum=google.devtools.build.v1.BuildEvent_BuildComponentStreamFinished_FinishType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *BuildEvent_BuildComponentStreamFinished) Reset() {
	*m = BuildEvent_BuildComponentStreamFinished{}
}
func (m *BuildEvent_BuildComponentStreamFinished) String() string { return proto.CompactTextString(m) }
func (*BuildEvent_BuildComponentStreamFinished) ProtoMessage()    {}
func (*BuildEvent_BuildComponentStreamFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e5c65e074f08f4, []int{0, 5}
}

func (m *BuildEvent_BuildComponentStreamFinished) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEvent_BuildComponentStreamFinished.Unmarshal(m, b)
}
func (m *BuildEvent_BuildComponentStreamFinished) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEvent_BuildComponentStreamFinished.Marshal(b, m, deterministic)
}
func (m *BuildEvent_BuildComponentStreamFinished) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEvent_BuildComponentStreamFinished.Merge(m, src)
}
func (m *BuildEvent_BuildComponentStreamFinished) XXX_Size() int {
	return xxx_messageInfo_BuildEvent_BuildComponentStreamFinished.Size(m)
}
func (m *BuildEvent_BuildComponentStreamFinished) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEvent_BuildComponentStreamFinished.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEvent_BuildComponentStreamFinished proto.InternalMessageInfo

func (m *BuildEvent_BuildComponentStreamFinished) GetType() BuildEvent_BuildComponentStreamFinished_FinishType {
	if m != nil {
		return m.Type
	}
	return BuildEvent_BuildComponentStreamFinished_FINISH_TYPE_UNSPECIFIED
}

// Unique identifier for a build event stream.
type StreamId struct {
	// The id of a Build message.
	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	// The unique invocation ID within this build.
	// It should be the same as {invocation} (below) during the migration.
	InvocationId string `protobuf:"bytes,6,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	// The component that emitted this event.
	Component            StreamId_BuildComponent `protobuf:"varint,3,opt,name=component,proto3,enum=google.devtools.build.v1.StreamId_BuildComponent" json:"component,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StreamId) Reset()         { *m = StreamId{} }
func (m *StreamId) String() string { return proto.CompactTextString(m) }
func (*StreamId) ProtoMessage()    {}
func (*StreamId) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e5c65e074f08f4, []int{1}
}

func (m *StreamId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamId.Unmarshal(m, b)
}
func (m *StreamId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamId.Marshal(b, m, deterministic)
}
func (m *StreamId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamId.Merge(m, src)
}
func (m *StreamId) XXX_Size() int {
	return xxx_messageInfo_StreamId.Size(m)
}
func (m *StreamId) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamId.DiscardUnknown(m)
}

var xxx_messageInfo_StreamId proto.InternalMessageInfo

func (m *StreamId) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

func (m *StreamId) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

func (m *StreamId) GetComponent() StreamId_BuildComponent {
	if m != nil {
		return m.Component
	}
	return StreamId_UNKNOWN_COMPONENT
}

func init() {
	proto.RegisterEnum("google.devtools.build.v1.ConsoleOutputStream", ConsoleOutputStream_name, ConsoleOutputStream_value)
	proto.RegisterEnum("google.devtools.build.v1.BuildEvent_BuildComponentStreamFinished_FinishType", BuildEvent_BuildComponentStreamFinished_FinishType_name, BuildEvent_BuildComponentStreamFinished_FinishType_value)
	proto.RegisterEnum("google.devtools.build.v1.StreamId_BuildComponent", StreamId_BuildComponent_name, StreamId_BuildComponent_value)
	proto.RegisterType((*BuildEvent)(nil), "google.devtools.build.v1.BuildEvent")
	proto.RegisterType((*BuildEvent_InvocationAttemptStarted)(nil), "google.devtools.build.v1.BuildEvent.InvocationAttemptStarted")
	proto.RegisterType((*BuildEvent_InvocationAttemptFinished)(nil), "google.devtools.build.v1.BuildEvent.InvocationAttemptFinished")
	proto.RegisterType((*BuildEvent_BuildEnqueued)(nil), "google.devtools.build.v1.BuildEvent.BuildEnqueued")
	proto.RegisterType((*BuildEvent_BuildFinished)(nil), "google.devtools.build.v1.BuildEvent.BuildFinished")
	proto.RegisterType((*BuildEvent_ConsoleOutput)(nil), "google.devtools.build.v1.BuildEvent.ConsoleOutput")
	proto.RegisterType((*BuildEvent_BuildComponentStreamFinished)(nil), "google.devtools.build.v1.BuildEvent.BuildComponentStreamFinished")
	proto.RegisterType((*StreamId)(nil), "google.devtools.build.v1.StreamId")
}

func init() {
	proto.RegisterFile("google/devtools/build/v1/build_events.proto", fileDescriptor_b1e5c65e074f08f4)
}

var fileDescriptor_b1e5c65e074f08f4 = []byte{
	// 896 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0x6d, 0x6f, 0xe3, 0x44,
	0x10, 0xc7, 0xe3, 0xf6, 0x48, 0xd3, 0x69, 0x13, 0x7c, 0x7b, 0x9c, 0xce, 0xf1, 0x55, 0x02, 0x8a,
	0x2a, 0x21, 0x10, 0x8e, 0xda, 0x03, 0x1d, 0x70, 0xf4, 0x50, 0x93, 0xb8, 0x8a, 0xb9, 0x9e, 0x1d,
	0x6d, 0x5c, 0x1d, 0x0f, 0x2f, 0x82, 0x63, 0x6f, 0x73, 0x96, 0x12, 0xaf, 0x2f, 0x5e, 0x47, 0x17,
	0x24, 0x04, 0x1f, 0x84, 0xd7, 0x7c, 0x14, 0x3e, 0x0c, 0x9f, 0x80, 0x97, 0xc8, 0xbb, 0xeb, 0xc4,
	0xb9, 0xab, 0xfb, 0x00, 0xef, 0x36, 0x33, 0xff, 0xf9, 0xcd, 0xcc, 0xee, 0x8c, 0x15, 0xf8, 0x74,
	0x4c, 0xe9, 0x78, 0x42, 0x5a, 0x01, 0x99, 0x33, 0x4a, 0x27, 0x49, 0x6b, 0x94, 0x86, 0x93, 0xa0,
	0x35, 0x3f, 0x14, 0x87, 0x21, 0x99, 0x93, 0x88, 0x25, 0x46, 0x3c, 0xa3, 0x8c, 0x22, 0x4d, 0x88,
	0x8d, 0x5c, 0x6c, 0x70, 0x8d, 0x31, 0x3f, 0xd4, 0xf7, 0x24, 0xc6, 0x8b, 0xc3, 0x96, 0x17, 0x45,
	0x94, 0x79, 0x2c, 0xa4, 0x91, 0x8c, 0xd3, 0xaf, 0x4b, 0x92, 0x30, 0x8f, 0xa5, 0xb9, 0xb8, 0x29,
	0xc5, 0xfc, 0xd7, 0x28, 0xbd, 0x68, 0x79, 0xd1, 0x42, 0xba, 0xde, 0x7f, 0xd3, 0xc5, 0xc2, 0x29,
	0x49, 0x98, 0x37, 0x8d, 0x85, 0x60, 0xff, 0x8f, 0x3a, 0x40, 0x3b, 0x43, 0x9a, 0x59, 0xd9, 0xe8,
	0x2b, 0x00, 0x5e, 0xff, 0x30, 0xd3, 0x69, 0xca, 0x07, 0xca, 0xc7, 0x3b, 0x47, 0xba, 0x21, 0x9b,
	0xc8, 0x21, 0x86, 0x9b, 0x43, 0xf0, 0x36, 0x57, 0x67, 0xbf, 0xd1, 0xaf, 0xa0, 0x87, 0xd1, 0x9c,
	0xfa, 0xbc, 0x8f, 0xa1, 0xc7, 0x18, 0x99, 0xc6, 0x2c, 0x2b, 0x74, 0xc6, 0x48, 0xa0, 0x3d, 0xe2,
	0xa8, 0x63, 0xa3, 0xec, 0x3e, 0x8c, 0x55, 0x11, 0x86, 0xb5, 0xc4, 0x9c, 0x08, 0xca, 0x40, 0x40,
	0x7a, 0x15, 0xac, 0x85, 0x25, 0x3e, 0xf4, 0xbb, 0x02, 0x0f, 0x2f, 0xc9, 0x7f, 0x11, 0x46, 0x61,
	0xf2, 0x92, 0x04, 0xda, 0xe7, 0xbc, 0x80, 0xa7, 0xff, 0xad, 0x80, 0x53, 0x49, 0xe9, 0x55, 0x70,
	0x33, 0x2c, 0x73, 0xa2, 0x9f, 0xa0, 0x21, 0x47, 0x20, 0x7a, 0x95, 0x92, 0x94, 0x04, 0xda, 0x17,
	0x3c, 0xe9, 0xd1, 0x8d, 0x92, 0x8a, 0xa3, 0x8c, 0xec, 0x55, 0x70, 0x7d, 0x54, 0x34, 0xac, 0xe0,
	0xcb, 0x8e, 0x1e, 0xdf, 0x16, 0x5e, 0xe8, 0x42, 0xc0, 0x8b, 0x95, 0xfb, 0x34, 0x4a, 0xe8, 0x84,
	0x0c, 0x69, 0xca, 0xe2, 0x94, 0x69, 0x5f, 0xde, 0x02, 0xde, 0x11, 0xa1, 0x0e, 0x8f, 0xcc, 0xe0,
	0x7e, 0xd1, 0x80, 0x7e, 0x83, 0xa6, 0x4f, 0xa7, 0x31, 0x8d, 0xb2, 0xb9, 0x4a, 0xd8, 0x8c, 0x78,
	0xd3, 0x55, 0x13, 0x4f, 0x78, 0x9e, 0x93, 0x9b, 0x37, 0xd1, 0xc9, 0x51, 0x03, 0x4e, 0x2a, 0xf4,
	0xf4, 0xc0, 0xbf, 0xdc, 0x85, 0x1e, 0xc3, 0xce, 0xc8, 0xfb, 0x85, 0x4c, 0xc4, 0x6a, 0x6a, 0xdf,
	0xf0, 0x94, 0xef, 0xbd, 0x35, 0xd5, 0x27, 0xd1, 0xa2, 0x57, 0xc1, 0xc0, 0xa5, 0x62, 0x1b, 0xbe,
	0x83, 0xfb, 0xf2, 0x41, 0x5f, 0x13, 0x3f, 0xe5, 0x73, 0x25, 0x10, 0xc7, 0x57, 0x22, 0xee, 0x89,
	0x97, 0xcb, 0x63, 0x04, 0xab, 0x0b, 0x28, 0xa1, 0xe9, 0xcc, 0x27, 0xc3, 0x0b, 0xc2, 0xfc, 0x97,
	0x12, 0xf4, 0xf4, 0x4a, 0x90, 0x2a, 0x22, 0x4e, 0xb3, 0x00, 0x4e, 0xd1, 0x5f, 0x81, 0x56, 0xb6,
	0x1d, 0xe8, 0x00, 0x1a, 0xf9, 0xd4, 0x47, 0xe9, 0x74, 0x44, 0x66, 0x7c, 0x7f, 0x37, 0x71, 0x5d,
	0x5a, 0x6d, 0x6e, 0x44, 0x06, 0x6c, 0x05, 0x84, 0x79, 0xe1, 0x24, 0xd1, 0x36, 0xca, 0xb3, 0xe3,
	0x5c, 0xa4, 0x53, 0x68, 0x96, 0xee, 0x03, 0xc2, 0x70, 0xb7, 0xb0, 0x74, 0xe2, 0xab, 0xa4, 0x6d,
	0x72, 0xec, 0xc1, 0x35, 0x6f, 0x3a, 0xe0, 0x62, 0xac, 0xae, 0xe2, 0x85, 0x45, 0xff, 0x16, 0xea,
	0x6b, 0xbb, 0x50, 0xac, 0x58, 0xb9, 0x49, 0xc5, 0xb6, 0x04, 0x2c, 0xab, 0x3c, 0x86, 0xaa, 0x2c,
	0x4d, 0xb9, 0x4d, 0x69, 0x32, 0x48, 0xff, 0x53, 0x81, 0xfa, 0xda, 0x8c, 0xa3, 0x13, 0xb8, 0xc3,
	0x16, 0xb1, 0xf8, 0x40, 0x36, 0x8e, 0x3e, 0x2b, 0xc7, 0xad, 0x85, 0x89, 0xb1, 0xc4, 0x3c, 0x14,
	0x7d, 0x08, 0x3b, 0x8c, 0xbc, 0x66, 0xf9, 0xbe, 0x65, 0x4f, 0xb1, 0x9d, 0x8d, 0x5f, 0x66, 0x94,
	0x59, 0x0e, 0xa0, 0x3e, 0x0a, 0x23, 0x6f, 0xb6, 0xc8, 0x45, 0xd9, 0xc5, 0xee, 0xf6, 0x2a, 0x78,
	0x57, 0x98, 0x85, 0xac, 0x5d, 0x83, 0xaa, 0xf0, 0xeb, 0x7f, 0x29, 0xb0, 0x77, 0xd5, 0x92, 0xa0,
	0x9f, 0xd7, 0xea, 0x3e, 0xfb, 0xdf, 0x5b, 0x67, 0x88, 0x83, 0xbb, 0x88, 0x89, 0x68, 0x6b, 0xbf,
	0x0b, 0xb0, 0xb2, 0xa1, 0x87, 0xf0, 0xe0, 0xd4, 0xb2, 0xad, 0x41, 0x6f, 0xe8, 0xfe, 0xd0, 0x37,
	0x87, 0xe7, 0xf6, 0xa0, 0x6f, 0x76, 0xac, 0x53, 0xcb, 0xec, 0xaa, 0x15, 0xb4, 0x0b, 0x35, 0xe1,
	0x34, 0xbb, 0xaa, 0x82, 0x76, 0x60, 0xcb, 0xfc, 0xbe, 0x6f, 0x61, 0xb3, 0xab, 0x6e, 0xb4, 0xb7,
	0xe0, 0x1d, 0xbe, 0x1f, 0xfb, 0x7f, 0x2b, 0x50, 0x13, 0x29, 0xad, 0x00, 0x35, 0xa1, 0x26, 0xd6,
	0x31, 0x0c, 0x78, 0x07, 0xdb, 0x78, 0x8b, 0xff, 0xb6, 0x02, 0xf4, 0x11, 0xd4, 0x0b, 0x73, 0x18,
	0x06, 0x5a, 0x95, 0xfb, 0x77, 0x57, 0x46, 0x2b, 0x40, 0x0e, 0x6c, 0x2f, 0x3f, 0x11, 0xfc, 0x2e,
	0x1b, 0x47, 0x87, 0xe5, 0x57, 0x90, 0xa7, 0x7d, 0xe3, 0x02, 0xf0, 0x8a, 0xb1, 0xff, 0x1c, 0x1a,
	0xeb, 0x4e, 0x74, 0x1f, 0xee, 0x9e, 0xdb, 0xcf, 0x6c, 0xe7, 0x85, 0x3d, 0xec, 0x38, 0xcf, 0xfb,
	0x8e, 0x6d, 0xda, 0xae, 0x5a, 0x41, 0x0d, 0x80, 0x8e, 0x63, 0xbb, 0xd8, 0x39, 0x3b, 0x33, 0xb1,
	0xaa, 0x20, 0x80, 0xea, 0x0b, 0x07, 0x3f, 0x33, 0xb1, 0xba, 0x81, 0x6a, 0x70, 0xc7, 0x75, 0x9c,
	0x33, 0x75, 0xf3, 0x93, 0xaf, 0xe1, 0xde, 0x25, 0xf3, 0x92, 0xdd, 0x8c, 0x64, 0xaa, 0x95, 0x2c,
	0x72, 0xe0, 0x76, 0x9d, 0x73, 0x57, 0x50, 0x06, 0x6e, 0xd7, 0xc4, 0x58, 0xdd, 0x68, 0x27, 0xb0,
	0xe7, 0xd3, 0x69, 0x69, 0x37, 0xed, 0x77, 0x57, 0x2f, 0xda, 0xcf, 0x96, 0xa6, 0xaf, 0xfc, 0x78,
	0x2c, 0xc5, 0x63, 0x3a, 0xf1, 0xa2, 0xb1, 0x41, 0x67, 0xe3, 0xd6, 0x98, 0x44, 0x7c, 0xa5, 0x5a,
	0xc2, 0xe5, 0xc5, 0x61, 0xf2, 0xf6, 0x5f, 0x90, 0x27, 0xfc, 0xf0, 0x8f, 0xa2, 0x8c, 0xaa, 0x5c,
	0xfc, 0xe8, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xcf, 0xda, 0xac, 0x13, 0x09, 0x00, 0x00,
}
