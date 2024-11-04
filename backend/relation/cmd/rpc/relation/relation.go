// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: relation.proto

package relation

import (
	"context"

	"github.com/ShellWen/GitPulse/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCreateRepoReq                 = pb.AddCreateRepoReq
	AddCreateRepoResp                = pb.AddCreateRepoResp
	AddFollowReq                     = pb.AddFollowReq
	AddFollowResp                    = pb.AddFollowResp
	AddForkReq                       = pb.AddForkReq
	AddForkResp                      = pb.AddForkResp
	AddStarReq                       = pb.AddStarReq
	AddStarResp                      = pb.AddStarResp
	BlockUntilCreatedRepoUpdatedReq  = pb.BlockUntilCreatedRepoUpdatedReq
	BlockUntilCreatedRepoUpdatedResp = pb.BlockUntilCreatedRepoUpdatedResp
	BlockUntilFollowerUpdatedReq     = pb.BlockUntilFollowerUpdatedReq
	BlockUntilFollowerUpdatedResp    = pb.BlockUntilFollowerUpdatedResp
	BlockUntilFollowingUpdatedReq    = pb.BlockUntilFollowingUpdatedReq
	BlockUntilFollowingUpdatedResp   = pb.BlockUntilFollowingUpdatedResp
	BlockUntilForkUpdatedReq         = pb.BlockUntilForkUpdatedReq
	BlockUntilForkUpdatedResp        = pb.BlockUntilForkUpdatedResp
	BlockUntilStarredRepoUpdatedReq  = pb.BlockUntilStarredRepoUpdatedReq
	BlockUntilStarredRepoUpdatedResp = pb.BlockUntilStarredRepoUpdatedResp
	CheckFollowResp                  = pb.CheckFollowResp
	CheckIfFollowReq                 = pb.CheckIfFollowReq
	CheckIfStarReq                   = pb.CheckIfStarReq
	CheckIfStarResp                  = pb.CheckIfStarResp
	CreateRepo                       = pb.CreateRepo
	DelAllCreatedRepoReq             = pb.DelAllCreatedRepoReq
	DelAllCreatedRepoResp            = pb.DelAllCreatedRepoResp
	DelAllFollowerReq                = pb.DelAllFollowerReq
	DelAllFollowerResp               = pb.DelAllFollowerResp
	DelAllFollowingReq               = pb.DelAllFollowingReq
	DelAllFollowingResp              = pb.DelAllFollowingResp
	DelAllForkReq                    = pb.DelAllForkReq
	DelAllForkResp                   = pb.DelAllForkResp
	DelAllStaringDevReq              = pb.DelAllStaringDevReq
	DelAllStaringDevResp             = pb.DelAllStaringDevResp
	DelAllStarredRepoReq             = pb.DelAllStarredRepoReq
	DelAllStarredRepoResp            = pb.DelAllStarredRepoResp
	DelCreateRepoReq                 = pb.DelCreateRepoReq
	DelCreateRepoResp                = pb.DelCreateRepoResp
	DelFollowReq                     = pb.DelFollowReq
	DelFollowResp                    = pb.DelFollowResp
	DelForkReq                       = pb.DelForkReq
	DelForkResp                      = pb.DelForkResp
	DelStarReq                       = pb.DelStarReq
	DelStarResp                      = pb.DelStarResp
	Follow                           = pb.Follow
	Fork                             = pb.Fork
	GetCreatorIdReq                  = pb.GetCreatorIdReq
	GetCreatorIdResp                 = pb.GetCreatorIdResp
	GetOriginReq                     = pb.GetOriginReq
	GetOriginResp                    = pb.GetOriginResp
	SearchCreatedRepoReq             = pb.SearchCreatedRepoReq
	SearchCreatedRepoResp            = pb.SearchCreatedRepoResp
	SearchFollowerByDeveloperIdReq   = pb.SearchFollowerByDeveloperIdReq
	SearchFollowerByDeveloperIdResp  = pb.SearchFollowerByDeveloperIdResp
	SearchFollowingByDeveloperIdReq  = pb.SearchFollowingByDeveloperIdReq
	SearchFollowingByDeveloperIdResp = pb.SearchFollowingByDeveloperIdResp
	SearchForkReq                    = pb.SearchForkReq
	SearchForkResp                   = pb.SearchForkResp
	SearchStaringDevReq              = pb.SearchStaringDevReq
	SearchStaringDevResp             = pb.SearchStaringDevResp
	SearchStarredRepoReq             = pb.SearchStarredRepoReq
	SearchStarredRepoResp            = pb.SearchStarredRepoResp
	Star                             = pb.Star

	Relation interface {
		// -----------------------createRepo-----------------------
		AddCreateRepo(ctx context.Context, in *AddCreateRepoReq, opts ...grpc.CallOption) (*AddCreateRepoResp, error)
		DelCreateRepo(ctx context.Context, in *DelCreateRepoReq, opts ...grpc.CallOption) (*DelCreateRepoResp, error)
		DelAllCreatedRepo(ctx context.Context, in *DelAllCreatedRepoReq, opts ...grpc.CallOption) (*DelAllCreatedRepoResp, error)
		GetCreatorId(ctx context.Context, in *GetCreatorIdReq, opts ...grpc.CallOption) (*GetCreatorIdResp, error)
		SearchCreatedRepo(ctx context.Context, in *SearchCreatedRepoReq, opts ...grpc.CallOption) (*SearchCreatedRepoResp, error)
		BlockUntilCreatedRepoUpdated(ctx context.Context, in *BlockUntilCreatedRepoUpdatedReq, opts ...grpc.CallOption) (*BlockUntilCreatedRepoUpdatedResp, error)
		// -----------------------follow-----------------------
		AddFollow(ctx context.Context, in *AddFollowReq, opts ...grpc.CallOption) (*AddFollowResp, error)
		DelFollow(ctx context.Context, in *DelFollowReq, opts ...grpc.CallOption) (*DelFollowResp, error)
		DelAllFollower(ctx context.Context, in *DelAllFollowerReq, opts ...grpc.CallOption) (*DelAllFollowerResp, error)
		DelAllFollowing(ctx context.Context, in *DelAllFollowingReq, opts ...grpc.CallOption) (*DelAllFollowingResp, error)
		CheckIfFollow(ctx context.Context, in *CheckIfFollowReq, opts ...grpc.CallOption) (*CheckFollowResp, error)
		SearchFollowingByDeveloperId(ctx context.Context, in *SearchFollowingByDeveloperIdReq, opts ...grpc.CallOption) (*SearchFollowingByDeveloperIdResp, error)
		SearchFollowerByDeveloperId(ctx context.Context, in *SearchFollowerByDeveloperIdReq, opts ...grpc.CallOption) (*SearchFollowerByDeveloperIdResp, error)
		BlockUntilFollowingUpdated(ctx context.Context, in *BlockUntilFollowingUpdatedReq, opts ...grpc.CallOption) (*BlockUntilFollowingUpdatedResp, error)
		BlockUntilFollowerUpdated(ctx context.Context, in *BlockUntilFollowerUpdatedReq, opts ...grpc.CallOption) (*BlockUntilFollowerUpdatedResp, error)
		// -----------------------fork-----------------------
		AddFork(ctx context.Context, in *AddForkReq, opts ...grpc.CallOption) (*AddForkResp, error)
		DelFork(ctx context.Context, in *DelForkReq, opts ...grpc.CallOption) (*DelForkResp, error)
		DelAllFork(ctx context.Context, in *DelAllForkReq, opts ...grpc.CallOption) (*DelAllForkResp, error)
		GetOrigin(ctx context.Context, in *GetOriginReq, opts ...grpc.CallOption) (*GetOriginResp, error)
		SearchFork(ctx context.Context, in *SearchForkReq, opts ...grpc.CallOption) (*SearchForkResp, error)
		BlockUntilForkUpdated(ctx context.Context, in *BlockUntilForkUpdatedReq, opts ...grpc.CallOption) (*BlockUntilForkUpdatedResp, error)
		// -----------------------star-----------------------
		AddStar(ctx context.Context, in *AddStarReq, opts ...grpc.CallOption) (*AddStarResp, error)
		DelStar(ctx context.Context, in *DelStarReq, opts ...grpc.CallOption) (*DelStarResp, error)
		DelAllStarredRepo(ctx context.Context, in *DelAllStarredRepoReq, opts ...grpc.CallOption) (*DelAllStarredRepoResp, error)
		DelAllStaringDev(ctx context.Context, in *DelAllStaringDevReq, opts ...grpc.CallOption) (*DelAllStaringDevResp, error)
		CheckIfStar(ctx context.Context, in *CheckIfStarReq, opts ...grpc.CallOption) (*CheckIfStarResp, error)
		SearchStarredRepo(ctx context.Context, in *SearchStarredRepoReq, opts ...grpc.CallOption) (*SearchStarredRepoResp, error)
		SearchStaringDev(ctx context.Context, in *SearchStaringDevReq, opts ...grpc.CallOption) (*SearchStaringDevResp, error)
		BlockUntilStarredRepoUpdated(ctx context.Context, in *BlockUntilStarredRepoUpdatedReq, opts ...grpc.CallOption) (*BlockUntilStarredRepoUpdatedResp, error)
	}

	defaultRelation struct {
		cli zrpc.Client
	}
)

