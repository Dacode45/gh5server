package main
import (
  "net/http/httptest"
  "net/http"
  "net/url"
  "testing"
  "fmt"
)

func TestHandlers(t *testing.T){
  fmt.Println("running test from gh5server")
  int currentId
  tests := []struct{
    Description string
    Handler func(http.ResponseWriter, *http.Request),
    Path string
    Method string
    Params url.Values
    Status int
    Payload []byte
    Match map[string]bool
  }{
  {
    {
      Description: "Get Index Page",
      Handler: Index,
      Method: "GET"
      Path: "/",
      Status: http.StatusOK,
    },
    {
      Description: "Get Court Index Page",
      Handler: CourtIndex,
      Method: "GET",
      Path:"/courts",
      Status: http.StatusOK,
    },
    {
      Description: "Create Invalid Court",
      Handler: CourtCreate,
      Method: "POST",
      Path:"/courts",
      Status: 422,
    }
  }
}

for _, test := range tests {
  record := httptest.NewRecord()
  var req *Request
  switch test.Method{
    case "GET":
      req = &http.Request{
        Method: test.Method,
        URL: &url.URL{Path: test.Path},
        Form: test.Params,
      }
    default:
      req, err = http.NewRequest(test.Method, &url.URL{Path: test.Path}, bytes.NewBuffer(test.Payload))
  }

  test.Handler(record, req)
  fmt.Println("Description : ", test.Description)
	fmt.Println("StatusCode : ", record.Code)
  if got, want := record.Code, test.Status; got != want{
    t.Errorf("%s: response code = %d, want %d", test.Description, got, want)
  }
  var objmap map[string]*json.RawMessage
  if err := json.Unmarshal(record.Body, &objmap); err != nil{
    t.Errorf("%s: response body not valid json %v", test.Description, record.Body)
  }
  for key, value := range test.Match{
    var str string
    if err := json.Unmarshal(*objmap[key], &str); err != nil{
      t.Errorf("%s: response obj does not contain key %v", test.Description, key)
    }
    if value != str {
      t.Errorf("%s: unexpected output. Wanted %v, got %v", value, str)
    }
    if key == "id" {
      currentId = value
    }

  }



}
}
