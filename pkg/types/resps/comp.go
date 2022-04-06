package resps

type CompCondResp struct {
	CompName  string `json:"compName" form:"compName"`
	CompType  string `json:"compType" form:"compType"`
	CompImg   string `json:"compImg" form:"compImg"`
	PeopleNum string `json:"peopleNum" form:"peopleNum"`
}
