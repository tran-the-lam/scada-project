Phần này xây dựng API để tương tác với các hệ thống bên ngoài. Có 2 API được xây dựng là api cập nhật dữ liệu và api lấy nội dung theo key

#### Clean data
```
cd ../../test-network
./network.sh down
```

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

==> All step:
```
cd ../../test-network
./network.sh down

./network.sh up createChannel

cd ../scada-project
../test-network/network.sh deployCC -ccn basic -ccp $PWD/smart-contract/chaincode -ccl go

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

TODO
Chuẩn bị API để ghép vào giao diện (thứ 3 phải xong)
[x] API get all event, save key allEvent => Call function API get history by Key
[x] API search => save key sensor_id, parameter => Call function get history by Key
[x] API update my password
[x] ADD user, 
remove_user,
[x] reset password

Chức năng chung:
- Xem lại các chỉ số giám sát được gửi từ cảm biến
- Màn profile
Chức năng admin:
- Thêm, reset mật khẩu cho user, tạm thời bỏ chức năng xóa user hoặc là có thể làm lưu danh sách xóa vào 1 key rồi tiến hành merge 2 mảng từ 'allUser' và 'deleteUser' => Cần nghĩ thêm


Tạo script để tạo dữ liệu mẫu

Luận văn sẽ nói về những vấn đề gì
+) 
+) Những phần mà ứng dụng có thể cải tiến, như là kết hợp với 1 vài cơ sở dữ liệu nữa để làm nhiệm vụ search, ...
+) Có thể xóa những dữ liệu không dùng bằng phương thức DelState => nhưng trước tiên cần lưu chúng ở một nơi nào đó an toàn
