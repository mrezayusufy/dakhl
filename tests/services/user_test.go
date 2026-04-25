package services

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"dakhl/app/services"
	"dakhl/tests"
)

/*
*********************************************
Please init .env file before running the test.
**********************************************
*/
type UserTestSuite struct {
	suite.Suite
	tests.TestCase
	user services.User
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, &UserTestSuite{
		user: services.NewUserImpl(),
	})
}

// SetupTest will run before each test in the suite.
func (s *UserTestSuite) SetupTest() {
}

// TearDownTest will run after each test in the suite.
func (s *UserTestSuite) TearDownTest() {
}

func (s *UserTestSuite) TestCreateByConfig() {
	user, err := s.user.Create("feature")
	s.Nil(err)
	s.True(user.ID.Len())
	s.Equal("feature", user.Name)
}
