/*
SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Gardener contributors

SPDX-License-Identifier: Apache-2.0
*/

package util_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/gardener/gardenctl-v2/internal/util"
)

var _ = Describe("ConfirmDialog utility", func() {
	var in *util.SafeBytesBuffer
	var out *util.SafeBytesBuffer
	var ioStreams util.IOStreams

	const question = "Is 42 the answer?"

	BeforeEach(func() {
		ioStreams, in, out, _ = util.NewTestIOStreams()
	})

	DescribeTable("Should return the default is user inputs nothing", func(userInput, choices string, defaultAnswer, expected bool) {
		_, err := in.Write([]byte(userInput))
		Expect(err).ShouldNot(HaveOccurred())

		result := util.ConfirmDialog(ioStreams, question, defaultAnswer)

		Expect(result).To(Equal(expected))
		Expect(out.String()).To(Equal(question + " " + choices + ": "))
	},
		Entry("Default true", "", "[n/Y]", true, true),
		Entry("Default false", "", "[y/N]", false, false),
	)

	DescribeTable("Should return the users input", func(userInput, choices string, defaultAnswer, expected bool) {
		_, err := in.Write([]byte(userInput))
		Expect(err).ShouldNot(HaveOccurred())

		result := util.ConfirmDialog(ioStreams, question, defaultAnswer)

		Expect(result).To(Equal(expected))
		Expect(out.String()).To(Equal(question + " " + choices + ": "))
	},
		Entry("user inputs 'y'", "y", "[n/Y]", true, true),
		Entry("user inputs 'n'", "n", "[n/Y]", true, false),
	)

	DescribeTable("Should print the choices correctly", func(userInput, choices string, defaultAnswer, expected bool) {
		_, err := in.Write([]byte(userInput))
		Expect(err).ShouldNot(HaveOccurred())

		result := util.ConfirmDialog(ioStreams, question, defaultAnswer)

		Expect(result).To(Equal(expected))
		Expect(out.String()).To(Equal(question + " " + choices + ": "))
	},
		Entry("with default answer true", "n", "[n/Y]", true, false),
		Entry("with default answer false", "n", "[y/N]", false, false),
	)

	It("Should repeat the question if user provides invalid input", func() {
		_, err := in.Write([]byte("invalid input\nn"))
		Expect(err).ShouldNot(HaveOccurred())

		result := util.ConfirmDialog(ioStreams, question, true)
		expectedPrompt := strings.Repeat(question+" [n/Y]: ", 2)

		Expect(result).To(Equal(false))
		Expect(out.String()).To(Equal(expectedPrompt))
	})
})