func NewRelation(cli zrpc.Client) Relation {
	return &defaultRelation{
		cli: cli,
	}
}

// -----------------------createRepo-----------------------
func (m *defaultRelation) AddCreateRepo(ctx context.Context, in *AddCreateRepoReq, opts ...grpc.CallOption) (*AddCreateRepoResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.AddCreateRepo(ctx, in, opts...)
}

func (m *defaultRelation) DelCreateRepo(ctx context.Context, in *DelCreateRepoReq, opts ...grpc.CallOption) (*DelCreateRepoResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.DelCreateRepo(ctx, in, opts...)
}

func (m *defaultRelation) DelAllCreatedRepo(ctx context.Context, in *DelAllCreatedRepoReq, opts ...grpc.CallOption) (*DelAllCreatedRepoResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.DelAllCreatedRepo(ctx, in, opts...)
}

func (m *defaultRelation) GetCreatorId(ctx context.Context, in *GetCreatorIdReq, opts ...grpc.CallOption) (*GetCreatorIdResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.GetCreatorId(ctx, in, opts...)
}

func (m *defaultRelation) SearchCreatedRepo(ctx context.Context, in *SearchCreatedRepoReq, opts ...grpc.CallOption) (*SearchCreatedRepoResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.SearchCreatedRepo(ctx, in, opts...)
}

