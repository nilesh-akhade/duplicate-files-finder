package checksum

type ChecksumCalc interface {
	Calculate(path string) (string, error)
}
