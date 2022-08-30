package word_generator

var Chains map[string]Chain

func UpdateChains(base string, nameBases map[string][]string) (err error) {
	Chains, err = UpdateChain(base, nameBases, Chains)
	return
}
