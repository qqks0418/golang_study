package user

// https://qiita.com/ryu3/items/a2e39157bf1d55be149f
// https://engineering.mercari.com/blog/entry/how_to_use_t_parallel/
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func BenchmarkUser(b *testing.B) {
	//go test -bench .
  var id int64 = 1
  var name string = "Nakata"
  var email string = "nakata@example.com"

  b.ResetTimer()
  for i:= 0; i < b.N; i++ {
    user:= NewUser(id, name, email)
  }
  b.StopTimer()

  if order == nil {
    b.Errorf("failed NewUser()")
  }

}
*/

func TestUser(t *testing.T) {

  t.Run("success NewUser()", func(t *testing.T){
    var id int64 = 1
    var name string = "Nakata"
    var email string = "nakata@example.com"
    user:= NewUser(id, name, email)

    if user == nil {
      t.Errorf("failed NewUser()")
    }

    assert.Equal(t, id, user.ID)
    assert.Equal(t, name, user.Name)
    assert.Equal(t, email, user.Email)

    t.Logf("user: %p", user)
    t.Logf("user.ID: %d", user.ID)
    t.Logf("user.Name: %s", user.Name)
    t.Logf("user.Email: %s", user.Email)
  })

  t.Run("success NewUser()", func(t *testing.T){
    var id int64 = 2
    var name string = "Suzuki"
    var email string = "suzuki@example.com"
    user:= NewUser(id, name, email)

    if user == nil {
      t.Errorf("failed NewUser()")
    }

    assert.Equal(t, id, user.ID)
    assert.Equal(t, name, user.Name)
    assert.Equal(t, email, user.Email)

    t.Logf("user: %p", user)
    t.Logf("user.ID: %d", user.ID)
    t.Logf("user.Name: %s", user.Name)
    t.Logf("user.Email: %s", user.Email)
  })
}