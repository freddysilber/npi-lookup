package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

/*
0 - Profession Code
1 - Profession Name
2 - File Number (Formerly known as Lic_id)
3 - Current License Expiration Date
4 - Original Date (Original Issuance of License)
5 - Profession Rank Code
6 - License Number
7 - Effective Date
8 - Board Action Indicator
9 - License Status Description
10 - Last Name
11 - First Name
12 - Middle Name Initial
13 - Name Suffix
14 - Business Name
15 - License Active Status Description
16 - County
17 - County Description
18 - Mailing Address Line1
19 - Mailing Address Line2
20 - Mailing Address City
21 - Mailing Address State
22 - Mailing Address ZIP Code
23 - Mailing Address Area Code
24 - Mailing Address Phone Number
25 - Mailing Address Phone Extension
26 - Practice Location Address Line1
27 - Practice Location Address Line2
28 - Practice Location Address City
29 - Practice Location Address State
30 - Practice Location Address ZIP Code
31 - Email Address
32 - Mod Code
33 - Prescribing Indicator
34 - Dispensing Indicator
*/

type Practitioner struct {
	firstName      string
	lastName       string
	email          string
	licenseNumber  string
	professionName string
}

func main() {
	readFile, err := os.Open("./input/input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var counties = []string{"palm beach", "martin", "saint lucie", "indian river", "okeechobee"}
	var parts []string
	var state string
	var county string
	var firstName string
	var lastName string
	var email string
	var licenseNumber string
	var professionName string

	for fileScanner.Scan() {
		parts = strings.Split(fileScanner.Text(), "|")

		state = strings.ToLower(parts[29])
		county = strings.ToLower(parts[17])
		firstName = parts[10]
		lastName = parts[11]
		email = parts[31]
		licenseNumber = parts[6]
		professionName = parts[1]

		if state == "fl" && slices.Contains(counties, county) {
			practitioner := Practitioner{
				firstName,
				lastName,
				email,
				licenseNumber,
				professionName,
			}
			fmt.Print(practitioner)
			fmt.Print("\n \n")
		}
	}

	readFile.Close()
}
