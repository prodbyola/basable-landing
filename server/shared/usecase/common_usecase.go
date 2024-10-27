package common

// uri is used for binding request params
type EntityID struct {
	ID string `uri:"id" binding:"required,uuid"`
} // @name EntityID

type SuccessResponseDto struct {
	StatusText string      `json:"statusText"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
} // @name SuccessResponseDto

type FailedResponseDto struct {
	StatusText string
	StatusCode int
	ErrorType  string
	Error      string
} // @name FailedResponseDto

type PaginateQuery struct {
	Page int `form:"page"`
	Size int `form:"size"`
} // @name PaginateQuery

type PaginateResponseDto struct {
	Content interface{} `json:"content"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	Total   int         `json:"total"`
} // @name PaginateResponseDto
