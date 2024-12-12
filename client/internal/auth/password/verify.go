package password

import "hetz-client/internal/models"

func Verify(user *models.GetUser, password string) bool {
	// TODO: @Kevin, Since you don't like having a global DB instance, advise on how to access the repository from here.
	// We could do some kind of a wrapper and call that function, but the repo is already a wrapper, and that'd be over-complicating it
	return true
}
