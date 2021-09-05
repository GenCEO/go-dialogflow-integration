
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)