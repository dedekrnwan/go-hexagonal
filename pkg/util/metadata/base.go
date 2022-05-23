package metadata

type key int32

// preventing context key collisions, i use an int for the sake of simplicity
// alternatively, we could use private struct for each key if we want to optimize runtime memory consumption
const (
	metadataUser key = iota
	metadataAdmin
)
