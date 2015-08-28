// Code generated by protoc-gen-gogo.
// source: desired_lrp_requests.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

import fmt "fmt"
import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type DesiredLRPLifecycleResponse struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *DesiredLRPLifecycleResponse) Reset()      { *m = DesiredLRPLifecycleResponse{} }
func (*DesiredLRPLifecycleResponse) ProtoMessage() {}

func (m *DesiredLRPLifecycleResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type DesiredLRPsResponse struct {
	Error       *Error        `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	DesiredLrps []*DesiredLRP `protobuf:"bytes,2,rep,name=desired_lrps" json:"desired_lrps,omitempty"`
}

func (m *DesiredLRPsResponse) Reset()      { *m = DesiredLRPsResponse{} }
func (*DesiredLRPsResponse) ProtoMessage() {}

func (m *DesiredLRPsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DesiredLRPsResponse) GetDesiredLrps() []*DesiredLRP {
	if m != nil {
		return m.DesiredLrps
	}
	return nil
}

type DesiredLRPsRequest struct {
	Domain string `protobuf:"bytes,1,opt,name=domain" json:"domain"`
}

func (m *DesiredLRPsRequest) Reset()      { *m = DesiredLRPsRequest{} }
func (*DesiredLRPsRequest) ProtoMessage() {}

func (m *DesiredLRPsRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type DesiredLRPResponse struct {
	Error      *Error      `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	DesiredLrp *DesiredLRP `protobuf:"bytes,2,opt,name=desired_lrp" json:"desired_lrp,omitempty"`
}

func (m *DesiredLRPResponse) Reset()      { *m = DesiredLRPResponse{} }
func (*DesiredLRPResponse) ProtoMessage() {}

func (m *DesiredLRPResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DesiredLRPResponse) GetDesiredLrp() *DesiredLRP {
	if m != nil {
		return m.DesiredLrp
	}
	return nil
}

type DesiredLRPByProcessGuidRequest struct {
	ProcessGuid string `protobuf:"bytes,1,opt,name=process_guid" json:"process_guid"`
}

func (m *DesiredLRPByProcessGuidRequest) Reset()      { *m = DesiredLRPByProcessGuidRequest{} }
func (*DesiredLRPByProcessGuidRequest) ProtoMessage() {}

func (m *DesiredLRPByProcessGuidRequest) GetProcessGuid() string {
	if m != nil {
		return m.ProcessGuid
	}
	return ""
}

type DesireLRPRequest struct {
	DesiredLrp *DesiredLRP `protobuf:"bytes,1,opt,name=desired_lrp" json:"desired_lrp,omitempty"`
}

func (m *DesireLRPRequest) Reset()      { *m = DesireLRPRequest{} }
func (*DesireLRPRequest) ProtoMessage() {}

func (m *DesireLRPRequest) GetDesiredLrp() *DesiredLRP {
	if m != nil {
		return m.DesiredLrp
	}
	return nil
}

type UpdateDesiredLRPRequest struct {
	ProcessGuid string            `protobuf:"bytes,1,opt,name=process_guid" json:"process_guid"`
	Update      *DesiredLRPUpdate `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
}

func (m *UpdateDesiredLRPRequest) Reset()      { *m = UpdateDesiredLRPRequest{} }
func (*UpdateDesiredLRPRequest) ProtoMessage() {}

func (m *UpdateDesiredLRPRequest) GetProcessGuid() string {
	if m != nil {
		return m.ProcessGuid
	}
	return ""
}

func (m *UpdateDesiredLRPRequest) GetUpdate() *DesiredLRPUpdate {
	if m != nil {
		return m.Update
	}
	return nil
}

type RemoveDesiredLRPRequest struct {
	ProcessGuid string `protobuf:"bytes,1,opt,name=process_guid" json:"process_guid"`
}

func (m *RemoveDesiredLRPRequest) Reset()      { *m = RemoveDesiredLRPRequest{} }
func (*RemoveDesiredLRPRequest) ProtoMessage() {}

func (m *RemoveDesiredLRPRequest) GetProcessGuid() string {
	if m != nil {
		return m.ProcessGuid
	}
	return ""
}

func (this *DesiredLRPLifecycleResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DesiredLRPLifecycleResponse)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Error.Equal(that1.Error) {
		return false
	}
	return true
}
func (this *DesiredLRPsResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DesiredLRPsResponse)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Error.Equal(that1.Error) {
		return false
	}
	if len(this.DesiredLrps) != len(that1.DesiredLrps) {
		return false
	}
	for i := range this.DesiredLrps {
		if !this.DesiredLrps[i].Equal(that1.DesiredLrps[i]) {
			return false
		}
	}
	return true
}
func (this *DesiredLRPsRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DesiredLRPsRequest)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Domain != that1.Domain {
		return false
	}
	return true
}
func (this *DesiredLRPResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DesiredLRPResponse)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Error.Equal(that1.Error) {
		return false
	}
	if !this.DesiredLrp.Equal(that1.DesiredLrp) {
		return false
	}
	return true
}
func (this *DesiredLRPByProcessGuidRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DesiredLRPByProcessGuidRequest)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ProcessGuid != that1.ProcessGuid {
		return false
	}
	return true
}
func (this *DesireLRPRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DesireLRPRequest)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.DesiredLrp.Equal(that1.DesiredLrp) {
		return false
	}
	return true
}
func (this *UpdateDesiredLRPRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*UpdateDesiredLRPRequest)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ProcessGuid != that1.ProcessGuid {
		return false
	}
	if !this.Update.Equal(that1.Update) {
		return false
	}
	return true
}
func (this *RemoveDesiredLRPRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RemoveDesiredLRPRequest)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ProcessGuid != that1.ProcessGuid {
		return false
	}
	return true
}
func (this *DesiredLRPLifecycleResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.DesiredLRPLifecycleResponse{` +
		`Error:` + fmt.Sprintf("%#v", this.Error) + `}`}, ", ")
	return s
}
func (this *DesiredLRPsResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.DesiredLRPsResponse{` +
		`Error:` + fmt.Sprintf("%#v", this.Error),
		`DesiredLrps:` + fmt.Sprintf("%#v", this.DesiredLrps) + `}`}, ", ")
	return s
}
func (this *DesiredLRPsRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.DesiredLRPsRequest{` +
		`Domain:` + fmt.Sprintf("%#v", this.Domain) + `}`}, ", ")
	return s
}
func (this *DesiredLRPResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.DesiredLRPResponse{` +
		`Error:` + fmt.Sprintf("%#v", this.Error),
		`DesiredLrp:` + fmt.Sprintf("%#v", this.DesiredLrp) + `}`}, ", ")
	return s
}
func (this *DesiredLRPByProcessGuidRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.DesiredLRPByProcessGuidRequest{` +
		`ProcessGuid:` + fmt.Sprintf("%#v", this.ProcessGuid) + `}`}, ", ")
	return s
}
func (this *DesireLRPRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.DesireLRPRequest{` +
		`DesiredLrp:` + fmt.Sprintf("%#v", this.DesiredLrp) + `}`}, ", ")
	return s
}
func (this *UpdateDesiredLRPRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.UpdateDesiredLRPRequest{` +
		`ProcessGuid:` + fmt.Sprintf("%#v", this.ProcessGuid),
		`Update:` + fmt.Sprintf("%#v", this.Update) + `}`}, ", ")
	return s
}
func (this *RemoveDesiredLRPRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.RemoveDesiredLRPRequest{` +
		`ProcessGuid:` + fmt.Sprintf("%#v", this.ProcessGuid) + `}`}, ", ")
	return s
}
func valueToGoStringDesiredLrpRequests(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringDesiredLrpRequests(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *DesiredLRPLifecycleResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DesiredLRPLifecycleResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		data[i] = 0xa
		i++
		i = encodeVarintDesiredLrpRequests(data, i, uint64(m.Error.Size()))
		n1, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *DesiredLRPsResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DesiredLRPsResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		data[i] = 0xa
		i++
		i = encodeVarintDesiredLrpRequests(data, i, uint64(m.Error.Size()))
		n2, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.DesiredLrps) > 0 {
		for _, msg := range m.DesiredLrps {
			data[i] = 0x12
			i++
			i = encodeVarintDesiredLrpRequests(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *DesiredLRPsRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DesiredLRPsRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintDesiredLrpRequests(data, i, uint64(len(m.Domain)))
	i += copy(data[i:], m.Domain)
	return i, nil
}

func (m *DesiredLRPResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DesiredLRPResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		data[i] = 0xa
		i++
		i = encodeVarintDesiredLrpRequests(data, i, uint64(m.Error.Size()))
		n3, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.DesiredLrp != nil {
		data[i] = 0x12
		i++
		i = encodeVarintDesiredLrpRequests(data, i, uint64(m.DesiredLrp.Size()))
		n4, err := m.DesiredLrp.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *DesiredLRPByProcessGuidRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DesiredLRPByProcessGuidRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintDesiredLrpRequests(data, i, uint64(len(m.ProcessGuid)))
	i += copy(data[i:], m.ProcessGuid)
	return i, nil
}

func (m *DesireLRPRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DesireLRPRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DesiredLrp != nil {
		data[i] = 0xa
		i++
		i = encodeVarintDesiredLrpRequests(data, i, uint64(m.DesiredLrp.Size()))
		n5, err := m.DesiredLrp.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *UpdateDesiredLRPRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *UpdateDesiredLRPRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintDesiredLrpRequests(data, i, uint64(len(m.ProcessGuid)))
	i += copy(data[i:], m.ProcessGuid)
	if m.Update != nil {
		data[i] = 0x12
		i++
		i = encodeVarintDesiredLrpRequests(data, i, uint64(m.Update.Size()))
		n6, err := m.Update.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *RemoveDesiredLRPRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RemoveDesiredLRPRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintDesiredLrpRequests(data, i, uint64(len(m.ProcessGuid)))
	i += copy(data[i:], m.ProcessGuid)
	return i, nil
}

func encodeFixed64DesiredLrpRequests(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32DesiredLrpRequests(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDesiredLrpRequests(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *DesiredLRPLifecycleResponse) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovDesiredLrpRequests(uint64(l))
	}
	return n
}

func (m *DesiredLRPsResponse) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovDesiredLrpRequests(uint64(l))
	}
	if len(m.DesiredLrps) > 0 {
		for _, e := range m.DesiredLrps {
			l = e.Size()
			n += 1 + l + sovDesiredLrpRequests(uint64(l))
		}
	}
	return n
}

func (m *DesiredLRPsRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Domain)
	n += 1 + l + sovDesiredLrpRequests(uint64(l))
	return n
}

func (m *DesiredLRPResponse) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovDesiredLrpRequests(uint64(l))
	}
	if m.DesiredLrp != nil {
		l = m.DesiredLrp.Size()
		n += 1 + l + sovDesiredLrpRequests(uint64(l))
	}
	return n
}

func (m *DesiredLRPByProcessGuidRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProcessGuid)
	n += 1 + l + sovDesiredLrpRequests(uint64(l))
	return n
}

func (m *DesireLRPRequest) Size() (n int) {
	var l int
	_ = l
	if m.DesiredLrp != nil {
		l = m.DesiredLrp.Size()
		n += 1 + l + sovDesiredLrpRequests(uint64(l))
	}
	return n
}

func (m *UpdateDesiredLRPRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProcessGuid)
	n += 1 + l + sovDesiredLrpRequests(uint64(l))
	if m.Update != nil {
		l = m.Update.Size()
		n += 1 + l + sovDesiredLrpRequests(uint64(l))
	}
	return n
}

func (m *RemoveDesiredLRPRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProcessGuid)
	n += 1 + l + sovDesiredLrpRequests(uint64(l))
	return n
}

func sovDesiredLrpRequests(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDesiredLrpRequests(x uint64) (n int) {
	return sovDesiredLrpRequests(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DesiredLRPLifecycleResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DesiredLRPLifecycleResponse{`,
		`Error:` + strings.Replace(fmt.Sprintf("%v", this.Error), "Error", "Error", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DesiredLRPsResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DesiredLRPsResponse{`,
		`Error:` + strings.Replace(fmt.Sprintf("%v", this.Error), "Error", "Error", 1) + `,`,
		`DesiredLrps:` + strings.Replace(fmt.Sprintf("%v", this.DesiredLrps), "DesiredLRP", "DesiredLRP", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DesiredLRPsRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DesiredLRPsRequest{`,
		`Domain:` + fmt.Sprintf("%v", this.Domain) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DesiredLRPResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DesiredLRPResponse{`,
		`Error:` + strings.Replace(fmt.Sprintf("%v", this.Error), "Error", "Error", 1) + `,`,
		`DesiredLrp:` + strings.Replace(fmt.Sprintf("%v", this.DesiredLrp), "DesiredLRP", "DesiredLRP", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DesiredLRPByProcessGuidRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DesiredLRPByProcessGuidRequest{`,
		`ProcessGuid:` + fmt.Sprintf("%v", this.ProcessGuid) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DesireLRPRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DesireLRPRequest{`,
		`DesiredLrp:` + strings.Replace(fmt.Sprintf("%v", this.DesiredLrp), "DesiredLRP", "DesiredLRP", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UpdateDesiredLRPRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpdateDesiredLRPRequest{`,
		`ProcessGuid:` + fmt.Sprintf("%v", this.ProcessGuid) + `,`,
		`Update:` + strings.Replace(fmt.Sprintf("%v", this.Update), "DesiredLRPUpdate", "DesiredLRPUpdate", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RemoveDesiredLRPRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RemoveDesiredLRPRequest{`,
		`ProcessGuid:` + fmt.Sprintf("%v", this.ProcessGuid) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDesiredLrpRequests(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DesiredLRPLifecycleResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipDesiredLrpRequests(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *DesiredLRPsResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DesiredLrps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DesiredLrps = append(m.DesiredLrps, &DesiredLRP{})
			if err := m.DesiredLrps[len(m.DesiredLrps)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipDesiredLrpRequests(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *DesiredLRPsRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if stringLen < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipDesiredLrpRequests(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *DesiredLRPResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DesiredLrp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DesiredLrp == nil {
				m.DesiredLrp = &DesiredLRP{}
			}
			if err := m.DesiredLrp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipDesiredLrpRequests(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *DesiredLRPByProcessGuidRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessGuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if stringLen < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProcessGuid = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipDesiredLrpRequests(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *DesireLRPRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DesiredLrp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DesiredLrp == nil {
				m.DesiredLrp = &DesiredLRP{}
			}
			if err := m.DesiredLrp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipDesiredLrpRequests(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *UpdateDesiredLRPRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessGuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if stringLen < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProcessGuid = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Update", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Update == nil {
				m.Update = &DesiredLRPUpdate{}
			}
			if err := m.Update.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipDesiredLrpRequests(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *RemoveDesiredLRPRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessGuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if stringLen < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProcessGuid = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipDesiredLrpRequests(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDesiredLrpRequests
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipDesiredLrpRequests(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDesiredLrpRequests
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipDesiredLrpRequests(data[start:])
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
	ErrInvalidLengthDesiredLrpRequests = fmt.Errorf("proto: negative length found during unmarshaling")
)
