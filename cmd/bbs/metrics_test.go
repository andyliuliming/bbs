package main_test

import (
	"code.cloudfoundry.org/bbs/cmd/bbs/testrunner"
	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/diego-logging-client/testhelpers"
	"code.cloudfoundry.org/locket"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/ginkgomon"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metrics", func() {
	BeforeEach(func() {
		bbsRunner = testrunner.New(bbsBinPath, bbsConfig)
	})

	JustBeforeEach(func() {
		bbsProcess = ginkgomon.Invoke(bbsRunner)
	})

	It("starts emitting metrics", func() {
		Eventually(testMetricsChan).Should(Receive())
	})

	It("starts emitting file descriptor count metrics", func() {
		Eventually(testMetricsChan).Should(Receive(
			testhelpers.MatchV2Metric(
				testhelpers.MetricAndValue{Name: "OpenFileDescriptors"},
			),
		))
	})

	Context("when the BBS instance isn't holding the lock", func() {
		var competingBBSLockProcess ifrit.Process

		BeforeEach(func() {
			competingBBSLock := locket.NewLock(logger, consulClient, locket.LockSchemaPath("bbs_lock"), []byte{}, clock.NewClock(), locket.RetryInterval, locket.DefaultSessionTTL, locket.WithMetronClient(&testhelpers.FakeIngressClient{}))
			competingBBSLockProcess = ifrit.Invoke(competingBBSLock)

			bbsRunner.StartCheck = "bbs.consul-lock.acquiring-lock"
		})

		AfterEach(func() {
			ginkgomon.Kill(competingBBSLockProcess)
		})

		It("still emits file descriptor count metrics", func() {
			Eventually(testMetricsChan).Should(Receive(
				testhelpers.MatchV2Metric(
					testhelpers.MetricAndValue{Name: "OpenFileDescriptors"},
				),
			))
		})
	})
})
