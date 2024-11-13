package db

import (
	"context"
	"testing"

	"github.com/mauro-moreno/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	usr1 := createRandomUser(t)
	usr2, err := testQueries.GetUser(context.Background(), usr1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, usr2)

	require.Equal(t, usr1.Username, usr2.Username)
	require.Equal(t, usr1.HashedPassword, usr2.HashedPassword)
	require.Equal(t, usr1.FullName, usr2.FullName)
	require.Equal(t, usr1.Email, usr2.Email)
	require.WithinDuration(t, usr1.PasswordChangedAt, usr2.PasswordChangedAt, 0)
	require.WithinDuration(t, usr1.CreatedAt, usr2.CreatedAt, 0)
}

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	user := User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	usr, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		FullName:       user.FullName,
		Email:          user.Email,
	})
	require.NoError(t, err)
	require.NotEmpty(t, usr)

	require.Equal(t, user.Username, usr.Username)
	require.Equal(t, user.HashedPassword, usr.HashedPassword)
	require.Equal(t, user.FullName, usr.FullName)
	require.Equal(t, user.Email, usr.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, usr.CreatedAt)

	return usr
}
