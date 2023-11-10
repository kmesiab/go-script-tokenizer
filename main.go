package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"sync"
)

const (
	SupportedFileExtensions = ".json"
	InputFolderName         = "./input"
	OutputFolderName        = "./output"
	outputFilePermissions   = 0777
)

var stopWords = []string{
	".",
	",",
	";",
	"um",
	"uh",
}

func main() {
	// get the ars for the folder to scan
	err := ScanForTranscriptFiles(InputFolderName, onFileDetected)
	if err != nil {
		return
	}
}

func ScanForTranscriptFiles(folder string, callback FileDetectedCallback) error {
	var err error
	var entries []os.DirEntry

	if entries, err = os.ReadDir(folder); err != nil {
		return err
	}

	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}

		if filepath.Ext(entry.Name()) != SupportedFileExtensions {
			continue
		}

		callback(&entry)

	}

	return nil
}

func printStats(transcriptFile *TranscriptFile) {
	Logf("✅ Transcription file parsed successfully").
		Add("job_name", transcriptFile.JobName).
		Add("account_id", transcriptFile.AccountID).
		Add("status", transcriptFile.Status).
		Add("total_transcripts", strconv.Itoa(len(transcriptFile.Results.Transcript))).
		Add("total_words", strconv.Itoa(len(transcriptFile.Results.Items))).
		Info()

}

func onFileDetected(file *os.DirEntry) {
	var err error
	var transcriptFile *TranscriptFile

	// read and parse the input file
	if transcriptFile, err = parseTranscriptFile(file); err != nil {
		log.Printf("%s", err.Error())

		return
	}

	printStats(transcriptFile)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// start the goroutine to tokenize the transcript
	go func() {
		defer wg.Done()
		processTranscript(transcriptFile, (*file).Name())
	}()

	wg.Wait()
}

func processTranscript(transcriptFile *TranscriptFile, filename string) {

	// Get all the utterances and put them in order according to the speaker.
	tokenizedConversation, err := tokenizeUtterances(transcriptFile)
	if err != nil {
		return
	}

	sentences, err := generateSentences(tokenizedConversation)

	err = saveTranscriptFile(*sentences, filename)

	if err != nil {
		Logf("⚠️  Error saving transcript\n").
			Add("filename", filename).
			AddError(err).
			Error()

		return
	}

	Logf("✅ Transcript saved successfully").Info()

}

func saveTranscriptFile(sentences []string, filename string) error {

	var contents string

	for _, sentence := range sentences {
		contents += sentence
	}

	filename = "processed_" + filename
	fullPath, err := filepath.Abs(filename)

	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, []byte(contents), outputFilePermissions)

	if err != nil {
		return err
	}

	return nil
}

func generateSentences(conversation *[][]Utterance) (*[]string, error) {

	var sentences []string

	// sentenceFormat = `[start_time, end_time, speaker_label, content]`
	const sentenceFormat = "%s: %s\n"

	for _, utterances := range *conversation {

		var speaker = utterances[0].SpeakerLabel

		var content string
		for _, utterance := range utterances {

			content += " " + utterance.Words
		}

		sentence := fmt.Sprintf(sentenceFormat, speaker, content)
		sentences = append(sentences, sentence)
	}

	return &sentences, nil
}

func tokenizeUtterances(transcriptFile *TranscriptFile) (*[][]Utterance, error) {

	var currentSpeaker string   // Keep track of the current speaker, to break the sentence
	var sentences [][]Utterance // A slice of pointers to slices of pointers to Utterances
	var sentence []Utterance    // A slice of pointers to Utterances

	for _, item := range transcriptFile.Results.Items {

		if len(item.Alternatives) != 1 {
			Logf("Found %d alternative pronunciations. Halting", len(item.Alternatives)).Fatal()
			return nil, fmt.Errorf("found %d alternative pronunciations", len(item.Alternatives))
		}

		if slices.Contains(stopWords, item.Alternatives[0].Content) {
			continue
		}

		if currentSpeaker != item.SpeakerLabel {

			// Reset the current speaker
			currentSpeaker = item.SpeakerLabel

			// If we detected some speech, add the sentence
			// to the list of sentences
			if len(sentence) > 0 {
				sentences = append(sentences, sentence)
			}

			// The current collection of utterances for currentSpeaker
			sentence = []Utterance{}
		}

		sentence = append(sentence, Utterance{
			SpeakerLabel: item.SpeakerLabel,
			StartTime:    item.StartTime,
			EndTime:      item.EndTime,
			Words:        item.Alternatives[0].Content,
		})
	}

	return &sentences, nil
}

func parseTranscriptFile(file *os.DirEntry) (*TranscriptFile, error) {
	var transcriptFile TranscriptFile

	filename := OutputFolderName + "/" + (*file).Name()
	var absFilename string
	var fileBytes []byte
	var err error

	if absFilename, err = filepath.Abs(filename); err != nil {
		return nil, err
	}

	if fileBytes, err = os.ReadFile(absFilename); err != nil {
		return nil, err
	}

	fileContents := string(fileBytes)

	if json.Unmarshal([]byte(fileContents), &transcriptFile) != nil {
		return nil, err
	}

	return &transcriptFile, nil
}
