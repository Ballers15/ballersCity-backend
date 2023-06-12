package clicker

import (
	"github.com/nostressdev/tvx/internal/storages"
	"github.com/nostressdev/tvx/pb"
)

type TvxServer struct {
	pb.UnimplementedClickerBackendServer
	Storage storages.TvxStorage
}

func NewTvxServer(storage storages.TvxStorage) TvxServer {
	return TvxServer{
		Storage: storage,
	}
}