func (m *defaultRelation) BlockUntilCreatedRepoUpdated(ctx context.Context, in *BlockUntilCreatedRepoUpdatedReq, opts ...grpc.CallOption) (*BlockUntilCreatedRepoUpdatedResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.BlockUntilCreatedRepoUpdated(ctx, in, opts...)
}

// -----------------------follow-----------------------
func (m *defaultRelation) AddFollow(ctx context.Context, in *AddFollowReq, opts ...grpc.CallOption) (*AddFollowResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.AddFollow(ctx, in, opts...)
}

func (m *defaultRelation) DelFollow(ctx context.Context, in *DelFollowReq, opts ...grpc.CallOption) (*DelFollowResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.DelFollow(ctx, in, opts...)
}

func (m *defaultRelation) DelAllFollower(ctx context.Context, in *DelAllFollowerReq, opts ...grpc.CallOption) (*DelAllFollowerResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.DelAllFollower(ctx, in, opts...)
}

func (m *defaultRelation) DelAllFollowing(ctx context.Context, in *DelAllFollowingReq, opts ...grpc.CallOption) (*DelAllFollowingResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.DelAllFollowing(ctx, in, opts...)
}

func (m *defaultRelation) CheckIfFollow(ctx context.Context, in *CheckIfFollowReq, opts ...grpc.CallOption) (*CheckFollowResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.CheckIfFollow(ctx, in, opts...)
}

func (m *defaultRelation) SearchFollowingByDeveloperId(ctx context.Context, in *SearchFollowingByDeveloperIdReq, opts ...grpc.CallOption) (*SearchFollowingByDeveloperIdResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.SearchFollowingByDeveloperId(ctx, in, opts...)
}

