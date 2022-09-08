package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Microsoft/hcsshim/ext4/tar2ext4"
	"github.com/Microsoft/hcsshim/internal/log"
	"github.com/urfave/cli"
)

var (
	input      = flag.String("i", "", "input file(s), comma separated")
	output     = flag.String("o", "", "output file")
	overlay    = flag.Bool("overlay", false, "produce overlayfs-compatible layer image")
	vhd        = flag.Bool("vhd", false, "add a VHD footer to the end of the image")
	inlineData = flag.Bool("inline", false, "write small file data into the inode; not compatible with DAX")

	read = flag.Bool("read", false, "if true, read from input file only")
)

/*func main() {
	flag.Parse()
	if flag.NArg() != 0 || len(*output) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	err := func() (err error) {
		inputs := []io.Reader{os.Stdin}
		// in := os.Stdin // TODO katiewasnothere: fix this later
		if *input != "" {
			paths := strings.Split(*input, ",")
			inputs = []io.Reader{}
			for _, p := range paths {
				in, err := os.Open(p)
				if err != nil {
					return err
				}
				inputs = append(inputs, in)
			}

		}
		out, err := os.Create(*output)
		if err != nil {
			return err
		}

		var opts []tar2ext4.Option
		if *overlay {
			opts = append(opts, tar2ext4.ConvertWhiteout)
		}
		if *vhd {
			opts = append(opts, tar2ext4.AppendVhdFooter)
		}
		if *inlineData {
			opts = append(opts, tar2ext4.InlineData)
		}
		if len(inputs) == 1 {
			err = tar2ext4.Convert(inputs[0], out, opts...)
			if err != nil {
				return err
			}
		} else {
			err = tar2ext4.ConvertMultiple(inputs, out, opts...)
			if err != nil {
				return err
			}
		}

		// Exhaust the tar stream.
		// TODO katiewasnothere: do we need this?
		// _, _ = io.Copy(ioutil.Discard, in)
		return nil
	}()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}*/

func main() {
	app := cli.NewApp()
	app.Name = "tar2ext4"
	app.Usage = "converts tar file(s) into vhd(s)"
	app.Commands = []cli.Command{
		readGPTCommand,
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output,o",
			Usage: "output file",
		},
		cli.BoolFlag{
			Name:  "overlay",
			Usage: "produce overlayfs-compatible layer image",
		},
		cli.BoolFlag{
			Name:  "vhd",
			Usage: "add a VHD footer to the end of the image",
		},
		cli.BoolFlag{
			Name:  "inline",
			Usage: "write small file data into the inode; not compatible with DAX",
		},
		cli.StringSliceFlag{
			Name:  "input,i",
			Usage: "input file(s)",
		},
	}
	app.Action = func(cliCtx *cli.Context) (err error) {
		defer func() {
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}()
		fileNames := cliCtx.StringSlice("input")
		inputs := []io.Reader{}
		if len(fileNames) == 0 {
			inputs = append(inputs, os.Stdin)
		}

		for _, p := range fileNames {
			in, err := os.Open(p)
			if err != nil {
				return err
			}
			inputs = append(inputs, in)
		}
		outputName := cliCtx.String("output")
		out, err := os.Create(outputName)
		if err != nil {
			return err
		}

		var opts []tar2ext4.Option
		if cliCtx.Bool("overlay") {
			opts = append(opts, tar2ext4.ConvertWhiteout)
		}
		if cliCtx.Bool("vhd") {
			opts = append(opts, tar2ext4.AppendVhdFooter)
		}
		if cliCtx.Bool("inline") {
			opts = append(opts, tar2ext4.InlineData)
		}
		if len(inputs) == 1 {
			err = tar2ext4.Convert(inputs[0], out, opts...)
			if err != nil {
				return err
			}
		} else {
			err = tar2ext4.ConvertMultiple(inputs, out, opts...)
			if err != nil {
				return err
			}
		}

		// Exhaust the tar stream.
		// TODO katiewasnothere: do we need this?
		// _, _ = io.Copy(ioutil.Discard, in)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var readGPTCommand = cli.Command{
	Name:  "read-gpt",
	Usage: "read various data structures from a gpt disk",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "input",
			Usage: "input file to read",
		},
		cli.BoolFlag{
			Name:  "pmbr",
			Usage: "read pmbr data structures",
		},
		cli.BoolFlag{
			Name:  "header",
			Usage: "read gpt header",
		},
		cli.BoolFlag{
			Name:  "alt-header",
			Usage: "read alt gpt header",
		},
		cli.BoolFlag{
			Name:  "entry",
			Usage: "read the main entry array",
		},
		cli.BoolFlag{
			Name:  "alt-entry",
			Usage: "read the alt entry array",
		},
		cli.BoolFlag{
			Name:  "all",
			Usage: "read all bytes",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		input := cliCtx.String("input")
		if input == "" {
			return fmt.Errorf("input must be specified ")
		}
		inFile, err := os.Open(input)
		if err != nil {
			return err
		}
		if cliCtx.Bool("pmbr") {
			pmbr, err := tar2ext4.ReadPMBR(inFile)
			if err != nil {
				return err
			}
			log.G(context.Background()).WithField("pmbr", pmbr).Info("file pmbr")
		}
		if cliCtx.Bool("header") {
			header, err := tar2ext4.ReadGPTHeader(inFile, 1)
			if err != nil {
				return err
			}
			log.G(context.Background()).WithField("header", header).Info("file header")
		}
		if cliCtx.Bool("entry") {
			header, err := tar2ext4.ReadGPTHeader(inFile, 1)
			if err != nil {
				return err
			}
			entry, err := tar2ext4.ReadGPTPartitionArray(inFile, header.PartitionEntryLBA, header.NumberOfPartitionEntries)
			if err != nil {
				return err
			}
			log.G(context.Background()).WithField("entry", entry).Info("file entry")

		}
		if cliCtx.Bool("alt-entry") {
			header, err := tar2ext4.ReadGPTHeader(inFile, 1)
			if err != nil {
				return err
			}
			altHeader, err := tar2ext4.ReadGPTHeader(inFile, header.AlternateLBA)
			if err != nil {
				return err
			}
			entry, err := tar2ext4.ReadGPTPartitionArray(inFile, altHeader.PartitionEntryLBA, altHeader.NumberOfPartitionEntries)
			if err != nil {
				return err
			}
			log.G(context.Background()).WithField("alt-entry", entry).Info("file alt-entry")

		}
		if cliCtx.Bool("alt-header") {
			header, err := tar2ext4.ReadGPTHeader(inFile, 1)
			if err != nil {
				return err
			}
			altHeader, err := tar2ext4.ReadGPTHeader(inFile, header.AlternateLBA)
			if err != nil {
				return err
			}
			log.G(context.Background()).WithField("altHeader", altHeader).Info("file altHeader")

		}
		if cliCtx.Bool("all") {
			all, err := io.ReadAll(inFile)
			if err != nil {
				return err
			}
			resultAsString := ""
			for _, b := range all {
				resultAsString += fmt.Sprintf(" %x", b)
			}
			log.G(context.Background()).WithField("byte", resultAsString).Infof("file all content")

		}
		return nil
	},
}
