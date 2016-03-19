package TSAuto//_test
import(
	"testing"
)
func Test_getAdmin(t *testing.T) {
	count := Insert()
	if (count != 1) {
		t.Error("insert  data error")
	}
}