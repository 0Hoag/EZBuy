package modules

type Paging struct {
	Page      int   `json:"page" form:"page"`
	Size      int   `json:"size" form:"size"`
	TotalPage int64 `json:"totalPage" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Size <= 0 || p.Page >= 100 {
		p.Size = 10
	}
}
