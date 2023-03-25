package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aphilas/aoc2022/day1"
	"github.com/aphilas/aoc2022/day2"
	"github.com/aphilas/aoc2022/day3"
	"github.com/aphilas/aoc2022/day4"
	"github.com/aphilas/aoc2022/day5"
	"github.com/aphilas/aoc2022/day6"
	"github.com/aphilas/aoc2022/day7"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cmd()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func solutions() {
	fmt.Println("day 1:")
	f1, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatalf("could not open input: %s", err)
	}
	defer f1.Close()

	fmt.Println(day1.CalCount(f1))
	f1.Seek(0, 0) // "Reset" cursor position
	fmt.Println(day1.CalRank(f1))
	fmt.Println()

	// Day 2

	fmt.Println("day 2:")
	f2, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatalf("could not open input: %s", err)
	}
	defer f2.Close()

	fmt.Println(day2.FixedChoice(f2))
	f2.Seek(0, 0) // "Reset" cursor position
	fmt.Println(day2.FixedEnding(f2))
	fmt.Println()

	// Day 3

	fmt.Println("day 3:")
	f3, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatalf("could not open input: %s", err)
	}
	defer f3.Close()

	fmt.Println(day3.CampDuplicates(f3))
	f3.Seek(0, 0)
	fmt.Println(day3.CampBadges(f3))
	fmt.Println()

	// Day 4

	fmt.Println("day 4:")
	f4, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatalf("could not open input: %s", err)
	}
	defer f4.Close()

	fmt.Println(day4.FullOverlaps(f4))
	f4.Seek(0, 0)
	fmt.Println(day4.PartialOverlaps(f4))
	fmt.Println()

	// Day 5

	fmt.Println("day 5:")
	f5, err := os.Open("./day5/input.txt")
	if err != nil {
		log.Fatalf("could not open input: %s", err)
	}
	defer f5.Close()

	fmt.Println(day5.MoveCratesSingly(f5, 9, 8))
	f5.Seek(0, 0)
	fmt.Println(day5.MoveCratesGrouped(f5, 9, 8))
	fmt.Println()

	// Day 6

	fmt.Println("day 6:")
	data, err := os.ReadFile("./day6/input.txt")
	if err != nil {
		log.Fatalf("could not open input: %s", err)
	}
	f6 := string(data)

	fmt.Println(day6.PacketMarker(f6))
	fmt.Println(day6.MessageMarker(f6))
	fmt.Println()

	// Day 7

	fmt.Println("day 7:")
	f7, err := os.Open("./day7/input.txt")
	if err != nil {
		log.Fatalf("could not open input: %s", err)
	}
	defer f7.Close()

	fmt.Println(day7.SpaceSaver(f7))
	// f7.Seek(0, 0)
	// fmt.Println(day7.MoveCratesGrouped(f7, 9, 8))
	// fmt.Println()
}

func cmd() *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "launch",
				Usage:   "create a new package",
				Aliases: []string{"l"},
				Action:  launch,
			},
			{
				Name:    "solutions",
				Usage:   "print solutions",
				Aliases: []string{"s"},
				Action: func(ctx *cli.Context) error {
					solutions()
					return nil
				},
			},
		},
	}
}

func launch(ctx *cli.Context) error {
	name := ctx.Args().Get(0)

	if name == "" {
		return fmt.Errorf("name must be a string")
	}

	err := os.Mkdir(name, 0750)
	if err != nil {
		return err
	}

	f1, err := os.Create(fmt.Sprintf("%v/%v.go", name, name))
	if err != nil {
		return err
	}
	defer f1.Close()

	f2, err := os.Create(fmt.Sprintf("%v/%v_test.go", name, name))
	if err != nil {
		return err
	}
	defer f2.Close()

	return nil
}
