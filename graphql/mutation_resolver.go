package main

import "context"

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) createAaccount(ctx context.Context, in AccountInput) (*Account, error) {

}

func (r *mutationResolver) createdProduct(ctx context.Context, in ProductInput) (*Product, error) {

}

func (r *mutationResolver) createOrder(ctx context.Context, in OrderInput) (*Order, error) {

}
