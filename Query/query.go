package query

type Results []QueryResult

type Query interface {
	/*
	 * TODO
	 * Add additional fields to the query, beyond these basics.
	 */
	BusinessName() string
	Address() string
	City() string
	State() string
	QueryWithBusinessName(name string) Query
	QueryWithAddress(address string) Query
	QueryWithCity(city string) Query
	QueryWithState(state string) Query
	Print()
	Execute() (*Results, error)
}

type QueryResult interface {
	Print()
}
