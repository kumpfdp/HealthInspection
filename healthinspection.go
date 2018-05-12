package healthinspection

import (
	"fmt"
)

type Query struct {
	businessName string
	address string
	city string
	state string
}

type QueryResult struct {
	inspectionType string
	violationDescription string
}

type Results []QueryResult

/*
 * Query
 */
func NewQuery() *Query {
	return &Query{}
}

func (q *Query) QueryWithBusinessName(name string) *Query {
	q.businessName = name
	return q
}

func (q *Query) QueryWithAddress(address string) *Query {
	q.address = address
	return q
}

func (q *Query) QueryWithCity(city string) *Query {
	q.city = city
	return q
}

func (q *Query) QueryWithState(state string) *Query {
	q.state = state
	return q
}

func (q *Query) Print() {
	fmt.Println("Query:")
	if (q.businessName != "") {
		fmt.Printf("BusinessName: %v\n", q.businessName)
	}
	if (q.address != "") {
		fmt.Printf("Address: %v\n", q.address)
	}
	if (q.city != "") {
		fmt.Printf("City: %v\n", q.city)
	}
	if (q.state != "") {
		fmt.Printf("State: %v\n", q.state)
	}
}

/*
 * Query Execution and Results.
 */
func (q *Query) Execute() *Results{
	return &Results{{inspectionType: "type", violationDescription: "description"}}
}

func (qr *QueryResult) Print() {
	fmt.Println("Query Result:");
	if (qr.inspectionType != "") {
		fmt.Printf("Inspection Type: %v\n", qr.inspectionType)
	}
	if (qr.violationDescription != "") {
		fmt.Printf("Violation Description: %v\n", qr.violationDescription)
	}
}

