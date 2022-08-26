package response

type (
	StaffResponse struct {
		ID         string             `json:"_id"`
		Name       string             `json:"name" form:"name"`
		Email      string             `json:"email" form:"email"`
		Password   string             `json:"password,omitempty" form:"password"`
		Department DepartmentResponse `json:"department" form:"department"`
	}

	StaffFindAllResponse struct {
		Data  []StaffResponse `json:"data" form:"data"`
		Limit int64           `json:"limit" form:"limit"`
		Page  int64           `json:"page" form:"page"`
	}
)
