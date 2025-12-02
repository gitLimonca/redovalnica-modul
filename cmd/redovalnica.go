package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/gitLimonca/redovalnica-modul/redovalnica"
)

func main() {

	// Začetni podatki
	studenti := map[string]redovalnica.Student{
		"63210001": {Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9, 8, 10, 9, 9}},
		"63210002": {Ime: "Boris", Priimek: "Kralj", Ocene: []int{6, 7, 5, 8, 6, 6}},
		"63210003": {Ime: "Janez", Priimek: "Novak", Ocene: []int{4, 5, 3, 5, 4, 4}},
	}

	app := &cli.App{
		Name:  "Redovalnica",
		Usage: "Aplikacija za upravljanje študentskih ocen",

		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "stOcen",
				Value:   6,
				Usage:   "Minimalno število ocen za izračun končnega uspeha",
				Aliases: []string{"s"},
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Value: 0,
				Usage: "Najmanjša dovoljena ocena",
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Value: 10,
				Usage: "Največja dovoljena ocena",
			},
		},

		Commands: []*cli.Command{
			{
				Name:  "izpis",
				Usage: "Izpiše vse študente in njihove ocene",
				Action: func(c *cli.Context) error {
					redovalnica.IzpisVsehOcen(studenti)
					return nil
				},
			},
			{
				Name:      "dodaj",
				Usage:     "Dodaj oceno izbranemu študentu",
				ArgsUsage: "<vpisna> <ocena>",
				Action: func(c *cli.Context) error {

					if c.Args().Len() < 2 {
						return fmt.Errorf("uporaba: dodaj <vpisna> <ocena>")
					}

					vpisna := c.Args().Get(0)
					ocena := c.Args().GetInt(1)

					err := redovalnica.DodajOceno(
						studenti,
						vpisna,
						ocena,
						c.Int("minOcena"),
						c.Int("maxOcena"),
					)

					if err != nil {
						return err
					}

					fmt.Println("Ocena uspešno dodana!")
					return nil
				},
			},
			{
				Name:  "uspeh",
				Usage: "Izpiše končni uspeh vseh študentov",
				Action: func(c *cli.Context) error {
					redovalnica.IzpisiKoncniUspeh(studenti, c.Int("stOcen"))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

