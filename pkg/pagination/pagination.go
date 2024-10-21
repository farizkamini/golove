package pag

import (
	"fmt"
	"net/http"
	"strconv"
)

type Pagination struct {
	TotalData   int32 `db:"total_data"   json:"total_data"`
	TotalPages  int32 `db:"total_pages"  json:"total_pages"`
	CurrentPage int32 `db:"current_page" json:"current_page"`
	NextPage    int32 `db:"next_page"    json:"next_page"`
	PrevPage    int32 `db:"prev_page"    json:"prev_page"`
	PageNumber  int32 `db:"page_number"  json:"page_number"`
}

func LimitOffset(
	r *http.Request,
	keyLimitParam, pageLimitParam string,
) (limit int, offset int, err error) {
	limitStr := r.URL.Query().Get(keyLimitParam)
	if len(limitStr) <= 0 {
		limitStr = "10"
	}
	offsetStr := r.URL.Query().Get(pageLimitParam)
	if len(offsetStr) <= 0 {
		offsetStr = "0"
	}
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return 0, 0, fmt.Errorf("error parse limit: %v", err)
	}
	offset, err = strconv.Atoi(offsetStr)
	if err != nil {
		return 0, 0, fmt.Errorf("error parse offset: %v", err)
	}
	finalOffset := (offset - 1) * (limit)
	return limit, finalOffset, nil
}
