package clb 

type LoadBalancerType int

const (
	Random     LoadBalancerType = iota
	RoundRobin LoadBalancerType = iota
)

type CacheType int

const (
	None CacheType = iota
	Ttl  CacheType = iota
)