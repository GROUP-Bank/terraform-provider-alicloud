package r_kvstore

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// TransformToPrePaid invokes the r_kvstore.TransformToPrePaid API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/transformtoprepaid.html
func (client *Client) TransformToPrePaid(request *TransformToPrePaidRequest) (response *TransformToPrePaidResponse, err error) {
	response = CreateTransformToPrePaidResponse()
	err = client.DoAction(request, response)
	return
}

// TransformToPrePaidWithChan invokes the r_kvstore.TransformToPrePaid API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/transformtoprepaid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TransformToPrePaidWithChan(request *TransformToPrePaidRequest) (<-chan *TransformToPrePaidResponse, <-chan error) {
	responseChan := make(chan *TransformToPrePaidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TransformToPrePaid(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// TransformToPrePaidWithCallback invokes the r_kvstore.TransformToPrePaid API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/transformtoprepaid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TransformToPrePaidWithCallback(request *TransformToPrePaidRequest, callback func(response *TransformToPrePaidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TransformToPrePaidResponse
		var err error
		defer close(result)
		response, err = client.TransformToPrePaid(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// TransformToPrePaidRequest is the request struct for api TransformToPrePaid
type TransformToPrePaidRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	Period               requests.Integer `position:"Query" name:"Period"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	FromApp              string           `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// TransformToPrePaidResponse is the response struct for api TransformToPrePaid
type TransformToPrePaidResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
	EndTime   string `json:"EndTime" xml:"EndTime"`
}

// CreateTransformToPrePaidRequest creates a request to invoke TransformToPrePaid API
func CreateTransformToPrePaidRequest() (request *TransformToPrePaidRequest) {
	request = &TransformToPrePaidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "TransformToPrePaid", "redisa", "openAPI")
	return
}

// CreateTransformToPrePaidResponse creates a response to parse from TransformToPrePaid response
func CreateTransformToPrePaidResponse() (response *TransformToPrePaidResponse) {
	response = &TransformToPrePaidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
