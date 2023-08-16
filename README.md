Project ứng dụng công nghệ blockchain [Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/release-2.5/index.html) để tăng cường tính bảo mật cho hệ thống `Scada`. 

## Kiến thức nền tảng

1. [Scada](https://docs.google.com/document/d/10R_ofWSwNjWEZ7i4tid5dPGmwIhwETDxzFmLMU1RotE/edit?usp=sharing)

2. [Hyperledger Fabric](https://docs.google.com/document/d/1OgtoTqUcE656rH7cmA90lOfzGgrw3QyokLfaWRzNBBM/edit?usp=sharing)

## Cài đặt

Bước 1: Cài đặt Hyperledger Fabric
```
./network.sh up createChannel
```

Bước 2: Deploy chaincode
```
./network.sh deployCC -ccn basic -ccp /Users/user/Documents/Master/LuanVan/Project/fabric-install/fabric-samples/scada-project/smart-contract/chaincode -ccl go
```

Bước 3: Run project
```
cd backend
go run cmd/main.go
```

[Postman API](https://documenter.getpostman.com/view/6827911/2s9Xy3sBBT)