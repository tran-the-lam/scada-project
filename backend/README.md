Phần này xây dựng API để tương tác với các hệ thống bên ngoài. Có 2 API được xây dựng là api cập nhật dữ liệu và api lấy nội dung theo key

#### Cài đặt:

Bước 1: Cài đặt Hyperledger Fabric
```
./network.sh up createChannel
```

Bước 2: Deploy chaincode
```
cd ../scada-project
../test-network/network.sh deployCC -ccn basic -ccp $PWD/smart-contract/chaincode -ccl go
```

Bước 3: Run project
```
cd backend
make run_api
```

[Postman API Documenter](https://documenter.getpostman.com/view/6827911/2s9Xy3sBBT)

#### Api cập nhật dữ liệu
```
curl --location '{{url}}/scadas/state' \
--header 'Content-Type: application/json' \
--data '{
    "function": "CreateKey",
    "args": ["k", "13"],
    "chain_code_id": "basic",
    "channel_id": "mychannel"
}'
```
| Các thuộc tính body | Mô tả |
|------------|-------|
| function | Tên hàm thực thi trong smart contract|
| args     | Các tham số của hàm thực thi |
| chain_code_id | Mã chain code |
| channel_id | Mã channel |

#### Api đọc dữ liệu
```
curl --location '{{url}}/scadas/state?chain_code_id=basic&channel_id=mychannel&function=QueryKey&args=k'
```
| Các thuộc tính query | Mô tả |
|------------|-------|
| function | Tên hàm thực thi trong smart contract|
| args     | Các tham số của hàm thực thi |
| chain_code_id | Mã chain code |
| channel_id | Mã channel |


#### Chạy test

```
make test
```
Kết qủa:
```
go clean -testcache && go test ./...
?       backend/cmd     [no test files]
?       backend/config  [no test files]
ok      backend/pkg/utils       0.165s
ok      backend/service 0.326s
```

#### Chạy test coverage
```
make test-cover
```

Kết quả:
```
go test -cover ./...
?       backend/cmd     [no test files]
?       backend/config  [no test files]
ok      backend/pkg/utils       0.290s  coverage: 100.0% of statements
ok      backend/service (cached)        coverage: 76.6% of statements
```
