// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lsp_mon.proto

/*
Package lsp_mon is a generated protocol buffer package.

It is generated from these files:
	lsp_mon.proto

It has these top-level messages:
	Key
	LspMonitorDataEvent
	EroTypeEntry
	EroIpv4Type
	RroTypeEntry
	RroIpv4Type
	LspMonitorDataProperty
	LspMon
*/
package lsp_mon

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LspEvent int32

const (
	LspEvent_INITIATED              LspEvent = 0
	LspEvent_CONCLUDED_UP           LspEvent = 1
	LspEvent_CONCLUDED_TORN_DOWN    LspEvent = 2
	LspEvent_PROTECTION_AVAILABLE   LspEvent = 3
	LspEvent_PROTECTION_UNAVAILABLE LspEvent = 4
	LspEvent_AUTOBW_SUCCESS         LspEvent = 5
	LspEvent_AUTOBW_FAIL            LspEvent = 6
	LspEvent_RESV_TEAR_RECEIVED     LspEvent = 7
	LspEvent_DESELECT_ACTIVE_PATH   LspEvent = 8
	LspEvent_CHANGE_ACTIVE_PATH     LspEvent = 9
	LspEvent_DETOUR_UP              LspEvent = 10
	LspEvent_DETOUR_DOWN            LspEvent = 11
	LspEvent_ORIGINATE_MBB          LspEvent = 12
	LspEvent_SELECT_ACTIVE_PATH     LspEvent = 13
	LspEvent_CSPF_NO_ROUTE          LspEvent = 14
	LspEvent_CSPF_SUCCESS           LspEvent = 15
	LspEvent_RESTART_RECOVERY_FAIL  LspEvent = 16
	LspEvent_PATHERR_RECEIVED       LspEvent = 17
	LspEvent_PATH_MTU_CHANGE        LspEvent = 18
	LspEvent_TUNNEL_LOCAL_REPAIRED  LspEvent = 19
)

var LspEvent_name = map[int32]string{
	0:  "INITIATED",
	1:  "CONCLUDED_UP",
	2:  "CONCLUDED_TORN_DOWN",
	3:  "PROTECTION_AVAILABLE",
	4:  "PROTECTION_UNAVAILABLE",
	5:  "AUTOBW_SUCCESS",
	6:  "AUTOBW_FAIL",
	7:  "RESV_TEAR_RECEIVED",
	8:  "DESELECT_ACTIVE_PATH",
	9:  "CHANGE_ACTIVE_PATH",
	10: "DETOUR_UP",
	11: "DETOUR_DOWN",
	12: "ORIGINATE_MBB",
	13: "SELECT_ACTIVE_PATH",
	14: "CSPF_NO_ROUTE",
	15: "CSPF_SUCCESS",
	16: "RESTART_RECOVERY_FAIL",
	17: "PATHERR_RECEIVED",
	18: "PATH_MTU_CHANGE",
	19: "TUNNEL_LOCAL_REPAIRED",
}
var LspEvent_value = map[string]int32{
	"INITIATED":              0,
	"CONCLUDED_UP":           1,
	"CONCLUDED_TORN_DOWN":    2,
	"PROTECTION_AVAILABLE":   3,
	"PROTECTION_UNAVAILABLE": 4,
	"AUTOBW_SUCCESS":         5,
	"AUTOBW_FAIL":            6,
	"RESV_TEAR_RECEIVED":     7,
	"DESELECT_ACTIVE_PATH":   8,
	"CHANGE_ACTIVE_PATH":     9,
	"DETOUR_UP":              10,
	"DETOUR_DOWN":            11,
	"ORIGINATE_MBB":          12,
	"SELECT_ACTIVE_PATH":     13,
	"CSPF_NO_ROUTE":          14,
	"CSPF_SUCCESS":           15,
	"RESTART_RECOVERY_FAIL":  16,
	"PATHERR_RECEIVED":       17,
	"PATH_MTU_CHANGE":        18,
	"TUNNEL_LOCAL_REPAIRED":  19,
}

func (x LspEvent) Enum() *LspEvent {
	p := new(LspEvent)
	*p = x
	return p
}
func (x LspEvent) String() string {
	return proto.EnumName(LspEvent_name, int32(x))
}
func (x *LspEvent) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LspEvent_value, data, "LspEvent")
	if err != nil {
		return err
	}
	*x = LspEvent(value)
	return nil
}
func (LspEvent) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EventSubcode int32

