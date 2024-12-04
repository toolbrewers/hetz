package repository

import (
	"hetz-client/internal/models"
	"testing"
)

var repo *Repository

const roleName = "test-role"

func openRepo() *Repository {
	localRepo := New("../../db/client.db")
	repo = localRepo
	return localRepo
}

func TestRole(t *testing.T) {
	openRepo()
	defer repo.Close()
	t.Run("CreateRole", testCreateRole)
	t.Run("GetRole", testGetRole)
	t.Run("DeleteRole", testDeleteRole)
}

func testCreateRole(t *testing.T) {

	role := &models.Role{
		Name: roleName,
	}

	name, err := repo.CreateRole(role)
	if err != nil {
		t.Fatalf("failed to create role: %v", err)
	}

	if roleName != role.Name {
		t.Fatalf("expected name to be %s, got %s", role.Name, name)
	}
}

func testGetRole(t *testing.T) {

	role, err := repo.GetRole(roleName)
	if err != nil {
		t.Fatalf("failed to get role: %v", err)
	}

	if role.Name != roleName {
		t.Fatalf("expected name to be %s, got %s", roleName, role.Name)
	}
}

func testDeleteRole(t *testing.T) {

	err := repo.DeleteRole(roleName)
	if err != nil {
		t.Fatalf("failed to delete role: %v", err)
	}
}

func BenchmarkCreateRole(b *testing.B) {
	repo = openRepo()
	defer repo.Close()

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		repo.CreateRole(&models.Role{
			Name: roleName,
		})
		b.StopTimer()
		repo.DeleteRole(roleName)
	}
}

func BenchmarkGetRole(b *testing.B) {
	repo = openRepo()
	defer repo.Close()

	repo.CreateRole(&models.Role{
		Name: roleName,
	})

	for i := 0; i < b.N; i++ {
		repo.GetRole(roleName)
	}
	repo.DeleteRole(roleName)
}

func BenchmarkDeleteRole(b *testing.B) {
	repo = openRepo()
	defer repo.Close()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		repo.CreateRole(&models.Role{
			Name: roleName,
		})
		b.StartTimer()
		repo.DeleteRole(roleName)
	}
}
