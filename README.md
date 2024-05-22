# EVM Chains Transfers History

Get a list of all ERC20 transfers where the from OR the to is part of a given list and have happened over the last X blocks.

The API implements the JSON-RPC specification and allow one single method: evm_transferHistory.

Thi app can be used for all EVM chains. (Chiliz EVM for exemple)

## evm_transferHistory

Returns list of ERC20 transfers.

### Parameters

* contractAddress: ERC20 Token Contract address (20 bytes).
* addresses: List of address involved in transfers
* lastBlocks: The X last blocks where transfers have happened

### Response

Returns an array of transfer object:

* from: transfer address from 
* to: transfer address to 
* value: the amount of tokens trasnfered
* transactionHash: the hash of the transaction
* blockHash: the hash of the block where transaction is stored
* blockNumber: the number of the block where transaction is stored

## How to use the app

You can use your desired rpc provider by updating the RPC_PROVIDER config in .env file.
Exemple: https://chiliz-rpc.publicnode.com for Chiliz Chain

### Run the tests

```bash
$ make test
```

### Start the app

```bash
$ make run
```

### Send a request

#### request

```bash
$ curl localhost:3000 \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"evm_transferHistory","params":[{"lastBlocks":20,"contractAddress":"0xdAC17F958D2ee523a2206206994597C13D831ec7", "addresses":["0x0a69babef1ca67a37ffaf7a485dfff3382056e78c"]}],"id":1}'
```

#### response

```bash
$ [{"from":"0x0a69babef1ca67a37ffaf7a485dfff3382056e78c","to":"0x011b815efb8f581194ae79006d24e0d814b7697f6","value":2147483647,"blockHash":"0x04ae1ee0792be04f491127d53d76bc77f37a709c08491f4fa9535ec07ec776b7","blockNumber":"0x1300c01","transactionHash":"0x27404e648ed94dfe8d8e0be8918ca594e133ca73ee98996aae7aa5cf30a40b52"},{"from":"0x0a69babef1ca67a37ffaf7a485dfff3382056e78c","to":"0x011b815efb8f581194ae79006d24e0d814b7697f6","value":2147483647,"blockHash":"0x3df3f5d03dcfb3c34046c99ebf62e07207c63cb115686c69e92c2896fdab5616","blockNumber":"0x1300c02","transactionHash":"0xd76dc4c6ca603bd997fdfba4578977c82078fe169b197a66f455103864abdce0"},{"from":"0x0a69babef1ca67a37ffaf7a485dfff3382056e78c","to":"0x0c7bbec68d12a0d1830360f8ec58fa599ba1b0e9b","value":2147483647,"blockHash":"0x3df3f5d03dcfb3c34046c99ebf62e07207c63cb115686c69e92c2896fdab5616","blockNumber":"0x1300c02","transactionHash":"0xc3bc7bfe3464b30201ecb68dec3757139bcb41c9b9814d50162fc4a6221e5693"},{"from":"0x0a69babef1ca67a37ffaf7a485dfff3382056e78c","to":"0x011b815efb8f581194ae79006d24e0d814b7697f6","value":2147483647,"blockHash":"0xe407e9b2cfb9bb0fdfa3c912312197cc817e4d8d708040ec0bddb7e896e8f256","blockNumber":"0x1300c03","transactionHash":"0xe736b7193b97d0f3073b2a2b2f12ba84772a4ff59b2d0eadc63a6d630437fb1d"},{"from":"0x0a69babef1ca67a37ffaf7a485dfff3382056e78c","to":"0x0bcc66fc7402daa98f5764057f95ac66b9391cd6b","value":2147483647,"blockHash":"0xe407e9b2cfb9bb0fdfa3c912312197cc817e4d8d708040ec0bddb7e896e8f256","blockNumber":"0x1300c03","transactionHash":"0x35a4868ea2d6d10a4f1ad0099095edf33db27f26c60180b892a55195e586e499"},{"from":"0x011b815efb8f581194ae79006d24e0d814b7697f6","to":"0x0a69babef1ca67a37ffaf7a485dfff3382056e78c","value":2147483647,"blockHash":"0xfd804897efd0f3109e8323f60881a27e515e48991e81b14f3dd068f531f8d17f","blockNumber":"0x1300c0d","transactionHash":"0xd6bdfacc505db39bf18bdce15b3904868c1579e10fde60d568ff9b665de4af53"}]
```