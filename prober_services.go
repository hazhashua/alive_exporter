package main

// import parseService "."

//    aPIV1Services, err := UnmarshalAPIV1Services(bytes)
//    bytes, err = aPIV1Services.Marshal()

func getServices(bytes []byte) {

	// 获取k8s service的返回内容

	UnmarshalAPIV1Services([]byte("hello go"))

}
