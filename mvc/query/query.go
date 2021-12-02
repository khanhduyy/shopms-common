package query

import (
	"github.com/khanhduyy/shopms-common/client/db"
	c "github.com/khanhduyy/shopms-common/mvc/query/criteria"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	DefaultSize   = 10
	DefaultOffset = 0
)

type Query struct {
	distinct interface{}
	fields   interface{}
	preloads []string
	from     string
	join     []c.Join
	leftJoin []c.Join
	where    []c.Expr
	order    interface{}
	offset   int
	limit    int

	client *db.Client
}

func NewQuery(client *db.Client) *Query {
	return &Query{
		client: client,
	}
}

func (q *Query) FindPage(value interface{}) (int64, error) {
	done := make(chan bool, 1)

	var count int64
	// Count the total records matching criteria
	go q.getTotalRecords(value, &count, done)
	// Find records matching criteria and pagination result
	cli := q.build(value)
	cli.Order(q.order)
	cli.Limit(q.limit)
	cli.Offset(q.offset)
	q.withReload(cli)
	result := cli.Find(value)

	<-done

	if err := result.Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (q *Query) FindAll(value interface{}) error {
	client := q.build(value)
	q.withReload(client)
	result := client.Find(value)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (q *Query) build(model interface{}) *gorm.DB {
	client := q.client.Model(model)
	client.Select(q.fields)
	if q.distinct != nil {
		client.Distinct()
	}
	client.Table(q.from)
	client.Clauses(
		clause.From{
			Joins: q.withJoin(),
		},
	)
	if len(q.where) > 0 {
		client.Clauses(clause.Where{
			Exprs: q.where,
		})
	}
	return client
}

func (q *Query) withReload(client *gorm.DB) {
	if len(q.preloads) > 0 {
		for _, m := range q.preloads {
			client.Preload(m)
		}
	}
}

func (q *Query) withJoin() []clause.Join {
	inner := buildJoin(q.join, clause.InnerJoin)
	left := buildJoin(q.leftJoin, clause.LeftJoin)
	return append(inner, left...)
}

func (q *Query) getTotalRecords(value interface{}, count *int64, done chan bool) {
	cli := q.build(value)
	if q.distinct != nil {
		cli.Distinct(q.distinct)
	}
	cli.Count(count)
	done <- true
}
