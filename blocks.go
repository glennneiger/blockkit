package blockkit

type SectionBlock struct {
	Type      string       `json:"type"`
	Text      TextObject   `json:"text"`
	BlockID   string       `json:"block_id"`
	Fields    []TextObject `json:"fields"`
	Accessory ok           `json:"accessory"`
}

type DividerBlock struct {
	Type    string `json:"type"`
	BlockID string `json:"block_id"`
}

type ImageBlock struct {
	Type     string     `json:"type"`
	ImageURL string     `json:"image_url"`
	AltText  string     `json:"alt_text"`
	Title    TextObject `json:"title"`
	BlockID  string     `json:"block_id"`
}

type ActionsBlock struct {
	Type     string `json:"type"`
	Elements []ok   `json:"elements"`
	BlockID  string `json:"block_id"`
}

type ContextBlock struct {
	Type     string `json:"type"`
	Elements []ok   `json:"elements"`
	BlockID  string `json:"block_id"`
}
