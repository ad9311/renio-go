package eval

import (
	"fmt"
	"regexp"
)

type ModelEval struct {
	Strings []String
	Floats  []Float
	issues  Issues
}

type Issues []string

type String struct {
	Name    string
	Value   string
	Fixed   int
	Min     int
	Max     int
	Pattern string
}

type Float struct {
	Name  string
	Value float32
	Fixed float32
	Min   float32
	Max   float32
}

func (m *ModelEval) Validate() Issues {
	for _, s := range m.Strings {
		issues := s.validate()
		m.issues = append(m.issues, issues...)
	}

	for _, s := range m.Floats {
		issues := s.validate()
		m.issues = append(m.issues, issues...)
	}

	if len(m.issues) > 0 {
		return m.issues
	}

	return nil
}

func (s String) validate() Issues {
	var issues []string

	if s.Pattern != "" {
		re := regexp.MustCompile(s.Pattern)
		if !re.MatchString(s.Value) {
			issue := fmt.Sprintf("%s: is not a valid value", s.Name)
			issues = append(issues, issue)
		}
	}

	size := len([]rune(s.Value))
	if s.Fixed > 0 && s.Fixed != size {
		issue := fmt.Sprintf("%s: must have a fixed length of %d characters", s.Name, s.Fixed)
		return append(issues, issue)
	}

	if s.Min > 0 && s.Min > size {
		issue := fmt.Sprintf("%s: must have a minimum length of %d characters", s.Name, s.Min)
		issues = append(issues, issue)
	}

	if s.Max > 0 && s.Max < size {
		issue := fmt.Sprintf("%s: must have a maximum length of %d characters", s.Name, s.Max)
		issues = append(issues, issue)
	}

	return issues
}

func (s Float) validate() Issues {
	var issues []string

	if s.Fixed > 0 && s.Fixed != s.Value {
		issue := fmt.Sprintf("%s: must have a fixed value of %f", s.Name, s.Fixed)
		return append(issues, issue)
	}

	if s.Min > 0 && s.Min > s.Value {
		issue := fmt.Sprintf("%s: must have a minimum value of %f", s.Name, s.Min)
		issues = append(issues, issue)
	}

	if s.Max > 0 && s.Max < s.Value {
		issue := fmt.Sprintf("%s: must have a maximum value of %f", s.Name, s.Max)
		issues = append(issues, issue)
	}

	return issues
}
