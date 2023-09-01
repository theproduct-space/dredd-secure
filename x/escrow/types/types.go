package types

// API CONDITIONS VALIDATION TYPES
type SubCondition struct {
	ConditionType string      `json:"conditionType"`
	DataType      string      `json:"dataType"`
	Name          string      `json:"name"`
	Path          []string    `json:"path"`
	Label         string      `json:"label"`
	Value         interface{} `json:"value"`
}

type TokenOfInterest struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type OracleCondition struct {
	Label           string          `json:"label"`
	Name            string          `json:"name"`
	Type            string          `json:"type"`
	SubConditions   []SubCondition  `json:"subConditions"`
	TokenOfInterest TokenOfInterest `json:"tokenOfInterest"`
}

type Header struct {
	Key   string
	Value string
}
