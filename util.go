package sestringgo

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/OTCompa/sestring-go/ffxiv"
)

func getInteger(bytes *bytes.Reader) (uint32, error) {
	marker, err := bytes.ReadByte()
	if err != nil {
		return 0, fmt.Errorf("Unexpected error")
	}

	if marker < 0xD0 {
		return uint32((marker - 1)), nil
	}

	marker = (marker + 1) & 0b1111

	var ret []byte = []byte{0, 0, 0, 0}
	for i := 3; i >= 0; i-- {
		if (marker & (1 << i)) == 0 {
			ret[i] = 0
		} else {
			ret[i], err = bytes.ReadByte()
			if err != nil {
				return 0, fmt.Errorf("Unexpected error")
			}
		}
	}

	return uint32(binary.LittleEndian.Uint32(ret)), nil
}

func resolvePayload(p payload, lang ffxiv.Language) (string, error) {
	reader := bytes.NewReader(p.Bytes)
	_, err := reader.ReadByte() // start byte
	if err != nil {
		return "", fmt.Errorf("Payload is empty")
	}

	payloadTypeByte, err := reader.ReadByte()
	if err != nil {
		return "", fmt.Errorf("Invalid format")
	}

	validPayload := isValidPayloadType(payloadTypeByte)
	if validPayload == false {
		return "", fmt.Errorf("Incorrect payload type")
	}

	_, err = reader.ReadByte()
	if err != nil {
		return "", fmt.Errorf("Invalid format")
	}

	rowGroupByte, err := reader.ReadByte()
	if err != nil {
		return "", fmt.Errorf("Invalid format")
	}
	rowGroup := uint32(rowGroupByte)

	rowId, err := getInteger(reader)
	if err != nil {
		return "", err
	}

	ret, err := ffxiv.AUTO_TRANSLATE[ffxiv.MapStruct{RowGroup: rowGroup, RowId: rowId}].Language(lang)
	if err != nil {
		return "", err
	}
	return ret, nil
}
