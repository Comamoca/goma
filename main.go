package main

import (
	"log"
	"os"

	"github.com/Comamoca/chuno"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "goma",
		Usage: "Simple MarkDown renderer✨",
		Commands: []*cli.Command{
			{
				Name:    "preview",
				Aliases: []string{"p"},
				Usage:   "Start the preview server.",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "isdark",
						Value: false,
						Usage: "Enable dark mode",
					},
					&cli.IntFlag{
						Name:  "port",
						Value: 3535,
						Usage: "Preview server port number",
					},
				},
				Action: func(cCtx *cli.Context) error {
					isDark := cCtx.Bool("isdark")
					port := cCtx.Int("port")

					file := cCtx.Args().Get(0)

					color.Magenta("Starting the preview server...🚀")
					err := chuno.LaunchPreviewServer(file, port, isDark)
					if err != nil {
						log.Fatal("An error occurred while starting the preview server.", err)
						return err
					}

					return nil
				},
			},
			{
				Name:    "render",
				Aliases: []string{"r"},
				Usage:   "Convert MarkDown to HTML and output.",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "isdark",
						Value: false,
						Usage: "Enable dark mode",
					},
				},
				Action: func(cCtx *cli.Context) error {
					file := cCtx.Args().Get(0)
					isDark := cCtx.Bool("isdark")

					err := renderCmd(file, isDark)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
