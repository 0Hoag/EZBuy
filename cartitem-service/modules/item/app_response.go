package modules

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter"`
}

func NewSuccessResponse(data, paging interface{}) *successRes {
	return &successRes{Data: data, Paging: paging}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return &successRes{Data: data, Paging: nil}
}
