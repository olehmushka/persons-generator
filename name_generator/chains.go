package name_generator

var Chains []Chain

func UpdateChains(base int, nameBases [][]string) {
	Chains = UpdateChain(base, nameBases, Chains)
}
