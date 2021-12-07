You will avoid commands below defining every time in services.

resp := Response{

   Code:    responseCode,

   Message: msg,

   Data:    data,

}

http.ResponseWriter.WriteHeader(http.StatusCode)

http.ResponseWriter.Header.Set("Content-Type", "application/json")

err := json.NewEncoder(w).Encode(resp)

if err != nil {

   fmt.Println(err)

   panic(err)

}



Example:

func AddDocument(w http.ResponseWriter, r *http.Request){
  var document <YOur DOcument Model>
  err := json.NewDecoder(r.Body).Decode(&document)
  responseData := <DocumentQueryToDb>
  
  //201 - http.StatusCode- Created
  leo.ResponseWriter(w, 201, "Message you want to print", responseData)
}



