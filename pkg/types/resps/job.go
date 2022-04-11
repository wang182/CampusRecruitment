package resps

import "CampusRecruitment/pkg/models"

type CreateJobResp struct {
	JobName   string    `json:"jobName" form:"jobName"`                          //职位名称
	PublishId models.Id `json:"publishId" form:"publishId" swaggerignore:"true"` //发布用户id
	CompId    models.Id `json:"compId" form:"compId" swaggerignore:"true"`       //公司id
	JobNum    int       `json:"jobNum" form:"jobNum"`                            //岗位数量
	City      string    `json:"city" form:"city"`                                //工作城市
	Address   string    `json:"address" form:"address"`                          //工作地址
	Tags      []string  `json:"tags" form:"tags"`                                //标签
}

type GetJobDetailResp struct {
	JobRespDetail
	JobDetailsUserResp
	CompCondResp
}

type JobRespDetail struct {
	JobName string   `json:"jobName,omitempty"`
	MinWage int      `json:"minWage,omitempty"`
	MaxWage int      `json:"maxWage,omitempty"`
	JobNum  string   `json:"jobNum,omitempty"`
	Desc    string   `json:"desc,omitempty"`
	City    string   `json:"city,omitempty"`
	Address string   `json:"address,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

type HotJobResp struct {
	JobName string `json:"jobName,omitempty"`
	MinWage int    `json:"minWage,omitempty"`
	MaxWage int    `json:"maxWage,omitempty"`
	JobNum  string `json:"jobNum,omitempty"`
	City    string `json:"city,omitempty"`
}
