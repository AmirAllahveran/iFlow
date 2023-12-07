package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/smf8/http-monitor/common"
	"github.com/smf8/http-monitor/model"
	"net/http"
)

type userAuthRequest struct {
	Username string `valid:"stringlength(4|32), alphanum" json:"username"`
	Password string `valid:"stringlength(4|32)" json:"password"`
}

// binding user auth request with model.User instance
func (r *userAuthRequest) bind(c echo.Context, user *model.User) error {
	if err := c.Bind(r); err != nil {
		return common.NewRequestError("error binding user request", err, http.StatusBadRequest)
	}
	if _, err := govalidator.ValidateStruct(r); err != nil {
		e := common.NewValidationError(err, "Error validating sign-up request")
		return e
	}
	user.Username = r.Username
	user.Password = r.Password
	return nil
}

type ExperimentRequest struct {
	Packages   PackageData    `json:"Packages"`
	Experiment ExperimentData `json:"Experiment"`
}

type PackageData struct {
	Name    string       `json:"Name"`
	Modules []ModuleData `json:"Modules"`
}

type ModuleData struct {
	Name  string `json:"Name"`
	Image string `json:"Image"`
}

type ExperimentData struct {
	Title         string             `json:"Title"`
	Description   string             `json:"Description"`
	ModuleConfigs []ModuleConfigData `json:"ModuleConfigs"`
}

type ModuleConfigData struct {
	//Cmd    string            `json:"Cmd"`
	Config map[string]string `json:"Config"`
	Count  int               `json:"Count"`
}
