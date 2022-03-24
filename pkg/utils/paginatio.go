package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/now"
)

const (
	IdType         string = "ID"
	NumberType     string = "NUMBER"
	StringType     string = "STRING"
	BoolType       string = "BOOL"
	DateType       string = "DATE"
	DatetimeType   string = "DATETIME"
	TimeStringType string = "TIMESTRING"
	JSONType       string = "JSON"
)

// PaginationConfig is interface for all paginated query or any custom query
type PaginationConfig interface {
	Limit() int
	Offset() int
	Order() string
	QueryMap() map[string][]string
	Scopes() []Scope
	MetaScopes() []Scope
	AddScope(scope Scope)
}

// Pagination struct implement PaginationConfig
type Pagination struct {
	limit      int
	offset     int
	order      string
	queryMap   map[string][]string
	scopes     []Scope
	metaScopes []Scope
}

// AddScope will add new scope to existing scope
func (p *Pagination) AddScope(scope Scope) {
	p.scopes = append(p.scopes, scope)
}

// Limit will return current limit of pagination
func (p *Pagination) Limit() (res int) {
	return p.limit
}

// Limit will return current order of pagination
func (p *Pagination) Order() string {
	return p.order
}

// Limit will return current offset of pagination
func (p *Pagination) Offset() (res int) {
	return p.offset
}

// Limit will return current query map
// will be nil if pagination is not initiated using NewRequestPaginationConfig
func (p *Pagination) QueryMap() (res map[string][]string) {
	return p.queryMap
}

// Scopes will return all scope in current pagination
func (p *Pagination) Scopes() []Scope {
	return p.scopes
}

// Scopes will return all scope in current pagination
func (p *Pagination) MetaScopes() []Scope {
	return p.metaScopes
}

// NewPaginationConfig will create new Pagination with limit, offset, order and any scopes
// will set query map to nil
func NewPaginationConfig(limit int, offset int, order string, scopes ...Scope) PaginationConfig {
	paginationConfig := Pagination{
		limit:      limit,
		offset:     offset,
		order:      order,
		queryMap:   nil,
		scopes:     make([]Scope, 0),
		metaScopes: make([]Scope, 0),
	}

	if len(scopes) > 0 {
		paginationConfig.scopes = append(paginationConfig.scopes, scopes...)
	}

	return injectMetaScope(paginationConfig)
}

// NewRequestPaginationConfig will create new Pagination with request condition and filterable list
// all resulted scope come from filterable field with conditions field data
// if any conditions field that is not declared in filterable field will be omitted
func NewRequestPaginationConfig(conditions map[string][]string, filterable map[string]string) PaginationConfig {
	paginationConfig := Pagination{
		limit:      buildLimit(conditions),
		offset:     buildOffset(conditions),
		order:      buildOrder(conditions),
		queryMap:   conditions,
		scopes:     buildScope(conditions, filterable),
		metaScopes: make([]Scope, 0),
	}

	return injectMetaScope(paginationConfig)
}

// NewDefaultPaginationConfig will create a default Pagination with zero scope and 20 limit
func NewDefaultPaginationConfig() PaginationConfig {
	return NewPaginationConfig(20, 0, "")
}

// BuildLimit build the limit with 100 threshold given the conditions
func buildLimit(conditions map[string][]string) int {
	res := 20
	if len(conditions["limit"]) > 0 {
		res, _ = strconv.Atoi(conditions["limit"][0])
		if res > 300 {
			res = 300
		}
	}
	return res
}

// BuildOffset build the offset given the conditions
func buildOffset(conditions map[string][]string) int {
	res := 0
	if len(conditions["offset"]) > 0 {
		res, _ = strconv.Atoi(conditions["offset"][0])
	}
	return res
}

// BuildOrder build the order given the conditions
func buildOrder(conditions map[string][]string) string {
	var orders string
	if len(conditions["sort"]) > 0 {
		orders = strings.Join(conditions["sort"], ",")
		return orders
	}
	return "id desc"
}

// BuildOrder build all scope given the conditions
func buildScope(conditions map[string][]string, filterable map[string]string) []Scope {
	scopes := make([]Scope, 0)

	for name, value := range filterable {
		if len(conditions[name]) > 0 {
			switch value {
			case IdType:
				scopes = append(scopes, WhereInScope(name, conditions[name]))
			case StringType:
				scopes = append(scopes, WhereLikeScope(name, conditions[name][0]))
			case BoolType:
				boolean := false
				if conditions[name][0] == "true" {
					boolean = true
				}
				scopes = append(scopes, WhereIsScope(name, boolean))
			case NumberType:
				minmax := strings.Split(conditions[name][0], ",")
				scopes = append(scopes, WhereBetweenScope(name, minmax[0], minmax[1]))
			case DateType:
				minmax := strings.Split(conditions[name][0], ",")
				min, _ := now.Parse(minmax[0])
				max, _ := now.Parse(minmax[1])
				scopes = append(scopes, WhereBetweenScope(
					name, now.New(min).BeginningOfDay(), now.New(max).EndOfDay(),
				))
			case DatetimeType:
				minmax := strings.Split(conditions[name][0], ",")
				min, _ := now.Parse(minmax[0])
				if !CheckContainsTime(minmax[0]) {
					min = now.New(min).BeginningOfDay()
				}

				max, _ := now.Parse(minmax[1])
				if !CheckContainsTime(minmax[1]) {
					max = now.New(max).EndOfDay()
				}

				scopes = append(scopes, WhereBetweenScope(
					name, min.UTC(), max.UTC(),
				))
			case JSONType:
				param := strings.Split(name, ".")
				queryJSON := fmt.Sprintf("JSON_EXTRACT(%s , \"$.%s\")", param[0], strings.Join(param[1:], "."))

				// attempt conversion to uint64
				var uintConditionValues []uint64
				for _, value := range conditions[name] {
					intVal, err := strconv.ParseUint(value, 10, 64)
					if err == nil {
						uintConditionValues = append(uintConditionValues, intVal)
					}
				}

				if len(uintConditionValues) == len(conditions[name]) {
					scopes = append(scopes, WhereInScope(queryJSON, uintConditionValues))
				} else {
					scopes = append(scopes, WhereInScope(queryJSON, conditions[name]))
				}
			}
		}
	}

	return scopes
}

// OverrideKey will override condition key with desired key
func OverrideKey(conditions map[string][]string, original string, replaceBy string) {
	if targetValue, ok := conditions[original]; ok {
		delete(conditions, original)
		conditions[replaceBy] = targetValue
	}
}

func injectMetaScope(paginationConfig Pagination) PaginationConfig {
	if paginationConfig.limit > 0 {
		paginationConfig.metaScopes = append(paginationConfig.metaScopes, LimitScope(paginationConfig.limit))
	}

	if paginationConfig.offset > 0 {
		paginationConfig.metaScopes = append(paginationConfig.metaScopes, OffsetScope(paginationConfig.offset))
	}

	if paginationConfig.order != "" {
		paginationConfig.metaScopes = append(paginationConfig.metaScopes, OrderScope(paginationConfig.order))
	}

	return &paginationConfig
}
