package getuuid

import(
	"github.com/satori/go.uuid"
	_ "fmt"
)

//	Acquire V4UUId
func GetV4() uuid.UUID{
	u1 := uuid.Must(uuid.NewV4())
	return u1

}