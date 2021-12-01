package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var solvers = map[int]struct {
	P1, P2 func(string)
}{}

var day int

func init() {
	rootCmd.Flags().IntVarP(&day, "day", "d", 0, "day to run")
}

var rootCmd = &cobra.Command{
	Use: "aoc2021",
	Run: func(cmd *cobra.Command, args []string) {
		loggerConf := zap.NewDevelopmentConfig()
		loggerConf.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05.000")
		loggerConf.EncoderConfig.EncodeCaller = nil
		if isatty.IsTerminal(os.Stdout.Fd()) {
			loggerConf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}

		logger, err := loggerConf.Build()
		if err != nil {
			log.Fatal("could not set up logger", err)
		}

		zap.ReplaceGlobals(logger)

		if day == 0 {
			logger.Fatal("day must be set")
		}

		inputBytes, err := ioutil.ReadFile(fmt.Sprintf("./inputs/day%v.txt", day))
		if err != nil {
			logger.Fatal("failed to read input file", zap.Error(err))
		}

		solver, ok := solvers[day]
		if !ok {
			logger.Fatal("no solver registered for day", zap.Int("day", day))
		}

		if solver.P1 != nil {
			solver.P1(string(inputBytes))
		}

		if solver.P2 != nil {
			solver.P2(string(inputBytes))
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
