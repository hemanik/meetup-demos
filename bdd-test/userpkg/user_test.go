package userpkg_test

import (
	. "github.com/hemanik/meetup-demos/bdd-test/userpkg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var u *User

	// Given
	BeforeEach(func() {
		u = New()
	})

	Describe("Full Name", func() {
		// when
		Context("With a first and last name", func() {
			// then
			It("should concatenate the names", func() {
				u.FirstName = "Peyton"
				u.LastName = "Manning"
				Expect(u.FullName()).To(Equal("PeytonManning"))
			})
		})

		// when
		Context("With only a first name", func() {
			// then
			It("should return the first name", func() {
				u.FirstName = "Peyton"
				Expect(u.FullName()).To(Equal("Peyton"))
			})
		})

		// when
		Context("With only a last name", func() {
			// then
			It("should return the last name", func() {
				u.LastName = "Manning"
				Expect(u.FullName()).To(Equal("Manning"))
			})
		})

		// when
		Context("When first and last name are missing", func() {
			// then
			It("should return the empty string", func() {
				Expect(u.FullName()).To(BeEmpty())
			})
		})
	})

})
