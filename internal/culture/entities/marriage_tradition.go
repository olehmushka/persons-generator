package entities

type MarriageTradition struct {
	Kind          string `json:"kind"`
	Bastardry     string `json:"bastardry"`
	Consanguinity string `json:"consanguinity"`
	Divorce       string `json:"divorce"`
}
