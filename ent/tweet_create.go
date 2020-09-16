// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/francismarcus/entgql/ent/tweet"
	"github.com/francismarcus/entgql/ent/user"
)

// TweetCreate is the builder for creating a Tweet entity.
type TweetCreate struct {
	config
	mutation *TweetMutation
	hooks    []Hook
}

// SetMessage sets the message field.
func (tc *TweetCreate) SetMessage(s string) *TweetCreate {
	tc.mutation.SetMessage(s)
	return tc
}

// SetCreatedAt sets the created_at field.
func (tc *TweetCreate) SetCreatedAt(t time.Time) *TweetCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (tc *TweetCreate) SetNillableCreatedAt(t *time.Time) *TweetCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetCreatorID sets the creator edge to User by id.
func (tc *TweetCreate) SetCreatorID(id int) *TweetCreate {
	tc.mutation.SetCreatorID(id)
	return tc
}

// SetNillableCreatorID sets the creator edge to User by id if the given value is not nil.
func (tc *TweetCreate) SetNillableCreatorID(id *int) *TweetCreate {
	if id != nil {
		tc = tc.SetCreatorID(*id)
	}
	return tc
}

// SetCreator sets the creator edge to User.
func (tc *TweetCreate) SetCreator(u *User) *TweetCreate {
	return tc.SetCreatorID(u.ID)
}

// Mutation returns the TweetMutation object of the builder.
func (tc *TweetCreate) Mutation() *TweetMutation {
	return tc.mutation
}

// Save creates the Tweet in the database.
func (tc *TweetCreate) Save(ctx context.Context) (*Tweet, error) {
	if err := tc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Tweet
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TweetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TweetCreate) SaveX(ctx context.Context) *Tweet {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TweetCreate) preSave() error {
	if _, ok := tc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New("ent: missing required field \"message\"")}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := tweet.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	return nil
}

func (tc *TweetCreate) sqlSave(ctx context.Context) (*Tweet, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TweetCreate) createSpec() (*Tweet, *sqlgraph.CreateSpec) {
	var (
		t     = &Tweet{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tweet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tweet.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tweet.FieldMessage,
		})
		t.Message = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tweet.FieldCreatedAt,
		})
		t.CreatedAt = value
	}
	if nodes := tc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.CreatorTable,
			Columns: []string{tweet.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return t, _spec
}

// TweetCreateBulk is the builder for creating a bulk of Tweet entities.
type TweetCreateBulk struct {
	config
	builders []*TweetCreate
}

// Save creates the Tweet entities in the database.
func (tcb *TweetCreateBulk) Save(ctx context.Context) ([]*Tweet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tweet, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*TweetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (tcb *TweetCreateBulk) SaveX(ctx context.Context) []*Tweet {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
