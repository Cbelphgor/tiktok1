// Copyright 2022 a76yyyy && CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * @Author: a76yyyy q981331502@163.com
 * @Date: 2022-06-12 09:11:35
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:02:00
 * @FilePath: /tiktok/cmd/feed/command/feed.go
 * @Description: 视频流 操作业务逻辑
 */

package command

import (
	"context"
	"sort"
	"time"

	"github.com/a76yyyy/tiktok/dal/pack"

	"github.com/a76yyyy/tiktok/kitex_gen/feed"

	"github.com/a76yyyy/tiktok/dal/db"
)

const (
	LIMIT = 30 // 单次返回最大视频数
)

type GetUserFeedService struct {
	ctx context.Context
}

// NewGetUserFeedService new GetUserFeedService
func NewGetUserFeedService(ctx context.Context) *GetUserFeedService {
	return &GetUserFeedService{ctx: ctx}
}

// GetUserFeed get feed info.
func (s *GetUserFeedService) GetUserFeed(req *feed.DouyinFeedRequest, fromID int64) (vis []*feed.Video, nextTime int64, err error) {
	videos, err := db.MGetVideos(s.ctx, LIMIT, req.LatestTime)
	if err != nil {
		return vis, nextTime, err
	}

	if len(videos) == 0 {
		return vis, nextTime, nil
	}

	if vis, err = pack.Videos(s.ctx, videos, &fromID); err != nil {
		return vis, nextTime, nil
	}

	if len(videos) > 0 {
		sort.Slice(videos, func(i, j int) bool {
			return videos[i].UpdatedAt.UnixMilli() > videos[j].UpdatedAt.UnixMilli()
		})
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	} else {
		nextTime = time.Now().UnixMilli()
	}

	return vis, nextTime, nil
}
