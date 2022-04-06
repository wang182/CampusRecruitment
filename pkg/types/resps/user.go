package resps

import "CampusRecruitment/pkg/models"

type UserLoginResp struct {
	Token string `json:"token"`
}

type PassUpdateResp struct {
	Id    models.Id `json:"id"`
	Name  string    `json:"name"`
	Role  string    `json:"role"`
	Email string    `json:"email"`
}
