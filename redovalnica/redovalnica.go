// Package redovalnica omogoča upravljanje študentskih ocen.
// Paket vsebuje funkcije za dodajanje ocen, izpis vseh ocen in izpis končnega uspeha študentov.
package redovalnica

import "fmt"

// Student predstavlja enega študenta z imenom, priimkom in ocenami.
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// DodajOceno doda oceno študentu znotraj določenega območja (minOcena, maxOcena).
// Vrne napako, če ocena ni veljavna ali študent ne obstaja.
func DodajOceno(studenti map[string]Student, vpisna string, ocena int, minOcena int, maxOcena int) error {
	if ocena < minOcena || ocena > maxOcena {
		return fmt.Errorf("ocena mora biti med %d in %d", minOcena, maxOcena)
	}
	student, obstaja := studenti[vpisna]
	if !obstaja {
		return fmt.Errorf("študent z vpisno številko %s ne obstaja", vpisna)
	}
	student.Ocene = append(student.Ocene, ocena)
	studenti[vpisna] = student
	return nil
}

// IzpisVsehOcen izpiše vse študente in njihove ocene.
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, s := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, s.Ime, s.Priimek, s.Ocene)
	}
}

// IzpisiKoncniUspeh izpiše končni uspeh vseh študentov glede na minimalno število ocen.
func IzpisiKoncniUspeh(studenti map[string]Student, stOcen int) {
	for _, s := range studenti {
		if len(s.Ocene) < stOcen {
			fmt.Printf("%s %s: premalo ocen za končni uspeh\n", s.Ime, s.Priimek)
			continue
		}
		p := povprecje(s)
		fmt.Printf("%s %s: povprečna ocena %.1f -> ", s.Ime, s.Priimek, p)
		switch {
		case p >= 9:
			fmt.Println("Odličen študent")
		case p >= 6:
			fmt.Println("Povprečen študent")
		default:
			fmt.Println("Neuspešen študent")
		}
	}
}

// povprecje izračuna povprečno oceno študenta.
// Funkcija ni izvožena in je namenjena internemu izračunu.
func povprecje(s Student) float64 {
	if len(s.Ocene) == 0 {
		return -1
	}
	sum := 0
	for _, o := range s.Ocene {
		sum += o
	}
	return float64(sum) / float64(len(s.Ocene))
}

// ExampleDodajOceno prikazuje primer uporabe funkcije DodajOceno.
func ExampleDodajOceno() {
	studenti := map[string]Student{
		"1": {Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9}},
	}
	_ = DodajOceno(studenti, "1", 8, 0, 10)
	IzpisVsehOcen(studenti)
	// Output:
	// 1 - Ana Novak: [10 9 8]
}

// ExampleIzpisiKoncniUspeh prikazuje primer uporabe funkcije IzpisiKoncniUspeh.
func ExampleIzpisiKoncniUspeh() {
	studenti := map[string]Student{
		"1": {Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9, 8}},
		"2": {Ime: "Boris", Priimek: "Kralj", Ocene: []int{5, 6}},
	}
	IzpisiKoncniUspeh(studenti, 3)
	// Output:
	// Ana Novak: povprečna ocena 9.0 -> Odličen študent
	// Boris Kralj: premalo ocen za končni uspeh
}
