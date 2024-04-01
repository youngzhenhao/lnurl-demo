# lnurl-demo

*This demo was just tested on one Windows only.*

## Test Steps

0. Build the executable program

```bash
go build -o ListAllUsers.exe main.go
go build -o ListAllInvoices.exe main.go
go build -o DecodeLnurl.exe main.go
go build -o ServerRun.exe main.go
go build -o PhoneRun.exe main.go
go build -o UploadUserInfoRun.exe main.go
go build -o PayToLnurlRun.exe main.go
```

1. In Alice's lnd configuration file, `lnd.conf`, write

```bash
# Alice-node
lnddir=C:/mySpace/walletDataDir/Alice-node
rpclisten=0.0.0.0:11009
restlisten=8180
listen=9835
[Bitcoin]
bitcoin.active=true
bitcoin.testnet=true
bitcoin.node=bitcoind
[Bitcoind]
bitcoind.rpchost=xxx.xxx.xxx.xxx
bitcoind.rpcuser=rpcuser
bitcoind.rpcpass=rpcpass
bitcoind.zmqpubrawblock=tcp://xxx.xxx.xxx.xxx:28332
bitcoind.zmqpubrawtx=tcp://xxx.xxx.xxx.xxx:28333
[rpcmiddleware]
rpcmiddleware.enable=true
[bolt]
db.bolt.auto-compact=true
```

2. Open a new pwsh terminal and start Alice's lnd

```bash
lnd --configfile C:/mySpace/walletDataDir/Alice-node/lnd.conf
```

3. In another terminal, Alice create the lnd wallet

```bash
lncli --network testnet --rpcserver 127.0.0.1:11009  --lnddir=C:/mySpace/walletDataDir/Alice-node create
```

Or unlock lnd wallet

```bash
lncli --network testnet --rpcserver 127.0.0.1:11009  --lnddir=C:/mySpace/walletDataDir/Alice-node unlock

```

4. In Bob's lnd configuration file, `lnd.conf`, write

```bash
# Bob-node
lnddir=C:/mySpace/walletDataDir/Bob-node
rpclisten=0.0.0.0:12009
restlisten=8280
listen=9935
[Bitcoin]
bitcoin.active=true
bitcoin.testnet=true
bitcoin.node=bitcoind
[Bitcoind]
# bitcoind.rpcpolling=true
bitcoind.rpchost=xxx.xxx.xxx.xxx
bitcoind.rpcuser=rpcuser
bitcoind.rpcpass=rpcpass
bitcoind.zmqpubrawblock=tcp://xxx.xxx.xxx.xxx:28332
bitcoind.zmqpubrawtx=tcp://xxx.xxx.xxx.xxx:28333
[rpcmiddleware]
rpcmiddleware.enable=true
[bolt]
db.bolt.auto-compact=true
```

5. Open a new pwsh terminal and start Bob's lnd

```bash
lnd --configfile C:/mySpace/walletDataDir/Bob-node/lnd.conf
```

6. In another terminal, Bob create the lnd wallet

```bash
lncli --network testnet --rpcserver 127.0.0.1:12009  --lnddir=C:/mySpace/walletDataDir/Bob-node create
```

Or unlock lnd wallet

```bash
lncli --network testnet --rpcserver 127.0.0.1:12009  --lnddir=C:/mySpace/walletDataDir/Bob-node unlock
```

7. Bob's lnd establish a connection to Alice's lnd

```bash
lncli --network testnet --rpcserver 127.0.0.1:12009  --lnddir=C:/mySpace/walletDataDir/Bob-node connect 03f2dfec54d4577a2808223e22cfb00353dfe78f831c2ee9d68884f7479164be7f@127.0.0.1:9835
```

8. Attention: Both Alice and Bob need to transfer some sats to the lnd node address

*When using testnet instead of regtest, Bob's lnd should open a channel to Alice's lnd with a balance of more than 30,000 sats (including the minimum open channel fee of 20,000 sats and 10,000 sats for `reserved_balance_anchor_chan`), and Alice's lnd should also have some balance.*

9. Bob's lnd open a channel to Alice's lnd

```bash
lncli --network testnet --rpcserver 127.0.0.1:12009  --lnddir=C:/mySpace/walletDataDir/Bob-node openchannel --node_key 03f2dfec54d4577a2808223e22cfb00353dfe78f831c2ee9d68884f7479164be7f --local_amt 20000
```

