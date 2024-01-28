package request

import "system-point/domain/entity"

type RequestAddKindActivityDTO struct {
	Name  string `json:"name"`
	Point int    `json:"point"`
}

func (input RequestAddKindActivityDTO) ToAddKindActivity() entity.KindActivity {
	return entity.KindActivity{
		Name:  input.Name,
		Point: input.Point,
	}
}
