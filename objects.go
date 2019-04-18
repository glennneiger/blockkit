package blockkit

type TextObject struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji"`
	Verbatim bool   `json:"verbatim"`
}

type ConfirmObject struct {
	Title   TextObject `json:"title"`
	Text    TextObject `json:"text"`
	Confirm TextObject `json:"confirm"`
	Deny    TextObject `json:"deny"`
}

type OptionObject struct {
	Text  TextObject `json:"text"`
	Value string     `json:"value"`
}

type OptionGroupObject struct {
	Label   TextObject     `json:"label"`
	Options []OptionObject `json:"options"`
}
