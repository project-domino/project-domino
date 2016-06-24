package vars

// Vars stores variables for use in templates.
type Vars map[string]interface{}

// New returns an empty Vars.
func New(others ...Vars) Vars {
	vars := make(Vars)
	for _, other := range others {
		vars.Merge(other)
	}
	return vars
}

// Merge merges in another Vars object and returns itself.
func (vars Vars) Merge(other Vars) Vars {
	for k, v := range other {
		vars[k] = v
	}
	return vars
}

// Set sets a variable in the Vars and returns itself.
func (vars Vars) Set(key string, value interface{}) Vars {
	vars[key] = value
	return vars
}
