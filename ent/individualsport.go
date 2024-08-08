// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/marko-durasic/olympics/ent/individualsport"
)

// IndividualSport is the model entity for the IndividualSport schema.
type IndividualSport struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Sport holds the value of the "sport" field.
	Sport string `json:"sport,omitempty"`
	// Gold holds the value of the "gold" field.
	Gold int `json:"gold,omitempty"`
	// Silver holds the value of the "silver" field.
	Silver int `json:"silver,omitempty"`
	// Bronze holds the value of the "bronze" field.
	Bronze int `json:"bronze,omitempty"`
	// Points holds the value of the "points" field.
	Points int `json:"points,omitempty"`
	// TotalScore holds the value of the "total_score" field.
	TotalScore int `json:"total_score,omitempty"`
	// Population holds the value of the "population" field.
	Population int `json:"population,omitempty"`
	// PerCapita holds the value of the "per_capita" field.
	PerCapita    float64 `json:"per_capita,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IndividualSport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case individualsport.FieldPerCapita:
			values[i] = new(sql.NullFloat64)
		case individualsport.FieldID, individualsport.FieldGold, individualsport.FieldSilver, individualsport.FieldBronze, individualsport.FieldPoints, individualsport.FieldTotalScore, individualsport.FieldPopulation:
			values[i] = new(sql.NullInt64)
		case individualsport.FieldCountry, individualsport.FieldSport:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IndividualSport fields.
func (is *IndividualSport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case individualsport.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			is.ID = int(value.Int64)
		case individualsport.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				is.Country = value.String
			}
		case individualsport.FieldSport:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sport", values[i])
			} else if value.Valid {
				is.Sport = value.String
			}
		case individualsport.FieldGold:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gold", values[i])
			} else if value.Valid {
				is.Gold = int(value.Int64)
			}
		case individualsport.FieldSilver:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field silver", values[i])
			} else if value.Valid {
				is.Silver = int(value.Int64)
			}
		case individualsport.FieldBronze:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field bronze", values[i])
			} else if value.Valid {
				is.Bronze = int(value.Int64)
			}
		case individualsport.FieldPoints:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field points", values[i])
			} else if value.Valid {
				is.Points = int(value.Int64)
			}
		case individualsport.FieldTotalScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_score", values[i])
			} else if value.Valid {
				is.TotalScore = int(value.Int64)
			}
		case individualsport.FieldPopulation:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field population", values[i])
			} else if value.Valid {
				is.Population = int(value.Int64)
			}
		case individualsport.FieldPerCapita:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field per_capita", values[i])
			} else if value.Valid {
				is.PerCapita = value.Float64
			}
		default:
			is.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IndividualSport.
// This includes values selected through modifiers, order, etc.
func (is *IndividualSport) Value(name string) (ent.Value, error) {
	return is.selectValues.Get(name)
}

// Update returns a builder for updating this IndividualSport.
// Note that you need to call IndividualSport.Unwrap() before calling this method if this IndividualSport
// was returned from a transaction, and the transaction was committed or rolled back.
func (is *IndividualSport) Update() *IndividualSportUpdateOne {
	return NewIndividualSportClient(is.config).UpdateOne(is)
}

// Unwrap unwraps the IndividualSport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (is *IndividualSport) Unwrap() *IndividualSport {
	_tx, ok := is.config.driver.(*txDriver)
	if !ok {
		panic("ent: IndividualSport is not a transactional entity")
	}
	is.config.driver = _tx.drv
	return is
}

// String implements the fmt.Stringer.
func (is *IndividualSport) String() string {
	var builder strings.Builder
	builder.WriteString("IndividualSport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", is.ID))
	builder.WriteString("country=")
	builder.WriteString(is.Country)
	builder.WriteString(", ")
	builder.WriteString("sport=")
	builder.WriteString(is.Sport)
	builder.WriteString(", ")
	builder.WriteString("gold=")
	builder.WriteString(fmt.Sprintf("%v", is.Gold))
	builder.WriteString(", ")
	builder.WriteString("silver=")
	builder.WriteString(fmt.Sprintf("%v", is.Silver))
	builder.WriteString(", ")
	builder.WriteString("bronze=")
	builder.WriteString(fmt.Sprintf("%v", is.Bronze))
	builder.WriteString(", ")
	builder.WriteString("points=")
	builder.WriteString(fmt.Sprintf("%v", is.Points))
	builder.WriteString(", ")
	builder.WriteString("total_score=")
	builder.WriteString(fmt.Sprintf("%v", is.TotalScore))
	builder.WriteString(", ")
	builder.WriteString("population=")
	builder.WriteString(fmt.Sprintf("%v", is.Population))
	builder.WriteString(", ")
	builder.WriteString("per_capita=")
	builder.WriteString(fmt.Sprintf("%v", is.PerCapita))
	builder.WriteByte(')')
	return builder.String()
}

// IndividualSports is a parsable slice of IndividualSport.
type IndividualSports []*IndividualSport