func (m *defaultRelation) SearchFollowerByDeveloperId(ctx context.Context, in *SearchFollowerByDeveloperIdReq, opts ...grpc.CallOption) (*SearchFollowerByDeveloperIdResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.SearchFollowerByDeveloperId(ctx, in, opts...)
}

func (m *defaultRelation) BlockUntilFollowingUpdated(ctx context.Context, in *BlockUntilFollowingUpdatedReq, opts ...grpc.CallOption) (*BlockUntilFollowingUpdatedResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.BlockUntilFollowingUpdated(ctx, in, opts...)
}

func (m *defaultRelation) BlockUntilFollowerUpdated(ctx context.Context, in *BlockUntilFollowerUpdatedReq, opts ...grpc.CallOption) (*BlockUntilFollowerUpdatedResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.BlockUntilFollowerUpdated(ctx, in, opts...)
}

// -----------------------fork-----------------------
func (m *defaultRelation) AddFork(ctx context.Context, in *AddForkReq, opts ...grpc.CallOption) (*AddForkResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.AddFork(ctx, in, opts...)
}

func (m *defaultRelation) DelFork(ctx context.Context, in *DelForkReq, opts ...grpc.CallOption) (*DelForkResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.DelFork(ctx, in, opts...)
}

func (m *defaultRelation) DelAllFork(ctx context.Context, in *DelAllForkReq, opts ...grpc.CallOption) (*DelAllForkResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.DelAllFork(ctx, in, opts...)
}

func (m *defaultRelation) GetOrigin(ctx context.Context, in *GetOriginReq, opts ...grpc.CallOption) (*GetOriginResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.GetOrigin(ctx, in, opts...)
}

func (m *defaultRelation) SearchFork(ctx context.Context, in *SearchForkReq, opts ...grpc.CallOption) (*SearchForkResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.SearchFork(ctx, in, opts...)
}

func (m *defaultRelation) BlockUntilForkUpdated(ctx context.Context, in *BlockUntilForkUpdatedReq, opts ...grpc.CallOption) (*BlockUntilForkUpdatedResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.BlockUntilForkUpdated(ctx, in, opts...)
}

// -----------------------star-----------------------
func (m *defaultRelation) AddStar(ctx context.Context, in *AddStarReq, opts ...grpc.CallOption) (*AddStarResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.AddStar(ctx, in, opts...)
}

func (m *defaultRelation) DelStar(ctx context.Context, in *DelStarReq, opts ...grpc.CallOption) (*DelStarResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.DelStar(ctx, in, opts...)
}

func (m *defaultRelation) DelAllStarredRepo(ctx context.Context, in *DelAllStarredRepoReq, opts ...grpc.CallOption) (*DelAllStarredRepoResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.DelAllStarredRepo(ctx, in, opts...)
}

func (m *defaultRelation) DelAllStaringDev(ctx context.Context, in *DelAllStaringDevReq, opts ...grpc.CallOption) (*DelAllStaringDevResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.DelAllStaringDev(ctx, in, opts...)
}

func (m *defaultRelation) CheckIfStar(ctx context.Context, in *CheckIfStarReq, opts ...grpc.CallOption) (*CheckIfStarResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.CheckIfStar(ctx, in, opts...)
}

func (m *defaultRelation) SearchStarredRepo(ctx context.Context, in *SearchStarredRepoReq, opts ...grpc.CallOption) (*SearchStarredRepoResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.SearchStarredRepo(ctx, in, opts...)
}

func (m *defaultRelation) SearchStaringDev(ctx context.Context, in *SearchStaringDevReq, opts ...grpc.CallOption) (*SearchStaringDevResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.SearchStaringDev(ctx, in, opts...)
}

func (m *defaultRelation) BlockUntilStarredRepoUpdated(ctx context.Context, in *BlockUntilStarredRepoUpdatedReq, opts ...grpc.CallOption) (*BlockUntilStarredRepoUpdatedResp, error) {
	client := pb.NewRelationClient(m.cli.Conn())
	return client.BlockUntilStarredRepoUpdated(ctx, in, opts...)
}
