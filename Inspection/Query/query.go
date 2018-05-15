package query

import "fmt"
import query "github.com/CincyGolangMeetup/HealthInspection/Query"
import rest "github.com/CincyGolangMeetup/HealthInspection/Inspection/Rest"

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

var (
	_ query.Query = &Query{}
)

/*
 * Query
 */
func NewQuery() query.Query {
	return &Query{}
}

func (q *Query) QueryWithBusinessName(name string) query.Query {
	q.businessName = name
	return q
}

func (q *Query) QueryWithAddress(address string) query.Query {
	q.address = address
	return q
}

func (q *Query) QueryWithCity(city string) query.Query {
	q.city = city
	return q
}

func (q *Query) QueryWithState(state string) query.Query {
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

func (q *Query) Execute() (*query.Results, error) {
	aq, err := rest.NewRestRepository(
		rest.WithName("testing"),
		rest.WithQuery(q),
		rest.WithLimit(1),
	)

	if (err != nil) {
		fmt.Println("Error: %v\n", err)
		return nil, err
	}

	results, err := aq.GetAll()
	if (err != nil) {
		fmt.Println("Error: %v\n", err)
		return nil, err
	}

	fmt.Printf("%v\n", results)

	return &query.Results{&QueryResult{inspectionType: "type", violationDescription: "description"}}, nil;
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
