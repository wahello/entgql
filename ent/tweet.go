// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/francismarcus/entgql/ent/tweet"
	"github.com/francismarcus/entgql/ent/user"
)

// Tweet is the model entity for the Tweet schema.
type Tweet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TweetQuery when eager-loading is set.
	Edges       TweetEdges `json:"edges"`
	user_tweets *int
}

// TweetEdges holds the relations/edges for other nodes in the graph.
type TweetEdges struct {
	// Creator holds the value of the creator edge.
	Creator *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TweetEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Creator == nil {
			// The edge creator was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tweet) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // message
		&sql.NullTime{},   // created_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Tweet) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_tweets
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tweet fields.
func (t *Tweet) assignValues(values ...interface{}) error {
	if m, n := len(values), len(tweet.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field message", values[0])
	} else if value.Valid {
		t.Message = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[1])
	} else if value.Valid {
		t.CreatedAt = value.Time
	}
	values = values[2:]
	if len(values) == len(tweet.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_tweets", value)
		} else if value.Valid {
			t.user_tweets = new(int)
			*t.user_tweets = int(value.Int64)
		}
	}
	return nil
}

// QueryCreator queries the creator edge of the Tweet.
func (t *Tweet) QueryCreator() *UserQuery {
	return (&TweetClient{config: t.config}).QueryCreator(t)
}

// Update returns a builder for updating this Tweet.
// Note that, you need to call Tweet.Unwrap() before calling this method, if this Tweet
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tweet) Update() *TweetUpdateOne {
	return (&TweetClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Tweet) Unwrap() *Tweet {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tweet is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tweet) String() string {
	var builder strings.Builder
	builder.WriteString("Tweet(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", message=")
	builder.WriteString(t.Message)
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Tweets is a parsable slice of Tweet.
type Tweets []*Tweet

func (t Tweets) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}