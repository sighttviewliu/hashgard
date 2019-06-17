package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/hashgard/hashgard/x/exchange"
)

// nolint
const (
	RestOrderId = "id"
	RestAddress = "address"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc(fmt.Sprintf("/exchange/order/{%s}", RestOrderId), queryOrderHandlerFn(cdc, cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/exchange/list/{%s}", RestAddress), queryOrdersByAddrHandlerFn(cdc, cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/exchange/frozen/{%s}", RestAddress), queryFrozenFundByAddrHandlerFn(cdc, cliCtx)).Methods("GET")

	r.HandleFunc("/exchange/make", postOrderMakeHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/exchange/take/{%s}", RestOrderId), postTakeHandlerFn(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/exchange/cancel/{%s}", RestOrderId), postCancelHandlerFn(cdc, cliCtx)).Methods("POST")
}

func queryOrderHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strOrderId := vars[RestOrderId]

		if len(strOrderId) == 0 {
			err := errors.New("id required but not specified")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		id, ok := rest.ParseUint64OrReturnBadRequest(w, strOrderId)
		if !ok {
			return
		}

		params := exchange.NewQueryOrderParams(id)

		bz, err := cdc.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", exchange.StoreKey, exchange.QueryOrder), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

func queryOrdersByAddrHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strAddress := vars[RestAddress]

		address, err := sdk.AccAddressFromBech32(strAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		params := exchange.NewQueryOrdersParams(address)

		bz, err := cdc.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", exchange.StoreKey, exchange.QueryAllOrdersByAddress), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

func queryFrozenFundByAddrHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strAddress := vars[RestAddress]

		address, err := sdk.AccAddressFromBech32(strAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		params := exchange.NewQueryFrozenFundParams(address)

		bz, err := cdc.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", exchange.StoreKey, exchange.QueryFrozenFund), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

type PostOrderReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Supply  sdk.Coin     `json:"supply"`
	Target  sdk.Coin     `json:"target"`
}

type TakeReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Amount  sdk.Coin     `json:"amount"`
}

type CancelReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
}

func postOrderMakeHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostOrderReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddress, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			return
		}

		// create the message
		msg := exchange.NewMsgMake(fromAddress, req.Supply, req.Target)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postTakeHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strOrderId := vars[RestOrderId]

		if len(strOrderId) == 0 {
			err := errors.New("id required but not specified")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		id, ok := rest.ParseUint64OrReturnBadRequest(w, strOrderId)
		if !ok {
			return
		}

		var req TakeReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddress, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			return
		}

		// create the message
		msg := exchange.NewMsgTake(id, fromAddress, req.Amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postCancelHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strOrderId := vars[RestOrderId]

		if len(strOrderId) == 0 {
			err := errors.New("id required but not specified")
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		id, ok := rest.ParseUint64OrReturnBadRequest(w, strOrderId)
		if !ok {
			return
		}

		var req CancelReq
		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddress, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			return
		}

		// create the message
		msg := exchange.NewMsgCancel(id, fromAddress)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
