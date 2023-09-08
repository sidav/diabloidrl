package static

type ItemStatsInterface interface {
	addEgo(*ego)
	GetName() string
	GetEgos() []*ego
	GetAffixDescriptions() []string
	addAffixDescription(string, ...interface{})
}
