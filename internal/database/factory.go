package database

import (
	"encoding/json"

	"github.com/meilisearch/meilisearch-go"

	"github.com/behnh/plane-bot/internal/logger"
)

type AircraftStoreType struct {
	Client meilisearch.ServiceManager
	Index  string
}

var AircraftStore AircraftStoreType

func NewMemeStore(url, apiKey, index string) AircraftStoreType {
	client := meilisearch.New(url, meilisearch.WithAPIKey(apiKey))
	AircraftStore = AircraftStoreType{
		Client: client,
		Index:  index,
	}
	return AircraftStore
}

func (m *AircraftStoreType) GetAircraftById(id string) (Aircraft, error) {
	var aircraft Aircraft
	err := m.Client.Index(m.Index).GetDocument(id, &meilisearch.DocumentQuery{}, &aircraft)
	if err != nil {
		logger.Log.Errorf("Failed to get meme: %v", err)
		return Aircraft{}, err
	}

	return aircraft, nil
}

func (m *AircraftStoreType) SearchAircraft(query string) ([]Aircraft, error) {
	data, err := m.Client.Index(m.Index).Search(query, &meilisearch.SearchRequest{
		Limit: 25,
	})

	if err != nil {
		logger.Log.Errorf("Failed to search memes: %v", err)
		return []Aircraft{}, err
	}

	var memes []Aircraft
	hits, err := json.Marshal(data.Hits)
	if err != nil {
		logger.Log.Errorf("failed to marshal hits: %v", err)
		return nil, err
	}
	err = json.Unmarshal(hits, &memes)
	if err != nil {
		logger.Log.Errorf("failed to unmarshal hits: %v", err)
		return nil, err
	}

	logger.Log.Infof("Hits: %v", memes)

	return memes, nil
}

func (m *AircraftStoreType) AddAircraft(aircraft Aircraft) error {
	_, err := m.Client.Index(m.Index).AddDocuments(aircraft)
	if err != nil {
		logger.Log.Errorf("Failed to add meme: %v", err)
		return err
	}
	return nil
}

func (m *AircraftStoreType) DeleteAircraft(id string) error {
	_, err := m.Client.Index(m.Index).DeleteDocument(id)
	if err != nil {
		logger.Log.Errorf("Failed to delete meme: %v", err)
		return err
	}

	return nil
}
