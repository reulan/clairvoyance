package extras

import (
	"math/rand"
	"time"
	//"fmt"
	//"github.com/kyokomi/emoji/v2"
)

/*
	Emoji
*/

var WinterEmoji = []string{":snowflake:", ":snowman:", ":Christmas:", "party popper"}
var FantasyEmoji = []string{":mage:", ":fairy:"}
var HalloweenEmoji = []string{
	":ghost:",
	":bone:",
	":vampire:",
	":zombie:",
	":spider:",
}

//":merperson:",
//":goblin:",
//":ogre:",
//":bomb:",
//":collision:",

var RobotEmoji = []string{
	":alien monster:",
	":robot:",
	":alien:",
}

var emojiTheme = [][]string{
	FantasyEmoji,
	HalloweenEmoji,
	RobotEmoji,
	WinterEmoji,
}

func emojiAmount() int {
	/*
		var emojiNumString string = os.Getenv("COLUMNS")
		var emojiNum int

		emojiNum, err := strconv.Atoi(emojiNumString)
		if err != nil {
			log.Println(err)
		} else {
			return 80
		}
		return 0
	*/
	return 20
}

// Pass in [][]string
// Get holiday festivity []string
// Find terminal width / character length
// for each num in witdhj choose a random emoji from theme
// return the string

func GetEmojiString() string {
	rand.Seed(time.Now().UnixNano())
	var randomTheme int = rand.Intn(2)
	var theme []string = emojiTheme[randomTheme]
	//log.Printf("randomTheme: %v\ntheme: %s", randomTheme, theme)

	var emojiString string
	var numEmoji int = 0

	for numEmoji < emojiAmount() {
		var emojiIndex = rand.Intn(len(theme))
		emojiString = emojiString + theme[emojiIndex]
		numEmoji = numEmoji + 1
	}
	return emojiString
}

var asciiArts = []string{Sleigh1Art, Sleigh2Art}

func GetAsciiArt() string {
	var index int = len(asciiArts)
	rand.Seed(time.Now().UnixNano())
	var randomArt int = rand.Intn(index)
	//log.Printf("randomArt: %v", randomArt)
	return asciiArts[randomArt]
}

/*
	ASCII art
*/

/*
	Art taken from: https://asciiart.website/index.php?art=holiday/christmas/santa

	sleigh1Art:
		From: u8211619@cc.nctu.edu.tw (Jurcy Hwang)
		Subject: [collection] my drawing
		Date: 10 Jan 1995 02:40:46 GMT
*/

var Sleigh1Art string = `
                                                       *
    *                                                          *
                                 *                  *        .--.
     \/ \/  \/  \/                                        ./   /=*
       \/     \/      *            *                ...  (_____)
        \ ^ ^/                                       \ \_((^o^))-.     *
        (o)(O)--)--------\.                           \   (   ) \  \._.
        |    |  ||================((~~~~~~~~~~~~~~~~~))|   ( )   |     \
         \__/             ,|        \. * * * * * * ./  (~~~~~~~~~~~)    \
  *        ||^||\.____./|| |          \___________/     ~||~~~~|~'\____/ *
           || ||     || || A            ||    ||          ||    |   jurcy
    *      <> <>     <> <>          (___||____||_____)   ((~~~~~|   *
`

var Sleigh2Art string = `
          __                                                      _.
 _---_.*~<('===          ,~~,         ,~~,         ,~~,           ,_)
(,    \ (__)=3--__._----_()'4__._----_()'4__._----_()'4__._,____.()'b
  \--------/-\  ~~(        ) ~~(        ) ~~(        )  ~~:       :'
   \_______|       (,_,,,_)     (,_,,,_)     (,_,,,_)     ;,,,,,,:
___I___I___I./     / /  \ \     / /  \ \     / /  \ \     / /  \ \
`
