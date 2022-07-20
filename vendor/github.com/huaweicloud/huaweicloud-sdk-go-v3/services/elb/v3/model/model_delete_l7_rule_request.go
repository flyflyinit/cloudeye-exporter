package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteL7RuleRequest struct {
	// 策略ID。

	L7policyId string `json:"l7policy_id"`
	// 规则ID。

	L7ruleId string `json:"l7rule_id"`
}

func (o DeleteL7RuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteL7RuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteL7RuleRequest", string(data)}, " ")
}