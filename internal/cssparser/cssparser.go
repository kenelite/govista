package cssparser

import (
	"strings"
)

// RuleSet represents a collection of CSS rules mapped to selectors.
type RuleSet map[string]map[string]string

// ParseCSS parses raw CSS content and returns a structured RuleSet.
func ParseCSS(css string) RuleSet {
	rules := make(RuleSet)
	css = strings.ReplaceAll(css, "\n", "")
	css = strings.ReplaceAll(css, "\t", "")
	blocks := strings.Split(css, "}")

	for _, block := range blocks {
		parts := strings.Split(block, "{")
		if len(parts) != 2 {
			continue
		}
		selector := strings.TrimSpace(parts[0])
		body := strings.TrimSpace(parts[1])
		props := strings.Split(body, ";")

		propMap := make(map[string]string)
		for _, prop := range props {
			kv := strings.SplitN(prop, ":", 2)
			if len(kv) != 2 {
				continue
			}
			key := strings.TrimSpace(kv[0])
			val := strings.TrimSpace(kv[1])
			propMap[key] = val
		}

		if len(propMap) > 0 {
			rules[selector] = propMap
		}
	}

	return rules
}

// GetStylesForSelector returns style properties for a given tag or class.
func (r RuleSet) GetStylesForSelector(selector string) map[string]string {
	if props, ok := r[selector]; ok {
		return props
	}
	// Handle class selectors like .header
	if strings.HasPrefix(selector, ".") {
		return r[selector]
	}
	// Handle element-based selectors like p, div
	return r[selector]
}
