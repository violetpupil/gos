package ocr

import (
	"context"
	"errors"
	"math"
	"time"

	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

// https://cloud.google.com/vision/docs/ocr?hl=zh-cn
func GoogleOcr(trace, cred string, imageBytes []byte) (*OCRResult, error) {
	log := zap.L().With(zap.String("traceId", trace))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsJSON([]byte(cred)))
	if err != nil {
		log.Info("new client error", zap.Error(err))
		return nil, OCRError{ErrorMsg: err.Error()}
	}

	image := &visionpb.Image{
		Content: imageBytes,
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 0)
	if err != nil {
		log.Info("detect text error", zap.Error(err))
		return nil, OCRError{ErrorMsg: err.Error()}
	}

	var result OCRResult
	for _, annotation := range annotations {
		if annotation == nil ||
			annotation.BoundingPoly == nil ||
			len(annotation.BoundingPoly.Vertices) < 4 {
			log.Error("info absent")
			return nil, errors.New("info absent")
		}

		log.Info("annotation result", zap.Any("annotation", annotation))

		// 坐标点顺序不固定
		var left, right, bottom, top int
		for _, vertex := range annotation.BoundingPoly.Vertices {
			if left == 0 {
				left = int(vertex.X)
			} else {
				left = int(math.Min(float64(left), float64(vertex.X)))
			}
			right = int(math.Max(float64(right), float64(vertex.X)))
			if bottom == 0 {
				bottom = int(vertex.Y)
			} else {
				bottom = int(math.Min(float64(bottom), float64(vertex.Y)))
			}
			top = int(math.Max(float64(top), float64(vertex.Y)))
		}
		width := right - left
		height := top - bottom
		if width <= 0 || height <= 0 {
			log.Error("info invalid")
			return nil, errors.New("info invalid")
		}

		result.WordsResult = append(result.WordsResult, WordsResult{
			Words: annotation.Description,
			Location: Location{
				Top:    uint32(top),
				Left:   uint32(left),
				Width:  uint32(width),
				Height: uint32(height),
			},
		})
	}

	return &result, nil
}
