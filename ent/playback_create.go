// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/playback"
	"github.com/zibbp/ganymede/internal/utils"
)

// PlaybackCreate is the builder for creating a Playback entity.
type PlaybackCreate struct {
	config
	mutation *PlaybackMutation
	hooks    []Hook
}

// SetVodID sets the "vod_id" field.
func (pc *PlaybackCreate) SetVodID(u uuid.UUID) *PlaybackCreate {
	pc.mutation.SetVodID(u)
	return pc
}

// SetUserID sets the "user_id" field.
func (pc *PlaybackCreate) SetUserID(u uuid.UUID) *PlaybackCreate {
	pc.mutation.SetUserID(u)
	return pc
}

// SetTime sets the "time" field.
func (pc *PlaybackCreate) SetTime(i int) *PlaybackCreate {
	pc.mutation.SetTime(i)
	return pc
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (pc *PlaybackCreate) SetNillableTime(i *int) *PlaybackCreate {
	if i != nil {
		pc.SetTime(*i)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *PlaybackCreate) SetStatus(us utils.PlaybackStatus) *PlaybackCreate {
	pc.mutation.SetStatus(us)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *PlaybackCreate) SetNillableStatus(us *utils.PlaybackStatus) *PlaybackCreate {
	if us != nil {
		pc.SetStatus(*us)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PlaybackCreate) SetUpdatedAt(t time.Time) *PlaybackCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PlaybackCreate) SetNillableUpdatedAt(t *time.Time) *PlaybackCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PlaybackCreate) SetCreatedAt(t time.Time) *PlaybackCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PlaybackCreate) SetNillableCreatedAt(t *time.Time) *PlaybackCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PlaybackCreate) SetID(u uuid.UUID) *PlaybackCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PlaybackCreate) SetNillableID(u *uuid.UUID) *PlaybackCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// Mutation returns the PlaybackMutation object of the builder.
func (pc *PlaybackCreate) Mutation() *PlaybackMutation {
	return pc.mutation
}

// Save creates the Playback in the database.
func (pc *PlaybackCreate) Save(ctx context.Context) (*Playback, error) {
	pc.defaults()
	return withHooks[*Playback, PlaybackMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlaybackCreate) SaveX(ctx context.Context) *Playback {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlaybackCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlaybackCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PlaybackCreate) defaults() {
	if _, ok := pc.mutation.Time(); !ok {
		v := playback.DefaultTime
		pc.mutation.SetTime(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := playback.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := playback.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := playback.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := playback.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlaybackCreate) check() error {
	if _, ok := pc.mutation.VodID(); !ok {
		return &ValidationError{Name: "vod_id", err: errors.New(`ent: missing required field "Playback.vod_id"`)}
	}
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Playback.user_id"`)}
	}
	if _, ok := pc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "Playback.time"`)}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := playback.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Playback.status": %w`, err)}
		}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Playback.updated_at"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Playback.created_at"`)}
	}
	return nil
}

func (pc *PlaybackCreate) sqlSave(ctx context.Context) (*Playback, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PlaybackCreate) createSpec() (*Playback, *sqlgraph.CreateSpec) {
	var (
		_node = &Playback{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(playback.Table, sqlgraph.NewFieldSpec(playback.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.VodID(); ok {
		_spec.SetField(playback.FieldVodID, field.TypeUUID, value)
		_node.VodID = value
	}
	if value, ok := pc.mutation.UserID(); ok {
		_spec.SetField(playback.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := pc.mutation.Time(); ok {
		_spec.SetField(playback.FieldTime, field.TypeInt, value)
		_node.Time = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(playback.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(playback.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(playback.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// PlaybackCreateBulk is the builder for creating many Playback entities in bulk.
type PlaybackCreateBulk struct {
	config
	builders []*PlaybackCreate
}

// Save creates the Playback entities in the database.
func (pcb *PlaybackCreateBulk) Save(ctx context.Context) ([]*Playback, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Playback, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlaybackMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlaybackCreateBulk) SaveX(ctx context.Context) []*Playback {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlaybackCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlaybackCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
