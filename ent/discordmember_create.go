// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lukepeltier/discordle/ent/discordmember"
	"github.com/lukepeltier/discordle/ent/message"
)

// DiscordMemberCreate is the builder for creating a DiscordMember entity.
type DiscordMemberCreate struct {
	config
	mutation *DiscordMemberMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (dmc *DiscordMemberCreate) SetUsername(s string) *DiscordMemberCreate {
	dmc.mutation.SetUsername(s)
	return dmc
}

// SetNicknames sets the "nicknames" field.
func (dmc *DiscordMemberCreate) SetNicknames(s []string) *DiscordMemberCreate {
	dmc.mutation.SetNicknames(s)
	return dmc
}

// SetDiscordID sets the "discord_id" field.
func (dmc *DiscordMemberCreate) SetDiscordID(s string) *DiscordMemberCreate {
	dmc.mutation.SetDiscordID(s)
	return dmc
}

// SetBlacklisted sets the "blacklisted" field.
func (dmc *DiscordMemberCreate) SetBlacklisted(b bool) *DiscordMemberCreate {
	dmc.mutation.SetBlacklisted(b)
	return dmc
}

// SetNillableBlacklisted sets the "blacklisted" field if the given value is not nil.
func (dmc *DiscordMemberCreate) SetNillableBlacklisted(b *bool) *DiscordMemberCreate {
	if b != nil {
		dmc.SetBlacklisted(*b)
	}
	return dmc
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (dmc *DiscordMemberCreate) AddMessageIDs(ids ...int) *DiscordMemberCreate {
	dmc.mutation.AddMessageIDs(ids...)
	return dmc
}

// AddMessages adds the "messages" edges to the Message entity.
func (dmc *DiscordMemberCreate) AddMessages(m ...*Message) *DiscordMemberCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return dmc.AddMessageIDs(ids...)
}

// Mutation returns the DiscordMemberMutation object of the builder.
func (dmc *DiscordMemberCreate) Mutation() *DiscordMemberMutation {
	return dmc.mutation
}

// Save creates the DiscordMember in the database.
func (dmc *DiscordMemberCreate) Save(ctx context.Context) (*DiscordMember, error) {
	dmc.defaults()
	return withHooks(ctx, dmc.sqlSave, dmc.mutation, dmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dmc *DiscordMemberCreate) SaveX(ctx context.Context) *DiscordMember {
	v, err := dmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmc *DiscordMemberCreate) Exec(ctx context.Context) error {
	_, err := dmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmc *DiscordMemberCreate) ExecX(ctx context.Context) {
	if err := dmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dmc *DiscordMemberCreate) defaults() {
	if _, ok := dmc.mutation.Blacklisted(); !ok {
		v := discordmember.DefaultBlacklisted
		dmc.mutation.SetBlacklisted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dmc *DiscordMemberCreate) check() error {
	if _, ok := dmc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "DiscordMember.username"`)}
	}
	if v, ok := dmc.mutation.Username(); ok {
		if err := discordmember.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "DiscordMember.username": %w`, err)}
		}
	}
	if _, ok := dmc.mutation.Nicknames(); !ok {
		return &ValidationError{Name: "nicknames", err: errors.New(`ent: missing required field "DiscordMember.nicknames"`)}
	}
	if _, ok := dmc.mutation.DiscordID(); !ok {
		return &ValidationError{Name: "discord_id", err: errors.New(`ent: missing required field "DiscordMember.discord_id"`)}
	}
	if _, ok := dmc.mutation.Blacklisted(); !ok {
		return &ValidationError{Name: "blacklisted", err: errors.New(`ent: missing required field "DiscordMember.blacklisted"`)}
	}
	return nil
}

func (dmc *DiscordMemberCreate) sqlSave(ctx context.Context) (*DiscordMember, error) {
	if err := dmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dmc.mutation.id = &_node.ID
	dmc.mutation.done = true
	return _node, nil
}

func (dmc *DiscordMemberCreate) createSpec() (*DiscordMember, *sqlgraph.CreateSpec) {
	var (
		_node = &DiscordMember{config: dmc.config}
		_spec = sqlgraph.NewCreateSpec(discordmember.Table, sqlgraph.NewFieldSpec(discordmember.FieldID, field.TypeInt))
	)
	if value, ok := dmc.mutation.Username(); ok {
		_spec.SetField(discordmember.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := dmc.mutation.Nicknames(); ok {
		_spec.SetField(discordmember.FieldNicknames, field.TypeJSON, value)
		_node.Nicknames = value
	}
	if value, ok := dmc.mutation.DiscordID(); ok {
		_spec.SetField(discordmember.FieldDiscordID, field.TypeString, value)
		_node.DiscordID = value
	}
	if value, ok := dmc.mutation.Blacklisted(); ok {
		_spec.SetField(discordmember.FieldBlacklisted, field.TypeBool, value)
		_node.Blacklisted = value
	}
	if nodes := dmc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discordmember.MessagesTable,
			Columns: []string{discordmember.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DiscordMemberCreateBulk is the builder for creating many DiscordMember entities in bulk.
type DiscordMemberCreateBulk struct {
	config
	err      error
	builders []*DiscordMemberCreate
}

// Save creates the DiscordMember entities in the database.
func (dmcb *DiscordMemberCreateBulk) Save(ctx context.Context) ([]*DiscordMember, error) {
	if dmcb.err != nil {
		return nil, dmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dmcb.builders))
	nodes := make([]*DiscordMember, len(dmcb.builders))
	mutators := make([]Mutator, len(dmcb.builders))
	for i := range dmcb.builders {
		func(i int, root context.Context) {
			builder := dmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiscordMemberMutation)
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
					_, err = mutators[i+1].Mutate(root, dmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, dmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dmcb *DiscordMemberCreateBulk) SaveX(ctx context.Context) []*DiscordMember {
	v, err := dmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmcb *DiscordMemberCreateBulk) Exec(ctx context.Context) error {
	_, err := dmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmcb *DiscordMemberCreateBulk) ExecX(ctx context.Context) {
	if err := dmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
