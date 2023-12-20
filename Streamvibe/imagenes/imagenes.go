package imagenes

import (
	"fmt"
	"io"
	"net/http"
)

func ShowImage(apiKey string) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/network/network_id/images?api_key=%s", apiKey)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
