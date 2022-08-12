package word_generator

var Chains map[string]Chain

func UpdateChains(base string, nameBases map[string][]string) {
	Chains = UpdateChain(base, nameBases, Chains)
}
