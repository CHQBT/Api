package service

import (
	"fmt"
	"log/slog"
	"milliy/logs"
	"milliy/model"
	"milliy/storage"
)

type TwitService struct {
	storage storage.IStorage
	log     *slog.Logger
}

func NewTwitService(storage storage.IStorage) *TwitService {
	return &TwitService{
		storage: storage,
		log:     logs.NewLogger(),
	}
}

func (s *TwitService) CreateTwit(req *model.CreateTwitRequest) (string, error) {
	s.log.Info("CreateTwit rpc started")
	twitID, err := s.storage.Twit().CreateTwit(req)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error creating Twit: %v", err))
		return "", err
	}
	s.log.Info("CreateTwit rpc finished")
	return twitID, nil
}

func (s *TwitService) GetTwit(id string) (*model.TwitResponse, error) {
	s.log.Info("GetTwit rpc started")
	twit, err := s.storage.Twit().GetTwitByID(id)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting Twit by ID: %v", err))
		return nil, err
	}

	// Get videos
	videos, err := s.storage.Video().GetByTwitID(id)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting Twit by ID %v", err))
		return nil, err
	}

	// Get photos
	photos, err := s.storage.Photo().GetByTwitID(id)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting Photo by ID %v", err))
		return nil, err
	}

	// Get musics
	musics, err := s.storage.Music().GetByTwitID(id)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting Music by ID %v", err))
		return nil, err
	}

	// Get locations
	locations, err := s.storage.Location().GetByTwitID(id)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting Location by ID %v", err))
		return nil, err
	}

	// Get URLs
	urls, err := s.storage.Url().GetByTwitID(id)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting Url by ID %v", err))
		return nil, err
	}

	response := &model.TwitResponse{
		ID:           twit.ID,
		UserID:       twit.UserID,
		PublisherFio: twit.PublisherFIO,
		Type:         twit.Type,
		Texts:        twit.Texts,
		Title:        twit.Title,
		ReadersCount: int32(twit.ReadersCount),
		CreatedAt:    twit.CreatedAt.String(),
	}

	// Add videos
	for _, v := range videos {
		response.Videos = append(response.Videos, &model.VideoInfo{VideoID: v.ID, VideoURL: v.Video})
	}

	// Add photos
	for _, p := range photos {
		response.Photos = append(response.Photos, &model.PhotoInfo{PhotoID: p.ID, PhotoURL: p.Photo})
	}

	// Add musics
	for _, m := range musics {
		response.Musics = append(response.Musics, &model.MusicInfo{MusicID: m.ID, MusicURL: m.MP3})
	}

	// Add locations
	for _, l := range locations {
		response.Locations = append(response.Locations, &model.LocationInfo{
			LocationID: l.ID,
			Latitude:   l.Lat,
			Longitude:  l.Lon,
		})
	}

	// Add URLs
	for _, u := range urls {
		response.Urls = append(response.Urls, &model.UrlInfo{UrlID: u.ID, Url: u.URL})
	}

	s.log.Info("GetTwit rpc finished")
	return response, nil
}

func (s *TwitService) DeleteTwit(twitID string) error {
	s.log.Info("DeleteTwit rpc started")
	err := s.storage.Location().DeleteByTwitID(twitID)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error deleting Locations: %v", err))
		return err
	}
	err = s.storage.Url().DeleteByTwitID(twitID)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error deleting URLs: %v", err))
		return err
	}
	err = s.storage.Music().DeleteByTwitID(twitID)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error deleting Musics: %v", err))
		return err
	}
	err = s.storage.Photo().DeleteByTwitID(twitID)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error deleting Photos: %v", err))
		return err
	}
	err = s.storage.Video().DeleteByTwitID(twitID)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error deleting Videos: %v", err))
		return err
	}
	err = s.storage.Twit().DeleteTwit(twitID)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error deleting Twit: %v", err))
		return err
	}
	s.log.Info("DeleteTwit rpc finished")
	return nil
}

func (s *TwitService) AddReadersCount(id string) error {
	s.log.Info("AddReadersCount rpc started")
	err := s.storage.Twit().AddReadersCount(id)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error adding readers count: %v", err))
		return err
	}
	s.log.Info("AddReadersCount rpc finished")
	return nil
}

func (s *TwitService) GetAllTwits() ([]string, error) {
	s.log.Info("GetAllTwits rpc started")
	ids, err := s.storage.Twit().GetAllTwits()
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting all Twits: %v", err))
		return nil, err
	}
	s.log.Info("GetAllTwits rpc finished")
	return ids, nil
}

