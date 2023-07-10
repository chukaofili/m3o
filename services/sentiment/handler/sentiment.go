package handler

import (
	"context"

	"m3o.dev/platform/service/errors"
	"m3o.dev/services/sentiment/model"
	pb "m3o.dev/services/sentiment/proto"
)

type Sentiment struct{}

func (e *Sentiment) Analyze(ctx context.Context, req *pb.AnalyzeRequest, rsp *pb.AnalyzeResponse) error {
	if len(req.Text) == 0 {
		return errors.BadRequest("sentiment.analyze", "text is blank")
	}

	if len(req.Lang) == 0 {
		req.Lang = "english"
	}

	if req.Lang != "english" {
		return errors.BadRequest("sentiment.analyze", "only support english")
	}

	rsp.Score = model.Analyze(req.Text)

	// TODO: more complex word scoring

	return nil
}
