Phần này xây dựng API để tương tác với các hệ thống bên ngoài. Có 2 API được xây dựng là api cập nhật dữ liệu và api lấy nội dung theo key

##### Api cập nhật dữ liệu
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

##### Api đọc dữ liệu
```
curl --location '{{url}}/scadas/state?chain_code_id=basic&channel_id=mychannel&function=QueryKey&args=k'
```
| Các thuộc tính query | Mô tả |
|------------|-------|
| function | Tên hàm thực thi trong smart contract|
| args     | Các tham số của hàm thực thi |
| chain_code_id | Mã chain code |
| channel_id | Mã channel |