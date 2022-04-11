package resps

import "CampusRecruitment/pkg/models"

type CompCondResp struct {
	Id        models.Id `json:"id"`
	CompName  string    `json:"compName" form:"compName"`
	CompType  string    `json:"compType" form:"compType"`
	Logo      string    `json:"logo" form:"logo"`
	PeopleNum string    `json:"peopleNum" form:"peopleNum"`
}

type CompDetailResp struct {
	CompName     string       `json:"compName,omitempty"`     //公司名
	Logo         string       `json:"logo,omitempty"`         //公司logo
	CompType     string       `json:"compType,omitempty"`     //公司类型
	PeopleNum    string       `json:"peopleNum,omitempty"`    //公司人数区间
	JobNum       int          `json:"jobNum,omitempty"`       //在招岗位数量
	UserNum      int          `json:"userNum,omitempty"`      //招聘者数量
	City         string       `json:"city,omitempty"`         //所在城市
	Introduction string       `json:"introduction,omitempty"` //公司简介
	Address      string       `json:"address,omitempty"`      //公司地址
	Url          string       `json:"url,omitempty"`          //公司官网
	HotJobs      []HotJobResp `json:"hotJobs"`                //热门职位
}
