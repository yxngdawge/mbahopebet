// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: bettor/v1alpha/bettor.proto

package bettorv1alphaconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1alpha "github.com/elh/bettor/api/bettor/v1alpha"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// BettorServiceName is the fully-qualified name of the BettorService service.
	BettorServiceName = "bettor.v1alpha.BettorService"
)

// BettorServiceClient is a client for the bettor.v1alpha.BettorService service.
type BettorServiceClient interface {
	// CreateUser creates a new user.
	CreateUser(context.Context, *connect_go.Request[v1alpha.CreateUserRequest]) (*connect_go.Response[v1alpha.CreateUserResponse], error)
	// GetUser returns a user by ID.
	GetUser(context.Context, *connect_go.Request[v1alpha.GetUserRequest]) (*connect_go.Response[v1alpha.GetUserResponse], error)
	// GetUserByUsername returns a user by ID.
	GetUserByUsername(context.Context, *connect_go.Request[v1alpha.GetUserByUsernameRequest]) (*connect_go.Response[v1alpha.GetUserByUsernameResponse], error)
	// ListUsers lists users by filters.
	ListUsers(context.Context, *connect_go.Request[v1alpha.ListUsersRequest]) (*connect_go.Response[v1alpha.ListUsersResponse], error)
	// CreateMarket creates a new betting market.
	CreateMarket(context.Context, *connect_go.Request[v1alpha.CreateMarketRequest]) (*connect_go.Response[v1alpha.CreateMarketResponse], error)
	// GetMarket gets a betting market by ID.
	GetMarket(context.Context, *connect_go.Request[v1alpha.GetMarketRequest]) (*connect_go.Response[v1alpha.GetMarketResponse], error)
	// ListMarkets lists markets by filters.
	ListMarkets(context.Context, *connect_go.Request[v1alpha.ListMarketsRequest]) (*connect_go.Response[v1alpha.ListMarketsResponse], error)
	// LockMarket locks a betting market preventing further bets.
	LockMarket(context.Context, *connect_go.Request[v1alpha.LockMarketRequest]) (*connect_go.Response[v1alpha.LockMarketResponse], error)
	// SettleMarket settles a betting market and pays out bets.
	SettleMarket(context.Context, *connect_go.Request[v1alpha.SettleMarketRequest]) (*connect_go.Response[v1alpha.SettleMarketResponse], error)
	// CreateBet places a bet on an open betting market.
	CreateBet(context.Context, *connect_go.Request[v1alpha.CreateBetRequest]) (*connect_go.Response[v1alpha.CreateBetResponse], error)
	// GetBet gets a bet.
	GetBet(context.Context, *connect_go.Request[v1alpha.GetBetRequest]) (*connect_go.Response[v1alpha.GetBetResponse], error)
	// ListBet lists bets by filters.
	ListBets(context.Context, *connect_go.Request[v1alpha.ListBetsRequest]) (*connect_go.Response[v1alpha.ListBetsResponse], error)
}

// NewBettorServiceClient constructs a client for the bettor.v1alpha.BettorService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBettorServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) BettorServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &bettorServiceClient{
		createUser: connect_go.NewClient[v1alpha.CreateUserRequest, v1alpha.CreateUserResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/CreateUser",
			opts...,
		),
		getUser: connect_go.NewClient[v1alpha.GetUserRequest, v1alpha.GetUserResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/GetUser",
			opts...,
		),
		getUserByUsername: connect_go.NewClient[v1alpha.GetUserByUsernameRequest, v1alpha.GetUserByUsernameResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/GetUserByUsername",
			opts...,
		),
		listUsers: connect_go.NewClient[v1alpha.ListUsersRequest, v1alpha.ListUsersResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/ListUsers",
			opts...,
		),
		createMarket: connect_go.NewClient[v1alpha.CreateMarketRequest, v1alpha.CreateMarketResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/CreateMarket",
			opts...,
		),
		getMarket: connect_go.NewClient[v1alpha.GetMarketRequest, v1alpha.GetMarketResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/GetMarket",
			opts...,
		),
		listMarkets: connect_go.NewClient[v1alpha.ListMarketsRequest, v1alpha.ListMarketsResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/ListMarkets",
			opts...,
		),
		lockMarket: connect_go.NewClient[v1alpha.LockMarketRequest, v1alpha.LockMarketResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/LockMarket",
			opts...,
		),
		settleMarket: connect_go.NewClient[v1alpha.SettleMarketRequest, v1alpha.SettleMarketResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/SettleMarket",
			opts...,
		),
		createBet: connect_go.NewClient[v1alpha.CreateBetRequest, v1alpha.CreateBetResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/CreateBet",
			opts...,
		),
		getBet: connect_go.NewClient[v1alpha.GetBetRequest, v1alpha.GetBetResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/GetBet",
			opts...,
		),
		listBets: connect_go.NewClient[v1alpha.ListBetsRequest, v1alpha.ListBetsResponse](
			httpClient,
			baseURL+"/bettor.v1alpha.BettorService/ListBets",
			opts...,
		),
	}
}

