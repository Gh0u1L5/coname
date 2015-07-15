// Code generated by protoc-gen-gogo.
// source: replication.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

// KeyserverStep denotes the input to a single step of the keyserver state
// machine. Linearizable high-availability replication is achieved by
// replicating an in-order log of all steps and having each replica reproduce
// the state from them.
type KeyserverStep struct {
	UID uint64 `protobuf:"varint,1,opt,proto3" json:"UID,omitempty"`
	// update is appended to the log when it is received from a client and
	// has passed pre-validation. However, since pre-validation is not
	// final, "success" should not be returned to the client until after the
	// update has been applied and ratified.
	// update is applied to the keyserver state as soon as it has been
	// committed to the log.
	Update *SignedEntryUpdate `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
	// epoch_delimiter is appended to the log by a leader (not necessarily
	// unique) node when at least the duration EPOCH_INTERVAL_MIN and at
	// most EPOCH_INTERVAL_MAX after the previous epoch_delimiter has passed.
	// Between these times, an epoch delimiter is appended as soon as an
	// update is committed.
	// As the leader requirement for appending an epoch_delimiter is soft,
	// it may happen that an epoch delimiter with a epoch number not higher
	// than the previous gets committed to the log. It must be ignored.
	EpochDelimiter *EpochDelimiter `protobuf:"bytes,3,opt,name=epoch_delimiter" json:"epoch_delimiter,omitempty"`
	// replica_ratification for the last epoch is appended to the log
	// when the epoch_delimiter is committed.
	// After some majority of the replicas has ratified the state, their
	// signatures make up the keyserver signature. As combining signatures
	// is deterministic, no new log entry is appended to record that.
	ReplicaRatification *SignedRatification `protobuf:"bytes,4,opt,name=replica_ratification" json:"replica_ratification,omitempty"`
	// verifier_ratification is appended for each new ratification received
	// from a verifier; these are used to provide proof of verification to
	// clients.
	VerifierRatification *SignedRatification `protobuf:"bytes,5,opt,name=verifier_ratification" json:"verifier_ratification,omitempty"`
}

func (m *KeyserverStep) Reset()         { *m = KeyserverStep{} }
func (m *KeyserverStep) String() string { return proto1.CompactTextString(m) }
func (*KeyserverStep) ProtoMessage()    {}

func (m *KeyserverStep) GetUpdate() *SignedEntryUpdate {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *KeyserverStep) GetEpochDelimiter() *EpochDelimiter {
	if m != nil {
		return m.EpochDelimiter
	}
	return nil
}

func (m *KeyserverStep) GetReplicaRatification() *SignedRatification {
	if m != nil {
		return m.ReplicaRatification
	}
	return nil
}

func (m *KeyserverStep) GetVerifierRatification() *SignedRatification {
	if m != nil {
		return m.VerifierRatification
	}
	return nil
}

type EpochDelimiter struct {
	EpochNumber uint64 `protobuf:"varint,1,opt,name=epoch_number,proto3" json:"epoch_number,omitempty"`
	Timestamp   Time   `protobuf:"bytes,4,opt,name=timestamp,customtype=Time" json:"timestamp"`
}

func (m *EpochDelimiter) Reset()         { *m = EpochDelimiter{} }
func (m *EpochDelimiter) String() string { return proto1.CompactTextString(m) }
func (*EpochDelimiter) ProtoMessage()    {}

func init() {
}
func (m *KeyserverStep) Unmarshal(data []byte) error {
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
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.UID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Update == nil {
				m.Update = &SignedEntryUpdate{}
			}
			if err := m.Update.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochDelimiter", wireType)
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
			if m.EpochDelimiter == nil {
				m.EpochDelimiter = &EpochDelimiter{}
			}
			if err := m.EpochDelimiter.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaRatification", wireType)
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
			if m.ReplicaRatification == nil {
				m.ReplicaRatification = &SignedRatification{}
			}
			if err := m.ReplicaRatification.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifierRatification", wireType)
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
			if m.VerifierRatification == nil {
				m.VerifierRatification = &SignedRatification{}
			}
			if err := m.VerifierRatification.Unmarshal(data[iNdEx:postIndex]); err != nil {
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
			skippy, err := skipReplication(data[iNdEx:])
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
func (m *EpochDelimiter) Unmarshal(data []byte) error {
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
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.EpochNumber |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
			if err := m.Timestamp.Unmarshal(data[iNdEx:postIndex]); err != nil {
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
			skippy, err := skipReplication(data[iNdEx:])
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
func skipReplication(data []byte) (n int, err error) {
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
				next, err := skipReplication(data[start:])
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
func (m *KeyserverStep) Size() (n int) {
	var l int
	_ = l
	if m.UID != 0 {
		n += 1 + sovReplication(uint64(m.UID))
	}
	if m.Update != nil {
		l = m.Update.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	if m.EpochDelimiter != nil {
		l = m.EpochDelimiter.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	if m.ReplicaRatification != nil {
		l = m.ReplicaRatification.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	if m.VerifierRatification != nil {
		l = m.VerifierRatification.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	return n
}

func (m *EpochDelimiter) Size() (n int) {
	var l int
	_ = l
	if m.EpochNumber != 0 {
		n += 1 + sovReplication(uint64(m.EpochNumber))
	}
	l = m.Timestamp.Size()
	n += 1 + l + sovReplication(uint64(l))
	return n
}

func sovReplication(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozReplication(x uint64) (n int) {
	return sovReplication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyserverStep) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *KeyserverStep) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintReplication(data, i, uint64(m.UID))
	}
	if m.Update != nil {
		data[i] = 0x12
		i++
		i = encodeVarintReplication(data, i, uint64(m.Update.Size()))
		n1, err := m.Update.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.EpochDelimiter != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintReplication(data, i, uint64(m.EpochDelimiter.Size()))
		n2, err := m.EpochDelimiter.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.ReplicaRatification != nil {
		data[i] = 0x22
		i++
		i = encodeVarintReplication(data, i, uint64(m.ReplicaRatification.Size()))
		n3, err := m.ReplicaRatification.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.VerifierRatification != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintReplication(data, i, uint64(m.VerifierRatification.Size()))
		n4, err := m.VerifierRatification.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *EpochDelimiter) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EpochDelimiter) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintReplication(data, i, uint64(m.EpochNumber))
	}
	data[i] = 0x22
	i++
	i = encodeVarintReplication(data, i, uint64(m.Timestamp.Size()))
	n5, err := m.Timestamp.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	return i, nil
}

func encodeFixed64Replication(data []byte, offset int, v uint64) int {
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
func encodeFixed32Replication(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintReplication(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
