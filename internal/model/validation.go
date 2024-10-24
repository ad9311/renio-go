package model

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/ad9311/renio-go/internal/console"
	"github.com/ad9311/renio-go/internal/vars"
)

type ModelValidation struct {
	String        StringValidation
	Int           IntValidation
	Float32       Float32Validation
	Time          TimeValidation
	errorMessages ErrorMessages
}

type (
	ValidationKey     int
	StringValidation  map[string]map[ValidationKey]int
	IntValidation     map[string]map[ValidationKey]int
	Float32Validation map[string]map[ValidationKey]float32
	TimeValidation    map[string]map[ValidationKey]time.Time
	ErrorMessages     []string
)

const (
	Min ValidationKey = iota
	Max
	Length
	Fixed
)

func (mv *ModelValidation) ValidateModel(value any) {
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
			mv.validateFloat32(name, float32(value.Float()))
		default:
			console.Fatal(fmt.Sprintf("wrong validation type for %s", modelName))
		}
	}
}

// --- Helpers --- //

func fatalAtWrongTypeName(name string) {
	console.Fatal(fmt.Sprintf("wrong validation type for %s", name))
}

func (es *ErrorMessages) appendString(errorMsg string) {
	*es = append(*es, errorMsg)
}

func (es *ErrorMessages) join() string {
	return strings.Join(*es, ", ")
}

func formatAndFilter(name string, value string) string {
	if vars.FilteredFields[name] {
		return "[FILTERED]"
	}

	return fmt.Sprintf("'%s'", value)
}

// --- String --- //

func (mv *ModelValidation) validateString(name string, str string) {
	validations := mv.String[name]
	filtered := formatAndFilter(name, str)

	fmt.Println(filtered)

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

// --- Float32 --- //

func (mv *ModelValidation) validateFloat32(name string, float float32) {
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
