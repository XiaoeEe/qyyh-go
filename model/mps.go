package model

import "qyyh-go/database/table"

type OcrParm struct {
	Img string `json:"img"`
}

type CheckNameParm struct {
	Name []string `json:"name"`
}

type GetMpsParm struct {
	Date string `json:"date"`
}

type AddMPSParm struct {
	Table []table.MPS `json:"table"`
	Date  string      `json:"date"`
}

type DelMPSParm struct {
	Ids []int64 `json:"ids"`
}

type GetMPSCountParm struct {
	Date string `json:"date"`
}
