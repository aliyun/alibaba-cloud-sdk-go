[English](./README.md) | [简体中文](./README-CN.md) | 日本語

<p align="center">
<a href=" https://www.alibabacloud.com"><img src="https://aliyunsdk-pages.alicdn.com/icons/Aliyun.svg"></a>
</p>

<h1 align="center">Alibaba Cloud SDK for Go</h1>

[![Go](https://github.com/aliyun/alibaba-cloud-sdk-go/actions/workflows/go.yml/badge.svg)](https://github.com/aliyun/alibaba-cloud-sdk-go/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/aliyun/alibaba-cloud-sdk-go/graph/badge.svg?token=kHbylWc7aV)](https://codecov.io/gh/aliyun/alibaba-cloud-sdk-go)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go.svg?type=shield&issueType=license)](https://app.fossa.io/projects/git%2Bgithub.com%2Falibaba-cloud-sdk-go?ref=badge_shield&issueType=license)

Alibaba Cloud SDK for Goを使用すると、複雑なプログラミングを行わずに、Elastic Compute Service（ECS）、Server Load Balancer（SLB）、CloudMonitorなどのAlibaba Cloudサービスにアクセスできます。
このドキュメントでは、[Alibaba Cloud SDK for Go][SDK]を取得して呼び出す方法を紹介します。

## トラブルシューティング

[Troubleshoot](https://troubleshoot.api.aliyun.com/?source=github_sdk)は、OpenAPI診断サービスを提供し、`RequestID`または`エラーメッセージ`を通じて、開発者が迅速に問題を特定し、解決策を提供します。

## オンラインデモ

[Alibaba Cloud OpenAPI Developer Portal][open-api-portal]は、クラウド製品のOpenAPIをオンラインで呼び出し、SDKのサンプルコードを動的に生成し、インターフェースを迅速に検索する機能を提供します。これにより、クラウドAPIの使用が大幅に簡素化されます。

## 必要条件

- システムが[必要条件][Requirements]を満たしていることを確認してください。たとえば、1.13.x以降のGo環境をインストールする必要があります。

## インストール

`go get`を使用してSDKをインストールします：

```sh
go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk
```

## クイックスタート

始める前に、Alibaba Cloudアカウントにサインアップし、[認証情報](https://usercenter.console.aliyun.com/#/manage/ak)を取得する必要があります。

### クライアントの作成

```go
package main

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk"

func main() {
  client, err := sdk.NewClientWithAccessKey("REGION_ID", "ACCESS_KEY_ID", "ACCESS_KEY_SECRET")
  if err != nil {
    // 例外処理
    panic(err)
  }
}
```

### ROAリクエスト

```go
package main

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

func main() {
  request := requests.NewCommonRequest()        // 共通リクエストを作成
  request.Method = "GET"                        // リクエストメソッドを設定
  request.Product = "CS"                        // 製品を指定
  request.Domain = "cs.aliyuncs.com"            // ホストを指定すると、ロケーションサービスは有効になりません。たとえば、認証タイプがBearer Tokenのサービスは指定する必要があります
  request.Version = "2015-12-15"                // 製品バージョンを指定
  request.PathPattern = "/clusters/[ClusterId]" // ROAスタイルのパスルールを指定
  request.Scheme = "https"                      // リクエストスキームを設定。デフォルトはhttp
  request.ApiName = "DescribeCluster"           // 製品インターフェースを指定
  request.QueryParams["ClusterId"] = "123456"   // パス内のパラメータに値を割り当て
  request.QueryParams["RegionId"] = "region_id" // リクエストされたregionIdを指定。指定しない場合は、クライアントのregionId、次にデフォルトのregionIdを使用
  request.TransToAcsRequest()                   // 共通リクエストをacsリクエストに変換。これはクライアントによって使用されます。
}
```

### RPCリクエスト

```go
package main

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

func main() {
  request := requests.NewCommonRequest()                // 共通リクエストを作成
  request.Method = "POST"                               // リクエストメソッドを設定
  request.Product = "Ecs"                               // 製品を指定
  request.Domain = "ecs.aliyuncs.com"                   // ホストを指定すると、ロケーションサービスは有効になりません。たとえば、認証タイプがBearer Tokenのサービスは指定する必要があります
  request.Version = "2014-05-26"                        // 製品バージョンを指定
  request.Scheme = "https"                              // リクエストスキームを設定。デフォルトはhttp
  request.ApiName = "CreateInstance"                    // 製品インターフェースを指定
  request.QueryParams["InstanceType"] = "ecs.g5.large"  // パス内のパラメータに値を割り当て
  request.QueryParams["RegionId"] = "region_id"         // リクエストされたregionIdを指定。指定しない場合は、クライアントのregionId、次にデフォルトのregionIdを使用
  request.TransToAcsRequest()                           // 共通リクエストをacsリクエストに変換。これはクライアントによって使用されます。
}
```

## ドキュメント

- [Requirements](docs/0-Requirements-EN.md)
- [Installation](docs/1-Installation-EN.md)
- [Client & Credentials](docs/2-Client-EN.md)
- [SSL Verify](docs/3-Verify-EN.md)
- [Proxy](docs/4-Proxy-EN.md)
- [Timeout](docs/5-Timeout-EN.md)
- [Debug](docs/6-Debug-EN.md)
- [Logger](docs/7-Logger-EN.md)
- [Concurrent](docs/8-Concurrent-EN.md)
- [Asynchronous Call](docs/9-Asynchronous-EN.md)
- [Package Management](docs/10-Package-Management-EN.md)
- [Endpoint](docs/11-Endpoint-EN.md)

## 問題

[Issueを開く][issue]、ガイドラインに準拠していない問題は直ちに閉じられる場合があります。

## 貢献

プルリクエストを作成する前に、[貢献ガイド](CONTRIBUTING.md)を必ずお読みください。

## 参考文献

- [Alibaba Cloud Regions & Endpoints][endpoints]
- [Alibaba Cloud OpenAPI Developer Portal][open-api-portal]
- [Go][go]
- [最新リリース][latest-release]

## ライセンス

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go?ref=badge_large)

[SDK]: https://github.com/aliyun/alibaba-cloud-sdk-go
[issue]: https://github.com/aliyun/alibaba-cloud-sdk-go/issues/new
[open-api-portal]: https://api.aliyun.com/
[latest-release]: https://github.com/aliyun/alibaba-cloud-sdk-go/releases
[go]: https://golang.org/dl/
[endpoints]: https://developer.aliyun.com/endpoints
[Requirements]: docs/0-Requirements-JP.md
