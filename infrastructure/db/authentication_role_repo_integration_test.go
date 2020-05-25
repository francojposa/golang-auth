package db

import (
	"github.com/stretchr/testify/assert"
	"golang-auth/usecases/repos"
	"golang-auth/usecases/resources"
	"testing"
)

func TestPGAuthNRoleRepo(t *testing.T) {
	assertions := assert.New(t)

	sqlxDB, closeDB := SetUpDB(t)
	defer closeDB(t, sqlxDB)
	roleRepo, _ := SetUpAuthNRoleRepo(t, sqlxDB)

	role := resources.NewAuthNRole("staff")

	t.Run("create role", func(t *testing.T) {
		createdRole, _ := roleRepo.Create(role)
		assertAuthNRole(assertions, role, createdRole)
	})

	t.Run("create already existing role - error", func(t *testing.T) {
		alreadyCreatedRole, err := roleRepo.Create(role)
		assertions.Nil(alreadyCreatedRole)
		assertions.IsType(&repos.DuplicateAuthNRole{}, err)
	})

	t.Run("get role", func(t *testing.T) {
		retrievedRole, err := roleRepo.Get(role.Role)
		if err != nil {
			t.Error(err)
		}
		assertAuthNRole(assertions, role, retrievedRole)
	})

	t.Run("get nonexistent role - error", func(t *testing.T) {
		nonexistentRole, err := roleRepo.Get("xxx")
		assertions.Nil(nonexistentRole, "expected nil struct, got: %q", nonexistentRole)
		assertions.IsType(&repos.AuthNRoleNotFoundError{}, err)
	})
}

func assertAuthNRole(a *assert.Assertions, want, got *resources.AuthNRole) {
	a.Equal(
		want, got, "expected equivalent structs, want: %q, got: %q", want, got,
	)
}
