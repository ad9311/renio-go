package eval

import (
	"fmt"
	"regexp"

	"github.com/ad9311/renio-go/internal/vars"
)

type ModelEval struct {
	Strings []String
	Floats  []Float
	errMsgs vars.ErrorMessages
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
	Name  string
	Value float32
	Fixed float32
	Min   float32
	Max   float32
}

func (m *ModelEval) Validate() vars.ErrorMessages {
	for _, s := range m.Strings {
		errMsgs := s.validate()
		m.errMsgs = append(m.errMsgs, errMsgs...)
	}

	for _, s := range m.Floats {
		errMsgs := s.validate()
		m.errMsgs = append(m.errMsgs, errMsgs...)
	}

	if len(m.errMsgs) > 0 {
		return m.errMsgs
	}

	return nil
}

func (s String) validate() vars.ErrorMessages {
	var errMsgs []string

	if s.Pattern != "" {
		re := regexp.MustCompile(s.Pattern)
		if !re.MatchString(s.Value) {
			errMsg := fmt.Sprintf("%s: is not a valid value", s.Name)
			errMsgs = append(errMsgs, errMsg)
		}
	}

	size := len([]rune(s.Value))
	if s.Fixed > 0 && s.Fixed != size {
		errMsg := fmt.Sprintf("%s: must have a fixed length of %d characters", s.Name, s.Fixed)
		return append(errMsgs, errMsg)
	}

	if s.Min > 0 && s.Min > size {
		errMsg := fmt.Sprintf("%s: must have a minimum length of %d characters", s.Name, s.Min)
		errMsgs = append(errMsgs, errMsg)
	}

	if s.Max > 0 && s.Max < size {
		errMsg := fmt.Sprintf("%s: must have a maximum length of %d characters", s.Name, s.Max)
		errMsgs = append(errMsgs, errMsg)
	}

	return errMsgs
}

func (s Float) validate() vars.ErrorMessages {
	var errMsgs []string

	if s.Fixed > 0 && s.Fixed != s.Value {
		errMsg := fmt.Sprintf("%s: must have a fixed value of %f", s.Name, s.Fixed)
		return append(errMsgs, errMsg)
	}

	if s.Min > 0 && s.Min > s.Value {
		errMsg := fmt.Sprintf("%s: must have a minimum value of %f", s.Name, s.Min)
		errMsgs = append(errMsgs, errMsg)
	}

	if s.Max > 0 && s.Max < s.Value {
		errMsg := fmt.Sprintf("%s: must have a maximum value of %f", s.Name, s.Max)
		errMsgs = append(errMsgs, errMsg)
	}

	return errMsgs
}
