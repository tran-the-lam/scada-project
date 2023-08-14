package config

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"backend/pkg/utils"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
)

type OrgSetup struct {
	PORT         string
	OrgName      string
	CryptoPath   string
	CertPath     string
	KeyPath      string
	TLSCertPath  string
	PeerEndpoint string
	GatewayPeer  string
	MSPID        string
	Gateway      client.Gateway
}

var (
	instance *OrgSetup = nil
	once     sync.Once
)

func getEnv(key string, fallback interface{}) interface{} {
	var rValue interface{}
	value, exists := os.LookupEnv(key)
	if !exists {
		rValue = fallback
	} else {
		rValue = value
	}
	return rValue
}

func InitConfig() *OrgSetup {
	log.Println("Initializing connection")
	cryptoPath := "/Users/user/Documents/Master/LuanVan/Project/fabric-install/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com"
	once.Do(
		func() {
			if err := godotenv.Load(utils.ProviderPath(".env")); err != nil {
				log.Fatal("Error loading .env file")
			}

			instance = &OrgSetup{
				PORT:         getEnv("PORT", "").(string),
				OrgName:      "Org1",
				MSPID:        "Org1MSP",
				CertPath:     cryptoPath + "/users/User1@org1.example.com/msp/signcerts/User1@org1.example.com-cert.pem",
				KeyPath:      cryptoPath + "/users/User1@org1.example.com/msp/keystore/",
				TLSCertPath:  cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt",
				PeerEndpoint: "localhost:7051",
				GatewayPeer:  "peer0.org1.example.com",
			}
		},
	)

	clientConnection := instance.newGrpcConnection()
	id := instance.newIdentity()
	sign := instance.newSign()

	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	instance.Gateway = *gateway
	log.Println("Initialization complete")

	return instance
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func (setup *OrgSetup) newSign() identity.Sign {
	files, err := ioutil.ReadDir(setup.KeyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(setup.KeyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func (setup *OrgSetup) newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(setup.CertPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(setup.MSPID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func (setup *OrgSetup) newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(setup.TLSCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, setup.GatewayPeer)

	connection, err := grpc.Dial(setup.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}
