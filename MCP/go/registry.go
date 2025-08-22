package main

import (
	"github.com/apidapp/mcp-server/config"
	"github.com/apidapp/mcp-server/models"
	tools_transaction "github.com/apidapp/mcp-server/tools/transaction"
	tools_version "github.com/apidapp/mcp-server/tools/version"
	tools_general "github.com/apidapp/mcp-server/tools/general"
	tools_account "github.com/apidapp/mcp-server/tools/account"
	tools_blockchain "github.com/apidapp/mcp-server/tools/blockchain"
	tools_wallet "github.com/apidapp/mcp-server/tools/wallet"
	tools_block "github.com/apidapp/mcp-server/tools/block"
	tools_erc20 "github.com/apidapp/mcp-server/tools/erc20"
	tools_key "github.com/apidapp/mcp-server/tools/key"
	tools_contract "github.com/apidapp/mcp-server/tools/contract"
	tools_echo "github.com/apidapp/mcp-server/tools/echo"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_transaction.CreateGet_transaction_hashTool(cfg),
		tools_transaction.CreateOptions_transaction_hashTool(cfg),
		tools_version.CreateOptions_versionTool(cfg),
		tools_version.CreateGet_versionTool(cfg),
		tools_general.CreateOptionsTool(cfg),
		tools_account.CreateOptions_account_idTool(cfg),
		tools_account.CreateGet_account_idTool(cfg),
		tools_blockchain.CreateOptions_blockchainTool(cfg),
		tools_blockchain.CreateGet_blockchainTool(cfg),
		tools_wallet.CreateGet_walletTool(cfg),
		tools_wallet.CreateOptions_walletTool(cfg),
		tools_wallet.CreatePost_walletTool(cfg),
		tools_block.CreateGet_blockTool(cfg),
		tools_block.CreateOptions_blockTool(cfg),
		tools_erc20.CreateOptions_erc20Tool(cfg),
		tools_erc20.CreatePost_erc20Tool(cfg),
		tools_erc20.CreateGet_erc20Tool(cfg),
		tools_erc20.CreateOptions_erc20_addressTool(cfg),
		tools_erc20.CreatePost_erc20_addressTool(cfg),
		tools_erc20.CreateGet_erc20_addressTool(cfg),
		tools_wallet.CreatePost_wallet_account_id_erc20Tool(cfg),
		tools_blockchain.CreateGet_blockchain_idTool(cfg),
		tools_blockchain.CreateOptions_blockchain_idTool(cfg),
		tools_account.CreateOptions_accountTool(cfg),
		tools_account.CreatePost_accountTool(cfg),
		tools_transaction.CreateGet_transaction_hash_receiptTool(cfg),
		tools_transaction.CreateOptions_transaction_hash_receiptTool(cfg),
		tools_wallet.CreateGet_wallet_accountTool(cfg),
		tools_wallet.CreateOptions_wallet_accountTool(cfg),
		tools_wallet.CreatePost_wallet_accountTool(cfg),
		tools_key.CreateOptions_keyTool(cfg),
		tools_key.CreatePost_keyTool(cfg),
		tools_key.CreateGet_keyTool(cfg),
		tools_key.CreateDelete_key_keyTool(cfg),
		tools_key.CreateOptions_key_keyTool(cfg),
		tools_wallet.CreateGet_wallet_account_idTool(cfg),
		tools_wallet.CreateOptions_wallet_account_idTool(cfg),
		tools_block.CreateGet_block_idTool(cfg),
		tools_block.CreateOptions_block_idTool(cfg),
		tools_contract.CreatePost_contractTool(cfg),
		tools_contract.CreateOptions_contractTool(cfg),
		tools_wallet.CreatePost_wallet_account_id_contractTool(cfg),
		tools_block.CreateGet_block_id_transactionTool(cfg),
		tools_block.CreateOptions_block_id_transactionTool(cfg),
		tools_block.CreateGet_block_id_transaction_indexTool(cfg),
		tools_block.CreateOptions_block_id_transaction_indexTool(cfg),
		tools_contract.CreateGet_contract_idTool(cfg),
		tools_contract.CreateOptions_contract_idTool(cfg),
		tools_contract.CreatePost_contract_idTool(cfg),
		tools_wallet.CreatePost_wallet_account_id_payTool(cfg),
		tools_wallet.CreateOptions_wallet_account_id_payTool(cfg),
		tools_echo.CreateOptions_echoTool(cfg),
		tools_transaction.CreateOptions_transactionTool(cfg),
		tools_transaction.CreatePost_transactionTool(cfg),
	}
}
