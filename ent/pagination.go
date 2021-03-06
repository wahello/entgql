// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/francismarcus/entgql/ent/program"
	"github.com/francismarcus/entgql/ent/tweet"
	"github.com/francismarcus/entgql/ent/user"
	"github.com/ugorji/go/codec"
)

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID int
}

// ErrInvalidPagination error is returned when paginating with invalid parameters.
var ErrInvalidPagination = errors.New("ent: invalid pagination parameters")

var quote = []byte(`"`)

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.StdEncoding, w)
	defer wc.Close()
	_ = codec.NewEncoder(wc, &codec.MsgpackHandle{}).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := codec.NewDecoder(
		base64.NewDecoder(
			base64.StdEncoding,
			strings.NewReader(s),
		),
		&codec.MsgpackHandle{},
	).Decode(c); err != nil {
		return fmt.Errorf("decode cursor: %w", err)
	}
	return nil
}

// ProgramEdge is the edge representation of Program.
type ProgramEdge struct {
	Node   *Program `json:"node"`
	Cursor Cursor   `json:"cursor"`
}

// ProgramConnection is the connection containing edges to Program.
type ProgramConnection struct {
	Edges    []*ProgramEdge `json:"edges"`
	PageInfo PageInfo       `json:"pageInfo"`
}

// Paginate executes the query and returns a relay based cursor connection to Program.
func (pr *ProgramQuery) Paginate(ctx context.Context, after *Cursor, first *int, before *Cursor, last *int) (*ProgramConnection, error) {
	if first != nil && last != nil {
		return nil, ErrInvalidPagination
	}
	if first != nil {
		if *first == 0 {
			return &ProgramConnection{
				Edges: []*ProgramEdge{},
			}, nil
		} else if *first < 0 {
			return nil, ErrInvalidPagination
		}
	}
	if last != nil {
		if *last == 0 {
			return &ProgramConnection{
				Edges: []*ProgramEdge{},
			}, nil
		} else if *last < 0 {
			return nil, ErrInvalidPagination
		}
	}

	if after != nil {
		pr = pr.Where(program.IDGT(after.ID))
	}
	if before != nil {
		pr = pr.Where(program.IDLT(before.ID))
	}
	if first != nil {
		pr = pr.Order(Asc(program.FieldID)).Limit(*first + 1)
	}
	if last != nil {
		pr = pr.Order(Desc(program.FieldID)).Limit(*last + 1)
	}

	nodes, err := pr.All(ctx)
	if err != nil || len(nodes) == 0 {
		return &ProgramConnection{
			Edges: []*ProgramEdge{},
		}, err
	}
	if last != nil {
		for left, right := 0, len(nodes)-1; left < right; left, right = left+1, right-1 {
			nodes[left], nodes[right] = nodes[right], nodes[left]
		}
	}

	var conn ProgramConnection
	if first != nil && len(nodes) > *first {
		conn.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && len(nodes) > *last {
		conn.PageInfo.HasPreviousPage = true
		nodes = nodes[1:]
	}
	conn.Edges = make([]*ProgramEdge, len(nodes))
	for i, node := range nodes {
		conn.Edges[i] = &ProgramEdge{
			Node: node,
			Cursor: Cursor{
				ID: node.ID,
			},
		}
	}
	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor

	return &conn, nil
}

// TweetEdge is the edge representation of Tweet.
type TweetEdge struct {
	Node   *Tweet `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// TweetConnection is the connection containing edges to Tweet.
type TweetConnection struct {
	Edges    []*TweetEdge `json:"edges"`
	PageInfo PageInfo     `json:"pageInfo"`
}

// Paginate executes the query and returns a relay based cursor connection to Tweet.
func (t *TweetQuery) Paginate(ctx context.Context, after *Cursor, first *int, before *Cursor, last *int) (*TweetConnection, error) {
	if first != nil && last != nil {
		return nil, ErrInvalidPagination
	}
	if first != nil {
		if *first == 0 {
			return &TweetConnection{
				Edges: []*TweetEdge{},
			}, nil
		} else if *first < 0 {
			return nil, ErrInvalidPagination
		}
	}
	if last != nil {
		if *last == 0 {
			return &TweetConnection{
				Edges: []*TweetEdge{},
			}, nil
		} else if *last < 0 {
			return nil, ErrInvalidPagination
		}
	}

	if after != nil {
		t = t.Where(tweet.IDGT(after.ID))
	}
	if before != nil {
		t = t.Where(tweet.IDLT(before.ID))
	}
	if first != nil {
		t = t.Order(Asc(tweet.FieldID)).Limit(*first + 1)
	}
	if last != nil {
		t = t.Order(Desc(tweet.FieldID)).Limit(*last + 1)
	}

	nodes, err := t.All(ctx)
	if err != nil || len(nodes) == 0 {
		return &TweetConnection{
			Edges: []*TweetEdge{},
		}, err
	}
	if last != nil {
		for left, right := 0, len(nodes)-1; left < right; left, right = left+1, right-1 {
			nodes[left], nodes[right] = nodes[right], nodes[left]
		}
	}

	var conn TweetConnection
	if first != nil && len(nodes) > *first {
		conn.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && len(nodes) > *last {
		conn.PageInfo.HasPreviousPage = true
		nodes = nodes[1:]
	}
	conn.Edges = make([]*TweetEdge, len(nodes))
	for i, node := range nodes {
		conn.Edges[i] = &TweetEdge{
			Node: node,
			Cursor: Cursor{
				ID: node.ID,
			},
		}
	}
	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor

	return &conn, nil
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges    []*UserEdge `json:"edges"`
	PageInfo PageInfo    `json:"pageInfo"`
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(ctx context.Context, after *Cursor, first *int, before *Cursor, last *int) (*UserConnection, error) {
	if first != nil && last != nil {
		return nil, ErrInvalidPagination
	}
	if first != nil {
		if *first == 0 {
			return &UserConnection{
				Edges: []*UserEdge{},
			}, nil
		} else if *first < 0 {
			return nil, ErrInvalidPagination
		}
	}
	if last != nil {
		if *last == 0 {
			return &UserConnection{
				Edges: []*UserEdge{},
			}, nil
		} else if *last < 0 {
			return nil, ErrInvalidPagination
		}
	}

	if after != nil {
		u = u.Where(user.IDGT(after.ID))
	}
	if before != nil {
		u = u.Where(user.IDLT(before.ID))
	}
	if first != nil {
		u = u.Order(Asc(user.FieldID)).Limit(*first + 1)
	}
	if last != nil {
		u = u.Order(Desc(user.FieldID)).Limit(*last + 1)
	}

	nodes, err := u.All(ctx)
	if err != nil || len(nodes) == 0 {
		return &UserConnection{
			Edges: []*UserEdge{},
		}, err
	}
	if last != nil {
		for left, right := 0, len(nodes)-1; left < right; left, right = left+1, right-1 {
			nodes[left], nodes[right] = nodes[right], nodes[left]
		}
	}

	var conn UserConnection
	if first != nil && len(nodes) > *first {
		conn.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && len(nodes) > *last {
		conn.PageInfo.HasPreviousPage = true
		nodes = nodes[1:]
	}
	conn.Edges = make([]*UserEdge, len(nodes))
	for i, node := range nodes {
		conn.Edges[i] = &UserEdge{
			Node: node,
			Cursor: Cursor{
				ID: node.ID,
			},
		}
	}
	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor

	return &conn, nil
}
