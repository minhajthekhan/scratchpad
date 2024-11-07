package legos

import (
	"fmt"
	"slices"
	"strings"
)

type Specification[T any] interface {
	IsSatisfiedBy(T) bool
	AsSQL() (query string, args []any)
}

type LegoSpecification struct {
	colors      []string
	dimensions  []LegoDimension
	notMoreThan int
}

func NewLegoSpecification(
	colors []string,
	dimensions []LegoDimension,
	notMoreThan int,
) Specification[Lego] {
	return &LegoSpecification{
		colors:      colors,
		dimensions:  dimensions,
		notMoreThan: notMoreThan,
	}
}

func (s *LegoSpecification) IsSatisfiedBy(lego Lego) bool {
	if !slices.Contains(s.colors, lego.Color) {
		return false
	}

	for _, dimension := range s.dimensions {
		if lego.Dimensions.Equals(dimension) {
			return true
		}
	}
	return false
}

func (s *LegoSpecification) AsSQL() (query string, args []any) {
	args = []any{}
	query = "SELECT shelf_position FROM legos"

	query, args = s.buildColorQuery(query, args)
	query, args = s.buildDimensionsQuery(query, args)
	query, args = s.buildNotMoreThanQuery(query, args)
	return query, args
}

func (s *LegoSpecification) buildDimensionsQuery(query string, args []any) (string, []any) {
	dimensionClauses := []string{}
	for _, dimension := range s.dimensions {

		sizeArg := len(args) + 1
		heightArg := len(args) + 2

		dimensionClauses = append(dimensionClauses, fmt.Sprintf("(size = $%d AND height < $%d)", sizeArg, heightArg))
		args = append(args, dimension.Size, dimension.Height)
	}

	// Add the OR conditions if there are any
	if len(dimensionClauses) > 0 {
		query += " AND (" + strings.Join(dimensionClauses, " OR ") + ")"
	}
	return query, args
}

func (s *LegoSpecification) buildColorQuery(query string, args []any) (string, []any) {
	query += `WHERE color = ANY($1)`
	args = append(args, s.colors)
	return query, args
}

func (s *LegoSpecification) buildNotMoreThanQuery(query string, args []any) (string, []any) {
	args = append(args, s.notMoreThan)
	query += fmt.Sprintf(" LIMIT $%d", len(args))
	return query, args
}
