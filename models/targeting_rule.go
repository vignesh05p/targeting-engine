package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// JSONStringArray is a custom type that implements sql.Scanner and driver.Valuer
type JSONStringArray []string

// Scan implements the sql.Scanner interface
func (a *JSONStringArray) Scan(value interface{}) error {
	if value == nil {
		*a = JSONStringArray{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		// Try parsing as JSON array first
		if err := json.Unmarshal(v, a); err != nil {
			// If that fails, try splitting the string
			*a = strings.Split(string(v), ",")
			// Trim spaces from each element
			for i := range *a {
				(*a)[i] = strings.TrimSpace((*a)[i])
			}
		}
		return nil
	case string:
		if err := json.Unmarshal([]byte(v), a); err != nil {
			*a = strings.Split(v, ",")
			for i := range *a {
				(*a)[i] = strings.TrimSpace((*a)[i])
			}
		}
		return nil
	case []string:
		*a = v
		return nil
	default:
		return fmt.Errorf("failed to scan JSONStringArray: unexpected type %T", value)
	}
}

// Value implements the driver.Valuer interface
func (a JSONStringArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(a)
}

type TargetingRule struct {
	ID         string          `json:"id"`
	CampaignID string          `json:"campaign_id"`
	Dimension  string          `json:"dimension"` // "app", "country", "os"
	RuleType   string          `json:"rule_type"` // "INCLUDE" or "EXCLUDE"
	Values     JSONStringArray `json:"values"`    // e.g., ["US", "Canada"]
}
