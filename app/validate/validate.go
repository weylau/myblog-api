package validate

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Validate struct {
	validator *validator.Validate
	trans ut.Translator
	errs validator.ValidationErrors
}

func Default() (*Validate, error){
	validate := &Validate{}
	validate.validator = validator.New()
	lang := zh.New()
	uni := ut.New(lang, lang)
	trans, _ := uni.GetTranslator("zh")
	validate.SetTrans(trans)
	err := validate.registerDefaultTranslations(validate.trans)
	return validate,err
}

func (this *Validate) SetTrans(trans ut.Translator) {
	this.trans = trans
}

func (this *Validate) registerDefaultTranslations(trans ut.Translator) error {
	return en_translations.RegisterDefaultTranslations(this.validator, trans)
}


func (this *Validate) CheckStruct(s interface{}) bool {
	err := this.validator.Struct(s)
	if err != nil {
		this.errs = err.(validator.ValidationErrors)
		return false
	}
	return true
}

func (this *Validate) GetAllError() []string{
	var errList  []string
	for _, e := range this.errs {
		// can translate each error one at a time.
		errList = append(errList, e.Translate(this.trans))
	}
	return errList
}

func (this *Validate) GetOneError() string{
	for _, e := range this.errs {
		// can translate each error one at a time.
		return e.Translate(this.trans)
	}
	return ""
}

