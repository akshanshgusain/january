package january

import (
	"github.com/CloudyKit/jet/v6"
	"os"
	"testing"
)

/*
When you run your tests using the go test command, it will discover and execute the test functions in the package,
and the TestMain function will be called before and after the tests, as specified in your code.
*/

// jet views
var jv = jet.NewSet(
	jet.NewOSFileSystemLoader("./testData/views"),
	jet.InDevelopmentMode(),
)

var te = TemplateEngine{
	TemplateEngine: "",
	RootPath:       "",
	JetViews:       jv,
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
