/*
Copyright 2017 The Fission Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/fission/fission/crd"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *API) RecorderApiList(w http.ResponseWriter, r*http.Request) {
	// TODO: Extract/set namespace

	log.Info("At the right place!")

	recorders, err := a.fissionClient.Recorders(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		log.Info("Couldn't obtain recorders")
		a.respondWithError(w, err)
		return
	}
	resp, err := json.Marshal(recorders.Items)
	if err != nil {
		log.Info("Couldn't unmarshall")
		a.respondWithError(w, err)
		return
	}
	a.respondWithSuccess(w, resp)
}

/*
func (a *API) MessageQueueTriggerApiList(w http.ResponseWriter, r *http.Request) {
	//mqType := r.FormValue("mqtype") // ignored for now
	ns := a.extractQueryParamFromRequest(r, "namespace")
	if len(ns) == 0 {
		ns = metav1.NamespaceAll
	}

	triggers, err := a.fissionClient.MessageQueueTriggers(ns).List(metav1.ListOptions{})
	if err != nil {
		a.respondWithError(w, err)
		return
	}
	resp, err := json.Marshal(triggers.Items)
	if err != nil {
		a.respondWithError(w, err)
		return
	}
	a.respondWithSuccess(w, resp)
}
*/

func (a *API) RecorderApiCreate(w http.ResponseWriter, r*http.Request) {
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//_, err := w.Write([]byte("At least you tried"))
	//return

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Info("In RecoderApiCreate, error with ioutil.ReadAll")	// TODO: Remove later
		a.respondWithError(w, err)
		return
	}

	var recorder crd.Recorder
	err = json.Unmarshal(body, &recorder)
	if err != nil {
		log.Info("In RecoderApiCreate, error with json.Unmarshal")
		a.respondWithError(w, err)
		return
	}

	// check if namespace exists, if not create it.
	/*
	err = a.createNsIfNotExists(mqTrigger.Metadata.Namespace)
	if err != nil {
		a.respondWithError(w, err)
		return
	}
	*/

	tnew, err := a.fissionClient.Recorders("default").Create(&recorder)
	if err != nil {
		a.respondWithError(w, err)
		return
	}

	resp, err := json.Marshal(tnew.Metadata)
	if err != nil {
		a.respondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	a.respondWithSuccess(w, resp)
}

func (a *API) RecorderApiGet(w http.ResponseWriter, r *http.Request) {

}
/*
func (a *API) MessageQueueTriggerApiGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["mqTrigger"]
	ns := a.extractQueryParamFromRequest(r, "namespace")
	if len(ns) == 0 {
		ns = metav1.NamespaceDefault
	}

	mqTrigger, err := a.fissionClient.MessageQueueTriggers(ns).Get(name)
	if err != nil {
		a.respondWithError(w, err)
		return
	}
	resp, err := json.Marshal(mqTrigger)
	if err != nil {
		a.respondWithError(w, err)
		return
	}
	a.respondWithSuccess(w, resp)
}
*/

func (a *API) RecorderApiUpdate(w http.ResponseWriter, r *http.Request) {

}
/*
func (a *API) MessageQueueTriggerApiUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["mqTrigger"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.respondWithError(w, err)
		return
	}

	var mqTrigger crd.MessageQueueTrigger
	err = json.Unmarshal(body, &mqTrigger)
	if err != nil {
		a.respondWithError(w, err)
		return
	}

	if name != mqTrigger.Metadata.Name {
		err = fission.MakeError(fission.ErrorInvalidArgument, "Message queue trigger name doesn't match URL")
		a.respondWithError(w, err)
		return
	}

	tnew, err := a.fissionClient.MessageQueueTriggers(mqTrigger.Metadata.Namespace).Update(&mqTrigger)
	if err != nil {
		a.respondWithError(w, err)
		return
	}

	resp, err := json.Marshal(tnew.Metadata)
	if err != nil {
		a.respondWithError(w, err)
		return
	}
	a.respondWithSuccess(w, resp)
}
*/

func (a *API) RecorderApiDelete(w http.ResponseWriter, r *http.Request) {

}

/*
func (a *API) MessageQueueTriggerApiDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["mqTrigger"]
	ns := a.extractQueryParamFromRequest(r, "namespace")
	if len(ns) == 0 {
		ns = metav1.NamespaceDefault
	}

	err := a.fissionClient.MessageQueueTriggers(ns).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		a.respondWithError(w, err)
		return
	}
	a.respondWithSuccess(w, []byte(""))
}
*/
