package function

import (
	"fmt"
)

func Handle(req []byte) string {
	request, err := NewRequest(req)
	if err != nil {
		return fmt.Sprintf("Request parse error: %v", err)
	}

	start, end, step, err := request.GetQueryRange()
	if err != nil {
		return fmt.Sprintf("Query range parse error: %v", err)
	}

	c, err := NewClient(request.Server, "", "")
	if err != nil {
		return fmt.Sprintf("HTTP client error: %v", err)
	}

	response, err := c.QueryRange(request.Query, start, end, step)
	if err != nil {
		return fmt.Sprintf("Query error: %v", err)
	}

	output, err := formatRespose(response, request.Format)
	if err != nil {
		return fmt.Sprintf("Query result parse error: %v", err)
	}

	return output
}
