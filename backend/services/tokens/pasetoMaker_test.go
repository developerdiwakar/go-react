package tokens

import (
	"strconv"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(faker.New().RandomStringWithLength(32))
	require.NoError(t, err)

	userId := faker.New().Int16Between(1, 9999)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(int(userId), duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.Verifytoken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, strconv.Itoa(int(userId)), payload.Subject)
	require.WithinDuration(t, issuedAt, payload.IssuedAt.Time, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiresAt.Time, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(faker.New().RandomStringWithLength(32))
	require.NoError(t, err)

	userId := faker.New().Int16Between(1, 9999)
	token, err := maker.CreateToken(int(userId), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.Verifytoken(token)
	require.Error(t, err)
	require.ErrorContains(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
