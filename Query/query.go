package query

type Results []QueryResult

type Query interface {
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
