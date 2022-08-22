package orchestrator

import "persons_generator/engine/entities/language"

func (o *Orchestrator) SearchLanguage(q string) (*language.Language, error) {
	return language.GetLanguageByName(q), nil
}
