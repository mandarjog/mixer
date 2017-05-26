// GENERATED FILE (simulated)

package quota

type BaseInstance struct {
	// Tamplate and Name
	// are available in all instances
	Template *Template
	Name     string
}

type Instance struct {
	BaseInstance
	Dimensions map[string]interface{}
}

type Config struct {
	MaxAmount  int
	Expiration time.Duration
}
