package retention

import (
	"flag"
	"testing"

	"github.com/mattermost/mattermost-server/v5/shared/mlog"
	"github.com/mattermost/mattermost-server/v5/testlib"
	"github.com/mattermost/mattermost-server/v5/api4"
)

var replicaFlag bool

// export MM_SERVER_PATH=~/go/src/mattermost-server
func TestMain(m *testing.M) {
	if f := flag.Lookup("mysql-replica"); f == nil {
		flag.BoolVar(&replicaFlag, "mysql-replica", false, "")
		flag.Parse()
	}

	var options = testlib.HelperOptions{
		EnableStore:     true,
		EnableResources: true,
		WithReadReplica: replicaFlag,
	}

	mlog.DisableZap()
        mlog.Debug("Entering TestMain 1")

        mainHelper := testlib.NewMainHelperWithOptions(&options)
        api4.SetMainHelper(mainHelper)
	defer mainHelper.Close()

        mlog.Debug("Entering TestMain 2")

	mainHelper.Main(m)
}

// import (
// 	"testing"
//         // "fmt"
// 
// 	"github.com/mattermost/mattermost-server/v5/shared/mlog"
// 	// "github.com/mattermost/mattermost-server/v5/store/sqlstore"
// 	"github.com/mattermost/mattermost-server/v5/testlib"
// 	// "github.com/mattermost/mattermost-server/v5/imports/retention"
// )
// 
// var mainHelper *testlib.MainHelper
// 
// func TestMain(m *testing.M) {
// 	mlog.DisableZap()
// 	mainHelper = testlib.NewMainHelperWithOptions(nil)
// 	defer mainHelper.Close()
// 
//         initStores()
//         // fmt.Println("initStores OK.")
// 
// 	mainHelper.Main(m)
//         tearDownStores()
// }
