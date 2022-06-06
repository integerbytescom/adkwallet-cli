package modules

// TOOLSET TO read FireFox Snappy Containers. Required to connect to Firefox Extension/Vault.

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/golang/snappy"
	"io"
)

var boutput *bufio.Writer

func DecompressSnappyFox(binput *bufio.Reader) ([]byte, error) {
	hbuf := bytes.Buffer{}
	boutput = bufio.NewWriter(&hbuf)

	SnappyDecompressFrame(binput, boutput)
	boutput.Flush()
	return hbuf.Bytes(), nil
}

func SnappyDecompressFrame(binput *bufio.Reader, boutput *bufio.Writer) {
	for {
		chunktype, err := binput.ReadByte()
		if err != nil {
			break
		}
		reterr := parse_chunk(binput, boutput, chunktype)
		if reterr != nil {
			break
		}
	}
}

func parse_chunk(binput *bufio.Reader, boutput *bufio.Writer, chunktype byte) error {
	var err_stop error
	err_stop = nil
	switch chunktype {
	case 0xff:
		return parse_stream_identifier(binput)
	case 0x00:
		return parse_compressed_data_chunk(binput, boutput)
	case 0x01:
		/* TODO */
		// return parse_uncompressed_data_chunk(in, out);
		err_stop = errors.New("not implemented A")
	case 0xfe:
		/* TODO */
		// return parse_padding(in, out);
		err_stop = errors.New("not implemented B")
	default:
		//return parse_unknown_chunktype(chunktype);
		err_stop = errors.New("not implemented C")
	}
	return err_stop
}

func parse_stream_identifier(binput *bufio.Reader) error {
	// sNaPpY
	checkSig := []byte{0x06, 0x00, 0x00, 0x73, 0x4e, 0x61, 0x50, 0x70, 0x59}
	bcheck := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	_, err := io.ReadFull(binput, bcheck)

	if err != nil || bytes.Compare(checkSig, bcheck) != 0 {
		return errors.New("invalid stream identifier signature: sNaPpY not found")
	}

	return nil
}

const MAX_COMPRESSED_DATA_SIZE = 16777211
const MAX_UNCOMPRESSED_DATA_SIZE = 65536

func parse_compressed_data_chunk(binput *bufio.Reader, boutput *bufio.Writer) error {
	//c_data := make([]byte, MAX_COMPRESSED_DATA_SIZE)
	//data := make([]byte, MAX_UNCOMPRESSED_DATA_SIZE)
	c_len_raw := []byte{0x00, 0x00, 0x00, 0x00}

	c_length_b, errA := readBytesExact(binput, 3)
	if errA != nil {
		return errA
	}
	c_len_raw[0] = c_length_b[0]
	c_len_raw[1] = c_length_b[1]
	c_len_raw[2] = c_length_b[2]

	c_crc_b, errB := readBytesExact(binput, 4)
	if errB != nil {
		return errB
	}

	c_length := binary.LittleEndian.Uint32(c_len_raw)
	//c_crc := binary.LittleEndian.Uint32(c_crc_b)
	_ = binary.LittleEndian.Uint32(c_crc_b) //crc
	if c_length > MAX_COMPRESSED_DATA_SIZE {
		return errors.New("c_length > MAX_COMPRESSED_DATA_SIZE")
	}
	c_length--

	chunk, errC := readBytesExact(binput, int(c_length-3))
	if errC != nil {
		return errC
	}

	return snappy_uncompress(chunk, boutput)

}

func snappy_uncompress(data []byte, boutput *bufio.Writer) error { // uncompress chunk
	decoded, err := snappy.Decode(nil, data)
	if err != nil {
		return err
	} else {
		boutput.Write(decoded)
	}
	return nil
}

func readBytesExact(binput *bufio.Reader, cnt int) ([]byte, error) {
	bb := make([]byte, cnt)
	_, err := io.ReadFull(binput, bb)

	if err != nil {
		return bb, errors.New("end of stream A")
	}

	return bb, nil

}
