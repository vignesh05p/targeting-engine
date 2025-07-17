package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"targeting-engine/db"
	"targeting-engine/models"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// DeliveryResponse represents the campaign delivery response
type DeliveryResponse struct {
	CID string `json:"cid"`
	Img string `json:"img"`
	CTA string `json:"cta"`
}

// writeJSONError writes a JSON error response
func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

// validateQueryParams validates required query parameters
func validateQueryParams(params map[string]string) error {
	for key, value := range params {
		if value == "" {
			return fmt.Errorf("missing %s param", key)
		}
	}
	return nil
}

// DeliveryHandler handles GET /v1/delivery requests
func DeliveryHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	params := map[string]string{
		"app":     r.URL.Query().Get("app"),
		"country": r.URL.Query().Get("country"),
		"os":      r.URL.Query().Get("os"),
	}

	// Validate parameters
	if err := validateQueryParams(params); err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Fetch active campaigns
	campaigns, err := fetchActiveCampaigns()
	if err != nil {
		log.Printf("Error fetching campaigns: %v", err)
		writeJSONError(w, "failed to fetch campaigns", http.StatusInternalServerError)
		return
	}

	// Match campaigns against rules
	matched := matchCampaigns(campaigns, params)

	// Return response
	if len(matched) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matched)
}

// fetchActiveCampaigns retrieves all active campaigns from the database
func fetchActiveCampaigns() ([]models.Campaign, error) {
	campaigns := []models.Campaign{}

	rows, err := db.DB.Query("SELECT id, name, image, cta, status FROM campaigns WHERE status = 'ACTIVE'")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Campaign
		if err := rows.Scan(&c.ID, &c.Name, &c.Image, &c.CTA, &c.Status); err != nil {
			return nil, fmt.Errorf("scan failed: %v", err)
		}
		campaigns = append(campaigns, c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %v", err)
	}

	return campaigns, nil
}

// fetchRules retrieves all targeting rules for a given campaign
func fetchRules(campaignID string) ([]models.TargetingRule, error) {
	rules := []models.TargetingRule{}

	rows, err := db.DB.Query("SELECT id, campaign_id, dimension, rule_type, values FROM targeting_rules WHERE campaign_id = $1", campaignID)
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var r models.TargetingRule
		if err := rows.Scan(&r.ID, &r.CampaignID, &r.Dimension, &r.RuleType, &r.Values); err != nil {
			return nil, fmt.Errorf("scan failed: %v", err)
		}
		rules = append(rules, r)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %v", err)
	}

	return rules, nil
}

// matchCampaigns evaluates campaigns against their targeting rules
func matchCampaigns(campaigns []models.Campaign, params map[string]string) []DeliveryResponse {
	var matched []DeliveryResponse

	for _, campaign := range campaigns {
		rules, err := fetchRules(campaign.ID)
		if err != nil {
			log.Printf("Error fetching rules for campaign %s: %v", campaign.ID, err)
			continue
		}

		if matchAll(rules, params) {
			matched = append(matched, DeliveryResponse{
				CID: campaign.ID,
				Img: campaign.Image,
				CTA: campaign.CTA,
			})
		}
	}

	return matched
}

// matchAll checks if all rules are satisfied for given parameters
func matchAll(rules []models.TargetingRule, params map[string]string) bool {
	for _, rule := range rules {
		dimension := strings.ToLower(rule.Dimension)
		inputVal := strings.ToLower(params[dimension])

		// Create matching map for rule values
		matching := make(map[string]bool)
		for _, v := range rule.Values {
			matching[strings.ToLower(v)] = true
		}

		log.Printf("Evaluating rule: Dimension=%s, RuleType=%s, InputValue=%s, MatchingValues=%v",
			rule.Dimension, rule.RuleType, inputVal, matching)

		switch rule.RuleType {
		case "INCLUDE":
			if !matching[inputVal] {
				log.Printf("INCLUDE rule failed: InputValue=%s not in MatchingValues", inputVal)
				return false
			}
		case "EXCLUDE":
			if matching[inputVal] {
				log.Printf("EXCLUDE rule failed: InputValue=%s is in MatchingValues", inputVal)
				return false
			}
		default:
			log.Printf("Unknown rule type: %s", rule.RuleType)
			return false
		}
	}

	return true
}
