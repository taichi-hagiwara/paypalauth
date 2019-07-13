package paypalauth

// Endpoint は、paypal API エンドポイントを表す。
type Endpoint struct {
	RootURL string
	AuthURL string
}

// SandboxEndpoint は、サンドボックス環境のエンドポイント。
var SandboxEndpoint = Endpoint{
	RootURL: "https://api.sandbox.paypal.com",
	AuthURL: "https://www.sandbox.paypal.com/connect",
}

// LiveEndpoint は、本番環境のエンドポイント。
var LiveEndpoint = Endpoint{
	RootURL: "https://api.paypal.com",
	AuthURL: "https://www.paypal.com/connect",
}
