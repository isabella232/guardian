package guardiancmd

type CleanupCommand struct {
	*CommonCommand
}

func (cmd *CleanupCommand) Execute(args []string) error {
	log, _ := cmd.Logger.Logger("guardian-cleanup")
	cmd.Containers.DestroyContainersOnStartup = true

	wiring, err := cmd.createWiring(log)
	if err != nil {
		return err
	}

	gardener := cmd.createGardener(wiring)

	if err := gardener.Cleanup(log); err != nil {
		return err
	}

	cmd.saveProperties(log, cmd.Containers.PropertiesPath, wiring.PropertiesManager)

	return nil
}
