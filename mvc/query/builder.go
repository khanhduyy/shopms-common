package query

import (
	c "github.com/khanhduyy/shopms-common/mvc/query/criteria"
	"gorm.io/gorm/clause"
)

// Select provides the set of column to be used to retrieve row selected from one or more tables.
func (q *Query) Select(columns interface{}) *Query {
	q.fields = columns
	return q
}

// Distinct provides the set of columns to be used to remove of duplicated from a result set.
func (q *Query) Distinct(distinct interface{}) *Query {
	q.distinct = distinct
	return q
}

// From provides the source table name to be selected and possibly other clauses.
func (q *Query) From(name, alias string) *Query {
	if len(name) > 0 {
		q.from = name
		if len(alias) > 0 {
			q.from = q.from + " AS " + alias
		}
	}
	return q
}

// Join provides the inner join expression.
func (q *Query) Join(join c.Join) *Query {
	q.join = append(q.join, join)
	return q
}

// LeftJoin provides the left join expression.
func (q *Query) LeftJoin(left c.Join) *Query {
	q.leftJoin = append(q.leftJoin, left)
	return q
}

// Where provides array of expressions with one more conditions
// that evaluate to true for each row to be selected.
func (q *Query) Where(exps ...c.Expr) *Query {
	var where []c.Expr
	for _, expr := range exps {
		if expr != nil {
			where = append(where, expr)
		}
	}
	q.where = where
	return q
}

// Order provides the ordering rows in a result set.
func (q *Query) Order(raw string) *Query {
	if len(raw) > 0 {
		q.order = raw
	}
	return q
}

// Offset provides the page size to be returned a result set.
// If absents or negative the given input, DefaultSize to be used.
func (q *Query) Offset(offset int) {
	q.offset = offset
	if offset < 0 {
		q.offset = DefaultOffset
	}
}

// Limit provides the page size to be returned a result set.
// If absents or negative the given input, DefaultSize to be used.
func (q *Query) Limit(limit int) {
	q.limit = limit
	if limit <= 0 {
		q.limit = DefaultSize
	}
}

// Page provides the basic to be returned a pagination result set on given the offset and limit.
func (q *Query) Page(offset, limit int64) *Query {
	q.Limit(int(limit))
	q.Offset(int(offset) * q.limit)
	return q
}

// Preloads provides entity models to be allowed eager reloading.
// See https://gorm.io/docs/preload.html for more usage.
func (q *Query) Preloads(models ...string) *Query {
	q.preloads = models
	return q
}

func buildJoin(joins []c.Join, joinType clause.JoinType) []clause.Join {
	var joinClauses []clause.Join
	for _, j := range joins {
		joinClauses = append(joinClauses, clause.Join{
			Type:       joinType,
			Table:      clause.Table{Name: j.Table, Alias: j.As},
			ON:         clause.Where{Exprs: []clause.Expression{clause.Expr{SQL: j.Cond}}},
			Using:      nil,
			Expression: nil,
		})
	}
	return joinClauses
}
