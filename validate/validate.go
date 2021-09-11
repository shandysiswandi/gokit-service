package validate

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/shandysiswandi/gokit-service/entity"
)

var (
	once     sync.Once
	validate *validator.Validate
)

func init() {
	once.Do(func() {
		validate = validator.New()
	})
}

// how to use error validation
//
// check if the error not from validator
// if _, ok := err.(*validator.InvalidValidationError); ok {
// 	fmt.Println(err)
// 	return
// }
//
// consume error from validator
// for _, err := range err.(validator.ValidationErrors) {
// 	fmt.Println(err.Namespace())
// 	fmt.Println(err.Field())
// 	fmt.Println(err.StructNamespace())
// 	fmt.Println(err.StructField())
// 	fmt.Println(err.Tag())
// 	fmt.Println(err.ActualTag())
// 	fmt.Println(err.Kind())
// 	fmt.Println(err.Type())
// 	fmt.Println(err.Value())
// 	fmt.Println(err.Param())
// }

// validator.ValidationErrors
func ValidateGetAllTodoTodo(req entity.GetAllTodoTodoRequest) error {
	return validate.Struct(req)
}

func ValidateGetTodoByIDTodo(req entity.GetTodoByIDTodoRequest) error {
	return validate.Struct(req)
}

func ValidateCreateTodo(req entity.CreateTodoRequest) error {
	return validate.Struct(req)
}

func ValidateUpdateTodo(req entity.UpdateTodoRequest) error {
	return validate.Struct(req)
}

func ValidateDeleteTodo(req entity.DeleteTodoRequest) error {
	return validate.Struct(req)
}
