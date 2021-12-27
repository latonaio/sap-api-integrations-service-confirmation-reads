# sap-api-integrations-service-confirmation-reads
sap-api-integrations-service-confirmation-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で サービス確認データを取得するマイクロサービスです。    
sap-api-integrations-service-confirmation-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-service-confirmation-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_SERVICE_CONFIRMATION_SRV_0001/overview  

## 動作環境  
sap-api-integrations-service-confirmation-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-service-confirmation-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-service-confirmation-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SERVICE_CONFIRMATION_SRV_0001/overview   
* APIサービス名(=baseURL): API_SERVICE_CONFIRMATION_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-service-confirmation-reads には、次の API をコールするためのリソースが含まれています。  

* A_ServiceConfirmation（サービス確認 - ヘッダ）※サービス確認の詳細データを取得するために、ToPersonResponsible、ToReferenceObject、ToItem、ToItemPricingElement、と合わせて利用されます。
* ToPersonResponsible（サービス確認 - 責任者）
* ToReferenceObject（サービス確認 - 参照対象）
* ToItem（サービス確認 - 明細）
* ToItemPricingElement（サービス確認 - 明細価格条件）

## API への 値入力条件 の 初期値
sap-api-integrations-service-confirmation-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ServiceConfirmation.ServiceConfirmation（サービス確認）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。

```
	"api_schema": "sap.s4.beh.serviceconfirmation.v1.ServiceConfirmation.Created.v1",
	"accepter": ["Header"],
	"service_confirmation": "5000000000",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "sap.s4.beh.serviceconfirmation.v1.ServiceConfirmation.Created.v1",
	"accepter": ["All"],
	"service_confirmation": "5000000000",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetServiceConfirmation(serviceConfirmation string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(serviceConfirmation)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP サービス確認 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ServiceConfirmation" ～ "to_Item" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-service-confirmation-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-service-confirmation-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"ServiceConfirmation": "5000000000",
			"ServiceConfirmationType": "RPC1",
			"ServiceConfirmationDescription": "Printer Professional IHR1911",
			"ServiceObjectType": "BUS2000117",
			"Language": "EN",
			"ServiceDocumentPriority": "0",
			"RequestedServiceStartDateTime": "/Date(1616976000000+0000)/",
			"RequestedServiceEndDateTime": "/Date(1617235200000+0000)/",
			"PurchaseOrderByCustomer": "",
			"CustomerPurchaseOrderDate": "",
			"ServiceConfirmationIsCompleted": "X",
			"ServiceConfirmationIsCanceled": false,
			"SalesOrganization": "1710",
			"DistributionChannel": "10",
			"Division": "00",
			"SalesOffice": "",
			"SalesGroup": "",
			"SoldToParty": "17100001",
			"ShipToParty": "17100001",
			"BillToParty": "17100001",
			"PayerParty": "17100001",
			"ContactPerson": "",
			"ReferenceServiceOrder": "4000000000",
			"ServiceConfirmationIsFinal": "N",
			"TransactionCurrency": "USD",
			"RespyMgmtServiceTeam": "50006529",
			"RespyMgmtServiceTeamName": "REPAIR_TEAM_US",
			"ServiceOrganization": "",
			"to_PersonResponsible": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SERVICE_CONFIRMATION_SRV/A_ServiceConfirmation('5000000000')/to_PersonResponsible",
			"to_ReferenceObject": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SERVICE_CONFIRMATION_SRV/A_ServiceConfirmation('5000000000')/to_ReferenceObject",
			"to_Item": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SERVICE_CONFIRMATION_SRV/A_ServiceConfirmation('5000000000')/to_Item"
		}
	],
	"time": "2021-12-27T19:57:26.247263+09:00"
}
```

