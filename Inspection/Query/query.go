package query

import "fmt"
import query "github.com/kumpfdp/HealthInspection/Query"
import rest "github.com/kumpfdp/HealthInspection/Inspection/Rest"

type Query struct {
	businessName string
	address      string
	city         string
	state        string
}

type QueryResult struct {
	businessName         string
	address              string
	city                 string
	state                string
	inspectionType       string
	violationDescription string
}

var (
	_ query.Query       = &Query{}
	_ query.QueryResult = &QueryResult{}
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

func (q *Query) BusinessName() string {
	return q.businessName
}

func (q *Query) Address() string {
	return q.address
}

func (q *Query) City() string {
	return q.city
}

func (q *Query) State() string {
	return q.state
}

func (q *Query) Print() {
	fmt.Println("Query:")
	if q.businessName != "" {
		fmt.Printf("BusinessName: %v\n", q.businessName)
	}
	if q.address != "" {
		fmt.Printf("Address: %v\n", q.address)
	}
	if q.city != "" {
		fmt.Printf("City: %v\n", q.city)
	}
	if q.state != "" {
		fmt.Printf("State: %v\n", q.state)
	}
}

func (q *Query) Execute() (*query.Results, error) {
	repo, err := rest.NewRestRepository(
		rest.WithName("testing"),
		rest.WithQuery(q),
		rest.WithLimit(1),
	)

	if err != nil {
		return nil, err
	}

	/*
	 * TODO
	 * Get all the results from the rest query
	 * and convert those into query.Results
	 * so that they can be returned to the executor.
	 */
	var inspections rest.RestInspections
	if inspections, err = repo.GetAll(); err != nil {
		return nil, err
	}

	results := make(query.Results, 0)
	for _, inspection := range inspections {
		results = append(results, &QueryResult{
			businessName:         inspection.BusinessName,
			address:              inspection.Address,
			city:                 inspection.City,
			state:                inspection.State,
			inspectionType:       inspection.InspType,
			violationDescription: inspection.ViolationDescription})
	}
	return &results, nil
}

func (qr *QueryResult) Print() {
	fmt.Println("Query Result:")
	if qr.businessName != "" {
		fmt.Printf("businessName: %v\n", qr.businessName)
	}
	if qr.address != "" {
		fmt.Printf("address: %v\n", qr.address)
	}
	if qr.city != "" {
		fmt.Printf("city: %v\n", qr.city)
	}
	if qr.state != "" {
		fmt.Printf("state: %v\n", qr.state)
	}
	if qr.inspectionType != "" {
		fmt.Printf("Inspection Type: %v\n", qr.inspectionType)
	}
	if qr.violationDescription != "" {
		fmt.Printf("Violation Description: %v\n", qr.violationDescription)
	}
}
