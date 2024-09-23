package hash

import "encoding/binary"

/*
MurmurHash3 is a fast, non-cryptographic hash function suitable for general hash-based lookup operations.
It was developed by Austin Appleby and is widely used in software applications where fast hashing is needed, such as hash tables, bloom filters, or data deduplication tasks.
The algorithm is called "Murmur" because it produces non-randomized, but still "scrambled" hash results, ensuring good distribution of keys for hashing.
*/
type Murmur3 struct {
	seed      uint32
	h1        uint32
	length    int
	tail      []byte
	blockSize int
}

func NewMurmur3() *Murmur3 {
	return NewMurmur3WithSeed(0)
}

func NewMurmur3WithSeed(seed uint32) *Murmur3 {
	return &Murmur3{
		seed:      seed,
		h1:        seed,
		blockSize: 4,
	}
}

// mixBlock mixes the current block into the hash state
func (m *Murmur3) mixBlock(h1, k1 uint32) uint32 {
	k1 *= 0xcc9e2d51
	k1 = (k1 << 15) | (k1 >> (32 - 15))
	k1 *= 0x1b873593

	h1 ^= k1
	h1 = (h1 << 13) | (h1 >> (32 - 13))
	h1 = h1*5 + 0xe6546b64

	return h1
}

// Write adds more data to the running hash (via the embedded io.Writer interface)
func (m *Murmur3) Write(data []byte) (int, error) {
	m.length += len(data)

	// Append tail data from the last call to form a full block
	if len(m.tail) > 0 {
		data = append(m.tail, data...)
		m.tail = nil
	}

	// Process full blocks
	nblocks := len(data) / 4
	for i := 0; i < nblocks; i++ {
		k1 := binary.LittleEndian.Uint32(data[i*4 : (i+1)*4])
		m.h1 = m.mixBlock(m.h1, k1)
	}

	// Store remaining bytes (tail) for next call
	if len(data)%4 != 0 {
		m.tail = append([]byte{}, data[nblocks*4:]...)
	}

	return len(data), nil
}

// BlockSize returns the hash's underlying block size
func (m *Murmur3) BlockSize() int {
	return m.blockSize
}

// Size returns the number of bytes Sum will return
func (m *Murmur3) Size() int {
	return 4
}

// Reset resets the hash to its initial state
func (m *Murmur3) Reset() {
	m.h1 = m.seed
	m.length = 0
	m.tail = nil
}

// Sum appends the current hash to b and returns the resulting slice
func (m *Murmur3) Sum(b []byte) []byte {
	h1 := m.h1

	// Process remaining tail bytes
	var k1 uint32
	switch len(m.tail) {
	case 3:
		k1 ^= uint32(m.tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(m.tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(m.tail[0])
		k1 *= 0xcc9e2d51
		k1 = (k1 << 15) | (k1 >> (32 - 15))
		k1 *= 0x1b873593
		h1 ^= k1
	}

	// Finalize the hash
	h1 ^= uint32(m.length)
	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	// Append the result to the provided byte slice
	result := make([]byte, 4)
	binary.LittleEndian.PutUint32(result, h1)
	return append(b, result...)
}
