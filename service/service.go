package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/drizzleent/nft-transactor/config"
	"github.com/drizzleent/nft-transactor/models"
	"github.com/drizzleent/nft-transactor/repository"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TokenService interface {
	CreateToken(context.Context, *models.CreateTokenRequest) (*models.Token, error)
	ListToken(context.Context, *models.RequestParams) (interface{}, error)
	TotalSupplyToken(context.Context) (interface{}, error)
}

type tokenService struct {
	repository repository.TokenRepository
}

func NewTokenService(r repository.TokenRepository) TokenService {
	return &tokenService{
		repository: r,
	}
}

func (s *tokenService) CreateToken(ctx context.Context, req *models.CreateTokenRequest) (*models.Token, error) {
	token := models.Token{
		UniqueHash: generateString(20),
		MediaUrl:   req.MediaUrl,
		Owner:      req.Owner,
	}
	fmt.Println(token.UniqueHash)

	client, err := ethclient.Dial(config.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed connect rpc: %s error: %s", config.RPC, err.Error())
	}

	privateKey, err := crypto.HexToECDSA(config.PRIVATE_KEY)
	if err != nil {
		return nil, fmt.Errorf("failed get private key from hex: %s", err.Error())
	}
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed get public key from private: %s", err.Error())
	}

	fromAddress := crypto.PubkeyToAddress(*publicKey)
	contractAddress := common.HexToAddress(config.CONTRACT_ADDRESS)
	ownerAddress := common.HexToAddress(req.Owner)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed get nonce: %s", err.Error())
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed get gas price: %s", err.Error())
	}

	parsedAbi, err := abi.JSON(strings.NewReader(config.MINT_ABI))
	if err != nil {
		return nil, fmt.Errorf("failed parse abi: %s", err.Error())
	}

	call, err := parsedAbi.Pack("mint", ownerAddress, token.MediaUrl, token.UniqueHash)
	if err != nil {
		return nil, fmt.Errorf("failed pack  parsedAbi: %s", err.Error())
	}

	callMsg := ethereum.CallMsg{
		From: fromAddress,
		To:   &contractAddress,
		Data: call,
	}

	gasLimit, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		return nil, fmt.Errorf("failed get gas limit: %s", err.Error())
	}

	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, call)

	chainId, err := client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed get network id: %s", err.Error())
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed sign tx: %s", err.Error())
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed send transaction: %s", err.Error())
	}

	txHash := signedTx.Hash().Hex()
	token.TxHash = txHash

	err = s.repository.CreateToken(ctx, &token)
	if err != nil {
		return nil, fmt.Errorf("failed create token in data base: %s", err.Error())
	}

	return &token, nil
}

func (s *tokenService) ListToken(ctx context.Context, params *models.RequestParams) (interface{}, error) {

	resp, err := s.repository.ListToken(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed get token list: %s", err.Error())
	}
	return resp, nil
}

func (s *tokenService) TotalSupplyToken(ctx context.Context) (interface{}, error) {
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

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateString(l int) string {
	var builder strings.Builder
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	builder.Grow(l)

	for i := 0; i < l; i++ {
		s := symbols[seed.Intn(len(symbols))]
		builder.WriteByte(s)
	}

	return builder.String()
}
