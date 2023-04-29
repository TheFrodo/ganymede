// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/live"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/ent/vod"
)

// ChannelUpdate is the builder for updating Channel entities.
type ChannelUpdate struct {
	config
	hooks    []Hook
	mutation *ChannelMutation
}

// Where appends a list predicates to the ChannelUpdate builder.
func (cu *ChannelUpdate) Where(ps ...predicate.Channel) *ChannelUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetExtID sets the "ext_id" field.
func (cu *ChannelUpdate) SetExtID(s string) *ChannelUpdate {
	cu.mutation.SetExtID(s)
	return cu
}

// SetNillableExtID sets the "ext_id" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableExtID(s *string) *ChannelUpdate {
	if s != nil {
		cu.SetExtID(*s)
	}
	return cu
}

// ClearExtID clears the value of the "ext_id" field.
func (cu *ChannelUpdate) ClearExtID() *ChannelUpdate {
	cu.mutation.ClearExtID()
	return cu
}

// SetName sets the "name" field.
func (cu *ChannelUpdate) SetName(s string) *ChannelUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDisplayName sets the "display_name" field.
func (cu *ChannelUpdate) SetDisplayName(s string) *ChannelUpdate {
	cu.mutation.SetDisplayName(s)
	return cu
}

// SetImagePath sets the "image_path" field.
func (cu *ChannelUpdate) SetImagePath(s string) *ChannelUpdate {
	cu.mutation.SetImagePath(s)
	return cu
}

// SetRetention sets the "retention" field.
func (cu *ChannelUpdate) SetRetention(b bool) *ChannelUpdate {
	cu.mutation.SetRetention(b)
	return cu
}

// SetNillableRetention sets the "retention" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableRetention(b *bool) *ChannelUpdate {
	if b != nil {
		cu.SetRetention(*b)
	}
	return cu
}

// ClearRetention clears the value of the "retention" field.
func (cu *ChannelUpdate) ClearRetention() *ChannelUpdate {
	cu.mutation.ClearRetention()
	return cu
}

// SetRetentionDays sets the "retention_days" field.
func (cu *ChannelUpdate) SetRetentionDays(s string) *ChannelUpdate {
	cu.mutation.SetRetentionDays(s)
	return cu
}

// SetNillableRetentionDays sets the "retention_days" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableRetentionDays(s *string) *ChannelUpdate {
	if s != nil {
		cu.SetRetentionDays(*s)
	}
	return cu
}

// ClearRetentionDays clears the value of the "retention_days" field.
func (cu *ChannelUpdate) ClearRetentionDays() *ChannelUpdate {
	cu.mutation.ClearRetentionDays()
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ChannelUpdate) SetUpdatedAt(t time.Time) *ChannelUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// AddVodIDs adds the "vods" edge to the Vod entity by IDs.
func (cu *ChannelUpdate) AddVodIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.AddVodIDs(ids...)
	return cu
}

// AddVods adds the "vods" edges to the Vod entity.
func (cu *ChannelUpdate) AddVods(v ...*Vod) *ChannelUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.AddVodIDs(ids...)
}

// AddLiveIDs adds the "live" edge to the Live entity by IDs.
func (cu *ChannelUpdate) AddLiveIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.AddLiveIDs(ids...)
	return cu
}

// AddLive adds the "live" edges to the Live entity.
func (cu *ChannelUpdate) AddLive(l ...*Live) *ChannelUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cu.AddLiveIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cu *ChannelUpdate) Mutation() *ChannelMutation {
	return cu.mutation
}

// ClearVods clears all "vods" edges to the Vod entity.
func (cu *ChannelUpdate) ClearVods() *ChannelUpdate {
	cu.mutation.ClearVods()
	return cu
}

// RemoveVodIDs removes the "vods" edge to Vod entities by IDs.
func (cu *ChannelUpdate) RemoveVodIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.RemoveVodIDs(ids...)
	return cu
}

// RemoveVods removes "vods" edges to Vod entities.
func (cu *ChannelUpdate) RemoveVods(v ...*Vod) *ChannelUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cu.RemoveVodIDs(ids...)
}

// ClearLive clears all "live" edges to the Live entity.
func (cu *ChannelUpdate) ClearLive() *ChannelUpdate {
	cu.mutation.ClearLive()
	return cu
}

