package request

import "system-point/domain/entity"

type RequestRegisterDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (input RequestRegisterDTO) ToRegister() entity.Register {
	return entity.Register{
		Email:    input.Email,
		Password: input.Password,
	}
}
