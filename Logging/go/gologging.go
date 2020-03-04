// https://medium.com/google-cloud/combining-correlated-log-lines-in-google-stackdriver-dd23284aeb29

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/logging"
	"golang.org/x/net/context"
	mrpb "google.golang.org/genproto/googleapis/api/monitoredres"

	logpb "google.golang.org/genproto/googleapis/logging/v2"
)

var (
	client *logging.Client
)

const (
	httpport      = ":8080"
	parentLogName = "structured-log"
	childLogName  = "child-log"
	appID         = "529837574568057089"
	appProducer   = "MyDivision.MyBigCompany.com"
	projectId     = "ca-lutao-test"
)

func fronthandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Main Handler")
	start := time.Now()

	s := r.Header.Get("X-Cloud-Trace-Context")
	log.Printf("X-Cloud-Trace-Context: %s", s)

	parentLogger := client.Logger(
		parentLogName,
		logging.CommonLabels(map[string]string{
			"commonLabelParent": "commonLabelParentValue",
		}),
		logging.CommonResource(&mrpb.MonitoredResource{
			Labels: map[string]string{
				"project_id":  projectId,
				"instance_id": appID,
				"zone":        "us-central1-a",
			},
			Type: "gce_instance",
		}),
	)

	childLogger := client.Logger(
		childLogName,
		logging.CommonLabels(map[string]string{
			"commonChildParent": "commonLabelChildValue",
		}),
		logging.CommonResource(&mrpb.MonitoredResource{
			Labels: map[string]string{
				"project_id":  projectId,
				"zone":        "us-central1-a",
				"instance_id": appID,
			},
			Type: "gce_instance",
		}),
	)

	ent2 := logging.Entry{
		Timestamp: time.Now(),
		Severity:  logging.Info,
		Trace:     r.Header.Get("X-Cloud-Trace-Context"),
		Payload:   "ChildTextPayload Child Log message",
	}
	childLogger.Log(ent2)

	responseString := "hello world"

	ent := logging.Entry{
		HTTPRequest: &logging.HTTPRequest{
			Request:      r,
			Latency:      time.Now().Sub(start),
			ResponseSize: int64(len(responseString)),
		},
		Timestamp: time.Now(),
		Severity:  logging.Info,
		Trace:     r.Header.Get("X-Cloud-Trace-Context"),
		Operation: &logpb.LogEntryOperation{
			Id:       appID,
			Producer: appProducer,
		},
		Payload: "ParentTextPayload Parent Log message",
	}
	parentLogger.Log(ent)
	parentLogger.Flush()
	fmt.Fprint(w, responseString)

}

func healthhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {

	ctx := context.Background()
	client, _ = logging.NewClient(ctx, projectId)

	http.HandleFunc("/", fronthandler)
	http.HandleFunc("/_ah/health", healthhandler)

	log.Printf("Starting server on: %v", httpport)
	srv := &http.Server{
		Addr: httpport,
	}

	srv.ListenAndServe()
}
