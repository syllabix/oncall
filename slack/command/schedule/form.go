package schedule

func NewForm() Form {
	return newForm
}

type Form struct {
	ID         int     `json:"id,omitempty"`
	CallbackID string  `json:"callback_id,omitempty"`
	Blocks     []Block `json:"blocks"`
}

type Block struct {
	Type    string          `json:"type"`
	BlockID string          `json:"block_id"`
	Element *FormElement    `json:"element,omitempty"`
	Label   *Label          `json:"label,omitempty"`
	Actions []ActionElement `json:"elements,omitempty"`
}

type FormElement struct {
	Type        string   `json:"type,omitempty"`
	ActionID    string   `json:"action_id,omitempty"`
	Placeholder *Label   `json:"placeholder,omitempty"`
	Options     []Option `json:"options,omitempty"`
}

type Option struct {
	Text  Label  `json:"text"`
	Value string `json:"value"`
}

type Label struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji *bool  `json:"emoji,omitempty"`
}

type ActionElement struct {
	Type     string `json:"type"`
	Text     Label  `json:"text"`
	Value    string `json:"value"`
	ActionID string `json:"action_id"`
}
