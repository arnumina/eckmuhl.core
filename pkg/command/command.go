/*
#######
##                __               __   __
##       ___ ____/ /__ __ _  __ __/ /  / /    _______  _______
##      / -_) __/  '_//  ' \/ // / _ \/ / _  / __/ _ \/ __/ -_)
##      \__/\__/_/\_\/_/_/_/\_,_/_//_/_/ (_) \__/\___/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// ErrStopApp AFAIRE.
var ErrStopApp = errors.New("stop application requested")

type (
	// Command AFAIRE.
	Command interface {
		Name() string
		Description() string
		Version() string
		BuiltAt() time.Time
		Run(args []string) error
	}

	// CmdFlag AFAIRE.
	CmdFlag struct {
		*flag.FlagSet
		command Command
	}
)

// UnixToTime AFAIRE.
func UnixToTime(unix string) time.Time {
	ts, err := strconv.ParseInt(unix, 0, 64)
	if err != nil {
		ts = 0
	}

	return time.Unix(ts, 0).Local()
}

func (cf *CmdFlag) printUsage() {
	fmt.Println()
	fmt.Println(cf.Name(), "-", cf.command.Description())
	fmt.Println("================================================================================")
	fmt.Println("Options:")
	cf.PrintDefaults()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(cf.command.Version(), "/", cf.command.BuiltAt().String())
	fmt.Println("================================================================================")
	fmt.Println()
}

// NewCmdFlag AFAIRE.
func NewCmdFlag(cmd Command) *CmdFlag {
	cf := &CmdFlag{
		FlagSet: flag.NewFlagSet(cmd.Name(), flag.ContinueOnError),
		command: cmd,
	}

	cf.FlagSet.SetOutput(os.Stdout)
	cf.FlagSet.Usage = cf.printUsage

	return cf
}

// Parse AFAIRE.
func (cf *CmdFlag) Parse(args []string) error {
	if err := cf.FlagSet.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return ErrStopApp
		}

		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
