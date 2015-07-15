// Code generated by protoc-gen-gogo.
// source: local.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

// ReplicaState contains the persistent internal state of a single replica.
type ReplicaState struct {
	NextIndexLog        uint64         `protobuf:"varint,1,opt,name=next_index_log,proto3" json:"next_index_log,omitempty"`
	NextIndexVerifier   uint64         `protobuf:"varint,2,opt,name=next_index_verifier,proto3" json:"next_index_verifier,omitempty"`
	PreviousSummaryHash []byte         `protobuf:"bytes,3,opt,name=previous_summary_hash,proto3" json:"previous_summary_hash,omitempty"`
	LastEpochDelimiter  EpochDelimiter `protobuf:"bytes,4,opt,name=last_epoch_delimiter" json:"last_epoch_delimiter"`
	PendingUpdates      bool           `protobuf:"varint,5,opt,name=pending_updates,proto3" json:"pending_updates,omitempty"`
}

func (m *ReplicaState) Reset()         { *m = ReplicaState{} }
func (m *ReplicaState) String() string { return proto1.CompactTextString(m) }
func (*ReplicaState) ProtoMessage()    {}

func (m *ReplicaState) GetLastEpochDelimiter() EpochDelimiter {
	if m != nil {
		return m.LastEpochDelimiter
	}
	return EpochDelimiter{}
}

// Verifier contains the persistent internal state of a verifier.
type VerifierState struct {
	NextIndex           uint64 `protobuf:"varint,1,opt,name=next_index,proto3" json:"next_index,omitempty"`
	NextEpoch           uint64 `protobuf:"varint,2,opt,name=next_epoch,proto3" json:"next_epoch,omitempty"`
	PreviousSummaryHash []byte `protobuf:"bytes,3,opt,name=previous_summary_hash,proto3" json:"previous_summary_hash,omitempty"`
}

func (m *VerifierState) Reset()         { *m = VerifierState{} }
func (m *VerifierState) String() string { return proto1.CompactTextString(m) }
func (*VerifierState) ProtoMessage()    {}

func init() {
}
func (m *ReplicaState) Unmarshal(data []byte) error {
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextIndexLog", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NextIndexLog |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextIndexVerifier", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NextIndexVerifier |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousSummaryHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousSummaryHash = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastEpochDelimiter", wireType)
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
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastEpochDelimiter.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingUpdates", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PendingUpdates = bool(v != 0)
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
			skippy, err := skipLocal(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *VerifierState) Unmarshal(data []byte) error {
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextIndex", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NextIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextEpoch", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NextEpoch |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousSummaryHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousSummaryHash = append([]byte{}, data[iNdEx:postIndex]...)
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
			skippy, err := skipLocal(data[iNdEx:])
			if err != nil {
				return err
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipLocal(data []byte) (n int, err error) {
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
				next, err := skipLocal(data[start:])
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
func (m *ReplicaState) Size() (n int) {
	var l int
	_ = l
	if m.NextIndexLog != 0 {
		n += 1 + sovLocal(uint64(m.NextIndexLog))
	}
	if m.NextIndexVerifier != 0 {
		n += 1 + sovLocal(uint64(m.NextIndexVerifier))
	}
	if m.PreviousSummaryHash != nil {
		l = len(m.PreviousSummaryHash)
		if l > 0 {
			n += 1 + l + sovLocal(uint64(l))
		}
	}
	l = m.LastEpochDelimiter.Size()
	n += 1 + l + sovLocal(uint64(l))
	if m.PendingUpdates {
		n += 2
	}
	return n
}

func (m *VerifierState) Size() (n int) {
	var l int
	_ = l
	if m.NextIndex != 0 {
		n += 1 + sovLocal(uint64(m.NextIndex))
	}
	if m.NextEpoch != 0 {
		n += 1 + sovLocal(uint64(m.NextEpoch))
	}
	if m.PreviousSummaryHash != nil {
		l = len(m.PreviousSummaryHash)
		if l > 0 {
			n += 1 + l + sovLocal(uint64(l))
		}
	}
	return n
}

func sovLocal(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLocal(x uint64) (n int) {
	return sovLocal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReplicaState) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ReplicaState) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NextIndexLog != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintLocal(data, i, uint64(m.NextIndexLog))
	}
	if m.NextIndexVerifier != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintLocal(data, i, uint64(m.NextIndexVerifier))
	}
	if m.PreviousSummaryHash != nil {
		if len(m.PreviousSummaryHash) > 0 {
			data[i] = 0x1a
			i++
			i = encodeVarintLocal(data, i, uint64(len(m.PreviousSummaryHash)))
			i += copy(data[i:], m.PreviousSummaryHash)
		}
	}
	data[i] = 0x22
	i++
	i = encodeVarintLocal(data, i, uint64(m.LastEpochDelimiter.Size()))
	n1, err := m.LastEpochDelimiter.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.PendingUpdates {
		data[i] = 0x28
		i++
		if m.PendingUpdates {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *VerifierState) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *VerifierState) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NextIndex != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintLocal(data, i, uint64(m.NextIndex))
	}
	if m.NextEpoch != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintLocal(data, i, uint64(m.NextEpoch))
	}
	if m.PreviousSummaryHash != nil {
		if len(m.PreviousSummaryHash) > 0 {
			data[i] = 0x1a
			i++
			i = encodeVarintLocal(data, i, uint64(len(m.PreviousSummaryHash)))
			i += copy(data[i:], m.PreviousSummaryHash)
		}
	}
	return i, nil
}

func encodeFixed64Local(data []byte, offset int, v uint64) int {
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
func encodeFixed32Local(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLocal(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
