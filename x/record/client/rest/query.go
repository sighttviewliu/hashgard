package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/hashgard/hashgard/x/record/params"
	"github.com/hashgard/hashgard/x/record/types"

	"github.com/hashgard/hashgard/x/record/client/queriers"
	recordutils "github.com/hashgard/hashgard/x/record/utils"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc(fmt.Sprintf("/%s/%s/{%s}", types.QuerierRoute, types.QueryRecord, Hash), queryRecordHandlerFn(cdc, cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/%s", types.QuerierRoute, types.QueryRecords), queryRecordsHandlerFn(cdc, cliCtx)).Methods("GET")
}
func queryRecordHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hash := vars[Hash]
		if err := recordutils.CheckRecordHash(hash); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		res, err := queriers.QueryRecord(hash, cliCtx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}
func queryRecordsHandlerFn(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sender, err := sdk.AccAddressFromBech32(r.URL.Query().Get(Sender))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		startId := r.URL.Query().Get(StartRecordID)
		if len(startId) > 0 {
			if  err := recordutils.CheckRecordId(startId); err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}
		recordQueryParams := params.RecordQueryParams{
			StartRecordId: 	startId,
			Sender: 		sender,
			Limit:        	30,
		}
		strNumLimit := r.URL.Query().Get(Limit)
		if len(strNumLimit) > 0 {
			limit, err := strconv.Atoi(r.URL.Query().Get(Limit))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			recordQueryParams.Limit = limit
		}

		res, err := queriers.QueryRecords(recordQueryParams, cdc, cliCtx)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}
