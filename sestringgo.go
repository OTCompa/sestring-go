package sestringgo

import (
	"fmt"

	"github.com/OTCompa/sestring-go/ffxiv"
)

const START_BYTE = 0x02
const END_BYTE = 0x03
const AUTOTRANSLATE_BYTE = 0x2E

type payload struct {
	StartPos int
	EndPos   int
	Bytes    []byte
}

func Parse(str []byte, lang ffxiv.Language) (string, error) {
	var payloads []payload
	cursor := 0

	// check if any sestring payloads exist
	for cursor < len(str) {
		if str[cursor] == START_BYTE {
			// seek chunkLength byte
			chunkLenPos := cursor + 2

			if len(str) <= chunkLenPos {
				return "", fmt.Errorf("Invalid format")
			}

			// bounds checking
			chunkLen := int(str[chunkLenPos])
			endPos := chunkLenPos + chunkLen + 1
			if len(str) < endPos {
				return "", fmt.Errorf("Payload too short")
			}

			// add new payload to the list
			payloads = append(payloads, payload{
				StartPos: cursor,
				EndPos:   endPos,
				Bytes:    str[cursor:endPos],
			})

			cursor = endPos
		} else {
			cursor++
		}
	}

	ret := ""
	currPos := 0
	// resolve payloads and form resulting string
	for _, payload := range payloads {
		ret += string(str[currPos:payload.StartPos])
		resolvedPayload, err := resolvePayload(payload, lang)
		if err != nil {
			return "", fmt.Errorf("Error resolving payload")
		}
		ret += resolvedPayload
		currPos = payload.EndPos
	}
	ret += string(str[currPos:])

	return ret, nil
}
