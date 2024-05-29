package rest

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/pircuser61/media_main_task/config"
	"github.com/pircuser61/media_main_task/internal/exchanges"
)

func rootHandler(log *slog.Logger, rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	type Req struct {
		Amount    int
		Banknotes []int
	}
	var data Req
	err := decoder.Decode(&data)
	if err != nil {
		log.Info("bad request")
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(struct{ Error string }{Error: err.Error()})
		return
	}
	log.Debug("request", slog.Any("body", data))

	result, err := exchanges.GetExchages(data.Amount, data.Banknotes)
	if err != nil {
		log.Debug("result", slog.String("error", err.Error()))
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(struct{ Error string }{Error: err.Error()})
		return
	}
	log.Debug("result", slog.Any("data", result))

	type Resp struct {
		Exchanges [][]int `json:"exchanges"`
	}
	var resp Resp
	/*
		можно было бы в GetExchanges сразу формировать ответ как в json
		но мне формат не понравился, не хочтся подстраиваться под ручку
	*/
	for _, resultRow := range result {
		respRow := make([]int, 0, len(resultRow))
		for _, item := range resultRow {
			for count := 0; count < item.Count; count++ {
				respRow = append(respRow, item.Val)
			}
		}
		resp.Exchanges = append(resp.Exchanges, respRow)
	}

	//time.Sleep(time.Second * 30)
	err = json.NewEncoder(rw).Encode(resp)
	if err != nil {
		log.Error("write response error:", slog.String("msg", err.Error()))
	}
	log.Debug("response", slog.Any("body", resp))
}

func withLog(log *slog.Logger, f func(log *slog.Logger, rw http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) { f(log, rw, r) }
}

func RunSv(ctx context.Context, log *slog.Logger) {
	http.HandleFunc("/", withLog(log, rootHandler))
	srv := http.Server{Addr: config.GetAddr()}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Error(err.Error())
		}
		log.Info("http server stopped listening")
	}()

	<-ctx.Done()
	log.Info("Stopping http server...")
	if err := srv.Shutdown(context.TODO()); err != nil {
		log.Error(err.Error())
	}
	log.Info("http server stopped")
}
