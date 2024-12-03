package sestringgo

type PayloadType uint8

const (
	PayloadAutotranslate PayloadType = 0x2E
)

func isValidPayloadType(payloadType byte) bool {
	switch PayloadType(payloadType) {
	default:
		return false
	case PayloadAutotranslate:
		return true
	}
}
