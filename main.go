// A quick and super dirty wordlist generator / mangler that takes a list of keywords and appends them with years, songs, seasons, nfl, nba, and nhl teams.
// [!] CAUTION: This will generate a large wordlist. Use at your own risk.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Variables
	outdir := ""
	list := ""
	leet := "no"

	flag.StringVar(&outdir, "o", "output", "Enter the directory to save wordlist.")
	flag.StringVar(&list, "i", "import", "Import list of use to generate list.")
	flag.StringVar(&leet, "l", "leet", "Add leet chars to wordlist (Default: No)")
	flag.Parse()

	if outdir != "" && list != "" {
		mangeled := ImportWords(list)
		mangleres := FormatWords(mangeled, leet)
		Finalmanagled := YearAdd(mangleres)
		ExportList(Finalmanagled)

	} else {
		fmt.Println("You must supply both a output directory and a list to import.")
	}
}

func ImportWords(list string) []string {
	var outwords []string

	//Opens user provided file
	file, err := os.Open(list)
	if err != nil {
		log.Fatal(err)
	}

	//Invoke Scanner and read in file
	scanner := bufio.NewScanner(file)
	// Split words on file
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	for _, line := range text {
		outwords = append(outwords, line)
	}

	return outwords
}

func FormatWords(mangled []string, leet string) (result []string) {

	// Create a string array of top songs of the year
	songs := []string{"Shape of You",
		"Despacito",
		"That's What I Like",
		"Humble",
		"Something Just Like This",
		"Bad and Boujee",
		"Closer",
		"Body Like a Back Road",
		"Believer",
		"Congratulations",
		"Mask Off",
		"XO TOUR Llif3",
		"Rockstar",
		"Unforgettable",
		"Havana",
		"Perfect",
		"Look What You Made Me Do",
		"Black Beatles",
		"Starboy",
		"1-800-273-8255",
		"Rockstar",
		"Attention",
		"Bad at Love",
		"Rake It Up",
		"I'm the One",
		"Wild Thoughts",
		"Bodak Yellow",
		"Bank Account",
		"Meant to Be",
		"Thunder",
		"Feel It Still",
		"Sorry Not Sorry",
		"Mi Gente",
		"Too Good at Goodbyes",
		"Strip That Down",
		"Slow Hands",
		"Juju on That Beat",
		"Look at Me!",
		"Let Me Love You",
		"X",
		"New Rules",
		"24K Magic",
		"Location",
		"Love on the Brain",
		"Praying",
		"Despacito",
		"Say You Won't Let Go",
		"Scars to Your Beautiful",
		"24K Magic",
		"Starving",
		"Paris",
		"Love Galore",
		"Castle on the Hill",
		"Fake Love",
		"iSpy",
		"Swang",
		"Side to Side",
		"Passionfruit",
		"Caroline",
		"Swalla",
		"Sign of the Times",
		"Bad Things",
		"Chained to the Rhythm",
		"Rolex",
		"DNA",
		"Slide",
		"1-800-273-8255",
		"Play That Song",
		"Redbone",
		"Starving",
		"Everyday We Lit",
		"Mercy",
	}

	// Add season, nfl, nba, and nhl teams to string array
	seasons := []string{"Summer", "Fall", "Spring", "Winter", "Autumn"}
	nfl_teams := []string{"Cardinals",
		"Falcons",
		"Ravens",
		"Bills",
		"Panthers",
		"Bears",
		"Bengals",
		"Browns",
		"Cowboys",
		"Broncos",
		"Lions",
		"Packers",
		"Texans",
		"Colts",
		"Jaguars",
		"Chiefs",
		"Chargers",
		"Rams",
		"Dolphins",
		"Vikings",
		"Patriots",
		"Saints",
		"Giants",
		"Jets",
		"Raiders",
		"Eagles",
		"Steelers",
		"49ers",
		"Seahawks",
		"Buccaneers",
		"Titans",
		"Redskins",
	}

	nhl_teams := []string{
		"Ducks",
		"Bruins",
		"Sabres",
		"Flames",
		"Hurricanes",
		"Blackhawks",
		"Avalanche",
		"BLUEJackets",
		"Stars",
		"RedWings",
		"Oilers",
		"Panthers",
		"Kings",
		"Wild",
		"Canadiens",
		"Predators",
		"Devils",
		"Islanders",
		"Rangers",
		"Senators",
		"Flyers",
		"Coyotes",
		"Penguins",
		"BLUEs",
		"Sharks",
		"Lighting",
		"Leafs",
		"Canucks",
		"GoldenKnights",
		"Capitals",
		"Jets",
	}

	nba_teams := []string{"Hawks",
		"Celtics",
		"Nets",
		"Hornets",
		"Bulls",
		"Cavaliers",
		"Mavericks",
		"Nuggets",
		"Pistons",
		"Warriors",
		"Rockets",
		"Pacers",
		"Clippers",
		"Lakers",
		"Grizzlies",
		"Heat",
		"Bucks",
		"Timberwolves",
		"Pelicans",
		"Knicks",
		"Thunder",
		"Magic",
		"76ers",
		"Suns",
		"Blazers",
		"Kings",
		"Spurs",
		"Raptors",
		"Jazz",
		"Wizards"}

	switch leet {
	case "yes":
		println("[*] Running with Leet Mode!." + "\n")

		// Add leet seasons
		for _, line := range mangled {
			for _, sline := range seasons {
				out := line + sline
				result = append(result, out)
				leet := strings.Replace(line, "a", "@", 3) + sline
				result = append(result, strings.ToUpper(leet))
				result = append(result, strings.ToLower(leet))
				result = append(result, strings.Title(leet))
				out2 := sline + line
				result = append(result, out2)
				leet2 := strings.Replace(sline, "a", "@", 3) + line
				result = append(result, strings.ToUpper(leet2))
				result = append(result, strings.ToLower(leet2))
				result = append(result, strings.Title(leet2))

			}
		}

		// Add leet NFL teams
		for _, line := range mangled {
			for _, nfl_line := range nfl_teams {
				out := line + nfl_line
				result = append(result, out)
				leet := strings.Replace(line, "a", "@", 3) + nfl_line
				result = append(result, strings.ToUpper(leet))
				result = append(result, strings.ToLower(leet))
				out2 := nfl_line + line
				result = append(result, out2)
				leet2 := strings.Replace(nfl_line, "a", "@", 3) + line
				result = append(result, strings.ToUpper(leet2))
				result = append(result, strings.ToLower(leet2))
				result = append(result, strings.Title(leet2))
			}
		}

		// Add leet NBA teams
		for _, line := range mangled {
			for _, nba_line := range nba_teams {
				out := line + nba_line
				result = append(result, out)
				leet := strings.Replace(line, "a", "@", 3) + nba_line
				result = append(result, strings.ToUpper(leet))
				result = append(result, strings.ToLower(leet))
				out2 := nba_line + line
				result = append(result, out2)
				leet2 := strings.Replace(nba_line, "a", "@", 3) + line
				result = append(result, strings.ToUpper(leet2))
				result = append(result, strings.ToLower(leet2))
				result = append(result, strings.Title(leet2))
			}
		}

		// Add leet NHL teams
		for _, line := range mangled {
			for _, nhl_line := range nhl_teams {
				out := line + nhl_line
				result = append(result, out)
				leet := strings.Replace(line, "a", "@", 3) + nhl_line
				result = append(result, strings.ToUpper(leet))
				result = append(result, strings.ToLower(leet))
				out2 := nhl_line + line
				result = append(result, out2)
				leet2 := strings.Replace(nhl_line, "a", "@", 3) + line
				result = append(result, strings.ToUpper(leet2))
				result = append(result, strings.ToLower(leet2))
				result = append(result, strings.Title(leet2))
			}
		}

		// Add leet songs
		for _, line := range mangled {
			for _, song := range songs {
				out := line + song
				out = strings.Replace(out, " ", "_", 3)
				result = append(result, out)
				leet := strings.Replace(line, "a", "@", 3) + song
				leet = strings.Replace(leet, " ", "_", 3)
				result = append(result, strings.ToUpper(leet))
				result = append(result, strings.ToLower(leet))
				out2 := song + line
				out2 = strings.Replace(out2, " ", "_", 3)
				result = append(result, out2)
				leet2 := song + strings.Replace(line, "a", "@", 3)
				leet2 = strings.Replace(leet2, " ", "_", 3)
				result = append(result, strings.ToUpper(leet2))
				result = append(result, strings.ToLower(leet2))
				result = append(result, strings.Title(leet2))
			}
		}

		// change characters for words in the mangled array
		for _, line := range result {
			replace := strings.Replace(line, "a", "@", 3)
			result = append(result, replace)
			replace2 := strings.Replace(line, "i", "!", 1)
			result = append(result, replace2)
		}

	default:
		fmt.Println("[!] Running without Leet Mode.")

		// Add leet seasons
		for _, line := range mangled {
			for _, sline := range seasons {
				out := line + sline
				result = append(result, strings.ToUpper(out))
				result = append(result, strings.ToLower(out))
				result = append(result, strings.Title(out))
				out2 := sline + line
				result = append(result, strings.ToUpper(out2))
				result = append(result, strings.ToLower(out2))
				result = append(result, strings.Title(out2))

			}
		}

		// Add NFL teams
		for _, line := range mangled {
			for _, nfl_line := range nfl_teams {
				out := line + nfl_line
				result = append(result, strings.ToUpper(out))
				result = append(result, strings.ToLower(out))
				result = append(result, strings.Title(out))
				out2 := nfl_line + line
				result = append(result, strings.ToUpper(out2))
				result = append(result, strings.ToLower(out2))
				result = append(result, strings.Title(out2))
			}
		}

		// Add NBA teams
		for _, line := range mangled {
			for _, nba_line := range nba_teams {
				out := line + nba_line
				result = append(result, strings.ToUpper(out))
				result = append(result, strings.ToLower(out))
				result = append(result, strings.Title(out))
				out2 := nba_line + line
				result = append(result, strings.ToUpper(out2))
				result = append(result, strings.ToLower(out2))
				result = append(result, strings.Title(out2))
			}
		}

		// Add NHL teams and year to each line in the appened list
		for _, line := range mangled {
			for _, nhl_line := range nhl_teams {
				out := line + nhl_line
				result = append(result, strings.ToUpper(out))
				result = append(result, strings.ToLower(out))
				result = append(result, strings.Title(out))
				out2 := nhl_line + line
				result = append(result, strings.ToUpper(out2))
				result = append(result, strings.ToLower(out2))
				result = append(result, strings.Title(out2))
			}
		}

		// Add songs to each line in the appened list
		for _, line := range mangled {
			for _, song := range songs {
				out := line + song
				out = strings.Replace(out, " ", "_", 3)
				result = append(result, strings.ToUpper(out))
				result = append(result, strings.ToLower(out))
				result = append(result, strings.Title(out))
				out2 := song + line
				out2 = strings.Replace(out2, " ", "_", 3)
				result = append(result, strings.ToUpper(out2))
				result = append(result, strings.ToLower(out2))
				result = append(result, strings.Title(out2))
			}
		}
	}

	return result

}

func YearAdd(managledres []string) (final []string) {
	// Adds years to each of the lines in the mangled array
	// Add years into string array for the next 5 years.
	dt := time.Now()
	date := dt.String()
	year := strings.Split(date, "-")[0]
	numyear, _ := strconv.Atoi(year)
	endyear := numyear + 5 // change years here if needed.

	var years []string
	for numyear <= endyear {
		years = append(years, strconv.Itoa(numyear))
		numyear++
	}

	// copy current line in managledres and add year to it.
	for _, line := range managledres {
		for _, year := range years {
			out := line + year
			original := &line
			final = append(final, *original)
			final = append(final, out)

		}
	}

	// Removing duplicates
	words_string := map[string]bool{}
	for i := range final {
		words_string[final[i]] = true
	}
	output := []string{}
	for j, _ := range words_string {
		output = append(output, j)
	}

	return output
}

// Export list to file
func ExportList(Finalmanagled []string) {

	file, err := os.Create("wordlist.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for _, line := range Finalmanagled {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