func (s *TwitService) GetTwitsByType(twitType string) ([]string, error) {
	s.log.Info("GetTwitsByType rpc started")
	ids, err := s.storage.Twit().GetTwitsByType(twitType)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting Twits by type: %v", err))
		return nil, err
	}
	s.log.Info("GetTwitsByType rpc finished")
	return ids, nil
}

func (s *TwitService) GetMostViewedTwit(limit int) ([]string, error) {
	s.log.Info("GetMostViewedTwit rpc started")
	ids, err := s.storage.Twit().GetMostViewedTwit(limit)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting most viewed Twits: %v", err))
		return nil, err
	}
	s.log.Info("GetMostViewedTwit rpc finished")
	return ids, nil
}

func (s *TwitService) GetLatestTwits(limit int) ([]string, error) {
	s.log.Info("GetLatestTwits rpc started")
	ids, err := s.storage.Twit().GetLatestTwits(limit)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting latest Twits: %v", err))
		return nil, err
	}
	s.log.Info("GetLatestTwits rpc finished")
	return ids, nil
}

func (s *TwitService) SearchTwit(keyword string) ([]string, error) {
	s.log.Info("SearchTwit rpc started")
	ids, err := s.storage.Twit().SearchTwit(keyword)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error searching Twits: %v", err))
		return nil, err
	}
	s.log.Info("SearchTwit rpc finished")
	return ids, nil
}

func (s *TwitService) CreateLocation(req *model.CreateLocationRequest) (string, error) {
	s.log.Info("CreateLocation rpc started")
	id, err := s.storage.Location().Create(req)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error creating Location: %v", err))
		return "", err
	}
	s.log.Info("CreateLocation rpc finished")
	return id, nil
}

func (s *TwitService) CreateMusic(req *model.CreateMusicRequest) (string, error) {
	s.log.Info("CreateMusic rpc started")
	id, err := s.storage.Music().Create(req)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error creating Music: %v", err))
		return "", err
	}
	s.log.Info("CreateMusic rpc finished")
	return id, nil
}

func (s *TwitService) CreatePhoto(req *model.CreatePhotoRequest) (string, error) {
	s.log.Info("CreatePhoto rpc started")
	id, err := s.storage.Photo().Create(req)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error creating Photo: %v", err))
		return "", err
	}
	s.log.Info("CreatePhoto rpc finished")
	return id, nil
}

func (s *TwitService) CreateVideo(req *model.CreateVideoRequest) (string, error) {
	s.log.Info("CreateVideo rpc started")
	id, err := s.storage.Video().Create(req)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error creating Video: %v", err))
		return "", err
	}
	s.log.Info("CreateVideo rpc finished")
	return id, nil
}

func (s *TwitService) CreateUrl(req *model.CreateURLRequest) (string, error) {
	s.log.Info("CreateUrl rpc started")
	id, err := s.storage.Url().Create(req)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error creating Url: %v", err))
		return "", err
	}
	s.log.Info("CreateUrl rpc finished")
	return id, nil
}

func (s *TwitService) GetUniqueTypes() ([]string, error) {
	s.log.Info("GetUniqueTypes rpc started")
	types, err := s.storage.Twit().GetUniqueTypes()
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting unique types: %v", err))
		return nil, err
	}
	s.log.Info("GetUniqueTypes rpc finished")
	return types, nil
}

func (s *TwitService) AddMainTwit(twitID, start_time, end_time string) error {
	s.log.Info("AddMainTwit rpc started")

	// Twitni bazaga qo‘shish
	err := s.storage.Twit().AddMainTwit(twitID, start_time, end_time)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error adding main twit: %v", err))
		return err
	}

	s.log.Info("AddMainTwit rpc finished")
	return nil
}

func (s *TwitService) GetMainTwit() ([]string, error) {
	s.log.Info("GetMainTwit rpc started")

	// Hozirgi vaqtga mos keladigan twitlarni olish
	twits, err := s.storage.Twit().GetMainTwit()
	if err != nil {
		s.log.Error(fmt.Sprintf("Error getting main twits: %v", err))
		return nil, err
	}

	s.log.Info("GetMainTwit rpc finished")
	return twits, nil
}

func (s *TwitService) DeleteMainTwit(twitID string) error {
	s.log.Info("DeleteMainTwit rpc started")

	// Twitni o‘chirish (soft delete)
	err := s.storage.Twit().DeleteMainTwit(twitID)
	if err != nil {
		s.log.Error(fmt.Sprintf("Error deleting main twit: %v", err))
		return err
	}

	s.log.Info("DeleteMainTwit rpc finished")
	return nil
}
