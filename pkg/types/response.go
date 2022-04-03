package types

type APIResponse struct {
	Code          int         `json:"code"`
	Message       string      `json:"message"`
	MessageDetail string      `json:"messageDetail,omitempty"`
	Result        interface{} `json:"result"`
}

type APIPageResult struct {
	List     interface{} `json:"list"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
}
