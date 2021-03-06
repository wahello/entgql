// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/francismarcus/entgql/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// FollowsCount holds the value of the "follows_count" field.
	FollowsCount int `json:"follows_count,omitempty"`
	// FollowersCount holds the value of the "followers_count" field.
	FollowersCount int `json:"followers_count,omitempty"`
	// TweetsCount holds the value of the "tweets_count" field.
	TweetsCount int `json:"tweets_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Followers holds the value of the followers edge.
	Followers []*User
	// Following holds the value of the following edge.
	Following []*User
	// Programs holds the value of the programs edge.
	Programs []*Program
	// Tweets holds the value of the tweets edge.
	Tweets []*Tweet
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// FollowersOrErr returns the Followers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Followers, nil
	}
	return nil, &NotLoadedError{edge: "followers"}
}

// FollowingOrErr returns the Following value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowingOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Following, nil
	}
	return nil, &NotLoadedError{edge: "following"}
}

// ProgramsOrErr returns the Programs value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ProgramsOrErr() ([]*Program, error) {
	if e.loadedTypes[2] {
		return e.Programs, nil
	}
	return nil, &NotLoadedError{edge: "programs"}
}

// TweetsOrErr returns the Tweets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TweetsOrErr() ([]*Tweet, error) {
	if e.loadedTypes[3] {
		return e.Tweets, nil
	}
	return nil, &NotLoadedError{edge: "tweets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // username
		&sql.NullString{}, // email
		&sql.NullString{}, // password
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullInt64{},  // follows_count
		&sql.NullInt64{},  // followers_count
		&sql.NullInt64{},  // tweets_count
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field username", values[0])
	} else if value.Valid {
		u.Username = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[1])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[2])
	} else if value.Valid {
		u.Password = value.String
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[3])
	} else if value.Valid {
		u.CreatedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[4])
	} else if value.Valid {
		u.UpdatedAt = value.Time
	}
	if value, ok := values[5].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field follows_count", values[5])
	} else if value.Valid {
		u.FollowsCount = int(value.Int64)
	}
	if value, ok := values[6].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field followers_count", values[6])
	} else if value.Valid {
		u.FollowersCount = int(value.Int64)
	}
	if value, ok := values[7].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field tweets_count", values[7])
	} else if value.Valid {
		u.TweetsCount = int(value.Int64)
	}
	return nil
}

// QueryFollowers queries the followers edge of the User.
func (u *User) QueryFollowers() *UserQuery {
	return (&UserClient{config: u.config}).QueryFollowers(u)
}

// QueryFollowing queries the following edge of the User.
func (u *User) QueryFollowing() *UserQuery {
	return (&UserClient{config: u.config}).QueryFollowing(u)
}

// QueryPrograms queries the programs edge of the User.
func (u *User) QueryPrograms() *ProgramQuery {
	return (&UserClient{config: u.config}).QueryPrograms(u)
}

// QueryTweets queries the tweets edge of the User.
func (u *User) QueryTweets() *TweetQuery {
	return (&UserClient{config: u.config}).QueryTweets(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", password=")
	builder.WriteString(u.Password)
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", follows_count=")
	builder.WriteString(fmt.Sprintf("%v", u.FollowsCount))
	builder.WriteString(", followers_count=")
	builder.WriteString(fmt.Sprintf("%v", u.FollowersCount))
	builder.WriteString(", tweets_count=")
	builder.WriteString(fmt.Sprintf("%v", u.TweetsCount))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
