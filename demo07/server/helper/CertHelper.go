package helper

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"google.golang.org/grpc/credentials"
)

//获取服务器端证书配置
func GetServerCreds() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	return creds
}

//获取客户端证书配置，这是给httpserver.go用的
/*当用户浏览器去访问http的时候，httpserver主动请求grpc，因此httpserver也需要向grpc提供客户端证书*/
func GetClientCreds() credentials.TransportCredentials {

	cert, _ := tls.LoadX509KeyPair("demo07/client/cert/client.pem", "demo07/client/cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("demo07/client/cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //服务端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	return creds
}
