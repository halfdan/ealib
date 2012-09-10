package optprobs

import (
	"errors"
	"fmt"
	"reflect"
)

type FunctionEvaluator interface {
	Evaluate([]float64) float64
	Name() string
}

var typeList []reflect.Type

// Init initializes the package and registers available
// optimization problems.
func init() {
	typeList = make([]reflect.Type,0, 10)
	typeList = append(typeList, reflect.TypeOf(AckleyFunction{}))
	typeList = append(typeList, reflect.TypeOf(GriewankFunction{}))
	typeList = append(typeList, reflect.TypeOf(RosenbrockFunction{}))
}

// Prints all registered Evaluators.
func PrintAvailable() {
	for i, t := range typeList {
		fmt.Printf("%d\t%s\n", i, string(t.Name()))
	}
}

// GetEvaluator returns an evaluator based on an index
func GetEvaluator(i int) (FunctionEvaluator, error) {
	if i < 0 || i >= len(typeList) {
		return nil, errors.New("Invalid index")
	}
	v := reflect.New(typeList[i])
	return v.Interface().(FunctionEvaluator), nil
}

// RegisterEvaluator registers a new FunctionEvaluator
func RegisterEvaluator(f FunctionEvaluator) {
	t := reflect.TypeOf(f)
	typeList = append(typeList, t)
}