// bettorServiceClient implements BettorServiceClient.
type bettorServiceClient struct {
	createUser        *connect_go.Client[v1alpha.CreateUserRequest, v1alpha.CreateUserResponse]
	getUser           *connect_go.Client[v1alpha.GetUserRequest, v1alpha.GetUserResponse]
	getUserByUsername *connect_go.Client[v1alpha.GetUserByUsernameRequest, v1alpha.GetUserByUsernameResponse]
	listUsers         *connect_go.Client[v1alpha.ListUsersRequest, v1alpha.ListUsersResponse]
	createMarket      *connect_go.Client[v1alpha.CreateMarketRequest, v1alpha.CreateMarketResponse]
	getMarket         *connect_go.Client[v1alpha.GetMarketRequest, v1alpha.GetMarketResponse]
	listMarkets       *connect_go.Client[v1alpha.ListMarketsRequest, v1alpha.ListMarketsResponse]
	lockMarket        *connect_go.Client[v1alpha.LockMarketRequest, v1alpha.LockMarketResponse]
	settleMarket      *connect_go.Client[v1alpha.SettleMarketRequest, v1alpha.SettleMarketResponse]
	createBet         *connect_go.Client[v1alpha.CreateBetRequest, v1alpha.CreateBetResponse]
	getBet            *connect_go.Client[v1alpha.GetBetRequest, v1alpha.GetBetResponse]
	listBets          *connect_go.Client[v1alpha.ListBetsRequest, v1alpha.ListBetsResponse]
}

