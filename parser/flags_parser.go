package parser

import (
	"flag"
	"fmt"
	"neuro-csv/utils"
)

type NeuroCsvFlags struct {
	Config, Input  string
	GenerateConfig bool
}

const configDescription = "Path to config"
const inputDescription = "CSV file to analyze"
const generateConfigDescription = "Generate config file"

const usageTemplate = `Usage of %s:
  -c,  --config              Path to config
 -gc,  --generate-config     Generate config file
  -i,  --input               CSV file to analyze
`

func ParseFlags() *NeuroCsvFlags {
	flags := new(NeuroCsvFlags)

	flag.StringVar(&flags.Config, "c", "", configDescription)
	flag.StringVar(&flags.Config, "config", "", configDescription)

	flag.StringVar(&flags.Input, "i", "", inputDescription)
	flag.StringVar(&flags.Input, "input", "", inputDescription)

	flag.BoolVar(&flags.GenerateConfig, "gc", false, generateConfigDescription)
	flag.BoolVar(&flags.GenerateConfig, "generate-config", false, generateConfigDescription)

	flag.Usage = func() {
		fmt.Println(
			fmt.Sprintf(usageTemplate, utils.GetExecutableName()))
	}

	flag.Parse()

	return flags
}
