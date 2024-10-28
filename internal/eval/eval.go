package eval

import (
	"fmt"
	"regexp"
	"strings"
)

type Issues []string

type ErrEval struct {
	Issues Issues
}

type ModelEval struct {
	Strings []String
	Floats  []Float
	Ints    []Int
	issues  Issues
}

type String struct {
	Name    string
	Value   string
	Fixed   int
	Min     int
	Max     int
	Pattern string
}

type Float struct {
	Name     string
	Value    float32
	Positive bool
	Negative bool
	Fixed    float32
	Min      float32
	Max      float32
}

type Int struct {
	Name     string
	Value    int
	Positive bool
	Negative bool
	Fixed    int
	Min      int
	Max      int
}

func (e *ErrEval) Error() string {
	return strings.Join(e.Issues, ", ")
}

func (m *ModelEval) Validate() error {
	for _, s := range m.Strings {
		issues := s.validate()
		m.issues = append(m.issues, issues...)
	}

	for _, s := range m.Floats {
		issues := s.validate()
		m.issues = append(m.issues, issues...)
	}

	for _, s := range m.Ints {
		issues := s.validate()
		m.issues = append(m.issues, issues...)
	}

	if len(m.issues) > 0 {
		err := ErrEval{
			Issues: m.issues,
		}
		return &err
	}

	return nil
}

func (s String) validate() Issues {
	var issues Issues

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

func (f Float) validate() Issues {
	var issues Issues

	if issue := numericFixedValidation(f.Name, f.Fixed, f.Value); issue != "" {
		return Issues{issue}
	}

	if issue := numericPositiveValidation(f.Name, f.Positive, f.Value); issue != "" {
		issues = append(issues, issue)
	}

	if issue := numericNegativeValidation(f.Name, f.Negative, f.Value); issue != "" {
		issues = append(issues, issue)
	}

	if issue := numericMinValidation(f.Name, f.Min, f.Value); issue != "" {
		issues = append(issues, issue)
	}

	if issue := numericMaxValidation(f.Name, f.Max, f.Value); issue != "" {
		issues = append(issues, issue)
	}

	return issues
}

func (i Int) validate() Issues {
	var issues Issues

	if issue := numericFixedValidation(i.Name, float32(i.Fixed), float32(i.Value)); issue != "" {
		return Issues{issue}
	}

	if issue := numericPositiveValidation(i.Name, i.Positive, float32(i.Value)); issue != "" {
		issues = append(issues, issue)
	}

	if issue := numericNegativeValidation(i.Name, i.Negative, float32(i.Value)); issue != "" {
		issues = append(issues, issue)
	}

	if issue := numericMinValidation(i.Name, float32(i.Min), float32(i.Value)); issue != "" {
		issues = append(issues, issue)
	}

	if issue := numericMaxValidation(i.Name, float32(i.Max), float32(i.Value)); issue != "" {
		issues = append(issues, issue)
	}

	return issues
}

// --- Helpers --- //

func numericFixedValidation(name string, fixed float32, value float32) string {
	var issue string

	if fixed > 0 && fixed != value {
		return fmt.Sprintf("%s: must have a fixed value of %f", name, fixed)
	}

	return issue
}

func numericPositiveValidation(name string, positive bool, value float32) string {
	var issue string

	if positive && value <= 0 {
		return fmt.Sprintf("%s: must be positive", name)
	}

	return issue
}

func numericNegativeValidation(name string, negative bool, value float32) string {
	var issue string

	if negative && value <= 0 {
		return fmt.Sprintf("%s: must be negative", name)
	}

	return issue
}

func numericMinValidation(name string, min float32, value float32) string {
	var issue string

	if min > 0 && min > value {
		return fmt.Sprintf("%s: must have a minimum value of %f", name, min)
	}

	return issue
}

func numericMaxValidation(name string, max float32, value float32) string {
	var issue string

	if max > 0 && max < value {
		return fmt.Sprintf("%s: must have a maximum value of %f", name, max)
	}

	return issue
}
