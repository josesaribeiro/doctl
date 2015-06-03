package main

import (
	"github.com/bryanl/docli/sshkeys"
	"github.com/codegangsta/cli"
)

func sshKeyCommands() cli.Command {
	return cli.Command{
		Name:  "ssh-key",
		Usage: "ssh key commands",
		Subcommands: []cli.Command{
			sshKeyList(),
			sshKeyCreate(),
			sshKeyGet(),
			sshKeyUpdate(),
			sshKeyDelete(),
		},
	}
}

func sshKeyList() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "list ssh keys",
		Action: sshkeys.List,
	}
}

func sshKeyCreate() cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "create ssh key",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name",
				Usage: "ssh key name",
			},
			cli.StringFlag{
				Name:  "public-key",
				Usage: "ssh public key",
			},
		},
		Action: sshkeys.Create,
	}
}

func sshKeyGet() cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "get ssh key",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "key",
				Usage: "ssh key id or fingerprint",
			},
		},
		Action: sshkeys.Get,
	}
}

func sshKeyUpdate() cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "update ssh key",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "ssh key id",
			},
			cli.StringFlag{
				Name:  "fingerprint",
				Usage: "ssh key fingerprint",
			},
			cli.StringFlag{
				Name:  "name",
				Usage: "ssh key name",
			},
		},
		Action: sshkeys.Update,
	}
}

func sshKeyDelete() cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "delete ssh key",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "key",
				Usage: "ssh key id or fingerprint",
			},
		},
		Action: sshkeys.Delete,
	}
}