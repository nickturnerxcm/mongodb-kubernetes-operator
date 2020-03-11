package replica_set_readiness_probe

import (
	"testing"

	mdbv1 "github.com/mongodb/mongodb-kubernetes-operator/pkg/apis/mongodb/v1"
	e2eutil "github.com/mongodb/mongodb-kubernetes-operator/test/e2e"
	"github.com/mongodb/mongodb-kubernetes-operator/test/e2e/mongodbtests"
	f "github.com/operator-framework/operator-sdk/pkg/test"
)

func TestMain(m *testing.M) {
	f.MainEntry(m)
}

func TestReplicaSetReadinessProbeScaling(t *testing.T) {
	ctx := f.NewTestCtx(t)
	defer ctx.Cleanup()

	// register our types with the testing framework
	if err := e2eutil.RegisterTypesWithFramework(&mdbv1.MongoDB{}); err != nil {
		t.Fatal(err)
	}

	mdb := e2eutil.NewTestMongoDB()
	t.Run("Create MongoDB Resource", mongodbtests.CreateResource(mdb, ctx))
	t.Run("Perform BasicFunctionality Checks", mongodbtests.BasicFunctionality(mdb))
	t.Run("Test Basic Connectivity", mongodbtests.BasicConnectivity(mdb))
	t.Run("Delete Pod", mongodbtests.DeletePod(mdb, 0))
	t.Run("Test Replica Set Recovers", mongodbtests.BasicFunctionality(mdb))
	t.Run("Test Recovered Replica Set Connectivity", mongodbtests.BasicConnectivity(mdb))
}
