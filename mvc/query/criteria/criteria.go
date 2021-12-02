package criteria

import (
	"strings"
	"time"

	"gorm.io/gorm/clause"
)

// Expr alias gorm clause expression
type Expr = clause.Expression

// Join type
type Join struct {
	Table, As, Cond string
}

// Table alias gorm table
type Table = clause.Table

func Or(exprs ...Expr) Expr {
	if len(exprs) == 0 {
		return nil
	}
	return clause.Or(exprs...)
}

func And(exprs ...Expr) Expr {
	if len(exprs) == 0 {
		return nil
	}
	return clause.And(exprs...)
}

func EqUint(column string, value uint) Expr {
	if value > 0 {
		return clauseEq(column, value)
	}
	return nil
}

func EqInt64(column string, value int64) Expr {
	if value > 0 {
		return clauseEq(column, value)
	}
	return nil
}

// Wildcard clauses is the search wildcard expression value.
// Ex: LIKE '%keyword%'
func Wildcard(column, keyword string) Expr {
	if len(keyword) > 0 {
		return clause.Like{Column: column, Value: "%" + keyword + "%"}
	}
	return nil
}

// SqlTimeExp clauses is SQL express with the datetime
func SqlTimeExp(sql string, value *time.Time, layout string) clause.Expression {
	if value != nil && len(layout) > 0 {
		var format = value.Format(layout)
		return clause.Expr{SQL: sql, Vars: []interface{}{format}}
	}
	return nil
}

func InStr(column string, keywords ...string) Expr {
	if len(keywords) > 0 {
		var words []interface{}
		for _, w := range keywords {
			if strings.TrimSpace(w) != "" {
				words = append(words, w)
			}
		}
		return clause.IN{Column: column, Values: words}
	}
	return nil
}

func InUint(column string, keywords ...uint) Expr {
	if len(keywords) > 0 {
		return clause.IN{Column: column, Values: []interface{}{keywords}}
	}
	return nil
}

func clauseEq(column string, value interface{}) clause.Eq {
	return clause.Eq{Column: column, Value: value}
}

func OrderBy(column, order string) (raw string) {
	if len(column) > 0 {
		raw = column
		if len(order) > 0 {
			raw = raw + " " + order
		}
	}
	if len(raw) > 0 {
		order = raw
	}
	return
}
