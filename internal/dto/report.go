package dto

type ReportRead struct {
	URL string `json:"url"`
}

type ReportFilter struct {
	Year  uint `json:"year" validate:"required,number,gt=0"`
	Month uint `json:"month" validate:"required,number,ge=1,le=12"`
}
