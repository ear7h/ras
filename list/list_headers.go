package list

import (
	"fmt"
	"encoding/binary"
)

// block header
const blockHeaderLen = 8

// 2 block headers, used when allocating,
const blockHeaderLen2 = 2 * blockHeaderLen

// This is not used but defines the block header
type BlockHeaderScheme struct {
	Len, Next uint32
}

type blockHeader struct {
	len, next uint
}

func (hdr *blockHeader) UnmarshalBinary(data []byte) error {
	if len(data) != blockHeaderLen {
		return fmt.Errorf("block header must be taken form %d bytes", blockHeaderLen)
	}

	hdr.len = uint(binary.LittleEndian.Uint32(data[:4]))
	hdr.next = uint(binary.LittleEndian.Uint32(data[4:]))

	return nil
}

func (hdr blockHeader) MarshalBinary() []byte {
	ret := make([]byte, blockHeaderLen)

	binary.LittleEndian.PutUint32(ret[:4], uint32(hdr.len))
	binary.LittleEndian.PutUint32(ret[4:], uint32(hdr.next))

	return ret
}



const magicOffset = 0
const magicLen = 2
const headPtrOffset = magicOffset + magicLen
const headPtrLen = 4

const usableOffset = headPtrOffset + headPtrLen


var Magic = [magicLen]byte{0xe7, 0x89}

// This is not used but defines the list header
type ListHeaderScheme struct {
	Magic [magicLen]byte
	HeadPtr uint32
	// version or allocer enum
	// this would differentiate between linked list
	// and other kinds of allocers
}
