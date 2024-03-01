package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"k8s.io/api/admission/v1beta1"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not read request body: %v", err), http.StatusInternalServerError)
		return
	}

	// 反序列化AdmissionReview请求
	var admissionReviewReq v1beta1.AdmissionReview
	if err := json.Unmarshal(body, &admissionReviewReq); err != nil {
		http.Error(w, fmt.Sprintf("could not unmarshal request: %v", err), http.StatusInternalServerError)
		return
	}

	// 准备AdmissionReview响应
	admissionReviewResp := v1beta1.AdmissionReview{
		Response: &v1beta1.AdmissionResponse{
			UID:     admissionReviewReq.Request.UID,
			Allowed: true, // 默认允许所有请求
		},
	}

	// 序列化AdmissionReview响应
	respBytes, err := json.Marshal(admissionReviewResp)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not marshal response: %v", err), http.StatusInternalServerError)
		return
	}

	// 返回响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}

func main() {
	http.HandleFunc("/validate", handleWebhook)
	http.ListenAndServeTLS(":8443", "/certs/tls.crt", "/certs/tls.key", nil)
}