// RemoveLiveIDs removes the "live" edge to Live entities by IDs.
func (cu *ChannelUpdate) RemoveLiveIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.RemoveLiveIDs(ids...)
	return cu
}

// RemoveLive removes "live" edges to Live entities.
func (cu *ChannelUpdate) RemoveLive(l ...*Live) *ChannelUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cu.RemoveLiveIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChannelUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks[int, ChannelMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChannelUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChannelUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChannelUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ChannelUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := channel.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *ChannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(channel.Table, channel.Columns, sqlgraph.NewFieldSpec(channel.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.ExtID(); ok {
		_spec.SetField(channel.FieldExtID, field.TypeString, value)
	}
	if cu.mutation.ExtIDCleared() {
		_spec.ClearField(channel.FieldExtID, field.TypeString)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(channel.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.DisplayName(); ok {
		_spec.SetField(channel.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := cu.mutation.ImagePath(); ok {
		_spec.SetField(channel.FieldImagePath, field.TypeString, value)
	}
	if value, ok := cu.mutation.Retention(); ok {
		_spec.SetField(channel.FieldRetention, field.TypeBool, value)
	}
	if cu.mutation.RetentionCleared() {
		_spec.ClearField(channel.FieldRetention, field.TypeBool)
	}
	if value, ok := cu.mutation.RetentionDays(); ok {
		_spec.SetField(channel.FieldRetentionDays, field.TypeString, value)
	}
	if cu.mutation.RetentionDaysCleared() {
		_spec.ClearField(channel.FieldRetentionDays, field.TypeString)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(channel.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedVodsIDs(); len(nodes) > 0 && !cu.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.VodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.LiveCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.LiveTable,
			Columns: []string{channel.LiveColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(live.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedLiveIDs(); len(nodes) > 0 && !cu.mutation.LiveCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.LiveTable,
			Columns: []string{channel.LiveColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(live.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.LiveIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.LiveTable,
			Columns: []string{channel.LiveColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(live.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChannelUpdateOne is the builder for updating a single Channel entity.
type ChannelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChannelMutation
}

// SetExtID sets the "ext_id" field.
func (cuo *ChannelUpdateOne) SetExtID(s string) *ChannelUpdateOne {
	cuo.mutation.SetExtID(s)
	return cuo
}

// SetNillableExtID sets the "ext_id" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableExtID(s *string) *ChannelUpdateOne {
	if s != nil {
		cuo.SetExtID(*s)
	}
	return cuo
}

// ClearExtID clears the value of the "ext_id" field.
func (cuo *ChannelUpdateOne) ClearExtID() *ChannelUpdateOne {
	cuo.mutation.ClearExtID()
	return cuo
}

// SetName sets the "name" field.
func (cuo *ChannelUpdateOne) SetName(s string) *ChannelUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDisplayName sets the "display_name" field.
func (cuo *ChannelUpdateOne) SetDisplayName(s string) *ChannelUpdateOne {
	cuo.mutation.SetDisplayName(s)
	return cuo
}

// SetImagePath sets the "image_path" field.
func (cuo *ChannelUpdateOne) SetImagePath(s string) *ChannelUpdateOne {
	cuo.mutation.SetImagePath(s)
	return cuo
}

// SetRetention sets the "retention" field.
func (cuo *ChannelUpdateOne) SetRetention(b bool) *ChannelUpdateOne {
	cuo.mutation.SetRetention(b)
	return cuo
}

// SetNillableRetention sets the "retention" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableRetention(b *bool) *ChannelUpdateOne {
	if b != nil {
		cuo.SetRetention(*b)
	}
	return cuo
}

// ClearRetention clears the value of the "retention" field.
func (cuo *ChannelUpdateOne) ClearRetention() *ChannelUpdateOne {
	cuo.mutation.ClearRetention()
	return cuo
}

// SetRetentionDays sets the "retention_days" field.
func (cuo *ChannelUpdateOne) SetRetentionDays(s string) *ChannelUpdateOne {
	cuo.mutation.SetRetentionDays(s)
	return cuo
}

// SetNillableRetentionDays sets the "retention_days" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableRetentionDays(s *string) *ChannelUpdateOne {
	if s != nil {
		cuo.SetRetentionDays(*s)
	}
	return cuo
}

// ClearRetentionDays clears the value of the "retention_days" field.
func (cuo *ChannelUpdateOne) ClearRetentionDays() *ChannelUpdateOne {
	cuo.mutation.ClearRetentionDays()
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ChannelUpdateOne) SetUpdatedAt(t time.Time) *ChannelUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// AddVodIDs adds the "vods" edge to the Vod entity by IDs.
func (cuo *ChannelUpdateOne) AddVodIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.AddVodIDs(ids...)
	return cuo
}

// AddVods adds the "vods" edges to the Vod entity.
func (cuo *ChannelUpdateOne) AddVods(v ...*Vod) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.AddVodIDs(ids...)
}

// AddLiveIDs adds the "live" edge to the Live entity by IDs.
func (cuo *ChannelUpdateOne) AddLiveIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.AddLiveIDs(ids...)
	return cuo
}

// AddLive adds the "live" edges to the Live entity.
func (cuo *ChannelUpdateOne) AddLive(l ...*Live) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cuo.AddLiveIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cuo *ChannelUpdateOne) Mutation() *ChannelMutation {
	return cuo.mutation
}

// ClearVods clears all "vods" edges to the Vod entity.
func (cuo *ChannelUpdateOne) ClearVods() *ChannelUpdateOne {
	cuo.mutation.ClearVods()
	return cuo
}

// RemoveVodIDs removes the "vods" edge to Vod entities by IDs.
func (cuo *ChannelUpdateOne) RemoveVodIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.RemoveVodIDs(ids...)
	return cuo
}

// RemoveVods removes "vods" edges to Vod entities.
func (cuo *ChannelUpdateOne) RemoveVods(v ...*Vod) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cuo.RemoveVodIDs(ids...)
}

// ClearLive clears all "live" edges to the Live entity.
func (cuo *ChannelUpdateOne) ClearLive() *ChannelUpdateOne {
	cuo.mutation.ClearLive()
	return cuo
}

// RemoveLiveIDs removes the "live" edge to Live entities by IDs.
func (cuo *ChannelUpdateOne) RemoveLiveIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.RemoveLiveIDs(ids...)
	return cuo
}

// RemoveLive removes "live" edges to Live entities.
func (cuo *ChannelUpdateOne) RemoveLive(l ...*Live) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cuo.RemoveLiveIDs(ids...)
}

// Where appends a list predicates to the ChannelUpdate builder.
func (cuo *ChannelUpdateOne) Where(ps ...predicate.Channel) *ChannelUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChannelUpdateOne) Select(field string, fields ...string) *ChannelUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Channel entity.
func (cuo *ChannelUpdateOne) Save(ctx context.Context) (*Channel, error) {
	cuo.defaults()
	return withHooks[*Channel, ChannelMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChannelUpdateOne) SaveX(ctx context.Context) *Channel {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChannelUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChannelUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ChannelUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := channel.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *ChannelUpdateOne) sqlSave(ctx context.Context) (_node *Channel, err error) {
	_spec := sqlgraph.NewUpdateSpec(channel.Table, channel.Columns, sqlgraph.NewFieldSpec(channel.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Channel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, channel.FieldID)
		for _, f := range fields {
			if !channel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != channel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.ExtID(); ok {
		_spec.SetField(channel.FieldExtID, field.TypeString, value)
	}
	if cuo.mutation.ExtIDCleared() {
		_spec.ClearField(channel.FieldExtID, field.TypeString)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(channel.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.DisplayName(); ok {
		_spec.SetField(channel.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ImagePath(); ok {
		_spec.SetField(channel.FieldImagePath, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Retention(); ok {
		_spec.SetField(channel.FieldRetention, field.TypeBool, value)
	}
	if cuo.mutation.RetentionCleared() {
		_spec.ClearField(channel.FieldRetention, field.TypeBool)
	}
	if value, ok := cuo.mutation.RetentionDays(); ok {
		_spec.SetField(channel.FieldRetentionDays, field.TypeString, value)
	}
	if cuo.mutation.RetentionDaysCleared() {
		_spec.ClearField(channel.FieldRetentionDays, field.TypeString)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(channel.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedVodsIDs(); len(nodes) > 0 && !cuo.mutation.VodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.VodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.LiveCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.LiveTable,
			Columns: []string{channel.LiveColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(live.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedLiveIDs(); len(nodes) > 0 && !cuo.mutation.LiveCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.LiveTable,
			Columns: []string{channel.LiveColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(live.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.LiveIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.LiveTable,
			Columns: []string{channel.LiveColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(live.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Channel{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
