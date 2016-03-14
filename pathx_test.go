package pathx

import . "github.com/pkg4go/assert"
import "testing"
import "strings"

func TestResolve(t *testing.T) {
	a := A{t}
	a.Equal(strings.HasSuffix(Resolve("", "c"), "pathx/c"), true)
	a.Equal(Resolve("/a/b", "./c"), "/a/b/c")
	a.Equal(Resolve("/a/b", "../c"), "/a/c")
	a.Equal(Resolve("/a/b", "c"), "/a/b/c")
	a.Equal(Resolve("/a/b", "/c"), "/c")
	a.Equal(Resolve("", "/c"), "/c")
}
