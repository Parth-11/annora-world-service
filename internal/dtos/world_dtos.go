package dtos

type CreateWorldRequestDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description,omitempty" validate:"max=1000"`
	CoverImg    string `json:"cover_img,omitempty" validate:"omitempty,url"`
}

type UpdateWorldRequestDTO struct {
	Description *string `json:"description,omitempty" validate:"max=1000"`
	CoverImg    *string `json:"cover_img,omitempty" validate:"omitempty,url"`
}
