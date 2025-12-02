# Redovalnica Modul

Aplikacija za upravljanje študentskih ocen v jeziku Go.  
Omogoča dodajanje ocen, izpis vseh ocen in izpis končnega uspeha študentov.

---
# Uporaba

Ukaz: go run ./cmd <command>

3-je commandi: 
"izpis" ---> izpiše vse študente in njihove ocene

"dodaj <vpisna> <ocena>" ---> doda oceno <ocena> študentu z vpisno številko <vpisna>
  možna 2 flaga: "--minOcena" ---> najmanjša možna ocena (privzeto 0)
                 "--maxOcena" ---> največja možna ocena (privzeto 10)

"uspeh" ---> izpiše končni uspeh vseh študentov glede na minimalno število ocen
  možni flag: "--stOcen" ---> minimalno število ocen za izračun končnega uspeha (privzeto 6)
