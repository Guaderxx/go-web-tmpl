package mixin

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/Guaderxx/gowebtmpl/ent/hook"
	"github.com/sony/sonyflake"
)

type IDMixin struct {
	mixin.Schema
}

// Fields of the Mixin.
func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
	}
}

// Hooks of the Mixin.
func (IDMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(IDHook(), ent.OpCreate),
	}
}

func IDHook() ent.Hook {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	type IDSetter interface {
		SetID(uint64)
	}

	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			is, ok := m.(IDSetter)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation %T", m)
			}
			id, err := sf.NextID()
			if err != nil {
				return nil, err
			}
			is.SetID(id)
			return next.Mutate(ctx, m)
		})
	}
}
