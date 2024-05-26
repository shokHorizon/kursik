package entity

type (
	StringFilter struct {
		Eq  *string
		Neq *string
	}
	IntFilter struct {
		Eq  *int
		Neq *int
		Gt  *int
		Gte *int
		Lt  *int
		Lte *int
	}
)