// CreateUser calls bettor.v1alpha.BettorService.CreateUser.
func (c *bettorServiceClient) CreateUser(ctx context.Context, req *connect_go.Request[v1alpha.CreateUserRequest]) (*connect_go.Response[v1alpha.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// GetUser calls bettor.v1alpha.BettorService.GetUser.
func (c *bettorServiceClient) GetUser(ctx context.Context, req *connect_go.Request[v1alpha.GetUserRequest]) (*connect_go.Response[v1alpha.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// GetUserByUsername calls bettor.v1alpha.BettorService.GetUserByUsername.
func (c *bettorServiceClient) GetUserByUsername(ctx context.Context, req *connect_go.Request[v1alpha.GetUserByUsernameRequest]) (*connect_go.Response[v1alpha.GetUserByUsernameResponse], error) {
	return c.getUserByUsername.CallUnary(ctx, req)
}

// ListUsers calls bettor.v1alpha.BettorService.ListUsers.
func (c *bettorServiceClient) ListUsers(ctx context.Context, req *connect_go.Request[v1alpha.ListUsersRequest]) (*connect_go.Response[v1alpha.ListUsersResponse], error) {
	return c.listUsers.CallUnary(ctx, req)
}

// CreateMarket calls bettor.v1alpha.BettorService.CreateMarket.
func (c *bettorServiceClient) CreateMarket(ctx context.Context, req *connect_go.Request[v1alpha.CreateMarketRequest]) (*connect_go.Response[v1alpha.CreateMarketResponse], error) {
	return c.createMarket.CallUnary(ctx, req)
}

// GetMarket calls bettor.v1alpha.BettorService.GetMarket.
func (c *bettorServiceClient) GetMarket(ctx context.Context, req *connect_go.Request[v1alpha.GetMarketRequest]) (*connect_go.Response[v1alpha.GetMarketResponse], error) {
	return c.getMarket.CallUnary(ctx, req)
}

// ListMarkets calls bettor.v1alpha.BettorService.ListMarkets.
func (c *bettorServiceClient) ListMarkets(ctx context.Context, req *connect_go.Request[v1alpha.ListMarketsRequest]) (*connect_go.Response[v1alpha.ListMarketsResponse], error) {
	return c.listMarkets.CallUnary(ctx, req)
}

// LockMarket calls bettor.v1alpha.BettorService.LockMarket.
func (c *bettorServiceClient) LockMarket(ctx context.Context, req *connect_go.Request[v1alpha.LockMarketRequest]) (*connect_go.Response[v1alpha.LockMarketResponse], error) {
	return c.lockMarket.CallUnary(ctx, req)
}

// SettleMarket calls bettor.v1alpha.BettorService.SettleMarket.
func (c *bettorServiceClient) SettleMarket(ctx context.Context, req *connect_go.Request[v1alpha.SettleMarketRequest]) (*connect_go.Response[v1alpha.SettleMarketResponse], error) {
	return c.settleMarket.CallUnary(ctx, req)
}

// CreateBet calls bettor.v1alpha.BettorService.CreateBet.
func (c *bettorServiceClient) CreateBet(ctx context.Context, req *connect_go.Request[v1alpha.CreateBetRequest]) (*connect_go.Response[v1alpha.CreateBetResponse], error) {
	return c.createBet.CallUnary(ctx, req)
}

// GetBet calls bettor.v1alpha.BettorService.GetBet.
func (c *bettorServiceClient) GetBet(ctx context.Context, req *connect_go.Request[v1alpha.GetBetRequest]) (*connect_go.Response[v1alpha.GetBetResponse], error) {
	return c.getBet.CallUnary(ctx, req)
}

// ListBets calls bettor.v1alpha.BettorService.ListBets.
func (c *bettorServiceClient) ListBets(ctx context.Context, req *connect_go.Request[v1alpha.ListBetsRequest]) (*connect_go.Response[v1alpha.ListBetsResponse], error) {
	return c.listBets.CallUnary(ctx, req)
}

// BettorServiceHandler is an implementation of the bettor.v1alpha.BettorService service.
type BettorServiceHandler interface {
	// CreateUser creates a new user.
	CreateUser(context.Context, *connect_go.Request[v1alpha.CreateUserRequest]) (*connect_go.Response[v1alpha.CreateUserResponse], error)
	// GetUser returns a user by ID.
	GetUser(context.Context, *connect_go.Request[v1alpha.GetUserRequest]) (*connect_go.Response[v1alpha.GetUserResponse], error)
	// GetUserByUsername returns a user by ID.
	GetUserByUsername(context.Context, *connect_go.Request[v1alpha.GetUserByUsernameRequest]) (*connect_go.Response[v1alpha.GetUserByUsernameResponse], error)
	// ListUsers lists users by filters.
	ListUsers(context.Context, *connect_go.Request[v1alpha.ListUsersRequest]) (*connect_go.Response[v1alpha.ListUsersResponse], error)
	// CreateMarket creates a new betting market.
	CreateMarket(context.Context, *connect_go.Request[v1alpha.CreateMarketRequest]) (*connect_go.Response[v1alpha.CreateMarketResponse], error)
	// GetMarket gets a betting market by ID.
	GetMarket(context.Context, *connect_go.Request[v1alpha.GetMarketRequest]) (*connect_go.Response[v1alpha.GetMarketResponse], error)
	// ListMarkets lists markets by filters.
	ListMarkets(context.Context, *connect_go.Request[v1alpha.ListMarketsRequest]) (*connect_go.Response[v1alpha.ListMarketsResponse], error)
	// LockMarket locks a betting market preventing further bets.
	LockMarket(context.Context, *connect_go.Request[v1alpha.LockMarketRequest]) (*connect_go.Response[v1alpha.LockMarketResponse], error)
	// SettleMarket settles a betting market and pays out bets.
	SettleMarket(context.Context, *connect_go.Request[v1alpha.SettleMarketRequest]) (*connect_go.Response[v1alpha.SettleMarketResponse], error)
	// CreateBet places a bet on an open betting market.
	CreateBet(context.Context, *connect_go.Request[v1alpha.CreateBetRequest]) (*connect_go.Response[v1alpha.CreateBetResponse], error)
	// GetBet gets a bet.
	GetBet(context.Context, *connect_go.Request[v1alpha.GetBetRequest]) (*connect_go.Response[v1alpha.GetBetResponse], error)
	// ListBet lists bets by filters.
	ListBets(context.Context, *connect_go.Request[v1alpha.ListBetsRequest]) (*connect_go.Response[v1alpha.ListBetsResponse], error)
}

// NewBettorServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBettorServiceHandler(svc BettorServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/bettor.v1alpha.BettorService/CreateUser", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/CreateUser",
		svc.CreateUser,
		opts...,
	))
	mux.Handle("/bettor.v1alpha.BettorService/GetUser", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/GetUser",
		svc.GetUser,
		opts...,
	))
	mux.Handle("/bettor.v1alpha.BettorService/GetUserByUsername", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/GetUserByUsername",
		svc.GetUserByUsername,
		opts...,
	))
	mux.Handle("/bettor.v1alpha.BettorService/ListUsers", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/ListUsers",
		svc.ListUsers,
		opts...,
	))
	mux.Handle("/bettor.v1alpha.BettorService/CreateMarket", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/CreateMarket",
		svc.CreateMarket,
		opts...,
	))
	mux.Handle("/bettor.v1alpha.BettorService/GetMarket", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/GetMarket",
		svc.GetMarket,
		opts...,
	))
	mux.Handle("/bettor.v1alpha.BettorService/ListMarkets", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/ListMarkets",
		svc.ListMarkets,
		opts...,
	))
	mux.Handle("/bettor.v1alpha.BettorService/LockMarket", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/LockMarket",
		svc.LockMarket,
		opts...,
	))
	mux.Handle("/bettor.v1alpha.BettorService/SettleMarket", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/SettleMarket",
		svc.SettleMarket,
		opts...,
	))
	mux.Handle("/bettor.v1alpha.BettorService/CreateBet", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/CreateBet",
		svc.CreateBet,
		opts...,
	))
	mux.Handle("/bettor.v1alpha.BettorService/GetBet", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/GetBet",
		svc.GetBet,
		opts...,
	))
	mux.Handle("/bettor.v1alpha.BettorService/ListBets", connect_go.NewUnaryHandler(
		"/bettor.v1alpha.BettorService/ListBets",
		svc.ListBets,
		opts...,
	))
	return "/bettor.v1alpha.BettorService/", mux
}

// UnimplementedBettorServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBettorServiceHandler struct{}

func (UnimplementedBettorServiceHandler) CreateUser(context.Context, *connect_go.Request[v1alpha.CreateUserRequest]) (*connect_go.Response[v1alpha.CreateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.CreateUser is not implemented"))
}

func (UnimplementedBettorServiceHandler) GetUser(context.Context, *connect_go.Request[v1alpha.GetUserRequest]) (*connect_go.Response[v1alpha.GetUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.GetUser is not implemented"))
}

func (UnimplementedBettorServiceHandler) GetUserByUsername(context.Context, *connect_go.Request[v1alpha.GetUserByUsernameRequest]) (*connect_go.Response[v1alpha.GetUserByUsernameResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.GetUserByUsername is not implemented"))
}

func (UnimplementedBettorServiceHandler) ListUsers(context.Context, *connect_go.Request[v1alpha.ListUsersRequest]) (*connect_go.Response[v1alpha.ListUsersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.ListUsers is not implemented"))
}

func (UnimplementedBettorServiceHandler) CreateMarket(context.Context, *connect_go.Request[v1alpha.CreateMarketRequest]) (*connect_go.Response[v1alpha.CreateMarketResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.CreateMarket is not implemented"))
}

func (UnimplementedBettorServiceHandler) GetMarket(context.Context, *connect_go.Request[v1alpha.GetMarketRequest]) (*connect_go.Response[v1alpha.GetMarketResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.GetMarket is not implemented"))
}

func (UnimplementedBettorServiceHandler) ListMarkets(context.Context, *connect_go.Request[v1alpha.ListMarketsRequest]) (*connect_go.Response[v1alpha.ListMarketsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.ListMarkets is not implemented"))
}

func (UnimplementedBettorServiceHandler) LockMarket(context.Context, *connect_go.Request[v1alpha.LockMarketRequest]) (*connect_go.Response[v1alpha.LockMarketResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.LockMarket is not implemented"))
}

func (UnimplementedBettorServiceHandler) SettleMarket(context.Context, *connect_go.Request[v1alpha.SettleMarketRequest]) (*connect_go.Response[v1alpha.SettleMarketResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.SettleMarket is not implemented"))
}

func (UnimplementedBettorServiceHandler) CreateBet(context.Context, *connect_go.Request[v1alpha.CreateBetRequest]) (*connect_go.Response[v1alpha.CreateBetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.CreateBet is not implemented"))
}

func (UnimplementedBettorServiceHandler) GetBet(context.Context, *connect_go.Request[v1alpha.GetBetRequest]) (*connect_go.Response[v1alpha.GetBetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.GetBet is not implemented"))
}

func (UnimplementedBettorServiceHandler) ListBets(context.Context, *connect_go.Request[v1alpha.ListBetsRequest]) (*connect_go.Response[v1alpha.ListBetsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bettor.v1alpha.BettorService.ListBets is not implemented"))
}
