package update

func (o Operation) resetReleaseNext() error {
	o.Printf(">>> Reset %s branch to upstream/%s.\n",
		o.Config.Branches.ReleaseNext, o.Config.Branches.Main)
	return nil
}
