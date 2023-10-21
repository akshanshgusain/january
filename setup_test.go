package january

import (
	"github.com/CloudyKit/jet/v6"
	"os"
	"testing"
)

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
