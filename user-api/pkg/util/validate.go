package util

import (
	"errors"
	"reflect"

	"github.com/1348453525/user-redeem-code-grpc/user-api/pkg/result"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zht "github.com/go-playground/validator/v10/translations/zh"
)

func Validate(c *gin.Context, dto interface{}) bool {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	// 注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	// 注册翻译器
	err := zht.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		// 记录错误日志并返回国际化错误信息
		// fmt.Println(err.Error())
		result.Error(c, 500,
			"注册翻译器失败")
		return false
	}

	// 验证参数
	err = validate.Struct(dto)
	if err != nil {
		var resError string
		resErrors := make(map[string]string)
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, fieldErr := range validationErrors {
				// fmt.Printf("字段 %s %s 违反了 %s 规则\n", fieldErr.StructField(), fieldErr.Field(), fieldErr.Tag())
				// fmt.Printf("实际值: %v\n", fieldErr.Value())
				// fmt.Println(fieldErr.Translate(trans))
				// 返回的错误数据
				resError = fieldErr.Translate(trans)
				// resErrors[fieldErr.StructField()] = fieldErr.Translate(trans)
				jsonName := GetJSONName(dto, fieldErr.StructField())
				resErrors[jsonName] = fieldErr.Translate(trans)
			}
		}
		result.Error(c, 400, resError, resErrors)
		return false
	}
	return true
}
