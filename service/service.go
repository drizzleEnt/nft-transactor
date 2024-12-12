package service

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/drizzleent/nft-transactor/config"
	"github.com/drizzleent/nft-transactor/models"
	"github.com/drizzleent/nft-transactor/repository"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TokenService interface {
	CreateToken(*models.CreateTokenRequest) (*models.Token, error)
	ListToken() error
	TotalSupplyToken() (interface{}, error)
}

type tokenService struct {
	repository repository.TokenRepository
}

func NewTokenService(r repository.TokenRepository) TokenService {
	return &tokenService{
		repository: r,
	}
}

func (s *tokenService) CreateToken(*models.CreateTokenRequest) (*models.Token, error) {

	return nil, nil
}

func (s *tokenService) ListToken() error {
	return nil
}

func (s *tokenService) TotalSupplyToken() (interface{}, error) {
	ctx := context.Background()
	client, err := ethclient.Dial(config.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed connect rpc: %s error: %s", config.RPC, err.Error())
	}

	contractAddress := common.HexToAddress(config.CONTRACT_ADDRESS)

	parsedAbi, err := abi.JSON(strings.NewReader(config.TOTAL_SUPPLY_ABI))
	if err != nil {
		return nil, fmt.Errorf("failed parse abi: %s", err.Error())
	}

	call, err := parsedAbi.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("failed pack  parsedAbi: %s", err.Error())
	}

	res, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &contractAddress,
		Data: call,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed call contract: %s", err.Error())
	}

	var totalSupply *big.Int

	err = parsedAbi.UnpackIntoInterface(&totalSupply, "totalSupply", res)
	if err != nil {
		return nil, fmt.Errorf("failed unpack parsedAbi: %s", err.Error())
	}

	return totalSupply, nil
}
