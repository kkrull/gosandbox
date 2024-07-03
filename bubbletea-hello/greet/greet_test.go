package greet_test

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kkrull/bubbletea-hello/greet"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("greet", func() {
	Describe("NewModel", func() {
		It("returns a model with a default message", func() {
			subject := greet.NewModel()
			Expect(subject.Message).To(Equal("Hello World!"))
		})
	})

	Describe("greeterModel", func() {
		Describe("#Init", func() {
			It("returns nil", func() {
				subject := greet.NewModel()
				Expect(subject.Init()).To(BeNil())
			})
		})

		Describe("#Update", func() {
			It("returns bubbletea.Quit, given a matching key press", func() {
				subject := greet.NewModel()

				message := tea.KeyMsg(tea.Key{Type: tea.KeyCtrlC})
				model, command := subject.Update(message)

				Expect(model).To(Equal(subject))
				// Expect(command).To(Equal(tea.Quit))
				Expect(command).NotTo(BeNil())
			})

			It("returns no change, given any other message", func() {
				subject := greet.NewModel()

				message := tea.KeyMsg(tea.Key{Type: tea.KeySpace})
				model, command := subject.Update(message)

				Expect(model).To(Equal(subject))
				Expect(command).To(BeNil())
			})
		})

		Describe("#View", func() {
			It("renders a greeting", func() {
				subject := greet.NewModel()
				Expect(subject.View()).To(ContainSubstring("Hello World!"))
			})
		})
	})
})
