package cli

type Module interface {
	AddContact(name, telephone string) error
	// ListContacts()
	// FindContact()
	// DeleteContact()
}

type Deps struct {
	Module Module
}

type CLI struct {
	Deps
	commandList []command
}

// NewCLI creates a command line interface
func NewCLI(d Deps) CLI {
	return CLI{
		Deps: d,
		commandList: []command {
			{
				name: 		 "help",
				description: "справка",
			},
			{
				name: 		 "addContact",
				description: "добавить контакт: использование add --name=SomeName --telephone=+7434543535",
			},
			{
				name: 		 "deleteContact",
				description: "удалить контакт: использование delete --name=SomeName",
			},
			{
				name: 		 "findContact",
				description: "найти контакт: использование find --name=SomeName",
			},
			{
				name: 		 "listContact",
				description: "список контакт: использование list",
			},
		}
	}
}

// Run ..
func (c CLI) Run() error {
	args := os.args[1:]
	if len(args) == 0 {
		return fmt.Errorf("command isn't set")
	}

	commandName := args[0]
	switch commandName {
	case help:
		c.help()
		return nil
	case addContact:
		return c.addContact(args[1:])
	case deleteContact:
		return c.deleteContact(args[1:])
	case listContact:
		return c.listContact()
	case findContact:
		return c.findContact(args[1:])
	}

	return fmt.Errorf("command isn't set")
}

func (c CLI) addContact(args []string) error {
	var name, telephone string

	fs := flag.NewFlagSet(addContact, flag.ContinueOnError)
	fs.StringVar(&name, "name", "", "use --name=SomeName")
	fs.StringVar(&telephone, "telephone", "", "use --telephone=+74549092356")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if len(name) == 0 {
		return errors.New("name is empty")
	}
	if len(telephone) == 0 {
		return errors.New("telephone is empty")
	}
	return c.Module.AddContact(name, telephone)
}

func (c CLI) listContact() error {
	// return c.Module.ListContact()
}

func (c CLI) deleteContact(args []string) error {
	return nil
}

func (c CLI) findContact(args []string) error {
	return nil
}

func (c CLI) help() {
	fmt.Println("command list:")
	for _, cmd := range c.commandList {
		fmt.Println("", cmd.name, cmd.description)
	}
	return
}