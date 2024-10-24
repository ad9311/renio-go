package eval

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/vars"
)

type ModelEval struct {
	String        StringEval
	Int           IntEval
	Float32       FloatEval
	Time          TimeEval
	errorMessages errorMessages
}

type (
	EvalKey       int
	StringEval    map[string]map[EvalKey]int
	IntEval       map[string]map[EvalKey]int
	FloatEval     map[string]map[EvalKey]float32
	TimeEval      map[string]map[EvalKey]time.Time
	errorMessages []string
)

const (
	Min EvalKey = iota
	Max
	Length
	Fixed
)

func (mv *ModelEval) ValidateModel(value any) error {
	model := reflect.ValueOf(value)
	modelName := model.Type().Name()
	kind := reflect.TypeOf(value)

	for i := 0; i < model.NumField(); i++ {
		value := model.Field(i)
		kind := kind.Field(i)
		name := kind.Name

		switch kind.Type {
		case reflect.TypeOf(""):
			mv.validateString(name, value.String())
		case reflect.TypeOf(float32(0)):
			mv.validateFloat(name, float32(value.Float()))
		default:
			console.Fatal(fmt.Sprintf("wrong validation type for %s", modelName))
		}
	}

	if len(mv.errorMessages) > 0 {
		errMsgs := mv.errorMessages.join()
		mv.errorMessages.flush()
		return fmt.Errorf("%s", errMsgs)
	}

	return nil
}

// --- Helpers --- //

func fatalAtWrongTypeName(name string) {
	console.Fatal(fmt.Sprintf("wrong validation type for %s", name))
}

func (es *errorMessages) appendString(errorMsg string) {
	*es = append(*es, errorMsg)
}

func (es *errorMessages) join() string {
	return strings.Join(*es, ", ")
}

func (es *errorMessages) flush() {
	*es = errorMessages{}
}

func formatAndFilter(name string, value string) string {
	if vars.FilteredFields[name] {
		return "[FILTERED]"
	}

	return fmt.Sprintf("'%s'", value)
}

// --- String --- //

func (mv *ModelEval) validateString(name string, str string) {
	validations := mv.String[name]
	filtered := formatAndFilter(name, str)

	for key, val := range validations {
		switch key {
		case Length:
			if len([]rune(str)) != val {
				errMsg := fmt.Sprintf("%s: value of %s is not of length %d", name, filtered, val)
				mv.errorMessages.appendString(errMsg)
			}
		case Min:
			if len([]rune(str)) < val {
				errMsg := fmt.Sprintf("%s: value of %s is less than %d", name, filtered, val)
				mv.errorMessages.appendString(errMsg)
			}
		case Max:
			if len([]rune(str)) > val {
				errMsg := fmt.Sprintf("%s: value of %s is greater than %d", name, filtered, val)
				mv.errorMessages.appendString(errMsg)
			}
		default:
			fatalAtWrongTypeName(name)
		}
	}
}

// --- Float --- //

func (mv *ModelEval) validateFloat(name string, float float32) {
	validations := mv.Float32[name]
	filtered := formatAndFilter(name, fmt.Sprintf("%f", float))

	for key, val := range validations {
		switch key {
		case Fixed:
			if float != val {
				errMsg := fmt.Sprintf("%s: value of %s is different than %f", name, filtered, val)
				mv.errorMessages.appendString(errMsg)
			}
		case Min:
			if float < val {
				errMsg := fmt.Sprintf("%s: value of %s is less than %f", name, filtered, val)
				mv.errorMessages.appendString(errMsg)
			}
		case Max:
			if float > val {
				errMsg := fmt.Sprintf("%s: value of %s is greater than %f", name, filtered, val)
				mv.errorMessages.appendString(errMsg)
			}
		default:
			fatalAtWrongTypeName(name)
		}
	}
}
