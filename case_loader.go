package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func LoadCase(path string) (*CaseDefinition, error) {
	data, err := os.ReadFile(filepath.Join(path, "case.json"))
	if err != nil {
		return nil, err
	}

	var caseDef CaseDefinition

	err = json.Unmarshal(data, &caseDef)
	if err != nil {
		return nil, err
	}

	clues, err := LoadClues(filepath.Join(path, "clues.json"))
	if err != nil {
		return nil, err
	}

	locations, err := LoadLocations(filepath.Join(path, "locations.json"))
	if err != nil {
		return nil, err
	}

	objects, err := LoadObjects(filepath.Join(path, "objects.json"))
	if err != nil {
		return nil, err
	}

	suspects, err := LoadSuspects(filepath.Join(path, "suspects.json"))
	if err != nil {
		return nil, err
	}

	dialogues, err := LoadDialogues(filepath.Join(path, "dialogues.json"))
	if err != nil {
		return nil, err
	}

	caseDef.Clues = clues
	caseDef.Locations = locations
	caseDef.Objects = objects
	caseDef.Suspects = suspects
	caseDef.Dialogues = dialogues

	return &caseDef, nil
}

func LoadClues(path string) ([]Clue, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var clues []Clue

	err = json.Unmarshal(data, &clues)
	if err != nil {
		return nil, err
	}

	return clues, nil
}

func LoadLocations(path string) ([]Location, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var locations []Location

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}

func LoadObjects(path string) ([]Object, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var objects []Object

	err = json.Unmarshal(data, &objects)
	if err != nil {
		return nil, err
	}

	return objects, nil
}

func LoadSuspects(path string) ([]Suspect, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var suspects []Suspect

	err = json.Unmarshal(data, &suspects)
	if err != nil {
		return nil, err
	}

	return suspects, nil
}

func LoadDialogues(path string) ([]Dialogue, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var dialogues []Dialogue

	err = json.Unmarshal(data, &dialogues)
	if err != nil {
		return nil, err
	}

	return dialogues, nil
}
