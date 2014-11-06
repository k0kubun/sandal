## 既存の障害とその実装
### Abrupt termination of a process
- **[process]** プロセスの遷移の変更 (終了状態への遷移の追加)

### Message drop in a channel
- **[channel]** そのチャネルを使ったsend statementの遷移の変更

### Timeout in receiving a message
- **[statement]** recv statementの遷移, 返り値の変更

## その他の障害と実装に必要なもの
### Omission Failure
- **[statement]** send statementの遷移の変更 (send-omission failure)
- **[statement]** recv statementの遷移, 返り値の変更 (recv-omission failure)
- **[channel]** そのチャネルを使ったsend, recvの遷移の変更 (channel omission failure)

### Timing failure
- **[statement]** recv statementの遷移の変更

### Byzantine Failure (Arbitrary failure, Commission failure)
- **[statement]** send statementの動作の変更 (チャネルに送信する値を変更)
- **[statement]** recv statementの動作の変更 (チャネルから受け取る値を変更)

### Value error
- **[variable]** 変数の参照結果の変更

### メッセージ順序の入れ替わり
- **[channel]** buffered channelから値を返す順番の変更

## 障害挿入対象
- **statement** (6種, 現在1)
  - 遷移, 動作, 返り値の変更
- **variable** (1種, 現在0)
  - 変数を参照した結果の書き換え
- **channel** (3種, 現在1)
  - チャネルを使うstatementの遷移の変更
- **process** (1種, 現在1)
  - 特定の遷移の追加

## 拡張案

```go
fault recv(recvCh channel { int }, variable *int) @omission {
  // receive nothing
}

fault send(sendCh channel { int }, num int) @commission {
  // send invalid value
  send(sendCh, num + 1)
}

proc RecvProc(recvCh channel { int }) {
  var variable int
  for {
    recv(recvCh, &variable) @omission
  }
}

proc SendProc(sendCh channel { int }) {
  for {
    send(sendCh, 1) @commission
  }
}
```

sendやrecvに対して非決定的に新たな遷移を加えられるfault markerを定義可能にする。  
必要に応じてsendやrecvはexpressionに変える。  
(process, channel, variableに起きる障害は第一要件から外す)

## マイルストン

1. statementの障害遷移を定義する構文を追加する。statementの還元規則でfault markerを追加可能にする。
2. recvやsend内でのchannel操作を再現するstatementを追加し、障害定義の幅を広げる。
3. 拡張した構文で追加した障害を用い、リードソロモン符号やByzantine Fault Tolerantなアルゴリズムを検査する。
4. 変数用の障害定義構文を導入する。
5. チャネルの挙動を再現するstatementを追加する。チャネル用の障害定義構文を導入する。
6. process以外の既存の障害をSandal実装に変える。
7. processの任意の障害を記述可能にし、process障害もSandal実装に変える。
