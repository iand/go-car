package main

import (
	"log"
	"os"

	"github.com/multiformats/go-multicodec"
	"github.com/urfave/cli/v2"
)

func main() { os.Exit(main1()) }

func main1() int {
	app := &cli.App{
		Name:  "car",
		Usage: "Utility for working with car files",
		Commands: []*cli.Command{
			{
				Name:    "convert",
				Usage:   "Convert a car file to given codec",
				Aliases: []string{"con"},
				Action:  ConvertCar,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "selector",
						Aliases: []string{"s"},
						Usage:   "A selector over the dag",
					},
				},
			},
			{
				Name:    "concatinate",
				Usage:   "Concatinate car files",
				Aliases: []string{"cat"},
				Action:  CatCar,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:      "file",
						Aliases:   []string{"f", "output", "o"},
						Usage:     "The car file to write to",
						TakesFile: true,
					},
					&cli.IntFlag{
						Name:  "version",
						Value: 2,
						Usage: "Write output as a v1 or v2 format car",
					},
				},
			},
			{
				Name:    "create",
				Usage:   "Create a car file",
				Aliases: []string{"c"},
				Action:  CreateCar,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:      "file",
						Aliases:   []string{"f", "output", "o"},
						Usage:     "The car file to write to",
						TakesFile: true,
					},
					&cli.IntFlag{
						Name:  "version",
						Value: 2,
						Usage: "Write output as a v1 or v2 format car",
					},
				},
			},
			{
				Name:   "detach-index",
				Usage:  "Detach an index to a detached file",
				Action: DetachCar,
				Subcommands: []*cli.Command{{
					Name:   "list",
					Usage:  "List a detached index",
					Action: DetachCarList,
				}},
			},
			{
				Name:    "extract",
				Aliases: []string{"x"},
				Usage:   "Extract the contents of a car when the car encodes UnixFS data",
				Action:  ExtractCar,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:      "file",
						Aliases:   []string{"f"},
						Usage:     "The car file to extract from",
						Required:  true,
						TakesFile: true,
					},
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Usage:   "Include verbose information about extracted contents",
					},
				},
			},
			{
				Name:    "filter",
				Aliases: []string{"f"},
				Usage:   "Filter the CIDs in a car",
				Action:  FilterCar,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:      "cid-file",
						Usage:     "A file to read CIDs from",
						TakesFile: true,
					},
					&cli.BoolFlag{
						Name:  "append",
						Usage: "Append cids to an existing output file",
					},
				},
			},
			{
				Name:    "get-block",
				Aliases: []string{"gb"},
				Usage:   "Get a block out of a car",
				Action:  GetCarBlock,
			},
			{
				Name:    "get-dag",
				Aliases: []string{"gd"},
				Usage:   "Get a dag out of a car",
				Action:  GetCarDag,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "selector",
						Aliases: []string{"s"},
						Usage:   "A selector over the dag",
					},
					&cli.BoolFlag{
						Name:  "strict",
						Usage: "Fail if the selector finds links to blocks not in the original car",
					},
					&cli.IntFlag{
						Name:  "version",
						Value: 2,
						Usage: "Write output as a v1 or v2 format car",
					},
				},
			},
			{
				Name:   "import",
				Usage:  "Import a block into a car file",
				Action: ImportCar,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "codec",
						Aliases: []string{"c"},
						Usage:   "The codec the block data should be interpreted with",
						Value:   multicodec.DagJson.String(),
					},
				},
			},
			{
				Name:    "index",
				Aliases: []string{"i"},
				Usage:   "write out the car with an index",
				Action:  IndexCar,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "codec",
						Aliases: []string{"c"},
						Usage:   "The type of index to write",
						Value:   multicodec.CarMultihashIndexSorted.String(),
					},
					&cli.IntFlag{
						Name:  "version",
						Value: 2,
						Usage: "Write output as a v1 or v2 format car",
					},
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls"},
				Usage:   "List the CIDs in a car",
				Action:  ListCar,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Usage:   "Include verbose information about contained blocks",
					},
					&cli.BoolFlag{
						Name:  "unixfs",
						Usage: "List unixfs filesystem from the root of the car",
					},
				},
			},
			{
				Name:   "root",
				Usage:  "Get the root CID of a car",
				Action: CarRoot,
			},
			{
				Name:    "verify",
				Aliases: []string{"v"},
				Usage:   "Verify a CAR is wellformed",
				Action:  VerifyCar,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
		return 1
	}
	return 0
}
