package mock

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigenda/api/grpc/churner"
	"github.com/Layr-Labs/eigenda/core"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
)

type MockTransactor struct {
	mock.Mock
}

var _ core.Transactor = (*MockTransactor)(nil)

func (t *MockTransactor) GetBlockStaleMeasure(ctx context.Context) (uint32, error) {
	args := t.Called()
	return *new(uint32), args.Error(0)
}

func (t *MockTransactor) GetStoreDurationBlocks(ctx context.Context) (uint32, error) {
	args := t.Called()
	return *new(uint32), args.Error(0)
}

func (t *MockTransactor) GetRegisteredQuorumIdsForOperator(ctx context.Context, operator core.OperatorID) ([]core.QuorumID, error) {
	args := t.Called()
	result := args.Get(0)
	return result.([]core.QuorumID), args.Error(1)
}

func (t *MockTransactor) RegisterBLSPublicKey(ctx context.Context, keypair *core.KeyPair) error {
	args := t.Called()
	return args.Error(0)
}

func (t *MockTransactor) RegisterOperator(ctx context.Context, pubkeyG1 *core.G1Point, socket string, quorumIds []core.QuorumID) error {
	args := t.Called()
	return args.Error(0)
}

func (t *MockTransactor) RegisterOperatorWithChurn(ctx context.Context, pubkeyG1 *core.G1Point, socket string, quorumIds []core.QuorumID, churnReply *churner.ChurnReply) error {
	args := t.Called()
	return args.Error(0)
}

func (t *MockTransactor) DeregisterOperator(ctx context.Context, pubkeyG1 *core.G1Point, blockNumber uint32) error {
	args := t.Called()
	return args.Error(0)
}

func (t *MockTransactor) UpdateOperatorSocket(ctx context.Context, socket string) error {
	args := t.Called()
	return args.Error(0)
}

func (t *MockTransactor) GetOperatorStakes(ctx context.Context, operatorId core.OperatorID, blockNumber uint32) ([][]core.OperatorStake, []core.QuorumID, error) {
	args := t.Called()
	result0 := args.Get(0)
	result1 := args.Get(1)
	return result0.([][]core.OperatorStake), result1.([]core.QuorumID), args.Error(1)
}

func (t *MockTransactor) GetOperatorStakesForQuorums(ctx context.Context, quorums []core.QuorumID, blockNumber uint32) ([][]core.OperatorStake, error) {
	args := t.Called()
	result := args.Get(0)
	return result.([][]core.OperatorStake), args.Error(1)
}

func (t *MockTransactor) ConfirmBatch(ctx context.Context, batchHeader core.BatchHeader, quorums map[core.QuorumID]*core.QuorumResult, signatureAggregation core.SignatureAggregation) (*types.Receipt, error) {
	args := t.Called()
	var receipt *types.Receipt
	if args.Get(0) != nil {
		receipt = args.Get(0).(*types.Receipt)
	}
	return receipt, args.Error(1)
}

func (t *MockTransactor) StakeRegistry(ctx context.Context) (gethcommon.Address, error) {
	args := t.Called()
	result := args.Get(0)
	return result.(gethcommon.Address), args.Error(1)
}

func (t *MockTransactor) OperatorIDToAddress(ctx context.Context, operatorId core.OperatorID) (gethcommon.Address, error) {
	args := t.Called()
	result := args.Get(0)
	return result.(gethcommon.Address), args.Error(1)
}

func (t *MockTransactor) GetCurrentQuorumBitmapByOperatorId(ctx context.Context, operatorId core.OperatorID) (*big.Int, error) {
	args := t.Called()
	result := args.Get(0)
	return result.(*big.Int), args.Error(1)
}

func (t *MockTransactor) GetOperatorSetParams(ctx context.Context, quorumID core.QuorumID) (*core.OperatorSetParam, error) {
	args := t.Called()
	result := args.Get(0)
	return result.(*core.OperatorSetParam), args.Error(1)
}

func (t *MockTransactor) GetNumberOfRegisteredOperatorForQuorum(ctx context.Context, quorumID core.QuorumID) (uint32, error) {
	args := t.Called()
	result := args.Get(0)
	return result.(uint32), args.Error(1)
}

func (t *MockTransactor) WeightOfOperatorForQuorum(ctx context.Context, quorumID core.QuorumID, operator gethcommon.Address) (*big.Int, error) {
	args := t.Called()
	result := args.Get(0)
	return result.(*big.Int), args.Error(1)
}

func (t *MockTransactor) CalculateOperatorChurnApprovalDigestHash(
	ctx context.Context,
	operatorId core.OperatorID,
	operatorsToChurn []core.OperatorToChurn,
	salt [32]byte,
	expiry *big.Int,
) ([32]byte, error) {
	args := t.Called()
	result := args.Get(0)
	return result.([32]byte), args.Error(1)
}

func (t *MockTransactor) GetCurrentBlockNumber(ctx context.Context) (uint32, error) {
	args := t.Called()
	result := args.Get(0)
	return result.(uint32), args.Error(1)
}

func (t *MockTransactor) GetQuorumCount(ctx context.Context, blockNumber uint32) (uint16, error) {
	args := t.Called()
	result := args.Get(0)
	return result.(uint16), args.Error(1)
}

func (t *MockTransactor) PubkeyHashToOperator(ctx context.Context, operatorId core.OperatorID) (gethcommon.Address, error) {
	args := t.Called()
	result := args.Get(0)
	return result.(gethcommon.Address), args.Error(1)
}