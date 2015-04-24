package grengine

type Flag struct {
	
}

type Request struct {
	RequestId string								`json:"id"`
	Expression string								`json:"e"`
	Engines []string								`json:"es"`
	Flags []string									`json:"f"`
	Payload interface{}							`json:"p"`
}

type Response struct {
	ResponseCode int								`json:"c"`
	Payload map[string]interface{}	`json:"p"`
}

type Engine interface{

	// Unique engine code
	GetEngineCode() string
	
	// Handles a request and prepare a response object
	Handle(req *Request, res *Response)
	
	// Returns optional javascript content to be loaded on client's browser
	// Refer to sample.js for basic example of an engine's script
	GetScript() string
}