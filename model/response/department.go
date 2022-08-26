package response

type DepartmentResponse struct {
	ID      string `json:"_id" form:"_id"`
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
}
