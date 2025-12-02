package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gitLimonca/redovalnica-modul/redovalnica"
	"github.com/urfave/cli/v3"
)

// Glavni program aplikacije Redovalnica.
// Podpira naslednje ukaze:
//   - izpis   : izpiše vse študente in njihove ocene
//   - dodaj   : doda oceno izbranemu študentu (ukaz: dodaj <vpisna> <ocena>)
//   - uspeh  : izpiše končni uspeh vseh študentov glede na minimalno število ocen
func main() {
	// Začetni podatki o študentih
	studenti := map[string]redovalnica.Student{
		"63210001": {Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9, 8, 10, 9, 9}},
		"63210002": {Ime: "Boris", Priimek: "Kralj", Ocene: []int{6, 7, 5, 8, 6, 6}},
		"63210003": {Ime: "Janez", Priimek: "Novak", Ocene: []int{4, 5, 3, 5, 4, 4}},
	}

	// Definicija ukazov in flagov
	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Aplikacija za upravljanje študentskih ocen",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Value: 6,
				Usage: "Minimalno število ocen za izračun končnega uspeha",
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
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Preverimo, ali je podan ukaz
			if cmd.Args().Len() < 1 {
				fmt.Println("Uporaba: izpis | dodaj <vpisna> <ocena> | uspeh")
				return nil
			}

			ukaz := cmd.Args().Get(0)

			switch ukaz {
			case "izpis":
				// Izpiše vse študente in njihove ocene
				redovalnica.IzpisVsehOcen(studenti)
			case "dodaj":
				// Dodaj oceno študentu
				if cmd.Args().Len() < 3 {
					fmt.Println("Uporaba: dodaj <vpisna> <ocena>")
					return nil
				}
				vpisna := cmd.Args().Get(1)
				ocena, _ := strconv.Atoi(cmd.Args().Get(2))
				minOcena := cmd.Int("minOcena")
				maxOcena := cmd.Int("maxOcena")
				if err := redovalnica.DodajOceno(studenti, vpisna, ocena, minOcena, maxOcena); err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Dodali smo oceno %d študentu %s\n", ocena, vpisna)
				}
			case "uspeh":
				// Izpiše končni uspeh vseh študentov glede na minimalno število ocen
				stOcen := cmd.Int("stOcen")
				redovalnica.IzpisiKoncniUspeh(studenti, stOcen)
			default:
				fmt.Println("Neznan ukaz. Na voljo so: izpis, dodaj, uspeh")
			}

			return nil
		},
	}

	// Zaženemo ukaz
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
