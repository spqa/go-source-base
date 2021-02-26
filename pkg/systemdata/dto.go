package systemdata

type DataUpdateReq struct {
	Value string `json:"value"`
}

type DataRes map[string]string
