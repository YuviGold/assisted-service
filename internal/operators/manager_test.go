package operators_test

import (
	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/host"
	"github.com/openshift/assisted-service/internal/operators"
	"github.com/openshift/assisted-service/internal/operators/lso"
	"github.com/openshift/assisted-service/internal/operators/ocs"
	"github.com/openshift/assisted-service/models"
	"github.com/sirupsen/logrus"
)

var (
	cluster     *common.Cluster
	ctrl        *gomock.Controller
	log         = logrus.New()
	mockHostAPI *host.MockAPI
	manager     operators.API
)

var _ = BeforeEach(func() {
	// create simple cluster
	clusterID := strfmt.UUID(uuid.New().String())
	cluster = &common.Cluster{
		Cluster: models.Cluster{
			ID: &clusterID,
		},
	}
	cluster.ImageInfo = &models.ImageInfo{}

	ctrl = gomock.NewController(GinkgoT())
	mockHostAPI = host.NewMockAPI(ctrl)
	manager = operators.NewManager(log)
})

var _ = AfterEach(func() {
	ctrl.Finish()
})

var _ = Describe("Operators manager", func() {

	It("should generate OCS and LSO manifests when OCS operator is enabled", func() {
		cluster.MonitoredOperators = []*models.MonitoredOperator{
			&ocs.Operator,
		}

		manifests, err := manager.GenerateManifests(cluster)

		Expect(err).NotTo(HaveOccurred())

		Expect(manifests).NotTo(BeNil())
		Expect(len(manifests)).To(Equal(10))
		Expect(manifests).To(HaveKey(ContainSubstring("openshift-ocs")))
		Expect(manifests).To(HaveKey(ContainSubstring("openshift-lso")))
	})

	It("should generate LSO manifests when LSO operator is enabled", func() {
		cluster.MonitoredOperators = []*models.MonitoredOperator{
			&lso.Operator,
		}

		manifests, err := manager.GenerateManifests(cluster)

		Expect(err).NotTo(HaveOccurred())

		Expect(manifests).NotTo(BeNil())
		Expect(len(manifests)).To(Equal(5))
		Expect(manifests).To(HaveKey(ContainSubstring("openshift-lso")))
	})

	It("should generate no manifests when no operator is present", func() {
		cluster.MonitoredOperators = []*models.MonitoredOperator{}
		manifests, err := manager.GenerateManifests(cluster)
		Expect(err).NotTo(HaveOccurred())

		Expect(manifests).NotTo(BeNil())
		Expect(manifests).To(BeEmpty())
	})

	table.DescribeTable("should report any operator enabled", func(operators []*models.MonitoredOperator, expected bool) {
		cluster.MonitoredOperators = operators

		results := manager.AnyOperatorEnabled(cluster)
		Expect(results).To(Equal(expected))
	},
		table.Entry("false for no operators", []*models.MonitoredOperator{}, false),
		table.Entry("true for lso operator", []*models.MonitoredOperator{
			&lso.Operator,
		}, true),
		table.Entry("true for ocs operator", []*models.MonitoredOperator{
			&ocs.Operator,
		}, true),
		table.Entry("true for lso and ocs operators", []*models.MonitoredOperator{
			&lso.Operator,
			&ocs.Operator,
		}, true),
	)

	It("should deem OCS operator valid when it's absent", func() {
		cluster.MonitoredOperators = []*models.MonitoredOperator{}

		valid := manager.ValidateOCSRequirements(cluster)

		Expect(valid).To(Equal("success"))
	})

	It("should deem OCS operator invalid when it's enabled and invalid", func() {
		cluster.MonitoredOperators = []*models.MonitoredOperator{
			&ocs.Operator,
		}

		valid := manager.ValidateOCSRequirements(cluster)

		Expect(valid).To(Equal("failure"))
	})
})
