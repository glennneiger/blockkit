package blockkit

type ImageElement struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

type ButtonElement struct {
	Type     string        `json:"type"`
	Text     TextObject    `json:"text"`
	ActionID string        `json:"action_id"`
	URL      string        `json:"url"`
	Value    string        `json:"value"`
	Style    string        `json:"style"`
	Confirm  ConfirmObject `json:"confirm"`
}

type SelectMenuStaticOptionsElement struct {
	Type          string              `json:"type"`
	Placeholder   TextObject          `json:"placeholder"`
	ActionID      string              `json:"action_id"`
	Options       []OptionObject      `json:"options"`
	OptionGroups  []OptionGroupObject `json:"option_groups"`
	InitialOption OptionObject        `json:"initial_option"`
	Confirm       ConfirmObject       `json:"confirm"`
}

type SelectMenuExternalSourceElement struct {
	Type           string        `json:"type"`
	Placeholder    TextObject    `json:"placeholder"`
	ActionID       string        `json:"action_id"`
	InitialOption  OptionObject  `json:"initial_option"`
	MinQueryLength int           `json:"min_query_length"`
	Confirm        ConfirmObject `json:"confirm"`
}

type SelectMenuUserListElement struct {
	Type        string        `json:"type"`
	Placeholder TextObject    `json:"placeholder"`
	ActionID    string        `json:"action_id"`
	InitialUser string        `json:"initial_user"`
	Confirm     ConfirmObject `json:"confirm"`
}

type SelectMenuConversationsListElement struct {
	Type                string        `json:"type"`
	Placeholder         TextObject    `json:"placeholder"`
	ActionID            string        `json:"action_id"`
	InitialConversation string        `json:"initial_conversation"`
	Confirm             ConfirmObject `json:"confirm"`
}

type SelectMenuChannelsListElement struct {
	Type           string        `json:"type"`
	Placeholder    TextObject    `json:"placeholder"`
	ActionID       string        `json:"action_id"`
	InitialChannel string        `json:"initial_channel"`
	Confirm        ConfirmObject `json:"confirm"`
}

type OverflowMenuElement struct {
	Type     string         `json:"type"`
	ActionID string         `json:"action_id"`
	Options  []OptionObject `json:"options"`
	Confirm  ConfirmObject  `json:"confirm"`
}

type DatePickerElement struct {
	Type        string        `json:"type"`
	ActionID    string        `json:"action_id"`
	Placeholder TextObject    `json:"placeholder"`
	InitialDate string        `json:"initial_data"`
	Confirm     ConfirmObject `json:"confirm"`
}
