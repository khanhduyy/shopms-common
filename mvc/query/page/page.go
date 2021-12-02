package page

type Result struct {
	Total int64
	Items interface{}
}

func EmptyResult() *Result {
	return &Result{
		Total: 0,
	}
}

func NewResult(total int64, items interface{}) *Result {
	if total == 0 {
		return EmptyResult()
	}
	return &Result{
		Total: total,
		Items: items,
	}
}