---

10. Go back to demo directory and configure .env

```text
SERVER_DOMAIN_OR_SOCKET="127.0.0.1:9080"
ALICE_RPC_SERVER="127.0.0.1:11009"
ALICE_TLS_CERT_PATH="C:/mySpace/walletDataDir/Alice-node/tls.cert"
ALICE_MACAROON_PATH="C:/mySpace/walletDataDir/Alice-node/data/chain/bitcoin/testnet/admin.macaroon"
BOB_RPC_SERVER="127.0.0.1:12009"
BOB_TLS_CERT_PATH="C:/mySpace/walletDataDir/Bob-node/tls.cert"
BOB_MACAROON_PATH="C:/mySpace/walletDataDir/Bob-node/data/chain/bitcoin/testnet/admin.macaroon"
```

11. Open a new pwsh terminal, Alice run service

```powershell
.\ServerRun.exe
```

12. Open a new pwsh terminal, Server run service

```powershell
.\PhoneRun.exe
```

13. In another terminal, Alice upload User info

```powershell
.\UploadUserInfoRun.exe -name "Alice" -socket "127.0.0.1:9090"
```

Return `LNURL`

```text
LNURL1DP68GUP69UHNZV3H9CCZUVPWXYARJVPCXQHHQCTE8A5KG0FCXVCRJEFSVDSJ6DEHV5EZ6DPJ8PSJ6CTR8QMZ6V3HX9NRWC3CX4SN2WRPCY5HJ7
```

14. Query and list all Users info

```powershell
.\ListAllUsers.exe
```

Return the Alice uploaded info

```text
{8309e0ca-77e2-428a-ac86-271f7b85a58a Alice 127.0.0.1:9090}
```

15. Test using `LNURL` to generate `QR Code`

```poershell
cd .\qrc\
node .\qrc.js --url LNURL1DP68GUP69UHNZV3H9CCZUVPWXYARJVPCXQHHQCTE8A5KG0FCXVCRJEFSVDSJ6DEHV5EZ6DPJ8PSJ6CTR8QMZ6V3HX9NRWC3CX4SN2WRPCY5HJ7
```

16. Test decoding `LNURL`

```powershell
.\DecodeLnurl.exe -lnu LNURL1DP68GUP69UHNZV3H9CCZUVPWXYARJVPCXQHHQCTE8A5KG0FCXVCRJEFSVDSJ6DEHV5EZ6DPJ8PSJ6CTR8QMZ6V3HX9NRWC3CX4SN2WRPCY5HJ7
```

Return url

```text
http://127.0.0.1:9080/pay?id=8309e0ca-77e2-428a-ac86-271f7b85a58a
```

17. Bob pay 100 sats to Alice

*(Bob has scanned the code to get `LNURL`)*

```powershell
.\PayToLnurlRun.exe -amount 100 -lnu LNURL1DP68GUP69UHNZV3H9CCZUVPWXYARJVPCXQHHQCTE8A5KG0FCXVCRJEFSVDSJ6DEHV5EZ6DPJ8PSJ6CTR8QMZ6V3HX9NRWC3CX4SN2WRPCY5HJ7
```

Return `PaymentHash` of the transaction

```text
2fc7a7f6259d1ab783a7523d6107c1277f3487d17c7c53acd522dcfd7512b1df
```

18. Query and list all Invoices info

```powershell
.\ListAllInvoices.exe
```

Return invoice info created by Alice

```text
{be04b2a4-45d8-4b48-91f4-2ef7f02837c9 100 LNTB1U1PNQ5R72PP59LR60A39N5DT0QA82G7KZP7PYALNFP7303798TX4YTW06AGJK80SDQQCQZZSXQYZ5VQSP5ACD5AKJTXZ828256XGJYEH8DQF0A2QJT3853HZTCRNACAPQPAHUQ9QYYSSQVLAZTLEU509FL0RHLD24EX87NG754AJXX6DFKQ3JQQMP0U04AQHYMXLDHACWF5CP85AHYTVVYH7ZF7VLNM4MHJYAFADCL5FNDQXTC0CQ4XR09M}
```

## Design

![LNURL](./LNURL.jpg)
