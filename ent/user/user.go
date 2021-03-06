// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldFollowsCount holds the string denoting the follows_count field in the database.
	FieldFollowsCount = "follows_count"
	// FieldFollowersCount holds the string denoting the followers_count field in the database.
	FieldFollowersCount = "followers_count"
	// FieldTweetsCount holds the string denoting the tweets_count field in the database.
	FieldTweetsCount = "tweets_count"

	// EdgeFollowers holds the string denoting the followers edge name in mutations.
	EdgeFollowers = "followers"
	// EdgeFollowing holds the string denoting the following edge name in mutations.
	EdgeFollowing = "following"
	// EdgePrograms holds the string denoting the programs edge name in mutations.
	EdgePrograms = "programs"
	// EdgeTweets holds the string denoting the tweets edge name in mutations.
	EdgeTweets = "tweets"

	// Table holds the table name of the user in the database.
	Table = "users"
	// FollowersTable is the table the holds the followers relation/edge. The primary key declared below.
	FollowersTable = "user_following"
	// FollowingTable is the table the holds the following relation/edge. The primary key declared below.
	FollowingTable = "user_following"
	// ProgramsTable is the table the holds the programs relation/edge.
	ProgramsTable = "programs"
	// ProgramsInverseTable is the table name for the Program entity.
	// It exists in this package in order to avoid circular dependency with the "program" package.
	ProgramsInverseTable = "programs"
	// ProgramsColumn is the table column denoting the programs relation/edge.
	ProgramsColumn = "user_programs"
	// TweetsTable is the table the holds the tweets relation/edge.
	TweetsTable = "tweets"
	// TweetsInverseTable is the table name for the Tweet entity.
	// It exists in this package in order to avoid circular dependency with the "tweet" package.
	TweetsInverseTable = "tweets"
	// TweetsColumn is the table column denoting the tweets relation/edge.
	TweetsColumn = "user_tweets"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldEmail,
	FieldPassword,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldFollowsCount,
	FieldFollowersCount,
	FieldTweetsCount,
}

var (
	// FollowersPrimaryKey and FollowersColumn2 are the table columns denoting the
	// primary key for the followers relation (M2M).
	FollowersPrimaryKey = []string{"user_id", "follower_id"}
	// FollowingPrimaryKey and FollowingColumn2 are the table columns denoting the
	// primary key for the following relation (M2M).
	FollowingPrimaryKey = []string{"user_id", "follower_id"}
)

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultFollowsCount holds the default value on creation for the follows_count field.
	DefaultFollowsCount int
	// DefaultFollowersCount holds the default value on creation for the followers_count field.
	DefaultFollowersCount int
	// DefaultTweetsCount holds the default value on creation for the tweets_count field.
	DefaultTweetsCount int
)
