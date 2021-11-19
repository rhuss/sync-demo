package update

import "github.com/cardil/deviate/pkg/log/color"

func (o Operation) triggerCI() error {
	o.Println("Trigger CI",
		color.Yellow("Not yet implemented"))
	return nil
}
