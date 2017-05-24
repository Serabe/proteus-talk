package gomad

import (
	"golang.org/x/net/context"
)

type proteustalkServiceServer struct {
}

func NewProteustalkServiceServer() *proteustalkServiceServer {
	return &proteustalkServiceServer{}
}
func (s *proteustalkServiceServer) AddIngredient(ctx context.Context, in *AddIngredientRequest) (result *Pizza, err error) {
	result = new(Pizza)
	result = AddIngredient(in.Arg1, in.Arg2, in.Arg3)
	return
}
func (s *proteustalkServiceServer) NutritionalReport(ctx context.Context, in *Pizza) (result *NutritionalReportResponse, err error) {
	result = new(NutritionalReportResponse)
	result.Result1 = NutritionalReport(in)
	return
}
func (s *proteustalkServiceServer) RemoveIngredient(ctx context.Context, in *RemoveIngredientRequest) (result *Pizza, err error) {
	result = new(Pizza)
	result, err = RemoveIngredient(in.Arg1, in.Arg2, in.Arg3)
	return
}
