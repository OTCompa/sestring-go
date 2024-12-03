package ffxiv

import "fmt"

// LocalisedText represents text in multiple languages
type LocalisedText struct {
	En string
	Ja string
	De string
	Fr string
}

func (t LocalisedText) Language(lang Language) (string, error) {
	switch lang {
	case LangEn:
		return t.En, nil
	case LangJa:
		return t.Ja, nil
	case LangDe:
		return t.De, nil
	case LangFr:
		return t.Fr, nil
	default:
		return "", fmt.Errorf("Invalid language selected")
	}
}