const (
	EventSubcode_ADMISSION_CONTROL_FAILURE       EventSubcode = 1
	EventSubcode_SESSION_PREEMPTED               EventSubcode = 2
	EventSubcode_BAD_LOOSE_ROUTE                 EventSubcode = 3
	EventSubcode_BAD_STRICT_ROUTE                EventSubcode = 4
	EventSubcode_LABEL_ALLOCATION_FAILURE        EventSubcode = 5
	EventSubcode_NON_RSVP_CAPABLE_ROUTER         EventSubcode = 6
	EventSubcode_TTL_EXPIRED                     EventSubcode = 7
	EventSubcode_ROUTING_LOOP_DETECTED           EventSubcode = 8
	EventSubcode_REQUESTED_BANDWIDTH_UNAVAILABLE EventSubcode = 9
)

var EventSubcode_name = map[int32]string{
	1: "ADMISSION_CONTROL_FAILURE",
	2: "SESSION_PREEMPTED",
	3: "BAD_LOOSE_ROUTE",
	4: "BAD_STRICT_ROUTE",
	5: "LABEL_ALLOCATION_FAILURE",
	6: "NON_RSVP_CAPABLE_ROUTER",
	7: "TTL_EXPIRED",
	8: "ROUTING_LOOP_DETECTED",
	9: "REQUESTED_BANDWIDTH_UNAVAILABLE",
}
var EventSubcode_value = map[string]int32{
	"ADMISSION_CONTROL_FAILURE":       1,
	"SESSION_PREEMPTED":               2,
	"BAD_LOOSE_ROUTE":                 3,
	"BAD_STRICT_ROUTE":                4,
	"LABEL_ALLOCATION_FAILURE":        5,
	"NON_RSVP_CAPABLE_ROUTER":         6,
	"TTL_EXPIRED":                     7,
	"ROUTING_LOOP_DETECTED":           8,
	"REQUESTED_BANDWIDTH_UNAVAILABLE": 9,
}

