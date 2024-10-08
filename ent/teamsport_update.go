// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marko-durasic/olympics/ent/predicate"
	"github.com/marko-durasic/olympics/ent/teamsport"
)

// TeamSportUpdate is the builder for updating TeamSport entities.
type TeamSportUpdate struct {
	config
	hooks    []Hook
	mutation *TeamSportMutation
}

// Where appends a list predicates to the TeamSportUpdate builder.
func (tsu *TeamSportUpdate) Where(ps ...predicate.TeamSport) *TeamSportUpdate {
	tsu.mutation.Where(ps...)
	return tsu
}

// SetCountry sets the "country" field.
func (tsu *TeamSportUpdate) SetCountry(s string) *TeamSportUpdate {
	tsu.mutation.SetCountry(s)
	return tsu
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (tsu *TeamSportUpdate) SetNillableCountry(s *string) *TeamSportUpdate {
	if s != nil {
		tsu.SetCountry(*s)
	}
	return tsu
}

// SetSport sets the "sport" field.
func (tsu *TeamSportUpdate) SetSport(s string) *TeamSportUpdate {
	tsu.mutation.SetSport(s)
	return tsu
}

// SetNillableSport sets the "sport" field if the given value is not nil.
func (tsu *TeamSportUpdate) SetNillableSport(s *string) *TeamSportUpdate {
	if s != nil {
		tsu.SetSport(*s)
	}
	return tsu
}

// SetGold sets the "gold" field.
func (tsu *TeamSportUpdate) SetGold(i int) *TeamSportUpdate {
	tsu.mutation.ResetGold()
	tsu.mutation.SetGold(i)
	return tsu
}

// SetNillableGold sets the "gold" field if the given value is not nil.
func (tsu *TeamSportUpdate) SetNillableGold(i *int) *TeamSportUpdate {
	if i != nil {
		tsu.SetGold(*i)
	}
	return tsu
}

// AddGold adds i to the "gold" field.
func (tsu *TeamSportUpdate) AddGold(i int) *TeamSportUpdate {
	tsu.mutation.AddGold(i)
	return tsu
}

// SetSilver sets the "silver" field.
func (tsu *TeamSportUpdate) SetSilver(i int) *TeamSportUpdate {
	tsu.mutation.ResetSilver()
	tsu.mutation.SetSilver(i)
	return tsu
}

// SetNillableSilver sets the "silver" field if the given value is not nil.
func (tsu *TeamSportUpdate) SetNillableSilver(i *int) *TeamSportUpdate {
	if i != nil {
		tsu.SetSilver(*i)
	}
	return tsu
}

// AddSilver adds i to the "silver" field.
func (tsu *TeamSportUpdate) AddSilver(i int) *TeamSportUpdate {
	tsu.mutation.AddSilver(i)
	return tsu
}

// SetBronze sets the "bronze" field.
func (tsu *TeamSportUpdate) SetBronze(i int) *TeamSportUpdate {
	tsu.mutation.ResetBronze()
	tsu.mutation.SetBronze(i)
	return tsu
}

// SetNillableBronze sets the "bronze" field if the given value is not nil.
func (tsu *TeamSportUpdate) SetNillableBronze(i *int) *TeamSportUpdate {
	if i != nil {
		tsu.SetBronze(*i)
	}
	return tsu
}

// AddBronze adds i to the "bronze" field.
func (tsu *TeamSportUpdate) AddBronze(i int) *TeamSportUpdate {
	tsu.mutation.AddBronze(i)
	return tsu
}

// SetPoints sets the "points" field.
func (tsu *TeamSportUpdate) SetPoints(i int) *TeamSportUpdate {
	tsu.mutation.ResetPoints()
	tsu.mutation.SetPoints(i)
	return tsu
}

// SetNillablePoints sets the "points" field if the given value is not nil.
func (tsu *TeamSportUpdate) SetNillablePoints(i *int) *TeamSportUpdate {
	if i != nil {
		tsu.SetPoints(*i)
	}
	return tsu
}

// AddPoints adds i to the "points" field.
func (tsu *TeamSportUpdate) AddPoints(i int) *TeamSportUpdate {
	tsu.mutation.AddPoints(i)
	return tsu
}

// SetTotalScore sets the "total_score" field.
func (tsu *TeamSportUpdate) SetTotalScore(i int) *TeamSportUpdate {
	tsu.mutation.ResetTotalScore()
	tsu.mutation.SetTotalScore(i)
	return tsu
}

// SetNillableTotalScore sets the "total_score" field if the given value is not nil.
func (tsu *TeamSportUpdate) SetNillableTotalScore(i *int) *TeamSportUpdate {
	if i != nil {
		tsu.SetTotalScore(*i)
	}
	return tsu
}

// AddTotalScore adds i to the "total_score" field.
func (tsu *TeamSportUpdate) AddTotalScore(i int) *TeamSportUpdate {
	tsu.mutation.AddTotalScore(i)
	return tsu
}

// SetPopulation sets the "population" field.
func (tsu *TeamSportUpdate) SetPopulation(i int) *TeamSportUpdate {
	tsu.mutation.ResetPopulation()
	tsu.mutation.SetPopulation(i)
	return tsu
}

// SetNillablePopulation sets the "population" field if the given value is not nil.
func (tsu *TeamSportUpdate) SetNillablePopulation(i *int) *TeamSportUpdate {
	if i != nil {
		tsu.SetPopulation(*i)
	}
	return tsu
}

// AddPopulation adds i to the "population" field.
func (tsu *TeamSportUpdate) AddPopulation(i int) *TeamSportUpdate {
	tsu.mutation.AddPopulation(i)
	return tsu
}

// SetPerCapita sets the "per_capita" field.
func (tsu *TeamSportUpdate) SetPerCapita(f float64) *TeamSportUpdate {
	tsu.mutation.ResetPerCapita()
	tsu.mutation.SetPerCapita(f)
	return tsu
}

// SetNillablePerCapita sets the "per_capita" field if the given value is not nil.
func (tsu *TeamSportUpdate) SetNillablePerCapita(f *float64) *TeamSportUpdate {
	if f != nil {
		tsu.SetPerCapita(*f)
	}
	return tsu
}

// AddPerCapita adds f to the "per_capita" field.
func (tsu *TeamSportUpdate) AddPerCapita(f float64) *TeamSportUpdate {
	tsu.mutation.AddPerCapita(f)
	return tsu
}

// Mutation returns the TeamSportMutation object of the builder.
func (tsu *TeamSportUpdate) Mutation() *TeamSportMutation {
	return tsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tsu *TeamSportUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tsu.sqlSave, tsu.mutation, tsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsu *TeamSportUpdate) SaveX(ctx context.Context) int {
	affected, err := tsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tsu *TeamSportUpdate) Exec(ctx context.Context) error {
	_, err := tsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsu *TeamSportUpdate) ExecX(ctx context.Context) {
	if err := tsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tsu *TeamSportUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(teamsport.Table, teamsport.Columns, sqlgraph.NewFieldSpec(teamsport.FieldID, field.TypeInt))
	if ps := tsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsu.mutation.Country(); ok {
		_spec.SetField(teamsport.FieldCountry, field.TypeString, value)
	}
	if value, ok := tsu.mutation.Sport(); ok {
		_spec.SetField(teamsport.FieldSport, field.TypeString, value)
	}
	if value, ok := tsu.mutation.Gold(); ok {
		_spec.SetField(teamsport.FieldGold, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.AddedGold(); ok {
		_spec.AddField(teamsport.FieldGold, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.Silver(); ok {
		_spec.SetField(teamsport.FieldSilver, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.AddedSilver(); ok {
		_spec.AddField(teamsport.FieldSilver, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.Bronze(); ok {
		_spec.SetField(teamsport.FieldBronze, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.AddedBronze(); ok {
		_spec.AddField(teamsport.FieldBronze, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.Points(); ok {
		_spec.SetField(teamsport.FieldPoints, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.AddedPoints(); ok {
		_spec.AddField(teamsport.FieldPoints, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.TotalScore(); ok {
		_spec.SetField(teamsport.FieldTotalScore, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.AddedTotalScore(); ok {
		_spec.AddField(teamsport.FieldTotalScore, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.Population(); ok {
		_spec.SetField(teamsport.FieldPopulation, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.AddedPopulation(); ok {
		_spec.AddField(teamsport.FieldPopulation, field.TypeInt, value)
	}
	if value, ok := tsu.mutation.PerCapita(); ok {
		_spec.SetField(teamsport.FieldPerCapita, field.TypeFloat64, value)
	}
	if value, ok := tsu.mutation.AddedPerCapita(); ok {
		_spec.AddField(teamsport.FieldPerCapita, field.TypeFloat64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teamsport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tsu.mutation.done = true
	return n, nil
}

// TeamSportUpdateOne is the builder for updating a single TeamSport entity.
type TeamSportUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeamSportMutation
}

// SetCountry sets the "country" field.
func (tsuo *TeamSportUpdateOne) SetCountry(s string) *TeamSportUpdateOne {
	tsuo.mutation.SetCountry(s)
	return tsuo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (tsuo *TeamSportUpdateOne) SetNillableCountry(s *string) *TeamSportUpdateOne {
	if s != nil {
		tsuo.SetCountry(*s)
	}
	return tsuo
}

// SetSport sets the "sport" field.
func (tsuo *TeamSportUpdateOne) SetSport(s string) *TeamSportUpdateOne {
	tsuo.mutation.SetSport(s)
	return tsuo
}

// SetNillableSport sets the "sport" field if the given value is not nil.
func (tsuo *TeamSportUpdateOne) SetNillableSport(s *string) *TeamSportUpdateOne {
	if s != nil {
		tsuo.SetSport(*s)
	}
	return tsuo
}

// SetGold sets the "gold" field.
func (tsuo *TeamSportUpdateOne) SetGold(i int) *TeamSportUpdateOne {
	tsuo.mutation.ResetGold()
	tsuo.mutation.SetGold(i)
	return tsuo
}

// SetNillableGold sets the "gold" field if the given value is not nil.
func (tsuo *TeamSportUpdateOne) SetNillableGold(i *int) *TeamSportUpdateOne {
	if i != nil {
		tsuo.SetGold(*i)
	}
	return tsuo
}

// AddGold adds i to the "gold" field.
func (tsuo *TeamSportUpdateOne) AddGold(i int) *TeamSportUpdateOne {
	tsuo.mutation.AddGold(i)
	return tsuo
}

// SetSilver sets the "silver" field.
func (tsuo *TeamSportUpdateOne) SetSilver(i int) *TeamSportUpdateOne {
	tsuo.mutation.ResetSilver()
	tsuo.mutation.SetSilver(i)
	return tsuo
}

// SetNillableSilver sets the "silver" field if the given value is not nil.
func (tsuo *TeamSportUpdateOne) SetNillableSilver(i *int) *TeamSportUpdateOne {
	if i != nil {
		tsuo.SetSilver(*i)
	}
	return tsuo
}

// AddSilver adds i to the "silver" field.
func (tsuo *TeamSportUpdateOne) AddSilver(i int) *TeamSportUpdateOne {
	tsuo.mutation.AddSilver(i)
	return tsuo
}

// SetBronze sets the "bronze" field.
func (tsuo *TeamSportUpdateOne) SetBronze(i int) *TeamSportUpdateOne {
	tsuo.mutation.ResetBronze()
	tsuo.mutation.SetBronze(i)
	return tsuo
}

// SetNillableBronze sets the "bronze" field if the given value is not nil.
func (tsuo *TeamSportUpdateOne) SetNillableBronze(i *int) *TeamSportUpdateOne {
	if i != nil {
		tsuo.SetBronze(*i)
	}
	return tsuo
}

// AddBronze adds i to the "bronze" field.
func (tsuo *TeamSportUpdateOne) AddBronze(i int) *TeamSportUpdateOne {
	tsuo.mutation.AddBronze(i)
	return tsuo
}

// SetPoints sets the "points" field.
func (tsuo *TeamSportUpdateOne) SetPoints(i int) *TeamSportUpdateOne {
	tsuo.mutation.ResetPoints()
	tsuo.mutation.SetPoints(i)
	return tsuo
}

// SetNillablePoints sets the "points" field if the given value is not nil.
func (tsuo *TeamSportUpdateOne) SetNillablePoints(i *int) *TeamSportUpdateOne {
	if i != nil {
		tsuo.SetPoints(*i)
	}
	return tsuo
}

// AddPoints adds i to the "points" field.
func (tsuo *TeamSportUpdateOne) AddPoints(i int) *TeamSportUpdateOne {
	tsuo.mutation.AddPoints(i)
	return tsuo
}

// SetTotalScore sets the "total_score" field.
func (tsuo *TeamSportUpdateOne) SetTotalScore(i int) *TeamSportUpdateOne {
	tsuo.mutation.ResetTotalScore()
	tsuo.mutation.SetTotalScore(i)
	return tsuo
}

// SetNillableTotalScore sets the "total_score" field if the given value is not nil.
func (tsuo *TeamSportUpdateOne) SetNillableTotalScore(i *int) *TeamSportUpdateOne {
	if i != nil {
		tsuo.SetTotalScore(*i)
	}
	return tsuo
}

// AddTotalScore adds i to the "total_score" field.
func (tsuo *TeamSportUpdateOne) AddTotalScore(i int) *TeamSportUpdateOne {
	tsuo.mutation.AddTotalScore(i)
	return tsuo
}

// SetPopulation sets the "population" field.
func (tsuo *TeamSportUpdateOne) SetPopulation(i int) *TeamSportUpdateOne {
	tsuo.mutation.ResetPopulation()
	tsuo.mutation.SetPopulation(i)
	return tsuo
}

// SetNillablePopulation sets the "population" field if the given value is not nil.
func (tsuo *TeamSportUpdateOne) SetNillablePopulation(i *int) *TeamSportUpdateOne {
	if i != nil {
		tsuo.SetPopulation(*i)
	}
	return tsuo
}

// AddPopulation adds i to the "population" field.
func (tsuo *TeamSportUpdateOne) AddPopulation(i int) *TeamSportUpdateOne {
	tsuo.mutation.AddPopulation(i)
	return tsuo
}

// SetPerCapita sets the "per_capita" field.
func (tsuo *TeamSportUpdateOne) SetPerCapita(f float64) *TeamSportUpdateOne {
	tsuo.mutation.ResetPerCapita()
	tsuo.mutation.SetPerCapita(f)
	return tsuo
}

// SetNillablePerCapita sets the "per_capita" field if the given value is not nil.
func (tsuo *TeamSportUpdateOne) SetNillablePerCapita(f *float64) *TeamSportUpdateOne {
	if f != nil {
		tsuo.SetPerCapita(*f)
	}
	return tsuo
}

// AddPerCapita adds f to the "per_capita" field.
func (tsuo *TeamSportUpdateOne) AddPerCapita(f float64) *TeamSportUpdateOne {
	tsuo.mutation.AddPerCapita(f)
	return tsuo
}

// Mutation returns the TeamSportMutation object of the builder.
func (tsuo *TeamSportUpdateOne) Mutation() *TeamSportMutation {
	return tsuo.mutation
}

// Where appends a list predicates to the TeamSportUpdate builder.
func (tsuo *TeamSportUpdateOne) Where(ps ...predicate.TeamSport) *TeamSportUpdateOne {
	tsuo.mutation.Where(ps...)
	return tsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tsuo *TeamSportUpdateOne) Select(field string, fields ...string) *TeamSportUpdateOne {
	tsuo.fields = append([]string{field}, fields...)
	return tsuo
}

// Save executes the query and returns the updated TeamSport entity.
func (tsuo *TeamSportUpdateOne) Save(ctx context.Context) (*TeamSport, error) {
	return withHooks(ctx, tsuo.sqlSave, tsuo.mutation, tsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsuo *TeamSportUpdateOne) SaveX(ctx context.Context) *TeamSport {
	node, err := tsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tsuo *TeamSportUpdateOne) Exec(ctx context.Context) error {
	_, err := tsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsuo *TeamSportUpdateOne) ExecX(ctx context.Context) {
	if err := tsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tsuo *TeamSportUpdateOne) sqlSave(ctx context.Context) (_node *TeamSport, err error) {
	_spec := sqlgraph.NewUpdateSpec(teamsport.Table, teamsport.Columns, sqlgraph.NewFieldSpec(teamsport.FieldID, field.TypeInt))
	id, ok := tsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TeamSport.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teamsport.FieldID)
		for _, f := range fields {
			if !teamsport.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != teamsport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsuo.mutation.Country(); ok {
		_spec.SetField(teamsport.FieldCountry, field.TypeString, value)
	}
	if value, ok := tsuo.mutation.Sport(); ok {
		_spec.SetField(teamsport.FieldSport, field.TypeString, value)
	}
	if value, ok := tsuo.mutation.Gold(); ok {
		_spec.SetField(teamsport.FieldGold, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.AddedGold(); ok {
		_spec.AddField(teamsport.FieldGold, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.Silver(); ok {
		_spec.SetField(teamsport.FieldSilver, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.AddedSilver(); ok {
		_spec.AddField(teamsport.FieldSilver, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.Bronze(); ok {
		_spec.SetField(teamsport.FieldBronze, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.AddedBronze(); ok {
		_spec.AddField(teamsport.FieldBronze, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.Points(); ok {
		_spec.SetField(teamsport.FieldPoints, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.AddedPoints(); ok {
		_spec.AddField(teamsport.FieldPoints, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.TotalScore(); ok {
		_spec.SetField(teamsport.FieldTotalScore, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.AddedTotalScore(); ok {
		_spec.AddField(teamsport.FieldTotalScore, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.Population(); ok {
		_spec.SetField(teamsport.FieldPopulation, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.AddedPopulation(); ok {
		_spec.AddField(teamsport.FieldPopulation, field.TypeInt, value)
	}
	if value, ok := tsuo.mutation.PerCapita(); ok {
		_spec.SetField(teamsport.FieldPerCapita, field.TypeFloat64, value)
	}
	if value, ok := tsuo.mutation.AddedPerCapita(); ok {
		_spec.AddField(teamsport.FieldPerCapita, field.TypeFloat64, value)
	}
	_node = &TeamSport{config: tsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teamsport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tsuo.mutation.done = true
	return _node, nil
}
