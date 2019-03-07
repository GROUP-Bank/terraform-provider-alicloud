package cr

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

// DeleteImage invokes the cr.DeleteImage API synchronously
// api document: https://help.aliyun.com/api/cr/deleteimage.html
func (client *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
	response = CreateDeleteImageResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteImageWithChan invokes the cr.DeleteImage API asynchronously
// api document: https://help.aliyun.com/api/cr/deleteimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteImageWithChan(request *DeleteImageRequest) (<-chan *DeleteImageResponse, <-chan error) {
	responseChan := make(chan *DeleteImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteImage(request)
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

// DeleteImageWithCallback invokes the cr.DeleteImage API asynchronously
// api document: https://help.aliyun.com/api/cr/deleteimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteImageWithCallback(request *DeleteImageRequest, callback func(response *DeleteImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteImageResponse
		var err error
		defer close(result)
		response, err = client.DeleteImage(request)
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

// DeleteImageRequest is the request struct for api DeleteImage
type DeleteImageRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	Tag           string `position:"Path" name:"Tag"`
}

// DeleteImageResponse is the response struct for api DeleteImage
type DeleteImageResponse struct {
	*responses.BaseResponse
}

// CreateDeleteImageRequest creates a request to invoke DeleteImage API
func CreateDeleteImageRequest() (request *DeleteImageRequest) {
	request = &DeleteImageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "DeleteImage", "/repos/[RepoNamespace]/[RepoName]/tags/[Tag]", "acr", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteImageResponse creates a response to parse from DeleteImage response
func CreateDeleteImageResponse() (response *DeleteImageResponse) {
	response = &DeleteImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
