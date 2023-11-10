package main

import "os"

type FileDetectedCallback func(file *os.DirEntry)

type TranscriptFile struct {
	JobName       string           `json:"jobName"`
	AccountID     string           `json:"accountId"`
	Status        string           `json:"status"`
	Results       TranscriptResult `json:"results"`
	SpeakerLabels SpeakerLabel     `json:"speaker_labels"`
}

type TranscriptResult struct {
	Transcript []Transcript     `json:"transcripts"`
	Items      []TranscriptItem `json:"items"`
}

type Transcript struct {
	Text string `json:"transcript"`
}

type TranscriptItem struct {
	Type         string        `json:"type"`
	SpeakerLabel string        `json:"speaker_label"`
	Alternatives []Alternative `json:"alternatives"`
	StartTime    string        `json:"start_time"`
	EndTime      string        `json:"end_time"`
}

type Alternative struct {
	Confidence string `json:"confidence"`
	Content    string `json:"content"`
}

type SpeakerLabel struct {
	StartTime    string              `json:"start_time"`
	EndTime      string              `json:"end_time"`
	SpeakerLabel string              `json:"speaker_label"`
	Items        []SpeakerLabelItems `json:"items"`
}

type SpeakerLabelItems struct {
	SpeakerLabel string `json:"speaker_label"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
}

type Utterance struct {
	SpeakerLabel string `json:"speaker_label"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Words        string `json:"words"`
}

type Sentence struct {
	Content   string `json:"content"`
	Speaker   string `json:"speaker"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
