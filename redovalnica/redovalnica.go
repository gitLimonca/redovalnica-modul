package redovalnica

import (
	"fmt"
)

type Student struct {
	ime     string
	priimek string
	ocene   []int
}


func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena < 0 || ocena > 10 {
		fmt.Println("POZOR: Ocena mora biti med 0 in 10.")
		return
	}

	student, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		fmt.Println("Študent z vpisno številko ", vpisnaStevilka, " ne obstaja.")
		return
	}

	student.ocene = append(student.ocene, ocena)
	studenti[vpisnaStevilka] = student
	fmt.Println("Študent z vpisno številko ", vpisnaStevilka, " dobil oceno ", ocena, ".")
}


func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		return -1.0
	}

	if len(student.ocene) < 6 {
		return 0.0
	}

	sum := 0
	for _, o := range student.ocene {
		sum += o
	}
	return float64(sum) / float64(len(student.ocene))
}


func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, student := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, student.ime, student.priimek, student.ocene)
	}
}


func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisna, student := range studenti {
		p := povprecje(studenti, vpisna)

		if p == -1.0 {
			fmt.Printf("%s %s: ni podatkov\n", student.ime, student.priimek)
			continue
		}

		fmt.Printf("%s %s: povprečna ocena %.1f -> ", student.ime, student.priimek, p)
		if p >= 9 {
			fmt.Println("Odličen študent!")
		} else if p >= 6 {
			fmt.Println("Povprečen študent")
		} else {
			fmt.Println("Neuspešen študent")
		}
	}
}



