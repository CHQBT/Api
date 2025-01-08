package service

import (
	"context"
	"fmt"
	"milliy/generated/api"
	"milliy/model"
	"milliy/storage"
	"milliy/upload"
)

type twitService struct {
	api.UnimplementedTwitServiceServer
	storage  storage.IStorage
	uploader *upload.MinioUploader
}

func NewTwitService(storage storage.IStorage, uploader *upload.MinioUploader) *twitService {
	return &twitService{
		storage:  storage,
		uploader: uploader,
	}
}

func (s *twitService) CreateTwit(ctx context.Context, req *api.CreateTwitReq) (*api.TWitId, error) {
	id, err := s.storage.Twit().CreateTwit(&model.CreateTwitRequest{
		UserID:       req.UserId,
		PublisherFIO: req.PublisherFio,
		Type:         req.Type,
		Texts:        req.Texts,
		Title:        req.Title,
	})
	if err != nil {
		return nil, err
	}
	return &api.TWitId{Id: id}, nil
}

func (s *twitService) GetTwit(ctx context.Context, req *api.TWitId) (*api.Twit, error) {
	twit, err := s.storage.Twit().GetTwitByID(req.Id)
	if err != nil {
		return nil, err
	}

	// Get videos
	videos, err := s.storage.Video().GetByTwitID(req.Id)
	if err != nil {
		return nil, err
	}

	// Get photos
	photos, err := s.storage.Photo().GetByTwitID(req.Id)
	if err != nil {
		return nil, err
	}

	// Get musics
	musics, err := s.storage.Music().GetByTwitID(req.Id)
	if err != nil {
		return nil, err
	}

	// Get locations
	locations, err := s.storage.Location().GetByTwitID(req.Id)
	if err != nil {
		return nil, err
	}

	// Get URLs
	urls, err := s.storage.Url().GetByTwitID(req.Id)
	if err != nil {
		return nil, err
	}

	response := &api.Twit{
		Id:           twit.ID,
		UserId:       twit.UserID,
		PublisherFio: twit.PublisherFIO,
		Type:         twit.Type,
		Texts:        twit.Texts,
		Title:        twit.Title,
		ReadersCount: int32(twit.ReadersCount),
		CreatedAt:    twit.CreatedAt.String(),
	}

	// Add videos
	for _, v := range videos {
		response.Videos = append(response.Videos, &api.VideoInfo{Video: v.Video})
	}

	// Add photos
	for _, p := range photos {
		response.Photos = append(response.Photos, &api.PhotoInfo{Photo: p.Photo})
	}

	// Add musics
	for _, m := range musics {
		response.Musics = append(response.Musics, &api.MusicInfo{Mp3: m.MP3})
	}

	// Add locations
	for _, l := range locations {
		response.Locations = append(response.Locations, &api.LocationInfo{
			Lat: l.Lat,
			Lon: l.Lon,
		})
	}

	// Add URLs
	for _, u := range urls {
		response.Urls = append(response.Urls, &api.UrlInfo{Url: u.URL})
	}

	return response, nil
}

func (s *twitService) GetAllTwits(ctx context.Context, req *api.Empty) (*api.TwitList, error) {
	ids, err := s.storage.Twit().GetAllTwits()
	if err != nil {
		return nil, err
	}
	return &api.TwitList{TwitIds: ids}, nil
}

func (s *twitService) GetTwitsByType(ctx context.Context, req *api.TypeRequest) (*api.TwitList, error) {
	ids, err := s.storage.Twit().GetTwitsByType(req.Type)
	if err != nil {
		return nil, err
	}
	return &api.TwitList{TwitIds: ids}, nil
}

func (s *twitService) GetMostViewedTwits(ctx context.Context, req *api.LimitRequest) (*api.TwitList, error) {
	ids, err := s.storage.Twit().GetMostViewedTwit(int(req.Limit))
	if err != nil {
		return nil, err
	}
	return &api.TwitList{TwitIds: ids}, nil
}

func (s *twitService) GetLatestTwits(ctx context.Context, req *api.LimitRequest) (*api.TwitList, error) {
	ids, err := s.storage.Twit().GetLatestTwits(int(req.Limit))
	if err != nil {
		return nil, err
	}
	return &api.TwitList{TwitIds: ids}, nil
}

func (s *twitService) SearchTwits(ctx context.Context, req *api.SearchRequest) (*api.TwitList, error) {
	ids, err := s.storage.Twit().SearchTwit(req.Keyword)
	if err != nil {
		return nil, err
	}
	return &api.TwitList{TwitIds: ids}, nil
}

func (s *twitService) DeleteTwit(ctx context.Context, req *api.TWitId) (*api.Empty, error) {
	err := s.storage.Twit().DeleteTwit(req.Id)
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}

func (s *twitService) AddCountToTwit(ctx context.Context, req *api.TWitId) (*api.Empty, error) {
	err := s.storage.Twit().AddReadersCount(req.Id)
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}

// Create methods for related entities
func (s *twitService) CreateVideo(ctx context.Context, req *api.CreateVideoReq) (*api.VideoId, error) {
	// Upload file to Minio
	url, err := s.uploader.UploadFile("videos", req.File, req.FileName)
	if err != nil {
		return nil, fmt.Errorf("failed to upload video: %v", err)
	}

	// Save to database
	id, err := s.storage.Video().Create(&model.CreateVideoRequest{
		TwitID: req.TwitId,
		Video:  url,
	})
	if err != nil {
		return nil, err
	}
	return &api.VideoId{Id: id, Url: url}, nil
}

func (s *twitService) CreatePhoto(ctx context.Context, req *api.CreatePhotoReq) (*api.PhotoId, error) {
	// Upload file to Minio
	url, err := s.uploader.UploadFile("photos", req.File, req.FileName)
	if err != nil {
		return nil, fmt.Errorf("failed to upload photo: %v", err)
	}

	// Save to database
	id, err := s.storage.Photo().Create(&model.CreatePhotoRequest{
		TwitID: req.TwitId,
		Photo:  url,
	})
	if err != nil {
		return nil, err
	}
	return &api.PhotoId{Id: id, Url: url}, nil
}

func (s *twitService) CreateMusic(ctx context.Context, req *api.CreateMusicReq) (*api.MusicId, error) {
	// Upload file to Minio
	url, err := s.uploader.UploadFile("musics", req.File, req.FileName)
	if err != nil {
		return nil, fmt.Errorf("failed to upload music: %v", err)
	}

	// Save to database
	id, err := s.storage.Music().Create(&model.CreateMusicRequest{
		TwitID: req.TwitId,
		MP3:    url,
	})
	if err != nil {
		return nil, err
	}
	return &api.MusicId{Id: id, Url: url}, nil
}

func (s *twitService) CreateLocation(ctx context.Context, req *api.CreateLocationReq) (*api.LocationId, error) {
	id, err := s.storage.Location().Create(&model.CreateLocationRequest{
		TwitID: req.TwitId,
		Lat:    req.Lat,
		Lon:    req.Lon,
	})
	if err != nil {
		return nil, err
	}
	return &api.LocationId{Id: id}, nil
}

func (s *twitService) CreateUrl(ctx context.Context, req *api.CreateUrlReq) (*api.UrlId, error) {
	id, err := s.storage.Url().Create(&model.CreateURLRequest{
		TwitID: req.TwitId,
		URL:    req.Url,
	})
	if err != nil {
		return nil, err
	}
	return &api.UrlId{Id: id}, nil
}
