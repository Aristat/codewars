/*

Introduction
There is a war and nobody knows - the alphabet war!
There are two groups of hostile letters. The tension between left side letters and right side letters was too high and the war began. The letters have discovered a new unit - a priest with Wo lo looooooo power.

Task
Write a function that accepts fight string consists of only small letters and return who wins the fight. When the left side wins return Left side wins!, when the right side wins return Right side wins!, in other case return Let's fight again!.

The left side letters and their power:

 w - 4
 p - 3
 b - 2
 s - 1
 t - 0 (but it's priest with Wo lo loooooooo power)
The right side letters and their power:

 m - 4
 q - 3
 d - 2
 z - 1
 j - 0 (but it's priest with Wo lo loooooooo power)
The other letters don't have power and are only victims.
The priest units t and j change the adjacent letters from hostile letters to friendly letters with the same power.

mtq => wtp
wjs => mjz
A letter with adjacent letters j and t is not converted i.e.:

tmj => tmj
jzt => jzt
The priests (j and t) do not convert the other priests ( jt => jt ).

*/

package kata

import (
	"strings"
)

func applyEffect(characters []string, charIndex int, charactersMapping map[string]string, priestCharacter string, enemyMap map[string]int) {
	prev := charIndex - 1
	beforePrev := charIndex - 2
	next := charIndex + 1
	afterNext := charIndex + 2

	if prev >= 0 && characters[prev] != priestCharacter {
		skip := false
		if beforePrev >= 0 && characters[beforePrev] == priestCharacter {
			skip = true
		}

		if _, ok := enemyMap[characters[prev]]; ok && !skip {
			characters[prev] = charactersMapping[characters[prev]]
		}
	}

	if next < len(characters) && characters[next] != priestCharacter {
		skip := false
		if afterNext < len(characters) && characters[afterNext] == priestCharacter {
			skip = true
		}

		if _, ok := enemyMap[characters[next]]; ok && !skip {
			characters[next] = charactersMapping[characters[next]]
		}
	}
}

func AlphabetWar(fight string) string {
	left := map[string]int{"w": 4, "p": 3, "b": 2, "s": 1, "t": 0}
	right := map[string]int{"m": 4, "q": 3, "d": 2, "z": 1, "j": 0}

	rightToLeft := map[string]string{"m": "w", "q": "p", "d": "b", "z": "s"}
	leftToRight := map[string]string{"w": "m", "p": "q", "b": "d", "s": "z"}

	characters := strings.Split(fight, "")
	for charIndex, char := range characters {
		switch char {
		case "t":
			applyEffect(characters, charIndex, rightToLeft, "j", right)
		case "j":
			applyEffect(characters, charIndex, leftToRight, "t", left)
		}
	}

	leftScore := 0
	rightScore := 0

	for _, char := range characters {
		if val, ok := left[char]; ok {
			leftScore += val
		} else if val, ok := right[char]; ok {
			rightScore += val
		}
	}

	if leftScore > rightScore {
		return "Left side wins!"
	} else if rightScore > leftScore {
		return "Right side wins!"
	}
	return "Let's fight again!"
}
