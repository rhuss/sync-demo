package update

func (o Operation) addForkFiles() error {
	o.Printf(">>> Add fork's files on top of %s branch\n",
		o.Config.Branches.ReleaseNext)
	return nil
}
