/*

Package scale provides a method that returns a mucic scale

I couldn't figure it out, but thanks to Kotus9. I promise I didn't copy and paste. I alse made a few adjustments

Kotus9 Solution - https://exercism.io/tracks/go/exercises/scale-generator/solutions/7141290a7c3545d2b2b73d182c78015a
*/
package scale

import "strings"

var intervalMatcher = map[rune]int{
	'm': 1,
	'M': 2,
	'A': 3,
}

var sharpNotes = []string{
	"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#",
}

var flatNotes = []string{
	"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab",
}

var useSharpNotes = []string{
	"C", "G", "D", "A", "a", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#",
}

func isNoteInSlice(note string, noteSlice []string) (isFound bool) {
	for _, noteInSlice := range noteSlice {
		if note == noteInSlice {
			return true
		}
	}
	return
}

func notesToUse(note string) []string {
	if isNoteInSlice(note, useSharpNotes) {
		return sharpNotes
	}
	return flatNotes
}

func getStartingPoint(note string, noteSlice []string) (index int) {
	for index, noteInSlice := range noteSlice {
		if strings.ToLower(note) == strings.ToLower(noteInSlice) {
			return index
		}
	}
	return
}

func indexToTake(interval string) []int {
	indexes := []int{0}
	nextIndex := 0

	for _, nextInterval := range interval[:(len(interval) - 1)] {
		nextIndex += intervalMatcher[nextInterval]
		indexes = append(indexes, nextIndex)
	}
	return indexes
}

//Scale returns a slice of strings
func Scale(tonic, interval string) (scale []string) {

	notes := notesToUse(tonic)

	startingPoint := getStartingPoint(tonic, notes)
	fullNoteScale := append(notes[startingPoint:], notes[:startingPoint]...)

	if interval == "" {
		return fullNoteScale
	}

	for _, index := range indexToTake(interval) {
		scale = append(scale, fullNoteScale[index])
	}
	return
}
