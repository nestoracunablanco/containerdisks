//nolint:lll
package tumbleweed

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"kubevirt.io/containerdisks/pkg/api"
	"kubevirt.io/containerdisks/pkg/common"
	"kubevirt.io/containerdisks/pkg/docs"
	"kubevirt.io/containerdisks/testutil"
)

var _ = Describe("OpenSUSE Tumbleweed", func() {
	DescribeTable("Inspect should be able to parse checksum files",
		func(arch, mockFile string, envVariables map[string]string, details *api.ArtifactDetails, metadata *api.Metadata) {
			c := New(arch, envVariables)
			c.getter = testutil.NewMockGetter(mockFile)
			got, err := c.Inspect()
			Expect(err).NotTo(HaveOccurred())
			Expect(got.ChecksumHash).ToNot(BeNil())
			Expect(got.Checksum).To(Equal(details.Checksum))
			Expect(got.DownloadURL).To(Equal(details.DownloadURL))
			Expect(got.AdditionalUniqueTags).To(Equal(details.AdditionalUniqueTags))
			Expect(got.ImageArchitecture).To(Equal(details.ImageArchitecture))
			Expect(got.Compression).To(Equal(details.Compression))
			Expect(c.Metadata()).To(Equal(metadata))
		},
		Entry("tumbleweed:1 x86_64", "x86_64", "testdata/tumbleweed.SHA256SUM",
			map[string]string{
				common.DefaultInstancetypeEnv: "u1.medium",
				common.DefaultPreferenceEnv:   "opensuse.tumbleweed",
			},
			&api.ArtifactDetails{
				Checksum:          "e8150b4a7ce5c56587492c930af094236c7a095149d714c015e6860ce6c58e66",
				DownloadURL:       "https://download.opensuse.org/tumbleweed/appliances/openSUSE-Tumbleweed-Minimal-VM.x86_64-1.0.0-Cloud-Snapshot20240629.qcow2",
				ImageArchitecture: "amd64",
			},
			&api.Metadata{
				Name:        "opensuse-tumbleweed",
				Version:     "1.0.0",
				Description: description,
				ExampleUserData: docs.UserData{
					Username: "opensuse",
				},
				EnvVariables: map[string]string{
					common.DefaultInstancetypeEnv: "u1.medium",
					common.DefaultPreferenceEnv:   "opensuse.tumbleweed",
				},
				Arch: "x86_64",
			},
		),
	)
})

func TestTumbleweed(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OpenSUSE Tumbleweed Suite")
}