func (x EventSubcode) Enum() *EventSubcode {
	p := new(EventSubcode)
	*p = x
	return p
}
func (x EventSubcode) String() string {
	return proto.EnumName(EventSubcode_name, int32(x))
}
func (x *EventSubcode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EventSubcode_value, data, "EventSubcode")
	if err != nil {
		return err
	}
	*x = EventSubcode(value)
	return nil
}
func (EventSubcode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Key struct {
	Name               *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	InstanceIdentifier *int32  `protobuf:"varint,2,req,name=instance_identifier" json:"instance_identifier,omitempty"`
	TimeStampg         *uint64 `protobuf:"varint,3,req,name=time_stampg" json:"time_stampg,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Key) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Key) GetInstanceIdentifier() int32 {
	if m != nil && m.InstanceIdentifier != nil {
		return *m.InstanceIdentifier
	}
	return 0
}

func (m *Key) GetTimeStampg() uint64 {
	if m != nil && m.TimeStampg != nil {
		return *m.TimeStampg
	}
	return 0
}

type LspMonitorDataEvent struct {
	EventIdentifier  *LspEvent     `protobuf:"varint,1,req,name=event_identifier,enum=LspEvent" json:"event_identifier,omitempty"`
	Subcode          *EventSubcode `protobuf:"varint,2,opt,name=subcode,enum=EventSubcode" json:"subcode,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *LspMonitorDataEvent) Reset()                    { *m = LspMonitorDataEvent{} }
func (m *LspMonitorDataEvent) String() string            { return proto.CompactTextString(m) }
func (*LspMonitorDataEvent) ProtoMessage()               {}
func (*LspMonitorDataEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LspMonitorDataEvent) GetEventIdentifier() LspEvent {
	if m != nil && m.EventIdentifier != nil {
		return *m.EventIdentifier
	}
	return LspEvent_INITIATED
}

func (m *LspMonitorDataEvent) GetSubcode() EventSubcode {
	if m != nil && m.Subcode != nil {
		return *m.Subcode
	}
	return EventSubcode_ADMISSION_CONTROL_FAILURE
}

type EroTypeEntry struct {
	Ip               *uint32 `protobuf:"varint,1,req,name=ip" json:"ip,omitempty"`
	Flags            *string `protobuf:"bytes,2,opt,name=flags" json:"flags,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EroTypeEntry) Reset()                    { *m = EroTypeEntry{} }
func (m *EroTypeEntry) String() string            { return proto.CompactTextString(m) }
func (*EroTypeEntry) ProtoMessage()               {}
func (*EroTypeEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EroTypeEntry) GetIp() uint32 {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return 0
}

func (m *EroTypeEntry) GetFlags() string {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return ""
}

type EroIpv4Type struct {
	Entry            []*EroTypeEntry `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *EroIpv4Type) Reset()                    { *m = EroIpv4Type{} }
func (m *EroIpv4Type) String() string            { return proto.CompactTextString(m) }
func (*EroIpv4Type) ProtoMessage()               {}
func (*EroIpv4Type) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EroIpv4Type) GetEntry() []*EroTypeEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type RroTypeEntry struct {
	Nodeid           *uint32 `protobuf:"varint,1,opt,name=nodeid" json:"nodeid,omitempty"`
	Flags            *uint32 `protobuf:"varint,2,opt,name=flags" json:"flags,omitempty"`
	IntfAddr         *uint32 `protobuf:"varint,3,opt,name=intf_addr" json:"intf_addr,omitempty"`
	Label            *uint32 `protobuf:"varint,4,opt,name=label" json:"label,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RroTypeEntry) Reset()                    { *m = RroTypeEntry{} }
func (m *RroTypeEntry) String() string            { return proto.CompactTextString(m) }
func (*RroTypeEntry) ProtoMessage()               {}
func (*RroTypeEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RroTypeEntry) GetNodeid() uint32 {
	if m != nil && m.Nodeid != nil {
		return *m.Nodeid
	}
	return 0
}

func (m *RroTypeEntry) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *RroTypeEntry) GetIntfAddr() uint32 {
	if m != nil && m.IntfAddr != nil {
		return *m.IntfAddr
	}
	return 0
}

func (m *RroTypeEntry) GetLabel() uint32 {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return 0
}

type RroIpv4Type struct {
	RroEntry         []*RroTypeEntry `protobuf:"bytes,1,rep,name=rro_entry" json:"rro_entry,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RroIpv4Type) Reset()                    { *m = RroIpv4Type{} }
func (m *RroIpv4Type) String() string            { return proto.CompactTextString(m) }
func (*RroIpv4Type) ProtoMessage()               {}
func (*RroIpv4Type) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RroIpv4Type) GetRroEntry() []*RroTypeEntry {
	if m != nil {
		return m.RroEntry
	}
	return nil
}

type LspMonitorDataProperty struct {
	Bandwidth        *uint64      `protobuf:"varint,1,opt,name=bandwidth" json:"bandwidth,omitempty"`
	PathName         *string      `protobuf:"bytes,2,opt,name=path_name" json:"path_name,omitempty"`
	Metric           *int32       `protobuf:"varint,3,opt,name=metric" json:"metric,omitempty"`
	MaxAvgBw         *float32     `protobuf:"fixed32,4,opt,name=max_avg_bw" json:"max_avg_bw,omitempty"`
	Ero              *EroIpv4Type `protobuf:"bytes,5,opt,name=ero" json:"ero,omitempty"`
	Rro              *RroIpv4Type `protobuf:"bytes,6,opt,name=rro" json:"rro,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *LspMonitorDataProperty) Reset()                    { *m = LspMonitorDataProperty{} }
func (m *LspMonitorDataProperty) String() string            { return proto.CompactTextString(m) }
func (*LspMonitorDataProperty) ProtoMessage()               {}
func (*LspMonitorDataProperty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LspMonitorDataProperty) GetBandwidth() uint64 {
	if m != nil && m.Bandwidth != nil {
		return *m.Bandwidth
	}
	return 0
}

func (m *LspMonitorDataProperty) GetPathName() string {
	if m != nil && m.PathName != nil {
		return *m.PathName
	}
	return ""
}

func (m *LspMonitorDataProperty) GetMetric() int32 {
	if m != nil && m.Metric != nil {
		return *m.Metric
	}
	return 0
}

func (m *LspMonitorDataProperty) GetMaxAvgBw() float32 {
	if m != nil && m.MaxAvgBw != nil {
		return *m.MaxAvgBw
	}
	return 0
}

func (m *LspMonitorDataProperty) GetEro() *EroIpv4Type {
	if m != nil {
		return m.Ero
	}
	return nil
}

func (m *LspMonitorDataProperty) GetRro() *RroIpv4Type {
	if m != nil {
		return m.Rro
	}
	return nil
}

type LspMon struct {
	KeyField         *Key                    `protobuf:"bytes,1,req,name=key_field" json:"key_field,omitempty"`
	EventField       *LspMonitorDataEvent    `protobuf:"bytes,2,opt,name=event_field" json:"event_field,omitempty"`
	PropertyField    *LspMonitorDataProperty `protobuf:"bytes,3,opt,name=property_field" json:"property_field,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *LspMon) Reset()                    { *m = LspMon{} }
func (m *LspMon) String() string            { return proto.CompactTextString(m) }
func (*LspMon) ProtoMessage()               {}
func (*LspMon) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LspMon) GetKeyField() *Key {
	if m != nil {
		return m.KeyField
	}
	return nil
}

func (m *LspMon) GetEventField() *LspMonitorDataEvent {
	if m != nil {
		return m.EventField
	}
	return nil
}

func (m *LspMon) GetPropertyField() *LspMonitorDataProperty {
	if m != nil {
		return m.PropertyField
	}
	return nil
}

func init() {
	proto.RegisterType((*Key)(nil), "key")
	proto.RegisterType((*LspMonitorDataEvent)(nil), "lsp_monitor_data_event")
	proto.RegisterType((*EroTypeEntry)(nil), "ero_type_entry")
	proto.RegisterType((*EroIpv4Type)(nil), "ero_ipv4_type")
	proto.RegisterType((*RroTypeEntry)(nil), "rro_type_entry")
	proto.RegisterType((*RroIpv4Type)(nil), "rro_ipv4_type")
	proto.RegisterType((*LspMonitorDataProperty)(nil), "lsp_monitor_data_property")
	proto.RegisterType((*LspMon)(nil), "lsp_mon")
	proto.RegisterEnum("LspEvent", LspEvent_name, LspEvent_value)
	proto.RegisterEnum("EventSubcode", EventSubcode_name, EventSubcode_value)
}

func init() { proto.RegisterFile("lsp_mon.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 845 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xdd, 0x72, 0x9b, 0x46,
	0x14, 0x2e, 0x20, 0xd9, 0xd1, 0x91, 0x91, 0xd7, 0xeb, 0xc4, 0xc6, 0x71, 0xdb, 0x68, 0xd4, 0x5e,
	0x68, 0xd2, 0x8e, 0x3a, 0xa3, 0xf6, 0x05, 0x56, 0xb0, 0xb1, 0x99, 0xc1, 0x40, 0x96, 0x45, 0x6e,
	0xaf, 0x76, 0xb0, 0x59, 0x3b, 0x4c, 0x24, 0xd0, 0x20, 0xea, 0x54, 0xf7, 0x7d, 0x8f, 0x3e, 0x48,
	0x5f, 0xad, 0x17, 0x9d, 0x5d, 0xc9, 0x63, 0xd3, 0x26, 0x77, 0xe2, 0x3b, 0x7b, 0xbe, 0x9f, 0x73,
	0x8e, 0xc0, 0x5e, 0xac, 0x57, 0x62, 0x59, 0x95, 0x93, 0x55, 0x5d, 0x35, 0xd5, 0x88, 0x82, 0xf5,
	0x51, 0x6e, 0xf0, 0x01, 0x74, 0xca, 0x6c, 0x29, 0x1d, 0x63, 0x68, 0x8e, 0x7b, 0xf8, 0x1c, 0x8e,
	0x8b, 0x72, 0xdd, 0x64, 0xe5, 0xad, 0x14, 0x45, 0x2e, 0xcb, 0xa6, 0xb8, 0x2b, 0x64, 0xed, 0x98,
	0x43, 0x73, 0xdc, 0xc5, 0xc7, 0xd0, 0x6f, 0x8a, 0xa5, 0x14, 0xeb, 0x26, 0x5b, 0xae, 0xee, 0x1d,
	0x6b, 0x68, 0x8e, 0x3b, 0x23, 0x01, 0x27, 0x3b, 0xde, 0xa2, 0xa9, 0x6a, 0x91, 0x67, 0x4d, 0x26,
	0xe4, 0x83, 0x2c, 0x1b, 0xfc, 0x3d, 0x20, 0xfd, 0xe3, 0x39, 0x91, 0x52, 0x19, 0x4c, 0x61, 0xa2,
	0x5a, 0xb6, 0xaf, 0xde, 0xc0, 0xfe, 0xfa, 0xf7, 0x9b, 0xdb, 0x2a, 0x97, 0x8e, 0x39, 0x34, 0xc6,
	0x83, 0xe9, 0x60, 0xb2, 0xed, 0xda, 0xa1, 0xa3, 0x1f, 0x60, 0x20, 0xeb, 0x4a, 0x34, 0x9b, 0x95,
	0x14, 0xb2, 0x6c, 0xea, 0x0d, 0x06, 0x30, 0x8b, 0x95, 0xa6, 0xb2, 0xb1, 0x0d, 0xdd, 0xbb, 0x45,
	0x76, 0xbf, 0xd6, 0xcd, 0xbd, 0xd1, 0x4f, 0x60, 0xab, 0xc7, 0xc5, 0xea, 0xe1, 0x17, 0xdd, 0x81,
	0xbf, 0x85, 0xae, 0x6e, 0x72, 0x8c, 0xa1, 0x35, 0xee, 0x4f, 0x0f, 0x27, 0x6d, 0xae, 0xd1, 0x7b,
	0x18, 0xd4, 0x6d, 0xf6, 0x01, 0xec, 0x95, 0x55, 0x2e, 0x8b, 0xdc, 0x31, 0x86, 0xc6, 0x7f, 0x15,
	0x6c, 0x7c, 0x04, 0xbd, 0xa2, 0x6c, 0xee, 0x44, 0x96, 0xe7, 0xb5, 0x63, 0x3d, 0xbe, 0x58, 0x64,
	0x37, 0x72, 0xe1, 0x74, 0xd4, 0xe7, 0xe8, 0x67, 0xb0, 0xeb, 0x96, 0x87, 0x11, 0xf4, 0x14, 0xd0,
	0xf6, 0xd1, 0x56, 0x1d, 0xfd, 0x65, 0xc0, 0xd9, 0xff, 0xe6, 0xb8, 0xaa, 0xab, 0x95, 0xac, 0x9b,
	0x8d, 0x12, 0xbd, 0xc9, 0xca, 0xfc, 0x53, 0x91, 0x37, 0x1f, 0xb4, 0xad, 0x8e, 0x82, 0x56, 0x59,
	0xf3, 0x41, 0xe8, 0xe5, 0xe9, 0xf0, 0xca, 0xf9, 0x52, 0x36, 0x75, 0x71, 0xab, 0x7d, 0x75, 0x31,
	0x06, 0x58, 0x66, 0x7f, 0x88, 0xec, 0xe1, 0x5e, 0xdc, 0x7c, 0xd2, 0xe6, 0x4c, 0x7c, 0x0e, 0x96,
	0xac, 0x2b, 0xa7, 0x3b, 0x34, 0xc6, 0x7d, 0x35, 0xea, 0x96, 0xd1, 0x73, 0xb0, 0xea, 0xba, 0x72,
	0xf6, 0x76, 0xc5, 0x56, 0x8a, 0xd1, 0x9f, 0x06, 0xec, 0xef, 0x1c, 0xe2, 0x53, 0xe8, 0x7d, 0x94,
	0x1b, 0x71, 0x57, 0xc8, 0x45, 0xae, 0x17, 0xd1, 0x9f, 0x76, 0x26, 0xea, 0x9a, 0x7e, 0x84, 0xfe,
	0x76, 0x7b, 0xdb, 0x92, 0xa9, 0x99, 0x4e, 0x27, 0x5f, 0xb8, 0x90, 0x29, 0x0c, 0x1e, 0x23, 0xee,
	0x1a, 0x2c, 0xdd, 0xf0, 0x7a, 0xf2, 0xc5, 0x51, 0xbc, 0xfd, 0xdb, 0x82, 0xde, 0xd3, 0xf5, 0xd8,
	0xd0, 0xf3, 0x43, 0x9f, 0xfb, 0x84, 0x53, 0x0f, 0x7d, 0x85, 0x11, 0x1c, 0xb8, 0x51, 0xe8, 0x06,
	0xa9, 0x47, 0x3d, 0x91, 0xc6, 0xc8, 0xc0, 0xa7, 0x70, 0xfc, 0x84, 0xf0, 0x88, 0x85, 0xc2, 0x8b,
	0xae, 0x43, 0x64, 0x62, 0x07, 0x5e, 0xc6, 0x2c, 0xe2, 0xd4, 0xe5, 0x7e, 0x14, 0x0a, 0x32, 0x27,
	0x7e, 0x40, 0x66, 0x01, 0x45, 0x16, 0x7e, 0x0d, 0x27, 0xcf, 0x2a, 0x69, 0xf8, 0x54, 0xeb, 0x60,
	0x0c, 0x03, 0x92, 0xf2, 0x68, 0x76, 0x2d, 0x92, 0xd4, 0x75, 0x69, 0x92, 0xa0, 0x2e, 0x3e, 0x84,
	0xfe, 0x0e, 0x7b, 0x47, 0xfc, 0x00, 0xed, 0xe1, 0x13, 0xc0, 0x8c, 0x26, 0x73, 0xc1, 0x29, 0x61,
	0x82, 0x51, 0x97, 0xfa, 0x73, 0xea, 0xa1, 0x7d, 0x25, 0xe9, 0xd1, 0x84, 0x06, 0xd4, 0xe5, 0x82,
	0xb8, 0xdc, 0x9f, 0x53, 0x11, 0x13, 0x7e, 0x89, 0x5e, 0xa8, 0x0e, 0xf7, 0x92, 0x84, 0x17, 0xb4,
	0x85, 0xf7, 0x54, 0x3c, 0x8f, 0xf2, 0x28, 0x65, 0x2a, 0x0c, 0x28, 0xa5, 0xdd, 0xa7, 0x0e, 0xd1,
	0xc7, 0x47, 0x60, 0x47, 0xcc, 0xbf, 0xf0, 0x43, 0xc2, 0xa9, 0xb8, 0x9a, 0xcd, 0xd0, 0x81, 0xa2,
	0xfa, 0x8c, 0x84, 0xba, 0x5b, 0xdb, 0x4d, 0xe2, 0x77, 0x22, 0x8c, 0x04, 0x8b, 0x52, 0x4e, 0xd1,
	0x40, 0x4f, 0x4b, 0x41, 0x8f, 0x51, 0x0e, 0xf1, 0x19, 0xbc, 0x62, 0x34, 0xe1, 0x84, 0x71, 0xe5,
	0x3b, 0x9a, 0x53, 0xf6, 0xdb, 0x36, 0x14, 0xc2, 0x2f, 0x01, 0x29, 0x26, 0xca, 0x9e, 0x45, 0x3a,
	0xc2, 0xc7, 0x70, 0xa8, 0x50, 0x71, 0xc5, 0x53, 0xb1, 0x4d, 0x80, 0xb0, 0x62, 0xe1, 0x69, 0x18,
	0xd2, 0x40, 0x04, 0x91, 0x4b, 0x02, 0xc1, 0x68, 0x4c, 0x7c, 0x46, 0x3d, 0x74, 0xfc, 0xf6, 0x1f,
	0x03, 0xec, 0xd6, 0xdf, 0x1b, 0x7f, 0x03, 0x67, 0xc4, 0xbb, 0xf2, 0x93, 0x44, 0x0d, 0xdb, 0x8d,
	0x42, 0xce, 0xa2, 0x40, 0x6b, 0xa6, 0x8c, 0x22, 0x03, 0xbf, 0x82, 0xa3, 0x84, 0x6e, 0x8b, 0x31,
	0xa3, 0xf4, 0x2a, 0x56, 0x8b, 0x36, 0x95, 0xee, 0x8c, 0x78, 0x22, 0x88, 0xa2, 0x84, 0xee, 0xf2,
	0x58, 0xca, 0xa2, 0x02, 0x13, 0xce, 0x7c, 0x97, 0xef, 0xd0, 0x0e, 0xfe, 0x1a, 0x9c, 0x80, 0xcc,
	0x68, 0x20, 0x48, 0xa0, 0xec, 0xe8, 0xa5, 0x3e, 0xf2, 0x77, 0xf1, 0x39, 0x9c, 0x86, 0x51, 0x28,
	0x58, 0x32, 0x8f, 0x85, 0x4b, 0x62, 0xb5, 0xe6, 0x6d, 0x27, 0x43, 0x7b, 0x6a, 0xde, 0x9c, 0x07,
	0x82, 0xfe, 0x1a, 0x6b, 0xfb, 0xfb, 0x7a, 0x3e, 0x51, 0xca, 0xfd, 0xf0, 0x42, 0x49, 0xc7, 0xc2,
	0xa3, 0xea, 0x4c, 0xa8, 0x87, 0x5e, 0xe0, 0xef, 0xe0, 0x0d, 0xa3, 0xef, 0x53, 0x9a, 0x70, 0xea,
	0x89, 0x19, 0x09, 0xbd, 0x6b, 0xdf, 0xe3, 0x97, 0xad, 0xf3, 0xe9, 0xcd, 0xcc, 0x4b, 0xeb, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x40, 0x91, 0x42, 0x87, 0x05, 0x00, 0x00,
}
